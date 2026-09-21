# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsAcceptCallerPendingMembership**](OrganizationsAPI.md#OrganizationsAcceptCallerPendingMembership) | **Post** /api/v1/organizations/me/pending-memberships/{organization_member_id}/accept | Accept Caller Pending Membership
[**OrganizationsCreateActiveOrganizationDomain**](OrganizationsAPI.md#OrganizationsCreateActiveOrganizationDomain) | **Post** /api/v1/organizations/me/domains | Create Active Organization Domain
[**OrganizationsCreateActiveOrganizationMember**](OrganizationsAPI.md#OrganizationsCreateActiveOrganizationMember) | **Post** /api/v1/organizations/me/members | Create Active Organization Member
[**OrganizationsCreateOrganization**](OrganizationsAPI.md#OrganizationsCreateOrganization) | **Post** /api/v1/organizations | Create Organization
[**OrganizationsDeclineCallerPendingMembership**](OrganizationsAPI.md#OrganizationsDeclineCallerPendingMembership) | **Post** /api/v1/organizations/me/pending-memberships/{organization_member_id}/decline | Decline Caller Pending Membership
[**OrganizationsDeleteActiveOrganizationDomain**](OrganizationsAPI.md#OrganizationsDeleteActiveOrganizationDomain) | **Delete** /api/v1/organizations/me/domains/{organization_domain_id} | Delete Active Organization Domain
[**OrganizationsGetActiveOrganizationContext**](OrganizationsAPI.md#OrganizationsGetActiveOrganizationContext) | **Get** /api/v1/organizations/me | Get Active Organization Context
[**OrganizationsInviteActiveOrganizationMember**](OrganizationsAPI.md#OrganizationsInviteActiveOrganizationMember) | **Post** /api/v1/organizations/me/member-invitations | Invite Active Organization Member
[**OrganizationsListActiveOrganizationDomains**](OrganizationsAPI.md#OrganizationsListActiveOrganizationDomains) | **Get** /api/v1/organizations/me/domains | List Active Organization Domains
[**OrganizationsListActiveOrganizationMembers**](OrganizationsAPI.md#OrganizationsListActiveOrganizationMembers) | **Get** /api/v1/organizations/me/members | List Active Organization Members
[**OrganizationsListCallerOrganizationMemberships**](OrganizationsAPI.md#OrganizationsListCallerOrganizationMemberships) | **Get** /api/v1/organizations/me/memberships | List Caller Organization Memberships
[**OrganizationsListCallerPendingMemberships**](OrganizationsAPI.md#OrganizationsListCallerPendingMemberships) | **Get** /api/v1/organizations/me/pending-memberships | List Caller Pending Memberships
[**OrganizationsRemoveActiveOrganizationMember**](OrganizationsAPI.md#OrganizationsRemoveActiveOrganizationMember) | **Delete** /api/v1/organizations/me/members/{organization_member_id} | Remove Active Organization Member
[**OrganizationsRevokeActiveOrganizationMemberInvitation**](OrganizationsAPI.md#OrganizationsRevokeActiveOrganizationMemberInvitation) | **Delete** /api/v1/organizations/me/member-invitations/{invitation_id} | Revoke Active Organization Member Invitation
[**OrganizationsSwitchActiveOrganization**](OrganizationsAPI.md#OrganizationsSwitchActiveOrganization) | **Post** /api/v1/organizations/me/switch | Switch Active Organization
[**OrganizationsUpdateActiveOrganization**](OrganizationsAPI.md#OrganizationsUpdateActiveOrganization) | **Patch** /api/v1/organizations/me | Update Active Organization
[**OrganizationsUpdateActiveOrganizationDomain**](OrganizationsAPI.md#OrganizationsUpdateActiveOrganizationDomain) | **Patch** /api/v1/organizations/me/domains/{organization_domain_id} | Update Active Organization Domain
[**OrganizationsUpdateActiveOrganizationMember**](OrganizationsAPI.md#OrganizationsUpdateActiveOrganizationMember) | **Patch** /api/v1/organizations/me/members/{organization_member_id} | Update Active Organization Member
[**OrganizationsVerifyActiveOrganizationDomain**](OrganizationsAPI.md#OrganizationsVerifyActiveOrganizationDomain) | **Post** /api/v1/organizations/me/domains/{organization_domain_id}/verify | Verify Active Organization Domain



## OrganizationsAcceptCallerPendingMembership

> AcceptInvitationResultPublic OrganizationsAcceptCallerPendingMembership(ctx, organizationMemberId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsAcceptCallerPendingMembership(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsAcceptCallerPendingMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsAcceptCallerPendingMembership`: AcceptInvitationResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsAcceptCallerPendingMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsAcceptCallerPendingMembershipRequest struct via the builder pattern


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


## OrganizationsCreateActiveOrganizationDomain

> OrganizationDomainPublic OrganizationsCreateActiveOrganizationDomain(ctx).OrganizationDomainCreateRequest(organizationDomainCreateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsCreateActiveOrganizationDomain(context.Background()).OrganizationDomainCreateRequest(organizationDomainCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsCreateActiveOrganizationDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsCreateActiveOrganizationDomain`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsCreateActiveOrganizationDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsCreateActiveOrganizationDomainRequest struct via the builder pattern


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


## OrganizationsCreateActiveOrganizationMember

> ActiveOrganizationMemberCreateResultPublic OrganizationsCreateActiveOrganizationMember(ctx).ActiveOrganizationMemberCreateRequest(activeOrganizationMemberCreateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsCreateActiveOrganizationMember(context.Background()).ActiveOrganizationMemberCreateRequest(activeOrganizationMemberCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsCreateActiveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsCreateActiveOrganizationMember`: ActiveOrganizationMemberCreateResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsCreateActiveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsCreateActiveOrganizationMemberRequest struct via the builder pattern


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


## OrganizationsCreateOrganization

> OrganizationPublic OrganizationsCreateOrganization(ctx).OrganizationCreateRequest(organizationCreateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsCreateOrganization(context.Background()).OrganizationCreateRequest(organizationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsCreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsCreateOrganization`: OrganizationPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsCreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsCreateOrganizationRequest struct via the builder pattern


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


## OrganizationsDeclineCallerPendingMembership

> Message OrganizationsDeclineCallerPendingMembership(ctx, organizationMemberId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsDeclineCallerPendingMembership(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsDeclineCallerPendingMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsDeclineCallerPendingMembership`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsDeclineCallerPendingMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsDeclineCallerPendingMembershipRequest struct via the builder pattern


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


## OrganizationsDeleteActiveOrganizationDomain

> Message OrganizationsDeleteActiveOrganizationDomain(ctx, organizationDomainId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsDeleteActiveOrganizationDomain(context.Background(), organizationDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsDeleteActiveOrganizationDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsDeleteActiveOrganizationDomain`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsDeleteActiveOrganizationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsDeleteActiveOrganizationDomainRequest struct via the builder pattern


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


## OrganizationsGetActiveOrganizationContext

> OrganizationMembershipContextPublic OrganizationsGetActiveOrganizationContext(ctx).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsGetActiveOrganizationContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsGetActiveOrganizationContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsGetActiveOrganizationContext`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsGetActiveOrganizationContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsGetActiveOrganizationContextRequest struct via the builder pattern


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


## OrganizationsInviteActiveOrganizationMember

> InviteOrganizationMemberResultPublic OrganizationsInviteActiveOrganizationMember(ctx).InviteOrganizationMemberRequest(inviteOrganizationMemberRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsInviteActiveOrganizationMember(context.Background()).InviteOrganizationMemberRequest(inviteOrganizationMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsInviteActiveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsInviteActiveOrganizationMember`: InviteOrganizationMemberResultPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsInviteActiveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsInviteActiveOrganizationMemberRequest struct via the builder pattern


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


## OrganizationsListActiveOrganizationDomains

> OrganizationDomainsPublic OrganizationsListActiveOrganizationDomains(ctx).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsListActiveOrganizationDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsListActiveOrganizationDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsListActiveOrganizationDomains`: OrganizationDomainsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsListActiveOrganizationDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsListActiveOrganizationDomainsRequest struct via the builder pattern


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


## OrganizationsListActiveOrganizationMembers

> ActiveOrganizationMembersPublic OrganizationsListActiveOrganizationMembers(ctx).Skip(skip).Limit(limit).Search(search).Execute()

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
	search := "search_example" // string | Narrow to members whose name or email contains this text, case-insensitively. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsListActiveOrganizationMembers(context.Background()).Skip(skip).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsListActiveOrganizationMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsListActiveOrganizationMembers`: ActiveOrganizationMembersPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsListActiveOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsListActiveOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Narrow to members whose name or email contains this text, case-insensitively. | 

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


## OrganizationsListCallerOrganizationMemberships

> CallerOrganizationMembershipsPublic OrganizationsListCallerOrganizationMemberships(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsListCallerOrganizationMemberships(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsListCallerOrganizationMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsListCallerOrganizationMemberships`: CallerOrganizationMembershipsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsListCallerOrganizationMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsListCallerOrganizationMembershipsRequest struct via the builder pattern


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


## OrganizationsListCallerPendingMemberships

> PendingOrganizationInvitationsPublic OrganizationsListCallerPendingMemberships(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsListCallerPendingMemberships(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsListCallerPendingMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsListCallerPendingMemberships`: PendingOrganizationInvitationsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsListCallerPendingMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsListCallerPendingMembershipsRequest struct via the builder pattern


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


## OrganizationsRemoveActiveOrganizationMember

> Message OrganizationsRemoveActiveOrganizationMember(ctx, organizationMemberId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsRemoveActiveOrganizationMember(context.Background(), organizationMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsRemoveActiveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsRemoveActiveOrganizationMember`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsRemoveActiveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsRemoveActiveOrganizationMemberRequest struct via the builder pattern


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


## OrganizationsRevokeActiveOrganizationMemberInvitation

> Message OrganizationsRevokeActiveOrganizationMemberInvitation(ctx, invitationId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsRevokeActiveOrganizationMemberInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsRevokeActiveOrganizationMemberInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsRevokeActiveOrganizationMemberInvitation`: Message
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsRevokeActiveOrganizationMemberInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsRevokeActiveOrganizationMemberInvitationRequest struct via the builder pattern


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


## OrganizationsSwitchActiveOrganization

> OrganizationMembershipContextPublic OrganizationsSwitchActiveOrganization(ctx).SwitchActiveOrganizationRequest(switchActiveOrganizationRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsSwitchActiveOrganization(context.Background()).SwitchActiveOrganizationRequest(switchActiveOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsSwitchActiveOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsSwitchActiveOrganization`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsSwitchActiveOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsSwitchActiveOrganizationRequest struct via the builder pattern


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


## OrganizationsUpdateActiveOrganization

> OrganizationMembershipContextPublic OrganizationsUpdateActiveOrganization(ctx).ActiveOrganizationUpdateRequest(activeOrganizationUpdateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsUpdateActiveOrganization(context.Background()).ActiveOrganizationUpdateRequest(activeOrganizationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsUpdateActiveOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsUpdateActiveOrganization`: OrganizationMembershipContextPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsUpdateActiveOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsUpdateActiveOrganizationRequest struct via the builder pattern


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


## OrganizationsUpdateActiveOrganizationDomain

> OrganizationDomainPublic OrganizationsUpdateActiveOrganizationDomain(ctx, organizationDomainId).OrganizationDomainUpdateRequest(organizationDomainUpdateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsUpdateActiveOrganizationDomain(context.Background(), organizationDomainId).OrganizationDomainUpdateRequest(organizationDomainUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsUpdateActiveOrganizationDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsUpdateActiveOrganizationDomain`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsUpdateActiveOrganizationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsUpdateActiveOrganizationDomainRequest struct via the builder pattern


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


## OrganizationsUpdateActiveOrganizationMember

> ActiveOrganizationMemberPublic OrganizationsUpdateActiveOrganizationMember(ctx, organizationMemberId).ActiveOrganizationMemberUpdateRequest(activeOrganizationMemberUpdateRequest).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsUpdateActiveOrganizationMember(context.Background(), organizationMemberId).ActiveOrganizationMemberUpdateRequest(activeOrganizationMemberUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsUpdateActiveOrganizationMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsUpdateActiveOrganizationMember`: ActiveOrganizationMemberPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsUpdateActiveOrganizationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsUpdateActiveOrganizationMemberRequest struct via the builder pattern


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


## OrganizationsVerifyActiveOrganizationDomain

> OrganizationDomainPublic OrganizationsVerifyActiveOrganizationDomain(ctx, organizationDomainId).Execute()

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
	resp, r, err := apiClient.OrganizationsAPI.OrganizationsVerifyActiveOrganizationDomain(context.Background(), organizationDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsVerifyActiveOrganizationDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsVerifyActiveOrganizationDomain`: OrganizationDomainPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsVerifyActiveOrganizationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsVerifyActiveOrganizationDomainRequest struct via the builder pattern


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

