# OverviewSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveKeys** | **int32** |  | 
**ActiveMembers** | **int32** | Active members of the named workspace; 0 when none is named. | 
**Budgets** | [**NullableAllocationHealthResponse**](AllocationHealthResponse.md) |  | 
**Ceilings** | [**NullableAllocationHealthResponse**](AllocationHealthResponse.md) |  | 

## Methods

### NewOverviewSummaryResponse

`func NewOverviewSummaryResponse(activeKeys int32, activeMembers int32, budgets NullableAllocationHealthResponse, ceilings NullableAllocationHealthResponse, ) *OverviewSummaryResponse`

NewOverviewSummaryResponse instantiates a new OverviewSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewSummaryResponseWithDefaults

`func NewOverviewSummaryResponseWithDefaults() *OverviewSummaryResponse`

NewOverviewSummaryResponseWithDefaults instantiates a new OverviewSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveKeys

`func (o *OverviewSummaryResponse) GetActiveKeys() int32`

GetActiveKeys returns the ActiveKeys field if non-nil, zero value otherwise.

### GetActiveKeysOk

`func (o *OverviewSummaryResponse) GetActiveKeysOk() (*int32, bool)`

GetActiveKeysOk returns a tuple with the ActiveKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveKeys

`func (o *OverviewSummaryResponse) SetActiveKeys(v int32)`

SetActiveKeys sets ActiveKeys field to given value.


### GetActiveMembers

`func (o *OverviewSummaryResponse) GetActiveMembers() int32`

GetActiveMembers returns the ActiveMembers field if non-nil, zero value otherwise.

### GetActiveMembersOk

`func (o *OverviewSummaryResponse) GetActiveMembersOk() (*int32, bool)`

GetActiveMembersOk returns a tuple with the ActiveMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMembers

`func (o *OverviewSummaryResponse) SetActiveMembers(v int32)`

SetActiveMembers sets ActiveMembers field to given value.


### GetBudgets

`func (o *OverviewSummaryResponse) GetBudgets() AllocationHealthResponse`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *OverviewSummaryResponse) GetBudgetsOk() (*AllocationHealthResponse, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *OverviewSummaryResponse) SetBudgets(v AllocationHealthResponse)`

SetBudgets sets Budgets field to given value.


### SetBudgetsNil

`func (o *OverviewSummaryResponse) SetBudgetsNil(b bool)`

 SetBudgetsNil sets the value for Budgets to be an explicit nil

### UnsetBudgets
`func (o *OverviewSummaryResponse) UnsetBudgets()`

UnsetBudgets ensures that no value is present for Budgets, not even an explicit nil
### GetCeilings

`func (o *OverviewSummaryResponse) GetCeilings() AllocationHealthResponse`

GetCeilings returns the Ceilings field if non-nil, zero value otherwise.

### GetCeilingsOk

`func (o *OverviewSummaryResponse) GetCeilingsOk() (*AllocationHealthResponse, bool)`

GetCeilingsOk returns a tuple with the Ceilings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeilings

`func (o *OverviewSummaryResponse) SetCeilings(v AllocationHealthResponse)`

SetCeilings sets Ceilings field to given value.


### SetCeilingsNil

`func (o *OverviewSummaryResponse) SetCeilingsNil(b bool)`

 SetCeilingsNil sets the value for Ceilings to be an explicit nil

### UnsetCeilings
`func (o *OverviewSummaryResponse) UnsetCeilings()`

UnsetCeilings ensures that no value is present for Ceilings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


