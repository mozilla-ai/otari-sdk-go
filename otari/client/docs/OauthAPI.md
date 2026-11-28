# \OauthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthExchangeToken**](OauthAPI.md#OauthExchangeToken) | **Post** /api/v1/oauth/token | Exchange Token



## OauthExchangeToken

> FederationTokenResponse OauthExchangeToken(ctx).FederationTokenRequest(federationTokenRequest).Execute()

Exchange Token



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
	federationTokenRequest := *openapiclient.NewFederationTokenRequest() // FederationTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.OauthExchangeToken(context.Background()).FederationTokenRequest(federationTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.OauthExchangeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthExchangeToken`: FederationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.OauthExchangeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthExchangeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **federationTokenRequest** | [**FederationTokenRequest**](FederationTokenRequest.md) |  | 

### Return type

[**FederationTokenResponse**](FederationTokenResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

