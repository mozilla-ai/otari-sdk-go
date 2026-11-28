# \CatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogGetCatalogModel**](CatalogAPI.md#CatalogGetCatalogModel) | **Get** /api/v1/catalog/models/{model_id} | Get Catalog Model
[**CatalogGetCatalogProvider**](CatalogAPI.md#CatalogGetCatalogProvider) | **Get** /api/v1/catalog/providers/{name} | Get Catalog Provider
[**CatalogListCatalog**](CatalogAPI.md#CatalogListCatalog) | **Get** /api/v1/catalog/models | List Catalog
[**CatalogListCatalogProviders**](CatalogAPI.md#CatalogListCatalogProviders) | **Get** /api/v1/catalog/providers | List Catalog Providers
[**CatalogRefreshSelectorIndex**](CatalogAPI.md#CatalogRefreshSelectorIndex) | **Post** /api/v1/catalog/selectors/refresh | Refresh Selector Index



## CatalogGetCatalogModel

> CatalogModelDetail CatalogGetCatalogModel(ctx, modelId).WorkspaceId(workspaceId).Execute()

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | For a signed-in caller, only the offerings a request billed to this workspace may call. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogGetCatalogModel(context.Background(), modelId).WorkspaceId(workspaceId).Execute()
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

 **workspaceId** | **string** | For a signed-in caller, only the offerings a request billed to this workspace may call. | 

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


## CatalogGetCatalogProvider

> CatalogProviderDetail CatalogGetCatalogProvider(ctx, name).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()

Get Catalog Provider



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
	name := "name_example" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | For a signed-in caller, only the models a request billed to this workspace may call. (optional)
	skip := int32(56) // int32 | Number of the provider's models to skip. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of the provider's models to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogGetCatalogProvider(context.Background(), name).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CatalogGetCatalogProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogGetCatalogProvider`: CatalogProviderDetail
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CatalogGetCatalogProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogGetCatalogProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | For a signed-in caller, only the models a request billed to this workspace may call. | 
 **skip** | **int32** | Number of the provider&#39;s models to skip. | [default to 0]
 **limit** | **int32** | Maximum number of the provider&#39;s models to return. | [default to 100]

### Return type

[**CatalogProviderDetail**](CatalogProviderDetail.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogListCatalog

> CatalogResponse CatalogListCatalog(ctx).AtContext(atContext).Search(search).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Provider(provider).Vendor(vendor).InputModality(inputModality).OutputModality(outputModality).Capability(capability).Type_(type_).MinContext(minContext).MaxInput(maxInput).Pricing(pricing).Source(source).ZeroDataRetention(zeroDataRetention).DataLocation(dataLocation).ReleasedWithinDays(releasedWithinDays).Sort(sort).Direction(direction).IncludeFacets(includeFacets).Execute()

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
	search := "search_example" // string | Case-insensitive words, each in a model's name, vendor, id, or a selector's last path segment; a pasted selector, or part of one with a colon or slash, matches too. Provider instances are the `provider` filter's. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | For a signed-in caller, list only what a request billed to this workspace may call, with its own disabled keys and model lists applied. A workspace the caller cannot see answers 404. Ignored for an API key, which already names its workspace. (optional)
	skip := int32(56) // int32 | Number of matching models to skip. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of models to return. (optional) (default to 100)
	provider := []*string{"Inner_example"} // []*string | Match any named provider instance. (optional)
	vendor := []*string{"Inner_example"} // []*string | Match any vendor; an empty value names unknown vendors. (optional)
	inputModality := []*string{"Inner_example"} // []*string | Require every input modality. (optional)
	outputModality := []*string{"Inner_example"} // []*string | Require every output modality. (optional)
	capability := []openapiclient.CatalogCapability{openapiclient.CatalogCapability("tool_call")} // []CatalogCapability | Require every capability. (optional)
	type_ := []string{"Type_example"} // []string | Match any kind; a model of unknown kind does not match. (optional)
	minContext := int32(56) // int32 | Minimum context window; unknown windows do not match. (optional) (default to 0)
	maxInput := float32(8.14) // float32 | Maximum cheapest input price per million tokens; unpriced models, and models priced in another unit, do not match. (optional)
	pricing := "pricing_example" // string |  (optional) (default to "all")
	source := "source_example" // string |  (optional) (default to "all")
	zeroDataRetention := true // bool | Only models at least one provider declared to keep no prompt or completion data offers. (optional) (default to false)
	dataLocation := []*string{"Inner_example"} // []*string | Only models a hosted provider serves from one of these data locations: ISO 3166-1 alpha-2 country codes, `EU` or `GLOBAL`. Matched as a workspace's `allowed_data_locations` is, so `EU` also matches each EU member state and `GLOBAL` only `GLOBAL`. A model whose location is undeclared matches none. (optional)
	releasedWithinDays := int32(56) // int32 | Release window ending today (UTC); zero disables it. Unknown and future releases do not match. (optional) (default to 0)
	sort := "sort_example" // string |  (optional) (default to "name")
	direction := "direction_example" // string |  (optional) (default to "asc")
	includeFacets := true // bool | Include the filter choices drawn from the whole authorized catalog. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogListCatalog(context.Background()).AtContext(atContext).Search(search).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Provider(provider).Vendor(vendor).InputModality(inputModality).OutputModality(outputModality).Capability(capability).Type_(type_).MinContext(minContext).MaxInput(maxInput).Pricing(pricing).Source(source).ZeroDataRetention(zeroDataRetention).DataLocation(dataLocation).ReleasedWithinDays(releasedWithinDays).Sort(sort).Direction(direction).IncludeFacets(includeFacets).Execute()
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
 **search** | **string** | Case-insensitive words, each in a model&#39;s name, vendor, id, or a selector&#39;s last path segment; a pasted selector, or part of one with a colon or slash, matches too. Provider instances are the &#x60;provider&#x60; filter&#39;s. | 
 **workspaceId** | **string** | For a signed-in caller, list only what a request billed to this workspace may call, with its own disabled keys and model lists applied. A workspace the caller cannot see answers 404. Ignored for an API key, which already names its workspace. | 
 **skip** | **int32** | Number of matching models to skip. | [default to 0]
 **limit** | **int32** | Maximum number of models to return. | [default to 100]
 **provider** | **[]string** | Match any named provider instance. | 
 **vendor** | **[]string** | Match any vendor; an empty value names unknown vendors. | 
 **inputModality** | **[]string** | Require every input modality. | 
 **outputModality** | **[]string** | Require every output modality. | 
 **capability** | [**[]CatalogCapability**](CatalogCapability.md) | Require every capability. | 
 **type_** | **[]string** | Match any kind; a model of unknown kind does not match. | 
 **minContext** | **int32** | Minimum context window; unknown windows do not match. | [default to 0]
 **maxInput** | **float32** | Maximum cheapest input price per million tokens; unpriced models, and models priced in another unit, do not match. | 
 **pricing** | **string** |  | [default to &quot;all&quot;]
 **source** | **string** |  | [default to &quot;all&quot;]
 **zeroDataRetention** | **bool** | Only models at least one provider declared to keep no prompt or completion data offers. | [default to false]
 **dataLocation** | **[]string** | Only models a hosted provider serves from one of these data locations: ISO 3166-1 alpha-2 country codes, &#x60;EU&#x60; or &#x60;GLOBAL&#x60;. Matched as a workspace&#39;s &#x60;allowed_data_locations&#x60; is, so &#x60;EU&#x60; also matches each EU member state and &#x60;GLOBAL&#x60; only &#x60;GLOBAL&#x60;. A model whose location is undeclared matches none. | 
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


## CatalogListCatalogProviders

> CatalogProvidersResponse CatalogListCatalogProviders(ctx).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()

List Catalog Providers



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | For a signed-in caller, only the providers a request billed to this workspace may call. (optional)
	skip := int32(56) // int32 | Number of providers to skip. (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of providers to return. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPI.CatalogListCatalogProviders(context.Background()).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPI.CatalogListCatalogProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogListCatalogProviders`: CatalogProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPI.CatalogListCatalogProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogListCatalogProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | For a signed-in caller, only the providers a request billed to this workspace may call. | 
 **skip** | **int32** | Number of providers to skip. | [default to 0]
 **limit** | **int32** | Maximum number of providers to return. | [default to 100]

### Return type

[**CatalogProvidersResponse**](CatalogProvidersResponse.md)

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

