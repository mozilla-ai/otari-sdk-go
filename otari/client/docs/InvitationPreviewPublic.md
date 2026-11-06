# InvitationPreviewPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSignIn** | **bool** | Whether the invited address can already sign in, with a password or through a provider. | 
**Email** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**NeedsPassword** | **bool** | Whether accepting should also set the invited address&#39;s first password. False when the address already has a way in, and also when another organization has it as a member or invitee, since this link may not choose a password for it then; can_sign_in tells the two apart. Accept refuses a password whenever this is false. | 
**OrganizationName** | **string** |  | 
**Role** | **string** |  | 
**WaysIn** | **[]string** | When the address cannot sign in yet and this link will not set its password, what this deployment offers instead once it has accepted: &#39;provider&#39; (sign in with Google or GitHub) and &#39;password_reset&#39; (choose a password with \&quot;Forgot your password?\&quot;). Empty otherwise, and empty when the deployment offers neither, in which case whoever administers it has to help. | 

## Methods

### NewInvitationPreviewPublic

`func NewInvitationPreviewPublic(canSignIn bool, email string, expiresAt time.Time, needsPassword bool, organizationName string, role string, waysIn []string, ) *InvitationPreviewPublic`

NewInvitationPreviewPublic instantiates a new InvitationPreviewPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationPreviewPublicWithDefaults

`func NewInvitationPreviewPublicWithDefaults() *InvitationPreviewPublic`

NewInvitationPreviewPublicWithDefaults instantiates a new InvitationPreviewPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSignIn

`func (o *InvitationPreviewPublic) GetCanSignIn() bool`

GetCanSignIn returns the CanSignIn field if non-nil, zero value otherwise.

### GetCanSignInOk

`func (o *InvitationPreviewPublic) GetCanSignInOk() (*bool, bool)`

GetCanSignInOk returns a tuple with the CanSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignIn

`func (o *InvitationPreviewPublic) SetCanSignIn(v bool)`

SetCanSignIn sets CanSignIn field to given value.


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


### GetNeedsPassword

`func (o *InvitationPreviewPublic) GetNeedsPassword() bool`

GetNeedsPassword returns the NeedsPassword field if non-nil, zero value otherwise.

### GetNeedsPasswordOk

`func (o *InvitationPreviewPublic) GetNeedsPasswordOk() (*bool, bool)`

GetNeedsPasswordOk returns a tuple with the NeedsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsPassword

`func (o *InvitationPreviewPublic) SetNeedsPassword(v bool)`

SetNeedsPassword sets NeedsPassword field to given value.


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


### GetWaysIn

`func (o *InvitationPreviewPublic) GetWaysIn() []string`

GetWaysIn returns the WaysIn field if non-nil, zero value otherwise.

### GetWaysInOk

`func (o *InvitationPreviewPublic) GetWaysInOk() (*[]string, bool)`

GetWaysInOk returns a tuple with the WaysIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaysIn

`func (o *InvitationPreviewPublic) SetWaysIn(v []string)`

SetWaysIn sets WaysIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


