package otari

import (
	"context"

	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// ControlPlaneClient is a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage).
//
// Each resource (Keys, Users, Budgets, Pricing, Usage) exposes ergonomic
// aliases (Create, Get, List, Update, Delete, ...) that return the generated
// request builders under short, generator-name-free names; finish a call with
// the fluent setters and Execute, for example:
//
//	created, _, err := cp.Keys.Create(ctx).
//	    CreateKeyRequest(client.CreateKeyRequest{}).Execute()
//
// The underlying generated *client.APIClient stays reachable via Raw, so the
// full generated surface (generator-named operations, pagination, filters)
// remains available as an escape hatch:
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

func (r KeysResource) Create(ctx context.Context) client.ApiCreateKeyV1KeysPostRequest {
	return r.api.CreateKeyV1KeysPost(ctx)
}

func (r KeysResource) Get(ctx context.Context, keyID string) client.ApiGetKeyV1KeysKeyIdGetRequest {
	return r.api.GetKeyV1KeysKeyIdGet(ctx, keyID)
}

func (r KeysResource) List(ctx context.Context) client.ApiListKeysV1KeysGetRequest {
	return r.api.ListKeysV1KeysGet(ctx)
}

func (r KeysResource) Update(ctx context.Context, keyID string) client.ApiUpdateKeyV1KeysKeyIdPatchRequest {
	return r.api.UpdateKeyV1KeysKeyIdPatch(ctx, keyID)
}

func (r KeysResource) Delete(ctx context.Context, keyID string) client.ApiDeleteKeyV1KeysKeyIdDeleteRequest {
	return r.api.DeleteKeyV1KeysKeyIdDelete(ctx, keyID)
}

// UsersResource exposes ergonomic aliases for the users management endpoints.
type UsersResource struct{ api *client.UsersAPIService }

func (r UsersResource) Create(ctx context.Context) client.ApiCreateUserV1UsersPostRequest {
	return r.api.CreateUserV1UsersPost(ctx)
}

func (r UsersResource) Get(ctx context.Context, userID string) client.ApiGetUserV1UsersUserIdGetRequest {
	return r.api.GetUserV1UsersUserIdGet(ctx, userID)
}

func (r UsersResource) List(ctx context.Context) client.ApiListUsersV1UsersGetRequest {
	return r.api.ListUsersV1UsersGet(ctx)
}

func (r UsersResource) Update(ctx context.Context, userID string) client.ApiUpdateUserV1UsersUserIdPatchRequest {
	return r.api.UpdateUserV1UsersUserIdPatch(ctx, userID)
}

func (r UsersResource) Delete(ctx context.Context, userID string) client.ApiDeleteUserV1UsersUserIdDeleteRequest {
	return r.api.DeleteUserV1UsersUserIdDelete(ctx, userID)
}

func (r UsersResource) GetUsage(ctx context.Context, userID string) client.ApiGetUserUsageV1UsersUserIdUsageGetRequest {
	return r.api.GetUserUsageV1UsersUserIdUsageGet(ctx, userID)
}

// BudgetsResource exposes ergonomic aliases for the budgets management endpoints.
type BudgetsResource struct{ api *client.BudgetsAPIService }

func (r BudgetsResource) Create(ctx context.Context) client.ApiCreateBudgetV1BudgetsPostRequest {
	return r.api.CreateBudgetV1BudgetsPost(ctx)
}

func (r BudgetsResource) Get(ctx context.Context, budgetID string) client.ApiGetBudgetV1BudgetsBudgetIdGetRequest {
	return r.api.GetBudgetV1BudgetsBudgetIdGet(ctx, budgetID)
}

func (r BudgetsResource) List(ctx context.Context) client.ApiListBudgetsV1BudgetsGetRequest {
	return r.api.ListBudgetsV1BudgetsGet(ctx)
}

func (r BudgetsResource) Update(ctx context.Context, budgetID string) client.ApiUpdateBudgetV1BudgetsBudgetIdPatchRequest {
	return r.api.UpdateBudgetV1BudgetsBudgetIdPatch(ctx, budgetID)
}

func (r BudgetsResource) Delete(ctx context.Context, budgetID string) client.ApiDeleteBudgetV1BudgetsBudgetIdDeleteRequest {
	return r.api.DeleteBudgetV1BudgetsBudgetIdDelete(ctx, budgetID)
}

// PricingResource exposes ergonomic aliases for the model-pricing management endpoints.
type PricingResource struct{ api *client.PricingAPIService }

func (r PricingResource) List(ctx context.Context) client.ApiListPricingV1PricingGetRequest {
	return r.api.ListPricingV1PricingGet(ctx)
}

func (r PricingResource) Get(ctx context.Context, modelKey string) client.ApiGetPricingV1PricingModelKeyGetRequest {
	return r.api.GetPricingV1PricingModelKeyGet(ctx, modelKey)
}

func (r PricingResource) Set(ctx context.Context) client.ApiSetPricingV1PricingPostRequest {
	return r.api.SetPricingV1PricingPost(ctx)
}

func (r PricingResource) Delete(ctx context.Context, modelKey string) client.ApiDeletePricingV1PricingModelKeyDeleteRequest {
	return r.api.DeletePricingV1PricingModelKeyDelete(ctx, modelKey)
}

func (r PricingResource) GetHistory(ctx context.Context, modelKey string) client.ApiGetPricingHistoryV1PricingModelKeyHistoryGetRequest {
	return r.api.GetPricingHistoryV1PricingModelKeyHistoryGet(ctx, modelKey)
}

// UsageResource exposes ergonomic aliases for the usage-log management endpoints.
type UsageResource struct{ api *client.UsageAPIService }

func (r UsageResource) List(ctx context.Context) client.ApiListUsageV1UsageGetRequest {
	return r.api.ListUsageV1UsageGet(ctx)
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
//	created, _, err := cp.Keys.Create(ctx).
//	    CreateKeyRequest(client.CreateKeyRequest{}).Execute()
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
