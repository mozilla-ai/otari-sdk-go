# \WorkspacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspacesAddWorkspaceMember**](WorkspacesAPI.md#WorkspacesAddWorkspaceMember) | **Post** /api/v1/workspaces/{workspace_id}/members/{user_id} | Add Workspace Member
[**WorkspacesCreateWorkspace**](WorkspacesAPI.md#WorkspacesCreateWorkspace) | **Post** /api/v1/workspaces | Create Workspace
[**WorkspacesDeleteWorkspace**](WorkspacesAPI.md#WorkspacesDeleteWorkspace) | **Delete** /api/v1/workspaces/{workspace_id} | Delete Workspace
[**WorkspacesGetWorkspace**](WorkspacesAPI.md#WorkspacesGetWorkspace) | **Get** /api/v1/workspaces/{workspace_id} | Get Workspace
[**WorkspacesListWorkspaceMembers**](WorkspacesAPI.md#WorkspacesListWorkspaceMembers) | **Get** /api/v1/workspaces/{workspace_id}/members | List Workspace Members
[**WorkspacesListWorkspaces**](WorkspacesAPI.md#WorkspacesListWorkspaces) | **Get** /api/v1/workspaces | List Workspaces
[**WorkspacesRemoveWorkspaceMember**](WorkspacesAPI.md#WorkspacesRemoveWorkspaceMember) | **Delete** /api/v1/workspaces/{workspace_id}/members/{user_id} | Remove Workspace Member
[**WorkspacesUpdateWorkspace**](WorkspacesAPI.md#WorkspacesUpdateWorkspace) | **Patch** /api/v1/workspaces/{workspace_id} | Update Workspace
[**WorkspacesUpdateWorkspaceMemberRole**](WorkspacesAPI.md#WorkspacesUpdateWorkspaceMemberRole) | **Patch** /api/v1/workspaces/{workspace_id}/members/{user_id} | Update Workspace Member Role



## WorkspacesAddWorkspaceMember

> WorkspaceMemberPublic WorkspacesAddWorkspaceMember(ctx, workspaceId, userId).Role(role).Execute()

Add Workspace Member



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	role := "role_example" // string | Role to assign in this workspace. (optional) (default to "member")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesAddWorkspaceMember(context.Background(), workspaceId, userId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesAddWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesAddWorkspaceMember`: WorkspaceMemberPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesAddWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesAddWorkspaceMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **string** | Role to assign in this workspace. | [default to &quot;member&quot;]

### Return type

[**WorkspaceMemberPublic**](WorkspaceMemberPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesCreateWorkspace

> WorkspacePublic WorkspacesCreateWorkspace(ctx).WorkspaceCreate(workspaceCreate).Execute()

Create Workspace



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
	workspaceCreate := *openapiclient.NewWorkspaceCreate("Name_example") // WorkspaceCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesCreateWorkspace(context.Background()).WorkspaceCreate(workspaceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesCreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesCreateWorkspace`: WorkspacePublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesCreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceCreate** | [**WorkspaceCreate**](WorkspaceCreate.md) |  | 

### Return type

[**WorkspacePublic**](WorkspacePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesDeleteWorkspace

> Message WorkspacesDeleteWorkspace(ctx, workspaceId).Execute()

Delete Workspace



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
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesDeleteWorkspace(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesDeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesDeleteWorkspace`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesDeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesDeleteWorkspaceRequest struct via the builder pattern


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


## WorkspacesGetWorkspace

> WorkspacePublic WorkspacesGetWorkspace(ctx, workspaceId).Execute()

Get Workspace



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
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesGetWorkspace(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesGetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesGetWorkspace`: WorkspacePublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesGetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspacePublic**](WorkspacePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesListWorkspaceMembers

> WorkspaceMembersPublic WorkspacesListWorkspaceMembers(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Workspace Members



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
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesListWorkspaceMembers(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesListWorkspaceMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesListWorkspaceMembers`: WorkspaceMembersPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesListWorkspaceMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesListWorkspaceMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceMembersPublic**](WorkspaceMembersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesListWorkspaces

> WorkspacesPublic WorkspacesListWorkspaces(ctx).Skip(skip).Limit(limit).Execute()

List Workspaces



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesListWorkspaces(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesListWorkspaces`: WorkspacesPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspacesPublic**](WorkspacesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesRemoveWorkspaceMember

> Message WorkspacesRemoveWorkspaceMember(ctx, workspaceId, userId).Execute()

Remove Workspace Member



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesRemoveWorkspaceMember(context.Background(), workspaceId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesRemoveWorkspaceMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesRemoveWorkspaceMember`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesRemoveWorkspaceMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesRemoveWorkspaceMemberRequest struct via the builder pattern


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


## WorkspacesUpdateWorkspace

> WorkspacePublic WorkspacesUpdateWorkspace(ctx, workspaceId).WorkspaceUpdate(workspaceUpdate).Execute()

Update Workspace



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
	workspaceUpdate := *openapiclient.NewWorkspaceUpdate() // WorkspaceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesUpdateWorkspace(context.Background(), workspaceId).WorkspaceUpdate(workspaceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesUpdateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesUpdateWorkspace`: WorkspacePublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesUpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceUpdate** | [**WorkspaceUpdate**](WorkspaceUpdate.md) |  | 

### Return type

[**WorkspacePublic**](WorkspacePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesUpdateWorkspaceMemberRole

> WorkspaceMemberPublic WorkspacesUpdateWorkspaceMemberRole(ctx, workspaceId, userId).Role(role).Execute()

Update Workspace Member Role



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	role := "role_example" // string | Role to assign in this workspace.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.WorkspacesUpdateWorkspaceMemberRole(context.Background(), workspaceId, userId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.WorkspacesUpdateWorkspaceMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesUpdateWorkspaceMemberRole`: WorkspaceMemberPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.WorkspacesUpdateWorkspaceMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesUpdateWorkspaceMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **string** | Role to assign in this workspace. | 

### Return type

[**WorkspaceMemberPublic**](WorkspaceMemberPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

