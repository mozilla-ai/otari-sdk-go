# ActiveOrganizationMemberCreateResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributionUserId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**Email** | **string** |  | 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**InvitationId** | Pointer to **NullableString** |  | [optional] 
**OrganizationMemberId** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewActiveOrganizationMemberCreateResultPublic

`func NewActiveOrganizationMemberCreateResultPublic(email string, role string, status string, ) *ActiveOrganizationMemberCreateResultPublic`

NewActiveOrganizationMemberCreateResultPublic instantiates a new ActiveOrganizationMemberCreateResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveOrganizationMemberCreateResultPublicWithDefaults

`func NewActiveOrganizationMemberCreateResultPublicWithDefaults() *ActiveOrganizationMemberCreateResultPublic`

NewActiveOrganizationMemberCreateResultPublicWithDefaults instantiates a new ActiveOrganizationMemberCreateResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributionUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) GetAttributionUserId() string`

GetAttributionUserId returns the AttributionUserId field if non-nil, zero value otherwise.

### GetAttributionUserIdOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetAttributionUserIdOk() (*string, bool)`

GetAttributionUserIdOk returns a tuple with the AttributionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) SetAttributionUserId(v string)`

SetAttributionUserId sets AttributionUserId field to given value.

### HasAttributionUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) HasAttributionUserId() bool`

HasAttributionUserId returns a boolean if a field has been set.

### SetAttributionUserIdNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetAttributionUserIdNil(b bool)`

 SetAttributionUserIdNil sets the value for AttributionUserId to be an explicit nil

### UnsetAttributionUserId
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetAttributionUserId()`

UnsetAttributionUserId ensures that no value is present for AttributionUserId, not even an explicit nil
### GetCreatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEmail

`func (o *ActiveOrganizationMemberCreateResultPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ActiveOrganizationMemberCreateResultPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpiresAt

`func (o *ActiveOrganizationMemberCreateResultPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ActiveOrganizationMemberCreateResultPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ActiveOrganizationMemberCreateResultPublic) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFullName

`func (o *ActiveOrganizationMemberCreateResultPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ActiveOrganizationMemberCreateResultPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ActiveOrganizationMemberCreateResultPublic) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetInvitationId

`func (o *ActiveOrganizationMemberCreateResultPublic) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *ActiveOrganizationMemberCreateResultPublic) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *ActiveOrganizationMemberCreateResultPublic) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### SetInvitationIdNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetInvitationIdNil(b bool)`

 SetInvitationIdNil sets the value for InvitationId to be an explicit nil

### UnsetInvitationId
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetInvitationId()`

UnsetInvitationId ensures that no value is present for InvitationId, not even an explicit nil
### GetOrganizationMemberId

`func (o *ActiveOrganizationMemberCreateResultPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *ActiveOrganizationMemberCreateResultPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.

### HasOrganizationMemberId

`func (o *ActiveOrganizationMemberCreateResultPublic) HasOrganizationMemberId() bool`

HasOrganizationMemberId returns a boolean if a field has been set.

### SetOrganizationMemberIdNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetOrganizationMemberIdNil(b bool)`

 SetOrganizationMemberIdNil sets the value for OrganizationMemberId to be an explicit nil

### UnsetOrganizationMemberId
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetOrganizationMemberId()`

UnsetOrganizationMemberId ensures that no value is present for OrganizationMemberId, not even an explicit nil
### GetRole

`func (o *ActiveOrganizationMemberCreateResultPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ActiveOrganizationMemberCreateResultPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ActiveOrganizationMemberCreateResultPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveOrganizationMemberCreateResultPublic) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ActiveOrganizationMemberCreateResultPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActiveOrganizationMemberCreateResultPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActiveOrganizationMemberCreateResultPublic) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ActiveOrganizationMemberCreateResultPublic) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ActiveOrganizationMemberCreateResultPublic) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


