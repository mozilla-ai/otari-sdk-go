# \RoutingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete**](RoutingAPI.md#DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete) | **Delete** /v1/organizations/me/routing-policies/{name} | Delete Organization Routing Policy
[**DeletePolicyV1RoutingPoliciesNameDelete**](RoutingAPI.md#DeletePolicyV1RoutingPoliciesNameDelete) | **Delete** /v1/routing/policies/{name} | Delete Policy
[**ExplainPolicyV1RoutingPoliciesExplainPost**](RoutingAPI.md#ExplainPolicyV1RoutingPoliciesExplainPost) | **Post** /v1/routing/policies/explain | Explain Policy
[**ListPoliciesV1RoutingPoliciesGet**](RoutingAPI.md#ListPoliciesV1RoutingPoliciesGet) | **Get** /v1/routing/policies | List Policies
[**ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet**](RoutingAPI.md#ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet) | **Get** /v1/organizations/me/routing-policies | List Visible Routing Policies
[**RankCandidatesV1RoutingPreferencesRankPost**](RoutingAPI.md#RankCandidatesV1RoutingPreferencesRankPost) | **Post** /v1/routing/preferences/rank | Rank Candidates
[**RoutingMemoryStatusV1RoutingStatusGet**](RoutingAPI.md#RoutingMemoryStatusV1RoutingStatusGet) | **Get** /v1/routing/status | Routing Memory Status
[**SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost**](RoutingAPI.md#SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost) | **Post** /v1/organizations/me/routing-policies | Set Organization Routing Policy
[**SetPolicyV1RoutingPoliciesPost**](RoutingAPI.md#SetPolicyV1RoutingPoliciesPost) | **Post** /v1/routing/policies | Set Policy



## DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete

> DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete(ctx, name).WorkspaceId(workspaceId).Execute()

Delete Organization Routing Policy



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Delete the policy in this workspace of the caller's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete(context.Background(), name).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.DeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | Delete the policy in this workspace of the caller&#39;s organization. | 

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


## DeletePolicyV1RoutingPoliciesNameDelete

> DeletePolicyV1RoutingPoliciesNameDelete(ctx, name).UserId(userId).WorkspaceId(workspaceId).Execute()

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
	userId := "userId_example" // string | Delete the policy scoped to this user. Omit to delete the workspace-wide one. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Delete the policy in this workspace. Omit for the deployment's default workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutingAPI.DeletePolicyV1RoutingPoliciesNameDelete(context.Background(), name).UserId(userId).WorkspaceId(workspaceId).Execute()
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

 **userId** | **string** | Delete the policy scoped to this user. Omit to delete the workspace-wide one. | 
 **workspaceId** | **string** | Delete the policy in this workspace. Omit for the deployment&#39;s default workspace. | 

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

> []PolicyResponse ListPoliciesV1RoutingPoliciesGet(ctx).WorkspaceId(workspaceId).Execute()

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only stored policies in this workspace. Config-file policies are always included, being deployment-wide. Omit to list the stored policies of every workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ListPoliciesV1RoutingPoliciesGet(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ListPoliciesV1RoutingPoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPoliciesV1RoutingPoliciesGet`: []PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ListPoliciesV1RoutingPoliciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesV1RoutingPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Only stored policies in this workspace. Config-file policies are always included, being deployment-wide. Omit to list the stored policies of every workspace. | 

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


## ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet

> []PolicyResponse ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet(ctx).Limit(limit).Execute()

List Visible Routing Policies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet`: []PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.ListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVisibleRoutingPoliciesV1OrganizationsMeRoutingPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return, stored and config-file together. | [default to 1000]

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


## RankCandidatesV1RoutingPreferencesRankPost

> RankResponse RankCandidatesV1RoutingPreferencesRankPost(ctx).RankRequest(rankRequest).Execute()

Rank Candidates



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
	rankRequest := *openapiclient.NewRankRequest([]openapiclient.ScoredExample{*openapiclient.NewScoredExample("Prompt_example", map[string]float32{"key": float32(123)})}, "UserId_example") // RankRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.RankCandidatesV1RoutingPreferencesRankPost(context.Background()).RankRequest(rankRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RankCandidatesV1RoutingPreferencesRankPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RankCandidatesV1RoutingPreferencesRankPost`: RankResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RankCandidatesV1RoutingPreferencesRankPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRankCandidatesV1RoutingPreferencesRankPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rankRequest** | [**RankRequest**](RankRequest.md) |  | 

### Return type

[**RankResponse**](RankResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoutingMemoryStatusV1RoutingStatusGet

> RouterStatus RoutingMemoryStatusV1RoutingStatusGet(ctx).UserId(userId).WorkspaceId(workspaceId).Execute()

Routing Memory Status



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
	userId := "userId_example" // string | Whose routing memory to report on.
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Which workspace's routing memory to report on. Omit for the default workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.RoutingMemoryStatusV1RoutingStatusGet(context.Background()).UserId(userId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingMemoryStatusV1RoutingStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingMemoryStatusV1RoutingStatusGet`: RouterStatus
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingMemoryStatusV1RoutingStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingMemoryStatusV1RoutingStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Whose routing memory to report on. | 
 **workspaceId** | **string** | Which workspace&#39;s routing memory to report on. Omit for the default workspace. | 

### Return type

[**RouterStatus**](RouterStatus.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost

> PolicyResponse SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost(ctx).PolicyRequest(policyRequest).Execute()

Set Organization Routing Policy



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
	resp, r, err := apiClient.RoutingAPI.SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.SetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizationRoutingPolicyV1OrganizationsMeRoutingPoliciesPostRequest struct via the builder pattern


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

