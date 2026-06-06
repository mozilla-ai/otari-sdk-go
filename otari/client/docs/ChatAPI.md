# \ChatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatCompletionsV1ChatCompletionsPost**](ChatAPI.md#ChatCompletionsV1ChatCompletionsPost) | **Post** /v1/chat/completions | Chat Completions



## ChatCompletionsV1ChatCompletionsPost

> ChatCompletion ChatCompletionsV1ChatCompletionsPost(ctx).ChatCompletionRequest(chatCompletionRequest).Execute()

Chat Completions



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
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest([]openapiclient.ChatMessageInput{*openapiclient.NewChatMessageInput("Content_example", "Role_example", "Name_example", "ToolCallId_example")}, "Model_example") // ChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ChatCompletionsV1ChatCompletionsPost(context.Background()).ChatCompletionRequest(chatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.ChatCompletionsV1ChatCompletionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatCompletionsV1ChatCompletionsPost`: ChatCompletion
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.ChatCompletionsV1ChatCompletionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCompletionsV1ChatCompletionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 

### Return type

[**ChatCompletion**](ChatCompletion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

