# AcceptInvitationResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextStep** | [**InvitationNextStep**](InvitationNextStep.md) | What the invited address does now that it has accepted. | 
**OrganizationName** | **string** |  | 
**Role** | **string** |  | 
**WaysIn** | **[]string** | What this deployment offers instead, where next_step names them, as the preview&#39;s ways_in says. | 

## Methods

### NewAcceptInvitationResultPublic

`func NewAcceptInvitationResultPublic(nextStep InvitationNextStep, organizationName string, role string, waysIn []string, ) *AcceptInvitationResultPublic`

NewAcceptInvitationResultPublic instantiates a new AcceptInvitationResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptInvitationResultPublicWithDefaults

`func NewAcceptInvitationResultPublicWithDefaults() *AcceptInvitationResultPublic`

NewAcceptInvitationResultPublicWithDefaults instantiates a new AcceptInvitationResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextStep

`func (o *AcceptInvitationResultPublic) GetNextStep() InvitationNextStep`

GetNextStep returns the NextStep field if non-nil, zero value otherwise.

### GetNextStepOk

`func (o *AcceptInvitationResultPublic) GetNextStepOk() (*InvitationNextStep, bool)`

GetNextStepOk returns a tuple with the NextStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStep

`func (o *AcceptInvitationResultPublic) SetNextStep(v InvitationNextStep)`

SetNextStep sets NextStep field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


