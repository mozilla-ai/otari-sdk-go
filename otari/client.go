package otari

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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

	// placeholderAPIKey satisfies the OpenAI SDK's requirement that an API key
	// be set. The SDK still sends it in the Authorization header, but in
	// non-platform mode real auth is carried by the otari header, so this
	// is a non-secret placeholder that the gateway ignores.
	placeholderAPIKey = "otari-no-key"

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
// The client supports two authentication modes:
//   - Platform mode: uses a platform token as standard Bearer auth
//   - Non-platform mode: sends an otari API key via a custom Otari-Key header
type Client struct {
	// openaiClient is the OpenAI SDK client configured to point at the
	// otari gateway's /v1 endpoint.
	openaiClient openai.Client

	// apiBase is the gateway base URL, used for batch API calls that bypass
	// the OpenAI SDK.
	apiBase string

	// apiKey is the otari API key for non-platform mode, or the platform
	// token in platform mode, used for batch API calls that bypass the
	// OpenAI SDK.
	apiKey string

	// baseURL is the resolved base URL used to address endpoints
	// handled directly by this package (e.g. /v1/moderations, /v1/rerank).
	baseURL string

	// httpClient is the HTTP client used for endpoints the embedded
	// openai-go SDK does not expose (e.g. /v1/moderations, rerank, batch).
	// In non-platform mode this is the header-injecting client so otari
	// auth is preserved.
	httpClient *http.Client

	// platformToken is the Bearer token to set on requests we issue
	// directly in platform mode; empty in non-platform mode.
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

	// Build OpenAI client options. User's opts are cloned and auth-specific
	// opts are layered on top (later options win).
	sdkOpts := make([]option.RequestOption, 0, 3)

	httpClient := cfg.httpClientValue()
	var sdkAPIKey string

	if platformMode {
		sdkAPIKey = platformToken
		// Wrap the HTTP client so raw HTTP calls (e.g. Rerank) that bypass
		// the OpenAI SDK also carry the platform Bearer token.
		httpClient = newBearerClient(httpClient, platformToken)
	} else {
		// Non-platform mode: override any user-supplied API key with the
		// placeholder so secrets can't leak via the Authorization header.
		sdkAPIKey = placeholderAPIKey
		if otariKey != "" {
			httpClient = newHeaderClient(httpClient, bearerPrefix+otariKey)
		}
	}

	sdkOpts = append(sdkOpts,
		option.WithAPIKey(sdkAPIKey),
		option.WithHTTPClient(httpClient),
		option.WithBaseURL(apiBase),
	)

	// Determine the key to use for batch API calls.
	var batchAPIKey string
	if platformMode {
		batchAPIKey = platformToken
	} else {
		batchAPIKey = otariKey
	}

	// rawBaseURL strips any trailing /v1 suffix so that raw HTTP endpoints
	// (e.g. /v1/rerank) can prepend the version prefix themselves.
	rawBaseURL := strings.TrimSuffix(apiBase, "/v1")
	rawBaseURL = strings.TrimSuffix(rawBaseURL, "/v1/")

	return &Client{
		openaiClient:  openai.NewClient(sdkOpts...),
		apiBase:       apiBase,
		apiKey:        batchAPIKey,
		baseURL:       rawBaseURL,
		httpClient:    httpClient,
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

// Completion performs a chat completion request.
// Otari-specific errors (402, 502, 504) are converted to typed errors.
func (c *Client) Completion(
	ctx context.Context,
	params CompletionParams,
) (*ChatCompletion, error) {
	if err := validateCompletionParams(params); err != nil {
		return nil, err
	}

	req := convertParams(params)
	resp, err := c.openaiClient.Chat.Completions.New(ctx, req)
	if err != nil {
		return nil, convertOpenAIError(err)
	}

	return convertResponse(resp), nil
}

// CompletionStream performs a streaming chat completion request.
// Otari-specific errors (402, 502, 504) are converted to typed errors.
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

		req := convertParams(params)
		stream := c.openaiClient.Chat.Completions.NewStreaming(ctx, req)

		for stream.Next() {
			chunk := stream.Current()
			select {
			case chunks <- convertChunk(&chunk):
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			}
		}

		if err := stream.Err(); err != nil {
			errs <- convertOpenAIError(err)
		}
	}()

	return chunks, errs
}

// ConvertError converts errors to otari-specific typed errors where
// applicable. This is exported so callers can use it for custom error handling.
func (c *Client) ConvertError(err error) error {
	return convertOpenAIError(err)
}

// Embedding generates embeddings for the given input.
// Otari-specific errors (402, 502, 504) are converted to typed errors.
func (c *Client) Embedding(
	ctx context.Context,
	params EmbeddingParams,
) (*EmbeddingResponse, error) {
	req := convertEmbeddingParams(params)

	resp, err := c.openaiClient.Embeddings.New(ctx, req)
	if err != nil {
		return nil, convertOpenAIError(err)
	}

	return convertEmbeddingResponse(resp), nil
}

// ListModels returns a list of available models from the gateway.
// Otari-specific errors (402, 502, 504) are converted to typed errors.
func (c *Client) ListModels(ctx context.Context) (*ModelsResponse, error) {
	resp, err := c.openaiClient.Models.List(ctx)
	if err != nil {
		return nil, convertOpenAIError(err)
	}

	models := make([]Model, 0, len(resp.Data))
	for _, model := range resp.Data {
		models = append(models, Model{
			ID:      model.ID,
			Object:  objectModel,
			Created: model.Created,
			OwnedBy: string(model.OwnedBy),
		})
	}

	return &ModelsResponse{
		Object: objectList,
		Data:   models,
	}, nil
}
