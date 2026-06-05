# UpdateBudgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | Pointer to **NullableInt32** |  | [optional] 
**MaxBudget** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewUpdateBudgetRequest

`func NewUpdateBudgetRequest() *UpdateBudgetRequest`

NewUpdateBudgetRequest instantiates a new UpdateBudgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBudgetRequestWithDefaults

`func NewUpdateBudgetRequestWithDefaults() *UpdateBudgetRequest`

NewUpdateBudgetRequestWithDefaults instantiates a new UpdateBudgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetDurationSec

`func (o *UpdateBudgetRequest) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *UpdateBudgetRequest) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *UpdateBudgetRequest) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.

### HasBudgetDurationSec

`func (o *UpdateBudgetRequest) HasBudgetDurationSec() bool`

HasBudgetDurationSec returns a boolean if a field has been set.

### SetBudgetDurationSecNil

`func (o *UpdateBudgetRequest) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *UpdateBudgetRequest) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetMaxBudget

`func (o *UpdateBudgetRequest) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *UpdateBudgetRequest) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *UpdateBudgetRequest) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *UpdateBudgetRequest) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### SetMaxBudgetNil

`func (o *UpdateBudgetRequest) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *UpdateBudgetRequest) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


