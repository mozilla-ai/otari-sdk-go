# \WorkspaceMemberBudgetPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost**](WorkspaceMemberBudgetPoliciesAPI.md#CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost) | **Post** /v1/workspaces/{workspace_id}/member-budget-policies | Create Workspace Budget Default
[**DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete**](WorkspaceMemberBudgetPoliciesAPI.md#DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete) | **Delete** /v1/workspaces/{workspace_id}/member-budget-policies/{default_id} | Delete Workspace Budget Default
[**ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet**](WorkspaceMemberBudgetPoliciesAPI.md#ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet) | **Get** /v1/workspaces/{workspace_id}/member-budget-policies | List Workspace Budget Defaults
[**UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch**](WorkspaceMemberBudgetPoliciesAPI.md#UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch) | **Patch** /v1/workspaces/{workspace_id}/member-budget-policies/{default_id} | Update Workspace Budget Default



## CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost

> WorkspaceMemberBudgetPolicyPublic CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost(ctx, workspaceId).WorkspaceMemberBudgetPolicyCreate(workspaceMemberBudgetPolicyCreate).Execute()

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
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost(context.Background(), workspaceId).WorkspaceMemberBudgetPolicyCreate(workspaceMemberBudgetPolicyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost`: WorkspaceMemberBudgetPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.CreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesPostRequest struct via the builder pattern


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


## DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete

> Message DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete(ctx, workspaceId, defaultId).Execute()

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
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete(context.Background(), workspaceId, defaultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.DeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**defaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdDeleteRequest struct via the builder pattern


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


## ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet

> WorkspaceMemberBudgetPoliciesPublic ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet`: WorkspaceMemberBudgetPoliciesPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.ListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceBudgetDefaultsV1WorkspacesWorkspaceIdMemberBudgetPoliciesGetRequest struct via the builder pattern


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


## UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch

> WorkspaceMemberBudgetPolicyPublic UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch(ctx, workspaceId, defaultId).WorkspaceMemberBudgetPolicyUpdate(workspaceMemberBudgetPolicyUpdate).Execute()

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
	resp, r, err := apiClient.WorkspaceMemberBudgetPoliciesAPI.UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch(context.Background(), workspaceId, defaultId).WorkspaceMemberBudgetPolicyUpdate(workspaceMemberBudgetPolicyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMemberBudgetPoliciesAPI.UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch`: WorkspaceMemberBudgetPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMemberBudgetPoliciesAPI.UpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**defaultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceBudgetDefaultV1WorkspacesWorkspaceIdMemberBudgetPoliciesDefaultIdPatchRequest struct via the builder pattern


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

