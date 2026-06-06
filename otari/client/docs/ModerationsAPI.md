# \ModerationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModerationV1ModerationsPost**](ModerationsAPI.md#CreateModerationV1ModerationsPost) | **Post** /v1/moderations | Create Moderation



## CreateModerationV1ModerationsPost

> ModerationResponse CreateModerationV1ModerationsPost(ctx).ModerationRequest(moderationRequest).IncludeRaw(includeRaw).Execute()

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
	resp, r, err := apiClient.ModerationsAPI.CreateModerationV1ModerationsPost(context.Background()).ModerationRequest(moderationRequest).IncludeRaw(includeRaw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModerationsAPI.CreateModerationV1ModerationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateModerationV1ModerationsPost`: ModerationResponse
	fmt.Fprintf(os.Stdout, "Response from `ModerationsAPI.CreateModerationV1ModerationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModerationV1ModerationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderationRequest** | [**ModerationRequest**](ModerationRequest.md) |  | 
 **includeRaw** | **bool** |  | [default to false]

### Return type

[**ModerationResponse**](ModerationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

