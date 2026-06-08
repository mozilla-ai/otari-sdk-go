# \PricingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePricingV1PricingModelKeyDelete**](PricingAPI.md#DeletePricingV1PricingModelKeyDelete) | **Delete** /v1/pricing/{model_key} | Delete Pricing
[**GetPricingHistoryV1PricingModelKeyHistoryGet**](PricingAPI.md#GetPricingHistoryV1PricingModelKeyHistoryGet) | **Get** /v1/pricing/{model_key}/history | Get Pricing History
[**GetPricingV1PricingModelKeyGet**](PricingAPI.md#GetPricingV1PricingModelKeyGet) | **Get** /v1/pricing/{model_key} | Get Pricing
[**ListPricingV1PricingGet**](PricingAPI.md#ListPricingV1PricingGet) | **Get** /v1/pricing | List Pricing
[**SetPricingV1PricingPost**](PricingAPI.md#SetPricingV1PricingPost) | **Post** /v1/pricing | Set Pricing



## DeletePricingV1PricingModelKeyDelete

> DeletePricingV1PricingModelKeyDelete(ctx, modelKey).EffectiveAt(effectiveAt).Execute()

Delete Pricing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	modelKey := "modelKey_example" // string | 
	effectiveAt := time.Now() // time.Time | ISO datetime identifying a specific pricing row to delete (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PricingAPI.DeletePricingV1PricingModelKeyDelete(context.Background(), modelKey).EffectiveAt(effectiveAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.DeletePricingV1PricingModelKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePricingV1PricingModelKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **effectiveAt** | **time.Time** | ISO datetime identifying a specific pricing row to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingHistoryV1PricingModelKeyHistoryGet

> []PricingResponse GetPricingHistoryV1PricingModelKeyHistoryGet(ctx, modelKey).Execute()

Get Pricing History



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
	modelKey := "modelKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.GetPricingHistoryV1PricingModelKeyHistoryGet(context.Background(), modelKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingHistoryV1PricingModelKeyHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingHistoryV1PricingModelKeyHistoryGet`: []PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingHistoryV1PricingModelKeyHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingHistoryV1PricingModelKeyHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PricingResponse**](PricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingV1PricingModelKeyGet

> PricingResponse GetPricingV1PricingModelKeyGet(ctx, modelKey).AsOf(asOf).Execute()

Get Pricing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	modelKey := "modelKey_example" // string | 
	asOf := time.Now() // time.Time | ISO datetime for effective lookup (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.GetPricingV1PricingModelKeyGet(context.Background(), modelKey).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingV1PricingModelKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingV1PricingModelKeyGet`: PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingV1PricingModelKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingV1PricingModelKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | ISO datetime for effective lookup | 

### Return type

[**PricingResponse**](PricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPricingV1PricingGet

> []PricingResponse ListPricingV1PricingGet(ctx).Skip(skip).Limit(limit).Execute()

List Pricing



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
	resp, r, err := apiClient.PricingAPI.ListPricingV1PricingGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.ListPricingV1PricingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPricingV1PricingGet`: []PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.ListPricingV1PricingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPricingV1PricingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]PricingResponse**](PricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPricingV1PricingPost

> PricingResponse SetPricingV1PricingPost(ctx).SetPricingRequest(setPricingRequest).Execute()

Set Pricing



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
	setPricingRequest := *openapiclient.NewSetPricingRequest(float32(123), "ModelKey_example", float32(123)) // SetPricingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.SetPricingV1PricingPost(context.Background()).SetPricingRequest(setPricingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.SetPricingV1PricingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPricingV1PricingPost`: PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.SetPricingV1PricingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPricingV1PricingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPricingRequest** | [**SetPricingRequest**](SetPricingRequest.md) |  | 

### Return type

[**PricingResponse**](PricingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

