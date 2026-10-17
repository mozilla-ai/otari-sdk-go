# \WebSearchKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebSearchKeysArchiveOrgWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysArchiveOrgWebSearchKey) | **Post** /api/v1/organizations/me/web-search-keys/{key_id}/archive | Archive Org Web Search Key
[**WebSearchKeysCreateOrgWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysCreateOrgWebSearchKey) | **Post** /api/v1/organizations/me/web-search-keys | Create Org Web Search Key
[**WebSearchKeysDeleteOrgWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysDeleteOrgWebSearchKey) | **Delete** /api/v1/organizations/me/web-search-keys/{key_id} | Delete Org Web Search Key
[**WebSearchKeysListOrgWebSearchKeys**](WebSearchKeysAPI.md#WebSearchKeysListOrgWebSearchKeys) | **Get** /api/v1/organizations/me/web-search-keys | List Org Web Search Keys
[**WebSearchKeysListWorkspaceWebSearchKeys**](WebSearchKeysAPI.md#WebSearchKeysListWorkspaceWebSearchKeys) | **Get** /api/v1/workspaces/{workspace_id}/web-search-keys | List Workspace Web Search Keys
[**WebSearchKeysResetWorkspaceWebSearchKeyOverride**](WebSearchKeysAPI.md#WebSearchKeysResetWorkspaceWebSearchKeyOverride) | **Delete** /api/v1/workspaces/{workspace_id}/web-search-keys/{key_id} | Reset Workspace Web Search Key Override
[**WebSearchKeysRestoreOrgWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysRestoreOrgWebSearchKey) | **Post** /api/v1/organizations/me/web-search-keys/{key_id}/restore | Restore Org Web Search Key
[**WebSearchKeysSetOrgDefaultWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysSetOrgDefaultWebSearchKey) | **Post** /api/v1/organizations/me/web-search-keys/{key_id}/default | Set Org Default Web Search Key
[**WebSearchKeysSetWorkspaceWebSearchKeyOverride**](WebSearchKeysAPI.md#WebSearchKeysSetWorkspaceWebSearchKeyOverride) | **Patch** /api/v1/workspaces/{workspace_id}/web-search-keys/{key_id} | Set Workspace Web Search Key Override
[**WebSearchKeysUpdateOrgWebSearchKey**](WebSearchKeysAPI.md#WebSearchKeysUpdateOrgWebSearchKey) | **Patch** /api/v1/organizations/me/web-search-keys/{key_id} | Update Org Web Search Key



## WebSearchKeysArchiveOrgWebSearchKey

> OrgWebSearchKeyPublic WebSearchKeysArchiveOrgWebSearchKey(ctx, keyId).Execute()

Archive Org Web Search Key



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
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysArchiveOrgWebSearchKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysArchiveOrgWebSearchKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysArchiveOrgWebSearchKey`: OrgWebSearchKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysArchiveOrgWebSearchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysArchiveOrgWebSearchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgWebSearchKeyPublic**](OrgWebSearchKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysCreateOrgWebSearchKey

> OrgWebSearchKeyPublic WebSearchKeysCreateOrgWebSearchKey(ctx).OrgWebSearchKeyCreateRequest(orgWebSearchKeyCreateRequest).Execute()

Create Org Web Search Key



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
	orgWebSearchKeyCreateRequest := *openapiclient.NewOrgWebSearchKeyCreateRequest("ApiKey_example", "Name_example", "Provider_example") // OrgWebSearchKeyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysCreateOrgWebSearchKey(context.Background()).OrgWebSearchKeyCreateRequest(orgWebSearchKeyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysCreateOrgWebSearchKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysCreateOrgWebSearchKey`: OrgWebSearchKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysCreateOrgWebSearchKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysCreateOrgWebSearchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgWebSearchKeyCreateRequest** | [**OrgWebSearchKeyCreateRequest**](OrgWebSearchKeyCreateRequest.md) |  | 

### Return type

[**OrgWebSearchKeyPublic**](OrgWebSearchKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysDeleteOrgWebSearchKey

> WebSearchKeysDeleteOrgWebSearchKey(ctx, keyId).Execute()

Delete Org Web Search Key



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
	r, err := apiClient.WebSearchKeysAPI.WebSearchKeysDeleteOrgWebSearchKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysDeleteOrgWebSearchKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWebSearchKeysDeleteOrgWebSearchKeyRequest struct via the builder pattern


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


## WebSearchKeysListOrgWebSearchKeys

> OrgWebSearchKeysPublic WebSearchKeysListOrgWebSearchKeys(ctx).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()

List Org Web Search Keys



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
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysListOrgWebSearchKeys(context.Background()).IncludeArchived(includeArchived).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysListOrgWebSearchKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysListOrgWebSearchKeys`: OrgWebSearchKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysListOrgWebSearchKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysListOrgWebSearchKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArchived** | **bool** | Include archived keys. | [default to false]
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrgWebSearchKeysPublic**](OrgWebSearchKeysPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysListWorkspaceWebSearchKeys

> WorkspaceWebSearchKeysPublic WebSearchKeysListWorkspaceWebSearchKeys(ctx, workspaceId).Skip(skip).Limit(limit).Execute()

List Workspace Web Search Keys



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysListWorkspaceWebSearchKeys(context.Background(), workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysListWorkspaceWebSearchKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysListWorkspaceWebSearchKeys`: WorkspaceWebSearchKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysListWorkspaceWebSearchKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysListWorkspaceWebSearchKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**WorkspaceWebSearchKeysPublic**](WorkspaceWebSearchKeysPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysResetWorkspaceWebSearchKeyOverride

> WorkspaceWebSearchKeysPublic WebSearchKeysResetWorkspaceWebSearchKeyOverride(ctx, workspaceId, keyId).Execute()

Reset Workspace Web Search Key Override



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
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysResetWorkspaceWebSearchKeyOverride(context.Background(), workspaceId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysResetWorkspaceWebSearchKeyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysResetWorkspaceWebSearchKeyOverride`: WorkspaceWebSearchKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysResetWorkspaceWebSearchKeyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysResetWorkspaceWebSearchKeyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkspaceWebSearchKeysPublic**](WorkspaceWebSearchKeysPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysRestoreOrgWebSearchKey

> OrgWebSearchKeyPublic WebSearchKeysRestoreOrgWebSearchKey(ctx, keyId).Execute()

Restore Org Web Search Key



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
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysRestoreOrgWebSearchKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysRestoreOrgWebSearchKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysRestoreOrgWebSearchKey`: OrgWebSearchKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysRestoreOrgWebSearchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysRestoreOrgWebSearchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgWebSearchKeyPublic**](OrgWebSearchKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysSetOrgDefaultWebSearchKey

> OrgWebSearchKeyPublic WebSearchKeysSetOrgDefaultWebSearchKey(ctx, keyId).Execute()

Set Org Default Web Search Key



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
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysSetOrgDefaultWebSearchKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysSetOrgDefaultWebSearchKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysSetOrgDefaultWebSearchKey`: OrgWebSearchKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysSetOrgDefaultWebSearchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysSetOrgDefaultWebSearchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgWebSearchKeyPublic**](OrgWebSearchKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysSetWorkspaceWebSearchKeyOverride

> WorkspaceWebSearchKeysPublic WebSearchKeysSetWorkspaceWebSearchKeyOverride(ctx, workspaceId, keyId).WorkspaceWebSearchKeyOverrideRequest(workspaceWebSearchKeyOverrideRequest).Execute()

Set Workspace Web Search Key Override



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
	workspaceWebSearchKeyOverrideRequest := *openapiclient.NewWorkspaceWebSearchKeyOverrideRequest() // WorkspaceWebSearchKeyOverrideRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysSetWorkspaceWebSearchKeyOverride(context.Background(), workspaceId, keyId).WorkspaceWebSearchKeyOverrideRequest(workspaceWebSearchKeyOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysSetWorkspaceWebSearchKeyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysSetWorkspaceWebSearchKeyOverride`: WorkspaceWebSearchKeysPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysSetWorkspaceWebSearchKeyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysSetWorkspaceWebSearchKeyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceWebSearchKeyOverrideRequest** | [**WorkspaceWebSearchKeyOverrideRequest**](WorkspaceWebSearchKeyOverrideRequest.md) |  | 

### Return type

[**WorkspaceWebSearchKeysPublic**](WorkspaceWebSearchKeysPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebSearchKeysUpdateOrgWebSearchKey

> OrgWebSearchKeyPublic WebSearchKeysUpdateOrgWebSearchKey(ctx, keyId).OrgWebSearchKeyUpdateRequest(orgWebSearchKeyUpdateRequest).Execute()

Update Org Web Search Key



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
	orgWebSearchKeyUpdateRequest := *openapiclient.NewOrgWebSearchKeyUpdateRequest() // OrgWebSearchKeyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSearchKeysAPI.WebSearchKeysUpdateOrgWebSearchKey(context.Background(), keyId).OrgWebSearchKeyUpdateRequest(orgWebSearchKeyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSearchKeysAPI.WebSearchKeysUpdateOrgWebSearchKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebSearchKeysUpdateOrgWebSearchKey`: OrgWebSearchKeyPublic
	fmt.Fprintf(os.Stdout, "Response from `WebSearchKeysAPI.WebSearchKeysUpdateOrgWebSearchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebSearchKeysUpdateOrgWebSearchKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgWebSearchKeyUpdateRequest** | [**OrgWebSearchKeyUpdateRequest**](OrgWebSearchKeyUpdateRequest.md) |  | 

### Return type

[**OrgWebSearchKeyPublic**](OrgWebSearchKeyPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

