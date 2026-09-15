# \McpAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**McpExecuteMcpTool**](McpAPI.md#McpExecuteMcpTool) | **Post** /api/v1/mcp/execute | Execute Mcp Tool
[**McpListMcpTools**](McpAPI.md#McpListMcpTools) | **Get** /api/v1/mcp/servers/{mcp_server_id}/tools | List Mcp Tools



## McpExecuteMcpTool

> CallToolResult McpExecuteMcpTool(ctx).McpExecuteRequest(mcpExecuteRequest).Execute()

Execute Mcp Tool



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
	mcpExecuteRequest := *openapiclient.NewMcpExecuteRequest("ClientExecutionId_example", "McpServerId_example", "ServerRevision_example", "ToolName_example") // McpExecuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpAPI.McpExecuteMcpTool(context.Background()).McpExecuteRequest(mcpExecuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpAPI.McpExecuteMcpTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpExecuteMcpTool`: CallToolResult
	fmt.Fprintf(os.Stdout, "Response from `McpAPI.McpExecuteMcpTool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcpExecuteMcpToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcpExecuteRequest** | [**McpExecuteRequest**](McpExecuteRequest.md) |  | 

### Return type

[**CallToolResult**](CallToolResult.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McpListMcpTools

> McpToolsResponse McpListMcpTools(ctx, mcpServerId).Execute()

List Mcp Tools



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
	mcpServerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.McpAPI.McpListMcpTools(context.Background(), mcpServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `McpAPI.McpListMcpTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `McpListMcpTools`: McpToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `McpAPI.McpListMcpTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mcpServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcpListMcpToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**McpToolsResponse**](McpToolsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

