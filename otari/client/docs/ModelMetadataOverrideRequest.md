# ModelMetadataOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextWindow** | Pointer to **NullableInt32** |  | [optional] 
**InputModalities** | Pointer to **[]string** | What the model reads; &#x60;text&#x60; unless it says otherwise. | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**Reasoning** | Pointer to **NullableBool** | Whether the model reasons: &#x60;reasoning&#x60;. | [optional] 
**StructuredOutput** | Pointer to **NullableBool** | Whether the model takes &#x60;response_format&#x60;. Never offered where Otari cannot send it. | [optional] 
**ToolCall** | Pointer to **NullableBool** | Whether the model calls tools: &#x60;tools&#x60;, &#x60;tool_choice&#x60;. | [optional] 

## Methods

### NewModelMetadataOverrideRequest

`func NewModelMetadataOverrideRequest() *ModelMetadataOverrideRequest`

NewModelMetadataOverrideRequest instantiates a new ModelMetadataOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataOverrideRequestWithDefaults

`func NewModelMetadataOverrideRequestWithDefaults() *ModelMetadataOverrideRequest`

NewModelMetadataOverrideRequestWithDefaults instantiates a new ModelMetadataOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextWindow

`func (o *ModelMetadataOverrideRequest) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *ModelMetadataOverrideRequest) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *ModelMetadataOverrideRequest) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *ModelMetadataOverrideRequest) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *ModelMetadataOverrideRequest) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *ModelMetadataOverrideRequest) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetInputModalities

`func (o *ModelMetadataOverrideRequest) GetInputModalities() []string`

GetInputModalities returns the InputModalities field if non-nil, zero value otherwise.

### GetInputModalitiesOk

`func (o *ModelMetadataOverrideRequest) GetInputModalitiesOk() (*[]string, bool)`

GetInputModalitiesOk returns a tuple with the InputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputModalities

`func (o *ModelMetadataOverrideRequest) SetInputModalities(v []string)`

SetInputModalities sets InputModalities field to given value.

### HasInputModalities

`func (o *ModelMetadataOverrideRequest) HasInputModalities() bool`

HasInputModalities returns a boolean if a field has been set.

### SetInputModalitiesNil

`func (o *ModelMetadataOverrideRequest) SetInputModalitiesNil(b bool)`

 SetInputModalitiesNil sets the value for InputModalities to be an explicit nil

### UnsetInputModalities
`func (o *ModelMetadataOverrideRequest) UnsetInputModalities()`

UnsetInputModalities ensures that no value is present for InputModalities, not even an explicit nil
### GetMaxOutputTokens

`func (o *ModelMetadataOverrideRequest) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ModelMetadataOverrideRequest) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ModelMetadataOverrideRequest) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ModelMetadataOverrideRequest) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *ModelMetadataOverrideRequest) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *ModelMetadataOverrideRequest) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetReasoning

`func (o *ModelMetadataOverrideRequest) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ModelMetadataOverrideRequest) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ModelMetadataOverrideRequest) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ModelMetadataOverrideRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *ModelMetadataOverrideRequest) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *ModelMetadataOverrideRequest) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetStructuredOutput

`func (o *ModelMetadataOverrideRequest) GetStructuredOutput() bool`

GetStructuredOutput returns the StructuredOutput field if non-nil, zero value otherwise.

### GetStructuredOutputOk

`func (o *ModelMetadataOverrideRequest) GetStructuredOutputOk() (*bool, bool)`

GetStructuredOutputOk returns a tuple with the StructuredOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredOutput

`func (o *ModelMetadataOverrideRequest) SetStructuredOutput(v bool)`

SetStructuredOutput sets StructuredOutput field to given value.

### HasStructuredOutput

`func (o *ModelMetadataOverrideRequest) HasStructuredOutput() bool`

HasStructuredOutput returns a boolean if a field has been set.

### SetStructuredOutputNil

`func (o *ModelMetadataOverrideRequest) SetStructuredOutputNil(b bool)`

 SetStructuredOutputNil sets the value for StructuredOutput to be an explicit nil

### UnsetStructuredOutput
`func (o *ModelMetadataOverrideRequest) UnsetStructuredOutput()`

UnsetStructuredOutput ensures that no value is present for StructuredOutput, not even an explicit nil
### GetToolCall

`func (o *ModelMetadataOverrideRequest) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *ModelMetadataOverrideRequest) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *ModelMetadataOverrideRequest) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.

### HasToolCall

`func (o *ModelMetadataOverrideRequest) HasToolCall() bool`

HasToolCall returns a boolean if a field has been set.

### SetToolCallNil

`func (o *ModelMetadataOverrideRequest) SetToolCallNil(b bool)`

 SetToolCallNil sets the value for ToolCall to be an explicit nil

### UnsetToolCall
`func (o *ModelMetadataOverrideRequest) UnsetToolCall()`

UnsetToolCall ensures that no value is present for ToolCall, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


