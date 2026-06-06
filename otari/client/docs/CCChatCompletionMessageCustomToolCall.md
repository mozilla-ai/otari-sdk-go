# CCChatCompletionMessageCustomToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Custom** | [**CCCustom**](CCCustom.md) |  | 
**Type** | **string** |  | 

## Methods

### NewCCChatCompletionMessageCustomToolCall

`func NewCCChatCompletionMessageCustomToolCall(id string, custom CCCustom, type_ string, ) *CCChatCompletionMessageCustomToolCall`

NewCCChatCompletionMessageCustomToolCall instantiates a new CCChatCompletionMessageCustomToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChatCompletionMessageCustomToolCallWithDefaults

`func NewCCChatCompletionMessageCustomToolCallWithDefaults() *CCChatCompletionMessageCustomToolCall`

NewCCChatCompletionMessageCustomToolCallWithDefaults instantiates a new CCChatCompletionMessageCustomToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CCChatCompletionMessageCustomToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCChatCompletionMessageCustomToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCChatCompletionMessageCustomToolCall) SetId(v string)`

SetId sets Id field to given value.


### GetCustom

`func (o *CCChatCompletionMessageCustomToolCall) GetCustom() CCCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *CCChatCompletionMessageCustomToolCall) GetCustomOk() (*CCCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *CCChatCompletionMessageCustomToolCall) SetCustom(v CCCustom)`

SetCustom sets Custom field to given value.


### GetType

`func (o *CCChatCompletionMessageCustomToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CCChatCompletionMessageCustomToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CCChatCompletionMessageCustomToolCall) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


