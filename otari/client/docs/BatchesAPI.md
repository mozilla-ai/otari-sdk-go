# \BatchesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchesCancelBatch**](BatchesAPI.md#BatchesCancelBatch) | **Post** /api/v1/batches/{batch_id}/cancel | Cancel Batch
[**BatchesCreateBatch**](BatchesAPI.md#BatchesCreateBatch) | **Post** /api/v1/batches | Create Batch
[**BatchesListBatches**](BatchesAPI.md#BatchesListBatches) | **Get** /api/v1/batches | List Batches
[**BatchesRetrieveBatch**](BatchesAPI.md#BatchesRetrieveBatch) | **Get** /api/v1/batches/{batch_id} | Retrieve Batch
[**BatchesRetrieveBatchResults**](BatchesAPI.md#BatchesRetrieveBatchResults) | **Get** /api/v1/batches/{batch_id}/results | Retrieve Batch Results



## BatchesCancelBatch

> interface{} BatchesCancelBatch(ctx, batchId).Provider(provider).Execute()

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
	resp, r, err := apiClient.BatchesAPI.BatchesCancelBatch(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.BatchesCancelBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchesCancelBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.BatchesCancelBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchesCancelBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchesCreateBatch

> interface{} BatchesCreateBatch(ctx).CreateBatchRequest(createBatchRequest).Execute()

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
	resp, r, err := apiClient.BatchesAPI.BatchesCreateBatch(context.Background()).CreateBatchRequest(createBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.BatchesCreateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchesCreateBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.BatchesCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchesCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchRequest** | [**CreateBatchRequest**](CreateBatchRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchesListBatches

> interface{} BatchesListBatches(ctx).Provider(provider).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.BatchesAPI.BatchesListBatches(context.Background()).Provider(provider).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.BatchesListBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchesListBatches`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.BatchesListBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchesListBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchesRetrieveBatch

> interface{} BatchesRetrieveBatch(ctx, batchId).Provider(provider).Execute()

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
	resp, r, err := apiClient.BatchesAPI.BatchesRetrieveBatch(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.BatchesRetrieveBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchesRetrieveBatch`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.BatchesRetrieveBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchesRetrieveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchesRetrieveBatchResults

> interface{} BatchesRetrieveBatchResults(ctx, batchId).Provider(provider).Execute()

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
	resp, r, err := apiClient.BatchesAPI.BatchesRetrieveBatchResults(context.Background(), batchId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchesAPI.BatchesRetrieveBatchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchesRetrieveBatchResults`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BatchesAPI.BatchesRetrieveBatchResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchesRetrieveBatchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | **string** |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

