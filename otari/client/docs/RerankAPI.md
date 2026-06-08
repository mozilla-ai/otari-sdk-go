# \RerankAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRerankV1RerankPost**](RerankAPI.md#CreateRerankV1RerankPost) | **Post** /v1/rerank | Create Rerank



## CreateRerankV1RerankPost

> RerankResponse CreateRerankV1RerankPost(ctx).RerankRequest(rerankRequest).Execute()

Create Rerank



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
	rerankRequest := *openapiclient.NewRerankRequest([]string{"Documents_example"}, "Model_example", "Query_example") // RerankRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RerankAPI.CreateRerankV1RerankPost(context.Background()).RerankRequest(rerankRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RerankAPI.CreateRerankV1RerankPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRerankV1RerankPost`: RerankResponse
	fmt.Fprintf(os.Stdout, "Response from `RerankAPI.CreateRerankV1RerankPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRerankV1RerankPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rerankRequest** | [**RerankRequest**](RerankRequest.md) |  | 

### Return type

[**RerankResponse**](RerankResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

