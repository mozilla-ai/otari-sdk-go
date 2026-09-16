# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesCreateImage**](ImagesAPI.md#ImagesCreateImage) | **Post** /api/v1/images/generations | Create Image



## ImagesCreateImage

> ImagesResponse ImagesCreateImage(ctx).ImageGenerationRequest(imageGenerationRequest).Execute()

Create Image



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
	imageGenerationRequest := *openapiclient.NewImageGenerationRequest("Model_example", "Prompt_example") // ImageGenerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesCreateImage(context.Background()).ImageGenerationRequest(imageGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesCreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesCreateImage`: ImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesCreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageGenerationRequest** | [**ImageGenerationRequest**](ImageGenerationRequest.md) |  | 

### Return type

[**ImagesResponse**](ImagesResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

