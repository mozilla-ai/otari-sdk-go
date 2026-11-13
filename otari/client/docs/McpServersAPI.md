# \McpServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpServersAddWorkspaceMcpServerGrants**](McpServersAPI.md#McpServersAddWorkspaceMcpServerGrants) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants | Add Workspace Mcp Server Grants
[**McpServersCheckWorkspaceMcpServer**](McpServersAPI.md#McpServersCheckWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/check | Check Workspace Mcp Server
[**McpServersCheckWorkspaceMcpServerAccess**](McpServersAPI.md#McpServersCheckWorkspaceMcpServerAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/check | Check Workspace Mcp Server Access
[**McpServersCheckWorkspaceMcpServerConnection**](McpServersAPI.md#McpServersCheckWorkspaceMcpServerConnection) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/check | Check Workspace Mcp Server Connection
[**McpServersCreateWorkspaceMcpServer**](McpServersAPI.md#McpServersCreateWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers | Create Workspace Mcp Server
[**McpServersDeleteWorkspaceMcpServer**](McpServersAPI.md#McpServersDeleteWorkspaceMcpServer) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Delete Workspace Mcp Server
[**McpServersDeleteWorkspaceMcpServerGrant**](McpServersAPI.md#McpServersDeleteWorkspaceMcpServerGrant) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants/{grant_id} | Delete Workspace Mcp Server Grant
[**McpServersGetWorkspaceMcpKeyAccess**](McpServersAPI.md#McpServersGetWorkspaceMcpKeyAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/keys/{api_key_id}/access | Get Workspace Mcp Key Access
[**McpServersGetWorkspaceMcpServer**](McpServersAPI.md#McpServersGetWorkspaceMcpServer) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Get Workspace Mcp Server
[**McpServersGetWorkspaceMcpServerAccess**](McpServersAPI.md#McpServersGetWorkspaceMcpServerAccess) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access | Get Workspace Mcp Server Access
[**McpServersListWorkspaceMcpServers**](McpServersAPI.md#McpServersListWorkspaceMcpServers) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers | List Workspace Mcp Servers
[**McpServersUpdateWorkspaceMcpServer**](McpServersAPI.md#McpServersUpdateWorkspaceMcpServer) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Update Workspace Mcp Server
[**McpServersUpdateWorkspaceMcpServerGrant**](McpServersAPI.md#McpServersUpdateWorkspaceMcpServerGrant) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/access/grants/{grant_id} | Update Workspace Mcp Server Grant



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

