package otari

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// Batch API path constants.
const (
	// batchesPath is the base path (relative to /v1) for batch operations.
	batchesPath = "/batches"

	// providerQueryParam is the query-string key identifying the upstream
	// provider for batch operations.
	providerQueryParam = "provider"
)

// CreateBatch creates a new batch job.
func (c *Client) CreateBatch(
	ctx context.Context,
	params CreateBatchParams,
) (*Batch, error) {
	var batch Batch
	if err := c.doJSON(ctx, http.MethodPost, batchesPath, params, &batch, ""); err != nil {
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
	if err := c.doJSON(ctx, http.MethodGet, path, nil, &batch, batchID); err != nil {
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
	if err := c.doJSON(ctx, http.MethodPost, path, nil, &batch, batchID); err != nil {
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

	var listResp struct {
		Data []Batch `json:"data"`
	}
	if err := c.doJSON(ctx, http.MethodGet, u.RequestURI(), nil, &listResp, ""); err != nil {
		return nil, err
	}
	return listResp.Data, nil
}

// RetrieveBatchResults retrieves the results of a completed batch job.
//
// Returns a *BatchNotCompleteError (matching ErrBatchNotComplete) when the
// gateway reports the batch is not yet complete (HTTP 409).
func (c *Client) RetrieveBatchResults(
	ctx context.Context,
	batchID string,
	provider string,
) (*BatchResult, error) {
	path := batchItemPath(batchID, "results", provider)
	var result BatchResult
	if err := c.doJSON(ctx, http.MethodGet, path, nil, &result, batchID); err != nil {
		return nil, err
	}
	return &result, nil
}

// batchItemPath builds a path for /batches/{id}[/action]?provider=X relative to
// the /v1 base. It escapes each path segment individually via url.URL.JoinPath
// so a batchID containing "/" or ".." is encoded as a single segment.
func batchItemPath(batchID, action, provider string) string {
	u := (&url.URL{Path: batchesPath}).JoinPath(batchID)
	if action != "" {
		u = u.JoinPath(action)
	}
	u.RawQuery = url.Values{providerQueryParam: {provider}}.Encode()
	return u.RequestURI()
}
