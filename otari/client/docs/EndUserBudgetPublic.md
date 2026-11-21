# EndUserBudgetPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedEndUserCount** | **int32** |  | 
**BudgetDurationSec** | Pointer to **NullableInt32** |  | [optional] 
**BudgetId** | **string** |  | 
**EndUserCount** | **int32** | End users of the key&#39;s user on this budget, blocked ones included | 
**IsDefault** | **bool** | Whether a new end user starts here when its request names no budget | 
**MaxBudget** | Pointer to **NullableFloat32** | Dollars each end user may spend over the period | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RequestLimit** | Pointer to **NullableInt32** |  | [optional] 
**ResetAlignment** | Pointer to **NullableString** |  | [optional] 
**RpmLimit** | Pointer to **NullableInt32** |  | [optional] 
**TokenLimit** | Pointer to **NullableInt32** |  | [optional] 
**TpmLimit** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEndUserBudgetPublic

`func NewEndUserBudgetPublic(blockedEndUserCount int32, budgetId string, endUserCount int32, isDefault bool, ) *EndUserBudgetPublic`

NewEndUserBudgetPublic instantiates a new EndUserBudgetPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserBudgetPublicWithDefaults

`func NewEndUserBudgetPublicWithDefaults() *EndUserBudgetPublic`

NewEndUserBudgetPublicWithDefaults instantiates a new EndUserBudgetPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedEndUserCount

`func (o *EndUserBudgetPublic) GetBlockedEndUserCount() int32`

GetBlockedEndUserCount returns the BlockedEndUserCount field if non-nil, zero value otherwise.

### GetBlockedEndUserCountOk

`func (o *EndUserBudgetPublic) GetBlockedEndUserCountOk() (*int32, bool)`

GetBlockedEndUserCountOk returns a tuple with the BlockedEndUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEndUserCount

`func (o *EndUserBudgetPublic) SetBlockedEndUserCount(v int32)`

SetBlockedEndUserCount sets BlockedEndUserCount field to given value.


### GetBudgetDurationSec

`func (o *EndUserBudgetPublic) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *EndUserBudgetPublic) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *EndUserBudgetPublic) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.

### HasBudgetDurationSec

`func (o *EndUserBudgetPublic) HasBudgetDurationSec() bool`

HasBudgetDurationSec returns a boolean if a field has been set.

### SetBudgetDurationSecNil

`func (o *EndUserBudgetPublic) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *EndUserBudgetPublic) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetBudgetId

`func (o *EndUserBudgetPublic) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *EndUserBudgetPublic) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *EndUserBudgetPublic) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetEndUserCount

`func (o *EndUserBudgetPublic) GetEndUserCount() int32`

GetEndUserCount returns the EndUserCount field if non-nil, zero value otherwise.

### GetEndUserCountOk

`func (o *EndUserBudgetPublic) GetEndUserCountOk() (*int32, bool)`

GetEndUserCountOk returns a tuple with the EndUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserCount

`func (o *EndUserBudgetPublic) SetEndUserCount(v int32)`

SetEndUserCount sets EndUserCount field to given value.


### GetIsDefault

`func (o *EndUserBudgetPublic) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EndUserBudgetPublic) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EndUserBudgetPublic) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetMaxBudget

`func (o *EndUserBudgetPublic) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *EndUserBudgetPublic) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *EndUserBudgetPublic) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *EndUserBudgetPublic) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### SetMaxBudgetNil

`func (o *EndUserBudgetPublic) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *EndUserBudgetPublic) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil
### GetName

`func (o *EndUserBudgetPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndUserBudgetPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndUserBudgetPublic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndUserBudgetPublic) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EndUserBudgetPublic) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EndUserBudgetPublic) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRequestLimit

`func (o *EndUserBudgetPublic) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *EndUserBudgetPublic) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *EndUserBudgetPublic) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.

### HasRequestLimit

`func (o *EndUserBudgetPublic) HasRequestLimit() bool`

HasRequestLimit returns a boolean if a field has been set.

### SetRequestLimitNil

`func (o *EndUserBudgetPublic) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *EndUserBudgetPublic) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetResetAlignment

`func (o *EndUserBudgetPublic) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *EndUserBudgetPublic) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *EndUserBudgetPublic) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.

### HasResetAlignment

`func (o *EndUserBudgetPublic) HasResetAlignment() bool`

HasResetAlignment returns a boolean if a field has been set.

### SetResetAlignmentNil

`func (o *EndUserBudgetPublic) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *EndUserBudgetPublic) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetRpmLimit

`func (o *EndUserBudgetPublic) GetRpmLimit() int32`

GetRpmLimit returns the RpmLimit field if non-nil, zero value otherwise.

### GetRpmLimitOk

`func (o *EndUserBudgetPublic) GetRpmLimitOk() (*int32, bool)`

GetRpmLimitOk returns a tuple with the RpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmLimit

`func (o *EndUserBudgetPublic) SetRpmLimit(v int32)`

SetRpmLimit sets RpmLimit field to given value.

### HasRpmLimit

`func (o *EndUserBudgetPublic) HasRpmLimit() bool`

HasRpmLimit returns a boolean if a field has been set.

### SetRpmLimitNil

`func (o *EndUserBudgetPublic) SetRpmLimitNil(b bool)`

 SetRpmLimitNil sets the value for RpmLimit to be an explicit nil

### UnsetRpmLimit
`func (o *EndUserBudgetPublic) UnsetRpmLimit()`

UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
### GetTokenLimit

`func (o *EndUserBudgetPublic) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *EndUserBudgetPublic) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *EndUserBudgetPublic) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.

### HasTokenLimit

`func (o *EndUserBudgetPublic) HasTokenLimit() bool`

HasTokenLimit returns a boolean if a field has been set.

### SetTokenLimitNil

`func (o *EndUserBudgetPublic) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *EndUserBudgetPublic) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil
### GetTpmLimit

`func (o *EndUserBudgetPublic) GetTpmLimit() int32`

GetTpmLimit returns the TpmLimit field if non-nil, zero value otherwise.

### GetTpmLimitOk

`func (o *EndUserBudgetPublic) GetTpmLimitOk() (*int32, bool)`

GetTpmLimitOk returns a tuple with the TpmLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmLimit

`func (o *EndUserBudgetPublic) SetTpmLimit(v int32)`

SetTpmLimit sets TpmLimit field to given value.

### HasTpmLimit

`func (o *EndUserBudgetPublic) HasTpmLimit() bool`

HasTpmLimit returns a boolean if a field has been set.

### SetTpmLimitNil

`func (o *EndUserBudgetPublic) SetTpmLimitNil(b bool)`

 SetTpmLimitNil sets the value for TpmLimit to be an explicit nil

### UnsetTpmLimit
`func (o *EndUserBudgetPublic) UnsetTpmLimit()`

UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


