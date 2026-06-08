package otari

import (
	"context"
	"fmt"
	"net/http"
)

// Rerank reranks documents by relevance to a query via POST /v1/rerank.
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

	var result RerankResponse
	if err := c.doJSON(ctx, http.MethodPost, "/rerank", params, &result, ""); err != nil {
		return nil, err
	}
	return &result, nil
}
