# \WebSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebSearchV1WebSearchSearchGet**](WebSearchAPI.md#WebSearchV1WebSearchSearchGet) | **Get** /v1/web-search/search | Web Search



## WebSearchV1WebSearchSearchGet

> WebSearchBackendResponse WebSearchV1WebSearchSearchGet(ctx).Q(q).MaxResults(maxResults).SearchDepth(searchDepth).Topic(topic).TimeRange(timeRange).IncludeAnswer(includeAnswer).XGatewayToken(xGatewayToken).Execute()

Web Search



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
	q := "q_example" // string | The search query.
	maxResults := int32(56) // int32 |  (optional)
	searchDepth := "searchDepth_example" // string |  (optional)
	topic := "topic_example" // string |  (optional)
	timeRange := "timeRange_example" // string |  (optional)
	includeAnswer := true // bool |  (optional)
	xGatewayToken := "xGatewayToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSearchAPI.WebSearchV1WebSearchSearchGet(context.Background()).Q(q).MaxResults(maxResults).SearchDepth(searchDepth).Topic(topic).TimeRange(timeRange).IncludeAnswer(includeAnswer).XGatewayToken(xGatewayToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchAPI.WebSearchV1WebSearchSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchV1WebSearchSearchGet`: WebSearchBackendResponse
	fmt.Fprintf(os.Stdout, "Response from `WebSearchAPI.WebSearchV1WebSearchSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchV1WebSearchSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The search query. | 
 **maxResults** | **int32** |  | 
 **searchDepth** | **string** |  | 
 **topic** | **string** |  | 
 **timeRange** | **string** |  | 
 **includeAnswer** | **bool** |  | 
 **xGatewayToken** | **string** |  | 

### Return type

[**WebSearchBackendResponse**](WebSearchBackendResponse.md)

### Authorization

[GatewayTokenAuth](../README.md#GatewayTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

