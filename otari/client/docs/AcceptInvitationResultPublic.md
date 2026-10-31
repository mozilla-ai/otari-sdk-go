# AcceptInvitationResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSignIn** | Pointer to **bool** | Whether the invited address can sign in now, with a password (this accept&#39;s included) or through a provider. False when it accepted without one and has never signed in. | [optional] [default to true]
**OrganizationName** | **string** |  | 
**PasswordSet** | Pointer to **bool** | Whether this accept set the identity&#39;s password. | [optional] [default to false]
**Role** | **string** |  | 
**VerificationPending** | Pointer to **bool** | Whether the invited address chose a password through signup and has not verified it yet, so it has to open its verification link before it can sign in. | [optional] [default to false]
**WaysIn** | Pointer to **[]string** | When can_sign_in is false, what this deployment offers instead, as the preview&#39;s ways_in says. | [optional] 

## Methods

### NewAcceptInvitationResultPublic

`func NewAcceptInvitationResultPublic(organizationName string, role string, ) *AcceptInvitationResultPublic`

NewAcceptInvitationResultPublic instantiates a new AcceptInvitationResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptInvitationResultPublicWithDefaults

`func NewAcceptInvitationResultPublicWithDefaults() *AcceptInvitationResultPublic`

NewAcceptInvitationResultPublicWithDefaults instantiates a new AcceptInvitationResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSignIn

`func (o *AcceptInvitationResultPublic) GetCanSignIn() bool`

GetCanSignIn returns the CanSignIn field if non-nil, zero value otherwise.

### GetCanSignInOk

`func (o *AcceptInvitationResultPublic) GetCanSignInOk() (*bool, bool)`

GetCanSignInOk returns a tuple with the CanSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSignIn

`func (o *AcceptInvitationResultPublic) SetCanSignIn(v bool)`

SetCanSignIn sets CanSignIn field to given value.

### HasCanSignIn

`func (o *AcceptInvitationResultPublic) HasCanSignIn() bool`

HasCanSignIn returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AcceptInvitationResultPublic) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AcceptInvitationResultPublic) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AcceptInvitationResultPublic) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPasswordSet

`func (o *AcceptInvitationResultPublic) GetPasswordSet() bool`

GetPasswordSet returns the PasswordSet field if non-nil, zero value otherwise.

### GetPasswordSetOk

`func (o *AcceptInvitationResultPublic) GetPasswordSetOk() (*bool, bool)`

GetPasswordSetOk returns a tuple with the PasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSet

`func (o *AcceptInvitationResultPublic) SetPasswordSet(v bool)`

SetPasswordSet sets PasswordSet field to given value.

### HasPasswordSet

`func (o *AcceptInvitationResultPublic) HasPasswordSet() bool`

HasPasswordSet returns a boolean if a field has been set.

### GetRole

`func (o *AcceptInvitationResultPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AcceptInvitationResultPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AcceptInvitationResultPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetVerificationPending

`func (o *AcceptInvitationResultPublic) GetVerificationPending() bool`

GetVerificationPending returns the VerificationPending field if non-nil, zero value otherwise.

### GetVerificationPendingOk

`func (o *AcceptInvitationResultPublic) GetVerificationPendingOk() (*bool, bool)`

GetVerificationPendingOk returns a tuple with the VerificationPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationPending

`func (o *AcceptInvitationResultPublic) SetVerificationPending(v bool)`

SetVerificationPending sets VerificationPending field to given value.

### HasVerificationPending

`func (o *AcceptInvitationResultPublic) HasVerificationPending() bool`

HasVerificationPending returns a boolean if a field has been set.

### GetWaysIn

`func (o *AcceptInvitationResultPublic) GetWaysIn() []string`

GetWaysIn returns the WaysIn field if non-nil, zero value otherwise.

### GetWaysInOk

`func (o *AcceptInvitationResultPublic) GetWaysInOk() (*[]string, bool)`

GetWaysInOk returns a tuple with the WaysIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaysIn

`func (o *AcceptInvitationResultPublic) SetWaysIn(v []string)`

SetWaysIn sets WaysIn field to given value.

### HasWaysIn

`func (o *AcceptInvitationResultPublic) HasWaysIn() bool`

HasWaysIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


