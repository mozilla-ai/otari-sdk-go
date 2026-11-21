# AnthropicError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Type** | **string** | Anthropic&#39;s error type, such as &#x60;invalid_request_error&#x60; or &#x60;rate_limit_error&#x60;. | 

## Methods

### NewAnthropicError

`func NewAnthropicError(message string, type_ string, ) *AnthropicError`

NewAnthropicError instantiates a new AnthropicError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicErrorWithDefaults

`func NewAnthropicErrorWithDefaults() *AnthropicError`

NewAnthropicErrorWithDefaults instantiates a new AnthropicError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AnthropicError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnthropicError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnthropicError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *AnthropicError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnthropicError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnthropicError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


