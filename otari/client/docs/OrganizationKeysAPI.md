# \OrganizationKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOwnKeyV1OrganizationsMeKeysPost**](OrganizationKeysAPI.md#CreateOwnKeyV1OrganizationsMeKeysPost) | **Post** /v1/organizations/me/keys | Create Own Key
[**DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete**](OrganizationKeysAPI.md#DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete) | **Delete** /v1/organizations/me/keys/{key_id} | Delete Own Key
[**ListOwnKeysV1OrganizationsMeKeysGet**](OrganizationKeysAPI.md#ListOwnKeysV1OrganizationsMeKeysGet) | **Get** /v1/organizations/me/keys | List Own Keys
[**RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost**](OrganizationKeysAPI.md#RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost) | **Post** /v1/organizations/me/keys/{key_id}/rotate | Rotate Own Key
[**UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch**](OrganizationKeysAPI.md#UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch) | **Patch** /v1/organizations/me/keys/{key_id} | Update Own Key



## CreateOwnKeyV1OrganizationsMeKeysPost

> CreateKeyResponse CreateOwnKeyV1OrganizationsMeKeysPost(ctx).CreateOwnKeyRequest(createOwnKeyRequest).Execute()

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
	resp, r, err := apiClient.OrganizationKeysAPI.CreateOwnKeyV1OrganizationsMeKeysPost(context.Background()).CreateOwnKeyRequest(createOwnKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.CreateOwnKeyV1OrganizationsMeKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOwnKeyV1OrganizationsMeKeysPost`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.CreateOwnKeyV1OrganizationsMeKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOwnKeyV1OrganizationsMeKeysPostRequest struct via the builder pattern


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


## DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete

> DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete(ctx, keyId).Execute()

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
	r, err := apiClient.OrganizationKeysAPI.DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.DeleteOwnKeyV1OrganizationsMeKeysKeyIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOwnKeyV1OrganizationsMeKeysKeyIdDeleteRequest struct via the builder pattern


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


## ListOwnKeysV1OrganizationsMeKeysGet

> []KeyInfo ListOwnKeysV1OrganizationsMeKeysGet(ctx).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()

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
	resp, r, err := apiClient.OrganizationKeysAPI.ListOwnKeysV1OrganizationsMeKeysGet(context.Background()).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.ListOwnKeysV1OrganizationsMeKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOwnKeysV1OrganizationsMeKeysGet`: []KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.ListOwnKeysV1OrganizationsMeKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOwnKeysV1OrganizationsMeKeysGetRequest struct via the builder pattern


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


## RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost

> CreateKeyResponse RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost(ctx, keyId).Execute()

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
	resp, r, err := apiClient.OrganizationKeysAPI.RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost`: CreateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.RotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateOwnKeyV1OrganizationsMeKeysKeyIdRotatePostRequest struct via the builder pattern


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


## UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch

> KeyInfo UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch(ctx, keyId).UpdateOwnKeyRequest(updateOwnKeyRequest).Execute()

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
	resp, r, err := apiClient.OrganizationKeysAPI.UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch(context.Background(), keyId).UpdateOwnKeyRequest(updateOwnKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationKeysAPI.UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch`: KeyInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationKeysAPI.UpdateOwnKeyV1OrganizationsMeKeysKeyIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnKeyV1OrganizationsMeKeysKeyIdPatchRequest struct via the builder pattern


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

