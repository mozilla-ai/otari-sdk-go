# \OrganizationTracesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationTracesCountTraces**](OrganizationTracesAPI.md#OrganizationTracesCountTraces) | **Get** /api/v1/organizations/me/traces/count | Count Traces
[**OrganizationTracesGetTrace**](OrganizationTracesAPI.md#OrganizationTracesGetTrace) | **Get** /api/v1/organizations/me/traces/{trace_id} | Get Trace
[**OrganizationTracesListTraces**](OrganizationTracesAPI.md#OrganizationTracesListTraces) | **Get** /api/v1/organizations/me/traces | List Traces
[**OrganizationTracesTraceOfRequest**](OrganizationTracesAPI.md#OrganizationTracesTraceOfRequest) | **Get** /api/v1/organizations/me/traces/requests/{request_id} | Trace Of Request
[**OrganizationTracesTraceSeries**](OrganizationTracesAPI.md#OrganizationTracesTraceSeries) | **Get** /api/v1/organizations/me/traces/series | Trace Series



## OrganizationTracesCountTraces

> TraceCountPublic OrganizationTracesCountTraces(ctx).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

Count Traces



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only sessions in this workspace. Narrows the caller's scope; never widens it. (optional)
	start := time.Now() // time.Time | Last activity at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Last activity before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | A trace id, or the start of one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesCountTraces(context.Background()).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesCountTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesCountTraces`: TraceCountPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesCountTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesCountTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Only sessions in this workspace. Narrows the caller&#39;s scope; never widens it. | 
 **start** | **time.Time** | Last activity at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Last activity before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | A trace id, or the start of one. | 

### Return type

[**TraceCountPublic**](TraceCountPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesGetTrace

> TraceDetailPublic OrganizationTracesGetTrace(ctx, traceId).WorkspaceId(workspaceId).Execute()

Get Trace



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
	traceId := "traceId_example" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workspace the trace is in, as its list entry names it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesGetTrace(context.Background(), traceId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesGetTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesGetTrace`: TraceDetailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesGetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 

### Return type

[**TraceDetailPublic**](TraceDetailPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesListTraces

> TraceListPublic OrganizationTracesListTraces(ctx).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

List Traces



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only sessions in this workspace. Narrows the caller's scope; never widens it. (optional)
	start := time.Now() // time.Time | Last activity at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Last activity before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | A trace id, or the start of one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesListTraces(context.Background()).Skip(skip).Limit(limit).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesListTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesListTraces`: TraceListPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesListTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesListTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]
 **workspaceId** | **string** | Only sessions in this workspace. Narrows the caller&#39;s scope; never widens it. | 
 **start** | **time.Time** | Last activity at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Last activity before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | A trace id, or the start of one. | 

### Return type

[**TraceListPublic**](TraceListPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesTraceOfRequest

> TraceSummaryPublic OrganizationTracesTraceOfRequest(ctx, requestId).Execute()

Trace Of Request



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
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesTraceOfRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesTraceOfRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesTraceOfRequest`: TraceSummaryPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesTraceOfRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesTraceOfRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TraceSummaryPublic**](TraceSummaryPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesTraceSeries

> TraceSeriesPublic OrganizationTracesTraceSeries(ctx).Bucket(bucket).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

Trace Series



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
	bucket := "bucket_example" // string |  (optional) (default to "hour")
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only sessions in this workspace. Narrows the caller's scope; never widens it. (optional)
	start := time.Now() // time.Time | Last activity at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Last activity before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | A trace id, or the start of one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesTraceSeries(context.Background()).Bucket(bucket).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesTraceSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesTraceSeries`: TraceSeriesPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesTraceSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesTraceSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string** |  | [default to &quot;hour&quot;]
 **workspaceId** | **string** | Only sessions in this workspace. Narrows the caller&#39;s scope; never widens it. | 
 **start** | **time.Time** | Last activity at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Last activity before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | A trace id, or the start of one. | 

### Return type

[**TraceSeriesPublic**](TraceSeriesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

