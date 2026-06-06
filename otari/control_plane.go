package otari

import (
	generated "github.com/mozilla-ai/otari-sdk-go/otari/generated"
)

// ControlPlane returns a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage), exposed as the KeysAPI, UsersAPI,
// BudgetsAPI, PricingAPI, and UsageAPI services on the returned client.
//
// These endpoints authenticate with Authorization: Bearer <admin/master key>,
// which is distinct from the Otari-Key header used for inference. Pass the
// gateway master key (or an admin token).
//
//	cp := client.ControlPlane("gateway-master-key")
//	created, _, err := cp.KeysAPI.CreateKeyV1KeysPost(ctx).
//	    CreateKeyRequest(generated.CreateKeyRequest{}).Execute()
func (c *Client) ControlPlane(adminKey string) *generated.APIClient {
	cfg := generated.NewConfiguration()
	cfg.Servers = generated.ServerConfigurations{{URL: c.baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+adminKey)
	return generated.NewAPIClient(cfg)
}
