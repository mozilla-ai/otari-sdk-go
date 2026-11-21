# \KeyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyGetKeyBudget**](KeyAPI.md#KeyGetKeyBudget) | **Get** /api/v1/key/budget | Get Key Budget
[**KeyGetKeyUsageSeries**](KeyAPI.md#KeyGetKeyUsageSeries) | **Get** /api/v1/key/usage/series | Get Key Usage Series
[**KeyGetKeyUsageSummary**](KeyAPI.md#KeyGetKeyUsageSummary) | **Get** /api/v1/key/usage/summary | Get Key Usage Summary



## KeyGetKeyBudget

> KeyBudgetPublic KeyGetKeyBudget(ctx).Execute()

Get Key Budget



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
	resp, r, err := apiClient.KeyAPI.KeyGetKeyBudget(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyGetKeyBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyGetKeyBudget`: KeyBudgetPublic
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.KeyGetKeyBudget`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetKeyBudgetRequest struct via the builder pattern


### Return type

[**KeyBudgetPublic**](KeyBudgetPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGetKeyUsageSeries

> UsageGroupedSeries KeyGetKeyUsageSeries(ctx).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).Tag(tag).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).Bucket(bucket).Execute()

Get Key Usage Series



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
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session (source_label): the session's trace id on a gateway request that named one, or the session/project label imported usage carried (optional)
	tag := []string{"Inner_example"} // []string | Filter by a request tag (what a request sent in its `metadata`), as `key:value`. Repeatable: values for the same key match any of them, and different keys must all match. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, web_fetch, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.KeyGetKeyUsageSeries(context.Background()).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).Tag(tag).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyGetKeyUsageSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyGetKeyUsageSeries`: UsageGroupedSeries
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.KeyGetKeyUsageSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetKeyUsageSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupBy** | **string** | Dimension to split the series by | 
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session (source_label): the session&#39;s trace id on a gateway request that named one, or the session/project label imported usage carried | 
 **tag** | **[]string** | Filter by a request tag (what a request sent in its &#x60;metadata&#x60;), as &#x60;key:value&#x60;. Repeatable: values for the same key match any of them, and different keys must all match. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, web_fetch, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
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


## KeyGetKeyUsageSummary

> UsageSummary KeyGetKeyUsageSummary(ctx).StartDate(startDate).EndDate(endDate).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).Tag(tag).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).Bucket(bucket).Dimensions(dimensions).GroupByTag(groupByTag).Execute()

Get Key Usage Summary



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
	status := "status_example" // string | Filter to a single status: 'success', 'error', or 'absorbed' (an attempt a routing policy recovered from, excluded from error_count and request_count) (optional)
	statusCode := int32(56) // int32 | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status='error' unless 'status' is given explicitly (optional)
	model := []string{"Inner_example"} // []string | Filter to one or more models; repeatable (model=a&model=b). Several values match any of them. At most 50 per call. (optional)
	endpoint := "endpoint_example" // string | Filter to a single endpoint (e.g. '/v1/chat/completions') (optional)
	provider := "provider_example" // string | Filter to a single provider (e.g. 'openai') (optional)
	source := "source_example" // string | Filter to a single provenance source (e.g. 'gateway' or 'claude_code') (optional)
	sourceLabel := "sourceLabel_example" // string | Filter to a single session (source_label): the session's trace id on a gateway request that named one, or the session/project label imported usage carried (optional)
	tag := []string{"Inner_example"} // []string | Filter by a request tag (what a request sent in its `metadata`), as `key:value`. Repeatable: values for the same key match any of them, and different keys must all match. At most 50 per call. (optional)
	priced := true // bool | Filter by token-pricing state: true = only rows whose model tokens were priced, false = only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. (optional)
	tool := "tool_example" // string | Filter to requests that ran a gateway-run tool. 'any' matches any tool; a tool name (web_search, web_fetch, code_execution) matches that tool specifically. (optional)
	countsTowardBudget := true // bool | Filter by budget participation, which is not the same question as provenance: true = only enforced gateway rows, false = every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")
	dimensions := []string{"Dimensions_example"} // []string | Which breakdowns to compute; repeatable (dimensions=model&dimensions=user). Each value names the 'by_<value>' response field it fills, except 'status_code', which fills the failure taxonomy in 'errors_by_status_code'. Omit for every breakdown (the default); pass 'none' for a totals-and-series-only response. Each dimension left out skips one GROUP BY scan, so a caller that reads only the tiles or the time series should say so. Fields that were not requested come back empty. (optional)
	groupByTag := "groupByTag_example" // string | A tag key to break spend down by, returned as `by_tag`. Rows that do not carry the tag group under a null key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.KeyGetKeyUsageSummary(context.Background()).StartDate(startDate).EndDate(endDate).Status(status).StatusCode(statusCode).Model(model).Endpoint(endpoint).Provider(provider).Source(source).SourceLabel(sourceLabel).Tag(tag).Priced(priced).Tool(tool).CountsTowardBudget(countsTowardBudget).Bucket(bucket).Dimensions(dimensions).GroupByTag(groupByTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyGetKeyUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyGetKeyUsageSummary`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.KeyGetKeyUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetKeyUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return logs with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return logs with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **status** | **string** | Filter to a single status: &#39;success&#39;, &#39;error&#39;, or &#39;absorbed&#39; (an attempt a routing policy recovered from, excluded from error_count and request_count) | 
 **statusCode** | **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | 
 **model** | **[]string** | Filter to one or more models; repeatable (model&#x3D;a&amp;model&#x3D;b). Several values match any of them. At most 50 per call. | 
 **endpoint** | **string** | Filter to a single endpoint (e.g. &#39;/v1/chat/completions&#39;) | 
 **provider** | **string** | Filter to a single provider (e.g. &#39;openai&#39;) | 
 **source** | **string** | Filter to a single provenance source (e.g. &#39;gateway&#39; or &#39;claude_code&#39;) | 
 **sourceLabel** | **string** | Filter to a single session (source_label): the session&#39;s trace id on a gateway request that named one, or the session/project label imported usage carried | 
 **tag** | **[]string** | Filter by a request tag (what a request sent in its &#x60;metadata&#x60;), as &#x60;key:value&#x60;. Repeatable: values for the same key match any of them, and different keys must all match. At most 50 per call. | 
 **priced** | **bool** | Filter by token-pricing state: true &#x3D; only rows whose model tokens were priced, false &#x3D; only rows that still need pricing (no cost at all, or tokens that were never metered because the model had no rate). A row charged only for gateway-run tool calls still counts as needing pricing. | 
 **tool** | **string** | Filter to requests that ran a gateway-run tool. &#39;any&#39; matches any tool; a tool name (web_search, web_fetch, code_execution) matches that tool specifically. | 
 **countsTowardBudget** | **bool** | Filter by budget participation, which is not the same question as provenance: true &#x3D; only enforced gateway rows, false &#x3D; every row that never touches a budget, meaning imported usage and also gateway traffic on a budget-exempt key | 
 **bucket** | **string** | Time-series granularity: &#39;hour&#39; or &#39;day&#39; | [default to &quot;day&quot;]
 **dimensions** | **[]string** | Which breakdowns to compute; repeatable (dimensions&#x3D;model&amp;dimensions&#x3D;user). Each value names the &#39;by_&lt;value&gt;&#39; response field it fills, except &#39;status_code&#39;, which fills the failure taxonomy in &#39;errors_by_status_code&#39;. Omit for every breakdown (the default); pass &#39;none&#39; for a totals-and-series-only response. Each dimension left out skips one GROUP BY scan, so a caller that reads only the tiles or the time series should say so. Fields that were not requested come back empty. | 
 **groupByTag** | **string** | A tag key to break spend down by, returned as &#x60;by_tag&#x60;. Rows that do not carry the tag group under a null key. | 

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

