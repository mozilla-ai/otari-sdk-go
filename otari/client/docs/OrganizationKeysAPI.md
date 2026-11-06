# \OrganizationKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationKeysCreateOwnFederationRule**](OrganizationKeysAPI.md#OrganizationKeysCreateOwnFederationRule) | **Post** /api/v1/organizations/me/keys/{key_id}/federation-rules | Create Own Federation Rule
[**OrganizationKeysCreateOwnKey**](OrganizationKeysAPI.md#OrganizationKeysCreateOwnKey) | **Post** /api/v1/organizations/me/keys | Create Own Key
[**OrganizationKeysDeleteOwnFederationRule**](OrganizationKeysAPI.md#OrganizationKeysDeleteOwnFederationRule) | **Delete** /api/v1/organizations/me/keys/{key_id}/federation-rules/{rule_id} | Delete Own Federation Rule
[**OrganizationKeysDeleteOwnKey**](OrganizationKeysAPI.md#OrganizationKeysDeleteOwnKey) | **Delete** /api/v1/organizations/me/keys/{key_id} | Delete Own Key
[**OrganizationKeysListOwnFederationRules**](OrganizationKeysAPI.md#OrganizationKeysListOwnFederationRules) | **Get** /api/v1/organizations/me/keys/{key_id}/federation-rules | List Own Federation Rules
[**OrganizationKeysListOwnKeys**](OrganizationKeysAPI.md#OrganizationKeysListOwnKeys) | **Get** /api/v1/organizations/me/keys | List Own Keys
[**OrganizationKeysRotateOwnKey**](OrganizationKeysAPI.md#OrganizationKeysRotateOwnKey) | **Post** /api/v1/organizations/me/keys/{key_id}/rotate | Rotate Own Key
[**OrganizationKeysUpdateOwnKey**](OrganizationKeysAPI.md#OrganizationKeysUpdateOwnKey) | **Patch** /api/v1/organizations/me/keys/{key_id} | Update Own Key



## OrganizationKeysCreateOwnFederationRule

> FederationRulePublic OrganizationKeysCreateOwnFederationRule(ctx, keyId).FederationRuleCreate(federationRuleCreate).Execute()

Create Own Federation Rule



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
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysCreateOwnFederationRule(context.Background(), keyId).FederationRuleCreate(federationRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysCreateOwnFederationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysCreateOwnFederationRule`: FederationRulePublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysCreateOwnFederationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysCreateOwnFederationRuleRequest struct via the builder pattern


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


## OrganizationKeysCreateOwnKey

> CreateKeyResponse OrganizationKeysCreateOwnKey(ctx).CreateOwnKeyRequest(createOwnKeyRequest).Execute()

Create Own Key



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
	createOwnKeyRequest := *openapiclient.NewCreateOwnKeyRequest() // CreateOwnKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysCreateOwnKey(context.Background()).CreateOwnKeyRequest(createOwnKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysCreateOwnKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysCreateOwnKey`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysCreateOwnKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysCreateOwnKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOwnKeyRequest** | [**CreateOwnKeyRequest**](CreateOwnKeyRequest.md) |  | 

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


## OrganizationKeysDeleteOwnFederationRule

> OrganizationKeysDeleteOwnFederationRule(ctx, keyId, ruleId).Execute()

Delete Own Federation Rule



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
	r, err := apiClient.OrganizationKeysAPI.OrganizationKeysDeleteOwnFederationRule(context.Background(), keyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysDeleteOwnFederationRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationKeysDeleteOwnFederationRuleRequest struct via the builder pattern


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


## OrganizationKeysDeleteOwnKey

> OrganizationKeysDeleteOwnKey(ctx, keyId).Execute()

Delete Own Key



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
	r, err := apiClient.OrganizationKeysAPI.OrganizationKeysDeleteOwnKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysDeleteOwnKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationKeysDeleteOwnKeyRequest struct via the builder pattern


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


## OrganizationKeysListOwnFederationRules

> []FederationRulePublic OrganizationKeysListOwnFederationRules(ctx, keyId).Execute()

List Own Federation Rules



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
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysListOwnFederationRules(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysListOwnFederationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysListOwnFederationRules`: []FederationRulePublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysListOwnFederationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysListOwnFederationRulesRequest struct via the builder pattern


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


## OrganizationKeysListOwnKeys

> []KeyInfo OrganizationKeysListOwnKeys(ctx).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()

List Own Keys



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
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysListOwnKeys(context.Background()).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysListOwnKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysListOwnKeys`: []KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysListOwnKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysListOwnKeysRequest struct via the builder pattern


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


## OrganizationKeysRotateOwnKey

> CreateKeyResponse OrganizationKeysRotateOwnKey(ctx, keyId).Execute()

Rotate Own Key



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
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysRotateOwnKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysRotateOwnKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysRotateOwnKey`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysRotateOwnKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysRotateOwnKeyRequest struct via the builder pattern


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


## OrganizationKeysUpdateOwnKey

> KeyInfo OrganizationKeysUpdateOwnKey(ctx, keyId).UpdateOwnKeyRequest(updateOwnKeyRequest).Execute()

Update Own Key



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
	updateOwnKeyRequest := *openapiclient.NewUpdateOwnKeyRequest() // UpdateOwnKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationKeysAPI.OrganizationKeysUpdateOwnKey(context.Background(), keyId).UpdateOwnKeyRequest(updateOwnKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.OrganizationKeysUpdateOwnKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationKeysUpdateOwnKey`: KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.OrganizationKeysUpdateOwnKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationKeysUpdateOwnKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOwnKeyRequest** | [**UpdateOwnKeyRequest**](UpdateOwnKeyRequest.md) |  | 

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

