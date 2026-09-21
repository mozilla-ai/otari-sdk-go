# PolicyCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedPaths** | Pointer to **[]string** | Repo-relative paths the caller observed changed (e.g. &#x60;git status --porcelain&#x60;). | [optional] 
**CommandScope** | Pointer to **string** | What &#x60;commands&#x60; covers: &#x60;call&#x60; for the single tool call about to run, &#x60;session&#x60; for every command the session has run so far. | [optional] [default to "call"]
**Commands** | Pointer to **[]string** | Shell commands the caller observed run or is about to run. | [optional] 
**JudgeResults** | Pointer to [**[]JudgeVerdictRequest**](JudgeVerdictRequest.md) | Model verdicts the caller collected for this request&#39;s judge gates. | [optional] 
**PolicyYaml** | **string** |  | 

## Methods

### NewPolicyCheckRequest

`func NewPolicyCheckRequest(policyYaml string, ) *PolicyCheckRequest`

NewPolicyCheckRequest instantiates a new PolicyCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCheckRequestWithDefaults

`func NewPolicyCheckRequestWithDefaults() *PolicyCheckRequest`

NewPolicyCheckRequestWithDefaults instantiates a new PolicyCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedPaths

`func (o *PolicyCheckRequest) GetChangedPaths() []string`

GetChangedPaths returns the ChangedPaths field if non-nil, zero value otherwise.

### GetChangedPathsOk

`func (o *PolicyCheckRequest) GetChangedPathsOk() (*[]string, bool)`

GetChangedPathsOk returns a tuple with the ChangedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedPaths

`func (o *PolicyCheckRequest) SetChangedPaths(v []string)`

SetChangedPaths sets ChangedPaths field to given value.

### HasChangedPaths

`func (o *PolicyCheckRequest) HasChangedPaths() bool`

HasChangedPaths returns a boolean if a field has been set.

### SetChangedPathsNil

`func (o *PolicyCheckRequest) SetChangedPathsNil(b bool)`

 SetChangedPathsNil sets the value for ChangedPaths to be an explicit nil

### UnsetChangedPaths
`func (o *PolicyCheckRequest) UnsetChangedPaths()`

UnsetChangedPaths ensures that no value is present for ChangedPaths, not even an explicit nil
### GetCommandScope

`func (o *PolicyCheckRequest) GetCommandScope() string`

GetCommandScope returns the CommandScope field if non-nil, zero value otherwise.

### GetCommandScopeOk

`func (o *PolicyCheckRequest) GetCommandScopeOk() (*string, bool)`

GetCommandScopeOk returns a tuple with the CommandScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandScope

`func (o *PolicyCheckRequest) SetCommandScope(v string)`

SetCommandScope sets CommandScope field to given value.

### HasCommandScope

`func (o *PolicyCheckRequest) HasCommandScope() bool`

HasCommandScope returns a boolean if a field has been set.

### GetCommands

`func (o *PolicyCheckRequest) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *PolicyCheckRequest) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *PolicyCheckRequest) SetCommands(v []string)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *PolicyCheckRequest) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *PolicyCheckRequest) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *PolicyCheckRequest) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetJudgeResults

`func (o *PolicyCheckRequest) GetJudgeResults() []JudgeVerdictRequest`

GetJudgeResults returns the JudgeResults field if non-nil, zero value otherwise.

### GetJudgeResultsOk

`func (o *PolicyCheckRequest) GetJudgeResultsOk() (*[]JudgeVerdictRequest, bool)`

GetJudgeResultsOk returns a tuple with the JudgeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeResults

`func (o *PolicyCheckRequest) SetJudgeResults(v []JudgeVerdictRequest)`

SetJudgeResults sets JudgeResults field to given value.

### HasJudgeResults

`func (o *PolicyCheckRequest) HasJudgeResults() bool`

HasJudgeResults returns a boolean if a field has been set.

### SetJudgeResultsNil

`func (o *PolicyCheckRequest) SetJudgeResultsNil(b bool)`

 SetJudgeResultsNil sets the value for JudgeResults to be an explicit nil

### UnsetJudgeResults
`func (o *PolicyCheckRequest) UnsetJudgeResults()`

UnsetJudgeResults ensures that no value is present for JudgeResults, not even an explicit nil
### GetPolicyYaml

`func (o *PolicyCheckRequest) GetPolicyYaml() string`

GetPolicyYaml returns the PolicyYaml field if non-nil, zero value otherwise.

### GetPolicyYamlOk

`func (o *PolicyCheckRequest) GetPolicyYamlOk() (*string, bool)`

GetPolicyYamlOk returns a tuple with the PolicyYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyYaml

`func (o *PolicyCheckRequest) SetPolicyYaml(v string)`

SetPolicyYaml sets PolicyYaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


