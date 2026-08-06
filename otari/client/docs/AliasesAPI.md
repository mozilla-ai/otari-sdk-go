# \AliasesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAliasV1AliasesNameDelete**](AliasesAPI.md#DeleteAliasV1AliasesNameDelete) | **Delete** /v1/aliases/{name} | Delete Alias
[**ListAliasesV1AliasesGet**](AliasesAPI.md#ListAliasesV1AliasesGet) | **Get** /v1/aliases | List Aliases
[**SetAliasV1AliasesPost**](AliasesAPI.md#SetAliasV1AliasesPost) | **Post** /v1/aliases | Set Alias



## DeleteAliasV1AliasesNameDelete

> DeleteAliasV1AliasesNameDelete(ctx, name).UserId(userId).Execute()

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
	userId := "userId_example" // string | Delete the alias scoped to this user. Omit to delete the global alias of that name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AliasesAPI.DeleteAliasV1AliasesNameDelete(context.Background(), name).UserId(userId).Execute()
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

 **userId** | **string** | Delete the alias scoped to this user. Omit to delete the global alias of that name. | 

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

> []AliasResponse ListAliasesV1AliasesGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasesAPI.ListAliasesV1AliasesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasesAPI.ListAliasesV1AliasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAliasesV1AliasesGet`: []AliasResponse
	fmt.Fprintf(os.Stdout, "Response from `AliasesAPI.ListAliasesV1AliasesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAliasesV1AliasesGetRequest struct via the builder pattern


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

