# \WorkspaceActivationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost**](WorkspaceActivationAPI.md#CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost) | **Post** /v1/workspaces/{workspace_id}/activation/key | Create Workspace Activation Key
[**DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost**](WorkspaceActivationAPI.md#DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost) | **Post** /v1/workspaces/{workspace_id}/activation/dismiss | Dismiss Workspace Activation
[**GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet**](WorkspaceActivationAPI.md#GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet) | **Get** /v1/workspaces/{workspace_id}/activation | Get Workspace Activation



## CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost

> ActivationApiKeyPublic CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost(ctx, workspaceId).Execute()

Create Workspace Activation Key



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
	resp, r, err := apiClient.WorkspaceActivationAPI.CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost`: ActivationApiKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.CreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceActivationKeyV1WorkspacesWorkspaceIdActivationKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivationApiKeyPublic**](ActivationApiKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost

> Message DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost(ctx, workspaceId).Execute()

Dismiss Workspace Activation



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
	resp, r, err := apiClient.WorkspaceActivationAPI.DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.DismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDismissWorkspaceActivationV1WorkspacesWorkspaceIdActivationDismissPostRequest struct via the builder pattern


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


## GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet

> WorkspaceActivationPublic GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet(ctx, workspaceId).Execute()

Get Workspace Activation



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
	resp, r, err := apiClient.WorkspaceActivationAPI.GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet`: WorkspaceActivationPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.GetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceActivationV1WorkspacesWorkspaceIdActivationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceActivationPublic**](WorkspaceActivationPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

