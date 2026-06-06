# \BudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBudgetV1BudgetsPost**](BudgetsAPI.md#CreateBudgetV1BudgetsPost) | **Post** /v1/budgets | Create Budget
[**DeleteBudgetV1BudgetsBudgetIdDelete**](BudgetsAPI.md#DeleteBudgetV1BudgetsBudgetIdDelete) | **Delete** /v1/budgets/{budget_id} | Delete Budget
[**GetBudgetV1BudgetsBudgetIdGet**](BudgetsAPI.md#GetBudgetV1BudgetsBudgetIdGet) | **Get** /v1/budgets/{budget_id} | Get Budget
[**ListBudgetsV1BudgetsGet**](BudgetsAPI.md#ListBudgetsV1BudgetsGet) | **Get** /v1/budgets | List Budgets
[**UpdateBudgetV1BudgetsBudgetIdPatch**](BudgetsAPI.md#UpdateBudgetV1BudgetsBudgetIdPatch) | **Patch** /v1/budgets/{budget_id} | Update Budget



## CreateBudgetV1BudgetsPost

> BudgetResponse CreateBudgetV1BudgetsPost(ctx).CreateBudgetRequest(createBudgetRequest).Execute()

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
	resp, r, err := apiClient.BudgetsAPI.CreateBudgetV1BudgetsPost(context.Background()).CreateBudgetRequest(createBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.CreateBudgetV1BudgetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBudgetV1BudgetsPost`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.CreateBudgetV1BudgetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetV1BudgetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBudgetRequest** | [**CreateBudgetRequest**](CreateBudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBudgetV1BudgetsBudgetIdDelete

> DeleteBudgetV1BudgetsBudgetIdDelete(ctx, budgetId).Execute()

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
	r, err := apiClient.BudgetsAPI.DeleteBudgetV1BudgetsBudgetIdDelete(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.DeleteBudgetV1BudgetsBudgetIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBudgetV1BudgetsBudgetIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetV1BudgetsBudgetIdGet

> BudgetResponse GetBudgetV1BudgetsBudgetIdGet(ctx, budgetId).Execute()

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
	resp, r, err := apiClient.BudgetsAPI.GetBudgetV1BudgetsBudgetIdGet(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetV1BudgetsBudgetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetV1BudgetsBudgetIdGet`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetV1BudgetsBudgetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetV1BudgetsBudgetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgetsV1BudgetsGet

> []BudgetResponse ListBudgetsV1BudgetsGet(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.BudgetsAPI.ListBudgetsV1BudgetsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.ListBudgetsV1BudgetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBudgetsV1BudgetsGet`: []BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.ListBudgetsV1BudgetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetsV1BudgetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudgetV1BudgetsBudgetIdPatch

> BudgetResponse UpdateBudgetV1BudgetsBudgetIdPatch(ctx, budgetId).UpdateBudgetRequest(updateBudgetRequest).Execute()

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
	resp, r, err := apiClient.BudgetsAPI.UpdateBudgetV1BudgetsBudgetIdPatch(context.Background(), budgetId).UpdateBudgetRequest(updateBudgetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UpdateBudgetV1BudgetsBudgetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBudgetV1BudgetsBudgetIdPatch`: BudgetResponse
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UpdateBudgetV1BudgetsBudgetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetV1BudgetsBudgetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBudgetRequest** | [**UpdateBudgetRequest**](UpdateBudgetRequest.md) |  | 

### Return type

[**BudgetResponse**](BudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

