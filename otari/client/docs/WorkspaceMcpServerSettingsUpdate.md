# WorkspaceMcpServerSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachByDefault** | Pointer to **NullableBool** | Whether every request of the workspace gets the server; null takes the server&#39;s | [optional] 
**Enabled** | Pointer to **bool** | The workspace&#39;s own switch: false turns the server off for this workspace | [optional] 
**EveryoneTools** | Pointer to **[]string** | Tool patterns every key of the workspace may call; [] denies everyone; null takes the server&#39;s | [optional] 

## Methods

### NewWorkspaceMcpServerSettingsUpdate

`func NewWorkspaceMcpServerSettingsUpdate() *WorkspaceMcpServerSettingsUpdate`

NewWorkspaceMcpServerSettingsUpdate instantiates a new WorkspaceMcpServerSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerSettingsUpdateWithDefaults

`func NewWorkspaceMcpServerSettingsUpdateWithDefaults() *WorkspaceMcpServerSettingsUpdate`

NewWorkspaceMcpServerSettingsUpdateWithDefaults instantiates a new WorkspaceMcpServerSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachByDefault

`func (o *WorkspaceMcpServerSettingsUpdate) GetAttachByDefault() bool`

GetAttachByDefault returns the AttachByDefault field if non-nil, zero value otherwise.

### GetAttachByDefaultOk

`func (o *WorkspaceMcpServerSettingsUpdate) GetAttachByDefaultOk() (*bool, bool)`

GetAttachByDefaultOk returns a tuple with the AttachByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachByDefault

`func (o *WorkspaceMcpServerSettingsUpdate) SetAttachByDefault(v bool)`

SetAttachByDefault sets AttachByDefault field to given value.

### HasAttachByDefault

`func (o *WorkspaceMcpServerSettingsUpdate) HasAttachByDefault() bool`

HasAttachByDefault returns a boolean if a field has been set.

### SetAttachByDefaultNil

`func (o *WorkspaceMcpServerSettingsUpdate) SetAttachByDefaultNil(b bool)`

 SetAttachByDefaultNil sets the value for AttachByDefault to be an explicit nil

### UnsetAttachByDefault
`func (o *WorkspaceMcpServerSettingsUpdate) UnsetAttachByDefault()`

UnsetAttachByDefault ensures that no value is present for AttachByDefault, not even an explicit nil
### GetEnabled

`func (o *WorkspaceMcpServerSettingsUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerSettingsUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerSettingsUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkspaceMcpServerSettingsUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEveryoneTools

`func (o *WorkspaceMcpServerSettingsUpdate) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *WorkspaceMcpServerSettingsUpdate) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *WorkspaceMcpServerSettingsUpdate) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.

### HasEveryoneTools

`func (o *WorkspaceMcpServerSettingsUpdate) HasEveryoneTools() bool`

HasEveryoneTools returns a boolean if a field has been set.

### SetEveryoneToolsNil

`func (o *WorkspaceMcpServerSettingsUpdate) SetEveryoneToolsNil(b bool)`

 SetEveryoneToolsNil sets the value for EveryoneTools to be an explicit nil

### UnsetEveryoneTools
`func (o *WorkspaceMcpServerSettingsUpdate) UnsetEveryoneTools()`

UnsetEveryoneTools ensures that no value is present for EveryoneTools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


