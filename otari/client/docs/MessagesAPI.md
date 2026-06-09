# \MessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountMessageTokensV1MessagesCountTokensPost**](MessagesAPI.md#CountMessageTokensV1MessagesCountTokensPost) | **Post** /v1/messages/count_tokens | Count Message Tokens
[**CreateMessageV1MessagesPost**](MessagesAPI.md#CreateMessageV1MessagesPost) | **Post** /v1/messages | Create Message



## CountMessageTokensV1MessagesCountTokensPost

> CountTokensResponse CountMessageTokensV1MessagesCountTokensPost(ctx).CountTokensRequest(countTokensRequest).Execute()

Count Message Tokens



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
	countTokensRequest := *openapiclient.NewCountTokensRequest([]*map[string]interface{}{nil}, "Model_example") // CountTokensRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.CountMessageTokensV1MessagesCountTokensPost(context.Background()).CountTokensRequest(countTokensRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CountMessageTokensV1MessagesCountTokensPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMessageTokensV1MessagesCountTokensPost`: CountTokensResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CountMessageTokensV1MessagesCountTokensPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountMessageTokensV1MessagesCountTokensPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countTokensRequest** | [**CountTokensRequest**](CountTokensRequest.md) |  | 

### Return type

[**CountTokensResponse**](CountTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	messagesRequest := *openapiclient.NewMessagesRequest(int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, "Model_example") // MessagesRequest | 

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

