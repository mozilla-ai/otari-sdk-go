# \HostedProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostedProvidersAddHostedModel**](HostedProvidersAPI.md#HostedProvidersAddHostedModel) | **Post** /api/v1/hosted-providers/{provider}/models | Add Hosted Model
[**HostedProvidersCreateHostedProvider**](HostedProvidersAPI.md#HostedProvidersCreateHostedProvider) | **Post** /api/v1/hosted-providers | Create Hosted Provider
[**HostedProvidersDeleteHostedProvider**](HostedProvidersAPI.md#HostedProvidersDeleteHostedProvider) | **Delete** /api/v1/hosted-providers/{provider} | Delete Hosted Provider
[**HostedProvidersListAvailableModels**](HostedProvidersAPI.md#HostedProvidersListAvailableModels) | **Get** /api/v1/hosted-providers/{provider}/available-models | List Available Models
[**HostedProvidersListHostedModels**](HostedProvidersAPI.md#HostedProvidersListHostedModels) | **Get** /api/v1/hosted-providers/{provider}/models | List Hosted Models
[**HostedProvidersListHostedProviders**](HostedProvidersAPI.md#HostedProvidersListHostedProviders) | **Get** /api/v1/hosted-providers | List Hosted Providers
[**HostedProvidersPreviewCatalogRefresh**](HostedProvidersAPI.md#HostedProvidersPreviewCatalogRefresh) | **Get** /api/v1/hosted-providers/refresh/preview | Preview Catalog Refresh
[**HostedProvidersRefreshCatalog**](HostedProvidersAPI.md#HostedProvidersRefreshCatalog) | **Post** /api/v1/hosted-providers/refresh | Refresh Catalog
[**HostedProvidersRefreshHostedModels**](HostedProvidersAPI.md#HostedProvidersRefreshHostedModels) | **Post** /api/v1/hosted-providers/{provider}/models/refresh | Refresh Hosted Models
[**HostedProvidersRemoveHostedModel**](HostedProvidersAPI.md#HostedProvidersRemoveHostedModel) | **Delete** /api/v1/hosted-providers/{provider}/models/{model_id} | Remove Hosted Model
[**HostedProvidersUpdateHostedModel**](HostedProvidersAPI.md#HostedProvidersUpdateHostedModel) | **Patch** /api/v1/hosted-providers/{provider}/models/{model_id} | Update Hosted Model
[**HostedProvidersUpdateHostedProvider**](HostedProvidersAPI.md#HostedProvidersUpdateHostedProvider) | **Patch** /api/v1/hosted-providers/{provider} | Update Hosted Provider



## HostedProvidersAddHostedModel

> HostedModelPublic HostedProvidersAddHostedModel(ctx, provider).HostedModelCreateRequest(hostedModelCreateRequest).Execute()

Add Hosted Model



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
	provider := "provider_example" // string | 
	hostedModelCreateRequest := *openapiclient.NewHostedModelCreateRequest("Model_example") // HostedModelCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersAddHostedModel(context.Background(), provider).HostedModelCreateRequest(hostedModelCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersAddHostedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersAddHostedModel`: HostedModelPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersAddHostedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersAddHostedModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostedModelCreateRequest** | [**HostedModelCreateRequest**](HostedModelCreateRequest.md) |  | 

### Return type

[**HostedModelPublic**](HostedModelPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersCreateHostedProvider

> HostedProviderPublic HostedProvidersCreateHostedProvider(ctx).HostedProviderCreateRequest(hostedProviderCreateRequest).Execute()

Create Hosted Provider



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
	hostedProviderCreateRequest := *openapiclient.NewHostedProviderCreateRequest("ApiKey_example", "Provider_example") // HostedProviderCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersCreateHostedProvider(context.Background()).HostedProviderCreateRequest(hostedProviderCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersCreateHostedProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersCreateHostedProvider`: HostedProviderPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersCreateHostedProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersCreateHostedProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostedProviderCreateRequest** | [**HostedProviderCreateRequest**](HostedProviderCreateRequest.md) |  | 

### Return type

[**HostedProviderPublic**](HostedProviderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersDeleteHostedProvider

> Message HostedProvidersDeleteHostedProvider(ctx, provider).Execute()

Delete Hosted Provider



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersDeleteHostedProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersDeleteHostedProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersDeleteHostedProvider`: Message
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersDeleteHostedProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersDeleteHostedProviderRequest struct via the builder pattern


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


## HostedProvidersListAvailableModels

> HostedAvailableModelsPublic HostedProvidersListAvailableModels(ctx, provider).Execute()

List Available Models



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersListAvailableModels(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersListAvailableModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersListAvailableModels`: HostedAvailableModelsPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersListAvailableModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersListAvailableModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedAvailableModelsPublic**](HostedAvailableModelsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersListHostedModels

> HostedModelsPublic HostedProvidersListHostedModels(ctx, provider).Skip(skip).Limit(limit).Execute()

List Hosted Models



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
	provider := "provider_example" // string | 
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersListHostedModels(context.Background(), provider).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersListHostedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersListHostedModels`: HostedModelsPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersListHostedModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersListHostedModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 500]

### Return type

[**HostedModelsPublic**](HostedModelsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersListHostedProviders

> HostedProvidersPublic HostedProvidersListHostedProviders(ctx).Skip(skip).Limit(limit).Execute()

List Hosted Providers



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersListHostedProviders(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersListHostedProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersListHostedProviders`: HostedProvidersPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersListHostedProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersListHostedProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**HostedProvidersPublic**](HostedProvidersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersPreviewCatalogRefresh

> HostedCatalogRefreshPublic HostedProvidersPreviewCatalogRefresh(ctx).Execute()

Preview Catalog Refresh



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
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersPreviewCatalogRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersPreviewCatalogRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersPreviewCatalogRefresh`: HostedCatalogRefreshPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersPreviewCatalogRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersPreviewCatalogRefreshRequest struct via the builder pattern


### Return type

[**HostedCatalogRefreshPublic**](HostedCatalogRefreshPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersRefreshCatalog

> HostedCatalogRefreshPublic HostedProvidersRefreshCatalog(ctx).Execute()

Refresh Catalog



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
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersRefreshCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersRefreshCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersRefreshCatalog`: HostedCatalogRefreshPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersRefreshCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersRefreshCatalogRequest struct via the builder pattern


### Return type

[**HostedCatalogRefreshPublic**](HostedCatalogRefreshPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersRefreshHostedModels

> HostedModelsRefreshPublic HostedProvidersRefreshHostedModels(ctx, provider).Execute()

Refresh Hosted Models



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersRefreshHostedModels(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersRefreshHostedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersRefreshHostedModels`: HostedModelsRefreshPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersRefreshHostedModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersRefreshHostedModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedModelsRefreshPublic**](HostedModelsRefreshPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersRemoveHostedModel

> Message HostedProvidersRemoveHostedModel(ctx, provider, modelId).Execute()

Remove Hosted Model



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
	provider := "provider_example" // string | 
	modelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersRemoveHostedModel(context.Background(), provider, modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersRemoveHostedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersRemoveHostedModel`: Message
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersRemoveHostedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersRemoveHostedModelRequest struct via the builder pattern


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


## HostedProvidersUpdateHostedModel

> HostedModelPublic HostedProvidersUpdateHostedModel(ctx, provider, modelId).HostedModelUpdateRequest(hostedModelUpdateRequest).Execute()

Update Hosted Model



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
	provider := "provider_example" // string | 
	modelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	hostedModelUpdateRequest := *openapiclient.NewHostedModelUpdateRequest() // HostedModelUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersUpdateHostedModel(context.Background(), provider, modelId).HostedModelUpdateRequest(hostedModelUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersUpdateHostedModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersUpdateHostedModel`: HostedModelPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersUpdateHostedModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersUpdateHostedModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostedModelUpdateRequest** | [**HostedModelUpdateRequest**](HostedModelUpdateRequest.md) |  | 

### Return type

[**HostedModelPublic**](HostedModelPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostedProvidersUpdateHostedProvider

> HostedProviderPublic HostedProvidersUpdateHostedProvider(ctx, provider).HostedProviderUpdateRequest(hostedProviderUpdateRequest).Execute()

Update Hosted Provider



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
	provider := "provider_example" // string | 
	hostedProviderUpdateRequest := *openapiclient.NewHostedProviderUpdateRequest() // HostedProviderUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedProvidersAPI.HostedProvidersUpdateHostedProvider(context.Background(), provider).HostedProviderUpdateRequest(hostedProviderUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedProvidersAPI.HostedProvidersUpdateHostedProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostedProvidersUpdateHostedProvider`: HostedProviderPublic
	fmt.Fprintf(os.Stdout, "Response from `HostedProvidersAPI.HostedProvidersUpdateHostedProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostedProvidersUpdateHostedProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostedProviderUpdateRequest** | [**HostedProviderUpdateRequest**](HostedProviderUpdateRequest.md) |  | 

### Return type

[**HostedProviderPublic**](HostedProviderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

