# MRBetaMCPToolset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**McpServerName** | **string** |  | 
**Type** | **string** |  | 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**Configs** | Pointer to [**map[string]MRBetaMCPToolConfig**](MRBetaMCPToolConfig.md) |  | [optional] 
**DefaultConfig** | Pointer to [**NullableMRBetaMCPToolDefaultConfig**](MRBetaMCPToolDefaultConfig.md) |  | [optional] 
**Tools** | Pointer to [**[]MRBetaMCPToolParam**](MRBetaMCPToolParam.md) |  | [optional] 

## Methods

### NewMRBetaMCPToolset

`func NewMRBetaMCPToolset(mcpServerName string, type_ string, ) *MRBetaMCPToolset`

NewMRBetaMCPToolset instantiates a new MRBetaMCPToolset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaMCPToolsetWithDefaults

`func NewMRBetaMCPToolsetWithDefaults() *MRBetaMCPToolset`

NewMRBetaMCPToolsetWithDefaults instantiates a new MRBetaMCPToolset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcpServerName

`func (o *MRBetaMCPToolset) GetMcpServerName() string`

GetMcpServerName returns the McpServerName field if non-nil, zero value otherwise.

### GetMcpServerNameOk

`func (o *MRBetaMCPToolset) GetMcpServerNameOk() (*string, bool)`

GetMcpServerNameOk returns a tuple with the McpServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerName

`func (o *MRBetaMCPToolset) SetMcpServerName(v string)`

SetMcpServerName sets McpServerName field to given value.


### GetType

`func (o *MRBetaMCPToolset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaMCPToolset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaMCPToolset) SetType(v string)`

SetType sets Type field to given value.


### GetCacheControl

`func (o *MRBetaMCPToolset) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaMCPToolset) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaMCPToolset) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaMCPToolset) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaMCPToolset) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaMCPToolset) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetConfigs

`func (o *MRBetaMCPToolset) GetConfigs() map[string]MRBetaMCPToolConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *MRBetaMCPToolset) GetConfigsOk() (*map[string]MRBetaMCPToolConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *MRBetaMCPToolset) SetConfigs(v map[string]MRBetaMCPToolConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *MRBetaMCPToolset) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### SetConfigsNil

`func (o *MRBetaMCPToolset) SetConfigsNil(b bool)`

 SetConfigsNil sets the value for Configs to be an explicit nil

### UnsetConfigs
`func (o *MRBetaMCPToolset) UnsetConfigs()`

UnsetConfigs ensures that no value is present for Configs, not even an explicit nil
### GetDefaultConfig

`func (o *MRBetaMCPToolset) GetDefaultConfig() MRBetaMCPToolDefaultConfig`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *MRBetaMCPToolset) GetDefaultConfigOk() (*MRBetaMCPToolDefaultConfig, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *MRBetaMCPToolset) SetDefaultConfig(v MRBetaMCPToolDefaultConfig)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *MRBetaMCPToolset) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### SetDefaultConfigNil

`func (o *MRBetaMCPToolset) SetDefaultConfigNil(b bool)`

 SetDefaultConfigNil sets the value for DefaultConfig to be an explicit nil

### UnsetDefaultConfig
`func (o *MRBetaMCPToolset) UnsetDefaultConfig()`

UnsetDefaultConfig ensures that no value is present for DefaultConfig, not even an explicit nil
### GetTools

`func (o *MRBetaMCPToolset) GetTools() []MRBetaMCPToolParam`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MRBetaMCPToolset) GetToolsOk() (*[]MRBetaMCPToolParam, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MRBetaMCPToolset) SetTools(v []MRBetaMCPToolParam)`

SetTools sets Tools field to given value.

### HasTools

`func (o *MRBetaMCPToolset) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *MRBetaMCPToolset) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *MRBetaMCPToolset) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


