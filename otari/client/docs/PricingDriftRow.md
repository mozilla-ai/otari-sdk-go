# PricingDriftRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultInputPricePerMillion** | **NullableFloat32** | What genai-prices would meter this key at today. Null when the dataset does not know it. | 
**DefaultOutputPricePerMillion** | **NullableFloat32** |  | 
**DefaultReference** | **NullableString** | The genai-prices entry the default came from. | 
**EffectiveAt** | **string** |  | 
**InputDeltaPercent** | **NullableFloat32** | (stored - default) / default, as a percentage. | 
**InputPricePerMillion** | **float32** |  | 
**ModelKey** | **string** |  | 
**Origin** | **NullableString** |  | 
**OutputDeltaPercent** | **NullableFloat32** |  | 
**OutputPricePerMillion** | **float32** |  | 
**Unit** | **string** |  | 

## Methods

### NewPricingDriftRow

`func NewPricingDriftRow(defaultInputPricePerMillion NullableFloat32, defaultOutputPricePerMillion NullableFloat32, defaultReference NullableString, effectiveAt string, inputDeltaPercent NullableFloat32, inputPricePerMillion float32, modelKey string, origin NullableString, outputDeltaPercent NullableFloat32, outputPricePerMillion float32, unit string, ) *PricingDriftRow`

NewPricingDriftRow instantiates a new PricingDriftRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingDriftRowWithDefaults

`func NewPricingDriftRowWithDefaults() *PricingDriftRow`

NewPricingDriftRowWithDefaults instantiates a new PricingDriftRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultInputPricePerMillion

`func (o *PricingDriftRow) GetDefaultInputPricePerMillion() float32`

GetDefaultInputPricePerMillion returns the DefaultInputPricePerMillion field if non-nil, zero value otherwise.

### GetDefaultInputPricePerMillionOk

`func (o *PricingDriftRow) GetDefaultInputPricePerMillionOk() (*float32, bool)`

GetDefaultInputPricePerMillionOk returns a tuple with the DefaultInputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInputPricePerMillion

`func (o *PricingDriftRow) SetDefaultInputPricePerMillion(v float32)`

SetDefaultInputPricePerMillion sets DefaultInputPricePerMillion field to given value.


### SetDefaultInputPricePerMillionNil

`func (o *PricingDriftRow) SetDefaultInputPricePerMillionNil(b bool)`

 SetDefaultInputPricePerMillionNil sets the value for DefaultInputPricePerMillion to be an explicit nil

### UnsetDefaultInputPricePerMillion
`func (o *PricingDriftRow) UnsetDefaultInputPricePerMillion()`

UnsetDefaultInputPricePerMillion ensures that no value is present for DefaultInputPricePerMillion, not even an explicit nil
### GetDefaultOutputPricePerMillion

`func (o *PricingDriftRow) GetDefaultOutputPricePerMillion() float32`

GetDefaultOutputPricePerMillion returns the DefaultOutputPricePerMillion field if non-nil, zero value otherwise.

### GetDefaultOutputPricePerMillionOk

`func (o *PricingDriftRow) GetDefaultOutputPricePerMillionOk() (*float32, bool)`

GetDefaultOutputPricePerMillionOk returns a tuple with the DefaultOutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOutputPricePerMillion

`func (o *PricingDriftRow) SetDefaultOutputPricePerMillion(v float32)`

SetDefaultOutputPricePerMillion sets DefaultOutputPricePerMillion field to given value.


### SetDefaultOutputPricePerMillionNil

`func (o *PricingDriftRow) SetDefaultOutputPricePerMillionNil(b bool)`

 SetDefaultOutputPricePerMillionNil sets the value for DefaultOutputPricePerMillion to be an explicit nil

### UnsetDefaultOutputPricePerMillion
`func (o *PricingDriftRow) UnsetDefaultOutputPricePerMillion()`

UnsetDefaultOutputPricePerMillion ensures that no value is present for DefaultOutputPricePerMillion, not even an explicit nil
### GetDefaultReference

`func (o *PricingDriftRow) GetDefaultReference() string`

GetDefaultReference returns the DefaultReference field if non-nil, zero value otherwise.

### GetDefaultReferenceOk

`func (o *PricingDriftRow) GetDefaultReferenceOk() (*string, bool)`

GetDefaultReferenceOk returns a tuple with the DefaultReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReference

`func (o *PricingDriftRow) SetDefaultReference(v string)`

SetDefaultReference sets DefaultReference field to given value.


### SetDefaultReferenceNil

`func (o *PricingDriftRow) SetDefaultReferenceNil(b bool)`

 SetDefaultReferenceNil sets the value for DefaultReference to be an explicit nil

### UnsetDefaultReference
`func (o *PricingDriftRow) UnsetDefaultReference()`

UnsetDefaultReference ensures that no value is present for DefaultReference, not even an explicit nil
### GetEffectiveAt

`func (o *PricingDriftRow) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *PricingDriftRow) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *PricingDriftRow) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetInputDeltaPercent

`func (o *PricingDriftRow) GetInputDeltaPercent() float32`

GetInputDeltaPercent returns the InputDeltaPercent field if non-nil, zero value otherwise.

### GetInputDeltaPercentOk

`func (o *PricingDriftRow) GetInputDeltaPercentOk() (*float32, bool)`

GetInputDeltaPercentOk returns a tuple with the InputDeltaPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDeltaPercent

`func (o *PricingDriftRow) SetInputDeltaPercent(v float32)`

SetInputDeltaPercent sets InputDeltaPercent field to given value.


### SetInputDeltaPercentNil

`func (o *PricingDriftRow) SetInputDeltaPercentNil(b bool)`

 SetInputDeltaPercentNil sets the value for InputDeltaPercent to be an explicit nil

### UnsetInputDeltaPercent
`func (o *PricingDriftRow) UnsetInputDeltaPercent()`

UnsetInputDeltaPercent ensures that no value is present for InputDeltaPercent, not even an explicit nil
### GetInputPricePerMillion

`func (o *PricingDriftRow) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *PricingDriftRow) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *PricingDriftRow) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetModelKey

`func (o *PricingDriftRow) GetModelKey() string`

GetModelKey returns the ModelKey field if non-nil, zero value otherwise.

### GetModelKeyOk

`func (o *PricingDriftRow) GetModelKeyOk() (*string, bool)`

GetModelKeyOk returns a tuple with the ModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelKey

`func (o *PricingDriftRow) SetModelKey(v string)`

SetModelKey sets ModelKey field to given value.


### GetOrigin

`func (o *PricingDriftRow) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PricingDriftRow) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PricingDriftRow) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### SetOriginNil

`func (o *PricingDriftRow) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *PricingDriftRow) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetOutputDeltaPercent

`func (o *PricingDriftRow) GetOutputDeltaPercent() float32`

GetOutputDeltaPercent returns the OutputDeltaPercent field if non-nil, zero value otherwise.

### GetOutputDeltaPercentOk

`func (o *PricingDriftRow) GetOutputDeltaPercentOk() (*float32, bool)`

GetOutputDeltaPercentOk returns a tuple with the OutputDeltaPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDeltaPercent

`func (o *PricingDriftRow) SetOutputDeltaPercent(v float32)`

SetOutputDeltaPercent sets OutputDeltaPercent field to given value.


### SetOutputDeltaPercentNil

`func (o *PricingDriftRow) SetOutputDeltaPercentNil(b bool)`

 SetOutputDeltaPercentNil sets the value for OutputDeltaPercent to be an explicit nil

### UnsetOutputDeltaPercent
`func (o *PricingDriftRow) UnsetOutputDeltaPercent()`

UnsetOutputDeltaPercent ensures that no value is present for OutputDeltaPercent, not even an explicit nil
### GetOutputPricePerMillion

`func (o *PricingDriftRow) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *PricingDriftRow) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *PricingDriftRow) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.


### GetUnit

`func (o *PricingDriftRow) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PricingDriftRow) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PricingDriftRow) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


