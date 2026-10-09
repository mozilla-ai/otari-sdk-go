# \ProviderEndpointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProviderEndpointsCreateProviderEndpoint**](ProviderEndpointsAPI.md#ProviderEndpointsCreateProviderEndpoint) | **Post** /api/v1/provider-endpoints | Create Provider Endpoint
[**ProviderEndpointsDeleteProviderEndpoint**](ProviderEndpointsAPI.md#ProviderEndpointsDeleteProviderEndpoint) | **Delete** /api/v1/provider-endpoints/{endpoint_id} | Delete Provider Endpoint
[**ProviderEndpointsGetProviderEndpoint**](ProviderEndpointsAPI.md#ProviderEndpointsGetProviderEndpoint) | **Get** /api/v1/provider-endpoints/{endpoint_id} | Get Provider Endpoint
[**ProviderEndpointsListProviderEndpoints**](ProviderEndpointsAPI.md#ProviderEndpointsListProviderEndpoints) | **Get** /api/v1/provider-endpoints | List Provider Endpoints
[**ProviderEndpointsUpdateProviderEndpoint**](ProviderEndpointsAPI.md#ProviderEndpointsUpdateProviderEndpoint) | **Patch** /api/v1/provider-endpoints/{endpoint_id} | Update Provider Endpoint



## ProviderEndpointsCreateProviderEndpoint

> ProviderEndpointPublic ProviderEndpointsCreateProviderEndpoint(ctx).ProviderEndpointCreateRequest(providerEndpointCreateRequest).Execute()

Create Provider Endpoint



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
	providerEndpointCreateRequest := *openapiclient.NewProviderEndpointCreateRequest("ApiBase_example", "Name_example", "Provider_example") // ProviderEndpointCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderEndpointsAPI.ProviderEndpointsCreateProviderEndpoint(context.Background()).ProviderEndpointCreateRequest(providerEndpointCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderEndpointsAPI.ProviderEndpointsCreateProviderEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderEndpointsCreateProviderEndpoint`: ProviderEndpointPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderEndpointsAPI.ProviderEndpointsCreateProviderEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderEndpointsCreateProviderEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerEndpointCreateRequest** | [**ProviderEndpointCreateRequest**](ProviderEndpointCreateRequest.md) |  | 

### Return type

[**ProviderEndpointPublic**](ProviderEndpointPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderEndpointsDeleteProviderEndpoint

> ProviderEndpointsDeleteProviderEndpoint(ctx, endpointId).Execute()

Delete Provider Endpoint



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
	endpointId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProviderEndpointsAPI.ProviderEndpointsDeleteProviderEndpoint(context.Background(), endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderEndpointsAPI.ProviderEndpointsDeleteProviderEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderEndpointsDeleteProviderEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderEndpointsGetProviderEndpoint

> ProviderEndpointPublic ProviderEndpointsGetProviderEndpoint(ctx, endpointId).Execute()

Get Provider Endpoint



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
	endpointId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderEndpointsAPI.ProviderEndpointsGetProviderEndpoint(context.Background(), endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderEndpointsAPI.ProviderEndpointsGetProviderEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderEndpointsGetProviderEndpoint`: ProviderEndpointPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderEndpointsAPI.ProviderEndpointsGetProviderEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderEndpointsGetProviderEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderEndpointPublic**](ProviderEndpointPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderEndpointsListProviderEndpoints

> ProviderEndpointsPublic ProviderEndpointsListProviderEndpoints(ctx).WorkspaceId(workspaceId).UserId(userId).Skip(skip).Limit(limit).Execute()

List Provider Endpoints



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only endpoints this workspace owns. (optional)
	userId := "userId_example" // string | Only endpoints this user owns. (optional)
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderEndpointsAPI.ProviderEndpointsListProviderEndpoints(context.Background()).WorkspaceId(workspaceId).UserId(userId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderEndpointsAPI.ProviderEndpointsListProviderEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderEndpointsListProviderEndpoints`: ProviderEndpointsPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderEndpointsAPI.ProviderEndpointsListProviderEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderEndpointsListProviderEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Only endpoints this workspace owns. | 
 **userId** | **string** | Only endpoints this user owns. | 
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**ProviderEndpointsPublic**](ProviderEndpointsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderEndpointsUpdateProviderEndpoint

> ProviderEndpointPublic ProviderEndpointsUpdateProviderEndpoint(ctx, endpointId).ProviderEndpointUpdateRequest(providerEndpointUpdateRequest).Execute()

Update Provider Endpoint



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
	endpointId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	providerEndpointUpdateRequest := *openapiclient.NewProviderEndpointUpdateRequest() // ProviderEndpointUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderEndpointsAPI.ProviderEndpointsUpdateProviderEndpoint(context.Background(), endpointId).ProviderEndpointUpdateRequest(providerEndpointUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderEndpointsAPI.ProviderEndpointsUpdateProviderEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderEndpointsUpdateProviderEndpoint`: ProviderEndpointPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderEndpointsAPI.ProviderEndpointsUpdateProviderEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderEndpointsUpdateProviderEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerEndpointUpdateRequest** | [**ProviderEndpointUpdateRequest**](ProviderEndpointUpdateRequest.md) |  | 

### Return type

[**ProviderEndpointPublic**](ProviderEndpointPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

