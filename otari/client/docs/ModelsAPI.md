# \ModelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetModelV1ModelsModelIdGet**](ModelsAPI.md#GetModelV1ModelsModelIdGet) | **Get** /v1/models/{model_id} | Get Model
[**ListDiscoverableModelsV1ModelsDiscoverableGet**](ModelsAPI.md#ListDiscoverableModelsV1ModelsDiscoverableGet) | **Get** /v1/models/discoverable | List Discoverable Models
[**ListModelMetadataV1ModelsMetadataGet**](ModelsAPI.md#ListModelMetadataV1ModelsMetadataGet) | **Get** /v1/models/metadata | List Model Metadata
[**ListModelsV1ModelsGet**](ModelsAPI.md#ListModelsV1ModelsGet) | **Get** /v1/models | List Models



## GetModelV1ModelsModelIdGet

> ModelObject GetModelV1ModelsModelIdGet(ctx, modelId).Execute()

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
	resp, r, err := apiClient.ModelsAPI.GetModelV1ModelsModelIdGet(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelV1ModelsModelIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelV1ModelsModelIdGet`: ModelObject
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetModelV1ModelsModelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelV1ModelsModelIdGetRequest struct via the builder pattern


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


## ListDiscoverableModelsV1ModelsDiscoverableGet

> DiscoverableModelsResponse ListDiscoverableModelsV1ModelsDiscoverableGet(ctx).Refresh(refresh).Execute()

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
	resp, r, err := apiClient.ModelsAPI.ListDiscoverableModelsV1ModelsDiscoverableGet(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ListDiscoverableModelsV1ModelsDiscoverableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiscoverableModelsV1ModelsDiscoverableGet`: DiscoverableModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ListDiscoverableModelsV1ModelsDiscoverableGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDiscoverableModelsV1ModelsDiscoverableGetRequest struct via the builder pattern


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


## ListModelMetadataV1ModelsMetadataGet

> ModelMetadataResponse ListModelMetadataV1ModelsMetadataGet(ctx).Execute()

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
	resp, r, err := apiClient.ModelsAPI.ListModelMetadataV1ModelsMetadataGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ListModelMetadataV1ModelsMetadataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelMetadataV1ModelsMetadataGet`: ModelMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ListModelMetadataV1ModelsMetadataGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListModelMetadataV1ModelsMetadataGetRequest struct via the builder pattern


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


## ListModelsV1ModelsGet

> ModelListResponse ListModelsV1ModelsGet(ctx).Provider(provider).Execute()

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
	resp, r, err := apiClient.ModelsAPI.ListModelsV1ModelsGet(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ListModelsV1ModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelsV1ModelsGet`: ModelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ListModelsV1ModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListModelsV1ModelsGetRequest struct via the builder pattern


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

