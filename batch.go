package otari

import (
	"bytes"
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Batch API path constants.
const (
	// batchesPath is the base path for batch operations on the gateway.
	batchesPath = "/v1/batches"

	// providerQueryParam is the query-string key identifying the upstream
	// provider for batch operations.
	providerQueryParam = "provider"
)

// CreateBatch creates a new batch job.
func (c *Client) CreateBatch(
	ctx context.Context,
	params CreateBatchParams,
) (*Batch, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, newInvalidRequestError(providerName, err)
	}

	var batch Batch
	if err := c.callBatchJSON(ctx, http.MethodPost, batchesPath, "", body, &batch); err != nil {
		return nil, err
	}

	return &batch, nil
}

// RetrieveBatch retrieves a batch job by ID.
func (c *Client) RetrieveBatch(
	ctx context.Context,
	batchID string,
	provider string,
) (*Batch, error) {
	path := batchItemPath(batchID, "", provider)

	var batch Batch
	if err := c.callBatchJSON(ctx, http.MethodGet, path, batchID, nil, &batch); err != nil {
		return nil, err
	}

	return &batch, nil
}

// CancelBatch cancels a batch job.
func (c *Client) CancelBatch(
	ctx context.Context,
	batchID string,
	provider string,
) (*Batch, error) {
	path := batchItemPath(batchID, "cancel", provider)

	var batch Batch
	if err := c.callBatchJSON(ctx, http.MethodPost, path, batchID, nil, &batch); err != nil {
		return nil, err
	}

	return &batch, nil
}

// ListBatches lists batch jobs for a provider.
func (c *Client) ListBatches(
	ctx context.Context,
	provider string,
	opts ListBatchesOptions,
) ([]Batch, error) {
	params := url.Values{providerQueryParam: {provider}}
	if opts.After != "" {
		params.Set("after", opts.After)
	}
	if opts.Limit != nil {
		params.Set("limit", strconv.Itoa(*opts.Limit))
	}

	u := url.URL{Path: batchesPath, RawQuery: params.Encode()}
	path := u.RequestURI()

	var listResp struct {
		Data []Batch `json:"data"`
	}
	if err := c.callBatchJSON(ctx, http.MethodGet, path, "", nil, &listResp); err != nil {
		return nil, err
	}

	return listResp.Data, nil
}

// RetrieveBatchResults retrieves the results of a completed batch job.
func (c *Client) RetrieveBatchResults(
	ctx context.Context,
	batchID string,
	provider string,
) (*BatchResult, error) {
	path := batchItemPath(batchID, "results", provider)

	var result BatchResult
	if err := c.callBatchJSON(ctx, http.MethodGet, path, batchID, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// batchItemPath builds a path for /v1/batches/{id}[/action]?provider=X.
// It uses url.URL.JoinPath which escapes each path segment individually,
// so a batchID containing "/" or ".." is encoded as a single segment rather
// than traversing into a sibling route.
func batchItemPath(batchID, action, provider string) string {
	u := (&url.URL{Path: batchesPath}).JoinPath(batchID)
	if action != "" {
		u = u.JoinPath(action)
	}
	u.RawQuery = url.Values{providerQueryParam: {provider}}.Encode()
	return u.RequestURI()
}

// callBatchJSON performs a batch HTTP request and decodes a JSON success
// response into out. Non-2xx responses are mapped to typed errors via
// handleBatchError. batchID is used to enrich 409 errors with the specific
// batch identifier the caller requested; pass "" when not applicable.
func (c *Client) callBatchJSON(
	ctx context.Context,
	method, path, batchID string,
	body []byte,
	out any,
) error {
	resp, err := c.doRequest(ctx, method, path, body)
	if err != nil {
		return err
	}
	defer closeBody(resp)

	if resp.StatusCode != http.StatusOK {
		return c.handleBatchError(resp, batchID)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return newProviderError(providerName, fmt.Errorf("failed to decode batch response: %w", err))
	}

	return nil
}

// closeBody closes an HTTP response body; intended for use in defer
// statements. The close error is deliberately discarded.
func closeBody(resp *http.Response) {
	_ = resp.Body.Close()
}

// doRequest sends an HTTP request to the otari API for batch operations.
func (c *Client) doRequest(
	ctx context.Context,
	method, path string,
	body []byte,
) (*http.Response, error) {
	fullURL := c.apiBase + path

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, newProviderError(providerName, err)
	}

	req.Header.Set(contentTypeHeader, contentTypeJSON)
	if c.apiKey != "" {
		if c.platformMode {
			req.Header.Set(authorizationHeader, bearerPrefix+c.apiKey)
		} else {
			req.Header.Set(apiKeyHeaderName, bearerPrefix+c.apiKey)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, newProviderError(providerName, err)
	}

	return resp, nil
}

// handleBatchError maps a non-2xx batch HTTP response to a typed error.
// batchID is the batch identifier the caller operated on (empty for
// create/list); it is used to enrich 409 errors without parsing
// free-text error messages.
//
// All callers of this function operate on /v1/batches, so a 404 is
// interpreted as "gateway does not expose the batch API" and is not
// mapped to ModelNotFoundError.
func (c *Client) handleBatchError(resp *http.Response, batchID string) error {
	// ReadAll error is ignored: if reading fails, we fall back to an empty
	// body and still produce a typed error based on status code.
	bodyBytes, _ := io.ReadAll(resp.Body)

	detail, parseErr := parseBatchError(bodyBytes)

	switch resp.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return newAuthenticationError(providerName,
			fmt.Errorf("unauthorized: %s", detail.message()))
	case http.StatusNotFound:
		return newProviderError(providerName,
			stderrors.New("this gateway does not support batch operations; upgrade your gateway"))
	case http.StatusConflict:
		// For 409 the Status field matters (it tells the caller why the
		// batch isn't ready). If the body wasn't valid JSON, wrap the
		// parse error so drift in the gateway's response format is
		// visible instead of silently returning an empty Status.
		if parseErr != nil {
			return newBatchNotCompleteError(batchID, "",
				fmt.Errorf("failed to parse batch error response: %w", parseErr))
		}
		return newBatchNotCompleteError(batchID, detail.Status, detail.toError())
	case http.StatusUnprocessableEntity:
		return newProviderError(providerName,
			fmt.Errorf("unprocessable request: %s", detail.message()))
	case http.StatusTooManyRequests:
		return newRateLimitError(providerName,
			fmt.Errorf("rate limit: %s", detail.message()))
	case http.StatusBadGateway:
		return newProviderError(providerName,
			fmt.Errorf("upstream provider error: %s", detail.message()))
	default:
		return newProviderError(providerName,
			fmt.Errorf("HTTP %d: %s", resp.StatusCode, detail.message()))
	}
}

// batchError is the structured JSON shape error responses may use.
// All fields are optional; callers must cope with arbitrary bodies.
type batchError struct {
	// Detail is the human-readable error message (FastAPI-style payload).
	Detail string `json:"detail"`

	// Status is the batch status when the gateway returns a 409 for
	// RetrieveBatchResults; empty when not applicable or not provided.
	Status string `json:"status"`

	// raw is the untouched response body, used as a last-resort message
	// when the body is not JSON or does not include Detail.
	raw string
}

// message returns the best human-readable description of the error body.
func (b batchError) message() string {
	if b.Detail != "" {
		return b.Detail
	}
	return b.raw
}

// toError wraps message() in an error value suitable for embedding in a
// typed error's cause chain, or nil when no detail is available.
func (b batchError) toError() error {
	msg := b.message()
	if msg == "" {
		return nil
	}
	return stderrors.New(msg)
}

// parseBatchError parses a gateway error response body. The parse error is
// returned so callers can decide per-status whether a JSON parse failure is
// tolerable or should be treated as a hard failure.
func parseBatchError(body []byte) (batchError, error) {
	out := batchError{raw: string(body)}
	err := json.Unmarshal(body, &out)
	return out, err
}
