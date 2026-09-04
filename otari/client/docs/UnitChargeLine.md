# UnitChargeLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | USD charged for this line. | 
**Meter** | **string** | What was metered, e.g. &#39;request&#39; or &#39;web_search_calls&#39;. | 
**UnitRate** | **float32** | USD per call. | 
**Units** | [**Units1**](Units1.md) |  | 

## Methods

### NewUnitChargeLine

`func NewUnitChargeLine(cost float32, meter string, unitRate float32, units Units1, ) *UnitChargeLine`

NewUnitChargeLine instantiates a new UnitChargeLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitChargeLineWithDefaults

`func NewUnitChargeLineWithDefaults() *UnitChargeLine`

NewUnitChargeLineWithDefaults instantiates a new UnitChargeLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *UnitChargeLine) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UnitChargeLine) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UnitChargeLine) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetMeter

`func (o *UnitChargeLine) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *UnitChargeLine) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *UnitChargeLine) SetMeter(v string)`

SetMeter sets Meter field to given value.


### GetUnitRate

`func (o *UnitChargeLine) GetUnitRate() float32`

GetUnitRate returns the UnitRate field if non-nil, zero value otherwise.

### GetUnitRateOk

`func (o *UnitChargeLine) GetUnitRateOk() (*float32, bool)`

GetUnitRateOk returns a tuple with the UnitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitRate

`func (o *UnitChargeLine) SetUnitRate(v float32)`

SetUnitRate sets UnitRate field to given value.


### GetUnits

`func (o *UnitChargeLine) GetUnits() Units1`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *UnitChargeLine) GetUnitsOk() (*Units1, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *UnitChargeLine) SetUnits(v Units1)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


