# MSGChatCompletionMessageFunctionToolCallParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Function** | [**MSGFunction**](MSGFunction.md) |  | 
**Type** | **string** |  | 

## Methods

### NewMSGChatCompletionMessageFunctionToolCallParam

`func NewMSGChatCompletionMessageFunctionToolCallParam(id string, function MSGFunction, type_ string, ) *MSGChatCompletionMessageFunctionToolCallParam`

NewMSGChatCompletionMessageFunctionToolCallParam instantiates a new MSGChatCompletionMessageFunctionToolCallParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionMessageFunctionToolCallParamWithDefaults

`func NewMSGChatCompletionMessageFunctionToolCallParamWithDefaults() *MSGChatCompletionMessageFunctionToolCallParam`

NewMSGChatCompletionMessageFunctionToolCallParamWithDefaults instantiates a new MSGChatCompletionMessageFunctionToolCallParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MSGChatCompletionMessageFunctionToolCallParam) SetId(v string)`

SetId sets Id field to given value.


### GetFunction

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetFunction() MSGFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetFunctionOk() (*MSGFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *MSGChatCompletionMessageFunctionToolCallParam) SetFunction(v MSGFunction)`

SetFunction sets Function field to given value.


### GetType

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MSGChatCompletionMessageFunctionToolCallParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MSGChatCompletionMessageFunctionToolCallParam) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


