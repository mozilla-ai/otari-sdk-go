# \OrganizationGuardrailDefinitionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition**](OrganizationGuardrailDefinitionsAPI.md#OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition) | **Post** /api/v1/organizations/me/guardrail-definitions | Create Organization Guardrail Definition
[**OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition**](OrganizationGuardrailDefinitionsAPI.md#OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition) | **Delete** /api/v1/organizations/me/guardrail-definitions/{definition_id} | Delete Organization Guardrail Definition
[**OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions**](OrganizationGuardrailDefinitionsAPI.md#OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions) | **Get** /api/v1/organizations/me/guardrail-definitions | List Organization Guardrail Definitions
[**OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition**](OrganizationGuardrailDefinitionsAPI.md#OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition) | **Post** /api/v1/organizations/me/guardrail-definitions/{definition_id}/test | Test Organization Guardrail Definition
[**OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition**](OrganizationGuardrailDefinitionsAPI.md#OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition) | **Patch** /api/v1/organizations/me/guardrail-definitions/{definition_id} | Update Organization Guardrail Definition



## OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition

> OrganizationGuardrailDefinitionPublic OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition(ctx).OrganizationGuardrailDefinitionCreate(organizationGuardrailDefinitionCreate).Execute()

Create Organization Guardrail Definition



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
	organizationGuardrailDefinitionCreate := *openapiclient.NewOrganizationGuardrailDefinitionCreate("GuardrailName_example", "Name_example") // OrganizationGuardrailDefinitionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition(context.Background()).OrganizationGuardrailDefinitionCreate(organizationGuardrailDefinitionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition`: OrganizationGuardrailDefinitionPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailDefinitionsCreateOrganizationGuardrailDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationGuardrailDefinitionCreate** | [**OrganizationGuardrailDefinitionCreate**](OrganizationGuardrailDefinitionCreate.md) |  | 

### Return type

[**OrganizationGuardrailDefinitionPublic**](OrganizationGuardrailDefinitionPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition

> Message OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition(ctx, definitionId).Execute()

Delete Organization Guardrail Definition



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
	definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition(context.Background(), definitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailDefinitionsDeleteOrganizationGuardrailDefinitionRequest struct via the builder pattern


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


## OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions

> OrganizationGuardrailDefinitionsPublic OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions(ctx).Skip(skip).Limit(limit).Execute()

List Organization Guardrail Definitions



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
	resp, r, err := apiClient.OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions`: OrganizationGuardrailDefinitionsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailDefinitionsListOrganizationGuardrailDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrganizationGuardrailDefinitionsPublic**](OrganizationGuardrailDefinitionsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition

> OrganizationGuardrailDefinitionTestResult OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition(ctx, definitionId).OrganizationGuardrailDefinitionTest(organizationGuardrailDefinitionTest).Execute()

Test Organization Guardrail Definition



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
	definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationGuardrailDefinitionTest := *openapiclient.NewOrganizationGuardrailDefinitionTest("Text_example") // OrganizationGuardrailDefinitionTest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition(context.Background(), definitionId).OrganizationGuardrailDefinitionTest(organizationGuardrailDefinitionTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition`: OrganizationGuardrailDefinitionTestResult
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailDefinitionsTestOrganizationGuardrailDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationGuardrailDefinitionTest** | [**OrganizationGuardrailDefinitionTest**](OrganizationGuardrailDefinitionTest.md) |  | 

### Return type

[**OrganizationGuardrailDefinitionTestResult**](OrganizationGuardrailDefinitionTestResult.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition

> OrganizationGuardrailDefinitionPublic OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition(ctx, definitionId).OrganizationGuardrailDefinitionUpdate(organizationGuardrailDefinitionUpdate).Execute()

Update Organization Guardrail Definition



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
	definitionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationGuardrailDefinitionUpdate := *openapiclient.NewOrganizationGuardrailDefinitionUpdate() // OrganizationGuardrailDefinitionUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition(context.Background(), definitionId).OrganizationGuardrailDefinitionUpdate(organizationGuardrailDefinitionUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition`: OrganizationGuardrailDefinitionPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationGuardrailDefinitionsAPI.OrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGuardrailDefinitionsUpdateOrganizationGuardrailDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationGuardrailDefinitionUpdate** | [**OrganizationGuardrailDefinitionUpdate**](OrganizationGuardrailDefinitionUpdate.md) |  | 

### Return type

[**OrganizationGuardrailDefinitionPublic**](OrganizationGuardrailDefinitionPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

