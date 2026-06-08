# ToolCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Function** | [**MSGFunction**](MSGFunction.md) |  | 
**Type** | **string** |  | 
**Custom** | [**MSGCustom**](MSGCustom.md) |  | 

## Methods

### NewToolCallsInner

`func NewToolCallsInner(id string, function MSGFunction, type_ string, custom MSGCustom, ) *ToolCallsInner`

NewToolCallsInner instantiates a new ToolCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCallsInnerWithDefaults

`func NewToolCallsInnerWithDefaults() *ToolCallsInner`

NewToolCallsInnerWithDefaults instantiates a new ToolCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolCallsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolCallsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolCallsInner) SetId(v string)`

SetId sets Id field to given value.


### GetFunction

`func (o *ToolCallsInner) GetFunction() MSGFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ToolCallsInner) GetFunctionOk() (*MSGFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ToolCallsInner) SetFunction(v MSGFunction)`

SetFunction sets Function field to given value.


### GetType

`func (o *ToolCallsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolCallsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolCallsInner) SetType(v string)`

SetType sets Type field to given value.


### GetCustom

`func (o *ToolCallsInner) GetCustom() MSGCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ToolCallsInner) GetCustomOk() (*MSGCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ToolCallsInner) SetCustom(v MSGCustom)`

SetCustom sets Custom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


