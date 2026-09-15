# \ProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvidersCreateStoredProvider**](ProvidersAPI.md#ProvidersCreateStoredProvider) | **Post** /api/v1/provider-credentials | Create Stored Provider
[**ProvidersDeleteStoredProvider**](ProvidersAPI.md#ProvidersDeleteStoredProvider) | **Delete** /api/v1/provider-credentials/{instance} | Delete Stored Provider
[**ProvidersListProviders**](ProvidersAPI.md#ProvidersListProviders) | **Get** /api/v1/providers | List Providers
[**ProvidersListStoredProviders**](ProvidersAPI.md#ProvidersListStoredProviders) | **Get** /api/v1/provider-credentials | List Stored Providers
[**ProvidersProviderCatalog**](ProvidersAPI.md#ProvidersProviderCatalog) | **Get** /api/v1/providers/catalog | Provider Catalog
[**ProvidersProviderCatalogDetail**](ProvidersAPI.md#ProvidersProviderCatalogDetail) | **Get** /api/v1/providers/catalog/{provider_id} | Provider Catalog Detail
[**ProvidersProviderHealth**](ProvidersAPI.md#ProvidersProviderHealth) | **Get** /api/v1/providers/health | Provider Health
[**ProvidersReencryptStoredProviderKeys**](ProvidersAPI.md#ProvidersReencryptStoredProviderKeys) | **Post** /api/v1/provider-credentials/reencrypt | Reencrypt Stored Provider Keys
[**ProvidersTestProviderConnection**](ProvidersAPI.md#ProvidersTestProviderConnection) | **Post** /api/v1/provider-credentials/test | Test Provider Connection
[**ProvidersTestStoredProvider**](ProvidersAPI.md#ProvidersTestStoredProvider) | **Post** /api/v1/provider-credentials/{instance}/test | Test Stored Provider
[**ProvidersUpdateStoredProvider**](ProvidersAPI.md#ProvidersUpdateStoredProvider) | **Patch** /api/v1/provider-credentials/{instance} | Update Stored Provider



## ProvidersCreateStoredProvider

> StoredProviderResponse ProvidersCreateStoredProvider(ctx).CreateStoredProviderRequest(createStoredProviderRequest).Execute()

Create Stored Provider



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
	createStoredProviderRequest := *openapiclient.NewCreateStoredProviderRequest("Instance_example") // CreateStoredProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersCreateStoredProvider(context.Background()).CreateStoredProviderRequest(createStoredProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersCreateStoredProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersCreateStoredProvider`: StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersCreateStoredProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersCreateStoredProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStoredProviderRequest** | [**CreateStoredProviderRequest**](CreateStoredProviderRequest.md) |  | 

### Return type

[**StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersDeleteStoredProvider

> ProvidersDeleteStoredProvider(ctx, instance).Execute()

Delete Stored Provider



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
	instance := "instance_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvidersAPI.ProvidersDeleteStoredProvider(context.Background(), instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersDeleteStoredProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersDeleteStoredProviderRequest struct via the builder pattern


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


## ProvidersListProviders

> ProvidersResponse ProvidersListProviders(ctx).Execute()

List Providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersListProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersListProviders`: ProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersListProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListProvidersRequest struct via the builder pattern


### Return type

[**ProvidersResponse**](ProvidersResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersListStoredProviders

> []StoredProviderResponse ProvidersListStoredProviders(ctx).Execute()

List Stored Providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersListStoredProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersListStoredProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersListStoredProviders`: []StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersListStoredProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListStoredProvidersRequest struct via the builder pattern


### Return type

[**[]StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProviderCatalog

> []KnownProviderSummarySchema ProvidersProviderCatalog(ctx).Execute()

Provider Catalog



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersProviderCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersProviderCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersProviderCatalog`: []KnownProviderSummarySchema
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersProviderCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProviderCatalogRequest struct via the builder pattern


### Return type

[**[]KnownProviderSummarySchema**](KnownProviderSummarySchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProviderCatalogDetail

> KnownProviderSchema ProvidersProviderCatalogDetail(ctx, providerId).Execute()

Provider Catalog Detail



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
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersProviderCatalogDetail(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersProviderCatalogDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersProviderCatalogDetail`: KnownProviderSchema
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersProviderCatalogDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProviderCatalogDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnownProviderSchema**](KnownProviderSchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProviderHealth

> ProviderHealthResponse ProvidersProviderHealth(ctx).Refresh(refresh).Execute()

Provider Health



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
	refresh := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersProviderHealth(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersProviderHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersProviderHealth`: ProviderHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersProviderHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProviderHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** |  | [default to false]

### Return type

[**ProviderHealthResponse**](ProviderHealthResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersReencryptStoredProviderKeys

> ReencryptProviderCredentialsResponse ProvidersReencryptStoredProviderKeys(ctx).Execute()

Reencrypt Stored Provider Keys



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersReencryptStoredProviderKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersReencryptStoredProviderKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersReencryptStoredProviderKeys`: ReencryptProviderCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersReencryptStoredProviderKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersReencryptStoredProviderKeysRequest struct via the builder pattern


### Return type

[**ReencryptProviderCredentialsResponse**](ReencryptProviderCredentialsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersTestProviderConnection

> TestProviderResponse ProvidersTestProviderConnection(ctx).TestProviderRequest(testProviderRequest).Execute()

Test Provider Connection



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
	testProviderRequest := *openapiclient.NewTestProviderRequest() // TestProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersTestProviderConnection(context.Background()).TestProviderRequest(testProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersTestProviderConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersTestProviderConnection`: TestProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersTestProviderConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersTestProviderConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testProviderRequest** | [**TestProviderRequest**](TestProviderRequest.md) |  | 

### Return type

[**TestProviderResponse**](TestProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersTestStoredProvider

> TestProviderResponse ProvidersTestStoredProvider(ctx, instance).Execute()

Test Stored Provider



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
	instance := "instance_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersTestStoredProvider(context.Background(), instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersTestStoredProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersTestStoredProvider`: TestProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersTestStoredProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersTestStoredProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestProviderResponse**](TestProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersUpdateStoredProvider

> StoredProviderResponse ProvidersUpdateStoredProvider(ctx, instance).UpdateStoredProviderRequest(updateStoredProviderRequest).Execute()

Update Stored Provider



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
	instance := "instance_example" // string | 
	updateStoredProviderRequest := *openapiclient.NewUpdateStoredProviderRequest() // UpdateStoredProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersUpdateStoredProvider(context.Background(), instance).UpdateStoredProviderRequest(updateStoredProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersUpdateStoredProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersUpdateStoredProvider`: StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersUpdateStoredProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersUpdateStoredProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoredProviderRequest** | [**UpdateStoredProviderRequest**](UpdateStoredProviderRequest.md) |  | 

### Return type

[**StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

