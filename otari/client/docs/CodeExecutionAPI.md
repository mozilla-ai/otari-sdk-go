# \CodeExecutionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CodeExecutionExecuteCode**](CodeExecutionAPI.md#CodeExecutionExecuteCode) | **Post** /api/v1/code/execute | Execute Code
[**CodeExecutionListCodeExecutionEnvironments**](CodeExecutionAPI.md#CodeExecutionListCodeExecutionEnvironments) | **Get** /api/v1/code/environments | List Code Execution Environments



## CodeExecutionExecuteCode

> CodeExecuteResponse CodeExecutionExecuteCode(ctx).CodeExecuteRequest(codeExecuteRequest).Execute()

Execute Code



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
	codeExecuteRequest := *openapiclient.NewCodeExecuteRequest("Code_example") // CodeExecuteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeExecutionAPI.CodeExecutionExecuteCode(context.Background()).CodeExecuteRequest(codeExecuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeExecutionAPI.CodeExecutionExecuteCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CodeExecutionExecuteCode`: CodeExecuteResponse
	fmt.Fprintf(os.Stdout, "Response from `CodeExecutionAPI.CodeExecutionExecuteCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodeExecutionExecuteCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeExecuteRequest** | [**CodeExecuteRequest**](CodeExecuteRequest.md) |  | 

### Return type

[**CodeExecuteResponse**](CodeExecuteResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CodeExecutionListCodeExecutionEnvironments

> CodeExecutionEnvironmentList CodeExecutionListCodeExecutionEnvironments(ctx).Skip(skip).Limit(limit).Execute()

List Code Execution Environments



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
	skip := int32(56) // int32 | Number of environments to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of environments to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodeExecutionAPI.CodeExecutionListCodeExecutionEnvironments(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeExecutionAPI.CodeExecutionListCodeExecutionEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CodeExecutionListCodeExecutionEnvironments`: CodeExecutionEnvironmentList
	fmt.Fprintf(os.Stdout, "Response from `CodeExecutionAPI.CodeExecutionListCodeExecutionEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCodeExecutionListCodeExecutionEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of environments to skip | [default to 0]
 **limit** | **int32** | Maximum number of environments to return | [default to 100]

### Return type

[**CodeExecutionEnvironmentList**](CodeExecutionEnvironmentList.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

