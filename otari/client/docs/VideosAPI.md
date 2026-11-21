# \VideosAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideosCreateVideo**](VideosAPI.md#VideosCreateVideo) | **Post** /api/v1/videos | Create Video
[**VideosListVideos**](VideosAPI.md#VideosListVideos) | **Get** /api/v1/videos | List Videos
[**VideosRetrieveVideo**](VideosAPI.md#VideosRetrieveVideo) | **Get** /api/v1/videos/{video_id} | Retrieve Video
[**VideosRetrieveVideoContent**](VideosAPI.md#VideosRetrieveVideoContent) | **Get** /api/v1/videos/{video_id}/content | Retrieve Video Content



## VideosCreateVideo

> VideoJobResponse VideosCreateVideo(ctx).VideoGenerationRequest(videoGenerationRequest).Execute()

Create Video



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
	videoGenerationRequest := *openapiclient.NewVideoGenerationRequest("Model_example", "Prompt_example") // VideoGenerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.VideosCreateVideo(context.Background()).VideoGenerationRequest(videoGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.VideosCreateVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VideosCreateVideo`: VideoJobResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.VideosCreateVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideosCreateVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoGenerationRequest** | [**VideoGenerationRequest**](VideoGenerationRequest.md) |  | 

### Return type

[**VideoJobResponse**](VideoJobResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideosListVideos

> VideoJobList VideosListVideos(ctx).Limit(limit).Execute()

List Videos



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
	limit := int32(56) // int32 | How many of the newest jobs to list (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.VideosListVideos(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.VideosListVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VideosListVideos`: VideoJobList
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.VideosListVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideosListVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | How many of the newest jobs to list | [default to 20]

### Return type

[**VideoJobList**](VideoJobList.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideosRetrieveVideo

> VideoJobResponse VideosRetrieveVideo(ctx, videoId).Execute()

Retrieve Video



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
	videoId := "videoId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.VideosRetrieveVideo(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.VideosRetrieveVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VideosRetrieveVideo`: VideoJobResponse
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.VideosRetrieveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosRetrieveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoJobResponse**](VideoJobResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideosRetrieveVideoContent

> VideosRetrieveVideoContent(ctx, videoId).Execute()

Retrieve Video Content



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
	videoId := "videoId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideosAPI.VideosRetrieveVideoContent(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.VideosRetrieveVideoContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosRetrieveVideoContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

