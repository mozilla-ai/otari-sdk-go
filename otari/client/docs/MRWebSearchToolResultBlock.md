# MRWebSearchToolResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Content** | [**Content8**](Content8.md) |  | 
**ToolUseId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRWebSearchToolResultBlock

`func NewMRWebSearchToolResultBlock(content Content8, toolUseId string, type_ string, ) *MRWebSearchToolResultBlock`

NewMRWebSearchToolResultBlock instantiates a new MRWebSearchToolResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRWebSearchToolResultBlockWithDefaults

`func NewMRWebSearchToolResultBlockWithDefaults() *MRWebSearchToolResultBlock`

NewMRWebSearchToolResultBlockWithDefaults instantiates a new MRWebSearchToolResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaller

`func (o *MRWebSearchToolResultBlock) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MRWebSearchToolResultBlock) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MRWebSearchToolResultBlock) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MRWebSearchToolResultBlock) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MRWebSearchToolResultBlock) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MRWebSearchToolResultBlock) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetContent

`func (o *MRWebSearchToolResultBlock) GetContent() Content8`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRWebSearchToolResultBlock) GetContentOk() (*Content8, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRWebSearchToolResultBlock) SetContent(v Content8)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *MRWebSearchToolResultBlock) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *MRWebSearchToolResultBlock) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *MRWebSearchToolResultBlock) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetType

`func (o *MRWebSearchToolResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRWebSearchToolResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRWebSearchToolResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


