# SetPricingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveAt** | Pointer to **NullableTime** | ISO 8601 datetime from which this price applies. Defaults to now if omitted. | [optional] 
**InputPricePerMillion** | **float32** | Price per 1M input tokens | 
**ModelKey** | **string** | Model identifier in format &#39;provider:model&#39; | 
**OutputPricePerMillion** | **float32** | Price per 1M output tokens | 

## Methods

### NewSetPricingRequest

`func NewSetPricingRequest(inputPricePerMillion float32, modelKey string, outputPricePerMillion float32, ) *SetPricingRequest`

NewSetPricingRequest instantiates a new SetPricingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPricingRequestWithDefaults

`func NewSetPricingRequestWithDefaults() *SetPricingRequest`

NewSetPricingRequestWithDefaults instantiates a new SetPricingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveAt

`func (o *SetPricingRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *SetPricingRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *SetPricingRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *SetPricingRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### SetEffectiveAtNil

`func (o *SetPricingRequest) SetEffectiveAtNil(b bool)`

 SetEffectiveAtNil sets the value for EffectiveAt to be an explicit nil

### UnsetEffectiveAt
`func (o *SetPricingRequest) UnsetEffectiveAt()`

UnsetEffectiveAt ensures that no value is present for EffectiveAt, not even an explicit nil
### GetInputPricePerMillion

`func (o *SetPricingRequest) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *SetPricingRequest) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *SetPricingRequest) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetModelKey

`func (o *SetPricingRequest) GetModelKey() string`

GetModelKey returns the ModelKey field if non-nil, zero value otherwise.

### GetModelKeyOk

`func (o *SetPricingRequest) GetModelKeyOk() (*string, bool)`

GetModelKeyOk returns a tuple with the ModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelKey

`func (o *SetPricingRequest) SetModelKey(v string)`

SetModelKey sets ModelKey field to given value.


### GetOutputPricePerMillion

`func (o *SetPricingRequest) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *SetPricingRequest) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *SetPricingRequest) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


