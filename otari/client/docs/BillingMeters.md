# BillingMeters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**map[string]ToolMeter**](ToolMeter.md) |  | [optional] 

## Methods

### NewBillingMeters

`func NewBillingMeters() *BillingMeters`

NewBillingMeters instantiates a new BillingMeters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingMetersWithDefaults

`func NewBillingMetersWithDefaults() *BillingMeters`

NewBillingMetersWithDefaults instantiates a new BillingMeters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *BillingMeters) GetTools() map[string]ToolMeter`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *BillingMeters) GetToolsOk() (*map[string]ToolMeter, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *BillingMeters) SetTools(v map[string]ToolMeter)`

SetTools sets Tools field to given value.

### HasTools

`func (o *BillingMeters) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


