# \FetchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCreateFetch**](FetchAPI.md#FetchCreateFetch) | **Post** /api/v1/fetch | Create Fetch
[**FetchCreateFetchForTool**](FetchAPI.md#FetchCreateFetchForTool) | **Post** /api/v1/fetch/{fetch_tool_name} | Create Fetch For Tool



## FetchCreateFetch

> FetchResponse FetchCreateFetch(ctx).FetchRequest(fetchRequest).Execute()

Create Fetch



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
	fetchRequest := *openapiclient.NewFetchRequest("Url_example") // FetchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FetchAPI.FetchCreateFetch(context.Background()).FetchRequest(fetchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FetchAPI.FetchCreateFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCreateFetch`: FetchResponse
	fmt.Fprintf(os.Stdout, "Response from `FetchAPI.FetchCreateFetch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchCreateFetchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchRequest** | [**FetchRequest**](FetchRequest.md) |  | 

### Return type

[**FetchResponse**](FetchResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCreateFetchForTool

> FetchResponse FetchCreateFetchForTool(ctx, fetchToolName).FetchRequest(fetchRequest).Execute()

Create Fetch For Tool



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
	fetchToolName := "fetchToolName_example" // string | Configured fetch tool to run on, builtin_fetch included
	fetchRequest := *openapiclient.NewFetchRequest("Url_example") // FetchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FetchAPI.FetchCreateFetchForTool(context.Background(), fetchToolName).FetchRequest(fetchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FetchAPI.FetchCreateFetchForTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCreateFetchForTool`: FetchResponse
	fmt.Fprintf(os.Stdout, "Response from `FetchAPI.FetchCreateFetchForTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fetchToolName** | **string** | Configured fetch tool to run on, builtin_fetch included | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCreateFetchForToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchRequest** | [**FetchRequest**](FetchRequest.md) |  | 

### Return type

[**FetchResponse**](FetchResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

