# \AudioAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpeechV1AudioSpeechPost**](AudioAPI.md#CreateSpeechV1AudioSpeechPost) | **Post** /v1/audio/speech | Create Speech
[**CreateTranscriptionV1AudioTranscriptionsPost**](AudioAPI.md#CreateTranscriptionV1AudioTranscriptionsPost) | **Post** /v1/audio/transcriptions | Create Transcription



## CreateSpeechV1AudioSpeechPost

> interface{} CreateSpeechV1AudioSpeechPost(ctx).AudioSpeechRequest(audioSpeechRequest).Execute()

Create Speech



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
	audioSpeechRequest := *openapiclient.NewAudioSpeechRequest("Input_example", "Model_example", "Voice_example") // AudioSpeechRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateSpeechV1AudioSpeechPost(context.Background()).AudioSpeechRequest(audioSpeechRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateSpeechV1AudioSpeechPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpeechV1AudioSpeechPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateSpeechV1AudioSpeechPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpeechV1AudioSpeechPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audioSpeechRequest** | [**AudioSpeechRequest**](AudioSpeechRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, audio/L16, audio/aac, audio/flac, audio/mpeg, audio/opus, audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTranscriptionV1AudioTranscriptionsPost

> interface{} CreateTranscriptionV1AudioTranscriptionsPost(ctx).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).User(user).Execute()

Create Transcription



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
	file := "file_example" // string | 
	model := "model_example" // string | 
	language := "language_example" // string |  (optional)
	prompt := "prompt_example" // string |  (optional)
	responseFormat := "responseFormat_example" // string |  (optional)
	temperature := float32(8.14) // float32 |  (optional)
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.CreateTranscriptionV1AudioTranscriptionsPost(context.Background()).File(file).Model(model).Language(language).Prompt(prompt).ResponseFormat(responseFormat).Temperature(temperature).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.CreateTranscriptionV1AudioTranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTranscriptionV1AudioTranscriptionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.CreateTranscriptionV1AudioTranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranscriptionV1AudioTranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** |  | 
 **model** | **string** |  | 
 **language** | **string** |  | 
 **prompt** | **string** |  | 
 **responseFormat** | **string** |  | 
 **temperature** | **float32** |  | 
 **user** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

