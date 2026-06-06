# CCChatCompletionMessageToolCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Function** | [**CCFunction**](CCFunction.md) |  | 
**Type** | **string** |  | 
**Custom** | [**CCCustom**](CCCustom.md) |  | 

## Methods

### NewCCChatCompletionMessageToolCallsInner

`func NewCCChatCompletionMessageToolCallsInner(id string, function CCFunction, type_ string, custom CCCustom, ) *CCChatCompletionMessageToolCallsInner`

NewCCChatCompletionMessageToolCallsInner instantiates a new CCChatCompletionMessageToolCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChatCompletionMessageToolCallsInnerWithDefaults

`func NewCCChatCompletionMessageToolCallsInnerWithDefaults() *CCChatCompletionMessageToolCallsInner`

NewCCChatCompletionMessageToolCallsInnerWithDefaults instantiates a new CCChatCompletionMessageToolCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CCChatCompletionMessageToolCallsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCChatCompletionMessageToolCallsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCChatCompletionMessageToolCallsInner) SetId(v string)`

SetId sets Id field to given value.


### GetFunction

`func (o *CCChatCompletionMessageToolCallsInner) GetFunction() CCFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CCChatCompletionMessageToolCallsInner) GetFunctionOk() (*CCFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CCChatCompletionMessageToolCallsInner) SetFunction(v CCFunction)`

SetFunction sets Function field to given value.


### GetType

`func (o *CCChatCompletionMessageToolCallsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CCChatCompletionMessageToolCallsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CCChatCompletionMessageToolCallsInner) SetType(v string)`

SetType sets Type field to given value.


### GetCustom

`func (o *CCChatCompletionMessageToolCallsInner) GetCustom() CCCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CCChatCompletionMessageToolCallsInner) GetCustomOk() (*CCCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CCChatCompletionMessageToolCallsInner) SetCustom(v CCCustom)`

SetCustom sets Custom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


