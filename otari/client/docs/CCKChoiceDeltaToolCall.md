# CCKChoiceDeltaToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Function** | Pointer to [**NullableCCKChoiceDeltaToolCallFunction**](CCKChoiceDeltaToolCallFunction.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCCKChoiceDeltaToolCall

`func NewCCKChoiceDeltaToolCall(index int32, ) *CCKChoiceDeltaToolCall`

NewCCKChoiceDeltaToolCall instantiates a new CCKChoiceDeltaToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKChoiceDeltaToolCallWithDefaults

`func NewCCKChoiceDeltaToolCallWithDefaults() *CCKChoiceDeltaToolCall`

NewCCKChoiceDeltaToolCallWithDefaults instantiates a new CCKChoiceDeltaToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *CCKChoiceDeltaToolCall) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CCKChoiceDeltaToolCall) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CCKChoiceDeltaToolCall) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetId

`func (o *CCKChoiceDeltaToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCKChoiceDeltaToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCKChoiceDeltaToolCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CCKChoiceDeltaToolCall) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CCKChoiceDeltaToolCall) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CCKChoiceDeltaToolCall) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFunction

`func (o *CCKChoiceDeltaToolCall) GetFunction() CCKChoiceDeltaToolCallFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CCKChoiceDeltaToolCall) GetFunctionOk() (*CCKChoiceDeltaToolCallFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CCKChoiceDeltaToolCall) SetFunction(v CCKChoiceDeltaToolCallFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *CCKChoiceDeltaToolCall) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### SetFunctionNil

`func (o *CCKChoiceDeltaToolCall) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *CCKChoiceDeltaToolCall) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil
### GetType

`func (o *CCKChoiceDeltaToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CCKChoiceDeltaToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CCKChoiceDeltaToolCall) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CCKChoiceDeltaToolCall) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CCKChoiceDeltaToolCall) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CCKChoiceDeltaToolCall) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


