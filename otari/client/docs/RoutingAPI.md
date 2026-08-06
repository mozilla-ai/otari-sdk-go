# \RoutingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicyV1RoutingPoliciesNameDelete**](RoutingAPI.md#DeletePolicyV1RoutingPoliciesNameDelete) | **Delete** /v1/routing/policies/{name} | Delete Policy
[**ExplainPolicyV1RoutingPoliciesExplainPost**](RoutingAPI.md#ExplainPolicyV1RoutingPoliciesExplainPost) | **Post** /v1/routing/policies/explain | Explain Policy
[**ListPoliciesV1RoutingPoliciesGet**](RoutingAPI.md#ListPoliciesV1RoutingPoliciesGet) | **Get** /v1/routing/policies | List Policies
[**SetPolicyV1RoutingPoliciesPost**](RoutingAPI.md#SetPolicyV1RoutingPoliciesPost) | **Post** /v1/routing/policies | Set Policy



## DeletePolicyV1RoutingPoliciesNameDelete

> DeletePolicyV1RoutingPoliciesNameDelete(ctx, name).UserId(userId).Execute()

Delete Policy



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
	userId := "userId_example" // string | Delete the policy scoped to this user. Omit to delete the global one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeletePolicyV1RoutingPoliciesNameDelete(context.Background(), name).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeletePolicyV1RoutingPoliciesNameDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePolicyV1RoutingPoliciesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Delete the policy scoped to this user. Omit to delete the global one. | 

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


## ExplainPolicyV1RoutingPoliciesExplainPost

> ExplainResponse ExplainPolicyV1RoutingPoliciesExplainPost(ctx).ExplainRequest(explainRequest).Execute()

Explain Policy



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
	explainRequest := *openapiclient.NewExplainRequest() // ExplainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ExplainPolicyV1RoutingPoliciesExplainPost(context.Background()).ExplainRequest(explainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ExplainPolicyV1RoutingPoliciesExplainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainPolicyV1RoutingPoliciesExplainPost`: ExplainResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ExplainPolicyV1RoutingPoliciesExplainPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExplainPolicyV1RoutingPoliciesExplainPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **explainRequest** | [**ExplainRequest**](ExplainRequest.md) |  | 

### Return type

[**ExplainResponse**](ExplainResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPoliciesV1RoutingPoliciesGet

> []PolicyResponse ListPoliciesV1RoutingPoliciesGet(ctx).Execute()

List Policies



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
	resp, r, err := apiClient.RoutingAPI.ListPoliciesV1RoutingPoliciesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ListPoliciesV1RoutingPoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPoliciesV1RoutingPoliciesGet`: []PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ListPoliciesV1RoutingPoliciesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesV1RoutingPoliciesGetRequest struct via the builder pattern


### Return type

[**[]PolicyResponse**](PolicyResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicyV1RoutingPoliciesPost

> PolicyResponse SetPolicyV1RoutingPoliciesPost(ctx).PolicyRequest(policyRequest).Execute()

Set Policy



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
	policyRequest := *openapiclient.NewPolicyRequest("Name_example", map[string]interface{}{"key": interface{}(123)}) // PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.SetPolicyV1RoutingPoliciesPost(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.SetPolicyV1RoutingPoliciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPolicyV1RoutingPoliciesPost`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.SetPolicyV1RoutingPoliciesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPolicyV1RoutingPoliciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyRequest** | [**PolicyRequest**](PolicyRequest.md) |  | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

