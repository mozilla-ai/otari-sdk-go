package otari

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
	"github.com/stretchr/testify/require"
)

// recordingGateway returns an httptest server that records the method, path,
// and Authorization header of the last request, replying with a JSON body that
// decodes into any control-plane response type.
func recordingGateway(t *testing.T, body string) (*httptest.Server, func() (method, path, auth string)) {
	t.Helper()
	var (
		mu                       sync.Mutex
		gotMethod, gotPath, gotA string
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		gotMethod, gotPath, gotA = r.Method, r.URL.Path, r.Header.Get(authorizationHeader)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv, func() (string, string, string) {
		mu.Lock()
		defer mu.Unlock()
		return gotMethod, gotPath, gotA
	}
}

func TestControlPlaneAliasesRouteToGeneratedOperations(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		body         string
		call         func(ctx context.Context, cp *ControlPlaneClient) error
		method, path string
	}{
		// These exercise the alias -> generated-operation wiring (method + path)
		// without a gateway. The body-carrying aliases (Create/Update/Set) and
		// response decoding are covered end to end by the integration test.
		{"keys.List", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Keys.List(ctx, nil, nil)
			return err
		}, http.MethodGet, "/v1/keys"},
		{"keys.Delete", "", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, err := cp.Keys.Delete(ctx, "k1")
			return err
		}, http.MethodDelete, "/v1/keys/k1"},
		{"users.List", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Users.List(ctx, nil, nil)
			return err
		}, http.MethodGet, "/v1/users"},
		{"users.GetUsage", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Users.GetUsage(ctx, "u1", nil, nil)
			return err
		}, http.MethodGet, "/v1/users/u1/usage"},
		{"budgets.List", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Budgets.List(ctx, nil, nil)
			return err
		}, http.MethodGet, "/v1/budgets"},
		{"pricing.List", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Pricing.List(ctx, nil, nil)
			return err
		}, http.MethodGet, "/v1/pricing"},
		{"pricing.GetHistory", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Pricing.GetHistory(ctx, "mk1")
			return err
		}, http.MethodGet, "/v1/pricing/mk1/history"},
		{"usage.List", "[]", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Usage.List(ctx, nil, nil, nil, nil, nil)
			return err
		}, http.MethodGet, "/v1/usage"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv, recorded := recordingGateway(t, tc.body)
			client, err := New(WithBaseURL(srv.URL), WithAPIKey("unused"))
			require.NoError(t, err)
			cp := client.ControlPlane("master")

			require.NoError(t, tc.call(context.Background(), cp))

			method, path, auth := recorded()
			require.Equal(t, tc.method, method)
			require.Equal(t, tc.path, path)
			require.Equal(t, bearerPrefix+"master", auth)
		})
	}
}

// errorGateway returns an httptest server that replies to every request with
// the given status code and JSON body, for exercising error mapping.
func errorGateway(t *testing.T, status int, body string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv
}

// TestControlPlaneMapsTypedErrors asserts that control-plane methods surface the
// same typed SDK errors as inference (via errors.Is/errors.As), rather than
// leaking the generated *client.GenericOpenAPIError. It covers a representative
// method per resource and the documented status -> error mappings.
func TestControlPlaneMapsTypedErrors(t *testing.T) {
	t.Parallel()

	statusCases := []struct {
		name    string
		status  int
		body    string
		wantErr error
	}{
		{"401 auth", http.StatusUnauthorized, `{"error":{"message":"invalid token"}}`, ErrAuthentication},
		{"403 auth", http.StatusForbidden, `{"error":{"message":"forbidden"}}`, ErrAuthentication},
		{"402 funds", http.StatusPaymentRequired, `{"error":{"message":"payment required"}}`, ErrInsufficientFunds},
		{"404 not found", http.StatusNotFound, `{"error":{"message":"not found"}}`, ErrModelNotFound},
		{"429 rate limit", http.StatusTooManyRequests, `{"error":{"message":"slow down"}}`, ErrRateLimit},
		{"502 upstream", http.StatusBadGateway, `{"error":{"message":"bad gateway"}}`, ErrUpstreamProvider},
		{"504 timeout", http.StatusGatewayTimeout, `{"error":{"message":"timed out"}}`, ErrTimeout},
	}

	resources := []struct {
		name string
		call func(ctx context.Context, cp *ControlPlaneClient) error
	}{
		{"keys.List", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Keys.List(ctx, nil, nil)
			return err
		}},
		{"keys.Delete", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, err := cp.Keys.Delete(ctx, "k1")
			return err
		}},
		{"users.Get", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Users.Get(ctx, "u1")
			return err
		}},
		{"budgets.List", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Budgets.List(ctx, nil, nil)
			return err
		}},
		{"pricing.List", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Pricing.List(ctx, nil, nil)
			return err
		}},
		{"usage.List", func(ctx context.Context, cp *ControlPlaneClient) error {
			_, _, err := cp.Usage.List(ctx, nil, nil, nil, nil, nil)
			return err
		}},
	}

	for _, sc := range statusCases {
		for _, rc := range resources {
			t.Run(sc.name+"/"+rc.name, func(t *testing.T) {
				t.Parallel()

				srv := errorGateway(t, sc.status, sc.body)
				cl, err := New(WithBaseURL(srv.URL), WithAPIKey("unused"))
				require.NoError(t, err)
				cp := cl.ControlPlane("master")

				gotErr := rc.call(context.Background(), cp)
				require.Error(t, gotErr)
				// Classifies as the SDK sentinel, like inference.
				require.ErrorIs(t, gotErr, sc.wantErr)
				// And does NOT leak the generated error type.
				var generic client.GenericOpenAPIError
				require.False(t, errors.As(gotErr, &generic),
					"control-plane error should not be a *client.GenericOpenAPIError")
			})
		}
	}
}

// TestControlPlaneAuthenticationErrorAs confirms callers can read typed fields
// off a control-plane error via errors.As (e.g. *AuthenticationError on 401).
func TestControlPlaneAuthenticationErrorAs(t *testing.T) {
	t.Parallel()

	srv := errorGateway(t, http.StatusUnauthorized, `{"error":{"message":"invalid token"}}`)
	cl, err := New(WithBaseURL(srv.URL), WithAPIKey("unused"))
	require.NoError(t, err)
	cp := cl.ControlPlane("master")

	_, _, gotErr := cp.Keys.Create(context.Background(), client.CreateKeyRequest{})
	require.Error(t, gotErr)

	var authErr *AuthenticationError
	require.True(t, errors.As(gotErr, &authErr),
		"expected *AuthenticationError, got %T", gotErr)
}

func TestControlPlaneRawEscapeHatch(t *testing.T) {
	t.Parallel()

	srv, recorded := recordingGateway(t, "[]")
	client, err := New(WithBaseURL(srv.URL), WithAPIKey("unused"))
	require.NoError(t, err)
	cp := client.ControlPlane("master")

	require.NotNil(t, cp.Raw)
	_, _, err = cp.Raw.KeysAPI.ListKeysV1KeysGet(context.Background()).Execute()
	require.NoError(t, err)

	method, path, _ := recorded()
	require.Equal(t, http.MethodGet, method)
	require.Equal(t, "/v1/keys", path)
}
