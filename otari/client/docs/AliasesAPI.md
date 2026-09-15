# \AliasesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasesDeleteAlias**](AliasesAPI.md#AliasesDeleteAlias) | **Delete** /api/v1/aliases/{name} | Delete Alias
[**AliasesDeleteOrganizationAlias**](AliasesAPI.md#AliasesDeleteOrganizationAlias) | **Delete** /api/v1/organizations/me/aliases/{name} | Delete Organization Alias
[**AliasesListAliases**](AliasesAPI.md#AliasesListAliases) | **Get** /api/v1/aliases | List Aliases
[**AliasesListVisibleAliases**](AliasesAPI.md#AliasesListVisibleAliases) | **Get** /api/v1/organizations/me/aliases | List Visible Aliases
[**AliasesSetAlias**](AliasesAPI.md#AliasesSetAlias) | **Post** /api/v1/aliases | Set Alias
[**AliasesSetOrganizationAlias**](AliasesAPI.md#AliasesSetOrganizationAlias) | **Post** /api/v1/organizations/me/aliases | Set Organization Alias



## AliasesDeleteAlias

> AliasesDeleteAlias(ctx, name).UserId(userId).WorkspaceId(workspaceId).Execute()

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
	r, err := apiClient.AliasesAPI.AliasesDeleteAlias(context.Background(), name).UserId(userId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesDeleteAlias``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAliasesDeleteAliasRequest struct via the builder pattern


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


## AliasesDeleteOrganizationAlias

> AliasesDeleteOrganizationAlias(ctx, name).WorkspaceId(workspaceId).Execute()

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
	r, err := apiClient.AliasesAPI.AliasesDeleteOrganizationAlias(context.Background(), name).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesDeleteOrganizationAlias``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAliasesDeleteOrganizationAliasRequest struct via the builder pattern


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


## AliasesListAliases

> []AliasResponse AliasesListAliases(ctx).WorkspaceId(workspaceId).Execute()

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
	resp, r, err := apiClient.AliasesAPI.AliasesListAliases(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesListAliases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasesListAliases`: []AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.AliasesListAliases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAliasesListAliasesRequest struct via the builder pattern


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


## AliasesListVisibleAliases

> []AliasResponse AliasesListVisibleAliases(ctx).Limit(limit).WorkspaceId(workspaceId).Execute()

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit for every workspace this caller may see. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.AliasesListVisibleAliases(context.Background()).Limit(limit).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesListVisibleAliases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasesListVisibleAliases`: []AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.AliasesListVisibleAliases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAliasesListVisibleAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return, stored and config-file together. | [default to 1000]
 **workspaceId** | **string** | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit for every workspace this caller may see. | 

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


## AliasesSetAlias

> AliasResponse AliasesSetAlias(ctx).AliasRequest(aliasRequest).Execute()

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
	resp, r, err := apiClient.AliasesAPI.AliasesSetAlias(context.Background()).AliasRequest(aliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesSetAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasesSetAlias`: AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.AliasesSetAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAliasesSetAliasRequest struct via the builder pattern


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


## AliasesSetOrganizationAlias

> AliasResponse AliasesSetOrganizationAlias(ctx).AliasRequest(aliasRequest).Execute()

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
	resp, r, err := apiClient.AliasesAPI.AliasesSetOrganizationAlias(context.Background()).AliasRequest(aliasRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.AliasesSetOrganizationAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasesSetOrganizationAlias`: AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.AliasesSetOrganizationAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAliasesSetOrganizationAliasRequest struct via the builder pattern


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

