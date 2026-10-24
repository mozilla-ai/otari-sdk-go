# \HealthChecksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthChecksDisableFailingModels**](HealthChecksAPI.md#HealthChecksDisableFailingModels) | **Post** /api/v1/health-checks/disable-failing | Disable Failing Models
[**HealthChecksGetHealthCheckSettings**](HealthChecksAPI.md#HealthChecksGetHealthCheckSettings) | **Get** /api/v1/health-checks/settings | Get Health Check Settings
[**HealthChecksGetHealthCheckStatus**](HealthChecksAPI.md#HealthChecksGetHealthCheckStatus) | **Get** /api/v1/health-checks | Get Health Check Status
[**HealthChecksRunHealthChecks**](HealthChecksAPI.md#HealthChecksRunHealthChecks) | **Post** /api/v1/health-checks/run | Run Health Checks
[**HealthChecksUpdateHealthCheckSettings**](HealthChecksAPI.md#HealthChecksUpdateHealthCheckSettings) | **Patch** /api/v1/health-checks/settings | Update Health Check Settings



## HealthChecksDisableFailingModels

> HealthCheckDisabled HealthChecksDisableFailingModels(ctx).Execute()

Disable Failing Models



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
	resp, r, err := apiClient.HealthChecksAPI.HealthChecksDisableFailingModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthChecksAPI.HealthChecksDisableFailingModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthChecksDisableFailingModels`: HealthCheckDisabled
	fmt.Fprintf(os.Stdout, "Response from `HealthChecksAPI.HealthChecksDisableFailingModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksDisableFailingModelsRequest struct via the builder pattern


### Return type

[**HealthCheckDisabled**](HealthCheckDisabled.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthChecksGetHealthCheckSettings

> HealthCheckSettingsPublic HealthChecksGetHealthCheckSettings(ctx).Execute()

Get Health Check Settings



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
	resp, r, err := apiClient.HealthChecksAPI.HealthChecksGetHealthCheckSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthChecksAPI.HealthChecksGetHealthCheckSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthChecksGetHealthCheckSettings`: HealthCheckSettingsPublic
	fmt.Fprintf(os.Stdout, "Response from `HealthChecksAPI.HealthChecksGetHealthCheckSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksGetHealthCheckSettingsRequest struct via the builder pattern


### Return type

[**HealthCheckSettingsPublic**](HealthCheckSettingsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthChecksGetHealthCheckStatus

> HealthCheckStatus HealthChecksGetHealthCheckStatus(ctx).Execute()

Get Health Check Status



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
	resp, r, err := apiClient.HealthChecksAPI.HealthChecksGetHealthCheckStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthChecksAPI.HealthChecksGetHealthCheckStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthChecksGetHealthCheckStatus`: HealthCheckStatus
	fmt.Fprintf(os.Stdout, "Response from `HealthChecksAPI.HealthChecksGetHealthCheckStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksGetHealthCheckStatusRequest struct via the builder pattern


### Return type

[**HealthCheckStatus**](HealthCheckStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthChecksRunHealthChecks

> HealthCheckSettingsPublic HealthChecksRunHealthChecks(ctx).Execute()

Run Health Checks



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
	resp, r, err := apiClient.HealthChecksAPI.HealthChecksRunHealthChecks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthChecksAPI.HealthChecksRunHealthChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthChecksRunHealthChecks`: HealthCheckSettingsPublic
	fmt.Fprintf(os.Stdout, "Response from `HealthChecksAPI.HealthChecksRunHealthChecks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksRunHealthChecksRequest struct via the builder pattern


### Return type

[**HealthCheckSettingsPublic**](HealthCheckSettingsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthChecksUpdateHealthCheckSettings

> HealthCheckSettingsPublic HealthChecksUpdateHealthCheckSettings(ctx).HealthCheckSettingsUpdate(healthCheckSettingsUpdate).Execute()

Update Health Check Settings



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
	healthCheckSettingsUpdate := *openapiclient.NewHealthCheckSettingsUpdate() // HealthCheckSettingsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthChecksAPI.HealthChecksUpdateHealthCheckSettings(context.Background()).HealthCheckSettingsUpdate(healthCheckSettingsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthChecksAPI.HealthChecksUpdateHealthCheckSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthChecksUpdateHealthCheckSettings`: HealthCheckSettingsPublic
	fmt.Fprintf(os.Stdout, "Response from `HealthChecksAPI.HealthChecksUpdateHealthCheckSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksUpdateHealthCheckSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **healthCheckSettingsUpdate** | [**HealthCheckSettingsUpdate**](HealthCheckSettingsUpdate.md) |  | 

### Return type

[**HealthCheckSettingsPublic**](HealthCheckSettingsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

