# ModelMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | Pointer to **bool** |  | [optional] [default to false]
**ContextWindow** | Pointer to **NullableInt32** |  | [optional] 
**CostInput** | Pointer to **NullableFloat32** |  | [optional] 
**CostOutput** | Pointer to **NullableFloat32** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **NullableString** |  | [optional] 
**Family** | Pointer to **NullableString** |  | [optional] 
**InputModalities** | Pointer to **[]string** |  | [optional] 
**KnowledgeCutoff** | Pointer to **NullableString** |  | [optional] 
**LastUpdated** | Pointer to **NullableString** |  | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OpenWeights** | Pointer to **bool** |  | [optional] [default to false]
**OutputModalities** | Pointer to **[]string** |  | [optional] 
**Reasoning** | Pointer to **bool** |  | [optional] [default to false]
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**StructuredOutput** | Pointer to **bool** |  | [optional] [default to false]
**Temperature** | Pointer to **bool** |  | [optional] [default to false]
**ToolCall** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewModelMetadata

`func NewModelMetadata() *ModelMetadata`

NewModelMetadata instantiates a new ModelMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataWithDefaults

`func NewModelMetadataWithDefaults() *ModelMetadata`

NewModelMetadataWithDefaults instantiates a new ModelMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *ModelMetadata) GetAttachment() bool`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *ModelMetadata) GetAttachmentOk() (*bool, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *ModelMetadata) SetAttachment(v bool)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *ModelMetadata) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetContextWindow

`func (o *ModelMetadata) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *ModelMetadata) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *ModelMetadata) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *ModelMetadata) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *ModelMetadata) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *ModelMetadata) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetCostInput

`func (o *ModelMetadata) GetCostInput() float32`

GetCostInput returns the CostInput field if non-nil, zero value otherwise.

### GetCostInputOk

`func (o *ModelMetadata) GetCostInputOk() (*float32, bool)`

GetCostInputOk returns a tuple with the CostInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostInput

`func (o *ModelMetadata) SetCostInput(v float32)`

SetCostInput sets CostInput field to given value.

### HasCostInput

`func (o *ModelMetadata) HasCostInput() bool`

HasCostInput returns a boolean if a field has been set.

### SetCostInputNil

`func (o *ModelMetadata) SetCostInputNil(b bool)`

 SetCostInputNil sets the value for CostInput to be an explicit nil

### UnsetCostInput
`func (o *ModelMetadata) UnsetCostInput()`

UnsetCostInput ensures that no value is present for CostInput, not even an explicit nil
### GetCostOutput

`func (o *ModelMetadata) GetCostOutput() float32`

GetCostOutput returns the CostOutput field if non-nil, zero value otherwise.

### GetCostOutputOk

`func (o *ModelMetadata) GetCostOutputOk() (*float32, bool)`

GetCostOutputOk returns a tuple with the CostOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOutput

`func (o *ModelMetadata) SetCostOutput(v float32)`

SetCostOutput sets CostOutput field to given value.

### HasCostOutput

`func (o *ModelMetadata) HasCostOutput() bool`

HasCostOutput returns a boolean if a field has been set.

### SetCostOutputNil

`func (o *ModelMetadata) SetCostOutputNil(b bool)`

 SetCostOutputNil sets the value for CostOutput to be an explicit nil

### UnsetCostOutput
`func (o *ModelMetadata) UnsetCostOutput()`

UnsetCostOutput ensures that no value is present for CostOutput, not even an explicit nil
### GetDeprecated

`func (o *ModelMetadata) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *ModelMetadata) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *ModelMetadata) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *ModelMetadata) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *ModelMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelMetadata) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelMetadata) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFamily

`func (o *ModelMetadata) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ModelMetadata) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ModelMetadata) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ModelMetadata) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *ModelMetadata) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *ModelMetadata) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetInputModalities

`func (o *ModelMetadata) GetInputModalities() []string`

GetInputModalities returns the InputModalities field if non-nil, zero value otherwise.

### GetInputModalitiesOk

`func (o *ModelMetadata) GetInputModalitiesOk() (*[]string, bool)`

GetInputModalitiesOk returns a tuple with the InputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputModalities

`func (o *ModelMetadata) SetInputModalities(v []string)`

SetInputModalities sets InputModalities field to given value.

### HasInputModalities

`func (o *ModelMetadata) HasInputModalities() bool`

HasInputModalities returns a boolean if a field has been set.

### GetKnowledgeCutoff

`func (o *ModelMetadata) GetKnowledgeCutoff() string`

GetKnowledgeCutoff returns the KnowledgeCutoff field if non-nil, zero value otherwise.

### GetKnowledgeCutoffOk

`func (o *ModelMetadata) GetKnowledgeCutoffOk() (*string, bool)`

GetKnowledgeCutoffOk returns a tuple with the KnowledgeCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCutoff

`func (o *ModelMetadata) SetKnowledgeCutoff(v string)`

SetKnowledgeCutoff sets KnowledgeCutoff field to given value.

### HasKnowledgeCutoff

`func (o *ModelMetadata) HasKnowledgeCutoff() bool`

HasKnowledgeCutoff returns a boolean if a field has been set.

### SetKnowledgeCutoffNil

`func (o *ModelMetadata) SetKnowledgeCutoffNil(b bool)`

 SetKnowledgeCutoffNil sets the value for KnowledgeCutoff to be an explicit nil

### UnsetKnowledgeCutoff
`func (o *ModelMetadata) UnsetKnowledgeCutoff()`

UnsetKnowledgeCutoff ensures that no value is present for KnowledgeCutoff, not even an explicit nil
### GetLastUpdated

`func (o *ModelMetadata) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ModelMetadata) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ModelMetadata) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ModelMetadata) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ModelMetadata) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ModelMetadata) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetMaxOutputTokens

`func (o *ModelMetadata) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ModelMetadata) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ModelMetadata) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ModelMetadata) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *ModelMetadata) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *ModelMetadata) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetName

`func (o *ModelMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelMetadata) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelMetadata) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOpenWeights

`func (o *ModelMetadata) GetOpenWeights() bool`

GetOpenWeights returns the OpenWeights field if non-nil, zero value otherwise.

### GetOpenWeightsOk

`func (o *ModelMetadata) GetOpenWeightsOk() (*bool, bool)`

GetOpenWeightsOk returns a tuple with the OpenWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWeights

`func (o *ModelMetadata) SetOpenWeights(v bool)`

SetOpenWeights sets OpenWeights field to given value.

### HasOpenWeights

`func (o *ModelMetadata) HasOpenWeights() bool`

HasOpenWeights returns a boolean if a field has been set.

### GetOutputModalities

`func (o *ModelMetadata) GetOutputModalities() []string`

GetOutputModalities returns the OutputModalities field if non-nil, zero value otherwise.

### GetOutputModalitiesOk

`func (o *ModelMetadata) GetOutputModalitiesOk() (*[]string, bool)`

GetOutputModalitiesOk returns a tuple with the OutputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputModalities

`func (o *ModelMetadata) SetOutputModalities(v []string)`

SetOutputModalities sets OutputModalities field to given value.

### HasOutputModalities

`func (o *ModelMetadata) HasOutputModalities() bool`

HasOutputModalities returns a boolean if a field has been set.

### GetReasoning

`func (o *ModelMetadata) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ModelMetadata) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ModelMetadata) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ModelMetadata) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ModelMetadata) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ModelMetadata) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ModelMetadata) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ModelMetadata) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ModelMetadata) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ModelMetadata) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetStructuredOutput

`func (o *ModelMetadata) GetStructuredOutput() bool`

GetStructuredOutput returns the StructuredOutput field if non-nil, zero value otherwise.

### GetStructuredOutputOk

`func (o *ModelMetadata) GetStructuredOutputOk() (*bool, bool)`

GetStructuredOutputOk returns a tuple with the StructuredOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredOutput

`func (o *ModelMetadata) SetStructuredOutput(v bool)`

SetStructuredOutput sets StructuredOutput field to given value.

### HasStructuredOutput

`func (o *ModelMetadata) HasStructuredOutput() bool`

HasStructuredOutput returns a boolean if a field has been set.

### GetTemperature

`func (o *ModelMetadata) GetTemperature() bool`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ModelMetadata) GetTemperatureOk() (*bool, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ModelMetadata) SetTemperature(v bool)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ModelMetadata) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetToolCall

`func (o *ModelMetadata) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *ModelMetadata) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *ModelMetadata) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.

### HasToolCall

`func (o *ModelMetadata) HasToolCall() bool`

HasToolCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


