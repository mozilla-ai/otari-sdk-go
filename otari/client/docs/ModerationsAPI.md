# \ModerationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModerationsCreateModeration**](ModerationsAPI.md#ModerationsCreateModeration) | **Post** /api/v1/moderations | Create Moderation



## ModerationsCreateModeration

> ModerationResponse ModerationsCreateModeration(ctx).ModerationRequest(moderationRequest).IncludeRaw(includeRaw).Execute()

Create Moderation



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
	moderationRequest := *openapiclient.NewModerationRequest(*openapiclient.NewInput1(), "Model_example") // ModerationRequest | 
	includeRaw := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModerationsAPI.ModerationsCreateModeration(context.Background()).ModerationRequest(moderationRequest).IncludeRaw(includeRaw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationsAPI.ModerationsCreateModeration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModerationsCreateModeration`: ModerationResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationsAPI.ModerationsCreateModeration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModerationsCreateModerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderationRequest** | [**ModerationRequest**](ModerationRequest.md) |  | 
 **includeRaw** | **bool** |  | [default to false]

### Return type

[**ModerationResponse**](ModerationResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

