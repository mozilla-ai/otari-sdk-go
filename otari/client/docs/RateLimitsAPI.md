# \RateLimitsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RateLimitsCreateRateLimitRule**](RateLimitsAPI.md#RateLimitsCreateRateLimitRule) | **Post** /api/v1/rate-limits | Create Rate Limit Rule
[**RateLimitsDeleteRateLimitRule**](RateLimitsAPI.md#RateLimitsDeleteRateLimitRule) | **Delete** /api/v1/rate-limits/{name} | Delete Rate Limit Rule
[**RateLimitsListRateLimitRules**](RateLimitsAPI.md#RateLimitsListRateLimitRules) | **Get** /api/v1/rate-limits | List Rate Limit Rules
[**RateLimitsUpdateRateLimitRule**](RateLimitsAPI.md#RateLimitsUpdateRateLimitRule) | **Patch** /api/v1/rate-limits/{name} | Update Rate Limit Rule



## RateLimitsCreateRateLimitRule

> RateLimitRulePublic RateLimitsCreateRateLimitRule(ctx).RateLimitRuleCreate(rateLimitRuleCreate).Execute()

Create Rate Limit Rule



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
	rateLimitRuleCreate := *openapiclient.NewRateLimitRuleCreate("Name_example", "Per_example") // RateLimitRuleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateLimitsAPI.RateLimitsCreateRateLimitRule(context.Background()).RateLimitRuleCreate(rateLimitRuleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsAPI.RateLimitsCreateRateLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateLimitsCreateRateLimitRule`: RateLimitRulePublic
	fmt.Fprintf(os.Stdout, "Response from `RateLimitsAPI.RateLimitsCreateRateLimitRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitsCreateRateLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rateLimitRuleCreate** | [**RateLimitRuleCreate**](RateLimitRuleCreate.md) |  | 

### Return type

[**RateLimitRulePublic**](RateLimitRulePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RateLimitsDeleteRateLimitRule

> RateLimitsDeleteRateLimitRule(ctx, name).Execute()

Delete Rate Limit Rule



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RateLimitsAPI.RateLimitsDeleteRateLimitRule(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsAPI.RateLimitsDeleteRateLimitRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRateLimitsDeleteRateLimitRuleRequest struct via the builder pattern


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


## RateLimitsListRateLimitRules

> RateLimitRulesPublic RateLimitsListRateLimitRules(ctx).Execute()

List Rate Limit Rules



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
	resp, r, err := apiClient.RateLimitsAPI.RateLimitsListRateLimitRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsAPI.RateLimitsListRateLimitRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateLimitsListRateLimitRules`: RateLimitRulesPublic
	fmt.Fprintf(os.Stdout, "Response from `RateLimitsAPI.RateLimitsListRateLimitRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitsListRateLimitRulesRequest struct via the builder pattern


### Return type

[**RateLimitRulesPublic**](RateLimitRulesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RateLimitsUpdateRateLimitRule

> RateLimitRulePublic RateLimitsUpdateRateLimitRule(ctx, name).RateLimitRuleUpdate(rateLimitRuleUpdate).Execute()

Update Rate Limit Rule



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
	rateLimitRuleUpdate := *openapiclient.NewRateLimitRuleUpdate() // RateLimitRuleUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RateLimitsAPI.RateLimitsUpdateRateLimitRule(context.Background(), name).RateLimitRuleUpdate(rateLimitRuleUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsAPI.RateLimitsUpdateRateLimitRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RateLimitsUpdateRateLimitRule`: RateLimitRulePublic
	fmt.Fprintf(os.Stdout, "Response from `RateLimitsAPI.RateLimitsUpdateRateLimitRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitsUpdateRateLimitRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateLimitRuleUpdate** | [**RateLimitRuleUpdate**](RateLimitRuleUpdate.md) |  | 

### Return type

[**RateLimitRulePublic**](RateLimitRulePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

