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
	"strings"

	"github.com/openai/openai-go"
)

// Moderation runs a content moderation check via POST /v1/moderations.
//
// The openai-go SDK exposes a moderations resource, but we issue the HTTP
// call directly here to control the include_raw query parameter and the
// otari-specific error shape (notably the "does not support moderation"
// 400 that maps to *UnsupportedCapabilityError).
//
// Returns an *UnsupportedCapabilityError (matching ErrUnsupported) when
// the gateway reports that the chosen backend provider does not support
// moderation.
func (c *Client) Moderation(
	ctx context.Context,
	params ModerationParams,
) (*ModerationResponse, error) {
	u := &url.URL{Path: "/v1/moderations"}
	if params.IncludeRaw {
		u.RawQuery = url.Values{"include_raw": {"true"}}.Encode()
	}
	endpoint := c.baseURL + u.RequestURI()

	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("otari: marshal moderation params: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("otari: build moderation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.platformMode && c.platformToken != "" {
		// Platform mode: replicate the Bearer auth that the embedded
		// openai-go SDK adds for its own requests. Non-platform mode
		// relies on headerTransport, which is wired into c.httpClient.
		req.Header.Set("Authorization", bearerPrefix+c.platformToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, convertOpenAIError(err)
	}
	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("otari: read moderation response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, c.convertHTTPModerationError(resp.StatusCode, bodyBytes)
	}

	var out ModerationResponse
	if err := json.Unmarshal(bodyBytes, &out); err != nil {
		return nil, fmt.Errorf("otari: unmarshal moderation response: %w", err)
	}
	return &out, nil
}

// convertHTTPModerationError maps a non-200 response from /v1/moderations to
// a typed error. 400 bodies containing "does not support moderation" are
// converted to *UnsupportedCapabilityError; other statuses go through the
// existing ConvertError path.
func (c *Client) convertHTTPModerationError(statusCode int, body []byte) error {
	parsed := parseModerationError(body)

	// Locked phrasings from the gateway:
	//   "Provider <name> does not support moderation"
	//   "Provider <name> does not support multimodal moderation input"
	if statusCode == http.StatusBadRequest && parsed.parseErr == nil &&
		strings.Contains(parsed.Detail, "does not support") &&
		strings.Contains(parsed.Detail, "moderation") {
		providerID := parseUnsupportedProvider(parsed.Detail)
		operation := "moderation"
		if strings.Contains(parsed.Detail, "multimodal") {
			operation = "multimodal_moderation"
		}
		return newUnsupportedCapabilityError(providerID, operation, stderrors.New(parsed.Detail))
	}

	// Delegate to the existing ConvertError for other statuses by
	// synthesizing an openai.Error that ConvertError understands.
	return convertOpenAIError(&openai.Error{
		StatusCode: statusCode,
	})
}

// moderationError holds the parsed fields from a gateway error response body.
type moderationError struct {
	Detail   string
	raw      string
	parseErr error
}

// parseModerationError parses a gateway error response. Non-JSON bodies are
// preserved via the raw field. parseErr is set when the body cannot be decoded.
func parseModerationError(body []byte) moderationError {
	raw := string(body)
	var wrapper struct {
		Detail string `json:"detail"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return moderationError{raw: raw, parseErr: fmt.Errorf("unmarshal error body: %w", err)}
	}
	if wrapper.Detail == "" {
		return moderationError{raw: raw, parseErr: stderrors.New("error body missing \"detail\" field")}
	}
	return moderationError{Detail: wrapper.Detail, raw: raw}
}

// parseUnsupportedProvider extracts the provider name from the locked error
// phrasing "Provider <name> does not support ..." and returns "unknown" on
// any parse failure.
func parseUnsupportedProvider(detail string) string {
	const prefix = "Provider "
	rest, ok := strings.CutPrefix(detail, prefix)
	if !ok {
		return "unknown"
	}
	end := strings.Index(rest, " does not")
	if end <= 0 {
		return "unknown"
	}
	return rest[:end]
}
