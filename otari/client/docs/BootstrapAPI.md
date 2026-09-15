# \BootstrapAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BootstrapGetBootstrap**](BootstrapAPI.md#BootstrapGetBootstrap) | **Get** /api/v1/bootstrap | Get Bootstrap



## BootstrapGetBootstrap

> DeploymentBootstrap BootstrapGetBootstrap(ctx).Execute()

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
	resp, r, err := apiClient.BootstrapAPI.BootstrapGetBootstrap(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BootstrapAPI.BootstrapGetBootstrap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BootstrapGetBootstrap`: DeploymentBootstrap
	fmt.Fprintf(os.Stdout, "Response from `BootstrapAPI.BootstrapGetBootstrap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBootstrapGetBootstrapRequest struct via the builder pattern


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

