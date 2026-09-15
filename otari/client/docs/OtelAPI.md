# \OtelAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtelReceiveLogs**](OtelAPI.md#OtelReceiveLogs) | **Post** /otlp/v1/logs | Receive Logs
[**OtelReceiveMetrics**](OtelAPI.md#OtelReceiveMetrics) | **Post** /otlp/v1/metrics | Receive Metrics
[**OtelReceiveTraces**](OtelAPI.md#OtelReceiveTraces) | **Post** /otlp/v1/traces | Receive Traces



## OtelReceiveLogs

> interface{} OtelReceiveLogs(ctx).Execute()

Receive Logs



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
	resp, r, err := apiClient.OtelAPI.OtelReceiveLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.OtelReceiveLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtelReceiveLogs`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.OtelReceiveLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOtelReceiveLogsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtelReceiveMetrics

> interface{} OtelReceiveMetrics(ctx).Execute()

Receive Metrics



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
	resp, r, err := apiClient.OtelAPI.OtelReceiveMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.OtelReceiveMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtelReceiveMetrics`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.OtelReceiveMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOtelReceiveMetricsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtelReceiveTraces

> interface{} OtelReceiveTraces(ctx).Execute()

Receive Traces



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
	resp, r, err := apiClient.OtelAPI.OtelReceiveTraces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.OtelReceiveTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OtelReceiveTraces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.OtelReceiveTraces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOtelReceiveTracesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

