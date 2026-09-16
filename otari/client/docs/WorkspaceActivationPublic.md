# WorkspaceActivationPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationAttempt** | Pointer to [**NullableActivationAttemptPublic**](ActivationAttemptPublic.md) | The first request that succeeded, which is the guide&#39;s receipt. Null until one does. | [optional] 
**Dismissed** | **bool** | Whether someone skipped the guide for this workspace. | 
**ExperienceEligible** | **bool** | Whether the dashboard should offer the guide to this caller right now: the deployment has it enabled, the workspace is classified for it, nobody dismissed it, no request has succeeded yet, and the caller may manage the workspace. | 
**LatestAttempt** | Pointer to [**NullableActivationAttemptPublic**](ActivationAttemptPublic.md) | The most recent request, so a failure can be reported while the guide keeps waiting. | [optional] 
**Status** | **string** | &#39;activated&#39; once a gateway request in this workspace has succeeded, &#39;failed&#39; when the last one failed and none has yet succeeded, &#39;waiting&#39; when there has been none at all. | 

## Methods

### NewWorkspaceActivationPublic

`func NewWorkspaceActivationPublic(dismissed bool, experienceEligible bool, status string, ) *WorkspaceActivationPublic`

NewWorkspaceActivationPublic instantiates a new WorkspaceActivationPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceActivationPublicWithDefaults

`func NewWorkspaceActivationPublicWithDefaults() *WorkspaceActivationPublic`

NewWorkspaceActivationPublicWithDefaults instantiates a new WorkspaceActivationPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationAttempt

`func (o *WorkspaceActivationPublic) GetActivationAttempt() ActivationAttemptPublic`

GetActivationAttempt returns the ActivationAttempt field if non-nil, zero value otherwise.

### GetActivationAttemptOk

`func (o *WorkspaceActivationPublic) GetActivationAttemptOk() (*ActivationAttemptPublic, bool)`

GetActivationAttemptOk returns a tuple with the ActivationAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationAttempt

`func (o *WorkspaceActivationPublic) SetActivationAttempt(v ActivationAttemptPublic)`

SetActivationAttempt sets ActivationAttempt field to given value.

### HasActivationAttempt

`func (o *WorkspaceActivationPublic) HasActivationAttempt() bool`

HasActivationAttempt returns a boolean if a field has been set.

### SetActivationAttemptNil

`func (o *WorkspaceActivationPublic) SetActivationAttemptNil(b bool)`

 SetActivationAttemptNil sets the value for ActivationAttempt to be an explicit nil

### UnsetActivationAttempt
`func (o *WorkspaceActivationPublic) UnsetActivationAttempt()`

UnsetActivationAttempt ensures that no value is present for ActivationAttempt, not even an explicit nil
### GetDismissed

`func (o *WorkspaceActivationPublic) GetDismissed() bool`

GetDismissed returns the Dismissed field if non-nil, zero value otherwise.

### GetDismissedOk

`func (o *WorkspaceActivationPublic) GetDismissedOk() (*bool, bool)`

GetDismissedOk returns a tuple with the Dismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissed

`func (o *WorkspaceActivationPublic) SetDismissed(v bool)`

SetDismissed sets Dismissed field to given value.


### GetExperienceEligible

`func (o *WorkspaceActivationPublic) GetExperienceEligible() bool`

GetExperienceEligible returns the ExperienceEligible field if non-nil, zero value otherwise.

### GetExperienceEligibleOk

`func (o *WorkspaceActivationPublic) GetExperienceEligibleOk() (*bool, bool)`

GetExperienceEligibleOk returns a tuple with the ExperienceEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceEligible

`func (o *WorkspaceActivationPublic) SetExperienceEligible(v bool)`

SetExperienceEligible sets ExperienceEligible field to given value.


### GetLatestAttempt

`func (o *WorkspaceActivationPublic) GetLatestAttempt() ActivationAttemptPublic`

GetLatestAttempt returns the LatestAttempt field if non-nil, zero value otherwise.

### GetLatestAttemptOk

`func (o *WorkspaceActivationPublic) GetLatestAttemptOk() (*ActivationAttemptPublic, bool)`

GetLatestAttemptOk returns a tuple with the LatestAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAttempt

`func (o *WorkspaceActivationPublic) SetLatestAttempt(v ActivationAttemptPublic)`

SetLatestAttempt sets LatestAttempt field to given value.

### HasLatestAttempt

`func (o *WorkspaceActivationPublic) HasLatestAttempt() bool`

HasLatestAttempt returns a boolean if a field has been set.

### SetLatestAttemptNil

`func (o *WorkspaceActivationPublic) SetLatestAttemptNil(b bool)`

 SetLatestAttemptNil sets the value for LatestAttempt to be an explicit nil

### UnsetLatestAttempt
`func (o *WorkspaceActivationPublic) UnsetLatestAttempt()`

UnsetLatestAttempt ensures that no value is present for LatestAttempt, not even an explicit nil
### GetStatus

`func (o *WorkspaceActivationPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkspaceActivationPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkspaceActivationPublic) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


