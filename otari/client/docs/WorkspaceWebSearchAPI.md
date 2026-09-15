# \WorkspaceWebSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceWebSearchClearWorkspaceWebSearchConfig**](WorkspaceWebSearchAPI.md#WorkspaceWebSearchClearWorkspaceWebSearchConfig) | **Delete** /api/v1/workspaces/{workspace_id}/web-search | Clear Workspace Web Search Config
[**WorkspaceWebSearchGetWorkspaceWebSearchConfig**](WorkspaceWebSearchAPI.md#WorkspaceWebSearchGetWorkspaceWebSearchConfig) | **Get** /api/v1/workspaces/{workspace_id}/web-search | Get Workspace Web Search Config
[**WorkspaceWebSearchSetWorkspaceWebSearchConfig**](WorkspaceWebSearchAPI.md#WorkspaceWebSearchSetWorkspaceWebSearchConfig) | **Put** /api/v1/workspaces/{workspace_id}/web-search | Set Workspace Web Search Config



## WorkspaceWebSearchClearWorkspaceWebSearchConfig

> WorkspaceWebSearchConfigPublic WorkspaceWebSearchClearWorkspaceWebSearchConfig(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceWebSearchAPI.WorkspaceWebSearchClearWorkspaceWebSearchConfig(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.WorkspaceWebSearchClearWorkspaceWebSearchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceWebSearchClearWorkspaceWebSearchConfig`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.WorkspaceWebSearchClearWorkspaceWebSearchConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceWebSearchClearWorkspaceWebSearchConfigRequest struct via the builder pattern


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


## WorkspaceWebSearchGetWorkspaceWebSearchConfig

> WorkspaceWebSearchConfigPublic WorkspaceWebSearchGetWorkspaceWebSearchConfig(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceWebSearchAPI.WorkspaceWebSearchGetWorkspaceWebSearchConfig(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.WorkspaceWebSearchGetWorkspaceWebSearchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceWebSearchGetWorkspaceWebSearchConfig`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.WorkspaceWebSearchGetWorkspaceWebSearchConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceWebSearchGetWorkspaceWebSearchConfigRequest struct via the builder pattern


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


## WorkspaceWebSearchSetWorkspaceWebSearchConfig

> WorkspaceWebSearchConfigPublic WorkspaceWebSearchSetWorkspaceWebSearchConfig(ctx, workspaceId).WorkspaceWebSearchConfigUpdate(workspaceWebSearchConfigUpdate).Execute()

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
	resp, r, err := apiClient.WorkspaceWebSearchAPI.WorkspaceWebSearchSetWorkspaceWebSearchConfig(context.Background(), workspaceId).WorkspaceWebSearchConfigUpdate(workspaceWebSearchConfigUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceWebSearchAPI.WorkspaceWebSearchSetWorkspaceWebSearchConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceWebSearchSetWorkspaceWebSearchConfig`: WorkspaceWebSearchConfigPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceWebSearchAPI.WorkspaceWebSearchSetWorkspaceWebSearchConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceWebSearchSetWorkspaceWebSearchConfigRequest struct via the builder pattern


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

