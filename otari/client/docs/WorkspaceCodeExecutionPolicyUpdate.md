# WorkspaceCodeExecutionPolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPurposeHint** | Pointer to **NullableString** | Hint used when a request declares otari_code_execution without one of its own | [optional] 
**Enabled** | **bool** | False refuses code execution for this workspace | 
**ExecTimeoutS** | Pointer to **NullableInt32** | Ceiling on one execution&#39;s runtime in seconds; only ever lowers the effective limit, so at most 60 | [optional] 
**Image** | Pointer to **NullableString** | Sandbox image this workspace&#39;s code runs in. Must be one the operator curated into sandbox_allowed_session_images (or the deployment&#39;s own sandbox_session_image); null uses the deployment&#39;s | [optional] 
**MaxIterations** | Pointer to **NullableInt32** | Ceiling on tool-loop iterations; only ever lowers the effective limit, so at most 25 | [optional] 
**Tools** | Pointer to **[]string** | Code-execution tool kinds this workspace may use, from code_execution, bash_code_execution, text_editor_code_execution. Only ever removes one the backend serves; null exposes whatever it serves | [optional] 

## Methods

### NewWorkspaceCodeExecutionPolicyUpdate

`func NewWorkspaceCodeExecutionPolicyUpdate(enabled bool, ) *WorkspaceCodeExecutionPolicyUpdate`

NewWorkspaceCodeExecutionPolicyUpdate instantiates a new WorkspaceCodeExecutionPolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCodeExecutionPolicyUpdateWithDefaults

`func NewWorkspaceCodeExecutionPolicyUpdateWithDefaults() *WorkspaceCodeExecutionPolicyUpdate`

NewWorkspaceCodeExecutionPolicyUpdateWithDefaults instantiates a new WorkspaceCodeExecutionPolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPurposeHint

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetDefaultPurposeHint() string`

GetDefaultPurposeHint returns the DefaultPurposeHint field if non-nil, zero value otherwise.

### GetDefaultPurposeHintOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetDefaultPurposeHintOk() (*string, bool)`

GetDefaultPurposeHintOk returns a tuple with the DefaultPurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPurposeHint

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetDefaultPurposeHint(v string)`

SetDefaultPurposeHint sets DefaultPurposeHint field to given value.

### HasDefaultPurposeHint

`func (o *WorkspaceCodeExecutionPolicyUpdate) HasDefaultPurposeHint() bool`

HasDefaultPurposeHint returns a boolean if a field has been set.

### SetDefaultPurposeHintNil

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetDefaultPurposeHintNil(b bool)`

 SetDefaultPurposeHintNil sets the value for DefaultPurposeHint to be an explicit nil

### UnsetDefaultPurposeHint
`func (o *WorkspaceCodeExecutionPolicyUpdate) UnsetDefaultPurposeHint()`

UnsetDefaultPurposeHint ensures that no value is present for DefaultPurposeHint, not even an explicit nil
### GetEnabled

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecTimeoutS

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetExecTimeoutS() int32`

GetExecTimeoutS returns the ExecTimeoutS field if non-nil, zero value otherwise.

### GetExecTimeoutSOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetExecTimeoutSOk() (*int32, bool)`

GetExecTimeoutSOk returns a tuple with the ExecTimeoutS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimeoutS

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetExecTimeoutS(v int32)`

SetExecTimeoutS sets ExecTimeoutS field to given value.

### HasExecTimeoutS

`func (o *WorkspaceCodeExecutionPolicyUpdate) HasExecTimeoutS() bool`

HasExecTimeoutS returns a boolean if a field has been set.

### SetExecTimeoutSNil

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetExecTimeoutSNil(b bool)`

 SetExecTimeoutSNil sets the value for ExecTimeoutS to be an explicit nil

### UnsetExecTimeoutS
`func (o *WorkspaceCodeExecutionPolicyUpdate) UnsetExecTimeoutS()`

UnsetExecTimeoutS ensures that no value is present for ExecTimeoutS, not even an explicit nil
### GetImage

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *WorkspaceCodeExecutionPolicyUpdate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *WorkspaceCodeExecutionPolicyUpdate) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetMaxIterations

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetMaxIterations() int32`

GetMaxIterations returns the MaxIterations field if non-nil, zero value otherwise.

### GetMaxIterationsOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetMaxIterationsOk() (*int32, bool)`

GetMaxIterationsOk returns a tuple with the MaxIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIterations

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetMaxIterations(v int32)`

SetMaxIterations sets MaxIterations field to given value.

### HasMaxIterations

`func (o *WorkspaceCodeExecutionPolicyUpdate) HasMaxIterations() bool`

HasMaxIterations returns a boolean if a field has been set.

### SetMaxIterationsNil

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetMaxIterationsNil(b bool)`

 SetMaxIterationsNil sets the value for MaxIterations to be an explicit nil

### UnsetMaxIterations
`func (o *WorkspaceCodeExecutionPolicyUpdate) UnsetMaxIterations()`

UnsetMaxIterations ensures that no value is present for MaxIterations, not even an explicit nil
### GetTools

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceCodeExecutionPolicyUpdate) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *WorkspaceCodeExecutionPolicyUpdate) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *WorkspaceCodeExecutionPolicyUpdate) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *WorkspaceCodeExecutionPolicyUpdate) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


