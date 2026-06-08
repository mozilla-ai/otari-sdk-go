# \EmbeddingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmbeddingV1EmbeddingsPost**](EmbeddingsAPI.md#CreateEmbeddingV1EmbeddingsPost) | **Post** /v1/embeddings | Create Embedding



## CreateEmbeddingV1EmbeddingsPost

> CreateEmbeddingResponse CreateEmbeddingV1EmbeddingsPost(ctx).EmbeddingRequest(embeddingRequest).Execute()

Create Embedding



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
	embeddingRequest := *openapiclient.NewEmbeddingRequest(*openapiclient.NewInput(), "Model_example") // EmbeddingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmbeddingsAPI.CreateEmbeddingV1EmbeddingsPost(context.Background()).EmbeddingRequest(embeddingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbeddingsAPI.CreateEmbeddingV1EmbeddingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmbeddingV1EmbeddingsPost`: CreateEmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `EmbeddingsAPI.CreateEmbeddingV1EmbeddingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmbeddingV1EmbeddingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeddingRequest** | [**EmbeddingRequest**](EmbeddingRequest.md) |  | 

### Return type

[**CreateEmbeddingResponse**](CreateEmbeddingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

