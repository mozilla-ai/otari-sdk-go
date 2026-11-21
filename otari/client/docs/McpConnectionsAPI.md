# \McpConnectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpConnectionsCompleteMcpConnection**](McpConnectionsAPI.md#McpConnectionsCompleteMcpConnection) | **Post** /api/v1/mcp-connections/complete | Complete Mcp Connection
[**McpConnectionsEndMyMcpConnection**](McpConnectionsAPI.md#McpConnectionsEndMyMcpConnection) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-connections/{server_id} | End My Mcp Connection
[**McpConnectionsListMcpServerConnections**](McpConnectionsAPI.md#McpConnectionsListMcpServerConnections) | **Get** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/connections | List Mcp Server Connections
[**McpConnectionsListMyMcpConnections**](McpConnectionsAPI.md#McpConnectionsListMyMcpConnections) | **Get** /api/v1/workspaces/{workspace_id}/mcp-connections | List My Mcp Connections
[**McpConnectionsRemoveMcpServerConnection**](McpConnectionsAPI.md#McpConnectionsRemoveMcpServerConnection) | **Delete** /api/v1/workspaces/{workspace_id}/mcp-servers/{server_id}/connections/{user_id} | Remove Mcp Server Connection
[**McpConnectionsStartMcpConnection**](McpConnectionsAPI.md#McpConnectionsStartMcpConnection) | **Post** /api/v1/workspaces/{workspace_id}/mcp-connections/{server_id} | Start Mcp Connection



## McpConnectionsCompleteMcpConnection

> McpConnectionCompleted McpConnectionsCompleteMcpConnection(ctx).McpConnectionComplete(mcpConnectionComplete).Execute()

Complete Mcp Connection



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
	mcpConnectionComplete := *openapiclient.NewMcpConnectionComplete("Code_example", "State_example") // McpConnectionComplete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpConnectionsAPI.McpConnectionsCompleteMcpConnection(context.Background()).McpConnectionComplete(mcpConnectionComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsCompleteMcpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpConnectionsCompleteMcpConnection`: McpConnectionCompleted
	fmt.Fprintf(os.Stdout, "Response from `McpConnectionsAPI.McpConnectionsCompleteMcpConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsCompleteMcpConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcpConnectionComplete** | [**McpConnectionComplete**](McpConnectionComplete.md) |  | 

### Return type

[**McpConnectionCompleted**](McpConnectionCompleted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpConnectionsEndMyMcpConnection

> McpConnectionsEndMyMcpConnection(ctx, workspaceId, serverId).Execute()

End My Mcp Connection



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
	r, err := apiClient.McpConnectionsAPI.McpConnectionsEndMyMcpConnection(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsEndMyMcpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsEndMyMcpConnectionRequest struct via the builder pattern


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


## McpConnectionsListMcpServerConnections

> McpServerConnectionsPublic McpConnectionsListMcpServerConnections(ctx, workspaceId, serverId).Skip(skip).Limit(limit).Execute()

List Mcp Server Connections



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpConnectionsAPI.McpConnectionsListMcpServerConnections(context.Background(), workspaceId, serverId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsListMcpServerConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpConnectionsListMcpServerConnections`: McpServerConnectionsPublic
	fmt.Fprintf(os.Stdout, "Response from `McpConnectionsAPI.McpConnectionsListMcpServerConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsListMcpServerConnectionsRequest struct via the builder pattern


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


## McpConnectionsListMyMcpConnections

> McpConnectionsPublic McpConnectionsListMyMcpConnections(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List My Mcp Connections



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
	resp, r, err := apiClient.McpConnectionsAPI.McpConnectionsListMyMcpConnections(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsListMyMcpConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpConnectionsListMyMcpConnections`: McpConnectionsPublic
	fmt.Fprintf(os.Stdout, "Response from `McpConnectionsAPI.McpConnectionsListMyMcpConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsListMyMcpConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**McpConnectionsPublic**](McpConnectionsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpConnectionsRemoveMcpServerConnection

> McpConnectionsRemoveMcpServerConnection(ctx, workspaceId, serverId, userId).Execute()

Remove Mcp Server Connection



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.McpConnectionsAPI.McpConnectionsRemoveMcpServerConnection(context.Background(), workspaceId, serverId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsRemoveMcpServerConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsRemoveMcpServerConnectionRequest struct via the builder pattern


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


## McpConnectionsStartMcpConnection

> McpConnectionStarted McpConnectionsStartMcpConnection(ctx, workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()

Start Mcp Connection



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
	resp, r, err := apiClient.McpConnectionsAPI.McpConnectionsStartMcpConnection(context.Background(), workspaceId, serverId).McpConnectionStart(mcpConnectionStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpConnectionsAPI.McpConnectionsStartMcpConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpConnectionsStartMcpConnection`: McpConnectionStarted
	fmt.Fprintf(os.Stdout, "Response from `McpConnectionsAPI.McpConnectionsStartMcpConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpConnectionsStartMcpConnectionRequest struct via the builder pattern


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

