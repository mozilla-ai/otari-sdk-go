# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministrationAccessV1AdminAccessGet**](AdminAPI.md#GetAdministrationAccessV1AdminAccessGet) | **Get** /v1/admin/access | Get Administration Access
[**ListDeploymentUsersV1AdminUsersGet**](AdminAPI.md#ListDeploymentUsersV1AdminUsersGet) | **Get** /v1/admin/users | List Deployment Users
[**UpdateDeploymentUserV1AdminUsersUserIdPatch**](AdminAPI.md#UpdateDeploymentUserV1AdminUsersUserIdPatch) | **Patch** /v1/admin/users/{user_id} | Update Deployment User



## GetAdministrationAccessV1AdminAccessGet

> DeploymentAdminAccessPublic GetAdministrationAccessV1AdminAccessGet(ctx).Execute()

Get Administration Access



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
	resp, r, err := apiClient.AdminAPI.GetAdministrationAccessV1AdminAccessGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdministrationAccessV1AdminAccessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdministrationAccessV1AdminAccessGet`: DeploymentAdminAccessPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdministrationAccessV1AdminAccessGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministrationAccessV1AdminAccessGetRequest struct via the builder pattern


### Return type

[**DeploymentAdminAccessPublic**](DeploymentAdminAccessPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeploymentUsersV1AdminUsersGet

> DeploymentUsersPublic ListDeploymentUsersV1AdminUsersGet(ctx).Skip(skip).Limit(limit).Execute()

List Deployment Users



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
	resp, r, err := apiClient.AdminAPI.ListDeploymentUsersV1AdminUsersGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListDeploymentUsersV1AdminUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeploymentUsersV1AdminUsersGet`: DeploymentUsersPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListDeploymentUsersV1AdminUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentUsersV1AdminUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**DeploymentUsersPublic**](DeploymentUsersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentUserV1AdminUsersUserIdPatch

> DeploymentUserPublic UpdateDeploymentUserV1AdminUsersUserIdPatch(ctx, userId).DeploymentUserUpdateRequest(deploymentUserUpdateRequest).Execute()

Update Deployment User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	deploymentUserUpdateRequest := *openapiclient.NewDeploymentUserUpdateRequest() // DeploymentUserUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.UpdateDeploymentUserV1AdminUsersUserIdPatch(context.Background(), userId).DeploymentUserUpdateRequest(deploymentUserUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateDeploymentUserV1AdminUsersUserIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeploymentUserV1AdminUsersUserIdPatch`: DeploymentUserPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.UpdateDeploymentUserV1AdminUsersUserIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentUserV1AdminUsersUserIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentUserUpdateRequest** | [**DeploymentUserUpdateRequest**](DeploymentUserUpdateRequest.md) |  | 

### Return type

[**DeploymentUserPublic**](DeploymentUserPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

