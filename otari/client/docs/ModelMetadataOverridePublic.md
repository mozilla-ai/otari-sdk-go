# ModelMetadataOverridePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextWindow** | Pointer to **NullableInt32** |  | [optional] 
**InputModalities** | Pointer to **[]string** | What the model reads; &#x60;text&#x60; unless it says otherwise. | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**Reasoning** | Pointer to **NullableBool** | Whether the model reasons: &#x60;reasoning&#x60;. | [optional] 
**Selector** | **string** | The &#x60;provider:model&#x60; the override describes. | 
**StructuredOutput** | Pointer to **NullableBool** | Whether the model takes &#x60;response_format&#x60;. Never offered where Otari cannot send it. | [optional] 
**ToolCall** | Pointer to **NullableBool** | Whether the model calls tools: &#x60;tools&#x60;, &#x60;tool_choice&#x60;. | [optional] 

## Methods

### NewModelMetadataOverridePublic

`func NewModelMetadataOverridePublic(selector string, ) *ModelMetadataOverridePublic`

NewModelMetadataOverridePublic instantiates a new ModelMetadataOverridePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataOverridePublicWithDefaults

`func NewModelMetadataOverridePublicWithDefaults() *ModelMetadataOverridePublic`

NewModelMetadataOverridePublicWithDefaults instantiates a new ModelMetadataOverridePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextWindow

`func (o *ModelMetadataOverridePublic) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *ModelMetadataOverridePublic) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *ModelMetadataOverridePublic) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *ModelMetadataOverridePublic) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *ModelMetadataOverridePublic) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *ModelMetadataOverridePublic) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetInputModalities

`func (o *ModelMetadataOverridePublic) GetInputModalities() []string`

GetInputModalities returns the InputModalities field if non-nil, zero value otherwise.

### GetInputModalitiesOk

`func (o *ModelMetadataOverridePublic) GetInputModalitiesOk() (*[]string, bool)`

GetInputModalitiesOk returns a tuple with the InputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputModalities

`func (o *ModelMetadataOverridePublic) SetInputModalities(v []string)`

SetInputModalities sets InputModalities field to given value.

### HasInputModalities

`func (o *ModelMetadataOverridePublic) HasInputModalities() bool`

HasInputModalities returns a boolean if a field has been set.

### SetInputModalitiesNil

`func (o *ModelMetadataOverridePublic) SetInputModalitiesNil(b bool)`

 SetInputModalitiesNil sets the value for InputModalities to be an explicit nil

### UnsetInputModalities
`func (o *ModelMetadataOverridePublic) UnsetInputModalities()`

UnsetInputModalities ensures that no value is present for InputModalities, not even an explicit nil
### GetMaxOutputTokens

`func (o *ModelMetadataOverridePublic) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ModelMetadataOverridePublic) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ModelMetadataOverridePublic) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ModelMetadataOverridePublic) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *ModelMetadataOverridePublic) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *ModelMetadataOverridePublic) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetReasoning

`func (o *ModelMetadataOverridePublic) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ModelMetadataOverridePublic) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ModelMetadataOverridePublic) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ModelMetadataOverridePublic) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *ModelMetadataOverridePublic) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *ModelMetadataOverridePublic) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetSelector

`func (o *ModelMetadataOverridePublic) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ModelMetadataOverridePublic) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ModelMetadataOverridePublic) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetStructuredOutput

`func (o *ModelMetadataOverridePublic) GetStructuredOutput() bool`

GetStructuredOutput returns the StructuredOutput field if non-nil, zero value otherwise.

### GetStructuredOutputOk

`func (o *ModelMetadataOverridePublic) GetStructuredOutputOk() (*bool, bool)`

GetStructuredOutputOk returns a tuple with the StructuredOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredOutput

`func (o *ModelMetadataOverridePublic) SetStructuredOutput(v bool)`

SetStructuredOutput sets StructuredOutput field to given value.

### HasStructuredOutput

`func (o *ModelMetadataOverridePublic) HasStructuredOutput() bool`

HasStructuredOutput returns a boolean if a field has been set.

### SetStructuredOutputNil

`func (o *ModelMetadataOverridePublic) SetStructuredOutputNil(b bool)`

 SetStructuredOutputNil sets the value for StructuredOutput to be an explicit nil

### UnsetStructuredOutput
`func (o *ModelMetadataOverridePublic) UnsetStructuredOutput()`

UnsetStructuredOutput ensures that no value is present for StructuredOutput, not even an explicit nil
### GetToolCall

`func (o *ModelMetadataOverridePublic) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *ModelMetadataOverridePublic) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *ModelMetadataOverridePublic) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.

### HasToolCall

`func (o *ModelMetadataOverridePublic) HasToolCall() bool`

HasToolCall returns a boolean if a field has been set.

### SetToolCallNil

`func (o *ModelMetadataOverridePublic) SetToolCallNil(b bool)`

 SetToolCallNil sets the value for ToolCall to be an explicit nil

### UnsetToolCall
`func (o *ModelMetadataOverridePublic) UnsetToolCall()`

UnsetToolCall ensures that no value is present for ToolCall, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


