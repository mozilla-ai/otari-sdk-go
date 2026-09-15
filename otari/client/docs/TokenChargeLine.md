# TokenChargeLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | USD charged for this line. | 
**Meter** | **string** | What was metered, e.g. &#39;input&#39; or &#39;cache_read&#39;. | 
**RatePerMillion** | **float32** | USD per million units. | 
**Units** | [**Units**](Units.md) |  | 

## Methods

### NewTokenChargeLine

`func NewTokenChargeLine(cost float32, meter string, ratePerMillion float32, units Units, ) *TokenChargeLine`

NewTokenChargeLine instantiates a new TokenChargeLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenChargeLineWithDefaults

`func NewTokenChargeLineWithDefaults() *TokenChargeLine`

NewTokenChargeLineWithDefaults instantiates a new TokenChargeLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *TokenChargeLine) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TokenChargeLine) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TokenChargeLine) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetMeter

`func (o *TokenChargeLine) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *TokenChargeLine) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *TokenChargeLine) SetMeter(v string)`

SetMeter sets Meter field to given value.


### GetRatePerMillion

`func (o *TokenChargeLine) GetRatePerMillion() float32`

GetRatePerMillion returns the RatePerMillion field if non-nil, zero value otherwise.

### GetRatePerMillionOk

`func (o *TokenChargeLine) GetRatePerMillionOk() (*float32, bool)`

GetRatePerMillionOk returns a tuple with the RatePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerMillion

`func (o *TokenChargeLine) SetRatePerMillion(v float32)`

SetRatePerMillion sets RatePerMillion field to given value.


### GetUnits

`func (o *TokenChargeLine) GetUnits() Units`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TokenChargeLine) GetUnitsOk() (*Units, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TokenChargeLine) SetUnits(v Units)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


