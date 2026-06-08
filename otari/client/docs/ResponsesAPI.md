# \ResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResponseV1ResponsesPost**](ResponsesAPI.md#CreateResponseV1ResponsesPost) | **Post** /v1/responses | Create Response



## CreateResponseV1ResponsesPost

> interface{} CreateResponseV1ResponsesPost(ctx).ResponsesRequest(responsesRequest).Execute()

Create Response



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
	responsesRequest := *openapiclient.NewResponsesRequest(interface{}(123), "Model_example") // ResponsesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResponsesAPI.CreateResponseV1ResponsesPost(context.Background()).ResponsesRequest(responsesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResponsesAPI.CreateResponseV1ResponsesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResponseV1ResponsesPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ResponsesAPI.CreateResponseV1ResponsesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResponseV1ResponsesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responsesRequest** | [**ResponsesRequest**](ResponsesRequest.md) |  | 

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

