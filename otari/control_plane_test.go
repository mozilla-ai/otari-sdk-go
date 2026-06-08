package otari

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

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
