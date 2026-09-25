# \HooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HooksCheckPolicy**](HooksAPI.md#HooksCheckPolicy) | **Post** /api/v1/hooks/check | Check Policy



## HooksCheckPolicy

> PolicyCheckResponse HooksCheckPolicy(ctx).PolicyCheckRequest(policyCheckRequest).Execute()

Check Policy



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
	policyCheckRequest := *openapiclient.NewPolicyCheckRequest("PolicyYaml_example") // PolicyCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.HooksCheckPolicy(context.Background()).PolicyCheckRequest(policyCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.HooksCheckPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HooksCheckPolicy`: PolicyCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.HooksCheckPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHooksCheckPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyCheckRequest** | [**PolicyCheckRequest**](PolicyCheckRequest.md) |  | 

### Return type

[**PolicyCheckResponse**](PolicyCheckResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

