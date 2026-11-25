package otari

import (
	"context"
	stderrors "errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRedirectIsRefused pins that no request path follows a redirect: the
// gateway answers 307 to another host, and that host must see nothing.
func TestRedirectIsRefused(t *testing.T) {
	t.Parallel()

	followAll := &http.Client{
		CheckRedirect: func(*http.Request, []*http.Request) error { return nil },
	}
	// Runs after the parallel subtests: the SDK must copy the caller's client.
	t.Cleanup(func() {
		require.NoError(t, followAll.CheckRedirect(nil, nil), "the caller's client was modified")
	})

	tests := []struct {
		name string
		opts []Option
		call func(ctx context.Context, c *Client) error
	}{
		{
			name: "generated core GET",
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListModels(ctx)
				return err
			},
		},
		{
			name: "hand-built POST",
			call: func(ctx context.Context, c *Client) error {
				_, err := c.Completion(ctx, mockCompletionParams())
				return err
			},
		},
		{
			name: "streaming POST",
			call: func(ctx context.Context, c *Client) error {
				chunks, errs := c.CompletionStream(ctx, mockCompletionParams())
				for range chunks {
				}
				return <-errs
			},
		},
		{
			name: "control plane",
			call: func(ctx context.Context, c *Client) error {
				_, _, err := c.ControlPlane("admin_key").Keys.List(ctx, nil, nil)
				return err
			},
		},
		{
			name: "custom client that would follow",
			opts: []Option{WithHTTPClient(followAll)},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListModels(ctx)
				return err
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var hits atomic.Int32
			other := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				hits.Add(1)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"object":"list","data":[]}`))
			}))
			t.Cleanup(other.Close)

			gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, other.URL+r.URL.Path, http.StatusTemporaryRedirect)
			}))
			t.Cleanup(gateway.Close)

			opts := append([]Option{WithBaseURL(gateway.URL), WithOtariKey("gw_test_key")}, tc.opts...)
			c, err := New(opts...)
			require.NoError(t, err)

			err = tc.call(context.Background(), c)
			require.Error(t, err)
			require.True(t, stderrors.Is(err, ErrRedirect), "got %v", err)
			require.Contains(t, err.Error(), other.URL)
			require.Zero(t, hits.Load(), "the redirect target received a request")
		})
	}
}
