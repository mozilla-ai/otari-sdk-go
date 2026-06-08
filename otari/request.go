package otari

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// doJSON issues a buffered JSON request through the generated client's
// configured transport (HTTP client + default auth headers + base server URL)
// and decodes a 2xx body into out. Non-2xx responses are routed through the
// unified mapHTTPError so both auth modes produce the same typed errors.
//
// This reuses the generated client.APIClient for configuration and auth wiring
// rather than the openai-go SDK, while keeping the SDK's stable
// public request/response types as the marshal/decode targets. batchID enriches
// 409 batch errors; pass "" when not applicable.
func (c *Client) doJSON(
	ctx context.Context,
	method, path string,
	in any,
	out any,
	batchID string,
) error {
	resp, body, err := c.doRaw(ctx, method, path, in, false)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= http.StatusBadRequest {
		return mapHTTPError(resp.StatusCode, resp.Header, body, batchID)
	}
	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return newProviderError(providerName, fmt.Errorf("decode response: %w", err))
		}
	}
	return nil
}

// doRaw builds and sends a request via the generated client's configuration,
// returning the response together with the fully-read body. When stream is
// true the Accept header requests text/event-stream and the body is NOT read
// (the caller streams resp.Body); body is returned nil in that case unless the
// response is an error (>=400), in which case it is read so the caller can map
// it. The generated config's default headers carry the per-mode auth.
func (c *Client) doRaw(
	ctx context.Context,
	method, path string,
	in any,
	stream bool,
) (*http.Response, []byte, error) {
	cfg := c.api.GetConfig()

	var bodyReader io.Reader
	if in != nil {
		raw, err := json.Marshal(in)
		if err != nil {
			return nil, nil, newInvalidRequestError(providerName, fmt.Errorf("marshal request: %w", err))
		}
		bodyReader = bytes.NewReader(raw)
	}

	url := c.serverURL() + path
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, nil, newProviderError(providerName, err)
	}

	if in != nil {
		req.Header.Set(contentTypeHeader, contentTypeJSON)
	}
	if stream {
		req.Header.Set("Accept", "text/event-stream")
	}
	for k, v := range cfg.DefaultHeader {
		req.Header.Set(k, v)
	}

	resp, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, newProviderError(providerName, err)
	}

	if !stream {
		buf, rerr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if rerr != nil {
			return resp, nil, newProviderError(providerName, fmt.Errorf("read response: %w", rerr))
		}
		// Re-attach so callers may still inspect Body if needed.
		resp.Body = io.NopCloser(bytes.NewReader(buf))
		return resp, buf, nil
	}

	if resp.StatusCode >= http.StatusBadRequest {
		buf, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return resp, buf, nil
	}
	return resp, nil, nil
}

// serverURL returns the configured base server URL (the gateway root with the
// /v1 suffix) from the generated client configuration.
func (c *Client) serverURL() string {
	cfg := c.api.GetConfig()
	if len(cfg.Servers) > 0 {
		return cfg.Servers[0].URL
	}
	return ""
}

// structToMap marshals v to JSON and back into a map so callers can layer in
// extra fields and stream flags while honoring the struct's omitempty tags.
func structToMap(v any) map[string]any {
	raw, err := json.Marshal(v)
	if err != nil {
		return map[string]any{}
	}
	out := map[string]any{}
	_ = json.Unmarshal(raw, &out)
	return out
}

// newAPIClient builds a generated client.APIClient pointed at baseURL with the
// given default headers (carrying the per-mode auth header) and HTTP client.
func newAPIClient(baseURL string, headers map[string]string, httpClient *http.Client) *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{{URL: baseURL}}
	cfg.HTTPClient = httpClient
	for k, v := range headers {
		cfg.AddDefaultHeader(k, v)
	}
	return client.NewAPIClient(cfg)
}
