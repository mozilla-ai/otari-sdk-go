# \AgentRecommendersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRecommendersCreateAgentRecommender**](AgentRecommendersAPI.md#AgentRecommendersCreateAgentRecommender) | **Post** /api/v1/routing/recommenders | Create Agent Recommender
[**AgentRecommendersDeleteAgentRecommender**](AgentRecommendersAPI.md#AgentRecommendersDeleteAgentRecommender) | **Delete** /api/v1/routing/recommenders/{recommender_id} | Delete Agent Recommender
[**AgentRecommendersListAgentRecommenders**](AgentRecommendersAPI.md#AgentRecommendersListAgentRecommenders) | **Get** /api/v1/routing/recommenders | List Agent Recommenders
[**AgentRecommendersUpdateAgentRecommender**](AgentRecommendersAPI.md#AgentRecommendersUpdateAgentRecommender) | **Patch** /api/v1/routing/recommenders/{recommender_id} | Update Agent Recommender



## AgentRecommendersCreateAgentRecommender

> AgentRecommenderPublic AgentRecommendersCreateAgentRecommender(ctx).AgentRecommenderCreate(agentRecommenderCreate).Execute()

Create Agent Recommender



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
	agentRecommenderCreate := *openapiclient.NewAgentRecommenderCreate("Model_example", "Provider_example") // AgentRecommenderCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRecommendersAPI.AgentRecommendersCreateAgentRecommender(context.Background()).AgentRecommenderCreate(agentRecommenderCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRecommendersAPI.AgentRecommendersCreateAgentRecommender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRecommendersCreateAgentRecommender`: AgentRecommenderPublic
	fmt.Fprintf(os.Stdout, "Response from `AgentRecommendersAPI.AgentRecommendersCreateAgentRecommender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentRecommendersCreateAgentRecommenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentRecommenderCreate** | [**AgentRecommenderCreate**](AgentRecommenderCreate.md) |  | 

### Return type

[**AgentRecommenderPublic**](AgentRecommenderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRecommendersDeleteAgentRecommender

> AgentRecommendersDeleteAgentRecommender(ctx, recommenderId).Execute()

Delete Agent Recommender



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
	recommenderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentRecommendersAPI.AgentRecommendersDeleteAgentRecommender(context.Background(), recommenderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRecommendersAPI.AgentRecommendersDeleteAgentRecommender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recommenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRecommendersDeleteAgentRecommenderRequest struct via the builder pattern


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


## AgentRecommendersListAgentRecommenders

> AgentRecommendersPublic AgentRecommendersListAgentRecommenders(ctx).Execute()

List Agent Recommenders



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
	resp, r, err := apiClient.AgentRecommendersAPI.AgentRecommendersListAgentRecommenders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRecommendersAPI.AgentRecommendersListAgentRecommenders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRecommendersListAgentRecommenders`: AgentRecommendersPublic
	fmt.Fprintf(os.Stdout, "Response from `AgentRecommendersAPI.AgentRecommendersListAgentRecommenders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRecommendersListAgentRecommendersRequest struct via the builder pattern


### Return type

[**AgentRecommendersPublic**](AgentRecommendersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRecommendersUpdateAgentRecommender

> AgentRecommenderPublic AgentRecommendersUpdateAgentRecommender(ctx, recommenderId).AgentRecommenderUpdate(agentRecommenderUpdate).Execute()

Update Agent Recommender



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
	recommenderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	agentRecommenderUpdate := *openapiclient.NewAgentRecommenderUpdate() // AgentRecommenderUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRecommendersAPI.AgentRecommendersUpdateAgentRecommender(context.Background(), recommenderId).AgentRecommenderUpdate(agentRecommenderUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRecommendersAPI.AgentRecommendersUpdateAgentRecommender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRecommendersUpdateAgentRecommender`: AgentRecommenderPublic
	fmt.Fprintf(os.Stdout, "Response from `AgentRecommendersAPI.AgentRecommendersUpdateAgentRecommender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recommenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRecommendersUpdateAgentRecommenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentRecommenderUpdate** | [**AgentRecommenderUpdate**](AgentRecommenderUpdate.md) |  | 

### Return type

[**AgentRecommenderPublic**](AgentRecommenderPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

