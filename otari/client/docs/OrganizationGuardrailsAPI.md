# \OrganizationGuardrailsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGuardrailsCreateOrganizationGuardrail**](OrganizationGuardrailsAPI.md#OrganizationGuardrailsCreateOrganizationGuardrail) | **Post** /api/v1/organizations/me/guardrails | Create Organization Guardrail
[**OrganizationGuardrailsDeleteOrganizationGuardrail**](OrganizationGuardrailsAPI.md#OrganizationGuardrailsDeleteOrganizationGuardrail) | **Delete** /api/v1/organizations/me/guardrails/{guardrail_id} | Delete Organization Guardrail
[**OrganizationGuardrailsListOrganizationGuardrails**](OrganizationGuardrailsAPI.md#OrganizationGuardrailsListOrganizationGuardrails) | **Get** /api/v1/organizations/me/guardrails | List Organization Guardrails
[**OrganizationGuardrailsUpdateOrganizationGuardrail**](OrganizationGuardrailsAPI.md#OrganizationGuardrailsUpdateOrganizationGuardrail) | **Patch** /api/v1/organizations/me/guardrails/{guardrail_id} | Update Organization Guardrail



## OrganizationGuardrailsCreateOrganizationGuardrail

> OrganizationGuardrailPublic OrganizationGuardrailsCreateOrganizationGuardrail(ctx).OrganizationGuardrailCreate(organizationGuardrailCreate).Execute()

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
	resp, r, err := apiClient.OrganizationGuardrailsAPI.OrganizationGuardrailsCreateOrganizationGuardrail(context.Background()).OrganizationGuardrailCreate(organizationGuardrailCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.OrganizationGuardrailsCreateOrganizationGuardrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailsCreateOrganizationGuardrail`: OrganizationGuardrailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.OrganizationGuardrailsCreateOrganizationGuardrail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailsCreateOrganizationGuardrailRequest struct via the builder pattern


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


## OrganizationGuardrailsDeleteOrganizationGuardrail

> Message OrganizationGuardrailsDeleteOrganizationGuardrail(ctx, guardrailId).Execute()

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
	resp, r, err := apiClient.OrganizationGuardrailsAPI.OrganizationGuardrailsDeleteOrganizationGuardrail(context.Background(), guardrailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.OrganizationGuardrailsDeleteOrganizationGuardrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailsDeleteOrganizationGuardrail`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.OrganizationGuardrailsDeleteOrganizationGuardrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guardrailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailsDeleteOrganizationGuardrailRequest struct via the builder pattern


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


## OrganizationGuardrailsListOrganizationGuardrails

> OrganizationGuardrailsPublic OrganizationGuardrailsListOrganizationGuardrails(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationGuardrailsAPI.OrganizationGuardrailsListOrganizationGuardrails(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.OrganizationGuardrailsListOrganizationGuardrails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailsListOrganizationGuardrails`: OrganizationGuardrailsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.OrganizationGuardrailsListOrganizationGuardrails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailsListOrganizationGuardrailsRequest struct via the builder pattern


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


## OrganizationGuardrailsUpdateOrganizationGuardrail

> OrganizationGuardrailPublic OrganizationGuardrailsUpdateOrganizationGuardrail(ctx, guardrailId).OrganizationGuardrailUpdate(organizationGuardrailUpdate).Execute()

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
	resp, r, err := apiClient.OrganizationGuardrailsAPI.OrganizationGuardrailsUpdateOrganizationGuardrail(context.Background(), guardrailId).OrganizationGuardrailUpdate(organizationGuardrailUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailsAPI.OrganizationGuardrailsUpdateOrganizationGuardrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailsUpdateOrganizationGuardrail`: OrganizationGuardrailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailsAPI.OrganizationGuardrailsUpdateOrganizationGuardrail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guardrailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailsUpdateOrganizationGuardrailRequest struct via the builder pattern


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

