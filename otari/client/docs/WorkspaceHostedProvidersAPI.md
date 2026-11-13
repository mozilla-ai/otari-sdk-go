# \WorkspaceHostedProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceHostedProvidersGetWorkspaceProviderPolicy**](WorkspaceHostedProvidersAPI.md#WorkspaceHostedProvidersGetWorkspaceProviderPolicy) | **Get** /api/v1/workspaces/{workspace_id}/provider-policy | Get Workspace Provider Policy
[**WorkspaceHostedProvidersListWorkspaceHostedProviders**](WorkspaceHostedProvidersAPI.md#WorkspaceHostedProvidersListWorkspaceHostedProviders) | **Get** /api/v1/workspaces/{workspace_id}/hosted-providers | List Workspace Hosted Providers
[**WorkspaceHostedProvidersResetWorkspaceHostedProvider**](WorkspaceHostedProvidersAPI.md#WorkspaceHostedProvidersResetWorkspaceHostedProvider) | **Delete** /api/v1/workspaces/{workspace_id}/hosted-providers/{provider} | Reset Workspace Hosted Provider
[**WorkspaceHostedProvidersSetWorkspaceHostedProvider**](WorkspaceHostedProvidersAPI.md#WorkspaceHostedProvidersSetWorkspaceHostedProvider) | **Put** /api/v1/workspaces/{workspace_id}/hosted-providers/{provider} | Set Workspace Hosted Provider
[**WorkspaceHostedProvidersSetWorkspaceProviderPolicy**](WorkspaceHostedProvidersAPI.md#WorkspaceHostedProvidersSetWorkspaceProviderPolicy) | **Put** /api/v1/workspaces/{workspace_id}/provider-policy | Set Workspace Provider Policy



## WorkspaceHostedProvidersGetWorkspaceProviderPolicy

> WorkspaceProviderPolicyPublic WorkspaceHostedProvidersGetWorkspaceProviderPolicy(ctx, workspaceId).Execute()

Get Workspace Provider Policy



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
	resp, r, err := apiClient.WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersGetWorkspaceProviderPolicy(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersGetWorkspaceProviderPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceHostedProvidersGetWorkspaceProviderPolicy`: WorkspaceProviderPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersGetWorkspaceProviderPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceHostedProvidersGetWorkspaceProviderPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceProviderPolicyPublic**](WorkspaceProviderPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceHostedProvidersListWorkspaceHostedProviders

> WorkspaceHostedProvidersPublic WorkspaceHostedProvidersListWorkspaceHostedProviders(ctx, workspaceId).Execute()

List Workspace Hosted Providers



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
	resp, r, err := apiClient.WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersListWorkspaceHostedProviders(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersListWorkspaceHostedProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceHostedProvidersListWorkspaceHostedProviders`: WorkspaceHostedProvidersPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersListWorkspaceHostedProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceHostedProvidersListWorkspaceHostedProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceHostedProvidersPublic**](WorkspaceHostedProvidersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceHostedProvidersResetWorkspaceHostedProvider

> WorkspaceHostedProviderPublic WorkspaceHostedProvidersResetWorkspaceHostedProvider(ctx, workspaceId, provider).Execute()

Reset Workspace Hosted Provider



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersResetWorkspaceHostedProvider(context.Background(), workspaceId, provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersResetWorkspaceHostedProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceHostedProvidersResetWorkspaceHostedProvider`: WorkspaceHostedProviderPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersResetWorkspaceHostedProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceHostedProvidersResetWorkspaceHostedProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceHostedProviderPublic**](WorkspaceHostedProviderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceHostedProvidersSetWorkspaceHostedProvider

> WorkspaceHostedProviderPublic WorkspaceHostedProvidersSetWorkspaceHostedProvider(ctx, workspaceId, provider).WorkspaceHostedProviderAccessRequest(workspaceHostedProviderAccessRequest).Execute()

Set Workspace Hosted Provider



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
	provider := "provider_example" // string | 
	workspaceHostedProviderAccessRequest := *openapiclient.NewWorkspaceHostedProviderAccessRequest(false) // WorkspaceHostedProviderAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceHostedProvider(context.Background(), workspaceId, provider).WorkspaceHostedProviderAccessRequest(workspaceHostedProviderAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceHostedProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceHostedProvidersSetWorkspaceHostedProvider`: WorkspaceHostedProviderPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceHostedProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceHostedProvidersSetWorkspaceHostedProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceHostedProviderAccessRequest** | [**WorkspaceHostedProviderAccessRequest**](WorkspaceHostedProviderAccessRequest.md) |  | 

### Return type

[**WorkspaceHostedProviderPublic**](WorkspaceHostedProviderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceHostedProvidersSetWorkspaceProviderPolicy

> WorkspaceProviderPolicyPublic WorkspaceHostedProvidersSetWorkspaceProviderPolicy(ctx, workspaceId).WorkspaceProviderPolicyRequest(workspaceProviderPolicyRequest).Execute()

Set Workspace Provider Policy



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
	workspaceProviderPolicyRequest := *openapiclient.NewWorkspaceProviderPolicyRequest(false) // WorkspaceProviderPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceProviderPolicy(context.Background(), workspaceId).WorkspaceProviderPolicyRequest(workspaceProviderPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceProviderPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceHostedProvidersSetWorkspaceProviderPolicy`: WorkspaceProviderPolicyPublic
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceHostedProvidersAPI.WorkspaceHostedProvidersSetWorkspaceProviderPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceHostedProvidersSetWorkspaceProviderPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceProviderPolicyRequest** | [**WorkspaceProviderPolicyRequest**](WorkspaceProviderPolicyRequest.md) |  | 

### Return type

[**WorkspaceProviderPolicyPublic**](WorkspaceProviderPolicyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

