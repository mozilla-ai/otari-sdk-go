# \InvitationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitationV1InvitationsAcceptPost**](InvitationsAPI.md#AcceptInvitationV1InvitationsAcceptPost) | **Post** /v1/invitations/accept | Accept Invitation
[**ValidateInvitationV1InvitationsValidatePost**](InvitationsAPI.md#ValidateInvitationV1InvitationsValidatePost) | **Post** /v1/invitations/validate | Validate Invitation



## AcceptInvitationV1InvitationsAcceptPost

> AcceptInvitationResultPublic AcceptInvitationV1InvitationsAcceptPost(ctx).AcceptInvitationRequest(acceptInvitationRequest).Execute()

Accept Invitation



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
	acceptInvitationRequest := *openapiclient.NewAcceptInvitationRequest("Token_example") // AcceptInvitationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.AcceptInvitationV1InvitationsAcceptPost(context.Background()).AcceptInvitationRequest(acceptInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.AcceptInvitationV1InvitationsAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptInvitationV1InvitationsAcceptPost`: AcceptInvitationResultPublic
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.AcceptInvitationV1InvitationsAcceptPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInvitationV1InvitationsAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptInvitationRequest** | [**AcceptInvitationRequest**](AcceptInvitationRequest.md) |  | 

### Return type

[**AcceptInvitationResultPublic**](AcceptInvitationResultPublic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateInvitationV1InvitationsValidatePost

> InvitationPreviewPublic ValidateInvitationV1InvitationsValidatePost(ctx).ValidateInvitationRequest(validateInvitationRequest).Execute()

Validate Invitation



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
	validateInvitationRequest := *openapiclient.NewValidateInvitationRequest("Token_example") // ValidateInvitationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.ValidateInvitationV1InvitationsValidatePost(context.Background()).ValidateInvitationRequest(validateInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.ValidateInvitationV1InvitationsValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateInvitationV1InvitationsValidatePost`: InvitationPreviewPublic
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.ValidateInvitationV1InvitationsValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateInvitationV1InvitationsValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateInvitationRequest** | [**ValidateInvitationRequest**](ValidateInvitationRequest.md) |  | 

### Return type

[**InvitationPreviewPublic**](InvitationPreviewPublic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

