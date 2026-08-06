# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchForToolV1SearchSearchToolNamePost**](SearchAPI.md#CreateSearchForToolV1SearchSearchToolNamePost) | **Post** /v1/search/{search_tool_name} | Create Search For Tool
[**CreateSearchV1SearchPost**](SearchAPI.md#CreateSearchV1SearchPost) | **Post** /v1/search | Create Search



## CreateSearchForToolV1SearchSearchToolNamePost

> SearchResponse CreateSearchForToolV1SearchSearchToolNamePost(ctx, searchToolName).SearchRequest(searchRequest).Execute()

Create Search For Tool



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
	searchToolName := "searchToolName_example" // string | Configured search tool to run against
	searchRequest := *openapiclient.NewSearchRequest("Query_example") // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CreateSearchForToolV1SearchSearchToolNamePost(context.Background(), searchToolName).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CreateSearchForToolV1SearchSearchToolNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearchForToolV1SearchSearchToolNamePost`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CreateSearchForToolV1SearchSearchToolNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchToolName** | **string** | Configured search tool to run against | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchForToolV1SearchSearchToolNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSearchV1SearchPost

> SearchResponse CreateSearchV1SearchPost(ctx).SearchRequest(searchRequest).Execute()

Create Search



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
	searchRequest := *openapiclient.NewSearchRequest("Query_example") // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.CreateSearchV1SearchPost(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CreateSearchV1SearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearchV1SearchPost`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CreateSearchV1SearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchV1SearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

