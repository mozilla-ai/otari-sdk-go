# \ProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStoredProviderV1ProviderCredentialsPost**](ProvidersAPI.md#CreateStoredProviderV1ProviderCredentialsPost) | **Post** /v1/provider-credentials | Create Stored Provider
[**DeleteStoredProviderV1ProviderCredentialsInstanceDelete**](ProvidersAPI.md#DeleteStoredProviderV1ProviderCredentialsInstanceDelete) | **Delete** /v1/provider-credentials/{instance} | Delete Stored Provider
[**ListProvidersV1ProvidersGet**](ProvidersAPI.md#ListProvidersV1ProvidersGet) | **Get** /v1/providers | List Providers
[**ListStoredProvidersV1ProviderCredentialsGet**](ProvidersAPI.md#ListStoredProvidersV1ProviderCredentialsGet) | **Get** /v1/provider-credentials | List Stored Providers
[**ProviderCatalogDetailV1ProvidersCatalogProviderIdGet**](ProvidersAPI.md#ProviderCatalogDetailV1ProvidersCatalogProviderIdGet) | **Get** /v1/providers/catalog/{provider_id} | Provider Catalog Detail
[**ProviderCatalogV1ProvidersCatalogGet**](ProvidersAPI.md#ProviderCatalogV1ProvidersCatalogGet) | **Get** /v1/providers/catalog | Provider Catalog
[**ProviderHealthV1ProvidersHealthGet**](ProvidersAPI.md#ProviderHealthV1ProvidersHealthGet) | **Get** /v1/providers/health | Provider Health
[**ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost**](ProvidersAPI.md#ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost) | **Post** /v1/provider-credentials/reencrypt | Reencrypt Stored Provider Keys
[**TestProviderConnectionV1ProviderCredentialsTestPost**](ProvidersAPI.md#TestProviderConnectionV1ProviderCredentialsTestPost) | **Post** /v1/provider-credentials/test | Test Provider Connection
[**TestStoredProviderV1ProviderCredentialsInstanceTestPost**](ProvidersAPI.md#TestStoredProviderV1ProviderCredentialsInstanceTestPost) | **Post** /v1/provider-credentials/{instance}/test | Test Stored Provider
[**UpdateStoredProviderV1ProviderCredentialsInstancePatch**](ProvidersAPI.md#UpdateStoredProviderV1ProviderCredentialsInstancePatch) | **Patch** /v1/provider-credentials/{instance} | Update Stored Provider



## CreateStoredProviderV1ProviderCredentialsPost

> StoredProviderResponse CreateStoredProviderV1ProviderCredentialsPost(ctx).CreateStoredProviderRequest(createStoredProviderRequest).Execute()

Create Stored Provider



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
	createStoredProviderRequest := *openapiclient.NewCreateStoredProviderRequest("Instance_example") // CreateStoredProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.CreateStoredProviderV1ProviderCredentialsPost(context.Background()).CreateStoredProviderRequest(createStoredProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.CreateStoredProviderV1ProviderCredentialsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStoredProviderV1ProviderCredentialsPost`: StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.CreateStoredProviderV1ProviderCredentialsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoredProviderV1ProviderCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStoredProviderRequest** | [**CreateStoredProviderRequest**](CreateStoredProviderRequest.md) |  | 

### Return type

[**StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoredProviderV1ProviderCredentialsInstanceDelete

> DeleteStoredProviderV1ProviderCredentialsInstanceDelete(ctx, instance).Execute()

Delete Stored Provider



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
	instance := "instance_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvidersAPI.DeleteStoredProviderV1ProviderCredentialsInstanceDelete(context.Background(), instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.DeleteStoredProviderV1ProviderCredentialsInstanceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoredProviderV1ProviderCredentialsInstanceDeleteRequest struct via the builder pattern


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


## ListProvidersV1ProvidersGet

> ProvidersResponse ListProvidersV1ProvidersGet(ctx).Execute()

List Providers



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
	resp, r, err := apiClient.ProvidersAPI.ListProvidersV1ProvidersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ListProvidersV1ProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProvidersV1ProvidersGet`: ProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ListProvidersV1ProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProvidersV1ProvidersGetRequest struct via the builder pattern


### Return type

[**ProvidersResponse**](ProvidersResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStoredProvidersV1ProviderCredentialsGet

> []StoredProviderResponse ListStoredProvidersV1ProviderCredentialsGet(ctx).Execute()

List Stored Providers



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
	resp, r, err := apiClient.ProvidersAPI.ListStoredProvidersV1ProviderCredentialsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ListStoredProvidersV1ProviderCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStoredProvidersV1ProviderCredentialsGet`: []StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ListStoredProvidersV1ProviderCredentialsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStoredProvidersV1ProviderCredentialsGetRequest struct via the builder pattern


### Return type

[**[]StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderCatalogDetailV1ProvidersCatalogProviderIdGet

> KnownProviderSchema ProviderCatalogDetailV1ProvidersCatalogProviderIdGet(ctx, providerId).Execute()

Provider Catalog Detail



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
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProviderCatalogDetailV1ProvidersCatalogProviderIdGet(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProviderCatalogDetailV1ProvidersCatalogProviderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderCatalogDetailV1ProvidersCatalogProviderIdGet`: KnownProviderSchema
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProviderCatalogDetailV1ProvidersCatalogProviderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderCatalogDetailV1ProvidersCatalogProviderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnownProviderSchema**](KnownProviderSchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderCatalogV1ProvidersCatalogGet

> []KnownProviderSummarySchema ProviderCatalogV1ProvidersCatalogGet(ctx).Execute()

Provider Catalog



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
	resp, r, err := apiClient.ProvidersAPI.ProviderCatalogV1ProvidersCatalogGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProviderCatalogV1ProvidersCatalogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderCatalogV1ProvidersCatalogGet`: []KnownProviderSummarySchema
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProviderCatalogV1ProvidersCatalogGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProviderCatalogV1ProvidersCatalogGetRequest struct via the builder pattern


### Return type

[**[]KnownProviderSummarySchema**](KnownProviderSummarySchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderHealthV1ProvidersHealthGet

> ProviderHealthResponse ProviderHealthV1ProvidersHealthGet(ctx).Refresh(refresh).Execute()

Provider Health



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
	refresh := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProviderHealthV1ProvidersHealthGet(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProviderHealthV1ProvidersHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderHealthV1ProvidersHealthGet`: ProviderHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProviderHealthV1ProvidersHealthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderHealthV1ProvidersHealthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** |  | [default to false]

### Return type

[**ProviderHealthResponse**](ProviderHealthResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost

> ReencryptProviderCredentialsResponse ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost(ctx).Execute()

Reencrypt Stored Provider Keys



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
	resp, r, err := apiClient.ProvidersAPI.ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost`: ReencryptProviderCredentialsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ReencryptStoredProviderKeysV1ProviderCredentialsReencryptPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReencryptStoredProviderKeysV1ProviderCredentialsReencryptPostRequest struct via the builder pattern


### Return type

[**ReencryptProviderCredentialsResponse**](ReencryptProviderCredentialsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestProviderConnectionV1ProviderCredentialsTestPost

> TestProviderResponse TestProviderConnectionV1ProviderCredentialsTestPost(ctx).TestProviderRequest(testProviderRequest).Execute()

Test Provider Connection



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
	testProviderRequest := *openapiclient.NewTestProviderRequest() // TestProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.TestProviderConnectionV1ProviderCredentialsTestPost(context.Background()).TestProviderRequest(testProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.TestProviderConnectionV1ProviderCredentialsTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestProviderConnectionV1ProviderCredentialsTestPost`: TestProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.TestProviderConnectionV1ProviderCredentialsTestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestProviderConnectionV1ProviderCredentialsTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testProviderRequest** | [**TestProviderRequest**](TestProviderRequest.md) |  | 

### Return type

[**TestProviderResponse**](TestProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestStoredProviderV1ProviderCredentialsInstanceTestPost

> TestProviderResponse TestStoredProviderV1ProviderCredentialsInstanceTestPost(ctx, instance).Execute()

Test Stored Provider



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
	instance := "instance_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.TestStoredProviderV1ProviderCredentialsInstanceTestPost(context.Background(), instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.TestStoredProviderV1ProviderCredentialsInstanceTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestStoredProviderV1ProviderCredentialsInstanceTestPost`: TestProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.TestStoredProviderV1ProviderCredentialsInstanceTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestStoredProviderV1ProviderCredentialsInstanceTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestProviderResponse**](TestProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStoredProviderV1ProviderCredentialsInstancePatch

> StoredProviderResponse UpdateStoredProviderV1ProviderCredentialsInstancePatch(ctx, instance).UpdateStoredProviderRequest(updateStoredProviderRequest).Execute()

Update Stored Provider



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
	instance := "instance_example" // string | 
	updateStoredProviderRequest := *openapiclient.NewUpdateStoredProviderRequest() // UpdateStoredProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.UpdateStoredProviderV1ProviderCredentialsInstancePatch(context.Background(), instance).UpdateStoredProviderRequest(updateStoredProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.UpdateStoredProviderV1ProviderCredentialsInstancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStoredProviderV1ProviderCredentialsInstancePatch`: StoredProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.UpdateStoredProviderV1ProviderCredentialsInstancePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoredProviderV1ProviderCredentialsInstancePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoredProviderRequest** | [**UpdateStoredProviderRequest**](UpdateStoredProviderRequest.md) |  | 

### Return type

[**StoredProviderResponse**](StoredProviderResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

