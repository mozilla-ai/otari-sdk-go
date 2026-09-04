# \AliasesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAliasV1AliasesNameDelete**](AliasesAPI.md#DeleteAliasV1AliasesNameDelete) | **Delete** /v1/aliases/{name} | Delete Alias
[**DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete**](AliasesAPI.md#DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete) | **Delete** /v1/organizations/me/aliases/{name} | Delete Organization Alias
[**ListAliasesV1AliasesGet**](AliasesAPI.md#ListAliasesV1AliasesGet) | **Get** /v1/aliases | List Aliases
[**ListVisibleAliasesV1OrganizationsMeAliasesGet**](AliasesAPI.md#ListVisibleAliasesV1OrganizationsMeAliasesGet) | **Get** /v1/organizations/me/aliases | List Visible Aliases
[**SetAliasV1AliasesPost**](AliasesAPI.md#SetAliasV1AliasesPost) | **Post** /v1/aliases | Set Alias
[**SetOrganizationAliasV1OrganizationsMeAliasesPost**](AliasesAPI.md#SetOrganizationAliasV1OrganizationsMeAliasesPost) | **Post** /v1/organizations/me/aliases | Set Organization Alias



## DeleteAliasV1AliasesNameDelete

> DeleteAliasV1AliasesNameDelete(ctx, name).UserId(userId).WorkspaceId(workspaceId).Execute()

Delete Alias



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
	name := "name_example" // string | 
	userId := "userId_example" // string | Delete the alias scoped to this user. Omit to delete the workspace-wide alias of that name. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Delete the alias in this workspace. Omit for the deployment's default workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AliasesAPI.DeleteAliasV1AliasesNameDelete(context.Background(), name).UserId(userId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.DeleteAliasV1AliasesNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasV1AliasesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Delete the alias scoped to this user. Omit to delete the workspace-wide alias of that name. | 
 **workspaceId** | **string** | Delete the alias in this workspace. Omit for the deployment&#39;s default workspace. | 

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


## DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete

> DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete(ctx, name).WorkspaceId(workspaceId).Execute()

Delete Organization Alias



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
	name := "name_example" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Delete the alias in this workspace of the caller's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AliasesAPI.DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete(context.Background(), name).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.DeleteOrganizationAliasV1OrganizationsMeAliasesNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAliasV1OrganizationsMeAliasesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | Delete the alias in this workspace of the caller&#39;s organization. | 

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


## ListAliasesV1AliasesGet

> []AliasResponse ListAliasesV1AliasesGet(ctx).WorkspaceId(workspaceId).Execute()

List Aliases



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit to list the stored entries of every workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.ListAliasesV1AliasesGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.ListAliasesV1AliasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAliasesV1AliasesGet`: []AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.ListAliasesV1AliasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAliasesV1AliasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit to list the stored entries of every workspace. | 

### Return type

[**[]AliasResponse**](AliasResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVisibleAliasesV1OrganizationsMeAliasesGet

> []AliasResponse ListVisibleAliasesV1OrganizationsMeAliasesGet(ctx).Limit(limit).Execute()

List Visible Aliases



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
	limit := int32(56) // int32 | Maximum entries to return, stored and config-file together. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.ListVisibleAliasesV1OrganizationsMeAliasesGet(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.ListVisibleAliasesV1OrganizationsMeAliasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVisibleAliasesV1OrganizationsMeAliasesGet`: []AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.ListVisibleAliasesV1OrganizationsMeAliasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVisibleAliasesV1OrganizationsMeAliasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return, stored and config-file together. | [default to 1000]

### Return type

[**[]AliasResponse**](AliasResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAliasV1AliasesPost

> AliasResponse SetAliasV1AliasesPost(ctx).AliasRequest(aliasRequest).Execute()

Set Alias



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
	aliasRequest := *openapiclient.NewAliasRequest("Name_example", "Target_example") // AliasRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.SetAliasV1AliasesPost(context.Background()).AliasRequest(aliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.SetAliasV1AliasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAliasV1AliasesPost`: AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.SetAliasV1AliasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAliasV1AliasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliasRequest** | [**AliasRequest**](AliasRequest.md) |  | 

### Return type

[**AliasResponse**](AliasResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrganizationAliasV1OrganizationsMeAliasesPost

> AliasResponse SetOrganizationAliasV1OrganizationsMeAliasesPost(ctx).AliasRequest(aliasRequest).Execute()

Set Organization Alias



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
	aliasRequest := *openapiclient.NewAliasRequest("Name_example", "Target_example") // AliasRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.SetOrganizationAliasV1OrganizationsMeAliasesPost(context.Background()).AliasRequest(aliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.SetOrganizationAliasV1OrganizationsMeAliasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrganizationAliasV1OrganizationsMeAliasesPost`: AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.SetOrganizationAliasV1OrganizationsMeAliasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizationAliasV1OrganizationsMeAliasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aliasRequest** | [**AliasRequest**](AliasRequest.md) |  | 

### Return type

[**AliasResponse**](AliasResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

