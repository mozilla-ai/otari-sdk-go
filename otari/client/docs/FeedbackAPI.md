# \FeedbackAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeedbackSubmitFeedback**](FeedbackAPI.md#FeedbackSubmitFeedback) | **Post** /api/v1/feedback | Submit Feedback



## FeedbackSubmitFeedback

> FeedbackSubmitFeedback(ctx).FeedbackSubmission(feedbackSubmission).Execute()

Submit Feedback



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
	feedbackSubmission := *openapiclient.NewFeedbackSubmission("Message_example") // FeedbackSubmission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeedbackAPI.FeedbackSubmitFeedback(context.Background()).FeedbackSubmission(feedbackSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.FeedbackSubmitFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedbackSubmitFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackSubmission** | [**FeedbackSubmission**](FeedbackSubmission.md) |  | 

### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

