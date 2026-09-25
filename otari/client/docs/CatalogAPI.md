# \CatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogGetCatalogModel**](CatalogAPI.md#CatalogGetCatalogModel) | **Get** /api/v1/catalog/models/{model_id} | Get Catalog Model
[**CatalogListCatalog**](CatalogAPI.md#CatalogListCatalog) | **Get** /api/v1/catalog/models | List Catalog
[**CatalogRefreshSelectorIndex**](CatalogAPI.md#CatalogRefreshSelectorIndex) | **Post** /api/v1/catalog/selectors/refresh | Refresh Selector Index



## CatalogGetCatalogModel

> CatalogModelDetail CatalogGetCatalogModel(ctx, modelId).Execute()

Get Catalog Model



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
	resp, r, err := apiClient.CatalogAPI.CatalogGetCatalogModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CatalogGetCatalogModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogGetCatalogModel`: CatalogModelDetail
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CatalogGetCatalogModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogGetCatalogModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogModelDetail**](CatalogModelDetail.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogListCatalog

> CatalogResponse CatalogListCatalog(ctx).AtContext(atContext).Search(search).Skip(skip).Limit(limit).Execute()

List Catalog



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
	atContext := int32(56) // int32 | Compare prices for a request of this many input tokens: each model's minimum is taken from the pricing tier that request would settle at. Omitted, the base rates compare. (optional)
	search := "search_example" // string | Narrow to models whose name, catalog id or any selector contains this text, case-insensitively. (optional)
	skip := int32(56) // int32 | Number of models to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of models to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogListCatalog(context.Background()).AtContext(atContext).Search(search).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CatalogListCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogListCatalog`: CatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CatalogListCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogListCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atContext** | **int32** | Compare prices for a request of this many input tokens: each model&#39;s minimum is taken from the pricing tier that request would settle at. Omitted, the base rates compare. | 
 **search** | **string** | Narrow to models whose name, catalog id or any selector contains this text, case-insensitively. | 
 **skip** | **int32** | Number of models to skip | [default to 0]
 **limit** | **int32** | Maximum number of models to return | [default to 100]

### Return type

[**CatalogResponse**](CatalogResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogRefreshSelectorIndex

> SelectorIndexResponse CatalogRefreshSelectorIndex(ctx).Execute()

Refresh Selector Index



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
	resp, r, err := apiClient.CatalogAPI.CatalogRefreshSelectorIndex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CatalogRefreshSelectorIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogRefreshSelectorIndex`: SelectorIndexResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CatalogRefreshSelectorIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogRefreshSelectorIndexRequest struct via the builder pattern


### Return type

[**SelectorIndexResponse**](SelectorIndexResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

