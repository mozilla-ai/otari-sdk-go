# \OtelAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceiveLogsV1LogsPost**](OtelAPI.md#ReceiveLogsV1LogsPost) | **Post** /v1/logs | Receive Logs
[**ReceiveMetricsV1MetricsPost**](OtelAPI.md#ReceiveMetricsV1MetricsPost) | **Post** /v1/metrics | Receive Metrics
[**ReceiveTracesV1TracesPost**](OtelAPI.md#ReceiveTracesV1TracesPost) | **Post** /v1/traces | Receive Traces



## ReceiveLogsV1LogsPost

> interface{} ReceiveLogsV1LogsPost(ctx).Execute()

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
	resp, r, err := apiClient.OtelAPI.ReceiveLogsV1LogsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.ReceiveLogsV1LogsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveLogsV1LogsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.ReceiveLogsV1LogsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveLogsV1LogsPostRequest struct via the builder pattern


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


## ReceiveMetricsV1MetricsPost

> interface{} ReceiveMetricsV1MetricsPost(ctx).Execute()

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
	resp, r, err := apiClient.OtelAPI.ReceiveMetricsV1MetricsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.ReceiveMetricsV1MetricsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveMetricsV1MetricsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.ReceiveMetricsV1MetricsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveMetricsV1MetricsPostRequest struct via the builder pattern


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


## ReceiveTracesV1TracesPost

> interface{} ReceiveTracesV1TracesPost(ctx).Execute()

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
	resp, r, err := apiClient.OtelAPI.ReceiveTracesV1TracesPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtelAPI.ReceiveTracesV1TracesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveTracesV1TracesPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtelAPI.ReceiveTracesV1TracesPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveTracesV1TracesPostRequest struct via the builder pattern


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

