package otari

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// config holds the configuration for the otari client.
type config struct {
	// apiKey is the API key for authentication.
	apiKey string

	// baseURL is the base URL for the API.
	baseURL string

	// extra holds provider-specific configuration options.
	extra map[string]any

	// timeout is the request timeout. If zero, a default timeout is used.
	timeout time.Duration

	// httpClient is a custom HTTP client. Access via httpClientValue() method which
	// handles lazy creation with the configured timeout if not explicitly set.
	httpClient     *http.Client
	httpClientOnce sync.Once
}

// Option is a function that modifies the config.
type Option func(*config) error

// newConfig creates a config with the given options applied.
func newConfig(opts ...Option) (*config, error) {
	cfg := &config{
		timeout: 120 * time.Second,
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// WithAPIKey sets the API key. Whitespace is automatically trimmed.
func WithAPIKey(key string) Option {
	return func(c *config) error {
		key = strings.TrimSpace(key)
		if key == "" {
			return fmt.Errorf("API key cannot be empty")
		}
		c.apiKey = key
		return nil
	}
}

// WithBaseURL sets the base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *config) error {
		baseURL = strings.TrimSpace(baseURL)
		if baseURL == "" {
			return fmt.Errorf("base URL cannot be empty")
		}

		parsed, err := url.Parse(baseURL)
		if err != nil {
			return fmt.Errorf("invalid base URL: %w", err)
		}

		if parsed.Scheme == "" || parsed.Host == "" {
			return fmt.Errorf("base URL must have scheme and host")
		}

		c.baseURL = baseURL
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *config) error {
		if client == nil {
			return fmt.Errorf("HTTP client cannot be nil")
		}
		c.httpClient = client
		return nil
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *config) error {
		if d <= 0 {
			return fmt.Errorf("timeout must be positive, got %v", d)
		}
		c.timeout = d
		return nil
	}
}

// WithOtariKey sets the otari API key for non-platform mode authentication.
// The key is sent as a Bearer-formatted value in the Otari-Key header.
func WithOtariKey(key string) Option {
	return withExtra(extraKeyOtariKey, key)
}

// WithPlatformMode explicitly enables platform mode authentication.
// In platform mode, the token (from WithAPIKey, the OTARI_AI_TOKEN env var,
// or the legacy GATEWAY_PLATFORM_TOKEN alias) is used as standard Bearer
// authentication, and the base URL defaults to the hosted gateway.
func WithPlatformMode() Option {
	return withExtra(extraKeyPlatformMode, true)
}

// withExtra sets extra provider-specific configuration.
func withExtra(key string, value any) Option {
	return func(c *config) error {
		key = strings.TrimSpace(key)
		if key == "" {
			return fmt.Errorf("extra key cannot be empty")
		}
		if c.extra == nil {
			c.extra = make(map[string]any)
		}
		c.extra[key] = value
		return nil
	}
}

// extraValue retrieves a provider-specific configuration value.
func (c *config) extraValue(key string) (any, bool) {
	if c.extra == nil {
		return nil, false
	}
	v, ok := c.extra[key]
	return v, ok
}

// httpClientValue returns the configured HTTP client, or lazily creates one
// using the configured timeout if no custom client was provided.
func (c *config) httpClientValue() *http.Client {
	c.httpClientOnce.Do(func() {
		if c.httpClient == nil {
			c.httpClient = &http.Client{Timeout: c.timeout}
		}
	})
	return c.httpClient
}

// resolveEnv returns the value of the specified environment variable,
// trimming whitespace. Returns empty string if the variable is not set or empty.
func resolveEnv(envVar string) string {
	if envVar == "" {
		return ""
	}
	return strings.TrimSpace(os.Getenv(envVar))
}

// resolvePlatformToken resolves the platform token from the environment,
// preferring the canonical OTARI_AI_TOKEN and falling back to the legacy
// GATEWAY_PLATFORM_TOKEN alias. Returns empty string if neither is set.
func resolvePlatformToken() string {
	if token := resolveEnv(envPlatformToken); token != "" {
		return token
	}
	return resolveEnv(envPlatformTokenLegacy)
}

// resolveBaseURL resolves the base URL from config, environment variable, or default value.
func (c *config) resolveBaseURL(envVar, defaultVal string) (string, error) {
	baseURL := c.baseURL
	if baseURL == "" {
		baseURL = resolveEnv(envVar)
	}
	if baseURL == "" {
		baseURL = defaultVal
	}
	if baseURL == "" {
		return "", nil
	}

	baseURL = strings.TrimSpace(baseURL)

	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL %q: %w", baseURL, err)
	}

	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("base URL %q must have scheme and host", baseURL)
	}

	return baseURL, nil
}

// resolveAPIKey returns the API key from config if set, otherwise falls back
// to the specified environment variable.
func (c *config) resolveAPIKey(envVar string) string {
	if c.apiKey != "" {
		return c.apiKey
	}
	return os.Getenv(envVar)
}
