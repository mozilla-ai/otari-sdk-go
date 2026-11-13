# MRBetaResponseTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSchema** | [**MRBetaResponseToolInputSchema**](MRBetaResponseToolInputSchema.md) |  | 
**Name** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EagerInputStreaming** | Pointer to **NullableBool** | Only end users that are, or are not, blocked. | [optional] 
**InputExamples** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMRBetaResponseTool

`func NewMRBetaResponseTool(inputSchema MRBetaResponseToolInputSchema, name string, ) *MRBetaResponseTool`

NewMRBetaResponseTool instantiates a new MRBetaResponseTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaResponseToolWithDefaults

`func NewMRBetaResponseToolWithDefaults() *MRBetaResponseTool`

NewMRBetaResponseToolWithDefaults instantiates a new MRBetaResponseTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSchema

`func (o *MRBetaResponseTool) GetInputSchema() MRBetaResponseToolInputSchema`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *MRBetaResponseTool) GetInputSchemaOk() (*MRBetaResponseToolInputSchema, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *MRBetaResponseTool) SetInputSchema(v MRBetaResponseToolInputSchema)`

SetInputSchema sets InputSchema field to given value.


### GetName

`func (o *MRBetaResponseTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaResponseTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaResponseTool) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedCallers

`func (o *MRBetaResponseTool) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaResponseTool) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaResponseTool) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaResponseTool) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaResponseTool) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaResponseTool) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaResponseTool) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaResponseTool) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaResponseTool) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaResponseTool) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaResponseTool) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaResponseTool) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetDescription

`func (o *MRBetaResponseTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MRBetaResponseTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MRBetaResponseTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MRBetaResponseTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MRBetaResponseTool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MRBetaResponseTool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEagerInputStreaming

`func (o *MRBetaResponseTool) GetEagerInputStreaming() bool`

GetEagerInputStreaming returns the EagerInputStreaming field if non-nil, zero value otherwise.

### GetEagerInputStreamingOk

`func (o *MRBetaResponseTool) GetEagerInputStreamingOk() (*bool, bool)`

GetEagerInputStreamingOk returns a tuple with the EagerInputStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEagerInputStreaming

`func (o *MRBetaResponseTool) SetEagerInputStreaming(v bool)`

SetEagerInputStreaming sets EagerInputStreaming field to given value.

### HasEagerInputStreaming

`func (o *MRBetaResponseTool) HasEagerInputStreaming() bool`

HasEagerInputStreaming returns a boolean if a field has been set.

### SetEagerInputStreamingNil

`func (o *MRBetaResponseTool) SetEagerInputStreamingNil(b bool)`

 SetEagerInputStreamingNil sets the value for EagerInputStreaming to be an explicit nil

### UnsetEagerInputStreaming
`func (o *MRBetaResponseTool) UnsetEagerInputStreaming()`

UnsetEagerInputStreaming ensures that no value is present for EagerInputStreaming, not even an explicit nil
### GetInputExamples

`func (o *MRBetaResponseTool) GetInputExamples() []map[string]interface{}`

GetInputExamples returns the InputExamples field if non-nil, zero value otherwise.

### GetInputExamplesOk

`func (o *MRBetaResponseTool) GetInputExamplesOk() (*[]map[string]interface{}, bool)`

GetInputExamplesOk returns a tuple with the InputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputExamples

`func (o *MRBetaResponseTool) SetInputExamples(v []map[string]interface{})`

SetInputExamples sets InputExamples field to given value.

### HasInputExamples

`func (o *MRBetaResponseTool) HasInputExamples() bool`

HasInputExamples returns a boolean if a field has been set.

### SetInputExamplesNil

`func (o *MRBetaResponseTool) SetInputExamplesNil(b bool)`

 SetInputExamplesNil sets the value for InputExamples to be an explicit nil

### UnsetInputExamples
`func (o *MRBetaResponseTool) UnsetInputExamples()`

UnsetInputExamples ensures that no value is present for InputExamples, not even an explicit nil
### GetStrict

`func (o *MRBetaResponseTool) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaResponseTool) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaResponseTool) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaResponseTool) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaResponseTool) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaResponseTool) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil
### GetType

`func (o *MRBetaResponseTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaResponseTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaResponseTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MRBetaResponseTool) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MRBetaResponseTool) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MRBetaResponseTool) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


