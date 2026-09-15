# UpdateBudgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | Pointer to **NullableInt32** |  | [optional] 
**MaxBudget** | Pointer to **NullableFloat32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RequestLimit** | Pointer to **NullableInt32** | Maximum requests over the period. Independent of max_budget; null is unlimited | [optional] 
**ResetAlignment** | Pointer to **NullableString** |  | [optional] 
**TokenLimit** | Pointer to **NullableInt32** | Maximum tokens over the period. Independent of max_budget; null is unlimited | [optional] 

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
### GetName

`func (o *UpdateBudgetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBudgetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBudgetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBudgetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateBudgetRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateBudgetRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRequestLimit

`func (o *UpdateBudgetRequest) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *UpdateBudgetRequest) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *UpdateBudgetRequest) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *UpdateBudgetRequest) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### SetRequestLimitNil

`func (o *UpdateBudgetRequest) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *UpdateBudgetRequest) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetResetAlignment

`func (o *UpdateBudgetRequest) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *UpdateBudgetRequest) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *UpdateBudgetRequest) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.

### HasResetAlignment

`func (o *UpdateBudgetRequest) HasResetAlignment() bool`

HasResetAlignment returns a boolean if a field has been set.

### SetResetAlignmentNil

`func (o *UpdateBudgetRequest) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *UpdateBudgetRequest) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetTokenLimit

`func (o *UpdateBudgetRequest) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *UpdateBudgetRequest) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *UpdateBudgetRequest) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *UpdateBudgetRequest) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### SetTokenLimitNil

`func (o *UpdateBudgetRequest) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *UpdateBudgetRequest) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


