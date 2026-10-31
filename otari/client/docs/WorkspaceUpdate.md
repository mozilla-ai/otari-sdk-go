# WorkspaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TraceContentEnabled** | Pointer to **NullableBool** | Keep what this workspace&#39;s agent sessions were about: the prompts, commands and file paths its agents report. Only where the deployment keeps content at all. | [optional] 

## Methods

### NewWorkspaceUpdate

`func NewWorkspaceUpdate() *WorkspaceUpdate`

NewWorkspaceUpdate instantiates a new WorkspaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceUpdateWithDefaults

`func NewWorkspaceUpdateWithDefaults() *WorkspaceUpdate`

NewWorkspaceUpdateWithDefaults instantiates a new WorkspaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkspaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspaceUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *WorkspaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkspaceUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkspaceUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTraceContentEnabled

`func (o *WorkspaceUpdate) GetTraceContentEnabled() bool`

GetTraceContentEnabled returns the TraceContentEnabled field if non-nil, zero value otherwise.

### GetTraceContentEnabledOk

`func (o *WorkspaceUpdate) GetTraceContentEnabledOk() (*bool, bool)`

GetTraceContentEnabledOk returns a tuple with the TraceContentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceContentEnabled

`func (o *WorkspaceUpdate) SetTraceContentEnabled(v bool)`

SetTraceContentEnabled sets TraceContentEnabled field to given value.

### HasTraceContentEnabled

`func (o *WorkspaceUpdate) HasTraceContentEnabled() bool`

HasTraceContentEnabled returns a boolean if a field has been set.

### SetTraceContentEnabledNil

`func (o *WorkspaceUpdate) SetTraceContentEnabledNil(b bool)`

 SetTraceContentEnabledNil sets the value for TraceContentEnabled to be an explicit nil

### UnsetTraceContentEnabled
`func (o *WorkspaceUpdate) UnsetTraceContentEnabled()`

UnsetTraceContentEnabled ensures that no value is present for TraceContentEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


