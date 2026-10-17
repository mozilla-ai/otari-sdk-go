# \ChatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatChatCompletions**](ChatAPI.md#ChatChatCompletions) | **Post** /api/v1/chat/completions | Chat Completions



## ChatChatCompletions

> ChatCompletion ChatChatCompletions(ctx).ChatCompletionRequest(chatCompletionRequest).IdempotencyKey(idempotencyKey).OtariTags(otariTags).Execute()

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
	otariTags := "otariTags_example" // string | Tags for cost attribution, as comma-separated `key=value` pairs (`agent=bug-fix,run=1234`), for a client that cannot set the body's `metadata`. Keys and values are percent-decoded, so either may carry an encoded `,` or `=`, and a non-ASCII character must be encoded. Whitespace around a pair is ignored, and a repeated header is read as one list. The bounds are the body's: up to 16 pairs, keys up to 64 characters and values up to 512, counted after merging with the body's tags, which win on the same key. Never forwarded to the provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ChatChatCompletions(context.Background()).ChatCompletionRequest(chatCompletionRequest).IdempotencyKey(idempotencyKey).OtariTags(otariTags).Execute()
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
 **otariTags** | **string** | Tags for cost attribution, as comma-separated &#x60;key&#x3D;value&#x60; pairs (&#x60;agent&#x3D;bug-fix,run&#x3D;1234&#x60;), for a client that cannot set the body&#39;s &#x60;metadata&#x60;. Keys and values are percent-decoded, so either may carry an encoded &#x60;,&#x60; or &#x60;&#x3D;&#x60;, and a non-ASCII character must be encoded. Whitespace around a pair is ignored, and a repeated header is read as one list. The bounds are the body&#39;s: up to 16 pairs, keys up to 64 characters and values up to 512, counted after merging with the body&#39;s tags, which win on the same key. Never forwarded to the provider. | 

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

