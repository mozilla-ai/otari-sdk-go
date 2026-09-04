# \BootstrapAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBootstrapV1BootstrapGet**](BootstrapAPI.md#GetBootstrapV1BootstrapGet) | **Get** /v1/bootstrap | Get Bootstrap



## GetBootstrapV1BootstrapGet

> DeploymentBootstrap GetBootstrapV1BootstrapGet(ctx).Execute()

Get Bootstrap



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
	resp, r, err := apiClient.BootstrapAPI.GetBootstrapV1BootstrapGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BootstrapAPI.GetBootstrapV1BootstrapGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBootstrapV1BootstrapGet`: DeploymentBootstrap
	fmt.Fprintf(os.Stdout, "Response from `BootstrapAPI.GetBootstrapV1BootstrapGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBootstrapV1BootstrapGetRequest struct via the builder pattern


### Return type

[**DeploymentBootstrap**](DeploymentBootstrap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

