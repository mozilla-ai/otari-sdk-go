# SubagentSavingsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByAgentType** | [**[]SavingsByAgentType**](SavingsByAgentType.md) | Largest saving first, at most 20. | 
**BySwap** | [**[]SavingsBySwap**](SavingsBySwap.md) | Largest saving first, at most 20. | 
**Estimated** | [**SavingsAmounts**](SavingsAmounts.md) | The decisions whose outcome is the harness&#39;s own count. | 
**Measured** | [**SavingsAmounts**](SavingsAmounts.md) | The decisions whose outcome is metered. | 
**RecommenderAvailable** | **bool** | Whether anything answers a recommendation on this deployment. | 
**Series** | [**[]SubagentSavingsPoint**](SubagentSavingsPoint.md) |  | 
**Totals** | [**SubagentSavingsTotals**](SubagentSavingsTotals.md) |  | 

## Methods

### NewSubagentSavingsPublic

`func NewSubagentSavingsPublic(byAgentType []SavingsByAgentType, bySwap []SavingsBySwap, estimated SavingsAmounts, measured SavingsAmounts, recommenderAvailable bool, series []SubagentSavingsPoint, totals SubagentSavingsTotals, ) *SubagentSavingsPublic`

NewSubagentSavingsPublic instantiates a new SubagentSavingsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubagentSavingsPublicWithDefaults

`func NewSubagentSavingsPublicWithDefaults() *SubagentSavingsPublic`

NewSubagentSavingsPublicWithDefaults instantiates a new SubagentSavingsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByAgentType

`func (o *SubagentSavingsPublic) GetByAgentType() []SavingsByAgentType`

GetByAgentType returns the ByAgentType field if non-nil, zero value otherwise.

### GetByAgentTypeOk

`func (o *SubagentSavingsPublic) GetByAgentTypeOk() (*[]SavingsByAgentType, bool)`

GetByAgentTypeOk returns a tuple with the ByAgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAgentType

`func (o *SubagentSavingsPublic) SetByAgentType(v []SavingsByAgentType)`

SetByAgentType sets ByAgentType field to given value.


### GetBySwap

`func (o *SubagentSavingsPublic) GetBySwap() []SavingsBySwap`

GetBySwap returns the BySwap field if non-nil, zero value otherwise.

### GetBySwapOk

`func (o *SubagentSavingsPublic) GetBySwapOk() (*[]SavingsBySwap, bool)`

GetBySwapOk returns a tuple with the BySwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySwap

`func (o *SubagentSavingsPublic) SetBySwap(v []SavingsBySwap)`

SetBySwap sets BySwap field to given value.


### GetEstimated

`func (o *SubagentSavingsPublic) GetEstimated() SavingsAmounts`

GetEstimated returns the Estimated field if non-nil, zero value otherwise.

### GetEstimatedOk

`func (o *SubagentSavingsPublic) GetEstimatedOk() (*SavingsAmounts, bool)`

GetEstimatedOk returns a tuple with the Estimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimated

`func (o *SubagentSavingsPublic) SetEstimated(v SavingsAmounts)`

SetEstimated sets Estimated field to given value.


### GetMeasured

`func (o *SubagentSavingsPublic) GetMeasured() SavingsAmounts`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *SubagentSavingsPublic) GetMeasuredOk() (*SavingsAmounts, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *SubagentSavingsPublic) SetMeasured(v SavingsAmounts)`

SetMeasured sets Measured field to given value.


### GetRecommenderAvailable

`func (o *SubagentSavingsPublic) GetRecommenderAvailable() bool`

GetRecommenderAvailable returns the RecommenderAvailable field if non-nil, zero value otherwise.

### GetRecommenderAvailableOk

`func (o *SubagentSavingsPublic) GetRecommenderAvailableOk() (*bool, bool)`

GetRecommenderAvailableOk returns a tuple with the RecommenderAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommenderAvailable

`func (o *SubagentSavingsPublic) SetRecommenderAvailable(v bool)`

SetRecommenderAvailable sets RecommenderAvailable field to given value.


### GetSeries

`func (o *SubagentSavingsPublic) GetSeries() []SubagentSavingsPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SubagentSavingsPublic) GetSeriesOk() (*[]SubagentSavingsPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SubagentSavingsPublic) SetSeries(v []SubagentSavingsPoint)`

SetSeries sets Series field to given value.


### GetTotals

`func (o *SubagentSavingsPublic) GetTotals() SubagentSavingsTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *SubagentSavingsPublic) GetTotalsOk() (*SubagentSavingsTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *SubagentSavingsPublic) SetTotals(v SubagentSavingsTotals)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


