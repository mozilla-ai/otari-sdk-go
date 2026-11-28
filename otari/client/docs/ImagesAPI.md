# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesCreateImage**](ImagesAPI.md#ImagesCreateImage) | **Post** /api/v1/images/generations | Create Image
[**ImagesEditImage**](ImagesAPI.md#ImagesEditImage) | **Post** /api/v1/images/edits | Edit Image



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


## ImagesEditImage

> interface{} ImagesEditImage(ctx).Model(model).Prompt(prompt).Image(image).Image2(image2).Mask(mask).N(n).Quality(quality).ResponseFormat(responseFormat).Size(size).User(user).Execute()

Edit Image



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
	model := "model_example" // string | 
	prompt := "prompt_example" // string | 
	image := []string{"Inner_example"} // []string | The image or images to edit. (optional)
	image2 := []string{"Inner_example"} // []string | Same as image. (optional)
	mask := "mask_example" // string | Where to edit: the transparent areas of this PNG. (optional)
	n := int32(56) // int32 |  (optional)
	quality := "quality_example" // string | Filter to a single event type or metric name (e.g. 'tool_result', 'claude_code.commit.count') (optional)
	responseFormat := "responseFormat_example" // string |  (optional)
	size := "size_example" // string | Filter to a single event type or metric name (e.g. 'tool_result', 'claude_code.commit.count') (optional)
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ImagesEditImage(context.Background()).Model(model).Prompt(prompt).Image(image).Image2(image2).Mask(mask).N(n).Quality(quality).ResponseFormat(responseFormat).Size(size).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ImagesEditImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagesEditImage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ImagesEditImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesEditImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** |  | 
 **prompt** | **string** |  | 
 **image** | **[]string** | The image or images to edit. | 
 **image2** | **[]string** | Same as image. | 
 **mask** | **string** | Where to edit: the transparent areas of this PNG. | 
 **n** | **int32** |  | 
 **quality** | **string** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | 
 **responseFormat** | **string** |  | 
 **size** | **string** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | 
 **user** | **string** |  | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

