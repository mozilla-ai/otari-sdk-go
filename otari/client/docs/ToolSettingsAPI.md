# \ToolSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToolSettingsGetToolSettings**](ToolSettingsAPI.md#ToolSettingsGetToolSettings) | **Get** /api/v1/tool-settings | Get Tool Settings
[**ToolSettingsListBuiltinGuardrails**](ToolSettingsAPI.md#ToolSettingsListBuiltinGuardrails) | **Get** /api/v1/tool-settings/guardrails/catalog | List Builtin Guardrails
[**ToolSettingsListGuardrailProfiles**](ToolSettingsAPI.md#ToolSettingsListGuardrailProfiles) | **Get** /api/v1/tool-settings/guardrails/profiles | List Guardrail Profiles
[**ToolSettingsTestService**](ToolSettingsAPI.md#ToolSettingsTestService) | **Post** /api/v1/tool-settings/{service}/test | Test Service
[**ToolSettingsUpdateToolSettings**](ToolSettingsAPI.md#ToolSettingsUpdateToolSettings) | **Patch** /api/v1/tool-settings | Update Tool Settings



## ToolSettingsGetToolSettings

> ToolSettingsResponse ToolSettingsGetToolSettings(ctx).Execute()

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
	resp, r, err := apiClient.ToolSettingsAPI.ToolSettingsGetToolSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.ToolSettingsGetToolSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolSettingsGetToolSettings`: ToolSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.ToolSettingsGetToolSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiToolSettingsGetToolSettingsRequest struct via the builder pattern


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


## ToolSettingsListBuiltinGuardrails

> BuiltInGuardrailCatalog ToolSettingsListBuiltinGuardrails(ctx).Execute()

List Builtin Guardrails



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
	resp, r, err := apiClient.ToolSettingsAPI.ToolSettingsListBuiltinGuardrails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.ToolSettingsListBuiltinGuardrails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolSettingsListBuiltinGuardrails`: BuiltInGuardrailCatalog
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.ToolSettingsListBuiltinGuardrails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiToolSettingsListBuiltinGuardrailsRequest struct via the builder pattern


### Return type

[**BuiltInGuardrailCatalog**](BuiltInGuardrailCatalog.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolSettingsListGuardrailProfiles

> GuardrailCatalog ToolSettingsListGuardrailProfiles(ctx).Execute()

List Guardrail Profiles



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
	resp, r, err := apiClient.ToolSettingsAPI.ToolSettingsListGuardrailProfiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.ToolSettingsListGuardrailProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolSettingsListGuardrailProfiles`: GuardrailCatalog
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.ToolSettingsListGuardrailProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiToolSettingsListGuardrailProfilesRequest struct via the builder pattern


### Return type

[**GuardrailCatalog**](GuardrailCatalog.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolSettingsTestService

> TestServiceResponse ToolSettingsTestService(ctx, service).TestServiceRequest(testServiceRequest).Execute()

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
	resp, r, err := apiClient.ToolSettingsAPI.ToolSettingsTestService(context.Background(), service).TestServiceRequest(testServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.ToolSettingsTestService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolSettingsTestService`: TestServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.ToolSettingsTestService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolSettingsTestServiceRequest struct via the builder pattern


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


## ToolSettingsUpdateToolSettings

> ToolSettingsResponse ToolSettingsUpdateToolSettings(ctx).UpdateToolSettingsRequest(updateToolSettingsRequest).Execute()

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
	resp, r, err := apiClient.ToolSettingsAPI.ToolSettingsUpdateToolSettings(context.Background()).UpdateToolSettingsRequest(updateToolSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolSettingsAPI.ToolSettingsUpdateToolSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolSettingsUpdateToolSettings`: ToolSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolSettingsAPI.ToolSettingsUpdateToolSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolSettingsUpdateToolSettingsRequest struct via the builder pattern


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

