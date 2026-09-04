# \WorkspaceCodeExecutionPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete**](WorkspaceCodeExecutionPolicyAPI.md#ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete) | **Delete** /v1/workspaces/{workspace_id}/code-execution-policy | Clear Workspace Code Execution Policy
[**GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet**](WorkspaceCodeExecutionPolicyAPI.md#GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet) | **Get** /v1/workspaces/{workspace_id}/code-execution-policy | Get Workspace Code Execution Policy
[**SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut**](WorkspaceCodeExecutionPolicyAPI.md#SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut) | **Put** /v1/workspaces/{workspace_id}/code-execution-policy | Set Workspace Code Execution Policy



## ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete

> WorkspaceCodeExecutionPolicyPublic ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete(ctx, workspaceId).Execute()

Clear Workspace Code Execution Policy



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
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.ClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceCodeExecutionPolicyPublic**](WorkspaceCodeExecutionPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet

> WorkspaceCodeExecutionPolicyPublic GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet(ctx, workspaceId).Execute()

Get Workspace Code Execution Policy



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
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.GetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceCodeExecutionPolicyPublic**](WorkspaceCodeExecutionPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut

> WorkspaceCodeExecutionPolicyPublic SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut(ctx, workspaceId).WorkspaceCodeExecutionPolicyUpdate(workspaceCodeExecutionPolicyUpdate).Execute()

Set Workspace Code Execution Policy



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
	workspaceCodeExecutionPolicyUpdate := *openapiclient.NewWorkspaceCodeExecutionPolicyUpdate(false) // WorkspaceCodeExecutionPolicyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut(context.Background(), workspaceId).WorkspaceCodeExecutionPolicyUpdate(workspaceCodeExecutionPolicyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.SetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorkspaceCodeExecutionPolicyV1WorkspacesWorkspaceIdCodeExecutionPolicyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceCodeExecutionPolicyUpdate** | [**WorkspaceCodeExecutionPolicyUpdate**](WorkspaceCodeExecutionPolicyUpdate.md) |  | 

### Return type

[**WorkspaceCodeExecutionPolicyPublic**](WorkspaceCodeExecutionPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

