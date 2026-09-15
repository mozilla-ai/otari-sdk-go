# MRBetaWebSearchToolResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content11**](Content11.md) |  | 
**ToolUseId** | **string** |  | 
**Type** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 

## Methods

### NewMRBetaWebSearchToolResultBlock

`func NewMRBetaWebSearchToolResultBlock(content Content11, toolUseId string, type_ string, ) *MRBetaWebSearchToolResultBlock`

NewMRBetaWebSearchToolResultBlock instantiates a new MRBetaWebSearchToolResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaWebSearchToolResultBlockWithDefaults

`func NewMRBetaWebSearchToolResultBlockWithDefaults() *MRBetaWebSearchToolResultBlock`

NewMRBetaWebSearchToolResultBlockWithDefaults instantiates a new MRBetaWebSearchToolResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaWebSearchToolResultBlock) GetContent() Content11`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaWebSearchToolResultBlock) GetContentOk() (*Content11, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaWebSearchToolResultBlock) SetContent(v Content11)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *MRBetaWebSearchToolResultBlock) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *MRBetaWebSearchToolResultBlock) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *MRBetaWebSearchToolResultBlock) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetType

`func (o *MRBetaWebSearchToolResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaWebSearchToolResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaWebSearchToolResultBlock) SetType(v string)`

SetType sets Type field to given value.


### GetCaller

`func (o *MRBetaWebSearchToolResultBlock) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MRBetaWebSearchToolResultBlock) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MRBetaWebSearchToolResultBlock) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MRBetaWebSearchToolResultBlock) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MRBetaWebSearchToolResultBlock) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MRBetaWebSearchToolResultBlock) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


