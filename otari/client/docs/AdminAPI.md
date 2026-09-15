# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGetAdministrationAccess**](AdminAPI.md#AdminGetAdministrationAccess) | **Get** /api/v1/admin/access | Get Administration Access
[**AdminListDeploymentUsers**](AdminAPI.md#AdminListDeploymentUsers) | **Get** /api/v1/admin/users | List Deployment Users
[**AdminUpdateDeploymentUser**](AdminAPI.md#AdminUpdateDeploymentUser) | **Patch** /api/v1/admin/users/{user_id} | Update Deployment User



## AdminGetAdministrationAccess

> DeploymentAdminAccessPublic AdminGetAdministrationAccess(ctx).Execute()

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
	resp, r, err := apiClient.AdminAPI.AdminGetAdministrationAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminGetAdministrationAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminGetAdministrationAccess`: DeploymentAdminAccessPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminGetAdministrationAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetAdministrationAccessRequest struct via the builder pattern


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


## AdminListDeploymentUsers

> DeploymentUsersPublic AdminListDeploymentUsers(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.AdminAPI.AdminListDeploymentUsers(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminListDeploymentUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminListDeploymentUsers`: DeploymentUsersPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminListDeploymentUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminListDeploymentUsersRequest struct via the builder pattern


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


## AdminUpdateDeploymentUser

> DeploymentUserPublic AdminUpdateDeploymentUser(ctx, userId).DeploymentUserUpdateRequest(deploymentUserUpdateRequest).Execute()

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
	resp, r, err := apiClient.AdminAPI.AdminUpdateDeploymentUser(context.Background(), userId).DeploymentUserUpdateRequest(deploymentUserUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminUpdateDeploymentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminUpdateDeploymentUser`: DeploymentUserPublic
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminUpdateDeploymentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateDeploymentUserRequest struct via the builder pattern


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

