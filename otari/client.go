package otari

import (
	"context"
	"fmt"
	"strings"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// Provider configuration constants.
const (
	// apiKeyHeaderName is the HTTP header that carries the otari API key in
	// non-platform mode.
	apiKeyHeaderName = "Otari-Key"

	// authorizationHeader is the standard HTTP Authorization header.
	authorizationHeader = "Authorization"

	// bearerPrefix is the Authorization scheme prefix applied to the otari
	// key before it is placed in the header (e.g. "Bearer <key>").
	bearerPrefix = "Bearer "

	// contentTypeHeader is the standard HTTP Content-Type header.
	contentTypeHeader = "Content-Type"

	// contentTypeJSON is the Content-Type value for JSON request bodies.
	contentTypeJSON = "application/json"

	// envAPIBase is the environment variable read for the base URL
	// when WithBaseURL is not passed to New.
	envAPIBase = "GATEWAY_API_BASE"

	// envAPIKey is the environment variable read for the API key used
	// in non-platform mode when WithOtariKey is not passed to New.
	envAPIKey = "GATEWAY_API_KEY"

	// envPlatformToken is the canonical environment variable read for the
	// platform token used as Bearer auth in platform mode.
	envPlatformToken = "OTARI_AI_TOKEN"

	// envPlatformTokenLegacy is the legacy alias for the platform token
	// environment variable. It is consulted only when envPlatformToken
	// (OTARI_AI_TOKEN) is unset; the canonical name takes precedence.
	envPlatformTokenLegacy = "GATEWAY_PLATFORM_TOKEN"

	// defaultPlatformBaseURL is the hosted gateway base URL used when the
	// client is in platform mode and no base URL is otherwise provided.
	defaultPlatformBaseURL = "https://api.otari.ai"

	// extraKeyOtariKey is the config extra key used to coordinate
	// WithOtariKey (writer) with the resolver logic in New (reader).
	extraKeyOtariKey = "otari_key"

	// extraKeyPlatformMode is the config extra key used to coordinate
	// WithPlatformMode (writer) with the resolver logic in New (reader).
	extraKeyPlatformMode = "platform_mode"

	// providerName is the value returned by Client.Name and embedded in
	// errors produced by this package.
	providerName = "otari"
)

// Client implements the otari SDK client. It connects to an otari gateway
// server, which proxies requests to underlying LLM providers.
//
// The client is a thin, ergonomic shell over the OpenAPI-generated
// core in the otari/client package. Non-streaming inference is issued through
// that generated core's configured transport; streaming is the hand-written SSE
// shim; generated/HTTP errors are mapped to the typed errors in this package.
//
// The client supports two authentication modes:
//   - Platform mode: uses a platform token as standard Bearer auth
//   - Non-platform mode: sends an otari API key via a custom Otari-Key header
type Client struct {
	// api is the generated client configured to point at the gateway's /v1
	// base, carrying the per-mode auth header as a default header.
	api *client.APIClient

	// apiBase is the gateway root URL (without /v1).
	apiBase string

	// baseURL is the gateway root URL, kept for ControlPlane wiring.
	baseURL string

	// adminAuth is the resolved Authorization bearer value for batch calls
	// in platform mode (empty in non-platform mode, where Otari-Key is used).
	platformToken string

	// platformMode indicates whether the client is operating in platform
	// mode (using platform token for Bearer auth) or non-platform mode
	// (using otari key in custom header).
	platformMode bool
}

// New creates a new otari client.
//
// In platform mode the base URL defaults to the hosted gateway at
// https://api.otari.ai, so it can be omitted. It can still be overridden via
// WithBaseURL() or the GATEWAY_API_BASE environment variable. In non-platform
// mode a base URL is required.
//
// Authentication mode is determined as follows:
//   - If WithPlatformMode() is passed, platform mode is used. The token is
//     resolved from WithAPIKey(), then the OTARI_AI_TOKEN environment variable
//     (legacy alias GATEWAY_PLATFORM_TOKEN).
//   - If a platform token env var (OTARI_AI_TOKEN, or legacy
//     GATEWAY_PLATFORM_TOKEN) is set and no explicit API key or otari key is
//     provided, platform mode is auto-detected.
//   - Otherwise, non-platform mode is used with the key from WithOtariKey()
//     or GATEWAY_API_KEY, sent via the Otari-Key header.
func New(opts ...Option) (*Client, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	// Resolve otari key from extras or GATEWAY_API_KEY env var.
	var otariKey string
	if v, ok := cfg.extraValue(extraKeyOtariKey); ok {
		if key, ok := v.(string); ok && key != "" {
			otariKey = key
		}
	}
	if otariKey == "" {
		otariKey = resolveEnv(envAPIKey)
	}

	// Determine authentication mode.
	platformMode := false
	var platformToken string

	// Explicit opt-in via WithPlatformMode().
	if v, ok := cfg.extraValue(extraKeyPlatformMode); ok {
		if enabled, ok := v.(bool); ok && enabled {
			platformMode = true
			platformToken = cfg.apiKey
			if platformToken == "" {
				platformToken = resolvePlatformToken()
			}
		}
	}

	// Auto-detect: a platform token is set in the environment and no explicit
	// API key or otari key configured. An otari key signals non-platform intent.
	if !platformMode {
		envToken := resolvePlatformToken()
		if envToken != "" && cfg.apiKey == "" && otariKey == "" {
			platformMode = true
			platformToken = envToken
		}
	}

	if platformMode && platformToken == "" {
		return nil, fmt.Errorf(
			"platform mode requires a token (pass WithAPIKey option or set %s env var)",
			envPlatformToken,
		)
	}

	// Resolve apiBase. In platform mode, fall back to the hosted gateway URL
	// when no base URL is provided; the non-platform path keeps requiring one.
	defaultBaseURL := ""
	if platformMode {
		defaultBaseURL = defaultPlatformBaseURL
	}
	apiBase, err := cfg.resolveBaseURL(envAPIBase, defaultBaseURL)
	if err != nil {
		return nil, err
	}
	if apiBase == "" {
		return nil, fmt.Errorf(
			"otari base URL is required (set via WithBaseURL option or %s env var)",
			envAPIBase,
		)
	}
	apiBase = strings.TrimRight(apiBase, "/")

	// Normalize to the gateway root (without a trailing /v1). The generated
	// client's base server URL is then <root>/v1, and inference/batch paths are
	// joined relative to it (e.g. "/chat/completions"). Normalizing here means
	// the hosted default and a self-hosted base URL both work whether or not the
	// caller includes /v1 (mirrors the TS and Python SDKs).
	apiBase = strings.TrimSuffix(apiBase, "/v1")
	apiBase = strings.TrimRight(apiBase, "/")

	// Assemble the per-mode default auth header fed into the generated client
	// configuration. Platform mode uses standard Bearer Authorization;
	// non-platform mode uses the custom Otari-Key header.
	httpClient := cfg.httpClientValue()
	headers := map[string]string{}
	if platformMode {
		headers[authorizationHeader] = bearerPrefix + platformToken
	} else if otariKey != "" {
		headers[apiKeyHeaderName] = bearerPrefix + otariKey
	}

	// The generated core's operation paths already include the /v1 prefix, so
	// its base server URL is the gateway root + /v1.
	api := newAPIClient(apiBase+"/v1", headers, httpClient)

	return &Client{
		api:           api,
		apiBase:       apiBase,
		baseURL:       apiBase,
		platformMode:  platformMode,
		platformToken: platformToken,
	}, nil
}

// Name returns the provider name.
func (c *Client) Name() string {
	return providerName
}

// Capabilities returns the full set of capabilities for the otari client.
// Since the gateway proxies to any backend provider, all features are
// optimistically marked as supported. Actual support depends on the
// underlying provider behind the gateway; consumers should handle
// unsupported-operation errors at call time.
func (c *Client) Capabilities() Capabilities {
	return Capabilities{
		Completion:          true,
		CompletionImage:     true,
		CompletionPDF:       true,
		CompletionReasoning: true,
		CompletionStreaming: true,
		CompletionTools:     true,
		Embedding:           true,
		ListModels:          true,
		Moderation:          true,
		Rerank:              true,
	}
}

// Completion performs a non-streaming chat completion request via
// POST /v1/chat/completions. Gateway errors are mapped to typed errors.
func (c *Client) Completion(
	ctx context.Context,
	params CompletionParams,
) (*ChatCompletion, error) {
	if err := validateCompletionParams(params); err != nil {
		return nil, err
	}

	body := completionBody(params, false)
	var out ChatCompletion
	if err := c.doJSON(ctx, "POST", "/chat/completions", body, &out, ""); err != nil {
		return nil, err
	}
	if out.Object == "" {
		out.Object = objectChatCompletion
	}
	return &out, nil
}

// CompletionStream performs a streaming chat completion request. It returns a
// channel of typed ChatCompletionChunk values and a buffered error channel.
// The SSE shim parses the gateway's text/event-stream and stops at [DONE].
func (c *Client) CompletionStream(
	ctx context.Context,
	params CompletionParams,
) (<-chan ChatCompletionChunk, <-chan error) {
	chunks := make(chan ChatCompletionChunk)
	errs := make(chan error, 1)

	go func() {
		defer close(chunks)
		defer close(errs)

		if err := validateCompletionParams(params); err != nil {
			errs <- err
			return
		}

		body := completionBody(params, true)
		if err := streamSSE(ctx, c, "/chat/completions", body, func(payload []byte) error {
			var chunk ChatCompletionChunk
			if err := jsonUnmarshal(payload, &chunk); err != nil {
				return err
			}
			if chunk.Object == "" {
				chunk.Object = objectChatCompletionChunk
			}
			select {
			case chunks <- chunk:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}); err != nil {
			errs <- err
		}
	}()

	return chunks, errs
}

// ConvertError is retained for API compatibility. It returns the error as-is
// (errors returned by this SDK are already typed via the unified mapper).
func (c *Client) ConvertError(err error) error {
	return err
}

// Embedding generates embeddings for the given input via POST /v1/embeddings.
func (c *Client) Embedding(
	ctx context.Context,
	params EmbeddingParams,
) (*EmbeddingResponse, error) {
	var out EmbeddingResponse
	if err := c.doJSON(ctx, "POST", "/embeddings", params, &out, ""); err != nil {
		return nil, err
	}
	if out.Object == "" {
		out.Object = objectList
	}
	return &out, nil
}

// ListModels returns the list of available models from GET /v1/models.
func (c *Client) ListModels(ctx context.Context) (*ModelsResponse, error) {
	var out ModelsResponse
	if err := c.doJSON(ctx, "GET", "/models", nil, &out, ""); err != nil {
		return nil, err
	}
	if out.Object == "" {
		out.Object = objectList
	}
	return &out, nil
}

// Message creates an Anthropic-shaped message via POST /v1/messages.
//
// This endpoint has no OpenAI-SDK seam and was previously missing from the SDK.
// Its response is opaque (response_model=None on the gateway), so the decoded
// JSON is returned as a map. max_tokens is required by the gateway.
func (c *Client) Message(
	ctx context.Context,
	params MessageParams,
) (map[string]any, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	if params.MaxTokens <= 0 {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("max_tokens is required"))
	}
	body := messageBody(params, false)
	var out map[string]any
	if err := c.doJSON(ctx, "POST", "/messages", body, &out, ""); err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStream streams an Anthropic-shaped message via POST /v1/messages.
// The /messages event stream has no single typed chunk model, so each event is
// delivered as the raw decoded JSON map.
func (c *Client) MessageStream(
	ctx context.Context,
	params MessageParams,
) (<-chan map[string]any, <-chan error) {
	return c.rawStream(ctx, "/messages", messageBody(params, true), func() error {
		if params.Model == "" {
			return newInvalidRequestError(providerName, fmt.Errorf("model is required"))
		}
		if params.MaxTokens <= 0 {
			return newInvalidRequestError(providerName, fmt.Errorf("max_tokens is required"))
		}
		return nil
	})
}

// Response creates a response via the OpenAI-style Responses API
// (POST /v1/responses). Its response is opaque on the gateway, so the decoded
// JSON is returned as a map.
func (c *Client) Response(
	ctx context.Context,
	params ResponseParams,
) (map[string]any, error) {
	if params.Model == "" {
		return nil, newInvalidRequestError(providerName, fmt.Errorf("model is required"))
	}
	body := responseBody(params, false)
	var out map[string]any
	if err := c.doJSON(ctx, "POST", "/responses", body, &out, ""); err != nil {
		return nil, err
	}
	return out, nil
}

// ResponseStream streams a response via POST /v1/responses, delivering each
// event as the raw decoded JSON map.
func (c *Client) ResponseStream(
	ctx context.Context,
	params ResponseParams,
) (<-chan map[string]any, <-chan error) {
	return c.rawStream(ctx, "/responses", responseBody(params, true), func() error {
		if params.Model == "" {
			return newInvalidRequestError(providerName, fmt.Errorf("model is required"))
		}
		return nil
	})
}

// rawStream runs the SSE shim for endpoints whose events have no typed chunk
// model, delivering each event as a decoded JSON map. validate runs first.
func (c *Client) rawStream(
	ctx context.Context,
	path string,
	body any,
	validate func() error,
) (<-chan map[string]any, <-chan error) {
	events := make(chan map[string]any)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		if err := validate(); err != nil {
			errs <- err
			return
		}
		if err := streamSSE(ctx, c, path, body, func(payload []byte) error {
			var evt map[string]any
			if err := jsonUnmarshal(payload, &evt); err != nil {
				return err
			}
			select {
			case events <- evt:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}); err != nil {
			errs <- err
		}
	}()

	return events, errs
}
