# \McpServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpServersAcknowledgeMcpServerRequest**](McpServersAPI.md#McpServersAcknowledgeMcpServerRequest) | **Post** /api/v1/workspaces/{workspace_id}/mcp-server-requests/{request_id}/acknowledge | Acknowledge Mcp Server Request
[**McpServersAddWorkspaceMcpServerGrants**](McpServersAPI.md#McpServersAddWorkspaceMcpServerGrants) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants | Add Workspace Mcp Server Grants
[**McpServersCheckOrganizationMcpServer**](McpServersAPI.md#McpServersCheckOrganizationMcpServer) | **Post** /api/v1/organizations/me/mcp-servers/{server_id}/check | Check Organization Mcp Server
[**McpServersCheckOrganizationMcpServerConnection**](McpServersAPI.md#McpServersCheckOrganizationMcpServerConnection) | **Post** /api/v1/organizations/me/mcp-servers/check | Check Organization Mcp Server Connection
[**McpServersCheckWorkspaceMcpServer**](McpServersAPI.md#McpServersCheckWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/check | Check Workspace Mcp Server
[**McpServersCheckWorkspaceMcpServerAccess**](McpServersAPI.md#McpServersCheckWorkspaceMcpServerAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/check | Check Workspace Mcp Server Access
[**McpServersCheckWorkspaceMcpServerConnection**](McpServersAPI.md#McpServersCheckWorkspaceMcpServerConnection) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/check | Check Workspace Mcp Server Connection
[**McpServersClearOrganizationMcpServerWorkspaceLimit**](McpServersAPI.md#McpServersClearOrganizationMcpServerWorkspaceLimit) | **Delete** /api/v1/organizations/me/mcp-servers/{server_id}/workspaces/{workspace_id} | Clear Organization Mcp Server Workspace Limit
[**McpServersClearWorkspaceMcpServerOverride**](McpServersAPI.md#McpServersClearWorkspaceMcpServerOverride) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/override | Clear Workspace Mcp Server Override
[**McpServersCreateOrganizationMcpServer**](McpServersAPI.md#McpServersCreateOrganizationMcpServer) | **Post** /api/v1/organizations/me/mcp-servers | Create Organization Mcp Server
[**McpServersCreateWorkspaceMcpServer**](McpServersAPI.md#McpServersCreateWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers | Create Workspace Mcp Server
[**McpServersDeclineMcpServerRequests**](McpServersAPI.md#McpServersDeclineMcpServerRequests) | **Post** /api/v1/workspaces/{workspace_id}/mcp-server-requests/decline | Decline Mcp Server Requests
[**McpServersDeleteOrganizationMcpServer**](McpServersAPI.md#McpServersDeleteOrganizationMcpServer) | **Delete** /api/v1/organizations/me/mcp-servers/{server_id} | Delete Organization Mcp Server
[**McpServersDeleteWorkspaceMcpServer**](McpServersAPI.md#McpServersDeleteWorkspaceMcpServer) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Delete Workspace Mcp Server
[**McpServersDeleteWorkspaceMcpServerGrant**](McpServersAPI.md#McpServersDeleteWorkspaceMcpServerGrant) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants/{grant_id} | Delete Workspace Mcp Server Grant
[**McpServersEndOrganizationMcpServerSharedAccount**](McpServersAPI.md#McpServersEndOrganizationMcpServerSharedAccount) | **Delete** /api/v1/organizations/me/mcp-servers/{server_id}/shared-account | End Organization Mcp Server Shared Account
[**McpServersEndWorkspaceMcpServerSharedAccount**](McpServersAPI.md#McpServersEndWorkspaceMcpServerSharedAccount) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/shared-account | End Workspace Mcp Server Shared Account
[**McpServersGetMyWorkspaceMcpServerTools**](McpServersAPI.md#McpServersGetMyWorkspaceMcpServerTools) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/my-tools | Get My Workspace Mcp Server Tools
[**McpServersGetOrganizationMcpServer**](McpServersAPI.md#McpServersGetOrganizationMcpServer) | **Get** /api/v1/organizations/me/mcp-servers/{server_id} | Get Organization Mcp Server
[**McpServersGetWorkspaceMcpKeyAccess**](McpServersAPI.md#McpServersGetWorkspaceMcpKeyAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/keys/{api_key_id}/access | Get Workspace Mcp Key Access
[**McpServersGetWorkspaceMcpServer**](McpServersAPI.md#McpServersGetWorkspaceMcpServer) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Get Workspace Mcp Server
[**McpServersGetWorkspaceMcpServerAccess**](McpServersAPI.md#McpServersGetWorkspaceMcpServerAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access | Get Workspace Mcp Server Access
[**McpServersListMcpServerRequesters**](McpServersAPI.md#McpServersListMcpServerRequesters) | **Get** /api/v1/workspaces/{workspace_id}/mcp-server-requests/requesters | List Mcp Server Requesters
[**McpServersListMcpServerRequests**](McpServersAPI.md#McpServersListMcpServerRequests) | **Get** /api/v1/workspaces/{workspace_id}/mcp-server-requests | List Mcp Server Requests
[**McpServersListMyMcpServerRequests**](McpServersAPI.md#McpServersListMyMcpServerRequests) | **Get** /api/v1/workspaces/{workspace_id}/mcp-server-requests/mine | List My Mcp Server Requests
[**McpServersListOrganizationMcpServerConnections**](McpServersAPI.md#McpServersListOrganizationMcpServerConnections) | **Get** /api/v1/organizations/me/mcp-servers/{server_id}/connections | List Organization Mcp Server Connections
[**McpServersListOrganizationMcpServerCopies**](McpServersAPI.md#McpServersListOrganizationMcpServerCopies) | **Post** /api/v1/organizations/me/mcp-servers/same-url | List Organization Mcp Server Copies
[**McpServersListOrganizationMcpServerWorkspaces**](McpServersAPI.md#McpServersListOrganizationMcpServerWorkspaces) | **Get** /api/v1/organizations/me/mcp-servers/{server_id}/workspaces | List Organization Mcp Server Workspaces
[**McpServersListOrganizationMcpServers**](McpServersAPI.md#McpServersListOrganizationMcpServers) | **Get** /api/v1/organizations/me/mcp-servers | List Organization Mcp Servers
[**McpServersListWorkspaceMcpServerLogos**](McpServersAPI.md#McpServersListWorkspaceMcpServerLogos) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/logos | List Workspace Mcp Server Logos
[**McpServersListWorkspaceMcpServers**](McpServersAPI.md#McpServersListWorkspaceMcpServers) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers | List Workspace Mcp Servers
[**McpServersPromoteWorkspaceMcpServer**](McpServersAPI.md#McpServersPromoteWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/promote | Promote Workspace Mcp Server
[**McpServersRemoveOrganizationMcpServerConnection**](McpServersAPI.md#McpServersRemoveOrganizationMcpServerConnection) | **Delete** /api/v1/organizations/me/mcp-servers/{server_id}/connections/{user_id} | Remove Organization Mcp Server Connection
[**McpServersRemoveWorkspaceMcpServerLogo**](McpServersAPI.md#McpServersRemoveWorkspaceMcpServerLogo) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/logo | Remove Workspace Mcp Server Logo
[**McpServersRequestMcpServer**](McpServersAPI.md#McpServersRequestMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-server-requests | Request Mcp Server
[**McpServersSetOrganizationMcpServerWorkspaceLimit**](McpServersAPI.md#McpServersSetOrganizationMcpServerWorkspaceLimit) | **Put** /api/v1/organizations/me/mcp-servers/{server_id}/workspaces/{workspace_id} | Set Organization Mcp Server Workspace Limit
[**McpServersSetWorkspaceMcpServerLogo**](McpServersAPI.md#McpServersSetWorkspaceMcpServerLogo) | **Put** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/logo | Set Workspace Mcp Server Logo
[**McpServersSetWorkspaceMcpServerOverride**](McpServersAPI.md#McpServersSetWorkspaceMcpServerOverride) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/override | Set Workspace Mcp Server Override
[**McpServersStartOrganizationMcpServerCheckSignIn**](McpServersAPI.md#McpServersStartOrganizationMcpServerCheckSignIn) | **Post** /api/v1/organizations/me/mcp-servers/{server_id}/check/sign-in | Start Organization Mcp Server Check Sign In
[**McpServersStartOrganizationMcpServerSharedSignIn**](McpServersAPI.md#McpServersStartOrganizationMcpServerSharedSignIn) | **Post** /api/v1/organizations/me/mcp-servers/{server_id}/shared-account | Start Organization Mcp Server Shared Sign In
[**McpServersStartWorkspaceMcpServerCheckSignIn**](McpServersAPI.md#McpServersStartWorkspaceMcpServerCheckSignIn) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/check/sign-in | Start Workspace Mcp Server Check Sign In
[**McpServersStartWorkspaceMcpServerSharedSignIn**](McpServersAPI.md#McpServersStartWorkspaceMcpServerSharedSignIn) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/shared-account | Start Workspace Mcp Server Shared Sign In
[**McpServersUpdateOrganizationMcpServer**](McpServersAPI.md#McpServersUpdateOrganizationMcpServer) | **Patch** /api/v1/organizations/me/mcp-servers/{server_id} | Update Organization Mcp Server
[**McpServersUpdateWorkspaceMcpServer**](McpServersAPI.md#McpServersUpdateWorkspaceMcpServer) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Update Workspace Mcp Server
[**McpServersUpdateWorkspaceMcpServerGrant**](McpServersAPI.md#McpServersUpdateWorkspaceMcpServerGrant) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants/{grant_id} | Update Workspace Mcp Server Grant
[**McpServersWithdrawMcpServerRequest**](McpServersAPI.md#McpServersWithdrawMcpServerRequest) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-server-requests/{request_id} | Withdraw Mcp Server Request



## McpServersAcknowledgeMcpServerRequest

> McpServersAcknowledgeMcpServerRequest(ctx, workspaceId, requestId).Execute()

Acknowledge Mcp Server Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.McpServersAPI.McpServersAcknowledgeMcpServerRequest(context.Background(), workspaceId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersAcknowledgeMcpServerRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersAcknowledgeMcpServerRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersAddWorkspaceMcpServerGrants

> WorkspaceMcpServerAccess McpServersAddWorkspaceMcpServerGrants(ctx, workspaceId, serverId).WorkspaceMcpServerGrantsCreate(workspaceMcpServerGrantsCreate).Execute()

Add Workspace Mcp Server Grants



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerGrantsCreate := *openapiclient.NewWorkspaceMcpServerGrantsCreate([]string{"Tools_example"}) // WorkspaceMcpServerGrantsCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersAddWorkspaceMcpServerGrants(context.Background(), workspaceId, serverId).WorkspaceMcpServerGrantsCreate(workspaceMcpServerGrantsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersAddWorkspaceMcpServerGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersAddWorkspaceMcpServerGrants`: WorkspaceMcpServerAccess
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersAddWorkspaceMcpServerGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersAddWorkspaceMcpServerGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceMcpServerGrantsCreate** | [**WorkspaceMcpServerGrantsCreate**](WorkspaceMcpServerGrantsCreate.md) |  | 

### Return type

[**WorkspaceMcpServerAccess**](WorkspaceMcpServerAccess.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCheckOrganizationMcpServer

> WorkspaceMcpServerPublic McpServersCheckOrganizationMcpServer(ctx, serverId).Execute()

Check Organization Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCheckOrganizationMcpServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCheckOrganizationMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCheckOrganizationMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCheckOrganizationMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCheckOrganizationMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCheckOrganizationMcpServerConnection

> WorkspaceMcpServerHealth McpServersCheckOrganizationMcpServerConnection(ctx).WorkspaceMcpServerCheckRequest(workspaceMcpServerCheckRequest).Execute()

Check Organization Mcp Server Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceMcpServerCheckRequest := *openapiclient.NewWorkspaceMcpServerCheckRequest("Url_example") // WorkspaceMcpServerCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCheckOrganizationMcpServerConnection(context.Background()).WorkspaceMcpServerCheckRequest(workspaceMcpServerCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCheckOrganizationMcpServerConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCheckOrganizationMcpServerConnection`: WorkspaceMcpServerHealth
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCheckOrganizationMcpServerConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCheckOrganizationMcpServerConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceMcpServerCheckRequest** | [**WorkspaceMcpServerCheckRequest**](WorkspaceMcpServerCheckRequest.md) |  | 

### Return type

[**WorkspaceMcpServerHealth**](WorkspaceMcpServerHealth.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCheckWorkspaceMcpServer

> WorkspaceMcpServerPublic McpServersCheckWorkspaceMcpServer(ctx, workspaceId, serverId).Execute()

Check Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCheckWorkspaceMcpServer(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCheckWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCheckWorkspaceMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCheckWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCheckWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCheckWorkspaceMcpServerAccess

> WorkspaceMcpServerAccessCheck McpServersCheckWorkspaceMcpServerAccess(ctx, workspaceId, serverId).ApiKeyId(apiKeyId).Tool(tool).Execute()

Check Workspace Mcp Server Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiKeyId := "apiKeyId_example" // string | The key to test
	tool := "tool_example" // string | The tool name to test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCheckWorkspaceMcpServerAccess(context.Background(), workspaceId, serverId).ApiKeyId(apiKeyId).Tool(tool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCheckWorkspaceMcpServerAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCheckWorkspaceMcpServerAccess`: WorkspaceMcpServerAccessCheck
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCheckWorkspaceMcpServerAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCheckWorkspaceMcpServerAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeyId** | **string** | The key to test | 
 **tool** | **string** | The tool name to test | 

### Return type

[**WorkspaceMcpServerAccessCheck**](WorkspaceMcpServerAccessCheck.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCheckWorkspaceMcpServerConnection

> WorkspaceMcpServerHealth McpServersCheckWorkspaceMcpServerConnection(ctx, workspaceId).WorkspaceMcpServerCheckRequest(workspaceMcpServerCheckRequest).Execute()

Check Workspace Mcp Server Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerCheckRequest := *openapiclient.NewWorkspaceMcpServerCheckRequest("Url_example") // WorkspaceMcpServerCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCheckWorkspaceMcpServerConnection(context.Background(), workspaceId).WorkspaceMcpServerCheckRequest(workspaceMcpServerCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCheckWorkspaceMcpServerConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCheckWorkspaceMcpServerConnection`: WorkspaceMcpServerHealth
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCheckWorkspaceMcpServerConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCheckWorkspaceMcpServerConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceMcpServerCheckRequest** | [**WorkspaceMcpServerCheckRequest**](WorkspaceMcpServerCheckRequest.md) |  | 

### Return type

[**WorkspaceMcpServerHealth**](WorkspaceMcpServerHealth.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersClearOrganizationMcpServerWorkspaceLimit

> OrganizationMcpServerWorkspace McpServersClearOrganizationMcpServerWorkspaceLimit(ctx, serverId, workspaceId).Execute()

Clear Organization Mcp Server Workspace Limit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersClearOrganizationMcpServerWorkspaceLimit(context.Background(), serverId, workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersClearOrganizationMcpServerWorkspaceLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersClearOrganizationMcpServerWorkspaceLimit`: OrganizationMcpServerWorkspace
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersClearOrganizationMcpServerWorkspaceLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersClearOrganizationMcpServerWorkspaceLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMcpServerWorkspace**](OrganizationMcpServerWorkspace.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersClearWorkspaceMcpServerOverride

> WorkspaceMcpServerPublic McpServersClearWorkspaceMcpServerOverride(ctx, workspaceId, serverId).Execute()

Clear Workspace Mcp Server Override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersClearWorkspaceMcpServerOverride(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersClearWorkspaceMcpServerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersClearWorkspaceMcpServerOverride`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersClearWorkspaceMcpServerOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersClearWorkspaceMcpServerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCreateOrganizationMcpServer

> WorkspaceMcpServerPublic McpServersCreateOrganizationMcpServer(ctx).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()

Create Organization Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceMcpServerCreate := *openapiclient.NewWorkspaceMcpServerCreate("Name_example", "Url_example") // WorkspaceMcpServerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCreateOrganizationMcpServer(context.Background()).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCreateOrganizationMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCreateOrganizationMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCreateOrganizationMcpServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCreateOrganizationMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceMcpServerCreate** | [**WorkspaceMcpServerCreate**](WorkspaceMcpServerCreate.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersCreateWorkspaceMcpServer

> WorkspaceMcpServerPublic McpServersCreateWorkspaceMcpServer(ctx, workspaceId).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()

Create Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerCreate := *openapiclient.NewWorkspaceMcpServerCreate("Name_example", "Url_example") // WorkspaceMcpServerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersCreateWorkspaceMcpServer(context.Background(), workspaceId).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersCreateWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersCreateWorkspaceMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersCreateWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersCreateWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceMcpServerCreate** | [**WorkspaceMcpServerCreate**](WorkspaceMcpServerCreate.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersDeclineMcpServerRequests

> McpServerRequestsDeclined McpServersDeclineMcpServerRequests(ctx, workspaceId).McpServerRequestDecline(mcpServerRequestDecline).Execute()

Decline Mcp Server Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpServerRequestDecline := *openapiclient.NewMcpServerRequestDecline("CatalogName_example") // McpServerRequestDecline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersDeclineMcpServerRequests(context.Background(), workspaceId).McpServerRequestDecline(mcpServerRequestDecline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersDeclineMcpServerRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersDeclineMcpServerRequests`: McpServerRequestsDeclined
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersDeclineMcpServerRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersDeclineMcpServerRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpServerRequestDecline** | [**McpServerRequestDecline**](McpServerRequestDecline.md) |  | 

### Return type

[**McpServerRequestsDeclined**](McpServerRequestsDeclined.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersDeleteOrganizationMcpServer

> McpServersDeleteOrganizationMcpServer(ctx, serverId).Execute()

Delete Organization Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.McpServersAPI.McpServersDeleteOrganizationMcpServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersDeleteOrganizationMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersDeleteOrganizationMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersDeleteWorkspaceMcpServer

> Message McpServersDeleteWorkspaceMcpServer(ctx, workspaceId, serverId).Execute()

Delete Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersDeleteWorkspaceMcpServer(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersDeleteWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersDeleteWorkspaceMcpServer`: Message
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersDeleteWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersDeleteWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersDeleteWorkspaceMcpServerGrant

> WorkspaceMcpServerAccess McpServersDeleteWorkspaceMcpServerGrant(ctx, workspaceId, serverId, grantId).Execute()

Delete Workspace Mcp Server Grant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	grantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersDeleteWorkspaceMcpServerGrant(context.Background(), workspaceId, serverId, grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersDeleteWorkspaceMcpServerGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersDeleteWorkspaceMcpServerGrant`: WorkspaceMcpServerAccess
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersDeleteWorkspaceMcpServerGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersDeleteWorkspaceMcpServerGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceMcpServerAccess**](WorkspaceMcpServerAccess.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersEndOrganizationMcpServerSharedAccount

> WorkspaceMcpServerPublic McpServersEndOrganizationMcpServerSharedAccount(ctx, serverId).Execute()

End Organization Mcp Server Shared Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersEndOrganizationMcpServerSharedAccount(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersEndOrganizationMcpServerSharedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersEndOrganizationMcpServerSharedAccount`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersEndOrganizationMcpServerSharedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersEndOrganizationMcpServerSharedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersEndWorkspaceMcpServerSharedAccount

> WorkspaceMcpServerPublic McpServersEndWorkspaceMcpServerSharedAccount(ctx, workspaceId, serverId).Execute()

End Workspace Mcp Server Shared Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersEndWorkspaceMcpServerSharedAccount(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersEndWorkspaceMcpServerSharedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersEndWorkspaceMcpServerSharedAccount`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersEndWorkspaceMcpServerSharedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersEndWorkspaceMcpServerSharedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersGetMyWorkspaceMcpServerTools

> WorkspaceMcpServerMyTools McpServersGetMyWorkspaceMcpServerTools(ctx, workspaceId, serverId).Execute()

Get My Workspace Mcp Server Tools



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersGetMyWorkspaceMcpServerTools(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersGetMyWorkspaceMcpServerTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersGetMyWorkspaceMcpServerTools`: WorkspaceMcpServerMyTools
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersGetMyWorkspaceMcpServerTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersGetMyWorkspaceMcpServerToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerMyTools**](WorkspaceMcpServerMyTools.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersGetOrganizationMcpServer

> WorkspaceMcpServerPublic McpServersGetOrganizationMcpServer(ctx, serverId).Execute()

Get Organization Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersGetOrganizationMcpServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersGetOrganizationMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersGetOrganizationMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersGetOrganizationMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersGetOrganizationMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersGetWorkspaceMcpKeyAccess

> []WorkspaceMcpServerKeyAccess McpServersGetWorkspaceMcpKeyAccess(ctx, workspaceId, apiKeyId).Execute()

Get Workspace Mcp Key Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersGetWorkspaceMcpKeyAccess(context.Background(), workspaceId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersGetWorkspaceMcpKeyAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersGetWorkspaceMcpKeyAccess`: []WorkspaceMcpServerKeyAccess
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersGetWorkspaceMcpKeyAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersGetWorkspaceMcpKeyAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WorkspaceMcpServerKeyAccess**](WorkspaceMcpServerKeyAccess.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersGetWorkspaceMcpServer

> WorkspaceMcpServerPublic McpServersGetWorkspaceMcpServer(ctx, workspaceId, serverId).Execute()

Get Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersGetWorkspaceMcpServer(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersGetWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersGetWorkspaceMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersGetWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersGetWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersGetWorkspaceMcpServerAccess

> WorkspaceMcpServerAccess McpServersGetWorkspaceMcpServerAccess(ctx, workspaceId, serverId).Execute()

Get Workspace Mcp Server Access



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersGetWorkspaceMcpServerAccess(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersGetWorkspaceMcpServerAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersGetWorkspaceMcpServerAccess`: WorkspaceMcpServerAccess
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersGetWorkspaceMcpServerAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersGetWorkspaceMcpServerAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerAccess**](WorkspaceMcpServerAccess.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListMcpServerRequesters

> McpServerRequestersPublic McpServersListMcpServerRequesters(ctx, workspaceId).CatalogName(catalogName).Skip(skip).Limit(limit).Execute()

List Mcp Server Requesters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	catalogName := "catalogName_example" // string | The catalog server asked for
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListMcpServerRequesters(context.Background(), workspaceId).CatalogName(catalogName).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListMcpServerRequesters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListMcpServerRequesters`: McpServerRequestersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListMcpServerRequesters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListMcpServerRequestersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catalogName** | **string** | The catalog server asked for | 
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**McpServerRequestersPublic**](McpServerRequestersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListMcpServerRequests

> McpServerRequestGroupsPublic McpServersListMcpServerRequests(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Mcp Server Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Number of catalog servers to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of catalog servers to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListMcpServerRequests(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListMcpServerRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListMcpServerRequests`: McpServerRequestGroupsPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListMcpServerRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListMcpServerRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of catalog servers to skip | [default to 0]
 **limit** | **int32** | Maximum number of catalog servers to return | [default to 100]

### Return type

[**McpServerRequestGroupsPublic**](McpServerRequestGroupsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListMyMcpServerRequests

> McpServerRequestsPublic McpServersListMyMcpServerRequests(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List My Mcp Server Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListMyMcpServerRequests(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListMyMcpServerRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListMyMcpServerRequests`: McpServerRequestsPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListMyMcpServerRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListMyMcpServerRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 50]

### Return type

[**McpServerRequestsPublic**](McpServerRequestsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListOrganizationMcpServerConnections

> McpServerConnectionsPublic McpServersListOrganizationMcpServerConnections(ctx, serverId).Skip(skip).Limit(limit).Execute()

List Organization Mcp Server Connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListOrganizationMcpServerConnections(context.Background(), serverId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListOrganizationMcpServerConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListOrganizationMcpServerConnections`: McpServerConnectionsPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListOrganizationMcpServerConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListOrganizationMcpServerConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**McpServerConnectionsPublic**](McpServerConnectionsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListOrganizationMcpServerCopies

> McpServerFoldCandidates McpServersListOrganizationMcpServerCopies(ctx).McpServerFoldLookup(mcpServerFoldLookup).Execute()

List Organization Mcp Server Copies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mcpServerFoldLookup := *openapiclient.NewMcpServerFoldLookup() // McpServerFoldLookup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListOrganizationMcpServerCopies(context.Background()).McpServerFoldLookup(mcpServerFoldLookup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListOrganizationMcpServerCopies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListOrganizationMcpServerCopies`: McpServerFoldCandidates
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListOrganizationMcpServerCopies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListOrganizationMcpServerCopiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcpServerFoldLookup** | [**McpServerFoldLookup**](McpServerFoldLookup.md) |  | 

### Return type

[**McpServerFoldCandidates**](McpServerFoldCandidates.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListOrganizationMcpServerWorkspaces

> OrganizationMcpServerWorkspaces McpServersListOrganizationMcpServerWorkspaces(ctx, serverId).Execute()

List Organization Mcp Server Workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListOrganizationMcpServerWorkspaces(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListOrganizationMcpServerWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListOrganizationMcpServerWorkspaces`: OrganizationMcpServerWorkspaces
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListOrganizationMcpServerWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListOrganizationMcpServerWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMcpServerWorkspaces**](OrganizationMcpServerWorkspaces.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListOrganizationMcpServers

> WorkspaceMcpServersPublic McpServersListOrganizationMcpServers(ctx).Skip(skip).Limit(limit).Execute()

List Organization Mcp Servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListOrganizationMcpServers(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListOrganizationMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListOrganizationMcpServers`: WorkspaceMcpServersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListOrganizationMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListOrganizationMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceMcpServersPublic**](WorkspaceMcpServersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListWorkspaceMcpServerLogos

> WorkspaceMcpServerLogosPublic McpServersListWorkspaceMcpServerLogos(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Workspace Mcp Server Logos



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListWorkspaceMcpServerLogos(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListWorkspaceMcpServerLogos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListWorkspaceMcpServerLogos`: WorkspaceMcpServerLogosPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListWorkspaceMcpServerLogos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListWorkspaceMcpServerLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceMcpServerLogosPublic**](WorkspaceMcpServerLogosPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersListWorkspaceMcpServers

> WorkspaceMcpServersPublic McpServersListWorkspaceMcpServers(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Workspace Mcp Servers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersListWorkspaceMcpServers(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersListWorkspaceMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersListWorkspaceMcpServers`: WorkspaceMcpServersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersListWorkspaceMcpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersListWorkspaceMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceMcpServersPublic**](WorkspaceMcpServersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersPromoteWorkspaceMcpServer

> WorkspaceMcpServerPublic McpServersPromoteWorkspaceMcpServer(ctx, workspaceId, serverId).McpServerPromote(mcpServerPromote).Execute()

Promote Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpServerPromote := *openapiclient.NewMcpServerPromote() // McpServerPromote | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersPromoteWorkspaceMcpServer(context.Background(), workspaceId, serverId).McpServerPromote(mcpServerPromote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersPromoteWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersPromoteWorkspaceMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersPromoteWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersPromoteWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mcpServerPromote** | [**McpServerPromote**](McpServerPromote.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersRemoveOrganizationMcpServerConnection

> McpServersRemoveOrganizationMcpServerConnection(ctx, serverId, userId).Execute()

Remove Organization Mcp Server Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.McpServersAPI.McpServersRemoveOrganizationMcpServerConnection(context.Background(), serverId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersRemoveOrganizationMcpServerConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersRemoveOrganizationMcpServerConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersRemoveWorkspaceMcpServerLogo

> WorkspaceMcpServerLogoPublic McpServersRemoveWorkspaceMcpServerLogo(ctx, workspaceId, serverId).Execute()

Remove Workspace Mcp Server Logo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersRemoveWorkspaceMcpServerLogo(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersRemoveWorkspaceMcpServerLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersRemoveWorkspaceMcpServerLogo`: WorkspaceMcpServerLogoPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersRemoveWorkspaceMcpServerLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersRemoveWorkspaceMcpServerLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceMcpServerLogoPublic**](WorkspaceMcpServerLogoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersRequestMcpServer

> McpServerRequestPublic McpServersRequestMcpServer(ctx, workspaceId).McpServerRequestCreate(mcpServerRequestCreate).Execute()

Request Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpServerRequestCreate := *openapiclient.NewMcpServerRequestCreate("CatalogName_example") // McpServerRequestCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersRequestMcpServer(context.Background(), workspaceId).McpServerRequestCreate(mcpServerRequestCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersRequestMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersRequestMcpServer`: McpServerRequestPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersRequestMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersRequestMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpServerRequestCreate** | [**McpServerRequestCreate**](McpServerRequestCreate.md) |  | 

### Return type

[**McpServerRequestPublic**](McpServerRequestPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersSetOrganizationMcpServerWorkspaceLimit

> OrganizationMcpServerWorkspace McpServersSetOrganizationMcpServerWorkspaceLimit(ctx, serverId, workspaceId).OrganizationMcpServerLimit(organizationMcpServerLimit).Execute()

Set Organization Mcp Server Workspace Limit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationMcpServerLimit := *openapiclient.NewOrganizationMcpServerLimit([]string{"Tools_example"}) // OrganizationMcpServerLimit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersSetOrganizationMcpServerWorkspaceLimit(context.Background(), serverId, workspaceId).OrganizationMcpServerLimit(organizationMcpServerLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersSetOrganizationMcpServerWorkspaceLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersSetOrganizationMcpServerWorkspaceLimit`: OrganizationMcpServerWorkspace
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersSetOrganizationMcpServerWorkspaceLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersSetOrganizationMcpServerWorkspaceLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationMcpServerLimit** | [**OrganizationMcpServerLimit**](OrganizationMcpServerLimit.md) |  | 

### Return type

[**OrganizationMcpServerWorkspace**](OrganizationMcpServerWorkspace.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersSetWorkspaceMcpServerLogo

> WorkspaceMcpServerLogoPublic McpServersSetWorkspaceMcpServerLogo(ctx, workspaceId, serverId).Tile(tile).File(file).Execute()

Set Workspace Mcp Server Logo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tile := "tile_example" // string | Whether the logo sits on a tile: where it would fade out, always, or never.
	file := "file_example" // string | A PNG of at most 512 KB to show instead of the server's own logo. Leave out to keep it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersSetWorkspaceMcpServerLogo(context.Background(), workspaceId, serverId).Tile(tile).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersSetWorkspaceMcpServerLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersSetWorkspaceMcpServerLogo`: WorkspaceMcpServerLogoPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersSetWorkspaceMcpServerLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersSetWorkspaceMcpServerLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tile** | **string** | Whether the logo sits on a tile: where it would fade out, always, or never. | 
 **file** | **string** | A PNG of at most 512 KB to show instead of the server&#39;s own logo. Leave out to keep it. | 

### Return type

[**WorkspaceMcpServerLogoPublic**](WorkspaceMcpServerLogoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersSetWorkspaceMcpServerOverride

> WorkspaceMcpServerPublic McpServersSetWorkspaceMcpServerOverride(ctx, workspaceId, serverId).WorkspaceMcpServerSettingsUpdate(workspaceMcpServerSettingsUpdate).Execute()

Set Workspace Mcp Server Override



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerSettingsUpdate := *openapiclient.NewWorkspaceMcpServerSettingsUpdate() // WorkspaceMcpServerSettingsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersSetWorkspaceMcpServerOverride(context.Background(), workspaceId, serverId).WorkspaceMcpServerSettingsUpdate(workspaceMcpServerSettingsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersSetWorkspaceMcpServerOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersSetWorkspaceMcpServerOverride`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersSetWorkspaceMcpServerOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersSetWorkspaceMcpServerOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceMcpServerSettingsUpdate** | [**WorkspaceMcpServerSettingsUpdate**](WorkspaceMcpServerSettingsUpdate.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersStartOrganizationMcpServerCheckSignIn

> McpConnectionStarted McpServersStartOrganizationMcpServerCheckSignIn(ctx, serverId).McpConnectionStart(mcpConnectionStart).Execute()

Start Organization Mcp Server Check Sign In



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpConnectionStart := *openapiclient.NewMcpConnectionStart() // McpConnectionStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersStartOrganizationMcpServerCheckSignIn(context.Background(), serverId).McpConnectionStart(mcpConnectionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersStartOrganizationMcpServerCheckSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersStartOrganizationMcpServerCheckSignIn`: McpConnectionStarted
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersStartOrganizationMcpServerCheckSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersStartOrganizationMcpServerCheckSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpConnectionStart** | [**McpConnectionStart**](McpConnectionStart.md) |  | 

### Return type

[**McpConnectionStarted**](McpConnectionStarted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersStartOrganizationMcpServerSharedSignIn

> McpConnectionStarted McpServersStartOrganizationMcpServerSharedSignIn(ctx, serverId).McpConnectionStart(mcpConnectionStart).Execute()

Start Organization Mcp Server Shared Sign In



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpConnectionStart := *openapiclient.NewMcpConnectionStart() // McpConnectionStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersStartOrganizationMcpServerSharedSignIn(context.Background(), serverId).McpConnectionStart(mcpConnectionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersStartOrganizationMcpServerSharedSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersStartOrganizationMcpServerSharedSignIn`: McpConnectionStarted
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersStartOrganizationMcpServerSharedSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersStartOrganizationMcpServerSharedSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mcpConnectionStart** | [**McpConnectionStart**](McpConnectionStart.md) |  | 

### Return type

[**McpConnectionStarted**](McpConnectionStarted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersStartWorkspaceMcpServerCheckSignIn

> McpConnectionStarted McpServersStartWorkspaceMcpServerCheckSignIn(ctx, workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()

Start Workspace Mcp Server Check Sign In



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpConnectionStart := *openapiclient.NewMcpConnectionStart() // McpConnectionStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersStartWorkspaceMcpServerCheckSignIn(context.Background(), workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersStartWorkspaceMcpServerCheckSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersStartWorkspaceMcpServerCheckSignIn`: McpConnectionStarted
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersStartWorkspaceMcpServerCheckSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersStartWorkspaceMcpServerCheckSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mcpConnectionStart** | [**McpConnectionStart**](McpConnectionStart.md) |  | 

### Return type

[**McpConnectionStarted**](McpConnectionStarted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersStartWorkspaceMcpServerSharedSignIn

> McpConnectionStarted McpServersStartWorkspaceMcpServerSharedSignIn(ctx, workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()

Start Workspace Mcp Server Shared Sign In



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	mcpConnectionStart := *openapiclient.NewMcpConnectionStart() // McpConnectionStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersStartWorkspaceMcpServerSharedSignIn(context.Background(), workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersStartWorkspaceMcpServerSharedSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersStartWorkspaceMcpServerSharedSignIn`: McpConnectionStarted
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersStartWorkspaceMcpServerSharedSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersStartWorkspaceMcpServerSharedSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mcpConnectionStart** | [**McpConnectionStart**](McpConnectionStart.md) |  | 

### Return type

[**McpConnectionStarted**](McpConnectionStarted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersUpdateOrganizationMcpServer

> WorkspaceMcpServerPublic McpServersUpdateOrganizationMcpServer(ctx, serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()

Update Organization Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerUpdate := *openapiclient.NewWorkspaceMcpServerUpdate() // WorkspaceMcpServerUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersUpdateOrganizationMcpServer(context.Background(), serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersUpdateOrganizationMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersUpdateOrganizationMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersUpdateOrganizationMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersUpdateOrganizationMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceMcpServerUpdate** | [**WorkspaceMcpServerUpdate**](WorkspaceMcpServerUpdate.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersUpdateWorkspaceMcpServer

> WorkspaceMcpServerPublic McpServersUpdateWorkspaceMcpServer(ctx, workspaceId, serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()

Update Workspace Mcp Server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerUpdate := *openapiclient.NewWorkspaceMcpServerUpdate() // WorkspaceMcpServerUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersUpdateWorkspaceMcpServer(context.Background(), workspaceId, serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersUpdateWorkspaceMcpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersUpdateWorkspaceMcpServer`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersUpdateWorkspaceMcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersUpdateWorkspaceMcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceMcpServerUpdate** | [**WorkspaceMcpServerUpdate**](WorkspaceMcpServerUpdate.md) |  | 

### Return type

[**WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersUpdateWorkspaceMcpServerGrant

> WorkspaceMcpServerAccess McpServersUpdateWorkspaceMcpServerGrant(ctx, workspaceId, serverId, grantId).WorkspaceMcpServerGrantUpdate(workspaceMcpServerGrantUpdate).Execute()

Update Workspace Mcp Server Grant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	serverId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	grantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceMcpServerGrantUpdate := *openapiclient.NewWorkspaceMcpServerGrantUpdate([]string{"Tools_example"}) // WorkspaceMcpServerGrantUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpServersAPI.McpServersUpdateWorkspaceMcpServerGrant(context.Background(), workspaceId, serverId, grantId).WorkspaceMcpServerGrantUpdate(workspaceMcpServerGrantUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersUpdateWorkspaceMcpServerGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpServersUpdateWorkspaceMcpServerGrant`: WorkspaceMcpServerAccess
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.McpServersUpdateWorkspaceMcpServerGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersUpdateWorkspaceMcpServerGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **workspaceMcpServerGrantUpdate** | [**WorkspaceMcpServerGrantUpdate**](WorkspaceMcpServerGrantUpdate.md) |  | 

### Return type

[**WorkspaceMcpServerAccess**](WorkspaceMcpServerAccess.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpServersWithdrawMcpServerRequest

> McpServersWithdrawMcpServerRequest(ctx, workspaceId, requestId).Execute()

Withdraw Mcp Server Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	requestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.McpServersAPI.McpServersWithdrawMcpServerRequest(context.Background(), workspaceId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.McpServersWithdrawMcpServerRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpServersWithdrawMcpServerRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

