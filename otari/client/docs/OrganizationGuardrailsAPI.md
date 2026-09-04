# \OrganizationGuardrailsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost**](OrganizationGuardrailsAPI.md#CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost) | **Post** /v1/organizations/me/guardrails | Create Organization Guardrail
[**DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete**](OrganizationGuardrailsAPI.md#DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete) | **Delete** /v1/organizations/me/guardrails/{guardrail_id} | Delete Organization Guardrail
[**ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet**](OrganizationGuardrailsAPI.md#ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet) | **Get** /v1/organizations/me/guardrails | List Organization Guardrails
[**UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch**](OrganizationGuardrailsAPI.md#UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch) | **Patch** /v1/organizations/me/guardrails/{guardrail_id} | Update Organization Guardrail



## CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost

> OrganizationGuardrailPublic CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost(ctx).OrganizationGuardrailCreate(organizationGuardrailCreate).Execute()

Create Organization Guardrail



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
	organizationGuardrailCreate := *openapiclient.NewOrganizationGuardrailCreate("Profile_example") // OrganizationGuardrailCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailsAPI.CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost(context.Background()).OrganizationGuardrailCreate(organizationGuardrailCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost`: OrganizationGuardrailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.CreateOrganizationGuardrailV1OrganizationsMeGuardrailsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationGuardrailV1OrganizationsMeGuardrailsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationGuardrailCreate** | [**OrganizationGuardrailCreate**](OrganizationGuardrailCreate.md) |  | 

### Return type

[**OrganizationGuardrailPublic**](OrganizationGuardrailPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete

> Message DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete(ctx, guardrailId).Execute()

Delete Organization Guardrail



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
	guardrailId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailsAPI.DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete(context.Background(), guardrailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.DeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guardrailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdDeleteRequest struct via the builder pattern


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


## ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet

> OrganizationGuardrailsPublic ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet(ctx).Skip(skip).Limit(limit).Execute()

List Organization Guardrails



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
	resp, r, err := apiClient.OrganizationGuardrailsAPI.ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet`: OrganizationGuardrailsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.ListOrganizationGuardrailsV1OrganizationsMeGuardrailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationGuardrailsV1OrganizationsMeGuardrailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrganizationGuardrailsPublic**](OrganizationGuardrailsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch

> OrganizationGuardrailPublic UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch(ctx, guardrailId).OrganizationGuardrailUpdate(organizationGuardrailUpdate).Execute()

Update Organization Guardrail



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
	guardrailId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationGuardrailUpdate := *openapiclient.NewOrganizationGuardrailUpdate() // OrganizationGuardrailUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailsAPI.UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch(context.Background(), guardrailId).OrganizationGuardrailUpdate(organizationGuardrailUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch`: OrganizationGuardrailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.UpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guardrailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationGuardrailV1OrganizationsMeGuardrailsGuardrailIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationGuardrailUpdate** | [**OrganizationGuardrailUpdate**](OrganizationGuardrailUpdate.md) |  | 

### Return type

[**OrganizationGuardrailPublic**](OrganizationGuardrailPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

