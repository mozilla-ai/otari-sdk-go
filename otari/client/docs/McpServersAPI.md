# \McpServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpServersCheckWorkspaceMcpServer**](McpServersAPI.md#McpServersCheckWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/check | Check Workspace Mcp Server
[**McpServersCheckWorkspaceMcpServerConnection**](McpServersAPI.md#McpServersCheckWorkspaceMcpServerConnection) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers/check | Check Workspace Mcp Server Connection
[**McpServersCreateWorkspaceMcpServer**](McpServersAPI.md#McpServersCreateWorkspaceMcpServer) | **Post** /api/v1/workspaces/{workspace_id}/mcp-servers | Create Workspace Mcp Server
[**McpServersDeleteWorkspaceMcpServer**](McpServersAPI.md#McpServersDeleteWorkspaceMcpServer) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Delete Workspace Mcp Server
[**McpServersListWorkspaceMcpServers**](McpServersAPI.md#McpServersListWorkspaceMcpServers) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers | List Workspace Mcp Servers
[**McpServersUpdateWorkspaceMcpServer**](McpServersAPI.md#McpServersUpdateWorkspaceMcpServer) | **Patch** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Update Workspace Mcp Server



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

