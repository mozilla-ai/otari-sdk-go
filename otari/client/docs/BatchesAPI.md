# \BatchesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBatchV1BatchesBatchIdCancelPost**](BatchesAPI.md#CancelBatchV1BatchesBatchIdCancelPost) | **Post** /v1/batches/{batch_id}/cancel | Cancel Batch
[**CreateBatchV1BatchesPost**](BatchesAPI.md#CreateBatchV1BatchesPost) | **Post** /v1/batches | Create Batch
[**ListBatchesV1BatchesGet**](BatchesAPI.md#ListBatchesV1BatchesGet) | **Get** /v1/batches | List Batches
[**RetrieveBatchResultsV1BatchesBatchIdResultsGet**](BatchesAPI.md#RetrieveBatchResultsV1BatchesBatchIdResultsGet) | **Get** /v1/batches/{batch_id}/results | Retrieve Batch Results
[**RetrieveBatchV1BatchesBatchIdGet**](BatchesAPI.md#RetrieveBatchV1BatchesBatchIdGet) | **Get** /v1/batches/{batch_id} | Retrieve Batch



## CancelBatchV1BatchesBatchIdCancelPost

> interface{} CancelBatchV1BatchesBatchIdCancelPost(ctx, batchId).Provider(provider).Execute()

Cancel Batch



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
	batchId := "batchId_example" // string | 
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.CancelBatchV1BatchesBatchIdCancelPost(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.CancelBatchV1BatchesBatchIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelBatchV1BatchesBatchIdCancelPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.CancelBatchV1BatchesBatchIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBatchV1BatchesBatchIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchV1BatchesPost

> interface{} CreateBatchV1BatchesPost(ctx).CreateBatchRequest(createBatchRequest).Execute()

Create Batch



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
	createBatchRequest := *openapiclient.NewCreateBatchRequest("Model_example", []openapiclient.BatchRequestItem{*openapiclient.NewBatchRequestItem(map[string]interface{}{"key": interface{}(123)}, "CustomId_example")}) // CreateBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.CreateBatchV1BatchesPost(context.Background()).CreateBatchRequest(createBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.CreateBatchV1BatchesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchV1BatchesPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.CreateBatchV1BatchesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchV1BatchesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchRequest** | [**CreateBatchRequest**](CreateBatchRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBatchesV1BatchesGet

> interface{} ListBatchesV1BatchesGet(ctx).Provider(provider).After(after).Limit(limit).Execute()

List Batches



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
	after := "after_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.ListBatchesV1BatchesGet(context.Background()).Provider(provider).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.ListBatchesV1BatchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBatchesV1BatchesGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.ListBatchesV1BatchesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBatchesV1BatchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBatchResultsV1BatchesBatchIdResultsGet

> interface{} RetrieveBatchResultsV1BatchesBatchIdResultsGet(ctx, batchId).Provider(provider).Execute()

Retrieve Batch Results



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
	batchId := "batchId_example" // string | 
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.RetrieveBatchResultsV1BatchesBatchIdResultsGet(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.RetrieveBatchResultsV1BatchesBatchIdResultsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBatchResultsV1BatchesBatchIdResultsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.RetrieveBatchResultsV1BatchesBatchIdResultsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBatchResultsV1BatchesBatchIdResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBatchV1BatchesBatchIdGet

> interface{} RetrieveBatchV1BatchesBatchIdGet(ctx, batchId).Provider(provider).Execute()

Retrieve Batch



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
	batchId := "batchId_example" // string | 
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchesAPI.RetrieveBatchV1BatchesBatchIdGet(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.RetrieveBatchV1BatchesBatchIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBatchV1BatchesBatchIdGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.RetrieveBatchV1BatchesBatchIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBatchV1BatchesBatchIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

