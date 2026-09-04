# \OrganizationUsageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrganizationUsageV1OrganizationsMeUsageCountGet**](OrganizationUsageAPI.md#CountOrganizationUsageV1OrganizationsMeUsageCountGet) | **Get** /v1/organizations/me/usage/count | Count Organization Usage
[**ListOrganizationUsageV1OrganizationsMeUsageGet**](OrganizationUsageAPI.md#ListOrganizationUsageV1OrganizationsMeUsageGet) | **Get** /v1/organizations/me/usage | List Organization Usage
[**OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet**](OrganizationUsageAPI.md#OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet) | **Get** /v1/organizations/me/usage/series | Organization Usage Series
[**OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet**](OrganizationUsageAPI.md#OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet) | **Get** /v1/organizations/me/usage/summary | Organization Usage Summary



## CountOrganizationUsageV1OrganizationsMeUsageCountGet

> UsageCount CountOrganizationUsageV1OrganizationsMeUsageCountGet(ctx).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).RequestGroupId(requestGroupId).WorkspaceId(workspaceId).Execute()

Count Organization Usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startDate := time.Now() // time.Time | Return logs with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return logs with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session/project label (the source_label carried by imported usage) (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	requestGroupId := []string{"Inner_example"} // []string | Filter to the rows of one or more request groups; repeatable (request_group_id=a&request_group_id=b). A routed request writes one row per attempt, all sharing a request_group_id, so this returns a request's whole plan: its absorbed attempts and the attempt that served it. Ignore ordering by timestamp and read attempt_position to reconstruct the plan. At most 1000 ids per call. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only usage recorded in this workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationUsageAPI.CountOrganizationUsageV1OrganizationsMeUsageCountGet(context.Background()).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).RequestGroupId(requestGroupId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationUsageAPI.CountOrganizationUsageV1OrganizationsMeUsageCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrganizationUsageV1OrganizationsMeUsageCountGet`: UsageCount
	fmt.Fprintf(os.Stdout, "Response from `OrganizationUsageAPI.CountOrganizationUsageV1OrganizationsMeUsageCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountOrganizationUsageV1OrganizationsMeUsageCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session/project label (the source_label carried by imported usage) | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
 **requestGroupId** | **[]string** | Filter to the rows of one or more request groups; repeatable (request_group_id&#x3D;a&amp;request_group_id&#x3D;b). A routed request writes one row per attempt, all sharing a request_group_id, so this returns a request&#39;s whole plan: its absorbed attempts and the attempt that served it. Ignore ordering by timestamp and read attempt_position to reconstruct the plan. At most 1000 ids per call. | 
 **workspaceId** | **string** | Only usage recorded in this workspace. | 

### Return type

[**UsageCount**](UsageCount.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationUsageV1OrganizationsMeUsageGet

> []UsageEntry ListOrganizationUsageV1OrganizationsMeUsageGet(ctx).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).RequestGroupId(requestGroupId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()

List Organization Usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startDate := time.Now() // time.Time | Return logs with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return logs with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session/project label (the source_label carried by imported usage) (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	requestGroupId := []string{"Inner_example"} // []string | Filter to the rows of one or more request groups; repeatable (request_group_id=a&request_group_id=b). A routed request writes one row per attempt, all sharing a request_group_id, so this returns a request's whole plan: its absorbed attempts and the attempt that served it. Ignore ordering by timestamp and read attempt_position to reconstruct the plan. At most 1000 ids per call. (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only usage recorded in this workspace. (optional)
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationUsageAPI.ListOrganizationUsageV1OrganizationsMeUsageGet(context.Background()).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).RequestGroupId(requestGroupId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationUsageAPI.ListOrganizationUsageV1OrganizationsMeUsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationUsageV1OrganizationsMeUsageGet`: []UsageEntry
	fmt.Fprintf(os.Stdout, "Response from `OrganizationUsageAPI.ListOrganizationUsageV1OrganizationsMeUsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationUsageV1OrganizationsMeUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session/project label (the source_label carried by imported usage) | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
 **requestGroupId** | **[]string** | Filter to the rows of one or more request groups; repeatable (request_group_id&#x3D;a&amp;request_group_id&#x3D;b). A routed request writes one row per attempt, all sharing a request_group_id, so this returns a request&#39;s whole plan: its absorbed attempts and the attempt that served it. Ignore ordering by timestamp and read attempt_position to reconstruct the plan. At most 1000 ids per call. | 
 **workspaceId** | **string** | Only usage recorded in this workspace. | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]UsageEntry**](UsageEntry.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet

> UsageGroupedSeries OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet(ctx).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).WorkspaceId(workspaceId).Bucket(bucket).Execute()

Organization Usage Series



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	groupBy := "groupBy_example" // string | Dimension to split the series by
	startDate := time.Now() // time.Time | Return logs with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return logs with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session/project label (the source_label carried by imported usage) (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only usage recorded in this workspace. (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationUsageAPI.OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet(context.Background()).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).WorkspaceId(workspaceId).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationUsageAPI.OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet`: UsageGroupedSeries
	fmt.Fprintf(os.Stdout, "Response from `OrganizationUsageAPI.OrganizationUsageSeriesV1OrganizationsMeUsageSeriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUsageSeriesV1OrganizationsMeUsageSeriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupBy** | **string** | Dimension to split the series by | 
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session/project label (the source_label carried by imported usage) | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
 **workspaceId** | **string** | Only usage recorded in this workspace. | 
 **bucket** | **string** | Time-series granularity: &#39;hour&#39; or &#39;day&#39; | [default to &quot;day&quot;]

### Return type

[**UsageGroupedSeries**](UsageGroupedSeries.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet

> UsageSummary OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet(ctx).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).WorkspaceId(workspaceId).Bucket(bucket).Dimensions(dimensions).Execute()

Organization Usage Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startDate := time.Now() // time.Time | Return logs with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return logs with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session/project label (the source_label carried by imported usage) (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only usage recorded in this workspace. (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")
	dimensions := []string{"Dimensions_example"} // []string | Which breakdowns to compute; repeatable (dimensions=model&dimensions=user). Each value names the 'by_<value>' response field it fills, except 'status_code', which fills the failure taxonomy in 'errors_by_status_code'. Omit for every breakdown (the default); pass 'none' for a totals-and-series-only response. Each dimension left out skips one GROUP BY scan, so a caller that reads only the tiles or the time series should say so. Fields that were not requested come back empty. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationUsageAPI.OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet(context.Background()).StartDate(startDate).EndDate(endDate).UserId(userId).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).ApiKeyId(apiKeyId).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).WorkspaceId(workspaceId).Bucket(bucket).Dimensions(dimensions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationUsageAPI.OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `OrganizationUsageAPI.OrganizationUsageSummaryV1OrganizationsMeUsageSummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationUsageSummaryV1OrganizationsMeUsageSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session/project label (the source_label carried by imported usage) | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
 **workspaceId** | **string** | Only usage recorded in this workspace. | 
 **bucket** | **string** | Time-series granularity: &#39;hour&#39; or &#39;day&#39; | [default to &quot;day&quot;]
 **dimensions** | **[]string** | Which breakdowns to compute; repeatable (dimensions&#x3D;model&amp;dimensions&#x3D;user). Each value names the &#39;by_&lt;value&gt;&#39; response field it fills, except &#39;status_code&#39;, which fills the failure taxonomy in &#39;errors_by_status_code&#39;. Omit for every breakdown (the default); pass &#39;none&#39; for a totals-and-series-only response. Each dimension left out skips one GROUP BY scan, so a caller that reads only the tiles or the time series should say so. Fields that were not requested come back empty. | 

### Return type

[**UsageSummary**](UsageSummary.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

