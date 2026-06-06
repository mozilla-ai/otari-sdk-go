# MRWebFetchToolResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Content** | [**Content7**](Content7.md) |  | 
**ToolUseId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRWebFetchToolResultBlock

`func NewMRWebFetchToolResultBlock(content Content7, toolUseId string, type_ string, ) *MRWebFetchToolResultBlock`

NewMRWebFetchToolResultBlock instantiates a new MRWebFetchToolResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRWebFetchToolResultBlockWithDefaults

`func NewMRWebFetchToolResultBlockWithDefaults() *MRWebFetchToolResultBlock`

NewMRWebFetchToolResultBlockWithDefaults instantiates a new MRWebFetchToolResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaller

`func (o *MRWebFetchToolResultBlock) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MRWebFetchToolResultBlock) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MRWebFetchToolResultBlock) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MRWebFetchToolResultBlock) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MRWebFetchToolResultBlock) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MRWebFetchToolResultBlock) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetContent

`func (o *MRWebFetchToolResultBlock) GetContent() Content7`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRWebFetchToolResultBlock) GetContentOk() (*Content7, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRWebFetchToolResultBlock) SetContent(v Content7)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *MRWebFetchToolResultBlock) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *MRWebFetchToolResultBlock) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *MRWebFetchToolResultBlock) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetType

`func (o *MRWebFetchToolResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRWebFetchToolResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRWebFetchToolResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


