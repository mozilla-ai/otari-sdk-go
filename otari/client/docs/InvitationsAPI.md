# \InvitationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvitationsAcceptInvitation**](InvitationsAPI.md#InvitationsAcceptInvitation) | **Post** /api/v1/invitations/accept | Accept Invitation
[**InvitationsValidateInvitation**](InvitationsAPI.md#InvitationsValidateInvitation) | **Post** /api/v1/invitations/validate | Validate Invitation



## InvitationsAcceptInvitation

> AcceptInvitationResultPublic InvitationsAcceptInvitation(ctx).AcceptInvitationRequest(acceptInvitationRequest).Execute()

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
	resp, r, err := apiClient.InvitationsAPI.InvitationsAcceptInvitation(context.Background()).AcceptInvitationRequest(acceptInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.InvitationsAcceptInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationsAcceptInvitation`: AcceptInvitationResultPublic
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.InvitationsAcceptInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsAcceptInvitationRequest struct via the builder pattern


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


## InvitationsValidateInvitation

> InvitationPreviewPublic InvitationsValidateInvitation(ctx).ValidateInvitationRequest(validateInvitationRequest).Execute()

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
	resp, r, err := apiClient.InvitationsAPI.InvitationsValidateInvitation(context.Background()).ValidateInvitationRequest(validateInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.InvitationsValidateInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitationsValidateInvitation`: InvitationPreviewPublic
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.InvitationsValidateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvitationsValidateInvitationRequest struct via the builder pattern


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

