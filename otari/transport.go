package otari

import "net/http"

// headerTransport wraps an http.RoundTripper to inject a custom header into
// every request. It is retained as a small utility (and exercised by tests);
// the client itself now sets per-mode auth as default headers on each request
// rather than wrapping the user's transport.
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
