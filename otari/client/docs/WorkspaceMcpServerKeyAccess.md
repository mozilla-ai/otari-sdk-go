# WorkspaceMcpServerKeyAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | **[]string** | The server&#39;s own ceiling; null exposes every tool | 
**DecidedBy** | **string** |  | 
**Enabled** | **bool** |  | 
**GrantId** | **NullableString** |  | 
**Name** | **string** |  | 
**ReadTools** | Pointer to **[]string** | The server&#39;s tools labeled read, which &#39;@read&#39; in tools stands for | [optional] 
**ServerId** | **string** |  | 
**Tools** | **[]string** | The deciding rule&#39;s tool patterns; [] means the key cannot reach the server | 

## Methods

### NewWorkspaceMcpServerKeyAccess

`func NewWorkspaceMcpServerKeyAccess(allowedTools []string, decidedBy string, enabled bool, grantId NullableString, name string, serverId string, tools []string, ) *WorkspaceMcpServerKeyAccess`

NewWorkspaceMcpServerKeyAccess instantiates a new WorkspaceMcpServerKeyAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerKeyAccessWithDefaults

`func NewWorkspaceMcpServerKeyAccessWithDefaults() *WorkspaceMcpServerKeyAccess`

NewWorkspaceMcpServerKeyAccessWithDefaults instantiates a new WorkspaceMcpServerKeyAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerKeyAccess) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerKeyAccess) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerKeyAccess) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.


### SetAllowedToolsNil

`func (o *WorkspaceMcpServerKeyAccess) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerKeyAccess) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetDecidedBy

`func (o *WorkspaceMcpServerKeyAccess) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *WorkspaceMcpServerKeyAccess) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *WorkspaceMcpServerKeyAccess) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.


### GetEnabled

`func (o *WorkspaceMcpServerKeyAccess) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerKeyAccess) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerKeyAccess) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGrantId

`func (o *WorkspaceMcpServerKeyAccess) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *WorkspaceMcpServerKeyAccess) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *WorkspaceMcpServerKeyAccess) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### SetGrantIdNil

`func (o *WorkspaceMcpServerKeyAccess) SetGrantIdNil(b bool)`

 SetGrantIdNil sets the value for GrantId to be an explicit nil

### UnsetGrantId
`func (o *WorkspaceMcpServerKeyAccess) UnsetGrantId()`

UnsetGrantId ensures that no value is present for GrantId, not even an explicit nil
### GetName

`func (o *WorkspaceMcpServerKeyAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceMcpServerKeyAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceMcpServerKeyAccess) SetName(v string)`

SetName sets Name field to given value.


### GetReadTools

`func (o *WorkspaceMcpServerKeyAccess) GetReadTools() []string`

GetReadTools returns the ReadTools field if non-nil, zero value otherwise.

### GetReadToolsOk

`func (o *WorkspaceMcpServerKeyAccess) GetReadToolsOk() (*[]string, bool)`

GetReadToolsOk returns a tuple with the ReadTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTools

`func (o *WorkspaceMcpServerKeyAccess) SetReadTools(v []string)`

SetReadTools sets ReadTools field to given value.

### HasReadTools

`func (o *WorkspaceMcpServerKeyAccess) HasReadTools() bool`

HasReadTools returns a boolean if a field has been set.

### GetServerId

`func (o *WorkspaceMcpServerKeyAccess) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *WorkspaceMcpServerKeyAccess) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *WorkspaceMcpServerKeyAccess) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetTools

`func (o *WorkspaceMcpServerKeyAccess) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerKeyAccess) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerKeyAccess) SetTools(v []string)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


