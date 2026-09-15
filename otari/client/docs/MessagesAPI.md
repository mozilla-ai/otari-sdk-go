# \MessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessagesCountMessageTokens**](MessagesAPI.md#MessagesCountMessageTokens) | **Post** /api/v1/messages/count_tokens | Count Message Tokens
[**MessagesCreateMessage**](MessagesAPI.md#MessagesCreateMessage) | **Post** /api/v1/messages | Create Message



## MessagesCountMessageTokens

> CountTokensResponse MessagesCountMessageTokens(ctx).CountTokensRequest(countTokensRequest).Execute()

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
	resp, r, err := apiClient.MessagesAPI.MessagesCountMessageTokens(context.Background()).CountTokensRequest(countTokensRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesCountMessageTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MessagesCountMessageTokens`: CountTokensResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesCountMessageTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesCountMessageTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countTokensRequest** | [**CountTokensRequest**](CountTokensRequest.md) |  | 

### Return type

[**CountTokensResponse**](CountTokensResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesCreateMessage

> MessageResponse MessagesCreateMessage(ctx).MessagesRequest(messagesRequest).Execute()

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
	resp, r, err := apiClient.MessagesAPI.MessagesCreateMessage(context.Background()).MessagesRequest(messagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesCreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MessagesCreateMessage`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesCreateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messagesRequest** | [**MessagesRequest**](MessagesRequest.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

