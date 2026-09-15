# \OrganizationBudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationBudgetsCreateOrganizationBudget**](OrganizationBudgetsAPI.md#OrganizationBudgetsCreateOrganizationBudget) | **Post** /api/v1/organizations/me/budgets | Create Organization Budget
[**OrganizationBudgetsCreateOrganizationSpendCeiling**](OrganizationBudgetsAPI.md#OrganizationBudgetsCreateOrganizationSpendCeiling) | **Post** /api/v1/organizations/me/spend-ceilings | Create Organization Spend Ceiling
[**OrganizationBudgetsDeleteOrganizationBudget**](OrganizationBudgetsAPI.md#OrganizationBudgetsDeleteOrganizationBudget) | **Delete** /api/v1/organizations/me/budgets/{budget_id} | Delete Organization Budget
[**OrganizationBudgetsDeleteOrganizationSpendCeiling**](OrganizationBudgetsAPI.md#OrganizationBudgetsDeleteOrganizationSpendCeiling) | **Delete** /api/v1/organizations/me/spend-ceilings/{ceiling_id} | Delete Organization Spend Ceiling
[**OrganizationBudgetsListOrganizationBudgets**](OrganizationBudgetsAPI.md#OrganizationBudgetsListOrganizationBudgets) | **Get** /api/v1/organizations/me/budgets | List Organization Budgets
[**OrganizationBudgetsListOrganizationSpendCeilings**](OrganizationBudgetsAPI.md#OrganizationBudgetsListOrganizationSpendCeilings) | **Get** /api/v1/organizations/me/spend-ceilings | List Organization Spend Ceilings
[**OrganizationBudgetsUpdateOrganizationBudget**](OrganizationBudgetsAPI.md#OrganizationBudgetsUpdateOrganizationBudget) | **Patch** /api/v1/organizations/me/budgets/{budget_id} | Update Organization Budget
[**OrganizationBudgetsUpdateOrganizationSpendCeiling**](OrganizationBudgetsAPI.md#OrganizationBudgetsUpdateOrganizationSpendCeiling) | **Patch** /api/v1/organizations/me/spend-ceilings/{ceiling_id} | Update Organization Spend Ceiling



## OrganizationBudgetsCreateOrganizationBudget

> OrganizationBudgetPublic OrganizationBudgetsCreateOrganizationBudget(ctx).OrganizationBudgetCreate(organizationBudgetCreate).Execute()

Create Organization Budget



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
	organizationBudgetCreate := *openapiclient.NewOrganizationBudgetCreate() // OrganizationBudgetCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationBudget(context.Background()).OrganizationBudgetCreate(organizationBudgetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsCreateOrganizationBudget`: OrganizationBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsCreateOrganizationBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationBudgetCreate** | [**OrganizationBudgetCreate**](OrganizationBudgetCreate.md) |  | 

### Return type

[**OrganizationBudgetPublic**](OrganizationBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsCreateOrganizationSpendCeiling

> OrganizationScopedBudgetPublic OrganizationBudgetsCreateOrganizationSpendCeiling(ctx).OrganizationScopedBudgetCreate(organizationScopedBudgetCreate).Execute()

Create Organization Spend Ceiling



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
	organizationScopedBudgetCreate := *openapiclient.NewOrganizationScopedBudgetCreate("BudgetId_example", "ScopeId_example", "ScopeType_example") // OrganizationScopedBudgetCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationSpendCeiling(context.Background()).OrganizationScopedBudgetCreate(organizationScopedBudgetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationSpendCeiling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsCreateOrganizationSpendCeiling`: OrganizationScopedBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsCreateOrganizationSpendCeiling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsCreateOrganizationSpendCeilingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationScopedBudgetCreate** | [**OrganizationScopedBudgetCreate**](OrganizationScopedBudgetCreate.md) |  | 

### Return type

[**OrganizationScopedBudgetPublic**](OrganizationScopedBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsDeleteOrganizationBudget

> Message OrganizationBudgetsDeleteOrganizationBudget(ctx, budgetId).Execute()

Delete Organization Budget



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
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationBudget(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsDeleteOrganizationBudget`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsDeleteOrganizationBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsDeleteOrganizationSpendCeiling

> Message OrganizationBudgetsDeleteOrganizationSpendCeiling(ctx, ceilingId).Execute()

Delete Organization Spend Ceiling



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
	ceilingId := "ceilingId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationSpendCeiling(context.Background(), ceilingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationSpendCeiling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsDeleteOrganizationSpendCeiling`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsDeleteOrganizationSpendCeiling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ceilingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsDeleteOrganizationSpendCeilingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsListOrganizationBudgets

> OrganizationBudgetsPublic OrganizationBudgetsListOrganizationBudgets(ctx).Skip(skip).Limit(limit).Execute()

List Organization Budgets



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationBudgets(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsListOrganizationBudgets`: OrganizationBudgetsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsListOrganizationBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrganizationBudgetsPublic**](OrganizationBudgetsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsListOrganizationSpendCeilings

> OrganizationScopedBudgetsPublic OrganizationBudgetsListOrganizationSpendCeilings(ctx).Skip(skip).Limit(limit).Execute()

List Organization Spend Ceilings



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationSpendCeilings(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationSpendCeilings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsListOrganizationSpendCeilings`: OrganizationScopedBudgetsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsListOrganizationSpendCeilings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsListOrganizationSpendCeilingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrganizationScopedBudgetsPublic**](OrganizationScopedBudgetsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsUpdateOrganizationBudget

> OrganizationBudgetPublic OrganizationBudgetsUpdateOrganizationBudget(ctx, budgetId).OrganizationBudgetUpdate(organizationBudgetUpdate).Execute()

Update Organization Budget



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
	organizationBudgetUpdate := *openapiclient.NewOrganizationBudgetUpdate() // OrganizationBudgetUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationBudget(context.Background(), budgetId).OrganizationBudgetUpdate(organizationBudgetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsUpdateOrganizationBudget`: OrganizationBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsUpdateOrganizationBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationBudgetUpdate** | [**OrganizationBudgetUpdate**](OrganizationBudgetUpdate.md) |  | 

### Return type

[**OrganizationBudgetPublic**](OrganizationBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationBudgetsUpdateOrganizationSpendCeiling

> OrganizationScopedBudgetPublic OrganizationBudgetsUpdateOrganizationSpendCeiling(ctx, ceilingId).OrganizationScopedBudgetUpdate(organizationScopedBudgetUpdate).Execute()

Update Organization Spend Ceiling



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
	ceilingId := "ceilingId_example" // string | 
	organizationScopedBudgetUpdate := *openapiclient.NewOrganizationScopedBudgetUpdate() // OrganizationScopedBudgetUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationSpendCeiling(context.Background(), ceilingId).OrganizationScopedBudgetUpdate(organizationScopedBudgetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationSpendCeiling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationBudgetsUpdateOrganizationSpendCeiling`: OrganizationScopedBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.OrganizationBudgetsUpdateOrganizationSpendCeiling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ceilingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationBudgetsUpdateOrganizationSpendCeilingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationScopedBudgetUpdate** | [**OrganizationScopedBudgetUpdate**](OrganizationScopedBudgetUpdate.md) |  | 

### Return type

[**OrganizationScopedBudgetPublic**](OrganizationScopedBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

