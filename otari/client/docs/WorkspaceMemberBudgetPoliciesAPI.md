# \WorkspaceMemberBudgetPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault**](WorkspaceMemberBudgetPoliciesAPI.md#WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault) | **Post** /api/v1/workspaces/{workspace_id}/member-budget-policies | Create Workspace Budget Default
[**WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault**](WorkspaceMemberBudgetPoliciesAPI.md#WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault) | **Delete** /api/v1/workspaces/{workspace_id}/member-budget-policies/{default_id} | Delete Workspace Budget Default
[**WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults**](WorkspaceMemberBudgetPoliciesAPI.md#WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults) | **Get** /api/v1/workspaces/{workspace_id}/member-budget-policies | List Workspace Budget Defaults
[**WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault**](WorkspaceMemberBudgetPoliciesAPI.md#WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault) | **Patch** /api/v1/workspaces/{workspace_id}/member-budget-policies/{default_id} | Update Workspace Budget Default



## WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault

> WorkspaceMemberBudgetPolicyPublic WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault(ctx, workspaceId).WorkspaceMemberBudgetPolicyCreate(workspaceMemberBudgetPolicyCreate).Execute()

Create Workspace Budget Default



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
	workspaceMemberBudgetPolicyCreate := *openapiclient.NewWorkspaceMemberBudgetPolicyCreate("BudgetId_example") // WorkspaceMemberBudgetPolicyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault(context.Background(), workspaceId).WorkspaceMemberBudgetPolicyCreate(workspaceMemberBudgetPolicyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault`: WorkspaceMemberBudgetPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceMemberBudgetPoliciesCreateWorkspaceBudgetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceMemberBudgetPolicyCreate** | [**WorkspaceMemberBudgetPolicyCreate**](WorkspaceMemberBudgetPolicyCreate.md) |  | 

### Return type

[**WorkspaceMemberBudgetPolicyPublic**](WorkspaceMemberBudgetPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault

> Message WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault(ctx, workspaceId, defaultId).Execute()

Delete Workspace Budget Default



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
	defaultId := "defaultId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault(context.Background(), workspaceId, defaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**defaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceMemberBudgetPoliciesDeleteWorkspaceBudgetDefaultRequest struct via the builder pattern


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


## WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults

> WorkspaceMemberBudgetPoliciesPublic WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Workspace Budget Defaults



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
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults`: WorkspaceMemberBudgetPoliciesPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceMemberBudgetPoliciesListWorkspaceBudgetDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceMemberBudgetPoliciesPublic**](WorkspaceMemberBudgetPoliciesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault

> WorkspaceMemberBudgetPolicyPublic WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault(ctx, workspaceId, defaultId).WorkspaceMemberBudgetPolicyUpdate(workspaceMemberBudgetPolicyUpdate).Execute()

Update Workspace Budget Default



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
	defaultId := "defaultId_example" // string | 
	workspaceMemberBudgetPolicyUpdate := *openapiclient.NewWorkspaceMemberBudgetPolicyUpdate("BudgetId_example") // WorkspaceMemberBudgetPolicyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault(context.Background(), workspaceId, defaultId).WorkspaceMemberBudgetPolicyUpdate(workspaceMemberBudgetPolicyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault`: WorkspaceMemberBudgetPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.WorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**defaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceMemberBudgetPoliciesUpdateWorkspaceBudgetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceMemberBudgetPolicyUpdate** | [**WorkspaceMemberBudgetPolicyUpdate**](WorkspaceMemberBudgetPolicyUpdate.md) |  | 

### Return type

[**WorkspaceMemberBudgetPolicyPublic**](WorkspaceMemberBudgetPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

