# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticatePasskeyV1AuthWebauthnAuthenticatePost**](AuthAPI.md#AuthenticatePasskeyV1AuthWebauthnAuthenticatePost) | **Post** /v1/auth/webauthn/authenticate | Authenticate Passkey
[**AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost**](AuthAPI.md#AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost) | **Post** /v1/auth/webauthn/authenticate/options | Authentication Options
[**AuthorizeV1AuthOauthProviderAuthorizeGet**](AuthAPI.md#AuthorizeV1AuthOauthProviderAuthorizeGet) | **Get** /v1/auth/oauth/{provider}/authorize | Authorize
[**CallbackV1AuthOauthProviderCallbackPost**](AuthAPI.md#CallbackV1AuthOauthProviderCallbackPost) | **Post** /v1/auth/oauth/{provider}/callback | Callback
[**ConfirmResetV1AuthPasswordResetConfirmPost**](AuthAPI.md#ConfirmResetV1AuthPasswordResetConfirmPost) | **Post** /v1/auth/password/reset/confirm | Confirm Reset
[**CreateSessionV1AuthSessionPost**](AuthAPI.md#CreateSessionV1AuthSessionPost) | **Post** /v1/auth/session | Create Session
[**DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete**](AuthAPI.md#DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete) | **Delete** /v1/auth/webauthn/credentials/{credential_id} | Delete Passkey
[**DeleteSessionV1AuthSessionDelete**](AuthAPI.md#DeleteSessionV1AuthSessionDelete) | **Delete** /v1/auth/session | Delete Session
[**ListPasskeysV1AuthWebauthnCredentialsGet**](AuthAPI.md#ListPasskeysV1AuthWebauthnCredentialsGet) | **Get** /v1/auth/webauthn/credentials | List Passkeys
[**RegisterPasskeyV1AuthWebauthnRegisterPost**](AuthAPI.md#RegisterPasskeyV1AuthWebauthnRegisterPost) | **Post** /v1/auth/webauthn/register | Register Passkey
[**RegistrationOptionsV1AuthWebauthnRegisterOptionsPost**](AuthAPI.md#RegistrationOptionsV1AuthWebauthnRegisterOptionsPost) | **Post** /v1/auth/webauthn/register/options | Registration Options
[**RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch**](AuthAPI.md#RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch) | **Patch** /v1/auth/webauthn/credentials/{credential_id} | Rename Passkey
[**RequestResetV1AuthPasswordResetPost**](AuthAPI.md#RequestResetV1AuthPasswordResetPost) | **Post** /v1/auth/password/reset | Request Reset
[**ResendVerificationV1AuthResendVerificationPost**](AuthAPI.md#ResendVerificationV1AuthResendVerificationPost) | **Post** /v1/auth/resend-verification | Resend Verification
[**SetDashboardPasswordV1AuthPasswordPut**](AuthAPI.md#SetDashboardPasswordV1AuthPasswordPut) | **Put** /v1/auth/password | Set Dashboard Password
[**SignupV1AuthSignupPost**](AuthAPI.md#SignupV1AuthSignupPost) | **Post** /v1/auth/signup | Signup
[**VerifyEmailRouteV1AuthVerifyEmailPost**](AuthAPI.md#VerifyEmailRouteV1AuthVerifyEmailPost) | **Post** /v1/auth/verify-email | Verify Email Route



## AuthenticatePasskeyV1AuthWebauthnAuthenticatePost

> PasskeySessionResponse AuthenticatePasskeyV1AuthWebauthnAuthenticatePost(ctx).AuthenticatePasskeyRequest(authenticatePasskeyRequest).Execute()

Authenticate Passkey



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
	authenticatePasskeyRequest := *openapiclient.NewAuthenticatePasskeyRequest(map[string]interface{}{"key": interface{}(123)}) // AuthenticatePasskeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthenticatePasskeyV1AuthWebauthnAuthenticatePost(context.Background()).AuthenticatePasskeyRequest(authenticatePasskeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthenticatePasskeyV1AuthWebauthnAuthenticatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticatePasskeyV1AuthWebauthnAuthenticatePost`: PasskeySessionResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthenticatePasskeyV1AuthWebauthnAuthenticatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatePasskeyV1AuthWebauthnAuthenticatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatePasskeyRequest** | [**AuthenticatePasskeyRequest**](AuthenticatePasskeyRequest.md) |  | 

### Return type

[**PasskeySessionResponse**](PasskeySessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost

> map[string]interface{} AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost(ctx).Execute()

Authentication Options



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
	resp, r, err := apiClient.AuthAPI.AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationOptionsV1AuthWebauthnAuthenticateOptionsPostRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeV1AuthOauthProviderAuthorizeGet

> AuthorizeResponse AuthorizeV1AuthOauthProviderAuthorizeGet(ctx, provider).Execute()

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
	provider := "provider_example" // string | Which OAuth provider to sign in with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthorizeV1AuthOauthProviderAuthorizeGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthorizeV1AuthOauthProviderAuthorizeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeV1AuthOauthProviderAuthorizeGet`: AuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthorizeV1AuthOauthProviderAuthorizeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Which OAuth provider to sign in with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeV1AuthOauthProviderAuthorizeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizeResponse**](AuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallbackV1AuthOauthProviderCallbackPost

> OAuthSessionResponse CallbackV1AuthOauthProviderCallbackPost(ctx, provider).OAuthCallbackRequest(oAuthCallbackRequest).Execute()

Callback



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
	provider := "provider_example" // string | Which OAuth provider to sign in with.
	oAuthCallbackRequest := *openapiclient.NewOAuthCallbackRequest("Code_example") // OAuthCallbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.CallbackV1AuthOauthProviderCallbackPost(context.Background(), provider).OAuthCallbackRequest(oAuthCallbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.CallbackV1AuthOauthProviderCallbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CallbackV1AuthOauthProviderCallbackPost`: OAuthSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.CallbackV1AuthOauthProviderCallbackPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Which OAuth provider to sign in with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCallbackV1AuthOauthProviderCallbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthCallbackRequest** | [**OAuthCallbackRequest**](OAuthCallbackRequest.md) |  | 

### Return type

[**OAuthSessionResponse**](OAuthSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmResetV1AuthPasswordResetConfirmPost

> ConfirmResetV1AuthPasswordResetConfirmPost(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Confirm Reset



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
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("NewPassword_example", "Token_example") // ResetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.ConfirmResetV1AuthPasswordResetConfirmPost(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ConfirmResetV1AuthPasswordResetConfirmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmResetV1AuthPasswordResetConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	createSessionRequest := *openapiclient.NewCreateSessionRequest() // CreateSessionRequest | 

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


## DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete

> DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete(ctx, credentialId).Execute()

Delete Passkey



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
	credentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthAPI.DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.DeletePasskeyV1AuthWebauthnCredentialsCredentialIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasskeyV1AuthWebauthnCredentialsCredentialIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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


## ListPasskeysV1AuthWebauthnCredentialsGet

> WebAuthnCredentialsPublic ListPasskeysV1AuthWebauthnCredentialsGet(ctx).Execute()

List Passkeys



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
	resp, r, err := apiClient.AuthAPI.ListPasskeysV1AuthWebauthnCredentialsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ListPasskeysV1AuthWebauthnCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPasskeysV1AuthWebauthnCredentialsGet`: WebAuthnCredentialsPublic
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.ListPasskeysV1AuthWebauthnCredentialsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPasskeysV1AuthWebauthnCredentialsGetRequest struct via the builder pattern


### Return type

[**WebAuthnCredentialsPublic**](WebAuthnCredentialsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPasskeyV1AuthWebauthnRegisterPost

> WebAuthnCredentialPublic RegisterPasskeyV1AuthWebauthnRegisterPost(ctx).RegisterPasskeyRequest(registerPasskeyRequest).Execute()

Register Passkey



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
	registerPasskeyRequest := *openapiclient.NewRegisterPasskeyRequest(map[string]interface{}{"key": interface{}(123)}) // RegisterPasskeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RegisterPasskeyV1AuthWebauthnRegisterPost(context.Background()).RegisterPasskeyRequest(registerPasskeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RegisterPasskeyV1AuthWebauthnRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterPasskeyV1AuthWebauthnRegisterPost`: WebAuthnCredentialPublic
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RegisterPasskeyV1AuthWebauthnRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPasskeyV1AuthWebauthnRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerPasskeyRequest** | [**RegisterPasskeyRequest**](RegisterPasskeyRequest.md) |  | 

### Return type

[**WebAuthnCredentialPublic**](WebAuthnCredentialPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationOptionsV1AuthWebauthnRegisterOptionsPost

> map[string]interface{} RegistrationOptionsV1AuthWebauthnRegisterOptionsPost(ctx).Execute()

Registration Options



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
	resp, r, err := apiClient.AuthAPI.RegistrationOptionsV1AuthWebauthnRegisterOptionsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RegistrationOptionsV1AuthWebauthnRegisterOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistrationOptionsV1AuthWebauthnRegisterOptionsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RegistrationOptionsV1AuthWebauthnRegisterOptionsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationOptionsV1AuthWebauthnRegisterOptionsPostRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch

> WebAuthnCredentialPublic RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch(ctx, credentialId).WebAuthnCredentialUpdate(webAuthnCredentialUpdate).Execute()

Rename Passkey



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
	credentialId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	webAuthnCredentialUpdate := *openapiclient.NewWebAuthnCredentialUpdate("Name_example") // WebAuthnCredentialUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch(context.Background(), credentialId).WebAuthnCredentialUpdate(webAuthnCredentialUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch`: WebAuthnCredentialPublic
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenamePasskeyV1AuthWebauthnCredentialsCredentialIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webAuthnCredentialUpdate** | [**WebAuthnCredentialUpdate**](WebAuthnCredentialUpdate.md) |  | 

### Return type

[**WebAuthnCredentialPublic**](WebAuthnCredentialPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestResetV1AuthPasswordResetPost

> RequestPasswordResetResponse RequestResetV1AuthPasswordResetPost(ctx).RequestPasswordResetRequest(requestPasswordResetRequest).Execute()

Request Reset



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
	requestPasswordResetRequest := *openapiclient.NewRequestPasswordResetRequest("Email_example") // RequestPasswordResetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.RequestResetV1AuthPasswordResetPost(context.Background()).RequestPasswordResetRequest(requestPasswordResetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.RequestResetV1AuthPasswordResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResetV1AuthPasswordResetPost`: RequestPasswordResetResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.RequestResetV1AuthPasswordResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestResetV1AuthPasswordResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestPasswordResetRequest** | [**RequestPasswordResetRequest**](RequestPasswordResetRequest.md) |  | 

### Return type

[**RequestPasswordResetResponse**](RequestPasswordResetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendVerificationV1AuthResendVerificationPost

> ResendVerificationResponse ResendVerificationV1AuthResendVerificationPost(ctx).ResendVerificationRequest(resendVerificationRequest).Execute()

Resend Verification



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
	resendVerificationRequest := *openapiclient.NewResendVerificationRequest("Email_example") // ResendVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.ResendVerificationV1AuthResendVerificationPost(context.Background()).ResendVerificationRequest(resendVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ResendVerificationV1AuthResendVerificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendVerificationV1AuthResendVerificationPost`: ResendVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.ResendVerificationV1AuthResendVerificationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendVerificationV1AuthResendVerificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendVerificationRequest** | [**ResendVerificationRequest**](ResendVerificationRequest.md) |  | 

### Return type

[**ResendVerificationResponse**](ResendVerificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDashboardPasswordV1AuthPasswordPut

> PasswordResponse SetDashboardPasswordV1AuthPasswordPut(ctx).SetPasswordRequest(setPasswordRequest).Execute()

Set Dashboard Password



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
	setPasswordRequest := *openapiclient.NewSetPasswordRequest("NewPassword_example") // SetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.SetDashboardPasswordV1AuthPasswordPut(context.Background()).SetPasswordRequest(setPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.SetDashboardPasswordV1AuthPasswordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDashboardPasswordV1AuthPasswordPut`: PasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.SetDashboardPasswordV1AuthPasswordPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDashboardPasswordV1AuthPasswordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPasswordRequest** | [**SetPasswordRequest**](SetPasswordRequest.md) |  | 

### Return type

[**PasswordResponse**](PasswordResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupV1AuthSignupPost

> SignupResponse SignupV1AuthSignupPost(ctx).SignupRequest(signupRequest).Execute()

Signup



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
	signupRequest := *openapiclient.NewSignupRequest("Email_example", "Password_example") // SignupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.SignupV1AuthSignupPost(context.Background()).SignupRequest(signupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.SignupV1AuthSignupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignupV1AuthSignupPost`: SignupResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.SignupV1AuthSignupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupV1AuthSignupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupRequest** | [**SignupRequest**](SignupRequest.md) |  | 

### Return type

[**SignupResponse**](SignupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailRouteV1AuthVerifyEmailPost

> VerifyEmailResponse VerifyEmailRouteV1AuthVerifyEmailPost(ctx).VerifyEmailRequest(verifyEmailRequest).Execute()

Verify Email Route



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
	verifyEmailRequest := *openapiclient.NewVerifyEmailRequest("Token_example") // VerifyEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.VerifyEmailRouteV1AuthVerifyEmailPost(context.Background()).VerifyEmailRequest(verifyEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.VerifyEmailRouteV1AuthVerifyEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyEmailRouteV1AuthVerifyEmailPost`: VerifyEmailResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.VerifyEmailRouteV1AuthVerifyEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailRouteV1AuthVerifyEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyEmailRequest** | [**VerifyEmailRequest**](VerifyEmailRequest.md) |  | 

### Return type

[**VerifyEmailResponse**](VerifyEmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

