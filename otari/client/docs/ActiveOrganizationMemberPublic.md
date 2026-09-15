# ActiveOrganizationMemberPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributionUserId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**InvitationId** | Pointer to **NullableString** |  | [optional] 
**OrganizationMemberId** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActiveOrganizationMemberPublic

`func NewActiveOrganizationMemberPublic(createdAt time.Time, role string, status string, ) *ActiveOrganizationMemberPublic`

NewActiveOrganizationMemberPublic instantiates a new ActiveOrganizationMemberPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveOrganizationMemberPublicWithDefaults

`func NewActiveOrganizationMemberPublicWithDefaults() *ActiveOrganizationMemberPublic`

NewActiveOrganizationMemberPublicWithDefaults instantiates a new ActiveOrganizationMemberPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributionUserId

`func (o *ActiveOrganizationMemberPublic) GetAttributionUserId() string`

GetAttributionUserId returns the AttributionUserId field if non-nil, zero value otherwise.

### GetAttributionUserIdOk

`func (o *ActiveOrganizationMemberPublic) GetAttributionUserIdOk() (*string, bool)`

GetAttributionUserIdOk returns a tuple with the AttributionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionUserId

`func (o *ActiveOrganizationMemberPublic) SetAttributionUserId(v string)`

SetAttributionUserId sets AttributionUserId field to given value.

### HasAttributionUserId

`func (o *ActiveOrganizationMemberPublic) HasAttributionUserId() bool`

HasAttributionUserId returns a boolean if a field has been set.

### SetAttributionUserIdNil

`func (o *ActiveOrganizationMemberPublic) SetAttributionUserIdNil(b bool)`

 SetAttributionUserIdNil sets the value for AttributionUserId to be an explicit nil

### UnsetAttributionUserId
`func (o *ActiveOrganizationMemberPublic) UnsetAttributionUserId()`

UnsetAttributionUserId ensures that no value is present for AttributionUserId, not even an explicit nil
### GetCreatedAt

`func (o *ActiveOrganizationMemberPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActiveOrganizationMemberPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActiveOrganizationMemberPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *ActiveOrganizationMemberPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ActiveOrganizationMemberPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ActiveOrganizationMemberPublic) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ActiveOrganizationMemberPublic) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ActiveOrganizationMemberPublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ActiveOrganizationMemberPublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *ActiveOrganizationMemberPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ActiveOrganizationMemberPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ActiveOrganizationMemberPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ActiveOrganizationMemberPublic) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ActiveOrganizationMemberPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ActiveOrganizationMemberPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetInvitationId

`func (o *ActiveOrganizationMemberPublic) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *ActiveOrganizationMemberPublic) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *ActiveOrganizationMemberPublic) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *ActiveOrganizationMemberPublic) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### SetInvitationIdNil

`func (o *ActiveOrganizationMemberPublic) SetInvitationIdNil(b bool)`

 SetInvitationIdNil sets the value for InvitationId to be an explicit nil

### UnsetInvitationId
`func (o *ActiveOrganizationMemberPublic) UnsetInvitationId()`

UnsetInvitationId ensures that no value is present for InvitationId, not even an explicit nil
### GetOrganizationMemberId

`func (o *ActiveOrganizationMemberPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *ActiveOrganizationMemberPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *ActiveOrganizationMemberPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.

### HasOrganizationMemberId

`func (o *ActiveOrganizationMemberPublic) HasOrganizationMemberId() bool`

HasOrganizationMemberId returns a boolean if a field has been set.

### SetOrganizationMemberIdNil

`func (o *ActiveOrganizationMemberPublic) SetOrganizationMemberIdNil(b bool)`

 SetOrganizationMemberIdNil sets the value for OrganizationMemberId to be an explicit nil

### UnsetOrganizationMemberId
`func (o *ActiveOrganizationMemberPublic) UnsetOrganizationMemberId()`

UnsetOrganizationMemberId ensures that no value is present for OrganizationMemberId, not even an explicit nil
### GetRole

`func (o *ActiveOrganizationMemberPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ActiveOrganizationMemberPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ActiveOrganizationMemberPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ActiveOrganizationMemberPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveOrganizationMemberPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveOrganizationMemberPublic) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ActiveOrganizationMemberPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActiveOrganizationMemberPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActiveOrganizationMemberPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActiveOrganizationMemberPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ActiveOrganizationMemberPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ActiveOrganizationMemberPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *ActiveOrganizationMemberPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActiveOrganizationMemberPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActiveOrganizationMemberPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActiveOrganizationMemberPublic) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ActiveOrganizationMemberPublic) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ActiveOrganizationMemberPublic) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


