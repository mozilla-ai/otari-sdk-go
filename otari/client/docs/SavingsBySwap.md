# SavingsBySwap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualCost** | **float32** | The same tokens priced at the models they ran on. | 
**BaselineCost** | **float32** | The subagents&#39; tokens priced at the models they would have started on. | 
**BaselineModel** | **string** | The priced model the baseline was read as, or the harness&#39;s name for it where it was read as none. | 
**Decisions** | **int32** | Recommendations answered. | 
**Estimated** | **int32** | Priced outcomes read from the harness&#39;s own token counts. | 
**GrossSaved** | **float32** | &#x60;baseline_cost&#x60; less &#x60;actual_cost&#x60;. | 
**Measured** | **int32** | Priced outcomes read from metered usage. | 
**NetSaved** | **float32** | &#x60;gross_saved&#x60; less &#x60;router_cost&#x60;. | 
**RanOn** | **string** | The model the subagent ran on, as the harness named it. | 
**RouterCost** | **float32** | What the recommendations themselves were charged. | 
**WithOutcome** | **int32** | Decisions whose subagent&#39;s consumption was reported. | 

## Methods

### NewSavingsBySwap

`func NewSavingsBySwap(actualCost float32, baselineCost float32, baselineModel string, decisions int32, estimated int32, grossSaved float32, measured int32, netSaved float32, ranOn string, routerCost float32, withOutcome int32, ) *SavingsBySwap`

NewSavingsBySwap instantiates a new SavingsBySwap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavingsBySwapWithDefaults

`func NewSavingsBySwapWithDefaults() *SavingsBySwap`

NewSavingsBySwapWithDefaults instantiates a new SavingsBySwap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualCost

`func (o *SavingsBySwap) GetActualCost() float32`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *SavingsBySwap) GetActualCostOk() (*float32, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *SavingsBySwap) SetActualCost(v float32)`

SetActualCost sets ActualCost field to given value.


### GetBaselineCost

`func (o *SavingsBySwap) GetBaselineCost() float32`

GetBaselineCost returns the BaselineCost field if non-nil, zero value otherwise.

### GetBaselineCostOk

`func (o *SavingsBySwap) GetBaselineCostOk() (*float32, bool)`

GetBaselineCostOk returns a tuple with the BaselineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCost

`func (o *SavingsBySwap) SetBaselineCost(v float32)`

SetBaselineCost sets BaselineCost field to given value.


### GetBaselineModel

`func (o *SavingsBySwap) GetBaselineModel() string`

GetBaselineModel returns the BaselineModel field if non-nil, zero value otherwise.

### GetBaselineModelOk

`func (o *SavingsBySwap) GetBaselineModelOk() (*string, bool)`

GetBaselineModelOk returns a tuple with the BaselineModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineModel

`func (o *SavingsBySwap) SetBaselineModel(v string)`

SetBaselineModel sets BaselineModel field to given value.


### GetDecisions

`func (o *SavingsBySwap) GetDecisions() int32`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *SavingsBySwap) GetDecisionsOk() (*int32, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *SavingsBySwap) SetDecisions(v int32)`

SetDecisions sets Decisions field to given value.


### GetEstimated

`func (o *SavingsBySwap) GetEstimated() int32`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *SavingsBySwap) GetEstimatedOk() (*int32, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *SavingsBySwap) SetEstimated(v int32)`

SetEstimated sets Estimated field to given value.


### GetGrossSaved

`func (o *SavingsBySwap) GetGrossSaved() float32`

GetGrossSaved returns the GrossSaved field if non-nil, zero value otherwise.

### GetGrossSavedOk

`func (o *SavingsBySwap) GetGrossSavedOk() (*float32, bool)`

GetGrossSavedOk returns a tuple with the GrossSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossSaved

`func (o *SavingsBySwap) SetGrossSaved(v float32)`

SetGrossSaved sets GrossSaved field to given value.


### GetMeasured

`func (o *SavingsBySwap) GetMeasured() int32`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *SavingsBySwap) GetMeasuredOk() (*int32, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *SavingsBySwap) SetMeasured(v int32)`

SetMeasured sets Measured field to given value.


### GetNetSaved

`func (o *SavingsBySwap) GetNetSaved() float32`

GetNetSaved returns the NetSaved field if non-nil, zero value otherwise.

### GetNetSavedOk

`func (o *SavingsBySwap) GetNetSavedOk() (*float32, bool)`

GetNetSavedOk returns a tuple with the NetSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSaved

`func (o *SavingsBySwap) SetNetSaved(v float32)`

SetNetSaved sets NetSaved field to given value.


### GetRanOn

`func (o *SavingsBySwap) GetRanOn() string`

GetRanOn returns the RanOn field if non-nil, zero value otherwise.

### GetRanOnOk

`func (o *SavingsBySwap) GetRanOnOk() (*string, bool)`

GetRanOnOk returns a tuple with the RanOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanOn

`func (o *SavingsBySwap) SetRanOn(v string)`

SetRanOn sets RanOn field to given value.


### GetRouterCost

`func (o *SavingsBySwap) GetRouterCost() float32`

GetRouterCost returns the RouterCost field if non-nil, zero value otherwise.

### GetRouterCostOk

`func (o *SavingsBySwap) GetRouterCostOk() (*float32, bool)`

GetRouterCostOk returns a tuple with the RouterCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCost

`func (o *SavingsBySwap) SetRouterCost(v float32)`

SetRouterCost sets RouterCost field to given value.


### GetWithOutcome

`func (o *SavingsBySwap) GetWithOutcome() int32`

GetWithOutcome returns the WithOutcome field if non-nil, zero value otherwise.

### GetWithOutcomeOk

`func (o *SavingsBySwap) GetWithOutcomeOk() (*int32, bool)`

GetWithOutcomeOk returns a tuple with the WithOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithOutcome

`func (o *SavingsBySwap) SetWithOutcome(v int32)`

SetWithOutcome sets WithOutcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


