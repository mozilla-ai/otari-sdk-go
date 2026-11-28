# AnthropicErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | The refusal&#39;s &#x60;Otari-Error-Code&#x60;, where it has one. | [optional] 
**Error** | [**AnthropicError**](AnthropicError.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "error"]

## Methods

### NewAnthropicErrorResponse

`func NewAnthropicErrorResponse(error_ AnthropicError, ) *AnthropicErrorResponse`

NewAnthropicErrorResponse instantiates a new AnthropicErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicErrorResponseWithDefaults

`func NewAnthropicErrorResponseWithDefaults() *AnthropicErrorResponse`

NewAnthropicErrorResponseWithDefaults instantiates a new AnthropicErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AnthropicErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AnthropicErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AnthropicErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AnthropicErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AnthropicErrorResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AnthropicErrorResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetError

`func (o *AnthropicErrorResponse) GetError() AnthropicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnthropicErrorResponse) GetErrorOk() (*AnthropicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnthropicErrorResponse) SetError(v AnthropicError)`

SetError sets Error field to given value.


### GetType

`func (o *AnthropicErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnthropicErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnthropicErrorResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnthropicErrorResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


