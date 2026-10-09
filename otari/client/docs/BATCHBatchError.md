# BATCHBatchError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Line** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**Message** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Param** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 

## Methods

### NewBATCHBatchError

`func NewBATCHBatchError() *BATCHBatchError`

NewBATCHBatchError instantiates a new BATCHBatchError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBATCHBatchErrorWithDefaults

`func NewBATCHBatchErrorWithDefaults() *BATCHBatchError`

NewBATCHBatchErrorWithDefaults instantiates a new BATCHBatchError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BATCHBatchError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BATCHBatchError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BATCHBatchError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BATCHBatchError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BATCHBatchError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BATCHBatchError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLine

`func (o *BATCHBatchError) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *BATCHBatchError) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *BATCHBatchError) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *BATCHBatchError) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *BATCHBatchError) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *BATCHBatchError) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetMessage

`func (o *BATCHBatchError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BATCHBatchError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BATCHBatchError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BATCHBatchError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BATCHBatchError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BATCHBatchError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetParam

`func (o *BATCHBatchError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *BATCHBatchError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *BATCHBatchError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *BATCHBatchError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *BATCHBatchError) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *BATCHBatchError) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


