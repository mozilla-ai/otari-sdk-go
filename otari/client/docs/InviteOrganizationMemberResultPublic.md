# InviteOrganizationMemberResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptLink** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Email** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**InvitationId** | **string** |  | 
**MailSent** | **bool** | Whether the invitation email was actually dispatched. False when mail is not configured, or the send itself failed; accept_link is set either way, so the operator can share it themselves rather than the invitation being a dead end. | 
**OrganizationMemberId** | **string** |  | 
**Role** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] [default to "invited"]

## Methods

### NewInviteOrganizationMemberResultPublic

`func NewInviteOrganizationMemberResultPublic(acceptLink string, createdAt time.Time, email string, expiresAt time.Time, invitationId string, mailSent bool, organizationMemberId string, role string, ) *InviteOrganizationMemberResultPublic`

NewInviteOrganizationMemberResultPublic instantiates a new InviteOrganizationMemberResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteOrganizationMemberResultPublicWithDefaults

`func NewInviteOrganizationMemberResultPublicWithDefaults() *InviteOrganizationMemberResultPublic`

NewInviteOrganizationMemberResultPublicWithDefaults instantiates a new InviteOrganizationMemberResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptLink

`func (o *InviteOrganizationMemberResultPublic) GetAcceptLink() string`

GetAcceptLink returns the AcceptLink field if non-nil, zero value otherwise.

### GetAcceptLinkOk

`func (o *InviteOrganizationMemberResultPublic) GetAcceptLinkOk() (*string, bool)`

GetAcceptLinkOk returns a tuple with the AcceptLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptLink

`func (o *InviteOrganizationMemberResultPublic) SetAcceptLink(v string)`

SetAcceptLink sets AcceptLink field to given value.


### GetCreatedAt

`func (o *InviteOrganizationMemberResultPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InviteOrganizationMemberResultPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InviteOrganizationMemberResultPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *InviteOrganizationMemberResultPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteOrganizationMemberResultPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteOrganizationMemberResultPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpiresAt

`func (o *InviteOrganizationMemberResultPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InviteOrganizationMemberResultPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InviteOrganizationMemberResultPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetInvitationId

`func (o *InviteOrganizationMemberResultPublic) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *InviteOrganizationMemberResultPublic) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *InviteOrganizationMemberResultPublic) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.


### GetMailSent

`func (o *InviteOrganizationMemberResultPublic) GetMailSent() bool`

GetMailSent returns the MailSent field if non-nil, zero value otherwise.

### GetMailSentOk

`func (o *InviteOrganizationMemberResultPublic) GetMailSentOk() (*bool, bool)`

GetMailSentOk returns a tuple with the MailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSent

`func (o *InviteOrganizationMemberResultPublic) SetMailSent(v bool)`

SetMailSent sets MailSent field to given value.


### GetOrganizationMemberId

`func (o *InviteOrganizationMemberResultPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *InviteOrganizationMemberResultPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *InviteOrganizationMemberResultPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.


### GetRole

`func (o *InviteOrganizationMemberResultPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteOrganizationMemberResultPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteOrganizationMemberResultPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *InviteOrganizationMemberResultPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InviteOrganizationMemberResultPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InviteOrganizationMemberResultPublic) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InviteOrganizationMemberResultPublic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


