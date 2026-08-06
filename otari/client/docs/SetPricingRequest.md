# SetPricingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cached-input tokens | [optional] 
**CacheWrite1hPricePerMillion** | Pointer to **NullableFloat32** | Price per 1M Anthropic 1-hour cache-write tokens | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cache-write (creation) tokens | [optional] 
**EffectiveAt** | Pointer to **NullableTime** | ISO 8601 datetime from which this price applies. Defaults to now if omitted. | [optional] 
**InputPricePerMillion** | **float32** | Price per 1M input tokens | 
**ModelKey** | **string** | Model identifier in format &#39;provider:model&#39; | 
**OutputPricePerMillion** | **float32** | Price per 1M output tokens | 
**PricingTiers** | Pointer to [**[]PricingTier**](PricingTier.md) | Whole-request context thresholds. Fields omitted by a tier inherit the base rate. | [optional] 

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

### GetCacheReadPricePerMillion

`func (o *SetPricingRequest) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *SetPricingRequest) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *SetPricingRequest) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *SetPricingRequest) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *SetPricingRequest) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *SetPricingRequest) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWrite1hPricePerMillion

`func (o *SetPricingRequest) GetCacheWrite1hPricePerMillion() float32`

GetCacheWrite1hPricePerMillion returns the CacheWrite1hPricePerMillion field if non-nil, zero value otherwise.

### GetCacheWrite1hPricePerMillionOk

`func (o *SetPricingRequest) GetCacheWrite1hPricePerMillionOk() (*float32, bool)`

GetCacheWrite1hPricePerMillionOk returns a tuple with the CacheWrite1hPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hPricePerMillion

`func (o *SetPricingRequest) SetCacheWrite1hPricePerMillion(v float32)`

SetCacheWrite1hPricePerMillion sets CacheWrite1hPricePerMillion field to given value.

### HasCacheWrite1hPricePerMillion

`func (o *SetPricingRequest) HasCacheWrite1hPricePerMillion() bool`

HasCacheWrite1hPricePerMillion returns a boolean if a field has been set.

### SetCacheWrite1hPricePerMillionNil

`func (o *SetPricingRequest) SetCacheWrite1hPricePerMillionNil(b bool)`

 SetCacheWrite1hPricePerMillionNil sets the value for CacheWrite1hPricePerMillion to be an explicit nil

### UnsetCacheWrite1hPricePerMillion
`func (o *SetPricingRequest) UnsetCacheWrite1hPricePerMillion()`

UnsetCacheWrite1hPricePerMillion ensures that no value is present for CacheWrite1hPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *SetPricingRequest) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *SetPricingRequest) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *SetPricingRequest) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *SetPricingRequest) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *SetPricingRequest) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *SetPricingRequest) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
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


### GetPricingTiers

`func (o *SetPricingRequest) GetPricingTiers() []PricingTier`

GetPricingTiers returns the PricingTiers field if non-nil, zero value otherwise.

### GetPricingTiersOk

`func (o *SetPricingRequest) GetPricingTiersOk() (*[]PricingTier, bool)`

GetPricingTiersOk returns a tuple with the PricingTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTiers

`func (o *SetPricingRequest) SetPricingTiers(v []PricingTier)`

SetPricingTiers sets PricingTiers field to given value.

### HasPricingTiers

`func (o *SetPricingRequest) HasPricingTiers() bool`

HasPricingTiers returns a boolean if a field has been set.

### SetPricingTiersNil

`func (o *SetPricingRequest) SetPricingTiersNil(b bool)`

 SetPricingTiersNil sets the value for PricingTiers to be an explicit nil

### UnsetPricingTiers
`func (o *SetPricingRequest) UnsetPricingTiers()`

UnsetPricingTiers ensures that no value is present for PricingTiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


