# \AgentTelemetryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentTelemetrySeriesV1AgentTelemetrySeriesGet**](AgentTelemetryAPI.md#AgentTelemetrySeriesV1AgentTelemetrySeriesGet) | **Get** /v1/agent-telemetry/series | Agent Telemetry Series
[**AgentTelemetrySummaryV1AgentTelemetrySummaryGet**](AgentTelemetryAPI.md#AgentTelemetrySummaryV1AgentTelemetrySummaryGet) | **Get** /v1/agent-telemetry/summary | Agent Telemetry Summary
[**CountAgentTelemetryV1AgentTelemetryCountGet**](AgentTelemetryAPI.md#CountAgentTelemetryV1AgentTelemetryCountGet) | **Get** /v1/agent-telemetry/count | Count Agent Telemetry
[**DeleteAgentTelemetryRowsV1AgentTelemetryDelete**](AgentTelemetryAPI.md#DeleteAgentTelemetryRowsV1AgentTelemetryDelete) | **Delete** /v1/agent-telemetry | Delete Agent Telemetry Rows



## AgentTelemetrySeriesV1AgentTelemetrySeriesGet

> AgentTelemetryGroupedSeries AgentTelemetrySeriesV1AgentTelemetrySeriesGet(ctx).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).Name(name).Bucket(bucket).Execute()

Agent Telemetry Series



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
	startDate := time.Now() // time.Time | Return rows with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return rows with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	name := "name_example" // string | Filter to a single event type or metric name (e.g. 'tool_result', 'claude_code.commit.count') (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTelemetryAPI.AgentTelemetrySeriesV1AgentTelemetrySeriesGet(context.Background()).GroupBy(groupBy).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).Name(name).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTelemetryAPI.AgentTelemetrySeriesV1AgentTelemetrySeriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentTelemetrySeriesV1AgentTelemetrySeriesGet`: AgentTelemetryGroupedSeries
	fmt.Fprintf(os.Stdout, "Response from `AgentTelemetryAPI.AgentTelemetrySeriesV1AgentTelemetrySeriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentTelemetrySeriesV1AgentTelemetrySeriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupBy** | **string** | Dimension to split the series by | 
 **startDate** | **time.Time** | Return rows with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return rows with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **name** | **string** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | 
 **bucket** | **string** | Time-series granularity: &#39;hour&#39; or &#39;day&#39; | [default to &quot;day&quot;]

### Return type

[**AgentTelemetryGroupedSeries**](AgentTelemetryGroupedSeries.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentTelemetrySummaryV1AgentTelemetrySummaryGet

> AgentTelemetrySummary AgentTelemetrySummaryV1AgentTelemetrySummaryGet(ctx).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).SessionLabel(sessionLabel).Bucket(bucket).Execute()

Agent Telemetry Summary



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
	startDate := time.Now() // time.Time | Return rows with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return rows with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	sessionLabel := "sessionLabel_example" // string | Filter to a single agent session. Matches agent_telemetry.session_label and, on the usage side of the join, the usage_logs.source_label that /v1/usage/summary filters on (optional)
	bucket := "bucket_example" // string | Time-series granularity: 'hour' or 'day' (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTelemetryAPI.AgentTelemetrySummaryV1AgentTelemetrySummaryGet(context.Background()).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).SessionLabel(sessionLabel).Bucket(bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTelemetryAPI.AgentTelemetrySummaryV1AgentTelemetrySummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentTelemetrySummaryV1AgentTelemetrySummaryGet`: AgentTelemetrySummary
	fmt.Fprintf(os.Stdout, "Response from `AgentTelemetryAPI.AgentTelemetrySummaryV1AgentTelemetrySummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentTelemetrySummaryV1AgentTelemetrySummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return rows with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return rows with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **sessionLabel** | **string** | Filter to a single agent session. Matches agent_telemetry.session_label and, on the usage side of the join, the usage_logs.source_label that /v1/usage/summary filters on | 
 **bucket** | **string** | Time-series granularity: &#39;hour&#39; or &#39;day&#39; | [default to &quot;day&quot;]

### Return type

[**AgentTelemetrySummary**](AgentTelemetrySummary.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAgentTelemetryV1AgentTelemetryCountGet

> AgentTelemetryCount CountAgentTelemetryV1AgentTelemetryCountGet(ctx).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).Name(name).Execute()

Count Agent Telemetry



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
	startDate := time.Now() // time.Time | Return rows with timestamp >= start_date (ISO 8601 or Unix epoch seconds) (optional)
	endDate := time.Now() // time.Time | Return rows with timestamp < end_date (ISO 8601 or Unix epoch seconds) (optional)
	userId := []string{"Inner_example"} // []string | Filter to one or more users; repeatable (user_id=a&user_id=b). Several values match any of them. At most 50 per call. (optional)
	apiKeyId := []string{"Inner_example"} // []string | Filter to one or more API key ids; repeatable (api_key_id=a&api_key_id=b). Several values match any of them. At most 50 per call. (optional)
	name := "name_example" // string | Filter to a single event type or metric name (e.g. 'tool_result', 'claude_code.commit.count') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTelemetryAPI.CountAgentTelemetryV1AgentTelemetryCountGet(context.Background()).StartDate(startDate).EndDate(endDate).UserId(userId).ApiKeyId(apiKeyId).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTelemetryAPI.CountAgentTelemetryV1AgentTelemetryCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAgentTelemetryV1AgentTelemetryCountGet`: AgentTelemetryCount
	fmt.Fprintf(os.Stdout, "Response from `AgentTelemetryAPI.CountAgentTelemetryV1AgentTelemetryCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountAgentTelemetryV1AgentTelemetryCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Return rows with timestamp &gt;&#x3D; start_date (ISO 8601 or Unix epoch seconds) | 
 **endDate** | **time.Time** | Return rows with timestamp &lt; end_date (ISO 8601 or Unix epoch seconds) | 
 **userId** | **[]string** | Filter to one or more users; repeatable (user_id&#x3D;a&amp;user_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **apiKeyId** | **[]string** | Filter to one or more API key ids; repeatable (api_key_id&#x3D;a&amp;api_key_id&#x3D;b). Several values match any of them. At most 50 per call. | 
 **name** | **string** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | 

### Return type

[**AgentTelemetryCount**](AgentTelemetryCount.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentTelemetryRowsV1AgentTelemetryDelete

> AgentTelemetryDeleteResult DeleteAgentTelemetryRowsV1AgentTelemetryDelete(ctx).AgentTelemetryDeleteRequest(agentTelemetryDeleteRequest).Execute()

Delete Agent Telemetry Rows



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
	agentTelemetryDeleteRequest := *openapiclient.NewAgentTelemetryDeleteRequest() // AgentTelemetryDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTelemetryAPI.DeleteAgentTelemetryRowsV1AgentTelemetryDelete(context.Background()).AgentTelemetryDeleteRequest(agentTelemetryDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTelemetryAPI.DeleteAgentTelemetryRowsV1AgentTelemetryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentTelemetryRowsV1AgentTelemetryDelete`: AgentTelemetryDeleteResult
	fmt.Fprintf(os.Stdout, "Response from `AgentTelemetryAPI.DeleteAgentTelemetryRowsV1AgentTelemetryDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentTelemetryRowsV1AgentTelemetryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentTelemetryDeleteRequest** | [**AgentTelemetryDeleteRequest**](AgentTelemetryDeleteRequest.md) |  | 

### Return type

[**AgentTelemetryDeleteResult**](AgentTelemetryDeleteResult.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

