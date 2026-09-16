# \RoutingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoutingDeleteOrganizationRoutingPolicy**](RoutingAPI.md#RoutingDeleteOrganizationRoutingPolicy) | **Delete** /api/v1/organizations/me/routing-policies/{name} | Delete Organization Routing Policy
[**RoutingDeletePolicy**](RoutingAPI.md#RoutingDeletePolicy) | **Delete** /api/v1/routing/policies/{name} | Delete Policy
[**RoutingExplainPolicy**](RoutingAPI.md#RoutingExplainPolicy) | **Post** /api/v1/routing/policies/explain | Explain Policy
[**RoutingListPolicies**](RoutingAPI.md#RoutingListPolicies) | **Get** /api/v1/routing/policies | List Policies
[**RoutingListVisibleRoutingPolicies**](RoutingAPI.md#RoutingListVisibleRoutingPolicies) | **Get** /api/v1/organizations/me/routing-policies | List Visible Routing Policies
[**RoutingRankCandidates**](RoutingAPI.md#RoutingRankCandidates) | **Post** /api/v1/routing/preferences/rank | Rank Candidates
[**RoutingRoutingMemoryStatus**](RoutingAPI.md#RoutingRoutingMemoryStatus) | **Get** /api/v1/routing/status | Routing Memory Status
[**RoutingSetOrganizationRoutingPolicy**](RoutingAPI.md#RoutingSetOrganizationRoutingPolicy) | **Post** /api/v1/organizations/me/routing-policies | Set Organization Routing Policy
[**RoutingSetPolicy**](RoutingAPI.md#RoutingSetPolicy) | **Post** /api/v1/routing/policies | Set Policy



## RoutingDeleteOrganizationRoutingPolicy

> RoutingDeleteOrganizationRoutingPolicy(ctx, name).WorkspaceId(workspaceId).Execute()

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
	r, err := apiClient.RoutingAPI.RoutingDeleteOrganizationRoutingPolicy(context.Background(), name).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingDeleteOrganizationRoutingPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRoutingDeleteOrganizationRoutingPolicyRequest struct via the builder pattern


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


## RoutingDeletePolicy

> RoutingDeletePolicy(ctx, name).UserId(userId).WorkspaceId(workspaceId).Execute()

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
	r, err := apiClient.RoutingAPI.RoutingDeletePolicy(context.Background(), name).UserId(userId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingDeletePolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRoutingDeletePolicyRequest struct via the builder pattern


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


## RoutingExplainPolicy

> ExplainResponse RoutingExplainPolicy(ctx).ExplainRequest(explainRequest).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingExplainPolicy(context.Background()).ExplainRequest(explainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingExplainPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingExplainPolicy`: ExplainResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingExplainPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingExplainPolicyRequest struct via the builder pattern


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


## RoutingListPolicies

> []PolicyResponse RoutingListPolicies(ctx).WorkspaceId(workspaceId).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingListPolicies(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingListPolicies`: []PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingListPoliciesRequest struct via the builder pattern


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


## RoutingListVisibleRoutingPolicies

> []PolicyResponse RoutingListVisibleRoutingPolicies(ctx).Limit(limit).WorkspaceId(workspaceId).Execute()

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit for every workspace this caller may see. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.RoutingListVisibleRoutingPolicies(context.Background()).Limit(limit).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingListVisibleRoutingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingListVisibleRoutingPolicies`: []PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingListVisibleRoutingPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingListVisibleRoutingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return, stored and config-file together. | [default to 1000]
 **workspaceId** | **string** | Only stored entries in this workspace. Config-file entries are always included, being deployment-wide. Omit for every workspace this caller may see. | 

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


## RoutingRankCandidates

> RankResponse RoutingRankCandidates(ctx).RankRequest(rankRequest).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingRankCandidates(context.Background()).RankRequest(rankRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingRankCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingRankCandidates`: RankResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingRankCandidates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingRankCandidatesRequest struct via the builder pattern


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


## RoutingRoutingMemoryStatus

> RouterStatus RoutingRoutingMemoryStatus(ctx).UserId(userId).WorkspaceId(workspaceId).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingRoutingMemoryStatus(context.Background()).UserId(userId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingRoutingMemoryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingRoutingMemoryStatus`: RouterStatus
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingRoutingMemoryStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingRoutingMemoryStatusRequest struct via the builder pattern


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


## RoutingSetOrganizationRoutingPolicy

> PolicyResponse RoutingSetOrganizationRoutingPolicy(ctx).PolicyRequest(policyRequest).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingSetOrganizationRoutingPolicy(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingSetOrganizationRoutingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingSetOrganizationRoutingPolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingSetOrganizationRoutingPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingSetOrganizationRoutingPolicyRequest struct via the builder pattern


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


## RoutingSetPolicy

> PolicyResponse RoutingSetPolicy(ctx).PolicyRequest(policyRequest).Execute()

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
	resp, r, err := apiClient.RoutingAPI.RoutingSetPolicy(context.Background()).PolicyRequest(policyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.RoutingSetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoutingSetPolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.RoutingSetPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoutingSetPolicyRequest struct via the builder pattern


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

