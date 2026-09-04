# \ProviderKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost**](ProviderKeysAPI.md#AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost) | **Post** /v1/workspaces/{workspace_id}/provider-keys/{key_id}/models | Add Workspace Provider Key Model Restriction
[**ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost**](ProviderKeysAPI.md#ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost) | **Post** /v1/organizations/me/provider-keys/{key_id}/archive | Archive Org Provider Key
[**CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost**](ProviderKeysAPI.md#CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost) | **Post** /v1/organizations/me/provider-keys | Create Org Provider Key
[**DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete**](ProviderKeysAPI.md#DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete) | **Delete** /v1/organizations/me/provider-keys/{key_id} | Delete Org Provider Key
[**ListOrgProviderKeysV1OrganizationsMeProviderKeysGet**](ProviderKeysAPI.md#ListOrgProviderKeysV1OrganizationsMeProviderKeysGet) | **Get** /v1/organizations/me/provider-keys | List Org Provider Keys
[**ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet**](ProviderKeysAPI.md#ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet) | **Get** /v1/workspaces/{workspace_id}/provider-keys/{key_id}/models | List Workspace Provider Key Model Restrictions
[**ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet**](ProviderKeysAPI.md#ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet) | **Get** /v1/workspaces/{workspace_id}/provider-keys | List Workspace Provider Keys
[**RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete**](ProviderKeysAPI.md#RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete) | **Delete** /v1/workspaces/{workspace_id}/provider-keys/{key_id}/models/{model} | Remove Workspace Provider Key Model Restriction
[**ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete**](ProviderKeysAPI.md#ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete) | **Delete** /v1/workspaces/{workspace_id}/provider-keys/{key_id} | Reset Workspace Provider Key Override
[**RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost**](ProviderKeysAPI.md#RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost) | **Post** /v1/organizations/me/provider-keys/{key_id}/restore | Restore Org Provider Key
[**SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost**](ProviderKeysAPI.md#SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost) | **Post** /v1/organizations/me/provider-keys/{key_id}/default | Set Org Provider Key Default
[**SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch**](ProviderKeysAPI.md#SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch) | **Patch** /v1/workspaces/{workspace_id}/provider-keys/{key_id} | Set Workspace Provider Key Override
[**UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch**](ProviderKeysAPI.md#UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch) | **Patch** /v1/organizations/me/provider-keys/{key_id} | Update Org Provider Key



## AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost

> Message AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost(ctx, workspaceId, keyId).WorkspaceProviderModelRestrictionRequest(workspaceProviderModelRestrictionRequest).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost(context.Background(), workspaceId, keyId).WorkspaceProviderModelRestrictionRequest(workspaceProviderModelRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.AddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsPostRequest struct via the builder pattern


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


## ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost

> OrgProviderKeyPublic ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost(ctx, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdArchivePostRequest struct via the builder pattern


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


## CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost

> OrgProviderKeyPublic CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost(ctx).OrgProviderKeyCreateRequest(orgProviderKeyCreateRequest).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost(context.Background()).OrgProviderKeyCreateRequest(orgProviderKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.CreateOrgProviderKeyV1OrganizationsMeProviderKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgProviderKeyV1OrganizationsMeProviderKeysPostRequest struct via the builder pattern


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


## DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete

> Message DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete(ctx, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.DeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdDeleteRequest struct via the builder pattern


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


## ListOrgProviderKeysV1OrganizationsMeProviderKeysGet

> OrgProviderKeysPublic ListOrgProviderKeysV1OrganizationsMeProviderKeysGet(ctx).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.ListOrgProviderKeysV1OrganizationsMeProviderKeysGet(context.Background()).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ListOrgProviderKeysV1OrganizationsMeProviderKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgProviderKeysV1OrganizationsMeProviderKeysGet`: OrgProviderKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ListOrgProviderKeysV1OrganizationsMeProviderKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgProviderKeysV1OrganizationsMeProviderKeysGetRequest struct via the builder pattern


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


## ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet

> WorkspaceProviderModelRestrictionsPublic ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet(ctx, workspaceId, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet(context.Background(), workspaceId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet`: WorkspaceProviderModelRestrictionsPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceProviderKeyModelRestrictionsV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsGetRequest struct via the builder pattern


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


## ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet

> WorkspaceProviderKeyOverridesPublic ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet(ctx, workspaceId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet`: WorkspaceProviderKeyOverridesPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceProviderKeysV1WorkspacesWorkspaceIdProviderKeysGetRequest struct via the builder pattern


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


## RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete

> Message RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete(ctx, workspaceId, keyId, model).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete(context.Background(), workspaceId, keyId, model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.RemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRemoveWorkspaceProviderKeyModelRestrictionV1WorkspacesWorkspaceIdProviderKeysKeyIdModelsModelDeleteRequest struct via the builder pattern


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


## ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete

> Message ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete(ctx, workspaceId, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete(context.Background(), workspaceId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.ResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdDeleteRequest struct via the builder pattern


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


## RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost

> OrgProviderKeyPublic RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost(ctx, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.RestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdRestorePostRequest struct via the builder pattern


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


## SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost

> OrgProviderKeyPublic SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost(ctx, keyId).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.SetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrgProviderKeyDefaultV1OrganizationsMeProviderKeysKeyIdDefaultPostRequest struct via the builder pattern


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


## SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch

> WorkspaceProviderKeyOverridePublic SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch(ctx, workspaceId, keyId).WorkspaceProviderKeyOverrideRequest(workspaceProviderKeyOverrideRequest).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch(context.Background(), workspaceId, keyId).WorkspaceProviderKeyOverrideRequest(workspaceProviderKeyOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch`: WorkspaceProviderKeyOverridePublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.SetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorkspaceProviderKeyOverrideV1WorkspacesWorkspaceIdProviderKeysKeyIdPatchRequest struct via the builder pattern


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


## UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch

> OrgProviderKeyPublic UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch(ctx, keyId).OrgProviderKeyUpdateRequest(orgProviderKeyUpdateRequest).Execute()

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
	resp, r, err := apiClient.ProviderKeysAPI.UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch(context.Background(), keyId).OrgProviderKeyUpdateRequest(orgProviderKeyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProviderKeysAPI.UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch`: OrgProviderKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `ProviderKeysAPI.UpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgProviderKeyV1OrganizationsMeProviderKeysKeyIdPatchRequest struct via the builder pattern


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

