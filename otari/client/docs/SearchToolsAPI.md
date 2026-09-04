# \SearchToolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchToolV1SearchToolsPost**](SearchToolsAPI.md#CreateSearchToolV1SearchToolsPost) | **Post** /v1/search-tools | Create Search Tool
[**DeleteStoredSearchToolV1SearchToolsNameDelete**](SearchToolsAPI.md#DeleteStoredSearchToolV1SearchToolsNameDelete) | **Delete** /v1/search-tools/{name} | Delete Stored Search Tool
[**ListAllSearchToolsV1SearchToolsGet**](SearchToolsAPI.md#ListAllSearchToolsV1SearchToolsGet) | **Get** /v1/search-tools | List All Search Tools
[**ListSearchProvidersV1SearchToolsProvidersGet**](SearchToolsAPI.md#ListSearchProvidersV1SearchToolsProvidersGet) | **Get** /v1/search-tools/providers | List Search Providers
[**ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost**](SearchToolsAPI.md#ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost) | **Post** /v1/search-tools/reencrypt | Reencrypt Stored Search Tool Keys
[**UpdateSearchToolV1SearchToolsNamePatch**](SearchToolsAPI.md#UpdateSearchToolV1SearchToolsNamePatch) | **Patch** /v1/search-tools/{name} | Update Search Tool



## CreateSearchToolV1SearchToolsPost

> StoredSearchToolSchema CreateSearchToolV1SearchToolsPost(ctx).CreateSearchToolRequest(createSearchToolRequest).Execute()

Create Search Tool



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
	createSearchToolRequest := *openapiclient.NewCreateSearchToolRequest("Name_example", "Provider_example") // CreateSearchToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchToolsAPI.CreateSearchToolV1SearchToolsPost(context.Background()).CreateSearchToolRequest(createSearchToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.CreateSearchToolV1SearchToolsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearchToolV1SearchToolsPost`: StoredSearchToolSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.CreateSearchToolV1SearchToolsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchToolV1SearchToolsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSearchToolRequest** | [**CreateSearchToolRequest**](CreateSearchToolRequest.md) |  | 

### Return type

[**StoredSearchToolSchema**](StoredSearchToolSchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoredSearchToolV1SearchToolsNameDelete

> DeleteStoredSearchToolV1SearchToolsNameDelete(ctx, name).Execute()

Delete Stored Search Tool



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchToolsAPI.DeleteStoredSearchToolV1SearchToolsNameDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.DeleteStoredSearchToolV1SearchToolsNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoredSearchToolV1SearchToolsNameDeleteRequest struct via the builder pattern


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


## ListAllSearchToolsV1SearchToolsGet

> SearchToolsResponse ListAllSearchToolsV1SearchToolsGet(ctx).Execute()

List All Search Tools



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
	resp, r, err := apiClient.SearchToolsAPI.ListAllSearchToolsV1SearchToolsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.ListAllSearchToolsV1SearchToolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllSearchToolsV1SearchToolsGet`: SearchToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.ListAllSearchToolsV1SearchToolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllSearchToolsV1SearchToolsGetRequest struct via the builder pattern


### Return type

[**SearchToolsResponse**](SearchToolsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSearchProvidersV1SearchToolsProvidersGet

> []SearchProviderSchema ListSearchProvidersV1SearchToolsProvidersGet(ctx).Execute()

List Search Providers



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
	resp, r, err := apiClient.SearchToolsAPI.ListSearchProvidersV1SearchToolsProvidersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.ListSearchProvidersV1SearchToolsProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSearchProvidersV1SearchToolsProvidersGet`: []SearchProviderSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.ListSearchProvidersV1SearchToolsProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSearchProvidersV1SearchToolsProvidersGetRequest struct via the builder pattern


### Return type

[**[]SearchProviderSchema**](SearchProviderSchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost

> ReencryptSearchToolsResponse ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost(ctx).Execute()

Reencrypt Stored Search Tool Keys



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
	resp, r, err := apiClient.SearchToolsAPI.ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost`: ReencryptSearchToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.ReencryptStoredSearchToolKeysV1SearchToolsReencryptPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReencryptStoredSearchToolKeysV1SearchToolsReencryptPostRequest struct via the builder pattern


### Return type

[**ReencryptSearchToolsResponse**](ReencryptSearchToolsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSearchToolV1SearchToolsNamePatch

> StoredSearchToolSchema UpdateSearchToolV1SearchToolsNamePatch(ctx, name).UpdateSearchToolRequest(updateSearchToolRequest).Execute()

Update Search Tool



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
	name := "name_example" // string | 
	updateSearchToolRequest := *openapiclient.NewUpdateSearchToolRequest() // UpdateSearchToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchToolsAPI.UpdateSearchToolV1SearchToolsNamePatch(context.Background(), name).UpdateSearchToolRequest(updateSearchToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.UpdateSearchToolV1SearchToolsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSearchToolV1SearchToolsNamePatch`: StoredSearchToolSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.UpdateSearchToolV1SearchToolsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSearchToolV1SearchToolsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSearchToolRequest** | [**UpdateSearchToolRequest**](UpdateSearchToolRequest.md) |  | 

### Return type

[**StoredSearchToolSchema**](StoredSearchToolSchema.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

