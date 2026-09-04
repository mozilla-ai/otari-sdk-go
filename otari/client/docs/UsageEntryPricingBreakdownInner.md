# UsageEntryPricingBreakdownInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | USD charged for this line. | 
**Meter** | **string** | What was metered, e.g. &#39;request&#39; or &#39;web_search_calls&#39;. | 
**RatePerMillion** | **float32** | USD per million units. | 
**Units** | [**Units1**](Units1.md) |  | 
**UnitRate** | **float32** | USD per call. | 

## Methods

### NewUsageEntryPricingBreakdownInner

`func NewUsageEntryPricingBreakdownInner(cost float32, meter string, ratePerMillion float32, units Units1, unitRate float32, ) *UsageEntryPricingBreakdownInner`

NewUsageEntryPricingBreakdownInner instantiates a new UsageEntryPricingBreakdownInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEntryPricingBreakdownInnerWithDefaults

`func NewUsageEntryPricingBreakdownInnerWithDefaults() *UsageEntryPricingBreakdownInner`

NewUsageEntryPricingBreakdownInnerWithDefaults instantiates a new UsageEntryPricingBreakdownInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *UsageEntryPricingBreakdownInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageEntryPricingBreakdownInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageEntryPricingBreakdownInner) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetMeter

`func (o *UsageEntryPricingBreakdownInner) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *UsageEntryPricingBreakdownInner) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *UsageEntryPricingBreakdownInner) SetMeter(v string)`

SetMeter sets Meter field to given value.


### GetRatePerMillion

`func (o *UsageEntryPricingBreakdownInner) GetRatePerMillion() float32`

GetRatePerMillion returns the RatePerMillion field if non-nil, zero value otherwise.

### GetRatePerMillionOk

`func (o *UsageEntryPricingBreakdownInner) GetRatePerMillionOk() (*float32, bool)`

GetRatePerMillionOk returns a tuple with the RatePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerMillion

`func (o *UsageEntryPricingBreakdownInner) SetRatePerMillion(v float32)`

SetRatePerMillion sets RatePerMillion field to given value.


### GetUnits

`func (o *UsageEntryPricingBreakdownInner) GetUnits() Units1`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *UsageEntryPricingBreakdownInner) GetUnitsOk() (*Units1, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *UsageEntryPricingBreakdownInner) SetUnits(v Units1)`

SetUnits sets Units field to given value.


### GetUnitRate

`func (o *UsageEntryPricingBreakdownInner) GetUnitRate() float32`

GetUnitRate returns the UnitRate field if non-nil, zero value otherwise.

### GetUnitRateOk

`func (o *UsageEntryPricingBreakdownInner) GetUnitRateOk() (*float32, bool)`

GetUnitRateOk returns a tuple with the UnitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitRate

`func (o *UsageEntryPricingBreakdownInner) SetUnitRate(v float32)`

SetUnitRate sets UnitRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


