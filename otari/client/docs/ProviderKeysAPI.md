# \ProviderKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProviderKeysAddWorkspaceProviderKeyModelRestriction**](ProviderKeysAPI.md#ProviderKeysAddWorkspaceProviderKeyModelRestriction) | **Post** /api/v1/workspaces/{workspace_id}/provider-keys/{key_id}/models | Add Workspace Provider Key Model Restriction
[**ProviderKeysArchiveOrgProviderKey**](ProviderKeysAPI.md#ProviderKeysArchiveOrgProviderKey) | **Post** /api/v1/organizations/me/provider-keys/{key_id}/archive | Archive Org Provider Key
[**ProviderKeysCreateOrgProviderKey**](ProviderKeysAPI.md#ProviderKeysCreateOrgProviderKey) | **Post** /api/v1/organizations/me/provider-keys | Create Org Provider Key
[**ProviderKeysDeleteOrgProviderKey**](ProviderKeysAPI.md#ProviderKeysDeleteOrgProviderKey) | **Delete** /api/v1/organizations/me/provider-keys/{key_id} | Delete Org Provider Key
[**ProviderKeysListOrgProviderKeys**](ProviderKeysAPI.md#ProviderKeysListOrgProviderKeys) | **Get** /api/v1/organizations/me/provider-keys | List Org Provider Keys
[**ProviderKeysListWorkspaceProviderKeyModelRestrictions**](ProviderKeysAPI.md#ProviderKeysListWorkspaceProviderKeyModelRestrictions) | **Get** /api/v1/workspaces/{workspace_id}/provider-keys/{key_id}/models | List Workspace Provider Key Model Restrictions
[**ProviderKeysListWorkspaceProviderKeys**](ProviderKeysAPI.md#ProviderKeysListWorkspaceProviderKeys) | **Get** /api/v1/workspaces/{workspace_id}/provider-keys | List Workspace Provider Keys
[**ProviderKeysRemoveWorkspaceProviderKeyModelRestriction**](ProviderKeysAPI.md#ProviderKeysRemoveWorkspaceProviderKeyModelRestriction) | **Delete** /api/v1/workspaces/{workspace_id}/provider-keys/{key_id}/models/{model} | Remove Workspace Provider Key Model Restriction
[**ProviderKeysResetWorkspaceProviderKeyOverride**](ProviderKeysAPI.md#ProviderKeysResetWorkspaceProviderKeyOverride) | **Delete** /api/v1/workspaces/{workspace_id}/provider-keys/{key_id} | Reset Workspace Provider Key Override
[**ProviderKeysRestoreOrgProviderKey**](ProviderKeysAPI.md#ProviderKeysRestoreOrgProviderKey) | **Post** /api/v1/organizations/me/provider-keys/{key_id}/restore | Restore Org Provider Key
[**ProviderKeysSetOrgProviderKeyDefault**](ProviderKeysAPI.md#ProviderKeysSetOrgProviderKeyDefault) | **Post** /api/v1/organizations/me/provider-keys/{key_id}/default | Set Org Provider Key Default
[**ProviderKeysSetWorkspaceProviderKeyOverride**](ProviderKeysAPI.md#ProviderKeysSetWorkspaceProviderKeyOverride) | **Patch** /api/v1/workspaces/{workspace_id}/provider-keys/{key_id} | Set Workspace Provider Key Override
[**ProviderKeysUpdateOrgProviderKey**](ProviderKeysAPI.md#ProviderKeysUpdateOrgProviderKey) | **Patch** /api/v1/organizations/me/provider-keys/{key_id} | Update Org Provider Key



## ProviderKeysAddWorkspaceProviderKeyModelRestriction

> Message ProviderKeysAddWorkspaceProviderKeyModelRestriction(ctx, workspaceId, keyId).WorkspaceProviderModelRestrictionRequest(workspaceProviderModelRestrictionRequest).Execute()

Add Workspace Provider Key Model Restriction



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceProviderModelRestrictionRequest := *openapiclient.NewWorkspaceProviderModelRestrictionRequest("Model_example") // WorkspaceProviderModelRestrictionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysAddWorkspaceProviderKeyModelRestriction(context.Background(), workspaceId, keyId).WorkspaceProviderModelRestrictionRequest(workspaceProviderModelRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysAddWorkspaceProviderKeyModelRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysAddWorkspaceProviderKeyModelRestriction`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysAddWorkspaceProviderKeyModelRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysAddWorkspaceProviderKeyModelRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceProviderModelRestrictionRequest** | [**WorkspaceProviderModelRestrictionRequest**](WorkspaceProviderModelRestrictionRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysArchiveOrgProviderKey

> OrgProviderKeyPublic ProviderKeysArchiveOrgProviderKey(ctx, keyId).Execute()

Archive Org Provider Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysArchiveOrgProviderKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysArchiveOrgProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysArchiveOrgProviderKey`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysArchiveOrgProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysArchiveOrgProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgProviderKeyPublic**](OrgProviderKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysCreateOrgProviderKey

> OrgProviderKeyPublic ProviderKeysCreateOrgProviderKey(ctx).OrgProviderKeyCreateRequest(orgProviderKeyCreateRequest).Execute()

Create Org Provider Key



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
	orgProviderKeyCreateRequest := *openapiclient.NewOrgProviderKeyCreateRequest("Name_example", "Provider_example") // OrgProviderKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysCreateOrgProviderKey(context.Background()).OrgProviderKeyCreateRequest(orgProviderKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysCreateOrgProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysCreateOrgProviderKey`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysCreateOrgProviderKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysCreateOrgProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgProviderKeyCreateRequest** | [**OrgProviderKeyCreateRequest**](OrgProviderKeyCreateRequest.md) |  | 

### Return type

[**OrgProviderKeyPublic**](OrgProviderKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysDeleteOrgProviderKey

> Message ProviderKeysDeleteOrgProviderKey(ctx, keyId).Execute()

Delete Org Provider Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysDeleteOrgProviderKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysDeleteOrgProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysDeleteOrgProviderKey`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysDeleteOrgProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysDeleteOrgProviderKeyRequest struct via the builder pattern


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


## ProviderKeysListOrgProviderKeys

> OrgProviderKeysPublic ProviderKeysListOrgProviderKeys(ctx).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()

List Org Provider Keys



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
	includeArchived := true // bool | Include archived keys. (optional) (default to false)
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysListOrgProviderKeys(context.Background()).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysListOrgProviderKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysListOrgProviderKeys`: OrgProviderKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysListOrgProviderKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysListOrgProviderKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArchived** | **bool** | Include archived keys. | [default to false]
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrgProviderKeysPublic**](OrgProviderKeysPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysListWorkspaceProviderKeyModelRestrictions

> WorkspaceProviderModelRestrictionsPublic ProviderKeysListWorkspaceProviderKeyModelRestrictions(ctx, workspaceId, keyId).Execute()

List Workspace Provider Key Model Restrictions



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeyModelRestrictions(context.Background(), workspaceId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeyModelRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysListWorkspaceProviderKeyModelRestrictions`: WorkspaceProviderModelRestrictionsPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeyModelRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysListWorkspaceProviderKeyModelRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceProviderModelRestrictionsPublic**](WorkspaceProviderModelRestrictionsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysListWorkspaceProviderKeys

> WorkspaceProviderKeyOverridesPublic ProviderKeysListWorkspaceProviderKeys(ctx, workspaceId).Execute()

List Workspace Provider Keys



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeys(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysListWorkspaceProviderKeys`: WorkspaceProviderKeyOverridesPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysListWorkspaceProviderKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysListWorkspaceProviderKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceProviderKeyOverridesPublic**](WorkspaceProviderKeyOverridesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysRemoveWorkspaceProviderKeyModelRestriction

> Message ProviderKeysRemoveWorkspaceProviderKeyModelRestriction(ctx, workspaceId, keyId, model).Execute()

Remove Workspace Provider Key Model Restriction



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysRemoveWorkspaceProviderKeyModelRestriction(context.Background(), workspaceId, keyId, model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysRemoveWorkspaceProviderKeyModelRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysRemoveWorkspaceProviderKeyModelRestriction`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysRemoveWorkspaceProviderKeyModelRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysRemoveWorkspaceProviderKeyModelRestrictionRequest struct via the builder pattern


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


## ProviderKeysResetWorkspaceProviderKeyOverride

> Message ProviderKeysResetWorkspaceProviderKeyOverride(ctx, workspaceId, keyId).Execute()

Reset Workspace Provider Key Override



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysResetWorkspaceProviderKeyOverride(context.Background(), workspaceId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysResetWorkspaceProviderKeyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysResetWorkspaceProviderKeyOverride`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysResetWorkspaceProviderKeyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysResetWorkspaceProviderKeyOverrideRequest struct via the builder pattern


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


## ProviderKeysRestoreOrgProviderKey

> OrgProviderKeyPublic ProviderKeysRestoreOrgProviderKey(ctx, keyId).Execute()

Restore Org Provider Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysRestoreOrgProviderKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysRestoreOrgProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysRestoreOrgProviderKey`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysRestoreOrgProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysRestoreOrgProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgProviderKeyPublic**](OrgProviderKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysSetOrgProviderKeyDefault

> OrgProviderKeyPublic ProviderKeysSetOrgProviderKeyDefault(ctx, keyId).Execute()

Set Org Provider Key Default



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysSetOrgProviderKeyDefault(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysSetOrgProviderKeyDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysSetOrgProviderKeyDefault`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysSetOrgProviderKeyDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysSetOrgProviderKeyDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgProviderKeyPublic**](OrgProviderKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysSetWorkspaceProviderKeyOverride

> WorkspaceProviderKeyOverridePublic ProviderKeysSetWorkspaceProviderKeyOverride(ctx, workspaceId, keyId).WorkspaceProviderKeyOverrideRequest(workspaceProviderKeyOverrideRequest).Execute()

Set Workspace Provider Key Override



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceProviderKeyOverrideRequest := *openapiclient.NewWorkspaceProviderKeyOverrideRequest() // WorkspaceProviderKeyOverrideRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysSetWorkspaceProviderKeyOverride(context.Background(), workspaceId, keyId).WorkspaceProviderKeyOverrideRequest(workspaceProviderKeyOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysSetWorkspaceProviderKeyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysSetWorkspaceProviderKeyOverride`: WorkspaceProviderKeyOverridePublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysSetWorkspaceProviderKeyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysSetWorkspaceProviderKeyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceProviderKeyOverrideRequest** | [**WorkspaceProviderKeyOverrideRequest**](WorkspaceProviderKeyOverrideRequest.md) |  | 

### Return type

[**WorkspaceProviderKeyOverridePublic**](WorkspaceProviderKeyOverridePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderKeysUpdateOrgProviderKey

> OrgProviderKeyPublic ProviderKeysUpdateOrgProviderKey(ctx, keyId).OrgProviderKeyUpdateRequest(orgProviderKeyUpdateRequest).Execute()

Update Org Provider Key



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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orgProviderKeyUpdateRequest := *openapiclient.NewOrgProviderKeyUpdateRequest() // OrgProviderKeyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProviderKeysAPI.ProviderKeysUpdateOrgProviderKey(context.Background(), keyId).OrgProviderKeyUpdateRequest(orgProviderKeyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ProviderKeysUpdateOrgProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderKeysUpdateOrgProviderKey`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ProviderKeysUpdateOrgProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderKeysUpdateOrgProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgProviderKeyUpdateRequest** | [**OrgProviderKeyUpdateRequest**](OrgProviderKeyUpdateRequest.md) |  | 

### Return type

[**OrgProviderKeyPublic**](OrgProviderKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

