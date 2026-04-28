package otari

import "net/http"

// headerTransport wraps an http.RoundTripper to inject custom headers into
// every request. Used in non-platform mode to add the otari authentication
// header.
type headerTransport struct {
	base   http.RoundTripper
	header string
	value  string
}

// RoundTrip implements http.RoundTripper by cloning the request and injecting
// the configured header before delegating to the base transport.
func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set(t.header, t.value)
	return t.base.RoundTrip(clone)
}

// newBearerClient wraps the given HTTP client's transport to inject a standard
// Authorization: Bearer <token> header into every request. Used in platform
// mode for raw HTTP calls (e.g. Rerank) that bypass the OpenAI SDK.
func newBearerClient(base *http.Client, token string) *http.Client {
	transport := base.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	return &http.Client{
		Timeout: base.Timeout,
		Transport: &headerTransport{
			base:   transport,
			header: "Authorization",
			value:  bearerPrefix + token,
		},
	}
}

// newHeaderClient wraps the given HTTP client's transport to inject the
// otari authentication header into every request, preserving the base
// client's timeout and transport settings.
func newHeaderClient(base *http.Client, headerValue string) *http.Client {
	transport := base.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	return &http.Client{
		Timeout: base.Timeout,
		Transport: &headerTransport{
			base:   transport,
			header: apiKeyHeaderName,
			value:  headerValue,
		},
	}
}
