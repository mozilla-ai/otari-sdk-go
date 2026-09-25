# \PricingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingConfirmPricingRefresh**](PricingAPI.md#PricingConfirmPricingRefresh) | **Post** /api/v1/pricing/refresh/confirm | Confirm Pricing Refresh
[**PricingDeletePricing**](PricingAPI.md#PricingDeletePricing) | **Delete** /api/v1/pricing/{model_key} | Delete Pricing
[**PricingGetPendingPricingRefresh**](PricingAPI.md#PricingGetPendingPricingRefresh) | **Get** /api/v1/pricing/refresh/pending | Get Pending Pricing Refresh
[**PricingGetPricing**](PricingAPI.md#PricingGetPricing) | **Get** /api/v1/pricing/{model_key} | Get Pricing
[**PricingGetPricingHistory**](PricingAPI.md#PricingGetPricingHistory) | **Get** /api/v1/pricing/{model_key}/history | Get Pricing History
[**PricingListCurrentPricing**](PricingAPI.md#PricingListCurrentPricing) | **Get** /api/v1/pricing/current | List Current Pricing
[**PricingListPricing**](PricingAPI.md#PricingListPricing) | **Get** /api/v1/pricing | List Pricing
[**PricingListPricingDrift**](PricingAPI.md#PricingListPricingDrift) | **Get** /api/v1/pricing/drift | List Pricing Drift
[**PricingListPricingSnapshots**](PricingAPI.md#PricingListPricingSnapshots) | **Get** /api/v1/pricing/snapshots | List Pricing Snapshots
[**PricingPreviewPricingRefresh**](PricingAPI.md#PricingPreviewPricingRefresh) | **Post** /api/v1/pricing/refresh | Preview Pricing Refresh
[**PricingRejectPricingRefresh**](PricingAPI.md#PricingRejectPricingRefresh) | **Post** /api/v1/pricing/refresh/reject | Reject Pricing Refresh
[**PricingSetPricing**](PricingAPI.md#PricingSetPricing) | **Post** /api/v1/pricing | Set Pricing



## PricingConfirmPricingRefresh

> PricingRefreshConfirmationResponse PricingConfirmPricingRefresh(ctx).Execute()

Confirm Pricing Refresh



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
	resp, r, err := apiClient.PricingAPI.PricingConfirmPricingRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingConfirmPricingRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingConfirmPricingRefresh`: PricingRefreshConfirmationResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingConfirmPricingRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingConfirmPricingRefreshRequest struct via the builder pattern


### Return type

[**PricingRefreshConfirmationResponse**](PricingRefreshConfirmationResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingDeletePricing

> PricingDeletePricing(ctx, modelKey).EffectiveAt(effectiveAt).Execute()

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
	r, err := apiClient.PricingAPI.PricingDeletePricing(context.Background(), modelKey).EffectiveAt(effectiveAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingDeletePricing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPricingDeletePricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **effectiveAt** | **time.Time** | ISO datetime identifying a specific pricing row to delete | 

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


## PricingGetPendingPricingRefresh

> PricingRefreshPreviewResponse PricingGetPendingPricingRefresh(ctx).Execute()

Get Pending Pricing Refresh



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
	resp, r, err := apiClient.PricingAPI.PricingGetPendingPricingRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingGetPendingPricingRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetPendingPricingRefresh`: PricingRefreshPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingGetPendingPricingRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetPendingPricingRefreshRequest struct via the builder pattern


### Return type

[**PricingRefreshPreviewResponse**](PricingRefreshPreviewResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetPricing

> PricingResponse PricingGetPricing(ctx, modelKey).AsOf(asOf).Execute()

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
	resp, r, err := apiClient.PricingAPI.PricingGetPricing(context.Background(), modelKey).AsOf(asOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingGetPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetPricing`: PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingGetPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asOf** | **time.Time** | ISO datetime for effective lookup | 

### Return type

[**PricingResponse**](PricingResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetPricingHistory

> []PricingResponse PricingGetPricingHistory(ctx, modelKey).Execute()

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
	resp, r, err := apiClient.PricingAPI.PricingGetPricingHistory(context.Background(), modelKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingGetPricingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetPricingHistory`: []PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingGetPricingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetPricingHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PricingResponse**](PricingResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListCurrentPricing

> CurrentPricingPage PricingListCurrentPricing(ctx).Skip(skip).Limit(limit).Execute()

List Current Pricing



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
	resp, r, err := apiClient.PricingAPI.PricingListCurrentPricing(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingListCurrentPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListCurrentPricing`: CurrentPricingPage
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingListCurrentPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingListCurrentPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**CurrentPricingPage**](CurrentPricingPage.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListPricing

> []PricingResponse PricingListPricing(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.PricingAPI.PricingListPricing(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingListPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListPricing`: []PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingListPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingListPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]PricingResponse**](PricingResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListPricingDrift

> []PricingDriftRow PricingListPricingDrift(ctx).Limit(limit).Execute()

List Pricing Drift



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
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.PricingListPricingDrift(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingListPricingDrift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListPricingDrift`: []PricingDriftRow
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingListPricingDrift`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingListPricingDriftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 200]

### Return type

[**[]PricingDriftRow**](PricingDriftRow.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListPricingSnapshots

> []AcceptedSnapshotResponse PricingListPricingSnapshots(ctx).Limit(limit).Execute()

List Pricing Snapshots



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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.PricingListPricingSnapshots(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingListPricingSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListPricingSnapshots`: []AcceptedSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingListPricingSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingListPricingSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]AcceptedSnapshotResponse**](AcceptedSnapshotResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingPreviewPricingRefresh

> PricingRefreshPreviewResponse PricingPreviewPricingRefresh(ctx).Execute()

Preview Pricing Refresh



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
	resp, r, err := apiClient.PricingAPI.PricingPreviewPricingRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingPreviewPricingRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingPreviewPricingRefresh`: PricingRefreshPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingPreviewPricingRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingPreviewPricingRefreshRequest struct via the builder pattern


### Return type

[**PricingRefreshPreviewResponse**](PricingRefreshPreviewResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingRejectPricingRefresh

> PricingRejectPricingRefresh(ctx).Execute()

Reject Pricing Refresh



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
	r, err := apiClient.PricingAPI.PricingRejectPricingRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingRejectPricingRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingRejectPricingRefreshRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingSetPricing

> PricingResponse PricingSetPricing(ctx).SetPricingRequest(setPricingRequest).Execute()

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
	resp, r, err := apiClient.PricingAPI.PricingSetPricing(context.Background()).SetPricingRequest(setPricingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PricingSetPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingSetPricing`: PricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PricingSetPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingSetPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPricingRequest** | [**SetPricingRequest**](SetPricingRequest.md) |  | 

### Return type

[**PricingResponse**](PricingResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

