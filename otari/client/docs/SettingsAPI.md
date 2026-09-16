# \SettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsGetMailSettings**](SettingsAPI.md#SettingsGetMailSettings) | **Get** /api/v1/settings/mail | Get Mail Settings
[**SettingsGetMaintenanceMode**](SettingsAPI.md#SettingsGetMaintenanceMode) | **Get** /api/v1/settings/maintenance-mode | Get Maintenance Mode
[**SettingsGetSettings**](SettingsAPI.md#SettingsGetSettings) | **Get** /api/v1/settings | Get Settings
[**SettingsRotateMasterKey**](SettingsAPI.md#SettingsRotateMasterKey) | **Post** /api/v1/settings/master-key/rotate | Rotate Master Key
[**SettingsSendTestMail**](SettingsAPI.md#SettingsSendTestMail) | **Post** /api/v1/settings/mail/test | Send Test Mail
[**SettingsUpdateMaintenanceMode**](SettingsAPI.md#SettingsUpdateMaintenanceMode) | **Patch** /api/v1/settings/maintenance-mode | Update Maintenance Mode
[**SettingsUpdateSettings**](SettingsAPI.md#SettingsUpdateSettings) | **Patch** /api/v1/settings | Update Settings



## SettingsGetMailSettings

> MailSettings SettingsGetMailSettings(ctx).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsGetMailSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsGetMailSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsGetMailSettings`: MailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsGetMailSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsGetMailSettingsRequest struct via the builder pattern


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


## SettingsGetMaintenanceMode

> MaintenanceMode SettingsGetMaintenanceMode(ctx).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsGetMaintenanceMode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsGetMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsGetMaintenanceMode`: MaintenanceMode
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsGetMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsGetMaintenanceModeRequest struct via the builder pattern


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


## SettingsGetSettings

> GatewaySettings SettingsGetSettings(ctx).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsGetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsGetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsGetSettings`: GatewaySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsGetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsGetSettingsRequest struct via the builder pattern


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


## SettingsRotateMasterKey

> RotateMasterKeyResponse SettingsRotateMasterKey(ctx).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsRotateMasterKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRotateMasterKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRotateMasterKey`: RotateMasterKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRotateMasterKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRotateMasterKeyRequest struct via the builder pattern


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


## SettingsSendTestMail

> SendTestMailResponse SettingsSendTestMail(ctx).SendTestMailRequest(sendTestMailRequest).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsSendTestMail(context.Background()).SendTestMailRequest(sendTestMailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSendTestMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSendTestMail`: SendTestMailResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSendTestMail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSendTestMailRequest struct via the builder pattern


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


## SettingsUpdateMaintenanceMode

> MaintenanceMode SettingsUpdateMaintenanceMode(ctx).UpdateMaintenanceModeRequest(updateMaintenanceModeRequest).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsUpdateMaintenanceMode(context.Background()).UpdateMaintenanceModeRequest(updateMaintenanceModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsUpdateMaintenanceMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsUpdateMaintenanceMode`: MaintenanceMode
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsUpdateMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsUpdateMaintenanceModeRequest struct via the builder pattern


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


## SettingsUpdateSettings

> GatewaySettings SettingsUpdateSettings(ctx).UpdateSettingsRequest(updateSettingsRequest).Execute()

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
	resp, r, err := apiClient.SettingsAPI.SettingsUpdateSettings(context.Background()).UpdateSettingsRequest(updateSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsUpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsUpdateSettings`: GatewaySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsUpdateSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsUpdateSettingsRequest struct via the builder pattern


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

