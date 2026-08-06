# \SettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsV1SettingsGet**](SettingsAPI.md#GetSettingsV1SettingsGet) | **Get** /v1/settings | Get Settings
[**RotateMasterKeyV1SettingsMasterKeyRotatePost**](SettingsAPI.md#RotateMasterKeyV1SettingsMasterKeyRotatePost) | **Post** /v1/settings/master-key/rotate | Rotate Master Key
[**UpdateSettingsV1SettingsPatch**](SettingsAPI.md#UpdateSettingsV1SettingsPatch) | **Patch** /v1/settings | Update Settings



## GetSettingsV1SettingsGet

> GatewaySettings GetSettingsV1SettingsGet(ctx).Execute()

Get Settings



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
	resp, r, err := apiClient.SettingsAPI.GetSettingsV1SettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettingsV1SettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsV1SettingsGet`: GatewaySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSettingsV1SettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsV1SettingsGetRequest struct via the builder pattern


### Return type

[**GatewaySettings**](GatewaySettings.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateMasterKeyV1SettingsMasterKeyRotatePost

> RotateMasterKeyResponse RotateMasterKeyV1SettingsMasterKeyRotatePost(ctx).Execute()

Rotate Master Key



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
	resp, r, err := apiClient.SettingsAPI.RotateMasterKeyV1SettingsMasterKeyRotatePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.RotateMasterKeyV1SettingsMasterKeyRotatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateMasterKeyV1SettingsMasterKeyRotatePost`: RotateMasterKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.RotateMasterKeyV1SettingsMasterKeyRotatePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateMasterKeyV1SettingsMasterKeyRotatePostRequest struct via the builder pattern


### Return type

[**RotateMasterKeyResponse**](RotateMasterKeyResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettingsV1SettingsPatch

> GatewaySettings UpdateSettingsV1SettingsPatch(ctx).UpdateSettingsRequest(updateSettingsRequest).Execute()

Update Settings



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
	updateSettingsRequest := *openapiclient.NewUpdateSettingsRequest() // UpdateSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UpdateSettingsV1SettingsPatch(context.Background()).UpdateSettingsRequest(updateSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateSettingsV1SettingsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettingsV1SettingsPatch`: GatewaySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateSettingsV1SettingsPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsV1SettingsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSettingsRequest** | [**UpdateSettingsRequest**](UpdateSettingsRequest.md) |  | 

### Return type

[**GatewaySettings**](GatewaySettings.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

