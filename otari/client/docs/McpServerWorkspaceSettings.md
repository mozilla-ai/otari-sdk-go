# McpServerWorkspaceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachByDefault** | **NullableBool** | Whether the workspace attaches it; null takes the server&#39;s own | 
**Enabled** | **bool** | The workspace&#39;s own switch | 
**EveryoneTools** | **[]string** | What every key of the workspace may call; null takes the server&#39;s own | 
**OrganizationTools** | **[]string** | The organization&#39;s limit for the workspace, which caps every rule set here; null sets none | 
**Reached** | **bool** | The workspace&#39;s requests reach the server: neither the organization nor the workspace turned it off here, and the server itself is on | 

## Methods

### NewMcpServerWorkspaceSettings

`func NewMcpServerWorkspaceSettings(attachByDefault NullableBool, enabled bool, everyoneTools []string, organizationTools []string, reached bool, ) *McpServerWorkspaceSettings`

NewMcpServerWorkspaceSettings instantiates a new McpServerWorkspaceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerWorkspaceSettingsWithDefaults

`func NewMcpServerWorkspaceSettingsWithDefaults() *McpServerWorkspaceSettings`

NewMcpServerWorkspaceSettingsWithDefaults instantiates a new McpServerWorkspaceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachByDefault

`func (o *McpServerWorkspaceSettings) GetAttachByDefault() bool`

GetAttachByDefault returns the AttachByDefault field if non-nil, zero value otherwise.

### GetAttachByDefaultOk

`func (o *McpServerWorkspaceSettings) GetAttachByDefaultOk() (*bool, bool)`

GetAttachByDefaultOk returns a tuple with the AttachByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachByDefault

`func (o *McpServerWorkspaceSettings) SetAttachByDefault(v bool)`

SetAttachByDefault sets AttachByDefault field to given value.


### SetAttachByDefaultNil

`func (o *McpServerWorkspaceSettings) SetAttachByDefaultNil(b bool)`

 SetAttachByDefaultNil sets the value for AttachByDefault to be an explicit nil

### UnsetAttachByDefault
`func (o *McpServerWorkspaceSettings) UnsetAttachByDefault()`

UnsetAttachByDefault ensures that no value is present for AttachByDefault, not even an explicit nil
### GetEnabled

`func (o *McpServerWorkspaceSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpServerWorkspaceSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpServerWorkspaceSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEveryoneTools

`func (o *McpServerWorkspaceSettings) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *McpServerWorkspaceSettings) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *McpServerWorkspaceSettings) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.


### SetEveryoneToolsNil

`func (o *McpServerWorkspaceSettings) SetEveryoneToolsNil(b bool)`

 SetEveryoneToolsNil sets the value for EveryoneTools to be an explicit nil

### UnsetEveryoneTools
`func (o *McpServerWorkspaceSettings) UnsetEveryoneTools()`

UnsetEveryoneTools ensures that no value is present for EveryoneTools, not even an explicit nil
### GetOrganizationTools

`func (o *McpServerWorkspaceSettings) GetOrganizationTools() []string`

GetOrganizationTools returns the OrganizationTools field if non-nil, zero value otherwise.

### GetOrganizationToolsOk

`func (o *McpServerWorkspaceSettings) GetOrganizationToolsOk() (*[]string, bool)`

GetOrganizationToolsOk returns a tuple with the OrganizationTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationTools

`func (o *McpServerWorkspaceSettings) SetOrganizationTools(v []string)`

SetOrganizationTools sets OrganizationTools field to given value.


### SetOrganizationToolsNil

`func (o *McpServerWorkspaceSettings) SetOrganizationToolsNil(b bool)`

 SetOrganizationToolsNil sets the value for OrganizationTools to be an explicit nil

### UnsetOrganizationTools
`func (o *McpServerWorkspaceSettings) UnsetOrganizationTools()`

UnsetOrganizationTools ensures that no value is present for OrganizationTools, not even an explicit nil
### GetReached

`func (o *McpServerWorkspaceSettings) GetReached() bool`

GetReached returns the Reached field if non-nil, zero value otherwise.

### GetReachedOk

`func (o *McpServerWorkspaceSettings) GetReachedOk() (*bool, bool)`

GetReachedOk returns a tuple with the Reached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReached

`func (o *McpServerWorkspaceSettings) SetReached(v bool)`

SetReached sets Reached field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


