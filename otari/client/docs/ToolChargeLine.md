# ToolChargeLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | USD charged for this line. | 
**Label** | **string** | What served the calls: &#39;&lt;provider&gt;:&lt;instance&gt;&#39; for a search or fetch instance, &#39;builtin:builtin_fetch&#39; for the built-in fetcher, &#39;org:&lt;provider&gt;&#39; for an organization&#39;s own key, or &#39;otari:&lt;tool&gt;&#39; for any other gateway-run tool. | 
**Meter** | **string** | What was metered, e.g. &#39;request&#39; or &#39;web_search_calls&#39;. | 
**UnitRate** | **float32** | USD per call. | 
**Units** | [**Units1**](Units1.md) |  | 

## Methods

### NewToolChargeLine

`func NewToolChargeLine(cost float32, label string, meter string, unitRate float32, units Units1, ) *ToolChargeLine`

NewToolChargeLine instantiates a new ToolChargeLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolChargeLineWithDefaults

`func NewToolChargeLineWithDefaults() *ToolChargeLine`

NewToolChargeLineWithDefaults instantiates a new ToolChargeLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *ToolChargeLine) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ToolChargeLine) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ToolChargeLine) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetLabel

`func (o *ToolChargeLine) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ToolChargeLine) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ToolChargeLine) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMeter

`func (o *ToolChargeLine) GetMeter() string`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *ToolChargeLine) GetMeterOk() (*string, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *ToolChargeLine) SetMeter(v string)`

SetMeter sets Meter field to given value.


### GetUnitRate

`func (o *ToolChargeLine) GetUnitRate() float32`

GetUnitRate returns the UnitRate field if non-nil, zero value otherwise.

### GetUnitRateOk

`func (o *ToolChargeLine) GetUnitRateOk() (*float32, bool)`

GetUnitRateOk returns a tuple with the UnitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitRate

`func (o *ToolChargeLine) SetUnitRate(v float32)`

SetUnitRate sets UnitRate field to given value.


### GetUnits

`func (o *ToolChargeLine) GetUnits() Units1`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ToolChargeLine) GetUnitsOk() (*Units1, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ToolChargeLine) SetUnits(v Units1)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


