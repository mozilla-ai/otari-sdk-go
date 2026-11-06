# RerankUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostUsd** | Pointer to **NullableString** | What the request cost, as a six-decimal USD string; absent when it was not priced | [optional] 
**PricingSource** | Pointer to **NullableString** | Which rate priced the request: &#x60;organization&#x60;, &#x60;deployment&#x60; or &#x60;defaults&#x60;; absent when it was not priced | [optional] 
**SearchUnits** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**TotalTokens** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 

## Methods

### NewRerankUsage

`func NewRerankUsage() *RerankUsage`

NewRerankUsage instantiates a new RerankUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankUsageWithDefaults

`func NewRerankUsageWithDefaults() *RerankUsage`

NewRerankUsageWithDefaults instantiates a new RerankUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostUsd

`func (o *RerankUsage) GetCostUsd() string`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *RerankUsage) GetCostUsdOk() (*string, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *RerankUsage) SetCostUsd(v string)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *RerankUsage) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### SetCostUsdNil

`func (o *RerankUsage) SetCostUsdNil(b bool)`

 SetCostUsdNil sets the value for CostUsd to be an explicit nil

### UnsetCostUsd
`func (o *RerankUsage) UnsetCostUsd()`

UnsetCostUsd ensures that no value is present for CostUsd, not even an explicit nil
### GetPricingSource

`func (o *RerankUsage) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *RerankUsage) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *RerankUsage) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.

### HasPricingSource

`func (o *RerankUsage) HasPricingSource() bool`

HasPricingSource returns a boolean if a field has been set.

### SetPricingSourceNil

`func (o *RerankUsage) SetPricingSourceNil(b bool)`

 SetPricingSourceNil sets the value for PricingSource to be an explicit nil

### UnsetPricingSource
`func (o *RerankUsage) UnsetPricingSource()`

UnsetPricingSource ensures that no value is present for PricingSource, not even an explicit nil
### GetSearchUnits

`func (o *RerankUsage) GetSearchUnits() int32`

GetSearchUnits returns the SearchUnits field if non-nil, zero value otherwise.

### GetSearchUnitsOk

`func (o *RerankUsage) GetSearchUnitsOk() (*int32, bool)`

GetSearchUnitsOk returns a tuple with the SearchUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchUnits

`func (o *RerankUsage) SetSearchUnits(v int32)`

SetSearchUnits sets SearchUnits field to given value.

### HasSearchUnits

`func (o *RerankUsage) HasSearchUnits() bool`

HasSearchUnits returns a boolean if a field has been set.

### SetSearchUnitsNil

`func (o *RerankUsage) SetSearchUnitsNil(b bool)`

 SetSearchUnitsNil sets the value for SearchUnits to be an explicit nil

### UnsetSearchUnits
`func (o *RerankUsage) UnsetSearchUnits()`

UnsetSearchUnits ensures that no value is present for SearchUnits, not even an explicit nil
### GetTotalTokens

`func (o *RerankUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *RerankUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *RerankUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *RerankUsage) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### SetTotalTokensNil

`func (o *RerankUsage) SetTotalTokensNil(b bool)`

 SetTotalTokensNil sets the value for TotalTokens to be an explicit nil

### UnsetTotalTokens
`func (o *RerankUsage) UnsetTotalTokens()`

UnsetTotalTokens ensures that no value is present for TotalTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


