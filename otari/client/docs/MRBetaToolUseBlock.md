# MRBetaToolUseBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Input** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**ToolsetName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMRBetaToolUseBlock

`func NewMRBetaToolUseBlock(id string, input map[string]interface{}, name string, type_ string, ) *MRBetaToolUseBlock`

NewMRBetaToolUseBlock instantiates a new MRBetaToolUseBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaToolUseBlockWithDefaults

`func NewMRBetaToolUseBlockWithDefaults() *MRBetaToolUseBlock`

NewMRBetaToolUseBlockWithDefaults instantiates a new MRBetaToolUseBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MRBetaToolUseBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MRBetaToolUseBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MRBetaToolUseBlock) SetId(v string)`

SetId sets Id field to given value.


### GetInput

`func (o *MRBetaToolUseBlock) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *MRBetaToolUseBlock) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *MRBetaToolUseBlock) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetName

`func (o *MRBetaToolUseBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaToolUseBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaToolUseBlock) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaToolUseBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaToolUseBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaToolUseBlock) SetType(v string)`

SetType sets Type field to given value.


### GetCaller

`func (o *MRBetaToolUseBlock) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *MRBetaToolUseBlock) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *MRBetaToolUseBlock) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *MRBetaToolUseBlock) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *MRBetaToolUseBlock) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *MRBetaToolUseBlock) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetToolsetName

`func (o *MRBetaToolUseBlock) GetToolsetName() string`

GetToolsetName returns the ToolsetName field if non-nil, zero value otherwise.

### GetToolsetNameOk

`func (o *MRBetaToolUseBlock) GetToolsetNameOk() (*string, bool)`

GetToolsetNameOk returns a tuple with the ToolsetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsetName

`func (o *MRBetaToolUseBlock) SetToolsetName(v string)`

SetToolsetName sets ToolsetName field to given value.

### HasToolsetName

`func (o *MRBetaToolUseBlock) HasToolsetName() bool`

HasToolsetName returns a boolean if a field has been set.

### SetToolsetNameNil

`func (o *MRBetaToolUseBlock) SetToolsetNameNil(b bool)`

 SetToolsetNameNil sets the value for ToolsetName to be an explicit nil

### UnsetToolsetName
`func (o *MRBetaToolUseBlock) UnsetToolsetName()`

UnsetToolsetName ensures that no value is present for ToolsetName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


