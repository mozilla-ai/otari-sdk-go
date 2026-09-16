# OfferingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheHitRate** | **NullableFloat32** | Cache-read tokens over prompt tokens. Null when no prompt tokens. | 
**CacheReadTokens** | **int32** |  | 
**EffectivePricePerMillion** | **NullableFloat32** | Spend over every token served, per million. Null when no tokens were served. | 
**Requests** | **int32** |  | 
**SpendUsd** | **float32** |  | 
**TotalTokens** | **int32** |  | 

## Methods

### NewOfferingUsage

`func NewOfferingUsage(cacheHitRate NullableFloat32, cacheReadTokens int32, effectivePricePerMillion NullableFloat32, requests int32, spendUsd float32, totalTokens int32, ) *OfferingUsage`

NewOfferingUsage instantiates a new OfferingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferingUsageWithDefaults

`func NewOfferingUsageWithDefaults() *OfferingUsage`

NewOfferingUsageWithDefaults instantiates a new OfferingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheHitRate

`func (o *OfferingUsage) GetCacheHitRate() float32`

GetCacheHitRate returns the CacheHitRate field if non-nil, zero value otherwise.

### GetCacheHitRateOk

`func (o *OfferingUsage) GetCacheHitRateOk() (*float32, bool)`

GetCacheHitRateOk returns a tuple with the CacheHitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHitRate

`func (o *OfferingUsage) SetCacheHitRate(v float32)`

SetCacheHitRate sets CacheHitRate field to given value.


### SetCacheHitRateNil

`func (o *OfferingUsage) SetCacheHitRateNil(b bool)`

 SetCacheHitRateNil sets the value for CacheHitRate to be an explicit nil

### UnsetCacheHitRate
`func (o *OfferingUsage) UnsetCacheHitRate()`

UnsetCacheHitRate ensures that no value is present for CacheHitRate, not even an explicit nil
### GetCacheReadTokens

`func (o *OfferingUsage) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *OfferingUsage) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *OfferingUsage) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### GetEffectivePricePerMillion

`func (o *OfferingUsage) GetEffectivePricePerMillion() float32`

GetEffectivePricePerMillion returns the EffectivePricePerMillion field if non-nil, zero value otherwise.

### GetEffectivePricePerMillionOk

`func (o *OfferingUsage) GetEffectivePricePerMillionOk() (*float32, bool)`

GetEffectivePricePerMillionOk returns a tuple with the EffectivePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePricePerMillion

`func (o *OfferingUsage) SetEffectivePricePerMillion(v float32)`

SetEffectivePricePerMillion sets EffectivePricePerMillion field to given value.


### SetEffectivePricePerMillionNil

`func (o *OfferingUsage) SetEffectivePricePerMillionNil(b bool)`

 SetEffectivePricePerMillionNil sets the value for EffectivePricePerMillion to be an explicit nil

### UnsetEffectivePricePerMillion
`func (o *OfferingUsage) UnsetEffectivePricePerMillion()`

UnsetEffectivePricePerMillion ensures that no value is present for EffectivePricePerMillion, not even an explicit nil
### GetRequests

`func (o *OfferingUsage) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *OfferingUsage) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *OfferingUsage) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetSpendUsd

`func (o *OfferingUsage) GetSpendUsd() float32`

GetSpendUsd returns the SpendUsd field if non-nil, zero value otherwise.

### GetSpendUsdOk

`func (o *OfferingUsage) GetSpendUsdOk() (*float32, bool)`

GetSpendUsdOk returns a tuple with the SpendUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendUsd

`func (o *OfferingUsage) SetSpendUsd(v float32)`

SetSpendUsd sets SpendUsd field to given value.


### GetTotalTokens

`func (o *OfferingUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *OfferingUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *OfferingUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


