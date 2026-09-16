package otari

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// translate maps a generated control-plane error onto the SDK's typed error
// hierarchy, so management calls (keys/users/budgets/pricing/usage) classify
// the same way as inference via errors.Is/errors.As. When the generated call
// returns an error for an HTTP error status (>= 400), it surfaces the
// corresponding typed error (e.g. *AuthenticationError, *RateLimitError)
// instead of the raw *client.GenericOpenAPIError. Transport-level failures
// (no response) and successes pass through unchanged.
func translate(resp *http.Response, err error) error {
	if err == nil {
		return nil
	}
	if resp == nil || resp.StatusCode < 400 {
		return err
	}
	var body []byte
	var apiErr *client.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		body = apiErr.Body()
	}
	return mapHTTPError(resp.StatusCode, resp.Header, body, "")
}

// ControlPlaneClient is a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage).
//
// Each resource (Keys, Users, Budgets, Pricing, Usage) exposes ergonomic
// aliases (Create, Get, List, Update, Delete, ...) with short,
// generator-name-free names that take their inputs directly and execute the
// request, for example:
//
//	created, _, err := cp.Keys.Create(ctx, client.CreateKeyRequest{})
//
// Optional query parameters (pagination, pricing as-of, usage filters) are
// passed as pointers; pass nil to omit them. The underlying generated
// *client.APIClient stays reachable via Raw, so the full generated surface
// (generator-named operations, fluent request builders) remains available as
// an escape hatch:
//
//	created, _, err := cp.Raw.KeysAPI.KeysCreateKey(ctx).
//	    CreateKeyRequest(client.CreateKeyRequest{}).Execute()
type ControlPlaneClient struct {
	// Raw is the underlying generated API client (escape hatch).
	Raw     *client.APIClient
	Keys    KeysResource
	Users   UsersResource
	Budgets BudgetsResource
	Pricing PricingResource
	Usage   UsageResource
}

// KeysResource exposes ergonomic aliases for the API-keys management endpoints.
type KeysResource struct{ api *client.KeysAPIService }

func (r KeysResource) Create(ctx context.Context, req client.CreateKeyRequest) (*client.CreateKeyResponse, *http.Response, error) {
	v, resp, err := r.api.KeysCreateKey(ctx).CreateKeyRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r KeysResource) Get(ctx context.Context, keyID string) (*client.KeyInfo, *http.Response, error) {
	v, resp, err := r.api.KeysGetKey(ctx, keyID).Execute()
	return v, resp, translate(resp, err)
}

func (r KeysResource) List(ctx context.Context, skip, limit *int32) ([]client.KeyInfo, *http.Response, error) {
	req := r.api.KeysListKeys(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

func (r KeysResource) Update(ctx context.Context, keyID string, req client.UpdateKeyRequest) (*client.KeyInfo, *http.Response, error) {
	v, resp, err := r.api.KeysUpdateKey(ctx, keyID).UpdateKeyRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r KeysResource) Delete(ctx context.Context, keyID string) (*http.Response, error) {
	resp, err := r.api.KeysDeleteKey(ctx, keyID).Execute()
	return resp, translate(resp, err)
}

// UsersResource exposes ergonomic aliases for the users management endpoints.
type UsersResource struct{ api *client.UsersAPIService }

func (r UsersResource) Create(ctx context.Context, req client.CreateUserRequest) (*client.UserResponse, *http.Response, error) {
	v, resp, err := r.api.UsersCreateUser(ctx).CreateUserRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r UsersResource) Get(ctx context.Context, userID string) (*client.UserResponse, *http.Response, error) {
	v, resp, err := r.api.UsersGetUser(ctx, userID).Execute()
	return v, resp, translate(resp, err)
}

func (r UsersResource) List(ctx context.Context, skip, limit *int32) ([]client.UserResponse, *http.Response, error) {
	req := r.api.UsersListUsers(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

func (r UsersResource) Update(ctx context.Context, userID string, req client.UpdateUserRequest) (*client.UserResponse, *http.Response, error) {
	v, resp, err := r.api.UsersUpdateUser(ctx, userID).UpdateUserRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r UsersResource) Delete(ctx context.Context, userID string) (*http.Response, error) {
	resp, err := r.api.UsersDeleteUser(ctx, userID).Execute()
	return resp, translate(resp, err)
}

func (r UsersResource) GetUsage(ctx context.Context, userID string, skip, limit *int32) ([]client.UsageLogResponse, *http.Response, error) {
	req := r.api.UsersGetUserUsage(ctx, userID)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

// BudgetsResource exposes ergonomic aliases for the budgets management endpoints.
type BudgetsResource struct{ api *client.BudgetsAPIService }

func (r BudgetsResource) Create(ctx context.Context, req client.CreateBudgetRequest) (*client.BudgetResponse, *http.Response, error) {
	v, resp, err := r.api.BudgetsCreateBudget(ctx).CreateBudgetRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r BudgetsResource) Get(ctx context.Context, budgetID string) (*client.BudgetResponse, *http.Response, error) {
	v, resp, err := r.api.BudgetsGetBudget(ctx, budgetID).Execute()
	return v, resp, translate(resp, err)
}

func (r BudgetsResource) List(ctx context.Context, skip, limit *int32) ([]client.BudgetResponse, *http.Response, error) {
	req := r.api.BudgetsListBudgets(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

func (r BudgetsResource) Update(ctx context.Context, budgetID string, req client.UpdateBudgetRequest) (*client.BudgetResponse, *http.Response, error) {
	v, resp, err := r.api.BudgetsUpdateBudget(ctx, budgetID).UpdateBudgetRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r BudgetsResource) Delete(ctx context.Context, budgetID string) (*http.Response, error) {
	resp, err := r.api.BudgetsDeleteBudget(ctx, budgetID).Execute()
	return resp, translate(resp, err)
}

// PricingResource exposes ergonomic aliases for the model-pricing management endpoints.
type PricingResource struct{ api *client.PricingAPIService }

func (r PricingResource) List(ctx context.Context, skip, limit *int32) ([]client.PricingResponse, *http.Response, error) {
	req := r.api.PricingListPricing(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

func (r PricingResource) Get(ctx context.Context, modelKey string, asOf *time.Time) (*client.PricingResponse, *http.Response, error) {
	req := r.api.PricingGetPricing(ctx, modelKey)
	if asOf != nil {
		req = req.AsOf(*asOf)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

func (r PricingResource) Set(ctx context.Context, req client.SetPricingRequest) (*client.PricingResponse, *http.Response, error) {
	v, resp, err := r.api.PricingSetPricing(ctx).SetPricingRequest(req).Execute()
	return v, resp, translate(resp, err)
}

func (r PricingResource) Delete(ctx context.Context, modelKey string, effectiveAt *time.Time) (*http.Response, error) {
	req := r.api.PricingDeletePricing(ctx, modelKey)
	if effectiveAt != nil {
		req = req.EffectiveAt(*effectiveAt)
	}
	resp, err := req.Execute()
	return resp, translate(resp, err)
}

func (r PricingResource) GetHistory(ctx context.Context, modelKey string) ([]client.PricingResponse, *http.Response, error) {
	v, resp, err := r.api.PricingGetPricingHistory(ctx, modelKey).Execute()
	return v, resp, translate(resp, err)
}

// UsageResource exposes ergonomic aliases for the usage-log management endpoints.
type UsageResource struct{ api *client.UsageAPIService }

func (r UsageResource) List(ctx context.Context, startDate, endDate *time.Time, userID *string, skip, limit *int32) ([]client.UsageEntry, *http.Response, error) {
	req := r.api.UsageListUsage(ctx)
	if startDate != nil {
		req = req.StartDate(*startDate)
	}
	if endDate != nil {
		req = req.EndDate(*endDate)
	}
	if userID != nil {
		// UserId is repeatable upstream (user_id=a&user_id=b, max 50). This alias
		// keeps its single-user signature and wraps; multi-user filtering is
		// reachable through the generated client.
		req = req.UserId([]string{*userID})
	}
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	v, resp, err := req.Execute()
	return v, resp, translate(resp, err)
}

// ControlPlane returns a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage).
//
// These endpoints authenticate with Authorization: Bearer <admin/master key>,
// which is distinct from the Otari-Key header used for inference. Pass the
// gateway master key (or an admin token). The generated control-plane operation
// paths already include the API root, so the configured server URL is the
// gateway origin.
//
//	cp := client.ControlPlane("gateway-master-key")
//	created, _, err := cp.Keys.Create(ctx, client.CreateKeyRequest{})
func (c *Client) ControlPlane(adminKey string) *ControlPlaneClient {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{{URL: c.baseURL}}
	cfg.AddDefaultHeader(authorizationHeader, bearerPrefix+adminKey)
	raw := client.NewAPIClient(cfg)
	return &ControlPlaneClient{
		Raw:     raw,
		Keys:    KeysResource{api: raw.KeysAPI},
		Users:   UsersResource{api: raw.UsersAPI},
		Budgets: BudgetsResource{api: raw.BudgetsAPI},
		Pricing: PricingResource{api: raw.PricingAPI},
		Usage:   UsageResource{api: raw.UsageAPI},
	}
}
