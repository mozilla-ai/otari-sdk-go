# MRServerToolUseBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Input** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRServerToolUseBlock

`func NewMRServerToolUseBlock(id string, input map[string]interface{}, name string, type_ string, ) *MRServerToolUseBlock`

NewMRServerToolUseBlock instantiates a new MRServerToolUseBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRServerToolUseBlockWithDefaults

`func NewMRServerToolUseBlockWithDefaults() *MRServerToolUseBlock`

NewMRServerToolUseBlockWithDefaults instantiates a new MRServerToolUseBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MRServerToolUseBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MRServerToolUseBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MRServerToolUseBlock) SetId(v string)`

SetId sets Id field to given value.


### GetCaller

`func (o *MRServerToolUseBlock) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MRServerToolUseBlock) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MRServerToolUseBlock) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MRServerToolUseBlock) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MRServerToolUseBlock) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MRServerToolUseBlock) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetInput

`func (o *MRServerToolUseBlock) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *MRServerToolUseBlock) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *MRServerToolUseBlock) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetName

`func (o *MRServerToolUseBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRServerToolUseBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRServerToolUseBlock) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRServerToolUseBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRServerToolUseBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRServerToolUseBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


