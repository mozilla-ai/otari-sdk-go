# \ModelOverridesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelOverridesClearModelOverride**](ModelOverridesAPI.md#ModelOverridesClearModelOverride) | **Delete** /api/v1/model-overrides/{selector} | Clear Model Override
[**ModelOverridesClearOrganizationModelOverride**](ModelOverridesAPI.md#ModelOverridesClearOrganizationModelOverride) | **Delete** /api/v1/organizations/me/model-overrides/{selector} | Clear Organization Model Override
[**ModelOverridesSetModelOverride**](ModelOverridesAPI.md#ModelOverridesSetModelOverride) | **Put** /api/v1/model-overrides/{selector} | Set Model Override
[**ModelOverridesSetOrganizationModelOverride**](ModelOverridesAPI.md#ModelOverridesSetOrganizationModelOverride) | **Put** /api/v1/organizations/me/model-overrides/{selector} | Set Organization Model Override



## ModelOverridesClearModelOverride

> ModelOverridesClearModelOverride(ctx, selector).Execute()

Clear Model Override



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
	selector := "selector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModelOverridesAPI.ModelOverridesClearModelOverride(context.Background(), selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelOverridesAPI.ModelOverridesClearModelOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelOverridesClearModelOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelOverridesClearOrganizationModelOverride

> ModelOverridesClearOrganizationModelOverride(ctx, selector).Execute()

Clear Organization Model Override



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
	selector := "selector_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModelOverridesAPI.ModelOverridesClearOrganizationModelOverride(context.Background(), selector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelOverridesAPI.ModelOverridesClearOrganizationModelOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelOverridesClearOrganizationModelOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelOverridesSetModelOverride

> ModelMetadataOverridePublic ModelOverridesSetModelOverride(ctx, selector).ModelMetadataOverrideRequest(modelMetadataOverrideRequest).Execute()

Set Model Override



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
	selector := "selector_example" // string | 
	modelMetadataOverrideRequest := *openapiclient.NewModelMetadataOverrideRequest() // ModelMetadataOverrideRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelOverridesAPI.ModelOverridesSetModelOverride(context.Background(), selector).ModelMetadataOverrideRequest(modelMetadataOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelOverridesAPI.ModelOverridesSetModelOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelOverridesSetModelOverride`: ModelMetadataOverridePublic
	fmt.Fprintf(os.Stdout, "Response from `ModelOverridesAPI.ModelOverridesSetModelOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelOverridesSetModelOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelMetadataOverrideRequest** | [**ModelMetadataOverrideRequest**](ModelMetadataOverrideRequest.md) |  | 

### Return type

[**ModelMetadataOverridePublic**](ModelMetadataOverridePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelOverridesSetOrganizationModelOverride

> ModelMetadataOverridePublic ModelOverridesSetOrganizationModelOverride(ctx, selector).ModelMetadataOverrideRequest(modelMetadataOverrideRequest).Execute()

Set Organization Model Override



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
	selector := "selector_example" // string | 
	modelMetadataOverrideRequest := *openapiclient.NewModelMetadataOverrideRequest() // ModelMetadataOverrideRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelOverridesAPI.ModelOverridesSetOrganizationModelOverride(context.Background(), selector).ModelMetadataOverrideRequest(modelMetadataOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelOverridesAPI.ModelOverridesSetOrganizationModelOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelOverridesSetOrganizationModelOverride`: ModelMetadataOverridePublic
	fmt.Fprintf(os.Stdout, "Response from `ModelOverridesAPI.ModelOverridesSetOrganizationModelOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selector** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelOverridesSetOrganizationModelOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelMetadataOverrideRequest** | [**ModelMetadataOverrideRequest**](ModelMetadataOverrideRequest.md) |  | 

### Return type

[**ModelMetadataOverridePublic**](ModelMetadataOverridePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

