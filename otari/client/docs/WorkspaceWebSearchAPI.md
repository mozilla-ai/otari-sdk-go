# \WorkspaceWebSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete**](WorkspaceWebSearchAPI.md#ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete) | **Delete** /v1/workspaces/{workspace_id}/web-search | Clear Workspace Web Search Config
[**GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet**](WorkspaceWebSearchAPI.md#GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet) | **Get** /v1/workspaces/{workspace_id}/web-search | Get Workspace Web Search Config
[**SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut**](WorkspaceWebSearchAPI.md#SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut) | **Put** /v1/workspaces/{workspace_id}/web-search | Set Workspace Web Search Config



## ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete

> WorkspaceWebSearchConfigPublic ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete(ctx, workspaceId).Execute()

Clear Workspace Web Search Config



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceWebSearchAPI.ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.ClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceWebSearchConfigPublic**](WorkspaceWebSearchConfigPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet

> WorkspaceWebSearchConfigPublic GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet(ctx, workspaceId).Execute()

Get Workspace Web Search Config



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceWebSearchAPI.GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.GetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceWebSearchConfigPublic**](WorkspaceWebSearchConfigPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut

> WorkspaceWebSearchConfigPublic SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut(ctx, workspaceId).WorkspaceWebSearchConfigUpdate(workspaceWebSearchConfigUpdate).Execute()

Set Workspace Web Search Config



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
	workspaceWebSearchConfigUpdate := *openapiclient.NewWorkspaceWebSearchConfigUpdate(false) // WorkspaceWebSearchConfigUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceWebSearchAPI.SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut(context.Background(), workspaceId).WorkspaceWebSearchConfigUpdate(workspaceWebSearchConfigUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.SetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorkspaceWebSearchConfigV1WorkspacesWorkspaceIdWebSearchPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceWebSearchConfigUpdate** | [**WorkspaceWebSearchConfigUpdate**](WorkspaceWebSearchConfigUpdate.md) |  | 

### Return type

[**WorkspaceWebSearchConfigPublic**](WorkspaceWebSearchConfigPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

