# \ScopedBudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScopedBudgetV1ScopedBudgetsPost**](ScopedBudgetsAPI.md#CreateScopedBudgetV1ScopedBudgetsPost) | **Post** /v1/scoped-budgets | Create Scoped Budget
[**DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete**](ScopedBudgetsAPI.md#DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete) | **Delete** /v1/scoped-budgets/{budget_id} | Delete Scoped Budget
[**GetScopedBudgetV1ScopedBudgetsBudgetIdGet**](ScopedBudgetsAPI.md#GetScopedBudgetV1ScopedBudgetsBudgetIdGet) | **Get** /v1/scoped-budgets/{budget_id} | Get Scoped Budget
[**ListScopedBudgetsV1ScopedBudgetsGet**](ScopedBudgetsAPI.md#ListScopedBudgetsV1ScopedBudgetsGet) | **Get** /v1/scoped-budgets | List Scoped Budgets
[**UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch**](ScopedBudgetsAPI.md#UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch) | **Patch** /v1/scoped-budgets/{budget_id} | Update Scoped Budget



## CreateScopedBudgetV1ScopedBudgetsPost

> ScopedBudgetResponse CreateScopedBudgetV1ScopedBudgetsPost(ctx).CreateScopedBudgetRequest(createScopedBudgetRequest).Execute()

Create Scoped Budget



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
	createScopedBudgetRequest := *openapiclient.NewCreateScopedBudgetRequest("BudgetId_example", "ScopeId_example", "ScopeType_example") // CreateScopedBudgetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedBudgetsAPI.CreateScopedBudgetV1ScopedBudgetsPost(context.Background()).CreateScopedBudgetRequest(createScopedBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.CreateScopedBudgetV1ScopedBudgetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScopedBudgetV1ScopedBudgetsPost`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.CreateScopedBudgetV1ScopedBudgetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScopedBudgetV1ScopedBudgetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createScopedBudgetRequest** | [**CreateScopedBudgetRequest**](CreateScopedBudgetRequest.md) |  | 

### Return type

[**ScopedBudgetResponse**](ScopedBudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete

> DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete(ctx, budgetId).Execute()

Delete Scoped Budget



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
	budgetId := "budgetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopedBudgetsAPI.DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.DeleteScopedBudgetV1ScopedBudgetsBudgetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScopedBudgetV1ScopedBudgetsBudgetIdDeleteRequest struct via the builder pattern


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


## GetScopedBudgetV1ScopedBudgetsBudgetIdGet

> ScopedBudgetResponse GetScopedBudgetV1ScopedBudgetsBudgetIdGet(ctx, budgetId).Execute()

Get Scoped Budget



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
	budgetId := "budgetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedBudgetsAPI.GetScopedBudgetV1ScopedBudgetsBudgetIdGet(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.GetScopedBudgetV1ScopedBudgetsBudgetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScopedBudgetV1ScopedBudgetsBudgetIdGet`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.GetScopedBudgetV1ScopedBudgetsBudgetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopedBudgetV1ScopedBudgetsBudgetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScopedBudgetResponse**](ScopedBudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScopedBudgetsV1ScopedBudgetsGet

> []ScopedBudgetResponse ListScopedBudgetsV1ScopedBudgetsGet(ctx).ScopeType(scopeType).ScopeId(scopeId).Skip(skip).Limit(limit).Execute()

List Scoped Budgets



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
	scopeType := "scopeType_example" // string |  (optional)
	scopeId := "scopeId_example" // string |  (optional)
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedBudgetsAPI.ListScopedBudgetsV1ScopedBudgetsGet(context.Background()).ScopeType(scopeType).ScopeId(scopeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ListScopedBudgetsV1ScopedBudgetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListScopedBudgetsV1ScopedBudgetsGet`: []ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.ListScopedBudgetsV1ScopedBudgetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScopedBudgetsV1ScopedBudgetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeType** | **string** |  | 
 **scopeId** | **string** |  | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ScopedBudgetResponse**](ScopedBudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch

> ScopedBudgetResponse UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch(ctx, budgetId).UpdateScopedBudgetRequest(updateScopedBudgetRequest).Execute()

Update Scoped Budget



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
	budgetId := "budgetId_example" // string | 
	updateScopedBudgetRequest := *openapiclient.NewUpdateScopedBudgetRequest() // UpdateScopedBudgetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopedBudgetsAPI.UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch(context.Background(), budgetId).UpdateScopedBudgetRequest(updateScopedBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.UpdateScopedBudgetV1ScopedBudgetsBudgetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScopedBudgetV1ScopedBudgetsBudgetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScopedBudgetRequest** | [**UpdateScopedBudgetRequest**](UpdateScopedBudgetRequest.md) |  | 

### Return type

[**ScopedBudgetResponse**](ScopedBudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

