# PolicyCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckResults** | Pointer to [**[]CheckVerdictRequest**](CheckVerdictRequest.md) | Verifier verdicts the caller collected for this request&#39;s verifier gates. | [optional] 
**CommandScope** | Pointer to **string** | What &#x60;commands&#x60; covers: &#x60;call&#x60; for the single tool call about to run, &#x60;session&#x60; for every command the session has run so far. | [optional] [default to "call"]
**Commands** | Pointer to **[]string** | Shell commands the caller observed run or is about to run. | [optional] 
**JudgeResults** | Pointer to [**[]JudgeVerdictRequest**](JudgeVerdictRequest.md) | Model verdicts the caller collected for this request&#39;s judge gates. | [optional] 
**PathSource** | Pointer to **NullableString** | Which moment &#x60;paths&#x60; was read at, matching the &#x60;runs&#x60; values a path gate declares. Only three of the six &#x60;runs&#x60; values are legal here, because only those three are moments a path can be read at: &#x60;pre_tool_use.edit_target&#x60; for a write tool&#39;s own target before it runs, &#x60;pre_tool_use.read_target&#x60; for a read tool&#39;s, and &#x60;stop.working_tree&#x60; for &#x60;git status&#x60; once the turn is over. Required whenever &#x60;paths&#x60; is non-empty, and rejected with a 422 if omitted or set to any other value: either would resolve every path gate &#x60;not_applicable&#x60;, which loses enforcement without reporting anything. An empty &#x60;paths&#x60; needs no source. | [optional] 
**Paths** | Pointer to **[]string** | Repo-relative paths this moment of the session puts in scope: what &#x60;git status --porcelain&#x60; reports, or the single target a tool call is about to write or read. &#x60;path_source&#x60; says which. | [optional] 
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

### GetCheckResults

`func (o *PolicyCheckRequest) GetCheckResults() []CheckVerdictRequest`

GetCheckResults returns the CheckResults field if non-nil, zero value otherwise.

### GetCheckResultsOk

`func (o *PolicyCheckRequest) GetCheckResultsOk() (*[]CheckVerdictRequest, bool)`

GetCheckResultsOk returns a tuple with the CheckResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResults

`func (o *PolicyCheckRequest) SetCheckResults(v []CheckVerdictRequest)`

SetCheckResults sets CheckResults field to given value.

### HasCheckResults

`func (o *PolicyCheckRequest) HasCheckResults() bool`

HasCheckResults returns a boolean if a field has been set.

### SetCheckResultsNil

`func (o *PolicyCheckRequest) SetCheckResultsNil(b bool)`

 SetCheckResultsNil sets the value for CheckResults to be an explicit nil

### UnsetCheckResults
`func (o *PolicyCheckRequest) UnsetCheckResults()`

UnsetCheckResults ensures that no value is present for CheckResults, not even an explicit nil
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
### GetPathSource

`func (o *PolicyCheckRequest) GetPathSource() string`

GetPathSource returns the PathSource field if non-nil, zero value otherwise.

### GetPathSourceOk

`func (o *PolicyCheckRequest) GetPathSourceOk() (*string, bool)`

GetPathSourceOk returns a tuple with the PathSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSource

`func (o *PolicyCheckRequest) SetPathSource(v string)`

SetPathSource sets PathSource field to given value.

### HasPathSource

`func (o *PolicyCheckRequest) HasPathSource() bool`

HasPathSource returns a boolean if a field has been set.

### SetPathSourceNil

`func (o *PolicyCheckRequest) SetPathSourceNil(b bool)`

 SetPathSourceNil sets the value for PathSource to be an explicit nil

### UnsetPathSource
`func (o *PolicyCheckRequest) UnsetPathSource()`

UnsetPathSource ensures that no value is present for PathSource, not even an explicit nil
### GetPaths

`func (o *PolicyCheckRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PolicyCheckRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PolicyCheckRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PolicyCheckRequest) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### SetPathsNil

`func (o *PolicyCheckRequest) SetPathsNil(b bool)`

 SetPathsNil sets the value for Paths to be an explicit nil

### UnsetPaths
`func (o *PolicyCheckRequest) UnsetPaths()`

UnsetPaths ensures that no value is present for Paths, not even an explicit nil
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


