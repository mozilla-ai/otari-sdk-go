package otari

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// Moderation runs a content moderation check via POST /v1/moderations.
//
// Returns an *UnsupportedCapabilityError (matching ErrUnsupported) when the
// gateway reports that the chosen backend provider does not support moderation;
// that mapping is handled by the unified error mapper in both auth modes.
func (c *Client) Moderation(
	ctx context.Context,
	params ModerationParams,
) (*ModerationResponse, error) {
	u := &url.URL{Path: "/moderations"}
	if params.IncludeRaw {
		u.RawQuery = url.Values{"include_raw": {"true"}}.Encode()
	}

	var out ModerationResponse
	if err := c.doJSON(ctx, http.MethodPost, u.RequestURI(), params, &out, ""); err != nil {
		return nil, err
	}
	return &out, nil
}

// parseUnsupportedProvider extracts the provider name from the locked error
// phrasing "Provider <name> does not support ..." and returns "unknown" on any
// parse failure.
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
