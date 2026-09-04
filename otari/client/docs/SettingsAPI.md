# \SettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMailSettingsV1SettingsMailGet**](SettingsAPI.md#GetMailSettingsV1SettingsMailGet) | **Get** /v1/settings/mail | Get Mail Settings
[**GetMaintenanceModeV1SettingsMaintenanceModeGet**](SettingsAPI.md#GetMaintenanceModeV1SettingsMaintenanceModeGet) | **Get** /v1/settings/maintenance-mode | Get Maintenance Mode
[**GetSettingsV1SettingsGet**](SettingsAPI.md#GetSettingsV1SettingsGet) | **Get** /v1/settings | Get Settings
[**RotateMasterKeyV1SettingsMasterKeyRotatePost**](SettingsAPI.md#RotateMasterKeyV1SettingsMasterKeyRotatePost) | **Post** /v1/settings/master-key/rotate | Rotate Master Key
[**SendTestMailV1SettingsMailTestPost**](SettingsAPI.md#SendTestMailV1SettingsMailTestPost) | **Post** /v1/settings/mail/test | Send Test Mail
[**UpdateMaintenanceModeV1SettingsMaintenanceModePatch**](SettingsAPI.md#UpdateMaintenanceModeV1SettingsMaintenanceModePatch) | **Patch** /v1/settings/maintenance-mode | Update Maintenance Mode
[**UpdateSettingsV1SettingsPatch**](SettingsAPI.md#UpdateSettingsV1SettingsPatch) | **Patch** /v1/settings | Update Settings



## GetMailSettingsV1SettingsMailGet

> MailSettings GetMailSettingsV1SettingsMailGet(ctx).Execute()

Get Mail Settings



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
	resp, r, err := apiClient.SettingsAPI.GetMailSettingsV1SettingsMailGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetMailSettingsV1SettingsMailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailSettingsV1SettingsMailGet`: MailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetMailSettingsV1SettingsMailGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailSettingsV1SettingsMailGetRequest struct via the builder pattern


### Return type

[**MailSettings**](MailSettings.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceModeV1SettingsMaintenanceModeGet

> MaintenanceMode GetMaintenanceModeV1SettingsMaintenanceModeGet(ctx).Execute()

Get Maintenance Mode



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
	resp, r, err := apiClient.SettingsAPI.GetMaintenanceModeV1SettingsMaintenanceModeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetMaintenanceModeV1SettingsMaintenanceModeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMaintenanceModeV1SettingsMaintenanceModeGet`: MaintenanceMode
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetMaintenanceModeV1SettingsMaintenanceModeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceModeV1SettingsMaintenanceModeGetRequest struct via the builder pattern


### Return type

[**MaintenanceMode**](MaintenanceMode.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## SendTestMailV1SettingsMailTestPost

> SendTestMailResponse SendTestMailV1SettingsMailTestPost(ctx).SendTestMailRequest(sendTestMailRequest).Execute()

Send Test Mail



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
	sendTestMailRequest := *openapiclient.NewSendTestMailRequest("To_example") // SendTestMailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SendTestMailV1SettingsMailTestPost(context.Background()).SendTestMailRequest(sendTestMailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SendTestMailV1SettingsMailTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestMailV1SettingsMailTestPost`: SendTestMailResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SendTestMailV1SettingsMailTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTestMailV1SettingsMailTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendTestMailRequest** | [**SendTestMailRequest**](SendTestMailRequest.md) |  | 

### Return type

[**SendTestMailResponse**](SendTestMailResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMaintenanceModeV1SettingsMaintenanceModePatch

> MaintenanceMode UpdateMaintenanceModeV1SettingsMaintenanceModePatch(ctx).UpdateMaintenanceModeRequest(updateMaintenanceModeRequest).Execute()

Update Maintenance Mode



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
	updateMaintenanceModeRequest := *openapiclient.NewUpdateMaintenanceModeRequest(false) // UpdateMaintenanceModeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UpdateMaintenanceModeV1SettingsMaintenanceModePatch(context.Background()).UpdateMaintenanceModeRequest(updateMaintenanceModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateMaintenanceModeV1SettingsMaintenanceModePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMaintenanceModeV1SettingsMaintenanceModePatch`: MaintenanceMode
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateMaintenanceModeV1SettingsMaintenanceModePatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMaintenanceModeV1SettingsMaintenanceModePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMaintenanceModeRequest** | [**UpdateMaintenanceModeRequest**](UpdateMaintenanceModeRequest.md) |  | 

### Return type

[**MaintenanceMode**](MaintenanceMode.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

