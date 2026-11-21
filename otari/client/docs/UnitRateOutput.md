# UnitRateOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**NullableCacheRatesOutput**](CacheRatesOutput.md) |  | [optional] 
**Overrides** | Pointer to [**[]RateOverrideOutput**](RateOverrideOutput.md) |  | [optional] 
**Price** | Pointer to **NullableFloat32** | USD per the unit&#39;s published quantity. | [optional] 
**Reasoning** | Pointer to **NullableFloat32** | Price of reasoning tokens, where apart. | [optional] 
**Unit** | Pointer to **NullableString** | What the price is per; token when omitted. | [optional] 

## Methods

### NewUnitRateOutput

`func NewUnitRateOutput() *UnitRateOutput`

NewUnitRateOutput instantiates a new UnitRateOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitRateOutputWithDefaults

`func NewUnitRateOutputWithDefaults() *UnitRateOutput`

NewUnitRateOutputWithDefaults instantiates a new UnitRateOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *UnitRateOutput) GetCache() CacheRatesOutput`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *UnitRateOutput) GetCacheOk() (*CacheRatesOutput, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *UnitRateOutput) SetCache(v CacheRatesOutput)`

SetCache sets Cache field to given value.

### HasCache

`func (o *UnitRateOutput) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *UnitRateOutput) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *UnitRateOutput) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetOverrides

`func (o *UnitRateOutput) GetOverrides() []RateOverrideOutput`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UnitRateOutput) GetOverridesOk() (*[]RateOverrideOutput, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UnitRateOutput) SetOverrides(v []RateOverrideOutput)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *UnitRateOutput) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPrice

`func (o *UnitRateOutput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UnitRateOutput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UnitRateOutput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UnitRateOutput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *UnitRateOutput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *UnitRateOutput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetReasoning

`func (o *UnitRateOutput) GetReasoning() float32`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *UnitRateOutput) GetReasoningOk() (*float32, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *UnitRateOutput) SetReasoning(v float32)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *UnitRateOutput) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *UnitRateOutput) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *UnitRateOutput) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetUnit

`func (o *UnitRateOutput) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UnitRateOutput) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UnitRateOutput) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UnitRateOutput) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *UnitRateOutput) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *UnitRateOutput) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


