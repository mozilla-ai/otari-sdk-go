# \OauthAppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthAppsAppDetails**](OauthAppsAPI.md#OauthAppsAppDetails) | **Get** /api/v1/oauth/apps/clients | App Details
[**OauthAppsAppWorkspaces**](OauthAppsAPI.md#OauthAppsAppWorkspaces) | **Get** /api/v1/oauth/apps/workspaces | App Workspaces
[**OauthAppsAuthorizationRedirect**](OauthAppsAPI.md#OauthAppsAuthorizationRedirect) | **Get** /api/v1/oauth/apps/authorize | Authorization Redirect
[**OauthAppsAuthorize**](OauthAppsAPI.md#OauthAppsAuthorize) | **Post** /api/v1/oauth/apps/authorize | Authorize
[**OauthAppsExchange**](OauthAppsAPI.md#OauthAppsExchange) | **Post** /api/v1/oauth/apps/token | Exchange
[**OauthAppsRegisterClient**](OauthAppsAPI.md#OauthAppsRegisterClient) | **Post** /api/v1/oauth/apps/register | Register Client
[**OauthAppsRevoke**](OauthAppsAPI.md#OauthAppsRevoke) | **Post** /api/v1/oauth/apps/revoke | Revoke



## OauthAppsAppDetails

> AppDetails OauthAppsAppDetails(ctx).ClientId(clientId).RedirectUri(redirectUri).Execute()

App Details



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
	clientId := "clientId_example" // string | 
	redirectUri := "redirectUri_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsAppDetails(context.Background()).ClientId(clientId).RedirectUri(redirectUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsAppDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsAppDetails`: AppDetails
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsAppDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsAppDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **redirectUri** | **string** |  | 

### Return type

[**AppDetails**](AppDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsAppWorkspaces

> WorkspacesPublic OauthAppsAppWorkspaces(ctx).Search(search).Skip(skip).Limit(limit).Execute()

App Workspaces



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
	search := "search_example" // string |  (optional) (default to "")
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsAppWorkspaces(context.Background()).Search(search).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsAppWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsAppWorkspaces`: WorkspacesPublic
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsAppWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsAppWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | [default to &quot;&quot;]
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**WorkspacesPublic**](WorkspacesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsAuthorizationRedirect

> OauthAppsAuthorizationRedirect(ctx).ClientId(clientId).RedirectUri(redirectUri).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).ResponseType(responseType).Limit(limit).Expiration(expiration).Execute()

Authorization Redirect



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
	clientId := "clientId_example" // string | 
	redirectUri := "redirectUri_example" // string | 
	state := "state_example" // string | 
	codeChallenge := "codeChallenge_example" // string | 
	codeChallengeMethod := "codeChallengeMethod_example" // string | 
	responseType := "responseType_example" // string |  (optional) (default to "code")
	limit := "limit_example" // string | A spending limit in USD to suggest for the key. The person can change it on the consent screen, and one the screen cannot use is left blank. (optional)
	expiration := "expiration_example" // string | An expiry to suggest for the key: 1-hour, 1-day, 7-days, 30-days, 90-days, 180-days, 1-year or never. The person can change it, and any other value leaves it at never. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OauthAppsAPI.OauthAppsAuthorizationRedirect(context.Background()).ClientId(clientId).RedirectUri(redirectUri).State(state).CodeChallenge(codeChallenge).CodeChallengeMethod(codeChallengeMethod).ResponseType(responseType).Limit(limit).Expiration(expiration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsAuthorizationRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsAuthorizationRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **redirectUri** | **string** |  | 
 **state** | **string** |  | 
 **codeChallenge** | **string** |  | 
 **codeChallengeMethod** | **string** |  | 
 **responseType** | **string** |  | [default to &quot;code&quot;]
 **limit** | **string** | A spending limit in USD to suggest for the key. The person can change it on the consent screen, and one the screen cannot use is left blank. | 
 **expiration** | **string** | An expiry to suggest for the key: 1-hour, 1-day, 7-days, 30-days, 90-days, 180-days, 1-year or never. The person can change it, and any other value leaves it at never. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsAuthorize

> AppAuthorizationResponse OauthAppsAuthorize(ctx).AppAuthorizationRequest(appAuthorizationRequest).Execute()

Authorize



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
	appAuthorizationRequest := *openapiclient.NewAppAuthorizationRequest("ClientId_example", "CodeChallenge_example", "CodeChallengeMethod_example", "RedirectUri_example", []openapiclient.KeyScope{openapiclient.KeyScope("inference:run")}, "State_example") // AppAuthorizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsAuthorize(context.Background()).AppAuthorizationRequest(appAuthorizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsAuthorize`: AppAuthorizationResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsAuthorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appAuthorizationRequest** | [**AppAuthorizationRequest**](AppAuthorizationRequest.md) |  | 

### Return type

[**AppAuthorizationResponse**](AppAuthorizationResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsExchange

> AppTokenResponse OauthAppsExchange(ctx).ClientId(clientId).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()

Exchange



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
	clientId := "clientId_example" // string | 
	grantType := "grantType_example" // string | 
	code := "code_example" // string |  (optional)
	codeVerifier := "codeVerifier_example" // string |  (optional)
	redirectUri := "redirectUri_example" // string |  (optional)
	refreshToken := "refreshToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsExchange(context.Background()).ClientId(clientId).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsExchange`: AppTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **grantType** | **string** |  | 
 **code** | **string** |  | 
 **codeVerifier** | **string** |  | 
 **redirectUri** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**AppTokenResponse**](AppTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsRegisterClient

> ClientRegistrationResponse OauthAppsRegisterClient(ctx).ClientRegistrationRequest(clientRegistrationRequest).Execute()

Register Client



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
	clientRegistrationRequest := *openapiclient.NewClientRegistrationRequest([]string{"RedirectUris_example"}) // ClientRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsRegisterClient(context.Background()).ClientRegistrationRequest(clientRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsRegisterClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsRegisterClient`: ClientRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsRegisterClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsRegisterClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRegistrationRequest** | [**ClientRegistrationRequest**](ClientRegistrationRequest.md) |  | 

### Return type

[**ClientRegistrationResponse**](ClientRegistrationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthAppsRevoke

> interface{} OauthAppsRevoke(ctx).Token(token).TokenTypeHint(tokenTypeHint).Execute()

Revoke



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
	token := "token_example" // string | 
	tokenTypeHint := "tokenTypeHint_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAppsAPI.OauthAppsRevoke(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAppsAPI.OauthAppsRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthAppsRevoke`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OauthAppsAPI.OauthAppsRevoke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthAppsRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 
 **tokenTypeHint** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

