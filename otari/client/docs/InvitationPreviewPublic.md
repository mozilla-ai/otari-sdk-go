# InvitationPreviewPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**OrganizationName** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewInvitationPreviewPublic

`func NewInvitationPreviewPublic(email string, expiresAt time.Time, organizationName string, role string, ) *InvitationPreviewPublic`

NewInvitationPreviewPublic instantiates a new InvitationPreviewPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationPreviewPublicWithDefaults

`func NewInvitationPreviewPublicWithDefaults() *InvitationPreviewPublic`

NewInvitationPreviewPublicWithDefaults instantiates a new InvitationPreviewPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InvitationPreviewPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationPreviewPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationPreviewPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpiresAt

`func (o *InvitationPreviewPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvitationPreviewPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvitationPreviewPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetOrganizationName

`func (o *InvitationPreviewPublic) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InvitationPreviewPublic) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InvitationPreviewPublic) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetRole

`func (o *InvitationPreviewPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InvitationPreviewPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InvitationPreviewPublic) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


