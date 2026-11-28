# \McpCatalogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpCatalogGetMcpCatalogSync**](McpCatalogAPI.md#McpCatalogGetMcpCatalogSync) | **Get** /api/v1/mcp-catalog/sync | Get Mcp Catalog Sync
[**McpCatalogListPopularMcpServers**](McpCatalogAPI.md#McpCatalogListPopularMcpServers) | **Get** /api/v1/mcp-catalog/popular | List Popular Mcp Servers
[**McpCatalogRequestMcpCatalogSync**](McpCatalogAPI.md#McpCatalogRequestMcpCatalogSync) | **Post** /api/v1/mcp-catalog/sync | Request Mcp Catalog Sync
[**McpCatalogSearchMcpCatalog**](McpCatalogAPI.md#McpCatalogSearchMcpCatalog) | **Get** /api/v1/mcp-catalog/servers | Search Mcp Catalog



## McpCatalogGetMcpCatalogSync

> McpCatalogSyncStatus McpCatalogGetMcpCatalogSync(ctx).Execute()

Get Mcp Catalog Sync



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
	resp, r, err := apiClient.McpCatalogAPI.McpCatalogGetMcpCatalogSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpCatalogAPI.McpCatalogGetMcpCatalogSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpCatalogGetMcpCatalogSync`: McpCatalogSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `McpCatalogAPI.McpCatalogGetMcpCatalogSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMcpCatalogGetMcpCatalogSyncRequest struct via the builder pattern


### Return type

[**McpCatalogSyncStatus**](McpCatalogSyncStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpCatalogListPopularMcpServers

> McpCatalogServersPublic McpCatalogListPopularMcpServers(ctx).WorkspaceId(workspaceId).Exclude(exclude).Execute()

List Popular Mcp Servers



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A workspace the caller can see. The servers it already has are left out. (optional)
	exclude := []*string{"Inner_example"} // []*string | Catalog names to leave out, such as the servers a caller has already picked. At most 50. The count leaves them out too, so it pages the same rows the caller sees. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpCatalogAPI.McpCatalogListPopularMcpServers(context.Background()).WorkspaceId(workspaceId).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpCatalogAPI.McpCatalogListPopularMcpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpCatalogListPopularMcpServers`: McpCatalogServersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpCatalogAPI.McpCatalogListPopularMcpServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpCatalogListPopularMcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | A workspace the caller can see. The servers it already has are left out. | 
 **exclude** | **[]string** | Catalog names to leave out, such as the servers a caller has already picked. At most 50. The count leaves them out too, so it pages the same rows the caller sees. | 

### Return type

[**McpCatalogServersPublic**](McpCatalogServersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpCatalogRequestMcpCatalogSync

> McpCatalogSyncStatus McpCatalogRequestMcpCatalogSync(ctx).McpCatalogSyncRequest(mcpCatalogSyncRequest).Execute()

Request Mcp Catalog Sync



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
	mcpCatalogSyncRequest := *openapiclient.NewMcpCatalogSyncRequest() // McpCatalogSyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpCatalogAPI.McpCatalogRequestMcpCatalogSync(context.Background()).McpCatalogSyncRequest(mcpCatalogSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpCatalogAPI.McpCatalogRequestMcpCatalogSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpCatalogRequestMcpCatalogSync`: McpCatalogSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `McpCatalogAPI.McpCatalogRequestMcpCatalogSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpCatalogRequestMcpCatalogSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcpCatalogSyncRequest** | [**McpCatalogSyncRequest**](McpCatalogSyncRequest.md) |  | 

### Return type

[**McpCatalogSyncStatus**](McpCatalogSyncStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpCatalogSearchMcpCatalog

> McpCatalogServersPublic McpCatalogSearchMcpCatalog(ctx).Q(q).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Exclude(exclude).Execute()

Search Mcp Catalog



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
	q := "q_example" // string | Text to find in a server's title, name or description (optional) (default to "")
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 25)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A workspace the caller can see. The servers it already has are left out. (optional)
	exclude := []string{"Inner_example"} // []string | Catalog names to leave out, such as the servers a caller has already picked. At most 50. The count leaves them out too, so it pages the same rows the caller sees. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpCatalogAPI.McpCatalogSearchMcpCatalog(context.Background()).Q(q).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpCatalogAPI.McpCatalogSearchMcpCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpCatalogSearchMcpCatalog`: McpCatalogServersPublic
	fmt.Fprintf(os.Stdout, "Response from `McpCatalogAPI.McpCatalogSearchMcpCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpCatalogSearchMcpCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Text to find in a server&#39;s title, name or description | [default to &quot;&quot;]
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 25]
 **workspaceId** | **string** | A workspace the caller can see. The servers it already has are left out. | 
 **exclude** | **[]string** | Catalog names to leave out, such as the servers a caller has already picked. At most 50. The count leaves them out too, so it pages the same rows the caller sees. | 

### Return type

[**McpCatalogServersPublic**](McpCatalogServersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

