# \ChatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatChatCompletions**](ChatAPI.md#ChatChatCompletions) | **Post** /api/v1/chat/completions | Chat Completions



## ChatChatCompletions

> ChatCompletion ChatChatCompletions(ctx).ChatCompletionRequest(chatCompletionRequest).IdempotencyKey(idempotencyKey).Execute()

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
	idempotencyKey := "idempotencyKey_example" // string | A unique value, such as a UUID, that makes a non-streaming request safe to retry. A retry with the same key and body returns the original response, request ID and cost without calling the provider or billing again. A retry while the original is still running is answered 409 with Retry-After. Reusing a key for a different body is refused with 422. Ignored for streaming requests, in hybrid mode, and on a deployment without OTARI_SECRET_KEY, which encrypts the stored response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ChatChatCompletions(context.Background()).ChatCompletionRequest(chatCompletionRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.ChatChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatChatCompletions`: ChatCompletion
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.ChatChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 
 **idempotencyKey** | **string** | A unique value, such as a UUID, that makes a non-streaming request safe to retry. A retry with the same key and body returns the original response, request ID and cost without calling the provider or billing again. A retry while the original is still running is answered 409 with Retry-After. Reusing a key for a different body is refused with 422. Ignored for streaming requests, in hybrid mode, and on a deployment without OTARI_SECRET_KEY, which encrypts the stored response. | 

### Return type

[**ChatCompletion**](ChatCompletion.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

