# \OrganizationSsoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationSsoDeleteOrganizationSso**](OrganizationSsoAPI.md#OrganizationSsoDeleteOrganizationSso) | **Delete** /api/v1/organizations/me/sso | Delete Organization Sso
[**OrganizationSsoGetOrganizationSso**](OrganizationSsoAPI.md#OrganizationSsoGetOrganizationSso) | **Get** /api/v1/organizations/me/sso | Get Organization Sso
[**OrganizationSsoGetOrganizationSsoPreflight**](OrganizationSsoAPI.md#OrganizationSsoGetOrganizationSsoPreflight) | **Get** /api/v1/organizations/me/sso/preflight | Get Organization Sso Preflight
[**OrganizationSsoPatchOrganizationSso**](OrganizationSsoAPI.md#OrganizationSsoPatchOrganizationSso) | **Patch** /api/v1/organizations/me/sso | Patch Organization Sso
[**OrganizationSsoPutOrganizationSso**](OrganizationSsoAPI.md#OrganizationSsoPutOrganizationSso) | **Put** /api/v1/organizations/me/sso | Put Organization Sso
[**OrganizationSsoPutOrganizationSsoEnforcement**](OrganizationSsoAPI.md#OrganizationSsoPutOrganizationSsoEnforcement) | **Put** /api/v1/organizations/me/sso/enforcement | Put Organization Sso Enforcement
[**OrganizationSsoTestOrganizationSso**](OrganizationSsoAPI.md#OrganizationSsoTestOrganizationSso) | **Post** /api/v1/organizations/me/sso/test | Test Organization Sso



## OrganizationSsoDeleteOrganizationSso

> OrganizationSsoDeleteOrganizationSso(ctx).Execute()

Delete Organization Sso



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
	r, err := apiClient.OrganizationSsoAPI.OrganizationSsoDeleteOrganizationSso(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoDeleteOrganizationSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoDeleteOrganizationSsoRequest struct via the builder pattern


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


## OrganizationSsoGetOrganizationSso

> OrganizationSsoPublic OrganizationSsoGetOrganizationSso(ctx).Execute()

Get Organization Sso



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
	resp, r, err := apiClient.OrganizationSsoAPI.OrganizationSsoGetOrganizationSso(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoGetOrganizationSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSsoGetOrganizationSso`: OrganizationSsoPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationSsoAPI.OrganizationSsoGetOrganizationSso`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoGetOrganizationSsoRequest struct via the builder pattern


### Return type

[**OrganizationSsoPublic**](OrganizationSsoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationSsoGetOrganizationSsoPreflight

> OrganizationSsoPreflightPublic OrganizationSsoGetOrganizationSsoPreflight(ctx).Execute()

Get Organization Sso Preflight



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
	resp, r, err := apiClient.OrganizationSsoAPI.OrganizationSsoGetOrganizationSsoPreflight(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoGetOrganizationSsoPreflight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSsoGetOrganizationSsoPreflight`: OrganizationSsoPreflightPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationSsoAPI.OrganizationSsoGetOrganizationSsoPreflight`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoGetOrganizationSsoPreflightRequest struct via the builder pattern


### Return type

[**OrganizationSsoPreflightPublic**](OrganizationSsoPreflightPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationSsoPatchOrganizationSso

> OrganizationSsoPublic OrganizationSsoPatchOrganizationSso(ctx).OIDCConnectionEnabledRequest(oIDCConnectionEnabledRequest).Execute()

Patch Organization Sso



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
	oIDCConnectionEnabledRequest := *openapiclient.NewOIDCConnectionEnabledRequest(false) // OIDCConnectionEnabledRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationSsoAPI.OrganizationSsoPatchOrganizationSso(context.Background()).OIDCConnectionEnabledRequest(oIDCConnectionEnabledRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoPatchOrganizationSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSsoPatchOrganizationSso`: OrganizationSsoPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationSsoAPI.OrganizationSsoPatchOrganizationSso`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoPatchOrganizationSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCConnectionEnabledRequest** | [**OIDCConnectionEnabledRequest**](OIDCConnectionEnabledRequest.md) |  | 

### Return type

[**OrganizationSsoPublic**](OrganizationSsoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationSsoPutOrganizationSso

> OrganizationSsoPublic OrganizationSsoPutOrganizationSso(ctx).OIDCConnectionRequest(oIDCConnectionRequest).Execute()

Put Organization Sso



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
	oIDCConnectionRequest := *openapiclient.NewOIDCConnectionRequest("ClientId_example", "DisplayName_example", "Issuer_example") // OIDCConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationSsoAPI.OrganizationSsoPutOrganizationSso(context.Background()).OIDCConnectionRequest(oIDCConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoPutOrganizationSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSsoPutOrganizationSso`: OrganizationSsoPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationSsoAPI.OrganizationSsoPutOrganizationSso`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoPutOrganizationSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oIDCConnectionRequest** | [**OIDCConnectionRequest**](OIDCConnectionRequest.md) |  | 

### Return type

[**OrganizationSsoPublic**](OrganizationSsoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationSsoPutOrganizationSsoEnforcement

> OrganizationSsoPutOrganizationSsoEnforcement(ctx).OrganizationSsoEnforcementRequest(organizationSsoEnforcementRequest).Execute()

Put Organization Sso Enforcement



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
	organizationSsoEnforcementRequest := *openapiclient.NewOrganizationSsoEnforcementRequest(false) // OrganizationSsoEnforcementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationSsoAPI.OrganizationSsoPutOrganizationSsoEnforcement(context.Background()).OrganizationSsoEnforcementRequest(organizationSsoEnforcementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoPutOrganizationSsoEnforcement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoPutOrganizationSsoEnforcementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationSsoEnforcementRequest** | [**OrganizationSsoEnforcementRequest**](OrganizationSsoEnforcementRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationSsoTestOrganizationSso

> OrganizationSsoPublic OrganizationSsoTestOrganizationSso(ctx).Execute()

Test Organization Sso



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
	resp, r, err := apiClient.OrganizationSsoAPI.OrganizationSsoTestOrganizationSso(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationSsoAPI.OrganizationSsoTestOrganizationSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationSsoTestOrganizationSso`: OrganizationSsoPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationSsoAPI.OrganizationSsoTestOrganizationSso`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationSsoTestOrganizationSsoRequest struct via the builder pattern


### Return type

[**OrganizationSsoPublic**](OrganizationSsoPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

