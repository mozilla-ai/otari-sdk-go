# \MessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageV1MessagesPost**](MessagesAPI.md#CreateMessageV1MessagesPost) | **Post** /v1/messages | Create Message



## CreateMessageV1MessagesPost

> MessageResponse CreateMessageV1MessagesPost(ctx).MessagesRequest(messagesRequest).Execute()

Create Message



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
	messagesRequest := *openapiclient.NewMessagesRequest(int32(123), []*map[string]interface{}{nil}, "Model_example") // MessagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CreateMessageV1MessagesPost(context.Background()).MessagesRequest(messagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateMessageV1MessagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessageV1MessagesPost`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateMessageV1MessagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageV1MessagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messagesRequest** | [**MessagesRequest**](MessagesRequest.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

