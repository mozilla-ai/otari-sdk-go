# OrganizationBudgetUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | Pointer to **NullableInt32** | Seconds between resets, counted from the last one. Mutually exclusive with reset_alignment | [optional] 
**MaxBudget** | Pointer to **NullableFloat32** | Maximum spend in USD over one period; null caps nothing | [optional] 
**Name** | Pointer to **NullableString** | Admin-facing label for the budget | [optional] 
**RequestLimit** | Pointer to **NullableInt32** | Maximum requests over one period; null caps nothing. Independent of max_budget | [optional] 
**ResetAlignment** | Pointer to **NullableString** | Reset on a UTC calendar boundary instead of a fixed number of seconds, which is the only way to express a calendar month. Mutually exclusive with budget_duration_sec | [optional] 
**TokenLimit** | Pointer to **NullableInt32** | Maximum tokens over one period; null caps nothing. Independent of max_budget | [optional] 

## Methods

### NewOrganizationBudgetUpdate

`func NewOrganizationBudgetUpdate() *OrganizationBudgetUpdate`

NewOrganizationBudgetUpdate instantiates a new OrganizationBudgetUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationBudgetUpdateWithDefaults

`func NewOrganizationBudgetUpdateWithDefaults() *OrganizationBudgetUpdate`

NewOrganizationBudgetUpdateWithDefaults instantiates a new OrganizationBudgetUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetDurationSec

`func (o *OrganizationBudgetUpdate) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *OrganizationBudgetUpdate) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *OrganizationBudgetUpdate) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.

### HasBudgetDurationSec

`func (o *OrganizationBudgetUpdate) HasBudgetDurationSec() bool`

HasBudgetDurationSec returns a boolean if a field has been set.

### SetBudgetDurationSecNil

`func (o *OrganizationBudgetUpdate) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *OrganizationBudgetUpdate) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetMaxBudget

`func (o *OrganizationBudgetUpdate) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *OrganizationBudgetUpdate) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *OrganizationBudgetUpdate) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *OrganizationBudgetUpdate) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### SetMaxBudgetNil

`func (o *OrganizationBudgetUpdate) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *OrganizationBudgetUpdate) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil
### GetName

`func (o *OrganizationBudgetUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationBudgetUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationBudgetUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationBudgetUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationBudgetUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationBudgetUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRequestLimit

`func (o *OrganizationBudgetUpdate) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *OrganizationBudgetUpdate) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *OrganizationBudgetUpdate) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *OrganizationBudgetUpdate) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### SetRequestLimitNil

`func (o *OrganizationBudgetUpdate) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *OrganizationBudgetUpdate) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetResetAlignment

`func (o *OrganizationBudgetUpdate) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *OrganizationBudgetUpdate) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *OrganizationBudgetUpdate) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.

### HasResetAlignment

`func (o *OrganizationBudgetUpdate) HasResetAlignment() bool`

HasResetAlignment returns a boolean if a field has been set.

### SetResetAlignmentNil

`func (o *OrganizationBudgetUpdate) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *OrganizationBudgetUpdate) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetTokenLimit

`func (o *OrganizationBudgetUpdate) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *OrganizationBudgetUpdate) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *OrganizationBudgetUpdate) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *OrganizationBudgetUpdate) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### SetTokenLimitNil

`func (o *OrganizationBudgetUpdate) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *OrganizationBudgetUpdate) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


