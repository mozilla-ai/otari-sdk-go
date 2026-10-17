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

> CatalogResponse CatalogListCatalog(ctx).AtContext(atContext).Search(search).Skip(skip).Limit(limit).Provider(provider).Vendor(vendor).InputModality(inputModality).OutputModality(outputModality).Capability(capability).MinContext(minContext).MaxInput(maxInput).Pricing(pricing).Source(source).ReleasedWithinDays(releasedWithinDays).Sort(sort).Direction(direction).IncludeFacets(includeFacets).Execute()

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
	search := "search_example" // string | Case-insensitive text in a model's name, vendor, id, selectors, or provider instances. (optional)
	skip := int32(56) // int32 | Number of matching models to skip. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of models to return. (optional) (default to 100)
	provider := []*string{"Inner_example"} // []*string | Match any named provider instance. (optional)
	vendor := []*string{"Inner_example"} // []*string | Match any vendor; an empty value names unknown vendors. (optional)
	inputModality := []*string{"Inner_example"} // []*string | Require every input modality. (optional)
	outputModality := []*string{"Inner_example"} // []*string | Require every output modality. (optional)
	capability := []openapiclient.CatalogCapability{openapiclient.CatalogCapability("tool_call")} // []CatalogCapability | Require every capability. (optional)
	minContext := int32(56) // int32 | Minimum context window; unknown windows do not match. (optional) (default to 0)
	maxInput := float32(8.14) // float32 | Maximum cheapest input price per million tokens; unpriced models do not match. (optional)
	pricing := "pricing_example" // string |  (optional) (default to "all")
	source := "source_example" // string |  (optional) (default to "all")
	releasedWithinDays := int32(56) // int32 | Release window ending today (UTC); zero disables it. Unknown and future releases do not match. (optional) (default to 0)
	sort := "sort_example" // string |  (optional) (default to "name")
	direction := "direction_example" // string |  (optional) (default to "asc")
	includeFacets := true // bool | Include the filter choices drawn from the whole authorized catalog. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogListCatalog(context.Background()).AtContext(atContext).Search(search).Skip(skip).Limit(limit).Provider(provider).Vendor(vendor).InputModality(inputModality).OutputModality(outputModality).Capability(capability).MinContext(minContext).MaxInput(maxInput).Pricing(pricing).Source(source).ReleasedWithinDays(releasedWithinDays).Sort(sort).Direction(direction).IncludeFacets(includeFacets).Execute()
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
 **search** | **string** | Case-insensitive text in a model&#39;s name, vendor, id, selectors, or provider instances. | 
 **skip** | **int32** | Number of matching models to skip. | [default to 0]
 **limit** | **int32** | Maximum number of models to return. | [default to 100]
 **provider** | **[]string** | Match any named provider instance. | 
 **vendor** | **[]string** | Match any vendor; an empty value names unknown vendors. | 
 **inputModality** | **[]string** | Require every input modality. | 
 **outputModality** | **[]string** | Require every output modality. | 
 **capability** | [**[]CatalogCapability**](CatalogCapability.md) | Require every capability. | 
 **minContext** | **int32** | Minimum context window; unknown windows do not match. | [default to 0]
 **maxInput** | **float32** | Maximum cheapest input price per million tokens; unpriced models do not match. | 
 **pricing** | **string** |  | [default to &quot;all&quot;]
 **source** | **string** |  | [default to &quot;all&quot;]
 **releasedWithinDays** | **int32** | Release window ending today (UTC); zero disables it. Unknown and future releases do not match. | [default to 0]
 **sort** | **string** |  | [default to &quot;name&quot;]
 **direction** | **string** |  | [default to &quot;asc&quot;]
 **includeFacets** | **bool** | Include the filter choices drawn from the whole authorized catalog. | [default to false]

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

