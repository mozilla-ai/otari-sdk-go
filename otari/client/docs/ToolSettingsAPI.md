# \ToolSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToolSettingsV1ToolSettingsGet**](ToolSettingsAPI.md#GetToolSettingsV1ToolSettingsGet) | **Get** /v1/tool-settings | Get Tool Settings
[**TestServiceV1ToolSettingsServiceTestPost**](ToolSettingsAPI.md#TestServiceV1ToolSettingsServiceTestPost) | **Post** /v1/tool-settings/{service}/test | Test Service
[**UpdateToolSettingsV1ToolSettingsPatch**](ToolSettingsAPI.md#UpdateToolSettingsV1ToolSettingsPatch) | **Patch** /v1/tool-settings | Update Tool Settings



## GetToolSettingsV1ToolSettingsGet

> ToolSettingsResponse GetToolSettingsV1ToolSettingsGet(ctx).Execute()

Get Tool Settings



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
	resp, r, err := apiClient.ToolSettingsAPI.GetToolSettingsV1ToolSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.GetToolSettingsV1ToolSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolSettingsV1ToolSettingsGet`: ToolSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.GetToolSettingsV1ToolSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolSettingsV1ToolSettingsGetRequest struct via the builder pattern


### Return type

[**ToolSettingsResponse**](ToolSettingsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceV1ToolSettingsServiceTestPost

> TestServiceResponse TestServiceV1ToolSettingsServiceTestPost(ctx, service).TestServiceRequest(testServiceRequest).Execute()

Test Service



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
	service := "service_example" // string | 
	testServiceRequest := *openapiclient.NewTestServiceRequest("Url_example") // TestServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolSettingsAPI.TestServiceV1ToolSettingsServiceTestPost(context.Background(), service).TestServiceRequest(testServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.TestServiceV1ToolSettingsServiceTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceV1ToolSettingsServiceTestPost`: TestServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.TestServiceV1ToolSettingsServiceTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceV1ToolSettingsServiceTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testServiceRequest** | [**TestServiceRequest**](TestServiceRequest.md) |  | 

### Return type

[**TestServiceResponse**](TestServiceResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToolSettingsV1ToolSettingsPatch

> ToolSettingsResponse UpdateToolSettingsV1ToolSettingsPatch(ctx).UpdateToolSettingsRequest(updateToolSettingsRequest).Execute()

Update Tool Settings



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
	updateToolSettingsRequest := *openapiclient.NewUpdateToolSettingsRequest() // UpdateToolSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolSettingsAPI.UpdateToolSettingsV1ToolSettingsPatch(context.Background()).UpdateToolSettingsRequest(updateToolSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.UpdateToolSettingsV1ToolSettingsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolSettingsV1ToolSettingsPatch`: ToolSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.UpdateToolSettingsV1ToolSettingsPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolSettingsV1ToolSettingsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateToolSettingsRequest** | [**UpdateToolSettingsRequest**](UpdateToolSettingsRequest.md) |  | 

### Return type

[**ToolSettingsResponse**](ToolSettingsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

