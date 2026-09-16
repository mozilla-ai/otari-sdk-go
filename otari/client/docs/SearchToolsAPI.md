# \SearchToolsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchToolsCreateSearchTool**](SearchToolsAPI.md#SearchToolsCreateSearchTool) | **Post** /api/v1/search-tools | Create Search Tool
[**SearchToolsDeleteStoredSearchTool**](SearchToolsAPI.md#SearchToolsDeleteStoredSearchTool) | **Delete** /api/v1/search-tools/{name} | Delete Stored Search Tool
[**SearchToolsListAllSearchTools**](SearchToolsAPI.md#SearchToolsListAllSearchTools) | **Get** /api/v1/search-tools | List All Search Tools
[**SearchToolsListSearchProviders**](SearchToolsAPI.md#SearchToolsListSearchProviders) | **Get** /api/v1/search-tools/providers | List Search Providers
[**SearchToolsReencryptStoredSearchToolKeys**](SearchToolsAPI.md#SearchToolsReencryptStoredSearchToolKeys) | **Post** /api/v1/search-tools/reencrypt | Reencrypt Stored Search Tool Keys
[**SearchToolsUpdateSearchTool**](SearchToolsAPI.md#SearchToolsUpdateSearchTool) | **Patch** /api/v1/search-tools/{name} | Update Search Tool



## SearchToolsCreateSearchTool

> StoredSearchToolSchema SearchToolsCreateSearchTool(ctx).CreateSearchToolRequest(createSearchToolRequest).Execute()

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
	resp, r, err := apiClient.SearchToolsAPI.SearchToolsCreateSearchTool(context.Background()).CreateSearchToolRequest(createSearchToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsCreateSearchTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchToolsCreateSearchTool`: StoredSearchToolSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.SearchToolsCreateSearchTool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchToolsCreateSearchToolRequest struct via the builder pattern


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


## SearchToolsDeleteStoredSearchTool

> SearchToolsDeleteStoredSearchTool(ctx, name).Execute()

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
	r, err := apiClient.SearchToolsAPI.SearchToolsDeleteStoredSearchTool(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsDeleteStoredSearchTool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSearchToolsDeleteStoredSearchToolRequest struct via the builder pattern


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


## SearchToolsListAllSearchTools

> SearchToolsResponse SearchToolsListAllSearchTools(ctx).Execute()

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
	resp, r, err := apiClient.SearchToolsAPI.SearchToolsListAllSearchTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsListAllSearchTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchToolsListAllSearchTools`: SearchToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.SearchToolsListAllSearchTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchToolsListAllSearchToolsRequest struct via the builder pattern


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


## SearchToolsListSearchProviders

> []SearchProviderSchema SearchToolsListSearchProviders(ctx).Execute()

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
	resp, r, err := apiClient.SearchToolsAPI.SearchToolsListSearchProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsListSearchProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchToolsListSearchProviders`: []SearchProviderSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.SearchToolsListSearchProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchToolsListSearchProvidersRequest struct via the builder pattern


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


## SearchToolsReencryptStoredSearchToolKeys

> ReencryptSearchToolsResponse SearchToolsReencryptStoredSearchToolKeys(ctx).Execute()

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
	resp, r, err := apiClient.SearchToolsAPI.SearchToolsReencryptStoredSearchToolKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsReencryptStoredSearchToolKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchToolsReencryptStoredSearchToolKeys`: ReencryptSearchToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.SearchToolsReencryptStoredSearchToolKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchToolsReencryptStoredSearchToolKeysRequest struct via the builder pattern


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


## SearchToolsUpdateSearchTool

> StoredSearchToolSchema SearchToolsUpdateSearchTool(ctx, name).UpdateSearchToolRequest(updateSearchToolRequest).Execute()

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
	resp, r, err := apiClient.SearchToolsAPI.SearchToolsUpdateSearchTool(context.Background(), name).UpdateSearchToolRequest(updateSearchToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchToolsAPI.SearchToolsUpdateSearchTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchToolsUpdateSearchTool`: StoredSearchToolSchema
	fmt.Fprintf(os.Stdout, "Response from `SearchToolsAPI.SearchToolsUpdateSearchTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchToolsUpdateSearchToolRequest struct via the builder pattern


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

