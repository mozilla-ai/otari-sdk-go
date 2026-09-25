# \OverviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OverviewGetOverview**](OverviewAPI.md#OverviewGetOverview) | **Get** /api/v1/overview | Get Overview



## OverviewGetOverview

> OverviewSummaryResponse OverviewGetOverview(ctx).WorkspaceId(workspaceId).Execute()

Get Overview



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Narrow the counts to one workspace of the caller's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverviewAPI.OverviewGetOverview(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.OverviewGetOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverviewGetOverview`: OverviewSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.OverviewGetOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOverviewGetOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Narrow the counts to one workspace of the caller&#39;s organization. | 

### Return type

[**OverviewSummaryResponse**](OverviewSummaryResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

