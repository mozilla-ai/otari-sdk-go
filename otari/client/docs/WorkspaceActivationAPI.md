# \WorkspaceActivationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceActivationCreateWorkspaceActivationKey**](WorkspaceActivationAPI.md#WorkspaceActivationCreateWorkspaceActivationKey) | **Post** /api/v1/workspaces/{workspace_id}/activation/key | Create Workspace Activation Key
[**WorkspaceActivationDismissWorkspaceActivation**](WorkspaceActivationAPI.md#WorkspaceActivationDismissWorkspaceActivation) | **Post** /api/v1/workspaces/{workspace_id}/activation/dismiss | Dismiss Workspace Activation
[**WorkspaceActivationGetWorkspaceActivation**](WorkspaceActivationAPI.md#WorkspaceActivationGetWorkspaceActivation) | **Get** /api/v1/workspaces/{workspace_id}/activation | Get Workspace Activation



## WorkspaceActivationCreateWorkspaceActivationKey

> ActivationApiKeyPublic WorkspaceActivationCreateWorkspaceActivationKey(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceActivationAPI.WorkspaceActivationCreateWorkspaceActivationKey(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.WorkspaceActivationCreateWorkspaceActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceActivationCreateWorkspaceActivationKey`: ActivationApiKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.WorkspaceActivationCreateWorkspaceActivationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceActivationCreateWorkspaceActivationKeyRequest struct via the builder pattern


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


## WorkspaceActivationDismissWorkspaceActivation

> Message WorkspaceActivationDismissWorkspaceActivation(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceActivationAPI.WorkspaceActivationDismissWorkspaceActivation(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.WorkspaceActivationDismissWorkspaceActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceActivationDismissWorkspaceActivation`: Message
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.WorkspaceActivationDismissWorkspaceActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceActivationDismissWorkspaceActivationRequest struct via the builder pattern


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


## WorkspaceActivationGetWorkspaceActivation

> WorkspaceActivationPublic WorkspaceActivationGetWorkspaceActivation(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.WorkspaceActivationAPI.WorkspaceActivationGetWorkspaceActivation(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceActivationAPI.WorkspaceActivationGetWorkspaceActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceActivationGetWorkspaceActivation`: WorkspaceActivationPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceActivationAPI.WorkspaceActivationGetWorkspaceActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceActivationGetWorkspaceActivationRequest struct via the builder pattern


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

