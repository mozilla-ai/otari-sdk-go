# BudgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | **NullableInt32** |  | 
**BudgetId** | **string** |  | 
**CreatedAt** | **string** |  | 
**MaxBudget** | **NullableFloat32** |  | 
**Name** | **NullableString** |  | 
**OrganizationId** | **NullableString** |  | 
**RequestLimit** | **NullableInt32** |  | 
**ResetAlignment** | **NullableString** |  | 
**TokenLimit** | **NullableInt32** |  | 
**TotalReserved** | Pointer to **float32** |  | [optional] [default to 0.0]
**TotalSpend** | Pointer to **float32** |  | [optional] [default to 0.0]
**UpdatedAt** | **string** |  | 
**UserCount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewBudgetResponse

`func NewBudgetResponse(budgetDurationSec NullableInt32, budgetId string, createdAt string, maxBudget NullableFloat32, name NullableString, organizationId NullableString, requestLimit NullableInt32, resetAlignment NullableString, tokenLimit NullableInt32, updatedAt string, ) *BudgetResponse`

NewBudgetResponse instantiates a new BudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetResponseWithDefaults

`func NewBudgetResponseWithDefaults() *BudgetResponse`

NewBudgetResponseWithDefaults instantiates a new BudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetDurationSec

`func (o *BudgetResponse) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *BudgetResponse) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *BudgetResponse) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.


### SetBudgetDurationSecNil

`func (o *BudgetResponse) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *BudgetResponse) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetBudgetId

`func (o *BudgetResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetCreatedAt

`func (o *BudgetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BudgetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BudgetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMaxBudget

`func (o *BudgetResponse) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *BudgetResponse) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *BudgetResponse) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.


### SetMaxBudgetNil

`func (o *BudgetResponse) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *BudgetResponse) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil
### GetName

`func (o *BudgetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BudgetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BudgetResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BudgetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BudgetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *BudgetResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BudgetResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BudgetResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *BudgetResponse) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *BudgetResponse) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetRequestLimit

`func (o *BudgetResponse) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *BudgetResponse) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *BudgetResponse) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.


### SetRequestLimitNil

`func (o *BudgetResponse) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *BudgetResponse) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetResetAlignment

`func (o *BudgetResponse) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *BudgetResponse) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *BudgetResponse) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.


### SetResetAlignmentNil

`func (o *BudgetResponse) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *BudgetResponse) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetTokenLimit

`func (o *BudgetResponse) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *BudgetResponse) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *BudgetResponse) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.


### SetTokenLimitNil

`func (o *BudgetResponse) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *BudgetResponse) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil
### GetTotalReserved

`func (o *BudgetResponse) GetTotalReserved() float32`

GetTotalReserved returns the TotalReserved field if non-nil, zero value otherwise.

### GetTotalReservedOk

`func (o *BudgetResponse) GetTotalReservedOk() (*float32, bool)`

GetTotalReservedOk returns a tuple with the TotalReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReserved

`func (o *BudgetResponse) SetTotalReserved(v float32)`

SetTotalReserved sets TotalReserved field to given value.

### HasTotalReserved

`func (o *BudgetResponse) HasTotalReserved() bool`

HasTotalReserved returns a boolean if a field has been set.

### GetTotalSpend

`func (o *BudgetResponse) GetTotalSpend() float32`

GetTotalSpend returns the TotalSpend field if non-nil, zero value otherwise.

### GetTotalSpendOk

`func (o *BudgetResponse) GetTotalSpendOk() (*float32, bool)`

GetTotalSpendOk returns a tuple with the TotalSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpend

`func (o *BudgetResponse) SetTotalSpend(v float32)`

SetTotalSpend sets TotalSpend field to given value.

### HasTotalSpend

`func (o *BudgetResponse) HasTotalSpend() bool`

HasTotalSpend returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BudgetResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BudgetResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BudgetResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserCount

`func (o *BudgetResponse) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *BudgetResponse) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *BudgetResponse) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *BudgetResponse) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


