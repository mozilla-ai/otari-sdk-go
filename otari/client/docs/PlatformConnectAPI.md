# \PlatformConnectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformConnectCancelPlatformConnect**](PlatformConnectAPI.md#PlatformConnectCancelPlatformConnect) | **Delete** /api/v1/platform/connect | Cancel Platform Connect
[**PlatformConnectGetPlatformConnect**](PlatformConnectAPI.md#PlatformConnectGetPlatformConnect) | **Get** /api/v1/platform/connect | Get Platform Connect
[**PlatformConnectStartPlatformConnect**](PlatformConnectAPI.md#PlatformConnectStartPlatformConnect) | **Post** /api/v1/platform/connect | Start Platform Connect



## PlatformConnectCancelPlatformConnect

> PlatformConnectCancelPlatformConnect(ctx).Execute()

Cancel Platform Connect



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
	r, err := apiClient.PlatformConnectAPI.PlatformConnectCancelPlatformConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConnectAPI.PlatformConnectCancelPlatformConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformConnectCancelPlatformConnectRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformConnectGetPlatformConnect

> PlatformConnectStatusPublic PlatformConnectGetPlatformConnect(ctx).Execute()

Get Platform Connect



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
	resp, r, err := apiClient.PlatformConnectAPI.PlatformConnectGetPlatformConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConnectAPI.PlatformConnectGetPlatformConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformConnectGetPlatformConnect`: PlatformConnectStatusPublic
	fmt.Fprintf(os.Stdout, "Response from `PlatformConnectAPI.PlatformConnectGetPlatformConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformConnectGetPlatformConnectRequest struct via the builder pattern


### Return type

[**PlatformConnectStatusPublic**](PlatformConnectStatusPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformConnectStartPlatformConnect

> PlatformConnectStartPublic PlatformConnectStartPlatformConnect(ctx).Execute()

Start Platform Connect



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
	resp, r, err := apiClient.PlatformConnectAPI.PlatformConnectStartPlatformConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformConnectAPI.PlatformConnectStartPlatformConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformConnectStartPlatformConnect`: PlatformConnectStartPublic
	fmt.Fprintf(os.Stdout, "Response from `PlatformConnectAPI.PlatformConnectStartPlatformConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformConnectStartPlatformConnectRequest struct via the builder pattern


### Return type

[**PlatformConnectStartPublic**](PlatformConnectStartPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

