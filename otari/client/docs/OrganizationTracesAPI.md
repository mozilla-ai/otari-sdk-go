# \OrganizationTracesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationTracesCountTraces**](OrganizationTracesAPI.md#OrganizationTracesCountTraces) | **Get** /api/v1/organizations/me/traces/count | Count Traces
[**OrganizationTracesEndTraceContentShare**](OrganizationTracesAPI.md#OrganizationTracesEndTraceContentShare) | **Delete** /api/v1/organizations/me/traces/{trace_id}/shares/{share_id} | End Trace Content Share
[**OrganizationTracesGetTrace**](OrganizationTracesAPI.md#OrganizationTracesGetTrace) | **Get** /api/v1/organizations/me/traces/{trace_id} | Get Trace
[**OrganizationTracesListOrganizationBreakGlassReads**](OrganizationTracesAPI.md#OrganizationTracesListOrganizationBreakGlassReads) | **Get** /api/v1/organizations/me/traces/audit/break-glass | List Organization Break Glass Reads
[**OrganizationTracesListTraceContentShareReads**](OrganizationTracesAPI.md#OrganizationTracesListTraceContentShareReads) | **Get** /api/v1/organizations/me/traces/{trace_id}/shares/{share_id}/reads | List Trace Content Share Reads
[**OrganizationTracesListTraceContentShares**](OrganizationTracesAPI.md#OrganizationTracesListTraceContentShares) | **Get** /api/v1/organizations/me/traces/{trace_id}/shares | List Trace Content Shares
[**OrganizationTracesListTraces**](OrganizationTracesAPI.md#OrganizationTracesListTraces) | **Get** /api/v1/organizations/me/traces | List Traces
[**OrganizationTracesOrganizationBreakGlass**](OrganizationTracesAPI.md#OrganizationTracesOrganizationBreakGlass) | **Post** /api/v1/organizations/me/traces/{trace_id}/break-glass | Organization Break Glass
[**OrganizationTracesShareTraceContent**](OrganizationTracesAPI.md#OrganizationTracesShareTraceContent) | **Post** /api/v1/organizations/me/traces/{trace_id}/shares | Share Trace Content
[**OrganizationTracesTraceOfRequest**](OrganizationTracesAPI.md#OrganizationTracesTraceOfRequest) | **Get** /api/v1/organizations/me/traces/requests/{request_id} | Trace Of Request
[**OrganizationTracesTraceSeries**](OrganizationTracesAPI.md#OrganizationTracesTraceSeries) | **Get** /api/v1/organizations/me/traces/series | Trace Series



## OrganizationTracesCountTraces

> TraceCountPublic OrganizationTracesCountTraces(ctx).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

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
	start := time.Now() // time.Time | Sessions active at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Sessions started before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	model := []string{"Inner_example"} // []string | Sessions with an LLM call recorded under any of these model ids (exact match). (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | Part of a trace id, session name, harness or user id, in any case. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesCountTraces(context.Background()).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
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
 **start** | **time.Time** | Sessions active at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Sessions started before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **model** | **[]string** | Sessions with an LLM call recorded under any of these model ids (exact match). | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | Part of a trace id, session name, harness or user id, in any case. | 

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


## OrganizationTracesEndTraceContentShare

> OrganizationTracesEndTraceContentShare(ctx, traceId, shareId).WorkspaceId(workspaceId).Execute()

End Trace Content Share



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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workspace the trace is in, as its list entry names it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationTracesAPI.OrganizationTracesEndTraceContentShare(context.Background(), traceId, shareId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesEndTraceContentShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesEndTraceContentShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 

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


## OrganizationTracesListOrganizationBreakGlassReads

> BreakGlassReadListPublic OrganizationTracesListOrganizationBreakGlassReads(ctx).Skip(skip).Limit(limit).Execute()

List Organization Break Glass Reads



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesListOrganizationBreakGlassReads(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesListOrganizationBreakGlassReads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesListOrganizationBreakGlassReads`: BreakGlassReadListPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesListOrganizationBreakGlassReads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesListOrganizationBreakGlassReadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**BreakGlassReadListPublic**](BreakGlassReadListPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesListTraceContentShareReads

> ContentReadListPublic OrganizationTracesListTraceContentShareReads(ctx, traceId, shareId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()

List Trace Content Share Reads



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
	shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workspace the trace is in, as its list entry names it. (optional)
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesListTraceContentShareReads(context.Background(), traceId, shareId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesListTraceContentShareReads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesListTraceContentShareReads`: ContentReadListPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesListTraceContentShareReads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesListTraceContentShareReadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**ContentReadListPublic**](ContentReadListPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesListTraceContentShares

> ContentShareListPublic OrganizationTracesListTraceContentShares(ctx, traceId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()

List Trace Content Shares



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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesListTraceContentShares(context.Background(), traceId).WorkspaceId(workspaceId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesListTraceContentShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesListTraceContentShares`: ContentShareListPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesListTraceContentShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesListTraceContentSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**ContentShareListPublic**](ContentShareListPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesListTraces

> TraceListPublic OrganizationTracesListTraces(ctx).Skip(skip).Limit(limit).Sort(sort).Order(order).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

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
	sort := "sort_example" // string | What to order by. Latency runs from a session's start to its last activity. (optional) (default to "last_activity")
	order := "order_example" // string | The direction of `sort`. (optional) (default to "desc")
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only sessions in this workspace. Narrows the caller's scope; never widens it. (optional)
	start := time.Now() // time.Time | Sessions active at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Sessions started before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	model := []string{"Inner_example"} // []string | Sessions with an LLM call recorded under any of these model ids (exact match). (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | Part of a trace id, session name, harness or user id, in any case. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesListTraces(context.Background()).Skip(skip).Limit(limit).Sort(sort).Order(order).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
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
 **sort** | **string** | What to order by. Latency runs from a session&#39;s start to its last activity. | [default to &quot;last_activity&quot;]
 **order** | **string** | The direction of &#x60;sort&#x60;. | [default to &quot;desc&quot;]
 **workspaceId** | **string** | Only sessions in this workspace. Narrows the caller&#39;s scope; never widens it. | 
 **start** | **time.Time** | Sessions active at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Sessions started before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **model** | **[]string** | Sessions with an LLM call recorded under any of these model ids (exact match). | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | Part of a trace id, session name, harness or user id, in any case. | 

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


## OrganizationTracesOrganizationBreakGlass

> TraceDetailPublic OrganizationTracesOrganizationBreakGlass(ctx, traceId).BreakGlassCreate(breakGlassCreate).WorkspaceId(workspaceId).Execute()

Organization Break Glass



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
	breakGlassCreate := *openapiclient.NewBreakGlassCreate("Reason_example") // BreakGlassCreate | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workspace the trace is in, as its list entry names it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesOrganizationBreakGlass(context.Background(), traceId).BreakGlassCreate(breakGlassCreate).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesOrganizationBreakGlass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesOrganizationBreakGlass`: TraceDetailPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesOrganizationBreakGlass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesOrganizationBreakGlassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **breakGlassCreate** | [**BreakGlassCreate**](BreakGlassCreate.md) |  | 
 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 

### Return type

[**TraceDetailPublic**](TraceDetailPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationTracesShareTraceContent

> ContentSharePublic OrganizationTracesShareTraceContent(ctx, traceId).ContentShareCreate(contentShareCreate).WorkspaceId(workspaceId).Execute()

Share Trace Content



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
	contentShareCreate := *openapiclient.NewContentShareCreate("Audience_example") // ContentShareCreate | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The workspace the trace is in, as its list entry names it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesShareTraceContent(context.Background(), traceId).ContentShareCreate(contentShareCreate).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTracesAPI.OrganizationTracesShareTraceContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationTracesShareTraceContent`: ContentSharePublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTracesAPI.OrganizationTracesShareTraceContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationTracesShareTraceContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentShareCreate** | [**ContentShareCreate**](ContentShareCreate.md) |  | 
 **workspaceId** | **string** | The workspace the trace is in, as its list entry names it. | 

### Return type

[**ContentSharePublic**](ContentSharePublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

> TraceSeriesPublic OrganizationTracesTraceSeries(ctx).Bucket(bucket).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()

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
	start := time.Now() // time.Time | Sessions active at or after this instant (ISO 8601). (optional)
	end := time.Now() // time.Time | Sessions started before this instant (ISO 8601). (optional)
	userId := []string{"Inner_example"} // []string | Owning users. (optional)
	apiKeyId := []string{"Inner_example"} // []string | API keys. (optional)
	model := []string{"Inner_example"} // []string | Sessions with an LLM call recorded under any of these model ids (exact match). (optional)
	harness := []string{"Inner_example"} // []string | Agent harnesses. (optional)
	sessionSource := []string{"SessionSource_example"} // []string | What grouped the session. (optional)
	hasError := true // bool | Only sessions with, or without, an unrecovered error. (optional)
	q := "q_example" // string | Part of a trace id, session name, harness or user id, in any case. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTracesAPI.OrganizationTracesTraceSeries(context.Background()).Bucket(bucket).WorkspaceId(workspaceId).Start(start).End(end).UserId(userId).ApiKeyId(apiKeyId).Model(model).Harness(harness).SessionSource(sessionSource).HasError(hasError).Q(q).Execute()
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
 **start** | **time.Time** | Sessions active at or after this instant (ISO 8601). | 
 **end** | **time.Time** | Sessions started before this instant (ISO 8601). | 
 **userId** | **[]string** | Owning users. | 
 **apiKeyId** | **[]string** | API keys. | 
 **model** | **[]string** | Sessions with an LLM call recorded under any of these model ids (exact match). | 
 **harness** | **[]string** | Agent harnesses. | 
 **sessionSource** | **[]string** | What grouped the session. | 
 **hasError** | **bool** | Only sessions with, or without, an unrecovered error. | 
 **q** | **string** | Part of a trace id, session name, harness or user id, in any case. | 

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

