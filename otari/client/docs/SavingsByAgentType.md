# SavingsByAgentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualCost** | **float32** | The same tokens priced at the models they ran on. | 
**AgentType** | **string** |  | 
**BaselineCost** | **float32** | The subagents&#39; tokens priced at the models they would have started on. | 
**Decisions** | **int32** | Recommendations answered. | 
**Estimated** | **int32** | Priced outcomes read from the harness&#39;s own token counts. | 
**GrossSaved** | **float32** | &#x60;baseline_cost&#x60; less &#x60;actual_cost&#x60;. | 
**Measured** | **int32** | Priced outcomes read from metered usage. | 
**NetSaved** | **float32** | &#x60;gross_saved&#x60; less &#x60;router_cost&#x60;. | 
**RouterCost** | **float32** | What the recommendations themselves were charged. | 
**WithOutcome** | **int32** | Decisions whose subagent&#39;s consumption was reported. | 

## Methods

### NewSavingsByAgentType

`func NewSavingsByAgentType(actualCost float32, agentType string, baselineCost float32, decisions int32, estimated int32, grossSaved float32, measured int32, netSaved float32, routerCost float32, withOutcome int32, ) *SavingsByAgentType`

NewSavingsByAgentType instantiates a new SavingsByAgentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavingsByAgentTypeWithDefaults

`func NewSavingsByAgentTypeWithDefaults() *SavingsByAgentType`

NewSavingsByAgentTypeWithDefaults instantiates a new SavingsByAgentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualCost

`func (o *SavingsByAgentType) GetActualCost() float32`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *SavingsByAgentType) GetActualCostOk() (*float32, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *SavingsByAgentType) SetActualCost(v float32)`

SetActualCost sets ActualCost field to given value.


### GetAgentType

`func (o *SavingsByAgentType) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *SavingsByAgentType) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *SavingsByAgentType) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetBaselineCost

`func (o *SavingsByAgentType) GetBaselineCost() float32`

GetBaselineCost returns the BaselineCost field if non-nil, zero value otherwise.

### GetBaselineCostOk

`func (o *SavingsByAgentType) GetBaselineCostOk() (*float32, bool)`

GetBaselineCostOk returns a tuple with the BaselineCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCost

`func (o *SavingsByAgentType) SetBaselineCost(v float32)`

SetBaselineCost sets BaselineCost field to given value.


### GetDecisions

`func (o *SavingsByAgentType) GetDecisions() int32`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *SavingsByAgentType) GetDecisionsOk() (*int32, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *SavingsByAgentType) SetDecisions(v int32)`

SetDecisions sets Decisions field to given value.


### GetEstimated

`func (o *SavingsByAgentType) GetEstimated() int32`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *SavingsByAgentType) GetEstimatedOk() (*int32, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *SavingsByAgentType) SetEstimated(v int32)`

SetEstimated sets Estimated field to given value.


### GetGrossSaved

`func (o *SavingsByAgentType) GetGrossSaved() float32`

GetGrossSaved returns the GrossSaved field if non-nil, zero value otherwise.

### GetGrossSavedOk

`func (o *SavingsByAgentType) GetGrossSavedOk() (*float32, bool)`

GetGrossSavedOk returns a tuple with the GrossSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossSaved

`func (o *SavingsByAgentType) SetGrossSaved(v float32)`

SetGrossSaved sets GrossSaved field to given value.


### GetMeasured

`func (o *SavingsByAgentType) GetMeasured() int32`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *SavingsByAgentType) GetMeasuredOk() (*int32, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *SavingsByAgentType) SetMeasured(v int32)`

SetMeasured sets Measured field to given value.


### GetNetSaved

`func (o *SavingsByAgentType) GetNetSaved() float32`

GetNetSaved returns the NetSaved field if non-nil, zero value otherwise.

### GetNetSavedOk

`func (o *SavingsByAgentType) GetNetSavedOk() (*float32, bool)`

GetNetSavedOk returns a tuple with the NetSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSaved

`func (o *SavingsByAgentType) SetNetSaved(v float32)`

SetNetSaved sets NetSaved field to given value.


### GetRouterCost

`func (o *SavingsByAgentType) GetRouterCost() float32`

GetRouterCost returns the RouterCost field if non-nil, zero value otherwise.

### GetRouterCostOk

`func (o *SavingsByAgentType) GetRouterCostOk() (*float32, bool)`

GetRouterCostOk returns a tuple with the RouterCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCost

`func (o *SavingsByAgentType) SetRouterCost(v float32)`

SetRouterCost sets RouterCost field to given value.


### GetWithOutcome

`func (o *SavingsByAgentType) GetWithOutcome() int32`

GetWithOutcome returns the WithOutcome field if non-nil, zero value otherwise.

### GetWithOutcomeOk

`func (o *SavingsByAgentType) GetWithOutcomeOk() (*int32, bool)`

GetWithOutcomeOk returns a tuple with the WithOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithOutcome

`func (o *SavingsByAgentType) SetWithOutcome(v int32)`

SetWithOutcome sets WithOutcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


