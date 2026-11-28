# SubagentSavingsTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualCost** | **float32** | The same tokens priced at the models they ran on. | 
**BaselineCost** | **float32** | The subagents&#39; tokens priced at the models they would have started on. | 
**Decisions** | **int32** | Recommendations answered. | 
**Estimated** | **int32** | Priced outcomes read from the harness&#39;s own token counts. | 
**GrossSaved** | **float32** | &#x60;baseline_cost&#x60; less &#x60;actual_cost&#x60;. | 
**Measured** | **int32** | Priced outcomes read from metered usage. | 
**NetSaved** | **float32** | &#x60;gross_saved&#x60; less &#x60;router_cost&#x60;. | 
**NoOutcome** | **int32** | Decisions with no consumption reported. Their router cost still counts. | 
**RouterCost** | **float32** | What the recommendations themselves were charged. | 
**RouterUnanswered** | **int32** | Recommendation requests the router failed to answer. | 
**Unpriced** | **int32** | Reported outcomes left out of the costs because a side has no price. | 
**WithOutcome** | **int32** | Decisions whose subagent&#39;s consumption was reported. | 

## Methods

### NewSubagentSavingsTotals

`func NewSubagentSavingsTotals(actualCost float32, baselineCost float32, decisions int32, estimated int32, grossSaved float32, measured int32, netSaved float32, noOutcome int32, routerCost float32, routerUnanswered int32, unpriced int32, withOutcome int32, ) *SubagentSavingsTotals`

NewSubagentSavingsTotals instantiates a new SubagentSavingsTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubagentSavingsTotalsWithDefaults

`func NewSubagentSavingsTotalsWithDefaults() *SubagentSavingsTotals`

NewSubagentSavingsTotalsWithDefaults instantiates a new SubagentSavingsTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualCost

`func (o *SubagentSavingsTotals) GetActualCost() float32`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *SubagentSavingsTotals) GetActualCostOk() (*float32, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *SubagentSavingsTotals) SetActualCost(v float32)`

SetActualCost sets ActualCost field to given value.


### GetBaselineCost

`func (o *SubagentSavingsTotals) GetBaselineCost() float32`

GetBaselineCost returns the BaselineCost field if non-nil, zero value otherwise.

### GetBaselineCostOk

`func (o *SubagentSavingsTotals) GetBaselineCostOk() (*float32, bool)`

GetBaselineCostOk returns a tuple with the BaselineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCost

`func (o *SubagentSavingsTotals) SetBaselineCost(v float32)`

SetBaselineCost sets BaselineCost field to given value.


### GetDecisions

`func (o *SubagentSavingsTotals) GetDecisions() int32`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *SubagentSavingsTotals) GetDecisionsOk() (*int32, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *SubagentSavingsTotals) SetDecisions(v int32)`

SetDecisions sets Decisions field to given value.


### GetEstimated

`func (o *SubagentSavingsTotals) GetEstimated() int32`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *SubagentSavingsTotals) GetEstimatedOk() (*int32, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *SubagentSavingsTotals) SetEstimated(v int32)`

SetEstimated sets Estimated field to given value.


### GetGrossSaved

`func (o *SubagentSavingsTotals) GetGrossSaved() float32`

GetGrossSaved returns the GrossSaved field if non-nil, zero value otherwise.

### GetGrossSavedOk

`func (o *SubagentSavingsTotals) GetGrossSavedOk() (*float32, bool)`

GetGrossSavedOk returns a tuple with the GrossSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossSaved

`func (o *SubagentSavingsTotals) SetGrossSaved(v float32)`

SetGrossSaved sets GrossSaved field to given value.


### GetMeasured

`func (o *SubagentSavingsTotals) GetMeasured() int32`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *SubagentSavingsTotals) GetMeasuredOk() (*int32, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *SubagentSavingsTotals) SetMeasured(v int32)`

SetMeasured sets Measured field to given value.


### GetNetSaved

`func (o *SubagentSavingsTotals) GetNetSaved() float32`

GetNetSaved returns the NetSaved field if non-nil, zero value otherwise.

### GetNetSavedOk

`func (o *SubagentSavingsTotals) GetNetSavedOk() (*float32, bool)`

GetNetSavedOk returns a tuple with the NetSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSaved

`func (o *SubagentSavingsTotals) SetNetSaved(v float32)`

SetNetSaved sets NetSaved field to given value.


### GetNoOutcome

`func (o *SubagentSavingsTotals) GetNoOutcome() int32`

GetNoOutcome returns the NoOutcome field if non-nil, zero value otherwise.

### GetNoOutcomeOk

`func (o *SubagentSavingsTotals) GetNoOutcomeOk() (*int32, bool)`

GetNoOutcomeOk returns a tuple with the NoOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOutcome

`func (o *SubagentSavingsTotals) SetNoOutcome(v int32)`

SetNoOutcome sets NoOutcome field to given value.


### GetRouterCost

`func (o *SubagentSavingsTotals) GetRouterCost() float32`

GetRouterCost returns the RouterCost field if non-nil, zero value otherwise.

### GetRouterCostOk

`func (o *SubagentSavingsTotals) GetRouterCostOk() (*float32, bool)`

GetRouterCostOk returns a tuple with the RouterCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCost

`func (o *SubagentSavingsTotals) SetRouterCost(v float32)`

SetRouterCost sets RouterCost field to given value.


### GetRouterUnanswered

`func (o *SubagentSavingsTotals) GetRouterUnanswered() int32`

GetRouterUnanswered returns the RouterUnanswered field if non-nil, zero value otherwise.

### GetRouterUnansweredOk

`func (o *SubagentSavingsTotals) GetRouterUnansweredOk() (*int32, bool)`

GetRouterUnansweredOk returns a tuple with the RouterUnanswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterUnanswered

`func (o *SubagentSavingsTotals) SetRouterUnanswered(v int32)`

SetRouterUnanswered sets RouterUnanswered field to given value.


### GetUnpriced

`func (o *SubagentSavingsTotals) GetUnpriced() int32`

GetUnpriced returns the Unpriced field if non-nil, zero value otherwise.

### GetUnpricedOk

`func (o *SubagentSavingsTotals) GetUnpricedOk() (*int32, bool)`

GetUnpricedOk returns a tuple with the Unpriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpriced

`func (o *SubagentSavingsTotals) SetUnpriced(v int32)`

SetUnpriced sets Unpriced field to given value.


### GetWithOutcome

`func (o *SubagentSavingsTotals) GetWithOutcome() int32`

GetWithOutcome returns the WithOutcome field if non-nil, zero value otherwise.

### GetWithOutcomeOk

`func (o *SubagentSavingsTotals) GetWithOutcomeOk() (*int32, bool)`

GetWithOutcomeOk returns a tuple with the WithOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithOutcome

`func (o *SubagentSavingsTotals) SetWithOutcome(v int32)`

SetWithOutcome sets WithOutcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


