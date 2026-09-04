# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost**](OrganizationsAPI.md#AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost) | **Post** /v1/organizations/me/pending-memberships/{organization_member_id}/accept | Accept Caller Pending Membership
[**CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost**](OrganizationsAPI.md#CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost) | **Post** /v1/organizations/me/domains | Create Active Organization Domain
[**CreateActiveOrganizationMemberV1OrganizationsMeMembersPost**](OrganizationsAPI.md#CreateActiveOrganizationMemberV1OrganizationsMeMembersPost) | **Post** /v1/organizations/me/members | Create Active Organization Member
[**CreateOrganizationV1OrganizationsPost**](OrganizationsAPI.md#CreateOrganizationV1OrganizationsPost) | **Post** /v1/organizations | Create Organization
[**DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost**](OrganizationsAPI.md#DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost) | **Post** /v1/organizations/me/pending-memberships/{organization_member_id}/decline | Decline Caller Pending Membership
[**DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete**](OrganizationsAPI.md#DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete) | **Delete** /v1/organizations/me/domains/{organization_domain_id} | Delete Active Organization Domain
[**GetActiveOrganizationContextV1OrganizationsMeGet**](OrganizationsAPI.md#GetActiveOrganizationContextV1OrganizationsMeGet) | **Get** /v1/organizations/me | Get Active Organization Context
[**InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost**](OrganizationsAPI.md#InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost) | **Post** /v1/organizations/me/member-invitations | Invite Active Organization Member
[**ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet**](OrganizationsAPI.md#ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet) | **Get** /v1/organizations/me/domains | List Active Organization Domains
[**ListActiveOrganizationMembersV1OrganizationsMeMembersGet**](OrganizationsAPI.md#ListActiveOrganizationMembersV1OrganizationsMeMembersGet) | **Get** /v1/organizations/me/members | List Active Organization Members
[**ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet**](OrganizationsAPI.md#ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet) | **Get** /v1/organizations/me/memberships | List Caller Organization Memberships
[**ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet**](OrganizationsAPI.md#ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet) | **Get** /v1/organizations/me/pending-memberships | List Caller Pending Memberships
[**RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete**](OrganizationsAPI.md#RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete) | **Delete** /v1/organizations/me/members/{organization_member_id} | Remove Active Organization Member
[**RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete**](OrganizationsAPI.md#RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete) | **Delete** /v1/organizations/me/member-invitations/{invitation_id} | Revoke Active Organization Member Invitation
[**SwitchActiveOrganizationV1OrganizationsMeSwitchPost**](OrganizationsAPI.md#SwitchActiveOrganizationV1OrganizationsMeSwitchPost) | **Post** /v1/organizations/me/switch | Switch Active Organization
[**UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch**](OrganizationsAPI.md#UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch) | **Patch** /v1/organizations/me/domains/{organization_domain_id} | Update Active Organization Domain
[**UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch**](OrganizationsAPI.md#UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch) | **Patch** /v1/organizations/me/members/{organization_member_id} | Update Active Organization Member
[**UpdateActiveOrganizationV1OrganizationsMePatch**](OrganizationsAPI.md#UpdateActiveOrganizationV1OrganizationsMePatch) | **Patch** /v1/organizations/me | Update Active Organization
[**VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost**](OrganizationsAPI.md#VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost) | **Post** /v1/organizations/me/domains/{organization_domain_id}/verify | Verify Active Organization Domain



## AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost

> AcceptInvitationResultPublic AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost(ctx, organizationMemberId).Execute()

Accept Caller Pending Membership



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
	organizationMemberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost`: AcceptInvitationResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AcceptInvitationResultPublic**](AcceptInvitationResultPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost

> OrganizationDomainPublic CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost(ctx).OrganizationDomainCreateRequest(organizationDomainCreateRequest).Execute()

Create Active Organization Domain



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
	organizationDomainCreateRequest := *openapiclient.NewOrganizationDomainCreateRequest("Domain_example") // OrganizationDomainCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost(context.Background()).OrganizationDomainCreateRequest(organizationDomainCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateActiveOrganizationDomainV1OrganizationsMeDomainsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActiveOrganizationDomainV1OrganizationsMeDomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationDomainCreateRequest** | [**OrganizationDomainCreateRequest**](OrganizationDomainCreateRequest.md) |  | 

### Return type

[**OrganizationDomainPublic**](OrganizationDomainPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActiveOrganizationMemberV1OrganizationsMeMembersPost

> ActiveOrganizationMemberCreateResultPublic CreateActiveOrganizationMemberV1OrganizationsMeMembersPost(ctx).ActiveOrganizationMemberCreateRequest(activeOrganizationMemberCreateRequest).Execute()

Create Active Organization Member



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
	activeOrganizationMemberCreateRequest := *openapiclient.NewActiveOrganizationMemberCreateRequest("Email_example") // ActiveOrganizationMemberCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateActiveOrganizationMemberV1OrganizationsMeMembersPost(context.Background()).ActiveOrganizationMemberCreateRequest(activeOrganizationMemberCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateActiveOrganizationMemberV1OrganizationsMeMembersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActiveOrganizationMemberV1OrganizationsMeMembersPost`: ActiveOrganizationMemberCreateResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateActiveOrganizationMemberV1OrganizationsMeMembersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActiveOrganizationMemberV1OrganizationsMeMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeOrganizationMemberCreateRequest** | [**ActiveOrganizationMemberCreateRequest**](ActiveOrganizationMemberCreateRequest.md) |  | 

### Return type

[**ActiveOrganizationMemberCreateResultPublic**](ActiveOrganizationMemberCreateResultPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationV1OrganizationsPost

> OrganizationPublic CreateOrganizationV1OrganizationsPost(ctx).OrganizationCreateRequest(organizationCreateRequest).Execute()

Create Organization



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
	organizationCreateRequest := *openapiclient.NewOrganizationCreateRequest("Name_example") // OrganizationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CreateOrganizationV1OrganizationsPost(context.Background()).OrganizationCreateRequest(organizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CreateOrganizationV1OrganizationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationV1OrganizationsPost`: OrganizationPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CreateOrganizationV1OrganizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationV1OrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationCreateRequest** | [**OrganizationCreateRequest**](OrganizationCreateRequest.md) |  | 

### Return type

[**OrganizationPublic**](OrganizationPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost

> Message DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost(ctx, organizationMemberId).Execute()

Decline Caller Pending Membership



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
	organizationMemberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineCallerPendingMembershipV1OrganizationsMePendingMembershipsOrganizationMemberIdDeclinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete

> Message DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete(ctx, organizationDomainId).Execute()

Delete Active Organization Domain



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
	organizationDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete(context.Background(), organizationDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveOrganizationContextV1OrganizationsMeGet

> OrganizationMembershipContextPublic GetActiveOrganizationContextV1OrganizationsMeGet(ctx).Execute()

Get Active Organization Context



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
	resp, r, err := apiClient.OrganizationsAPI.GetActiveOrganizationContextV1OrganizationsMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetActiveOrganizationContextV1OrganizationsMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveOrganizationContextV1OrganizationsMeGet`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetActiveOrganizationContextV1OrganizationsMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveOrganizationContextV1OrganizationsMeGetRequest struct via the builder pattern


### Return type

[**OrganizationMembershipContextPublic**](OrganizationMembershipContextPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost

> InviteOrganizationMemberResultPublic InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost(ctx).InviteOrganizationMemberRequest(inviteOrganizationMemberRequest).Execute()

Invite Active Organization Member



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
	inviteOrganizationMemberRequest := *openapiclient.NewInviteOrganizationMemberRequest("Email_example") // InviteOrganizationMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost(context.Background()).InviteOrganizationMemberRequest(inviteOrganizationMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost`: InviteOrganizationMemberResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.InviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteActiveOrganizationMemberV1OrganizationsMeMemberInvitationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteOrganizationMemberRequest** | [**InviteOrganizationMemberRequest**](InviteOrganizationMemberRequest.md) |  | 

### Return type

[**InviteOrganizationMemberResultPublic**](InviteOrganizationMemberResultPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet

> OrganizationDomainsPublic ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet(ctx).Execute()

List Active Organization Domains



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
	resp, r, err := apiClient.OrganizationsAPI.ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet`: OrganizationDomainsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListActiveOrganizationDomainsV1OrganizationsMeDomainsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListActiveOrganizationDomainsV1OrganizationsMeDomainsGetRequest struct via the builder pattern


### Return type

[**OrganizationDomainsPublic**](OrganizationDomainsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveOrganizationMembersV1OrganizationsMeMembersGet

> ActiveOrganizationMembersPublic ListActiveOrganizationMembersV1OrganizationsMeMembersGet(ctx).Skip(skip).Limit(limit).Execute()

List Active Organization Members



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListActiveOrganizationMembersV1OrganizationsMeMembersGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListActiveOrganizationMembersV1OrganizationsMeMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActiveOrganizationMembersV1OrganizationsMeMembersGet`: ActiveOrganizationMembersPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListActiveOrganizationMembersV1OrganizationsMeMembersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActiveOrganizationMembersV1OrganizationsMeMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**ActiveOrganizationMembersPublic**](ActiveOrganizationMembersPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet

> CallerOrganizationMembershipsPublic ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet(ctx).Skip(skip).Limit(limit).Execute()

List Caller Organization Memberships



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet`: CallerOrganizationMembershipsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCallerOrganizationMembershipsV1OrganizationsMeMembershipsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**CallerOrganizationMembershipsPublic**](CallerOrganizationMembershipsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet

> PendingOrganizationInvitationsPublic ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet(ctx).Skip(skip).Limit(limit).Execute()

List Caller Pending Memberships



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet`: PendingOrganizationInvitationsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCallerPendingMembershipsV1OrganizationsMePendingMembershipsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**PendingOrganizationInvitationsPublic**](PendingOrganizationInvitationsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete

> Message RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete(ctx, organizationMemberId).Execute()

Remove Active Organization Member



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
	organizationMemberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete

> Message RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete(ctx, invitationId).Execute()

Revoke Active Organization Member Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeActiveOrganizationMemberInvitationV1OrganizationsMeMemberInvitationsInvitationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Message**](Message.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchActiveOrganizationV1OrganizationsMeSwitchPost

> OrganizationMembershipContextPublic SwitchActiveOrganizationV1OrganizationsMeSwitchPost(ctx).SwitchActiveOrganizationRequest(switchActiveOrganizationRequest).Execute()

Switch Active Organization



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
	switchActiveOrganizationRequest := *openapiclient.NewSwitchActiveOrganizationRequest("OrganizationId_example") // SwitchActiveOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.SwitchActiveOrganizationV1OrganizationsMeSwitchPost(context.Background()).SwitchActiveOrganizationRequest(switchActiveOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.SwitchActiveOrganizationV1OrganizationsMeSwitchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwitchActiveOrganizationV1OrganizationsMeSwitchPost`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.SwitchActiveOrganizationV1OrganizationsMeSwitchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwitchActiveOrganizationV1OrganizationsMeSwitchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **switchActiveOrganizationRequest** | [**SwitchActiveOrganizationRequest**](SwitchActiveOrganizationRequest.md) |  | 

### Return type

[**OrganizationMembershipContextPublic**](OrganizationMembershipContextPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch

> OrganizationDomainPublic UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch(ctx, organizationDomainId).OrganizationDomainUpdateRequest(organizationDomainUpdateRequest).Execute()

Update Active Organization Domain



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
	organizationDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationDomainUpdateRequest := *openapiclient.NewOrganizationDomainUpdateRequest() // OrganizationDomainUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch(context.Background(), organizationDomainId).OrganizationDomainUpdateRequest(organizationDomainUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationDomainUpdateRequest** | [**OrganizationDomainUpdateRequest**](OrganizationDomainUpdateRequest.md) |  | 

### Return type

[**OrganizationDomainPublic**](OrganizationDomainPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch

> ActiveOrganizationMemberPublic UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch(ctx, organizationMemberId).ActiveOrganizationMemberUpdateRequest(activeOrganizationMemberUpdateRequest).Execute()

Update Active Organization Member



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
	organizationMemberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	activeOrganizationMemberUpdateRequest := *openapiclient.NewActiveOrganizationMemberUpdateRequest() // ActiveOrganizationMemberUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch(context.Background(), organizationMemberId).ActiveOrganizationMemberUpdateRequest(activeOrganizationMemberUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch`: ActiveOrganizationMemberPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveOrganizationMemberV1OrganizationsMeMembersOrganizationMemberIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeOrganizationMemberUpdateRequest** | [**ActiveOrganizationMemberUpdateRequest**](ActiveOrganizationMemberUpdateRequest.md) |  | 

### Return type

[**ActiveOrganizationMemberPublic**](ActiveOrganizationMemberPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveOrganizationV1OrganizationsMePatch

> OrganizationMembershipContextPublic UpdateActiveOrganizationV1OrganizationsMePatch(ctx).ActiveOrganizationUpdateRequest(activeOrganizationUpdateRequest).Execute()

Update Active Organization



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
	activeOrganizationUpdateRequest := *openapiclient.NewActiveOrganizationUpdateRequest("Name_example") // ActiveOrganizationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.UpdateActiveOrganizationV1OrganizationsMePatch(context.Background()).ActiveOrganizationUpdateRequest(activeOrganizationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateActiveOrganizationV1OrganizationsMePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActiveOrganizationV1OrganizationsMePatch`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateActiveOrganizationV1OrganizationsMePatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveOrganizationV1OrganizationsMePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeOrganizationUpdateRequest** | [**ActiveOrganizationUpdateRequest**](ActiveOrganizationUpdateRequest.md) |  | 

### Return type

[**OrganizationMembershipContextPublic**](OrganizationMembershipContextPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost

> OrganizationDomainPublic VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost(ctx, organizationDomainId).Execute()

Verify Active Organization Domain



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
	organizationDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost(context.Background(), organizationDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.VerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyActiveOrganizationDomainV1OrganizationsMeDomainsOrganizationDomainIdVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationDomainPublic**](OrganizationDomainPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

