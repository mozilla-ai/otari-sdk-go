# \WorkspaceCodeExecutionPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy**](WorkspaceCodeExecutionPolicyAPI.md#WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy) | **Delete** /api/v1/workspaces/{workspace_id}/code-execution-policy | Clear Workspace Code Execution Policy
[**WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy**](WorkspaceCodeExecutionPolicyAPI.md#WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy) | **Get** /api/v1/workspaces/{workspace_id}/code-execution-policy | Get Workspace Code Execution Policy
[**WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy**](WorkspaceCodeExecutionPolicyAPI.md#WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy) | **Put** /api/v1/workspaces/{workspace_id}/code-execution-policy | Set Workspace Code Execution Policy



## WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy

> WorkspaceCodeExecutionPolicyPublic WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceCodeExecutionPolicyClearWorkspaceCodeExecutionPolicyRequest struct via the builder pattern


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


## WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy

> WorkspaceCodeExecutionPolicyPublic WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceCodeExecutionPolicyGetWorkspaceCodeExecutionPolicyRequest struct via the builder pattern


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


## WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy

> WorkspaceCodeExecutionPolicyPublic WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy(ctx, workspaceId).WorkspaceCodeExecutionPolicyUpdate(workspaceCodeExecutionPolicyUpdate).Execute()

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
	resp, r, err := apiClient.WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy(context.Background(), workspaceId).WorkspaceCodeExecutionPolicyUpdate(workspaceCodeExecutionPolicyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy`: WorkspaceCodeExecutionPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceCodeExecutionPolicyAPI.WorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceCodeExecutionPolicySetWorkspaceCodeExecutionPolicyRequest struct via the builder pattern


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

