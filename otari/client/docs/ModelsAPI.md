# \ModelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelsGetModel**](ModelsAPI.md#ModelsGetModel) | **Get** /api/v1/models/{model_id} | Get Model
[**ModelsListDiscoverableModels**](ModelsAPI.md#ModelsListDiscoverableModels) | **Get** /api/v1/models/discoverable | List Discoverable Models
[**ModelsListModelMetadata**](ModelsAPI.md#ModelsListModelMetadata) | **Get** /api/v1/models/metadata | List Model Metadata
[**ModelsListModels**](ModelsAPI.md#ModelsListModels) | **Get** /api/v1/models | List Models



## ModelsGetModel

> ModelObject ModelsGetModel(ctx, modelId).Execute()

Get Model



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
	modelId := "modelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.ModelsGetModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ModelsGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelsGetModel`: ModelObject
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ModelsGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelsGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelObject**](ModelObject.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsListDiscoverableModels

> DiscoverableModelsResponse ModelsListDiscoverableModels(ctx).Refresh(refresh).Execute()

List Discoverable Models



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
	refresh := true // bool | Re-dial every provider instead of answering from the discovery cache. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.ModelsListDiscoverableModels(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ModelsListDiscoverableModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelsListDiscoverableModels`: DiscoverableModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ModelsListDiscoverableModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModelsListDiscoverableModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Re-dial every provider instead of answering from the discovery cache. | [default to false]

### Return type

[**DiscoverableModelsResponse**](DiscoverableModelsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsListModelMetadata

> ModelMetadataResponse ModelsListModelMetadata(ctx).Execute()

List Model Metadata



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
	resp, r, err := apiClient.ModelsAPI.ModelsListModelMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ModelsListModelMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelsListModelMetadata`: ModelMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ModelsListModelMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiModelsListModelMetadataRequest struct via the builder pattern


### Return type

[**ModelMetadataResponse**](ModelMetadataResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsListModels

> ModelListResponse ModelsListModels(ctx).Provider(provider).Execute()

List Models



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
	provider := "provider_example" // string | Filter models by provider name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.ModelsListModels(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ModelsListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelsListModels`: ModelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ModelsListModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModelsListModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Filter models by provider name | 

### Return type

[**ModelListResponse**](ModelListResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

