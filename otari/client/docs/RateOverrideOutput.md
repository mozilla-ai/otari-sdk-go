# RateOverrideOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**NullableCacheRatesOutput**](CacheRatesOutput.md) |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**Reasoning** | Pointer to **NullableFloat32** |  | [optional] 
**When** | **map[string]interface{}** | The condition, as models.dev writes it. | 

## Methods

### NewRateOverrideOutput

`func NewRateOverrideOutput(when map[string]interface{}, ) *RateOverrideOutput`

NewRateOverrideOutput instantiates a new RateOverrideOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateOverrideOutputWithDefaults

`func NewRateOverrideOutputWithDefaults() *RateOverrideOutput`

NewRateOverrideOutputWithDefaults instantiates a new RateOverrideOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *RateOverrideOutput) GetCache() CacheRatesOutput`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *RateOverrideOutput) GetCacheOk() (*CacheRatesOutput, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *RateOverrideOutput) SetCache(v CacheRatesOutput)`

SetCache sets Cache field to given value.

### HasCache

`func (o *RateOverrideOutput) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *RateOverrideOutput) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *RateOverrideOutput) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetPrice

`func (o *RateOverrideOutput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RateOverrideOutput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RateOverrideOutput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RateOverrideOutput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *RateOverrideOutput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *RateOverrideOutput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetReasoning

`func (o *RateOverrideOutput) GetReasoning() float32`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *RateOverrideOutput) GetReasoningOk() (*float32, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *RateOverrideOutput) SetReasoning(v float32)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *RateOverrideOutput) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *RateOverrideOutput) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *RateOverrideOutput) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetWhen

`func (o *RateOverrideOutput) GetWhen() map[string]interface{}`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *RateOverrideOutput) GetWhenOk() (*map[string]interface{}, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *RateOverrideOutput) SetWhen(v map[string]interface{})`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


