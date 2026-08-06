# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionV1AuthSessionPost**](AuthAPI.md#CreateSessionV1AuthSessionPost) | **Post** /v1/auth/session | Create Session
[**DeleteSessionV1AuthSessionDelete**](AuthAPI.md#DeleteSessionV1AuthSessionDelete) | **Delete** /v1/auth/session | Delete Session



## CreateSessionV1AuthSessionPost

> SessionResponse CreateSessionV1AuthSessionPost(ctx).CreateSessionRequest(createSessionRequest).Execute()

Create Session



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
	createSessionRequest := *openapiclient.NewCreateSessionRequest("MasterKey_example") // CreateSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.CreateSessionV1AuthSessionPost(context.Background()).CreateSessionRequest(createSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.CreateSessionV1AuthSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSessionV1AuthSessionPost`: SessionResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.CreateSessionV1AuthSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionV1AuthSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSessionRequest** | [**CreateSessionRequest**](CreateSessionRequest.md) |  | 

### Return type

[**SessionResponse**](SessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionV1AuthSessionDelete

> DeleteSessionV1AuthSessionDelete(ctx).Execute()

Delete Session



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
	r, err := apiClient.AuthAPI.DeleteSessionV1AuthSessionDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.DeleteSessionV1AuthSessionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionV1AuthSessionDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

