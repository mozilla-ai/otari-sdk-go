# OrganizationModelPricingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cached-input tokens | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cache-write (creation) tokens | [optional] 
**EffectiveFrom** | **time.Time** | ISO 8601 datetime from which this rate applies, inclusive. Required on a replacement, so an omitted value cannot silently move a stored period to the present. | 
**EffectiveTo** | Pointer to **NullableTime** | ISO 8601 datetime at which this rate stops applying, exclusive. Null leaves it open ended. Because the end is exclusive, the next period may begin at exactly this instant without overlapping. | [optional] 
**InputPricePerMillion** | **float32** | Price per 1M of the rate&#39;s unit: input tokens, requests, images or characters | 
**OutputPricePerMillion** | **float32** | Price per 1M output tokens; unused by a rate per request or per image | 
**PricingTiers** | Pointer to [**[]PricingTier**](PricingTier.md) | Whole-request context thresholds. Fields omitted by a tier inherit the base rate. | [optional] 
**Rates** | Pointer to [**NullableRatesDocumentInput**](RatesDocumentInput.md) | Every price of the row as a models.dev cost block: each modality&#39;s input and output price with its unit, cache and overrides, web_search beside the modalities, and multipliers. A price is per its unit as published (one image, one million tokens, one thousand searches). Given instead of the per-million fields, which are then read out of it; a row always reads back with both. | [optional] 
**Unit** | Pointer to **string** | What the rates are per: token for a model, request, image or character for a non-token endpoint. A completion model priced per request is charged one request per successful call. Every rate is USD per million of its unit. A rate in a unit that has no meter yet is stored and prices nothing. The plural names tokens, requests and images are still read for one release. | [optional] [default to "token"]

## Methods

### NewOrganizationModelPricingUpdate

`func NewOrganizationModelPricingUpdate(effectiveFrom time.Time, inputPricePerMillion float32, outputPricePerMillion float32, ) *OrganizationModelPricingUpdate`

NewOrganizationModelPricingUpdate instantiates a new OrganizationModelPricingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationModelPricingUpdateWithDefaults

`func NewOrganizationModelPricingUpdateWithDefaults() *OrganizationModelPricingUpdate`

NewOrganizationModelPricingUpdateWithDefaults instantiates a new OrganizationModelPricingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *OrganizationModelPricingUpdate) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *OrganizationModelPricingUpdate) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *OrganizationModelPricingUpdate) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *OrganizationModelPricingUpdate) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *OrganizationModelPricingUpdate) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *OrganizationModelPricingUpdate) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *OrganizationModelPricingUpdate) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *OrganizationModelPricingUpdate) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *OrganizationModelPricingUpdate) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *OrganizationModelPricingUpdate) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *OrganizationModelPricingUpdate) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *OrganizationModelPricingUpdate) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetEffectiveFrom

`func (o *OrganizationModelPricingUpdate) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *OrganizationModelPricingUpdate) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *OrganizationModelPricingUpdate) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *OrganizationModelPricingUpdate) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *OrganizationModelPricingUpdate) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *OrganizationModelPricingUpdate) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *OrganizationModelPricingUpdate) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *OrganizationModelPricingUpdate) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *OrganizationModelPricingUpdate) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetInputPricePerMillion

`func (o *OrganizationModelPricingUpdate) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *OrganizationModelPricingUpdate) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *OrganizationModelPricingUpdate) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetOutputPricePerMillion

`func (o *OrganizationModelPricingUpdate) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *OrganizationModelPricingUpdate) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *OrganizationModelPricingUpdate) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.


### GetPricingTiers

`func (o *OrganizationModelPricingUpdate) GetPricingTiers() []PricingTier`

GetPricingTiers returns the PricingTiers field if non-nil, zero value otherwise.

### GetPricingTiersOk

`func (o *OrganizationModelPricingUpdate) GetPricingTiersOk() (*[]PricingTier, bool)`

GetPricingTiersOk returns a tuple with the PricingTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTiers

`func (o *OrganizationModelPricingUpdate) SetPricingTiers(v []PricingTier)`

SetPricingTiers sets PricingTiers field to given value.

### HasPricingTiers

`func (o *OrganizationModelPricingUpdate) HasPricingTiers() bool`

HasPricingTiers returns a boolean if a field has been set.

### SetPricingTiersNil

`func (o *OrganizationModelPricingUpdate) SetPricingTiersNil(b bool)`

 SetPricingTiersNil sets the value for PricingTiers to be an explicit nil

### UnsetPricingTiers
`func (o *OrganizationModelPricingUpdate) UnsetPricingTiers()`

UnsetPricingTiers ensures that no value is present for PricingTiers, not even an explicit nil
### GetRates

`func (o *OrganizationModelPricingUpdate) GetRates() RatesDocumentInput`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *OrganizationModelPricingUpdate) GetRatesOk() (*RatesDocumentInput, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *OrganizationModelPricingUpdate) SetRates(v RatesDocumentInput)`

SetRates sets Rates field to given value.

### HasRates

`func (o *OrganizationModelPricingUpdate) HasRates() bool`

HasRates returns a boolean if a field has been set.

### SetRatesNil

`func (o *OrganizationModelPricingUpdate) SetRatesNil(b bool)`

 SetRatesNil sets the value for Rates to be an explicit nil

### UnsetRates
`func (o *OrganizationModelPricingUpdate) UnsetRates()`

UnsetRates ensures that no value is present for Rates, not even an explicit nil
### GetUnit

`func (o *OrganizationModelPricingUpdate) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OrganizationModelPricingUpdate) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OrganizationModelPricingUpdate) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OrganizationModelPricingUpdate) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


