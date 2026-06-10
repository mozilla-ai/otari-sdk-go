package otari

import (
	"context"
	"net/http"
	"time"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

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
//	created, _, err := cp.Raw.KeysAPI.CreateKeyV1KeysPost(ctx).
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
	return r.api.CreateKeyV1KeysPost(ctx).CreateKeyRequest(req).Execute()
}

func (r KeysResource) Get(ctx context.Context, keyID string) (*client.KeyInfo, *http.Response, error) {
	return r.api.GetKeyV1KeysKeyIdGet(ctx, keyID).Execute()
}

func (r KeysResource) List(ctx context.Context, skip, limit *int32) ([]client.KeyInfo, *http.Response, error) {
	req := r.api.ListKeysV1KeysGet(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

func (r KeysResource) Update(ctx context.Context, keyID string, req client.UpdateKeyRequest) (*client.KeyInfo, *http.Response, error) {
	return r.api.UpdateKeyV1KeysKeyIdPatch(ctx, keyID).UpdateKeyRequest(req).Execute()
}

func (r KeysResource) Delete(ctx context.Context, keyID string) (*http.Response, error) {
	return r.api.DeleteKeyV1KeysKeyIdDelete(ctx, keyID).Execute()
}

// UsersResource exposes ergonomic aliases for the users management endpoints.
type UsersResource struct{ api *client.UsersAPIService }

func (r UsersResource) Create(ctx context.Context, req client.CreateUserRequest) (*client.UserResponse, *http.Response, error) {
	return r.api.CreateUserV1UsersPost(ctx).CreateUserRequest(req).Execute()
}

func (r UsersResource) Get(ctx context.Context, userID string) (*client.UserResponse, *http.Response, error) {
	return r.api.GetUserV1UsersUserIdGet(ctx, userID).Execute()
}

func (r UsersResource) List(ctx context.Context, skip, limit *int32) ([]client.UserResponse, *http.Response, error) {
	req := r.api.ListUsersV1UsersGet(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

func (r UsersResource) Update(ctx context.Context, userID string, req client.UpdateUserRequest) (*client.UserResponse, *http.Response, error) {
	return r.api.UpdateUserV1UsersUserIdPatch(ctx, userID).UpdateUserRequest(req).Execute()
}

func (r UsersResource) Delete(ctx context.Context, userID string) (*http.Response, error) {
	return r.api.DeleteUserV1UsersUserIdDelete(ctx, userID).Execute()
}

func (r UsersResource) GetUsage(ctx context.Context, userID string, skip, limit *int32) ([]client.UsageLogResponse, *http.Response, error) {
	req := r.api.GetUserUsageV1UsersUserIdUsageGet(ctx, userID)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

// BudgetsResource exposes ergonomic aliases for the budgets management endpoints.
type BudgetsResource struct{ api *client.BudgetsAPIService }

func (r BudgetsResource) Create(ctx context.Context, req client.CreateBudgetRequest) (*client.BudgetResponse, *http.Response, error) {
	return r.api.CreateBudgetV1BudgetsPost(ctx).CreateBudgetRequest(req).Execute()
}

func (r BudgetsResource) Get(ctx context.Context, budgetID string) (*client.BudgetResponse, *http.Response, error) {
	return r.api.GetBudgetV1BudgetsBudgetIdGet(ctx, budgetID).Execute()
}

func (r BudgetsResource) List(ctx context.Context, skip, limit *int32) ([]client.BudgetResponse, *http.Response, error) {
	req := r.api.ListBudgetsV1BudgetsGet(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

func (r BudgetsResource) Update(ctx context.Context, budgetID string, req client.UpdateBudgetRequest) (*client.BudgetResponse, *http.Response, error) {
	return r.api.UpdateBudgetV1BudgetsBudgetIdPatch(ctx, budgetID).UpdateBudgetRequest(req).Execute()
}

func (r BudgetsResource) Delete(ctx context.Context, budgetID string) (*http.Response, error) {
	return r.api.DeleteBudgetV1BudgetsBudgetIdDelete(ctx, budgetID).Execute()
}

// PricingResource exposes ergonomic aliases for the model-pricing management endpoints.
type PricingResource struct{ api *client.PricingAPIService }

func (r PricingResource) List(ctx context.Context, skip, limit *int32) ([]client.PricingResponse, *http.Response, error) {
	req := r.api.ListPricingV1PricingGet(ctx)
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

func (r PricingResource) Get(ctx context.Context, modelKey string, asOf *time.Time) (*client.PricingResponse, *http.Response, error) {
	req := r.api.GetPricingV1PricingModelKeyGet(ctx, modelKey)
	if asOf != nil {
		req = req.AsOf(*asOf)
	}
	return req.Execute()
}

func (r PricingResource) Set(ctx context.Context, req client.SetPricingRequest) (*client.PricingResponse, *http.Response, error) {
	return r.api.SetPricingV1PricingPost(ctx).SetPricingRequest(req).Execute()
}

func (r PricingResource) Delete(ctx context.Context, modelKey string, effectiveAt *time.Time) (*http.Response, error) {
	req := r.api.DeletePricingV1PricingModelKeyDelete(ctx, modelKey)
	if effectiveAt != nil {
		req = req.EffectiveAt(*effectiveAt)
	}
	return req.Execute()
}

func (r PricingResource) GetHistory(ctx context.Context, modelKey string) ([]client.PricingResponse, *http.Response, error) {
	return r.api.GetPricingHistoryV1PricingModelKeyHistoryGet(ctx, modelKey).Execute()
}

// UsageResource exposes ergonomic aliases for the usage-log management endpoints.
type UsageResource struct{ api *client.UsageAPIService }

func (r UsageResource) List(ctx context.Context, startDate, endDate *time.Time, userID *string, skip, limit *int32) ([]client.UsageEntry, *http.Response, error) {
	req := r.api.ListUsageV1UsageGet(ctx)
	if startDate != nil {
		req = req.StartDate(*startDate)
	}
	if endDate != nil {
		req = req.EndDate(*endDate)
	}
	if userID != nil {
		req = req.UserId(*userID)
	}
	if skip != nil {
		req = req.Skip(*skip)
	}
	if limit != nil {
		req = req.Limit(*limit)
	}
	return req.Execute()
}

// ControlPlane returns a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage).
//
// These endpoints authenticate with Authorization: Bearer <admin/master key>,
// which is distinct from the Otari-Key header used for inference. Pass the
// gateway master key (or an admin token). The generated control-plane operation
// paths already include the /v1 prefix, so the configured server URL is the
// gateway root.
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
