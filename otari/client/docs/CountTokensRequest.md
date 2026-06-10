# CountTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheControl** | Pointer to **map[string]interface{}** |  | [optional] 
**Messages** | **[]map[string]interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Model** | **string** |  | 
**System** | Pointer to [**NullableSystem**](System.md) |  | [optional] 
**Thinking** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolChoice** | Pointer to **map[string]interface{}** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewCountTokensRequest

`func NewCountTokensRequest(messages []*map[string]interface{}, model string, ) *CountTokensRequest`

NewCountTokensRequest instantiates a new CountTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountTokensRequestWithDefaults

`func NewCountTokensRequestWithDefaults() *CountTokensRequest`

NewCountTokensRequestWithDefaults instantiates a new CountTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheControl

`func (o *CountTokensRequest) GetCacheControl() map[string]interface{}`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *CountTokensRequest) GetCacheControlOk() (*map[string]interface{}, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *CountTokensRequest) SetCacheControl(v map[string]interface{})`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *CountTokensRequest) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *CountTokensRequest) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *CountTokensRequest) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetMessages

`func (o *CountTokensRequest) GetMessages() []*map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CountTokensRequest) GetMessagesOk() (*[]*map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CountTokensRequest) SetMessages(v []*map[string]interface{})`

SetMessages sets Messages field to given value.


### GetMetadata

`func (o *CountTokensRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CountTokensRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CountTokensRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CountTokensRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CountTokensRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CountTokensRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *CountTokensRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CountTokensRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CountTokensRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSystem

`func (o *CountTokensRequest) GetSystem() System`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *CountTokensRequest) GetSystemOk() (*System, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *CountTokensRequest) SetSystem(v System)`

SetSystem sets System field to given value.

### HasSystem

`func (o *CountTokensRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *CountTokensRequest) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *CountTokensRequest) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetThinking

`func (o *CountTokensRequest) GetThinking() map[string]interface{}`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *CountTokensRequest) GetThinkingOk() (*map[string]interface{}, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *CountTokensRequest) SetThinking(v map[string]interface{})`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *CountTokensRequest) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### SetThinkingNil

`func (o *CountTokensRequest) SetThinkingNil(b bool)`

 SetThinkingNil sets the value for Thinking to be an explicit nil

### UnsetThinking
`func (o *CountTokensRequest) UnsetThinking()`

UnsetThinking ensures that no value is present for Thinking, not even an explicit nil
### GetToolChoice

`func (o *CountTokensRequest) GetToolChoice() map[string]interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *CountTokensRequest) GetToolChoiceOk() (*map[string]interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *CountTokensRequest) SetToolChoice(v map[string]interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *CountTokensRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *CountTokensRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *CountTokensRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *CountTokensRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CountTokensRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CountTokensRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *CountTokensRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *CountTokensRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *CountTokensRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


