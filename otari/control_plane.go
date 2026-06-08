package otari

import (
	"github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// ControlPlane returns a typed client for the gateway management endpoints
// (API keys, users, budgets, pricing, usage), exposed as the KeysAPI, UsersAPI,
// BudgetsAPI, PricingAPI, and UsageAPI services on the returned client.
//
// These endpoints authenticate with Authorization: Bearer <admin/master key>,
// which is distinct from the Otari-Key header used for inference. Pass the
// gateway master key (or an admin token). The generated control-plane operation
// paths already include the /v1 prefix, so the configured server URL is the
// gateway root.
//
//	cp := client.ControlPlane("gateway-master-key")
//	created, _, err := cp.KeysAPI.CreateKeyV1KeysPost(ctx).
//	    CreateKeyRequest(client.CreateKeyRequest{}).Execute()
func (c *Client) ControlPlane(adminKey string) *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{{URL: c.baseURL}}
	cfg.AddDefaultHeader(authorizationHeader, bearerPrefix+adminKey)
	return client.NewAPIClient(cfg)
}
