# \BudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BudgetsCreateBudget**](BudgetsAPI.md#BudgetsCreateBudget) | **Post** /api/v1/budgets | Create Budget
[**BudgetsDeleteBudget**](BudgetsAPI.md#BudgetsDeleteBudget) | **Delete** /api/v1/budgets/{budget_id} | Delete Budget
[**BudgetsGetBudget**](BudgetsAPI.md#BudgetsGetBudget) | **Get** /api/v1/budgets/{budget_id} | Get Budget
[**BudgetsListBudgetResetLogs**](BudgetsAPI.md#BudgetsListBudgetResetLogs) | **Get** /api/v1/budgets/{budget_id}/reset-logs | List Budget Reset Logs
[**BudgetsListBudgets**](BudgetsAPI.md#BudgetsListBudgets) | **Get** /api/v1/budgets | List Budgets
[**BudgetsUpdateBudget**](BudgetsAPI.md#BudgetsUpdateBudget) | **Patch** /api/v1/budgets/{budget_id} | Update Budget



## BudgetsCreateBudget

> BudgetResponse BudgetsCreateBudget(ctx).CreateBudgetRequest(createBudgetRequest).Execute()

Create Budget



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
	createBudgetRequest := *openapiclient.NewCreateBudgetRequest() // CreateBudgetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.BudgetsCreateBudget(context.Background()).CreateBudgetRequest(createBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsCreateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BudgetsCreateBudget`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.BudgetsCreateBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBudgetsCreateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBudgetRequest** | [**CreateBudgetRequest**](CreateBudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BudgetsDeleteBudget

> BudgetsDeleteBudget(ctx, budgetId).Execute()

Delete Budget



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
	r, err := apiClient.BudgetsAPI.BudgetsDeleteBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsDeleteBudget``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBudgetsDeleteBudgetRequest struct via the builder pattern


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


## BudgetsGetBudget

> BudgetResponse BudgetsGetBudget(ctx, budgetId).Execute()

Get Budget



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
	resp, r, err := apiClient.BudgetsAPI.BudgetsGetBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsGetBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BudgetsGetBudget`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.BudgetsGetBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBudgetsGetBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BudgetsListBudgetResetLogs

> []BudgetResetLogResponse BudgetsListBudgetResetLogs(ctx, budgetId).Skip(skip).Limit(limit).Execute()

List Budget Reset Logs



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.BudgetsListBudgetResetLogs(context.Background(), budgetId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsListBudgetResetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BudgetsListBudgetResetLogs`: []BudgetResetLogResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.BudgetsListBudgetResetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBudgetsListBudgetResetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]BudgetResetLogResponse**](BudgetResetLogResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BudgetsListBudgets

> []BudgetResponse BudgetsListBudgets(ctx).Skip(skip).Limit(limit).Execute()

List Budgets



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.BudgetsListBudgets(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsListBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BudgetsListBudgets`: []BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.BudgetsListBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBudgetsListBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]BudgetResponse**](BudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BudgetsUpdateBudget

> BudgetResponse BudgetsUpdateBudget(ctx, budgetId).UpdateBudgetRequest(updateBudgetRequest).Execute()

Update Budget



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
	updateBudgetRequest := *openapiclient.NewUpdateBudgetRequest() // UpdateBudgetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.BudgetsUpdateBudget(context.Background(), budgetId).UpdateBudgetRequest(updateBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.BudgetsUpdateBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BudgetsUpdateBudget`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.BudgetsUpdateBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBudgetsUpdateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBudgetRequest** | [**UpdateBudgetRequest**](UpdateBudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

