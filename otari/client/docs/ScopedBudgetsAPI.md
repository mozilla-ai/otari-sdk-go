# \ScopedBudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScopedBudgetsCreateScopedBudget**](ScopedBudgetsAPI.md#ScopedBudgetsCreateScopedBudget) | **Post** /api/v1/scoped-budgets | Create Scoped Budget
[**ScopedBudgetsDeleteScopedBudget**](ScopedBudgetsAPI.md#ScopedBudgetsDeleteScopedBudget) | **Delete** /api/v1/scoped-budgets/{budget_id} | Delete Scoped Budget
[**ScopedBudgetsGetScopedBudget**](ScopedBudgetsAPI.md#ScopedBudgetsGetScopedBudget) | **Get** /api/v1/scoped-budgets/{budget_id} | Get Scoped Budget
[**ScopedBudgetsListScopedBudgets**](ScopedBudgetsAPI.md#ScopedBudgetsListScopedBudgets) | **Get** /api/v1/scoped-budgets | List Scoped Budgets
[**ScopedBudgetsUpdateScopedBudget**](ScopedBudgetsAPI.md#ScopedBudgetsUpdateScopedBudget) | **Patch** /api/v1/scoped-budgets/{budget_id} | Update Scoped Budget



## ScopedBudgetsCreateScopedBudget

> ScopedBudgetResponse ScopedBudgetsCreateScopedBudget(ctx).CreateScopedBudgetRequest(createScopedBudgetRequest).Execute()

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
	resp, r, err := apiClient.ScopedBudgetsAPI.ScopedBudgetsCreateScopedBudget(context.Background()).CreateScopedBudgetRequest(createScopedBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ScopedBudgetsCreateScopedBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScopedBudgetsCreateScopedBudget`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.ScopedBudgetsCreateScopedBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScopedBudgetsCreateScopedBudgetRequest struct via the builder pattern


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


## ScopedBudgetsDeleteScopedBudget

> ScopedBudgetsDeleteScopedBudget(ctx, budgetId).Execute()

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
	r, err := apiClient.ScopedBudgetsAPI.ScopedBudgetsDeleteScopedBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ScopedBudgetsDeleteScopedBudget``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScopedBudgetsDeleteScopedBudgetRequest struct via the builder pattern


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


## ScopedBudgetsGetScopedBudget

> ScopedBudgetResponse ScopedBudgetsGetScopedBudget(ctx, budgetId).Execute()

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
	resp, r, err := apiClient.ScopedBudgetsAPI.ScopedBudgetsGetScopedBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ScopedBudgetsGetScopedBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScopedBudgetsGetScopedBudget`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.ScopedBudgetsGetScopedBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScopedBudgetsGetScopedBudgetRequest struct via the builder pattern


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


## ScopedBudgetsListScopedBudgets

> []ScopedBudgetResponse ScopedBudgetsListScopedBudgets(ctx).ScopeType(scopeType).ScopeId(scopeId).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.ScopedBudgetsAPI.ScopedBudgetsListScopedBudgets(context.Background()).ScopeType(scopeType).ScopeId(scopeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ScopedBudgetsListScopedBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScopedBudgetsListScopedBudgets`: []ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.ScopedBudgetsListScopedBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScopedBudgetsListScopedBudgetsRequest struct via the builder pattern


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


## ScopedBudgetsUpdateScopedBudget

> ScopedBudgetResponse ScopedBudgetsUpdateScopedBudget(ctx, budgetId).UpdateScopedBudgetRequest(updateScopedBudgetRequest).Execute()

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
	resp, r, err := apiClient.ScopedBudgetsAPI.ScopedBudgetsUpdateScopedBudget(context.Background(), budgetId).UpdateScopedBudgetRequest(updateScopedBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopedBudgetsAPI.ScopedBudgetsUpdateScopedBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScopedBudgetsUpdateScopedBudget`: ScopedBudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `ScopedBudgetsAPI.ScopedBudgetsUpdateScopedBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScopedBudgetsUpdateScopedBudgetRequest struct via the builder pattern


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

