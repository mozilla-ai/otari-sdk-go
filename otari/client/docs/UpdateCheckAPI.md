# \UpdateCheckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCheckAnswerUpdateCheck**](UpdateCheckAPI.md#UpdateCheckAnswerUpdateCheck) | **Put** /api/v1/update-check | Answer Update Check
[**UpdateCheckGetUpdateCheck**](UpdateCheckAPI.md#UpdateCheckGetUpdateCheck) | **Get** /api/v1/update-check | Get Update Check



## UpdateCheckAnswerUpdateCheck

> UpdateCheckStatus UpdateCheckAnswerUpdateCheck(ctx).UpdateCheckAnswer(updateCheckAnswer).Execute()

Answer Update Check



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
	updateCheckAnswer := *openapiclient.NewUpdateCheckAnswer(false) // UpdateCheckAnswer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateCheckAPI.UpdateCheckAnswerUpdateCheck(context.Background()).UpdateCheckAnswer(updateCheckAnswer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateCheckAPI.UpdateCheckAnswerUpdateCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCheckAnswerUpdateCheck`: UpdateCheckStatus
	fmt.Fprintf(os.Stdout, "Response from `UpdateCheckAPI.UpdateCheckAnswerUpdateCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckAnswerUpdateCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCheckAnswer** | [**UpdateCheckAnswer**](UpdateCheckAnswer.md) |  | 

### Return type

[**UpdateCheckStatus**](UpdateCheckStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckGetUpdateCheck

> UpdateCheckStatus UpdateCheckGetUpdateCheck(ctx).Execute()

Get Update Check



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateCheckAPI.UpdateCheckGetUpdateCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateCheckAPI.UpdateCheckGetUpdateCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCheckGetUpdateCheck`: UpdateCheckStatus
	fmt.Fprintf(os.Stdout, "Response from `UpdateCheckAPI.UpdateCheckGetUpdateCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckGetUpdateCheckRequest struct via the builder pattern


### Return type

[**UpdateCheckStatus**](UpdateCheckStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

