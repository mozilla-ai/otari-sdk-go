# InvitationPreviewPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**NextStep** | [**InvitationNextStep**](InvitationNextStep.md) | What the invited address does to get in. Accept refuses a password unless this is &#39;set_password&#39;. | 
**OrganizationName** | **string** |  | 
**Role** | **string** |  | 
**WaysIn** | **[]string** | What this deployment offers instead, where next_step names them: &#39;provider&#39; (sign in with Google or GitHub) and &#39;password_reset&#39; (choose a password with \&quot;Forgot your password?\&quot;). Empty for the other steps, and empty where the deployment offers neither, in which case whoever administers it has to help. | 

## Methods

### NewInvitationPreviewPublic

`func NewInvitationPreviewPublic(email string, expiresAt time.Time, nextStep InvitationNextStep, organizationName string, role string, waysIn []string, ) *InvitationPreviewPublic`

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


### GetNextStep

`func (o *InvitationPreviewPublic) GetNextStep() InvitationNextStep`

GetNextStep returns the NextStep field if non-nil, zero value otherwise.

### GetNextStepOk

`func (o *InvitationPreviewPublic) GetNextStepOk() (*InvitationNextStep, bool)`

GetNextStepOk returns a tuple with the NextStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStep

`func (o *InvitationPreviewPublic) SetNextStep(v InvitationNextStep)`

SetNextStep sets NextStep field to given value.


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


