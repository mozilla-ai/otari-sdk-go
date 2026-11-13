# \KeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeysCreateFederationRule**](KeysAPI.md#KeysCreateFederationRule) | **Post** /api/v1/keys/{key_id}/federation-rules | Create Federation Rule
[**KeysCreateKey**](KeysAPI.md#KeysCreateKey) | **Post** /api/v1/keys | Create Key
[**KeysDeleteFederationRule**](KeysAPI.md#KeysDeleteFederationRule) | **Delete** /api/v1/keys/{key_id}/federation-rules/{rule_id} | Delete Federation Rule
[**KeysDeleteKey**](KeysAPI.md#KeysDeleteKey) | **Delete** /api/v1/keys/{key_id} | Delete Key
[**KeysGetEndUser**](KeysAPI.md#KeysGetEndUser) | **Get** /api/v1/keys/{key_id}/end-users/{external_id} | Get End User
[**KeysGetKey**](KeysAPI.md#KeysGetKey) | **Get** /api/v1/keys/{key_id} | Get Key
[**KeysListEndUsers**](KeysAPI.md#KeysListEndUsers) | **Get** /api/v1/keys/{key_id}/end-users | List End Users
[**KeysListFederationRules**](KeysAPI.md#KeysListFederationRules) | **Get** /api/v1/keys/{key_id}/federation-rules | List Federation Rules
[**KeysListKeys**](KeysAPI.md#KeysListKeys) | **Get** /api/v1/keys | List Keys
[**KeysPutEndUser**](KeysAPI.md#KeysPutEndUser) | **Put** /api/v1/keys/{key_id}/end-users/{external_id} | Put End User
[**KeysPutEndUserByQuery**](KeysAPI.md#KeysPutEndUserByQuery) | **Put** /api/v1/keys/{key_id}/end-users | Put End User By Query
[**KeysRotateKey**](KeysAPI.md#KeysRotateKey) | **Post** /api/v1/keys/{key_id}/rotate | Rotate Key
[**KeysUpdateEndUser**](KeysAPI.md#KeysUpdateEndUser) | **Patch** /api/v1/keys/{key_id}/end-users/{external_id} | Update End User
[**KeysUpdateEndUserByQuery**](KeysAPI.md#KeysUpdateEndUserByQuery) | **Patch** /api/v1/keys/{key_id}/end-users | Update End User By Query
[**KeysUpdateKey**](KeysAPI.md#KeysUpdateKey) | **Patch** /api/v1/keys/{key_id} | Update Key



## KeysCreateFederationRule

> FederationRulePublic KeysCreateFederationRule(ctx, keyId).FederationRuleCreate(federationRuleCreate).Execute()

Create Federation Rule



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
	keyId := "keyId_example" // string | 
	federationRuleCreate := *openapiclient.NewFederationRuleCreate("Audience_example", "Issuer_example") // FederationRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysCreateFederationRule(context.Background(), keyId).FederationRuleCreate(federationRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysCreateFederationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysCreateFederationRule`: FederationRulePublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysCreateFederationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysCreateFederationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **federationRuleCreate** | [**FederationRuleCreate**](FederationRuleCreate.md) |  | 

### Return type

[**FederationRulePublic**](FederationRulePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysCreateKey

> CreateKeyResponse KeysCreateKey(ctx).CreateKeyRequest(createKeyRequest).Execute()

Create Key



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
	createKeyRequest := *openapiclient.NewCreateKeyRequest() // CreateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysCreateKey(context.Background()).CreateKeyRequest(createKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysCreateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysCreateKey`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysCreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeysCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKeyRequest** | [**CreateKeyRequest**](CreateKeyRequest.md) |  | 

### Return type

[**CreateKeyResponse**](CreateKeyResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysDeleteFederationRule

> KeysDeleteFederationRule(ctx, keyId, ruleId).Execute()

Delete Federation Rule



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
	keyId := "keyId_example" // string | 
	ruleId := "ruleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.KeysDeleteFederationRule(context.Background(), keyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysDeleteFederationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysDeleteFederationRuleRequest struct via the builder pattern


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


## KeysDeleteKey

> KeysDeleteKey(ctx, keyId).Execute()

Delete Key



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
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.KeysDeleteKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysDeleteKeyRequest struct via the builder pattern


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


## KeysGetEndUser

> EndUserPublic KeysGetEndUser(ctx, keyId, externalId).Execute()

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
	keyId := "keyId_example" // string | 
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysGetEndUser(context.Background(), keyId, externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysGetEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysGetEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysGetEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysGetEndUserRequest struct via the builder pattern


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


## KeysGetKey

> KeyInfo KeysGetKey(ctx, keyId).Execute()

Get Key



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
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysGetKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysGetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysGetKey`: KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysGetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyInfo**](KeyInfo.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysListEndUsers

> []EndUserPublic KeysListEndUsers(ctx, keyId).Skip(skip).Limit(limit).BudgetId(budgetId).Blocked(blocked).ExternalId(externalId).IncludeTotal(includeTotal).Execute()

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
	keyId := "keyId_example" // string | 
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)
	budgetId := "budgetId_example" // string | Only end users on this budget. (optional)
	blocked := true // bool | Only end users that are, or are not, blocked. (optional)
	externalId := "externalId_example" // string | Only the end user the service names with this `user` value: zero or one. (optional)
	includeTotal := true // bool | Also count every matching end user, in the Otari-Total-Count response header. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysListEndUsers(context.Background(), keyId).Skip(skip).Limit(limit).BudgetId(budgetId).Blocked(blocked).ExternalId(externalId).IncludeTotal(includeTotal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysListEndUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysListEndUsers`: []EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysListEndUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysListEndUsersRequest struct via the builder pattern


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


## KeysListFederationRules

> []FederationRulePublic KeysListFederationRules(ctx, keyId).Execute()

List Federation Rules



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
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysListFederationRules(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysListFederationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysListFederationRules`: []FederationRulePublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysListFederationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysListFederationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FederationRulePublic**](FederationRulePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysListKeys

> []KeyInfo KeysListKeys(ctx).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()

List Keys



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only keys in this workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysListKeys(context.Background()).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysListKeys`: []KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysListKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeysListKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]
 **workspaceId** | **string** | Only keys in this workspace. | 

### Return type

[**[]KeyInfo**](KeyInfo.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysPutEndUser

> EndUserPublic KeysPutEndUser(ctx, keyId, externalId).EndUserPut(endUserPut).Execute()

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
	keyId := "keyId_example" // string | 
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field
	endUserPut := *openapiclient.NewEndUserPut("BudgetId_example") // EndUserPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysPutEndUser(context.Background(), keyId, externalId).EndUserPut(endUserPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysPutEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysPutEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysPutEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysPutEndUserRequest struct via the builder pattern


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


## KeysPutEndUserByQuery

> EndUserPublic KeysPutEndUserByQuery(ctx, keyId).ExternalId(externalId).EndUserPut(endUserPut).Execute()

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
	keyId := "keyId_example" // string | 
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field, any id, '/' included
	endUserPut := *openapiclient.NewEndUserPut("BudgetId_example") // EndUserPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysPutEndUserByQuery(context.Background(), keyId).ExternalId(externalId).EndUserPut(endUserPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysPutEndUserByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysPutEndUserByQuery`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysPutEndUserByQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysPutEndUserByQueryRequest struct via the builder pattern


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


## KeysRotateKey

> CreateKeyResponse KeysRotateKey(ctx, keyId).Execute()

Rotate Key



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
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysRotateKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysRotateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysRotateKey`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysRotateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysRotateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateKeyResponse**](CreateKeyResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeysUpdateEndUser

> EndUserPublic KeysUpdateEndUser(ctx, keyId, externalId).EndUserUpdate(endUserUpdate).Execute()

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
	keyId := "keyId_example" // string | 
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field
	endUserUpdate := *openapiclient.NewEndUserUpdate() // EndUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysUpdateEndUser(context.Background(), keyId, externalId).EndUserUpdate(endUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysUpdateEndUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysUpdateEndUser`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysUpdateEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 
**externalId** | **string** | The id the service names the end user by in a request&#39;s user field | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysUpdateEndUserRequest struct via the builder pattern


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


## KeysUpdateEndUserByQuery

> EndUserPublic KeysUpdateEndUserByQuery(ctx, keyId).ExternalId(externalId).EndUserUpdate(endUserUpdate).Execute()

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
	keyId := "keyId_example" // string | 
	externalId := "externalId_example" // string | The id the service names the end user by in a request's user field, any id, '/' included
	endUserUpdate := *openapiclient.NewEndUserUpdate() // EndUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysUpdateEndUserByQuery(context.Background(), keyId).ExternalId(externalId).EndUserUpdate(endUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysUpdateEndUserByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysUpdateEndUserByQuery`: EndUserPublic
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysUpdateEndUserByQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysUpdateEndUserByQueryRequest struct via the builder pattern


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


## KeysUpdateKey

> KeyInfo KeysUpdateKey(ctx, keyId).UpdateKeyRequest(updateKeyRequest).Execute()

Update Key



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
	keyId := "keyId_example" // string | 
	updateKeyRequest := *openapiclient.NewUpdateKeyRequest() // UpdateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KeysUpdateKey(context.Background(), keyId).UpdateKeyRequest(updateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KeysUpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeysUpdateKey`: KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KeysUpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeysUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateKeyRequest** | [**UpdateKeyRequest**](UpdateKeyRequest.md) |  | 

### Return type

[**KeyInfo**](KeyInfo.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

