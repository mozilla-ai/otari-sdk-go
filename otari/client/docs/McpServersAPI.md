# \McpServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost**](McpServersAPI.md#CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost) | **Post** /v1/workspaces/{workspace_id}/mcp-servers | Create Workspace Mcp Server
[**DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete**](McpServersAPI.md#DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete) | **Delete** /v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Delete Workspace Mcp Server
[**ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet**](McpServersAPI.md#ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet) | **Get** /v1/workspaces/{workspace_id}/mcp-servers | List Workspace Mcp Servers
[**UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch**](McpServersAPI.md#UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch) | **Patch** /v1/workspaces/{workspace_id}/mcp-servers/{server_id} | Update Workspace Mcp Server



## CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost

> WorkspaceMcpServerPublic CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost(ctx, workspaceId).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()

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
	resp, r, err := apiClient.McpServersAPI.CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost(context.Background(), workspaceId).WorkspaceMcpServerCreate(workspaceMcpServerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.CreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersPostRequest struct via the builder pattern


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


## DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete

> Message DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete(ctx, workspaceId, serverId).Execute()

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
	resp, r, err := apiClient.McpServersAPI.DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete(context.Background(), workspaceId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.DeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdDeleteRequest struct via the builder pattern


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


## ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet

> WorkspaceMcpServersPublic ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.McpServersAPI.ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet`: WorkspaceMcpServersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.ListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceMcpServersV1WorkspacesWorkspaceIdMcpServersGetRequest struct via the builder pattern


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


## UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch

> WorkspaceMcpServerPublic UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch(ctx, workspaceId, serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()

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
	resp, r, err := apiClient.McpServersAPI.UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch(context.Background(), workspaceId, serverId).WorkspaceMcpServerUpdate(workspaceMcpServerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpServersAPI.UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch`: WorkspaceMcpServerPublic
	fmt.Fprintf(os.Stdout, "Response from `McpServersAPI.UpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**serverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceMcpServerV1WorkspacesWorkspaceIdMcpServersServerIdPatchRequest struct via the builder pattern


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

