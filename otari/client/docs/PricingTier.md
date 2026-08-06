# PricingTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWrite1hPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**InputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**MinInputTokens** | **int32** |  | 
**OutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewPricingTier

`func NewPricingTier(minInputTokens int32, ) *PricingTier`

NewPricingTier instantiates a new PricingTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingTierWithDefaults

`func NewPricingTierWithDefaults() *PricingTier`

NewPricingTierWithDefaults instantiates a new PricingTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *PricingTier) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *PricingTier) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *PricingTier) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *PricingTier) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *PricingTier) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *PricingTier) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWrite1hPricePerMillion

`func (o *PricingTier) GetCacheWrite1hPricePerMillion() float32`

GetCacheWrite1hPricePerMillion returns the CacheWrite1hPricePerMillion field if non-nil, zero value otherwise.

### GetCacheWrite1hPricePerMillionOk

`func (o *PricingTier) GetCacheWrite1hPricePerMillionOk() (*float32, bool)`

GetCacheWrite1hPricePerMillionOk returns a tuple with the CacheWrite1hPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hPricePerMillion

`func (o *PricingTier) SetCacheWrite1hPricePerMillion(v float32)`

SetCacheWrite1hPricePerMillion sets CacheWrite1hPricePerMillion field to given value.

### HasCacheWrite1hPricePerMillion

`func (o *PricingTier) HasCacheWrite1hPricePerMillion() bool`

HasCacheWrite1hPricePerMillion returns a boolean if a field has been set.

### SetCacheWrite1hPricePerMillionNil

`func (o *PricingTier) SetCacheWrite1hPricePerMillionNil(b bool)`

 SetCacheWrite1hPricePerMillionNil sets the value for CacheWrite1hPricePerMillion to be an explicit nil

### UnsetCacheWrite1hPricePerMillion
`func (o *PricingTier) UnsetCacheWrite1hPricePerMillion()`

UnsetCacheWrite1hPricePerMillion ensures that no value is present for CacheWrite1hPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *PricingTier) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *PricingTier) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *PricingTier) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *PricingTier) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *PricingTier) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *PricingTier) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetInputPricePerMillion

`func (o *PricingTier) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *PricingTier) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *PricingTier) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *PricingTier) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### SetInputPricePerMillionNil

`func (o *PricingTier) SetInputPricePerMillionNil(b bool)`

 SetInputPricePerMillionNil sets the value for InputPricePerMillion to be an explicit nil

### UnsetInputPricePerMillion
`func (o *PricingTier) UnsetInputPricePerMillion()`

UnsetInputPricePerMillion ensures that no value is present for InputPricePerMillion, not even an explicit nil
### GetMinInputTokens

`func (o *PricingTier) GetMinInputTokens() int32`

GetMinInputTokens returns the MinInputTokens field if non-nil, zero value otherwise.

### GetMinInputTokensOk

`func (o *PricingTier) GetMinInputTokensOk() (*int32, bool)`

GetMinInputTokensOk returns a tuple with the MinInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInputTokens

`func (o *PricingTier) SetMinInputTokens(v int32)`

SetMinInputTokens sets MinInputTokens field to given value.


### GetOutputPricePerMillion

`func (o *PricingTier) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *PricingTier) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *PricingTier) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *PricingTier) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### SetOutputPricePerMillionNil

`func (o *PricingTier) SetOutputPricePerMillionNil(b bool)`

 SetOutputPricePerMillionNil sets the value for OutputPricePerMillion to be an explicit nil

### UnsetOutputPricePerMillion
`func (o *PricingTier) UnsetOutputPricePerMillion()`

UnsetOutputPricePerMillion ensures that no value is present for OutputPricePerMillion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


