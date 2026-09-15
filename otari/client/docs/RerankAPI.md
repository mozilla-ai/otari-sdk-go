# \RerankAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RerankCreateRerank**](RerankAPI.md#RerankCreateRerank) | **Post** /api/v1/rerank | Create Rerank



## RerankCreateRerank

> RerankResponse RerankCreateRerank(ctx).RerankRequest(rerankRequest).Execute()

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
	resp, r, err := apiClient.RerankAPI.RerankCreateRerank(context.Background()).RerankRequest(rerankRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RerankAPI.RerankCreateRerank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RerankCreateRerank`: RerankResponse
	fmt.Fprintf(os.Stdout, "Response from `RerankAPI.RerankCreateRerank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRerankCreateRerankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rerankRequest** | [**RerankRequest**](RerankRequest.md) |  | 

### Return type

[**RerankResponse**](RerankResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

