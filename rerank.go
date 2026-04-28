package otari

import (
	"bytes"
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
)

// Rerank reranks documents by relevance to a query via the /v1/rerank endpoint.
// The response contains results sorted by relevance_score in descending order.
func (c *Client) Rerank(ctx context.Context, params RerankParams) (*RerankResponse, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if params.Query == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("query is required"))
	}
	if len(params.Documents) == 0 {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("at least one document is required"))
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshaling rerank request: %w", err)
	}

	reqURL := c.baseURL + "/v1/rerank"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating rerank request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, convertOpenAIError(err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleRerankErrorResponse(resp)
	}

	var result RerankResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding rerank response: %w", err)
	}
	return &result, nil
}

// handleRerankErrorResponse parses an HTTP error response from the /v1/rerank
// endpoint and returns a typed error.
func (c *Client) handleRerankErrorResponse(resp *http.Response) error {
	// ReadAll error is ignored: if reading fails, we fall back to an empty
	// body and still produce a typed error based on the status code.
	body, _ := io.ReadAll(resp.Body)

	parsed := parseRerankError(body)
	msg := parsed.message()

	switch resp.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		return newAuthenticationError(providerName, fmt.Errorf("%s", msg))
	case http.StatusNotFound:
		return newModelNotFoundError(providerName, fmt.Errorf("%s", msg))
	case http.StatusPaymentRequired:
		return newInsufficientFundsError(providerName, fmt.Errorf("%s", msg))
	case http.StatusTooManyRequests:
		return newRateLimitError(providerName, fmt.Errorf("%s", msg))
	case http.StatusBadGateway:
		return &UpstreamProviderError{OtariError: newOtariError(codeUpstreamProvider, providerName, fmt.Errorf("%s", msg), ErrUpstreamProvider)}
	case http.StatusGatewayTimeout:
		return &TimeoutError{OtariError: newOtariError(codeTimeout, providerName, fmt.Errorf("%s", msg), ErrTimeout)}
	default:
		return newProviderError(providerName, fmt.Errorf("%s", msg))
	}
}

// rerankError holds the parsed fields from a rerank error response.
type rerankError struct {
	Detail   string
	raw      string
	parseErr error
}

// message returns the best human-readable description of the error body.
func (e rerankError) message() string {
	if e.Detail != "" {
		return e.Detail
	}
	return e.raw
}

// parseRerankError parses a rerank error response. Non-JSON bodies
// are preserved via the raw field.
func parseRerankError(body []byte) rerankError {
	raw := string(body)
	var wrapper struct {
		Detail string `json:"detail"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return rerankError{raw: raw, parseErr: fmt.Errorf("unmarshal error body: %w", err)}
	}
	if wrapper.Detail == "" {
		return rerankError{raw: raw, parseErr: stderrors.New("error body missing \"detail\" field")}
	}
	return rerankError{Detail: wrapper.Detail, raw: raw}
}
