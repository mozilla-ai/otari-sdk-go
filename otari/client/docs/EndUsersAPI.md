# \EndUsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndUsersDeleteEndUser**](EndUsersAPI.md#EndUsersDeleteEndUser) | **Delete** /api/v1/end-users/{external_id} | Delete End User
[**EndUsersDeleteEndUsers**](EndUsersAPI.md#EndUsersDeleteEndUsers) | **Delete** /api/v1/end-users | Delete End Users
[**EndUsersGetEndUser**](EndUsersAPI.md#EndUsersGetEndUser) | **Get** /api/v1/end-users/{external_id} | Get End User
[**EndUsersListEndUserBudgets**](EndUsersAPI.md#EndUsersListEndUserBudgets) | **Get** /api/v1/end-user-budgets | List End User Budgets
[**EndUsersListEndUsers**](EndUsersAPI.md#EndUsersListEndUsers) | **Get** /api/v1/end-users | List End Users
[**EndUsersPutEndUser**](EndUsersAPI.md#EndUsersPutEndUser) | **Put** /api/v1/end-users/{external_id} | Put End User
[**EndUsersPutEndUserByQuery**](EndUsersAPI.md#EndUsersPutEndUserByQuery) | **Put** /api/v1/end-users | Put End User By Query
[**EndUsersUpdateEndUser**](EndUsersAPI.md#EndUsersUpdateEndUser) | **Patch** /api/v1/end-users/{external_id} | Update End User
[**EndUsersUpdateEndUserByQuery**](EndUsersAPI.md#EndUsersUpdateEndUserByQuery) | **Patch** /api/v1/end-users | Update End User By Query



## EndUsersDeleteEndUser

> EndUsersDeleteEndUser(ctx, externalId).Execute()

Delete End User



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndUsersAPI.EndUsersDeleteEndUser(context.Background(), externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersDeleteEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersDeleteEndUserRequest struct via the builder pattern


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


## EndUsersDeleteEndUsers

> EndUsersDeleted EndUsersDeleteEndUsers(ctx).ExternalId(externalId).CreatedBefore(createdBefore).Limit(limit).Execute()

Delete End Users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	externalId := "externalId_example" // string | Delete the end user the service names with this `user` value. (optional)
	createdBefore := time.Now() // time.Time | Delete the end users created before this time, oldest first, up to `limit` of them. (optional)
	limit := int32(56) // int32 | The most end users one call deletes with `created_before`. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersDeleteEndUsers(context.Background()).ExternalId(externalId).CreatedBefore(createdBefore).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersDeleteEndUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersDeleteEndUsers`: EndUsersDeleted
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersDeleteEndUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersDeleteEndUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | Delete the end user the service names with this &#x60;user&#x60; value. | 
 **createdBefore** | **time.Time** | Delete the end users created before this time, oldest first, up to &#x60;limit&#x60; of them. | 
 **limit** | **int32** | The most end users one call deletes with &#x60;created_before&#x60;. | [default to 100]

### Return type

[**EndUsersDeleted**](EndUsersDeleted.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersGetEndUser

> EndUserPublic EndUsersGetEndUser(ctx, externalId).Execute()

Get End User



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersGetEndUser(context.Background(), externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersGetEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersGetEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersGetEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersGetEndUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersListEndUserBudgets

> []EndUserBudgetPublic EndUsersListEndUserBudgets(ctx).Execute()

List End User Budgets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersListEndUserBudgets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersListEndUserBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersListEndUserBudgets`: []EndUserBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersListEndUserBudgets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersListEndUserBudgetsRequest struct via the builder pattern


### Return type

[**[]EndUserBudgetPublic**](EndUserBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersListEndUsers

> []EndUserPublic EndUsersListEndUsers(ctx).Skip(skip).Limit(limit).BudgetId(budgetId).Blocked(blocked).ExternalId(externalId).IncludeTotal(includeTotal).Execute()

List End Users



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
	budgetId := "budgetId_example" // string | Only end users on this budget. (optional)
	blocked := true // bool | Only end users that are, or are not, blocked. (optional)
	externalId := "externalId_example" // string | Only the end user the service names with this `user` value: zero or one. (optional)
	includeTotal := true // bool | Also count every matching end user, in the Otari-Total-Count response header. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersListEndUsers(context.Background()).Skip(skip).Limit(limit).BudgetId(budgetId).Blocked(blocked).ExternalId(externalId).IncludeTotal(includeTotal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersListEndUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersListEndUsers`: []EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersListEndUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersListEndUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **budgetId** | **string** | Only end users on this budget. | 
 **blocked** | **bool** | Only end users that are, or are not, blocked. | 
 **externalId** | **string** | Only the end user the service names with this &#x60;user&#x60; value: zero or one. | 
 **includeTotal** | **bool** | Also count every matching end user, in the Otari-Total-Count response header. | [default to false]

### Return type

[**[]EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersPutEndUser

> EndUserPublic EndUsersPutEndUser(ctx, externalId).EndUserPut(endUserPut).Execute()

Put End User



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field
	endUserPut := *openapiclient.NewEndUserPut("BudgetId_example") // EndUserPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersPutEndUser(context.Background(), externalId).EndUserPut(endUserPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersPutEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersPutEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersPutEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersPutEndUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endUserPut** | [**EndUserPut**](EndUserPut.md) |  | 

### Return type

[**EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersPutEndUserByQuery

> EndUserPublic EndUsersPutEndUserByQuery(ctx).ExternalId(externalId).EndUserPut(endUserPut).Execute()

Put End User By Query



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field, any id, '/' included
	endUserPut := *openapiclient.NewEndUserPut("BudgetId_example") // EndUserPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersPutEndUserByQuery(context.Background()).ExternalId(externalId).EndUserPut(endUserPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersPutEndUserByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersPutEndUserByQuery`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersPutEndUserByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersPutEndUserByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | The id the service names the end user by in a request&#39;s user field, any id, &#39;/&#39; included | 
 **endUserPut** | [**EndUserPut**](EndUserPut.md) |  | 

### Return type

[**EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersUpdateEndUser

> EndUserPublic EndUsersUpdateEndUser(ctx, externalId).EndUserUpdate(endUserUpdate).Execute()

Update End User



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field
	endUserUpdate := *openapiclient.NewEndUserUpdate() // EndUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersUpdateEndUser(context.Background(), externalId).EndUserUpdate(endUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersUpdateEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersUpdateEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersUpdateEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersUpdateEndUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endUserUpdate** | [**EndUserUpdate**](EndUserUpdate.md) |  | 

### Return type

[**EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUsersUpdateEndUserByQuery

> EndUserPublic EndUsersUpdateEndUserByQuery(ctx).ExternalId(externalId).EndUserUpdate(endUserUpdate).Execute()

Update End User By Query



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
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field, any id, '/' included
	endUserUpdate := *openapiclient.NewEndUserUpdate() // EndUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndUsersAPI.EndUsersUpdateEndUserByQuery(context.Background()).ExternalId(externalId).EndUserUpdate(endUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndUsersAPI.EndUsersUpdateEndUserByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersUpdateEndUserByQuery`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `EndUsersAPI.EndUsersUpdateEndUserByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUsersUpdateEndUserByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | The id the service names the end user by in a request&#39;s user field, any id, &#39;/&#39; included | 
 **endUserUpdate** | [**EndUserUpdate**](EndUserUpdate.md) |  | 

### Return type

[**EndUserPublic**](EndUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

