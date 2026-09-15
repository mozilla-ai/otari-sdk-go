# \ResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResponsesCreateResponse**](ResponsesAPI.md#ResponsesCreateResponse) | **Post** /api/v1/responses | Create Response



## ResponsesCreateResponse

> interface{} ResponsesCreateResponse(ctx).ResponsesRequest(responsesRequest).Execute()

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
	resp, r, err := apiClient.ResponsesAPI.ResponsesCreateResponse(context.Background()).ResponsesRequest(responsesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResponsesAPI.ResponsesCreateResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResponsesCreateResponse`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ResponsesAPI.ResponsesCreateResponse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResponsesCreateResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responsesRequest** | [**ResponsesRequest**](ResponsesRequest.md) |  | 

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

