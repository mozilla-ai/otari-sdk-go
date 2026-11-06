# UnitRateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**NullableCacheRatesInput**](CacheRatesInput.md) |  | [optional] 
**Overrides** | Pointer to [**[]RateOverrideInput**](RateOverrideInput.md) |  | [optional] 
**Price** | Pointer to **NullableFloat32** | USD per the unit&#39;s published quantity. | [optional] 
**Reasoning** | Pointer to **NullableFloat32** | Price of reasoning tokens, where apart. | [optional] 
**Unit** | Pointer to **NullableString** | What the price is per; token when omitted. | [optional] 

## Methods

### NewUnitRateInput

`func NewUnitRateInput() *UnitRateInput`

NewUnitRateInput instantiates a new UnitRateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitRateInputWithDefaults

`func NewUnitRateInputWithDefaults() *UnitRateInput`

NewUnitRateInputWithDefaults instantiates a new UnitRateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *UnitRateInput) GetCache() CacheRatesInput`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *UnitRateInput) GetCacheOk() (*CacheRatesInput, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *UnitRateInput) SetCache(v CacheRatesInput)`

SetCache sets Cache field to given value.

### HasCache

`func (o *UnitRateInput) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *UnitRateInput) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *UnitRateInput) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetOverrides

`func (o *UnitRateInput) GetOverrides() []RateOverrideInput`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UnitRateInput) GetOverridesOk() (*[]RateOverrideInput, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UnitRateInput) SetOverrides(v []RateOverrideInput)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *UnitRateInput) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPrice

`func (o *UnitRateInput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UnitRateInput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UnitRateInput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UnitRateInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *UnitRateInput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *UnitRateInput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetReasoning

`func (o *UnitRateInput) GetReasoning() float32`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *UnitRateInput) GetReasoningOk() (*float32, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *UnitRateInput) SetReasoning(v float32)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *UnitRateInput) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *UnitRateInput) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *UnitRateInput) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetUnit

`func (o *UnitRateInput) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UnitRateInput) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UnitRateInput) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UnitRateInput) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *UnitRateInput) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *UnitRateInput) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


