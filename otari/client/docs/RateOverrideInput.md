# RateOverrideInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**NullableCacheRatesInput**](CacheRatesInput.md) |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**Reasoning** | Pointer to **NullableFloat32** |  | [optional] 
**When** | **map[string]interface{}** | The condition, as models.dev writes it. | 

## Methods

### NewRateOverrideInput

`func NewRateOverrideInput(when map[string]interface{}, ) *RateOverrideInput`

NewRateOverrideInput instantiates a new RateOverrideInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateOverrideInputWithDefaults

`func NewRateOverrideInputWithDefaults() *RateOverrideInput`

NewRateOverrideInputWithDefaults instantiates a new RateOverrideInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *RateOverrideInput) GetCache() CacheRatesInput`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *RateOverrideInput) GetCacheOk() (*CacheRatesInput, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *RateOverrideInput) SetCache(v CacheRatesInput)`

SetCache sets Cache field to given value.

### HasCache

`func (o *RateOverrideInput) HasCache() bool`

HasCache returns a boolean if a field has been set.

### SetCacheNil

`func (o *RateOverrideInput) SetCacheNil(b bool)`

 SetCacheNil sets the value for Cache to be an explicit nil

### UnsetCache
`func (o *RateOverrideInput) UnsetCache()`

UnsetCache ensures that no value is present for Cache, not even an explicit nil
### GetPrice

`func (o *RateOverrideInput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RateOverrideInput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RateOverrideInput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *RateOverrideInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *RateOverrideInput) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *RateOverrideInput) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetReasoning

`func (o *RateOverrideInput) GetReasoning() float32`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *RateOverrideInput) GetReasoningOk() (*float32, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *RateOverrideInput) SetReasoning(v float32)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *RateOverrideInput) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *RateOverrideInput) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *RateOverrideInput) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetWhen

`func (o *RateOverrideInput) GetWhen() map[string]interface{}`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *RateOverrideInput) GetWhenOk() (*map[string]interface{}, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *RateOverrideInput) SetWhen(v map[string]interface{})`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


