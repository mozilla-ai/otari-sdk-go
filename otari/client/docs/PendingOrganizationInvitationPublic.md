# PendingOrganizationInvitationPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Email** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**InvitationId** | **string** |  | 
**OrganizationId** | **string** |  | 
**OrganizationMemberId** | **string** |  | 
**OrganizationName** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewPendingOrganizationInvitationPublic

`func NewPendingOrganizationInvitationPublic(createdAt time.Time, email string, expiresAt time.Time, invitationId string, organizationId string, organizationMemberId string, organizationName string, role string, ) *PendingOrganizationInvitationPublic`

NewPendingOrganizationInvitationPublic instantiates a new PendingOrganizationInvitationPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingOrganizationInvitationPublicWithDefaults

`func NewPendingOrganizationInvitationPublicWithDefaults() *PendingOrganizationInvitationPublic`

NewPendingOrganizationInvitationPublicWithDefaults instantiates a new PendingOrganizationInvitationPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PendingOrganizationInvitationPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PendingOrganizationInvitationPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PendingOrganizationInvitationPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *PendingOrganizationInvitationPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PendingOrganizationInvitationPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PendingOrganizationInvitationPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpiresAt

`func (o *PendingOrganizationInvitationPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PendingOrganizationInvitationPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PendingOrganizationInvitationPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetInvitationId

`func (o *PendingOrganizationInvitationPublic) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *PendingOrganizationInvitationPublic) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *PendingOrganizationInvitationPublic) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.


### GetOrganizationId

`func (o *PendingOrganizationInvitationPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PendingOrganizationInvitationPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PendingOrganizationInvitationPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationMemberId

`func (o *PendingOrganizationInvitationPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *PendingOrganizationInvitationPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *PendingOrganizationInvitationPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.


### GetOrganizationName

`func (o *PendingOrganizationInvitationPublic) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *PendingOrganizationInvitationPublic) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *PendingOrganizationInvitationPublic) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetRole

`func (o *PendingOrganizationInvitationPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PendingOrganizationInvitationPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PendingOrganizationInvitationPublic) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


