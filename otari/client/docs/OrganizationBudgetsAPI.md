# \OrganizationBudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationBudgetV1OrganizationsMeBudgetsPost**](OrganizationBudgetsAPI.md#CreateOrganizationBudgetV1OrganizationsMeBudgetsPost) | **Post** /v1/organizations/me/budgets | Create Organization Budget
[**CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost**](OrganizationBudgetsAPI.md#CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost) | **Post** /v1/organizations/me/spend-ceilings | Create Organization Spend Ceiling
[**DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete**](OrganizationBudgetsAPI.md#DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete) | **Delete** /v1/organizations/me/budgets/{budget_id} | Delete Organization Budget
[**DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete**](OrganizationBudgetsAPI.md#DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete) | **Delete** /v1/organizations/me/spend-ceilings/{ceiling_id} | Delete Organization Spend Ceiling
[**ListOrganizationBudgetsV1OrganizationsMeBudgetsGet**](OrganizationBudgetsAPI.md#ListOrganizationBudgetsV1OrganizationsMeBudgetsGet) | **Get** /v1/organizations/me/budgets | List Organization Budgets
[**ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet**](OrganizationBudgetsAPI.md#ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet) | **Get** /v1/organizations/me/spend-ceilings | List Organization Spend Ceilings
[**UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch**](OrganizationBudgetsAPI.md#UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch) | **Patch** /v1/organizations/me/budgets/{budget_id} | Update Organization Budget
[**UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch**](OrganizationBudgetsAPI.md#UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch) | **Patch** /v1/organizations/me/spend-ceilings/{ceiling_id} | Update Organization Spend Ceiling



## CreateOrganizationBudgetV1OrganizationsMeBudgetsPost

> OrganizationBudgetPublic CreateOrganizationBudgetV1OrganizationsMeBudgetsPost(ctx).OrganizationBudgetCreate(organizationBudgetCreate).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.CreateOrganizationBudgetV1OrganizationsMeBudgetsPost(context.Background()).OrganizationBudgetCreate(organizationBudgetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.CreateOrganizationBudgetV1OrganizationsMeBudgetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationBudgetV1OrganizationsMeBudgetsPost`: OrganizationBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.CreateOrganizationBudgetV1OrganizationsMeBudgetsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationBudgetV1OrganizationsMeBudgetsPostRequest struct via the builder pattern


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


## CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost

> OrganizationScopedBudgetPublic CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost(ctx).OrganizationScopedBudgetCreate(organizationScopedBudgetCreate).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost(context.Background()).OrganizationScopedBudgetCreate(organizationScopedBudgetCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost`: OrganizationScopedBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.CreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsPostRequest struct via the builder pattern


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


## DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete

> Message DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete(ctx, budgetId).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete(context.Background(), budgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.DeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdDeleteRequest struct via the builder pattern


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


## DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete

> Message DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete(ctx, ceilingId).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete(context.Background(), ceilingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.DeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ceilingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdDeleteRequest struct via the builder pattern


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


## ListOrganizationBudgetsV1OrganizationsMeBudgetsGet

> OrganizationBudgetsPublic ListOrganizationBudgetsV1OrganizationsMeBudgetsGet(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.ListOrganizationBudgetsV1OrganizationsMeBudgetsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.ListOrganizationBudgetsV1OrganizationsMeBudgetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationBudgetsV1OrganizationsMeBudgetsGet`: OrganizationBudgetsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.ListOrganizationBudgetsV1OrganizationsMeBudgetsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationBudgetsV1OrganizationsMeBudgetsGetRequest struct via the builder pattern


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


## ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet

> OrganizationScopedBudgetsPublic ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet`: OrganizationScopedBudgetsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.ListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationSpendCeilingsV1OrganizationsMeSpendCeilingsGetRequest struct via the builder pattern


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


## UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch

> OrganizationBudgetPublic UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch(ctx, budgetId).OrganizationBudgetUpdate(organizationBudgetUpdate).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch(context.Background(), budgetId).OrganizationBudgetUpdate(organizationBudgetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch`: OrganizationBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.UpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBudgetV1OrganizationsMeBudgetsBudgetIdPatchRequest struct via the builder pattern


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


## UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch

> OrganizationScopedBudgetPublic UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch(ctx, ceilingId).OrganizationScopedBudgetUpdate(organizationScopedBudgetUpdate).Execute()

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
	resp, r, err := apiClient.OrganizationBudgetsAPI.UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch(context.Background(), ceilingId).OrganizationScopedBudgetUpdate(organizationScopedBudgetUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationBudgetsAPI.UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch`: OrganizationScopedBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationBudgetsAPI.UpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ceilingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSpendCeilingV1OrganizationsMeSpendCeilingsCeilingIdPatchRequest struct via the builder pattern


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

