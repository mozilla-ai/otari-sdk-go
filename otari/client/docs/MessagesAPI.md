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

> MessageResponse MessagesCreateMessage(ctx).MessagesRequest(messagesRequest).IdempotencyKey(idempotencyKey).OtariTags(otariTags).Execute()

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
	idempotencyKey := "idempotencyKey_example" // string | A unique value, such as a UUID, that makes a non-streaming request safe to retry. A retry with the same key and body returns the original response, request ID and cost without calling the provider or billing again. A retry while the original is still running is answered 409 with Retry-After. Reusing a key for a different body is refused with 422. Ignored for streaming requests and on a deployment without OTARI_SECRET_KEY, which encrypts the stored response. (optional)
	otariTags := "otariTags_example" // string | Tags for cost attribution, as comma-separated `key=value` pairs (`agent=bug-fix,run=1234`), for a client that cannot set the body's `metadata`. Keys and values are percent-decoded, so either may carry an encoded `,` or `=`, and a non-ASCII character must be encoded. Whitespace around a pair is ignored, and a repeated header is read as one list. The bounds are the body's: up to 16 pairs, keys up to 64 characters and values up to 512, counted after merging with the body's tags, which win on the same key. Never forwarded to the provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagesAPI.MessagesCreateMessage(context.Background()).MessagesRequest(messagesRequest).IdempotencyKey(idempotencyKey).OtariTags(otariTags).Execute()
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
 **idempotencyKey** | **string** | A unique value, such as a UUID, that makes a non-streaming request safe to retry. A retry with the same key and body returns the original response, request ID and cost without calling the provider or billing again. A retry while the original is still running is answered 409 with Retry-After. Reusing a key for a different body is refused with 422. Ignored for streaming requests and on a deployment without OTARI_SECRET_KEY, which encrypts the stored response. | 
 **otariTags** | **string** | Tags for cost attribution, as comma-separated &#x60;key&#x3D;value&#x60; pairs (&#x60;agent&#x3D;bug-fix,run&#x3D;1234&#x60;), for a client that cannot set the body&#39;s &#x60;metadata&#x60;. Keys and values are percent-decoded, so either may carry an encoded &#x60;,&#x60; or &#x60;&#x3D;&#x60;, and a non-ASCII character must be encoded. Whitespace around a pair is ignored, and a repeated header is read as one list. The bounds are the body&#39;s: up to 16 pairs, keys up to 64 characters and values up to 512, counted after merging with the body&#39;s tags, which win on the same key. Never forwarded to the provider. | 

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

