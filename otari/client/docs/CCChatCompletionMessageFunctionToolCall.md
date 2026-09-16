# CCChatCompletionMessageFunctionToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Function** | [**CCFunction**](CCFunction.md) |  | 
**Type** | **string** |  | 
**ExtraContent** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCCChatCompletionMessageFunctionToolCall

`func NewCCChatCompletionMessageFunctionToolCall(id string, function CCFunction, type_ string, ) *CCChatCompletionMessageFunctionToolCall`

NewCCChatCompletionMessageFunctionToolCall instantiates a new CCChatCompletionMessageFunctionToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChatCompletionMessageFunctionToolCallWithDefaults

`func NewCCChatCompletionMessageFunctionToolCallWithDefaults() *CCChatCompletionMessageFunctionToolCall`

NewCCChatCompletionMessageFunctionToolCallWithDefaults instantiates a new CCChatCompletionMessageFunctionToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CCChatCompletionMessageFunctionToolCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCChatCompletionMessageFunctionToolCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCChatCompletionMessageFunctionToolCall) SetId(v string)`

SetId sets Id field to given value.


### GetFunction

`func (o *CCChatCompletionMessageFunctionToolCall) GetFunction() CCFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *CCChatCompletionMessageFunctionToolCall) GetFunctionOk() (*CCFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *CCChatCompletionMessageFunctionToolCall) SetFunction(v CCFunction)`

SetFunction sets Function field to given value.


### GetType

`func (o *CCChatCompletionMessageFunctionToolCall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CCChatCompletionMessageFunctionToolCall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CCChatCompletionMessageFunctionToolCall) SetType(v string)`

SetType sets Type field to given value.


### GetExtraContent

`func (o *CCChatCompletionMessageFunctionToolCall) GetExtraContent() map[string]interface{}`

GetExtraContent returns the ExtraContent field if non-nil, zero value otherwise.

### GetExtraContentOk

`func (o *CCChatCompletionMessageFunctionToolCall) GetExtraContentOk() (*map[string]interface{}, bool)`

GetExtraContentOk returns a tuple with the ExtraContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraContent

`func (o *CCChatCompletionMessageFunctionToolCall) SetExtraContent(v map[string]interface{})`

SetExtraContent sets ExtraContent field to given value.

### HasExtraContent

`func (o *CCChatCompletionMessageFunctionToolCall) HasExtraContent() bool`

HasExtraContent returns a boolean if a field has been set.

### SetExtraContentNil

`func (o *CCChatCompletionMessageFunctionToolCall) SetExtraContentNil(b bool)`

 SetExtraContentNil sets the value for ExtraContent to be an explicit nil

### UnsetExtraContent
`func (o *CCChatCompletionMessageFunctionToolCall) UnsetExtraContent()`

UnsetExtraContent ensures that no value is present for ExtraContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


