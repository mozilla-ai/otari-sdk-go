package otari

import (
	"context"
	"fmt"
	"net/http"
)

// ImageGenerationParams are the parameters for POST /v1/images/generations.
// Extra carries any additional OpenAI-compatible image fields (n, size,
// quality, response_format, style, user, ...) forwarded verbatim.
type ImageGenerationParams struct {
	Model  string         `json:"model"`
	Prompt string         `json:"prompt"`
	Extra  map[string]any `json:"-"`
}

// imageGenerationBody builds the JSON request body for POST
// /v1/images/generations from the ergonomic ImageGenerationParams.
func imageGenerationBody(params ImageGenerationParams) map[string]any {
	body := structToMap(params)
	for k, v := range params.Extra {
		body[k] = v
	}
	return body
}

// ImageGeneration generates images from a text prompt via
// POST /v1/images/generations.
//
// The gateway returns an OpenAI-compatible image payload
// (a map with "created" and "data"). The generated core models this response as
// an opaque object, so the decoded JSON is returned as a map (matching Message
// and Response).
func (c *Client) ImageGeneration(
	ctx context.Context,
	params ImageGenerationParams,
) (map[string]any, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if params.Prompt == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("prompt is required"))
	}

	var out map[string]any
	if err := c.doJSON(ctx, http.MethodPost, "/images/generations", imageGenerationBody(params), &out, ""); err != nil {
		return nil, err
	}
	return out, nil
}
