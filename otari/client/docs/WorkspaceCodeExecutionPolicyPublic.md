# WorkspaceCodeExecutionPolicyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedImages** | **[]string** |  | 
**AvailableTools** | **[]string** |  | 
**Configured** | **bool** |  | 
**CreatedAt** | **NullableString** |  | 
**DefaultPurposeHint** | **NullableString** |  | 
**Enabled** | **bool** |  | 
**ExecTimeoutS** | **NullableInt32** |  | 
**Image** | **NullableString** |  | 
**MaxIterations** | **NullableInt32** |  | 
**SandboxConfigured** | **bool** |  | 
**Tools** | **[]string** |  | 
**UpdatedAt** | **NullableString** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceCodeExecutionPolicyPublic

`func NewWorkspaceCodeExecutionPolicyPublic(allowedImages []string, availableTools []string, configured bool, createdAt NullableString, defaultPurposeHint NullableString, enabled bool, execTimeoutS NullableInt32, image NullableString, maxIterations NullableInt32, sandboxConfigured bool, tools []string, updatedAt NullableString, workspaceId string, ) *WorkspaceCodeExecutionPolicyPublic`

NewWorkspaceCodeExecutionPolicyPublic instantiates a new WorkspaceCodeExecutionPolicyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCodeExecutionPolicyPublicWithDefaults

`func NewWorkspaceCodeExecutionPolicyPublicWithDefaults() *WorkspaceCodeExecutionPolicyPublic`

NewWorkspaceCodeExecutionPolicyPublicWithDefaults instantiates a new WorkspaceCodeExecutionPolicyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedImages

`func (o *WorkspaceCodeExecutionPolicyPublic) GetAllowedImages() []string`

GetAllowedImages returns the AllowedImages field if non-nil, zero value otherwise.

### GetAllowedImagesOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetAllowedImagesOk() (*[]string, bool)`

GetAllowedImagesOk returns a tuple with the AllowedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedImages

`func (o *WorkspaceCodeExecutionPolicyPublic) SetAllowedImages(v []string)`

SetAllowedImages sets AllowedImages field to given value.


### GetAvailableTools

`func (o *WorkspaceCodeExecutionPolicyPublic) GetAvailableTools() []string`

GetAvailableTools returns the AvailableTools field if non-nil, zero value otherwise.

### GetAvailableToolsOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetAvailableToolsOk() (*[]string, bool)`

GetAvailableToolsOk returns a tuple with the AvailableTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTools

`func (o *WorkspaceCodeExecutionPolicyPublic) SetAvailableTools(v []string)`

SetAvailableTools sets AvailableTools field to given value.


### GetConfigured

`func (o *WorkspaceCodeExecutionPolicyPublic) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *WorkspaceCodeExecutionPolicyPublic) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.


### GetCreatedAt

`func (o *WorkspaceCodeExecutionPolicyPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceCodeExecutionPolicyPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDefaultPurposeHint

`func (o *WorkspaceCodeExecutionPolicyPublic) GetDefaultPurposeHint() string`

GetDefaultPurposeHint returns the DefaultPurposeHint field if non-nil, zero value otherwise.

### GetDefaultPurposeHintOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetDefaultPurposeHintOk() (*string, bool)`

GetDefaultPurposeHintOk returns a tuple with the DefaultPurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPurposeHint

`func (o *WorkspaceCodeExecutionPolicyPublic) SetDefaultPurposeHint(v string)`

SetDefaultPurposeHint sets DefaultPurposeHint field to given value.


### SetDefaultPurposeHintNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetDefaultPurposeHintNil(b bool)`

 SetDefaultPurposeHintNil sets the value for DefaultPurposeHint to be an explicit nil

### UnsetDefaultPurposeHint
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetDefaultPurposeHint()`

UnsetDefaultPurposeHint ensures that no value is present for DefaultPurposeHint, not even an explicit nil
### GetEnabled

`func (o *WorkspaceCodeExecutionPolicyPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceCodeExecutionPolicyPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExecTimeoutS

`func (o *WorkspaceCodeExecutionPolicyPublic) GetExecTimeoutS() int32`

GetExecTimeoutS returns the ExecTimeoutS field if non-nil, zero value otherwise.

### GetExecTimeoutSOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetExecTimeoutSOk() (*int32, bool)`

GetExecTimeoutSOk returns a tuple with the ExecTimeoutS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimeoutS

`func (o *WorkspaceCodeExecutionPolicyPublic) SetExecTimeoutS(v int32)`

SetExecTimeoutS sets ExecTimeoutS field to given value.


### SetExecTimeoutSNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetExecTimeoutSNil(b bool)`

 SetExecTimeoutSNil sets the value for ExecTimeoutS to be an explicit nil

### UnsetExecTimeoutS
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetExecTimeoutS()`

UnsetExecTimeoutS ensures that no value is present for ExecTimeoutS, not even an explicit nil
### GetImage

`func (o *WorkspaceCodeExecutionPolicyPublic) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WorkspaceCodeExecutionPolicyPublic) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetMaxIterations

`func (o *WorkspaceCodeExecutionPolicyPublic) GetMaxIterations() int32`

GetMaxIterations returns the MaxIterations field if non-nil, zero value otherwise.

### GetMaxIterationsOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetMaxIterationsOk() (*int32, bool)`

GetMaxIterationsOk returns a tuple with the MaxIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIterations

`func (o *WorkspaceCodeExecutionPolicyPublic) SetMaxIterations(v int32)`

SetMaxIterations sets MaxIterations field to given value.


### SetMaxIterationsNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetMaxIterationsNil(b bool)`

 SetMaxIterationsNil sets the value for MaxIterations to be an explicit nil

### UnsetMaxIterations
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetMaxIterations()`

UnsetMaxIterations ensures that no value is present for MaxIterations, not even an explicit nil
### GetSandboxConfigured

`func (o *WorkspaceCodeExecutionPolicyPublic) GetSandboxConfigured() bool`

GetSandboxConfigured returns the SandboxConfigured field if non-nil, zero value otherwise.

### GetSandboxConfiguredOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetSandboxConfiguredOk() (*bool, bool)`

GetSandboxConfiguredOk returns a tuple with the SandboxConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxConfigured

`func (o *WorkspaceCodeExecutionPolicyPublic) SetSandboxConfigured(v bool)`

SetSandboxConfigured sets SandboxConfigured field to given value.


### GetTools

`func (o *WorkspaceCodeExecutionPolicyPublic) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceCodeExecutionPolicyPublic) SetTools(v []string)`

SetTools sets Tools field to given value.


### SetToolsNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetUpdatedAt

`func (o *WorkspaceCodeExecutionPolicyPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceCodeExecutionPolicyPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *WorkspaceCodeExecutionPolicyPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WorkspaceCodeExecutionPolicyPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWorkspaceId

`func (o *WorkspaceCodeExecutionPolicyPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceCodeExecutionPolicyPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceCodeExecutionPolicyPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


