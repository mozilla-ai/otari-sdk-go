# ScopedBudgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | **NullableInt32** |  | 
**BudgetId** | **string** |  | 
**CreatedAt** | **string** |  | 
**CurrentRequests** | **int32** |  | 
**CurrentSpend** | **float32** |  | 
**CurrentTokens** | **int32** |  | 
**Id** | **string** |  | 
**MaxBudget** | **NullableFloat32** |  | 
**Name** | **NullableString** |  | 
**PeriodEnd** | **NullableString** |  | 
**PeriodStart** | **NullableString** |  | 
**ProviderKeyId** | **NullableString** |  | 
**RequestLimit** | **NullableInt32** |  | 
**ReservedRequests** | **int32** |  | 
**ReservedSpend** | **float32** |  | 
**ReservedTokens** | **int32** |  | 
**ResetAlignment** | **NullableString** |  | 
**ScopeId** | **string** |  | 
**ScopeType** | **string** |  | 
**TokenLimit** | **NullableInt32** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewScopedBudgetResponse

`func NewScopedBudgetResponse(budgetDurationSec NullableInt32, budgetId string, createdAt string, currentRequests int32, currentSpend float32, currentTokens int32, id string, maxBudget NullableFloat32, name NullableString, periodEnd NullableString, periodStart NullableString, providerKeyId NullableString, requestLimit NullableInt32, reservedRequests int32, reservedSpend float32, reservedTokens int32, resetAlignment NullableString, scopeId string, scopeType string, tokenLimit NullableInt32, updatedAt string, ) *ScopedBudgetResponse`

NewScopedBudgetResponse instantiates a new ScopedBudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedBudgetResponseWithDefaults

`func NewScopedBudgetResponseWithDefaults() *ScopedBudgetResponse`

NewScopedBudgetResponseWithDefaults instantiates a new ScopedBudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetDurationSec

`func (o *ScopedBudgetResponse) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *ScopedBudgetResponse) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *ScopedBudgetResponse) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.


### SetBudgetDurationSecNil

`func (o *ScopedBudgetResponse) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *ScopedBudgetResponse) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetBudgetId

`func (o *ScopedBudgetResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *ScopedBudgetResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *ScopedBudgetResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetCreatedAt

`func (o *ScopedBudgetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScopedBudgetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScopedBudgetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrentRequests

`func (o *ScopedBudgetResponse) GetCurrentRequests() int32`

GetCurrentRequests returns the CurrentRequests field if non-nil, zero value otherwise.

### GetCurrentRequestsOk

`func (o *ScopedBudgetResponse) GetCurrentRequestsOk() (*int32, bool)`

GetCurrentRequestsOk returns a tuple with the CurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRequests

`func (o *ScopedBudgetResponse) SetCurrentRequests(v int32)`

SetCurrentRequests sets CurrentRequests field to given value.


### GetCurrentSpend

`func (o *ScopedBudgetResponse) GetCurrentSpend() float32`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *ScopedBudgetResponse) GetCurrentSpendOk() (*float32, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *ScopedBudgetResponse) SetCurrentSpend(v float32)`

SetCurrentSpend sets CurrentSpend field to given value.


### GetCurrentTokens

`func (o *ScopedBudgetResponse) GetCurrentTokens() int32`

GetCurrentTokens returns the CurrentTokens field if non-nil, zero value otherwise.

### GetCurrentTokensOk

`func (o *ScopedBudgetResponse) GetCurrentTokensOk() (*int32, bool)`

GetCurrentTokensOk returns a tuple with the CurrentTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTokens

`func (o *ScopedBudgetResponse) SetCurrentTokens(v int32)`

SetCurrentTokens sets CurrentTokens field to given value.


### GetId

`func (o *ScopedBudgetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopedBudgetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopedBudgetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxBudget

`func (o *ScopedBudgetResponse) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *ScopedBudgetResponse) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *ScopedBudgetResponse) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.


### SetMaxBudgetNil

`func (o *ScopedBudgetResponse) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *ScopedBudgetResponse) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil
### GetName

`func (o *ScopedBudgetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopedBudgetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopedBudgetResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ScopedBudgetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScopedBudgetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPeriodEnd

`func (o *ScopedBudgetResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ScopedBudgetResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ScopedBudgetResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### SetPeriodEndNil

`func (o *ScopedBudgetResponse) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *ScopedBudgetResponse) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetPeriodStart

`func (o *ScopedBudgetResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ScopedBudgetResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ScopedBudgetResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### SetPeriodStartNil

`func (o *ScopedBudgetResponse) SetPeriodStartNil(b bool)`

 SetPeriodStartNil sets the value for PeriodStart to be an explicit nil

### UnsetPeriodStart
`func (o *ScopedBudgetResponse) UnsetPeriodStart()`

UnsetPeriodStart ensures that no value is present for PeriodStart, not even an explicit nil
### GetProviderKeyId

`func (o *ScopedBudgetResponse) GetProviderKeyId() string`

GetProviderKeyId returns the ProviderKeyId field if non-nil, zero value otherwise.

### GetProviderKeyIdOk

`func (o *ScopedBudgetResponse) GetProviderKeyIdOk() (*string, bool)`

GetProviderKeyIdOk returns a tuple with the ProviderKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKeyId

`func (o *ScopedBudgetResponse) SetProviderKeyId(v string)`

SetProviderKeyId sets ProviderKeyId field to given value.


### SetProviderKeyIdNil

`func (o *ScopedBudgetResponse) SetProviderKeyIdNil(b bool)`

 SetProviderKeyIdNil sets the value for ProviderKeyId to be an explicit nil

### UnsetProviderKeyId
`func (o *ScopedBudgetResponse) UnsetProviderKeyId()`

UnsetProviderKeyId ensures that no value is present for ProviderKeyId, not even an explicit nil
### GetRequestLimit

`func (o *ScopedBudgetResponse) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *ScopedBudgetResponse) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *ScopedBudgetResponse) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.


### SetRequestLimitNil

`func (o *ScopedBudgetResponse) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *ScopedBudgetResponse) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetReservedRequests

`func (o *ScopedBudgetResponse) GetReservedRequests() int32`

GetReservedRequests returns the ReservedRequests field if non-nil, zero value otherwise.

### GetReservedRequestsOk

`func (o *ScopedBudgetResponse) GetReservedRequestsOk() (*int32, bool)`

GetReservedRequestsOk returns a tuple with the ReservedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRequests

`func (o *ScopedBudgetResponse) SetReservedRequests(v int32)`

SetReservedRequests sets ReservedRequests field to given value.


### GetReservedSpend

`func (o *ScopedBudgetResponse) GetReservedSpend() float32`

GetReservedSpend returns the ReservedSpend field if non-nil, zero value otherwise.

### GetReservedSpendOk

`func (o *ScopedBudgetResponse) GetReservedSpendOk() (*float32, bool)`

GetReservedSpendOk returns a tuple with the ReservedSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedSpend

`func (o *ScopedBudgetResponse) SetReservedSpend(v float32)`

SetReservedSpend sets ReservedSpend field to given value.


### GetReservedTokens

`func (o *ScopedBudgetResponse) GetReservedTokens() int32`

GetReservedTokens returns the ReservedTokens field if non-nil, zero value otherwise.

### GetReservedTokensOk

`func (o *ScopedBudgetResponse) GetReservedTokensOk() (*int32, bool)`

GetReservedTokensOk returns a tuple with the ReservedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedTokens

`func (o *ScopedBudgetResponse) SetReservedTokens(v int32)`

SetReservedTokens sets ReservedTokens field to given value.


### GetResetAlignment

`func (o *ScopedBudgetResponse) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *ScopedBudgetResponse) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *ScopedBudgetResponse) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.


### SetResetAlignmentNil

`func (o *ScopedBudgetResponse) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *ScopedBudgetResponse) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetScopeId

`func (o *ScopedBudgetResponse) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ScopedBudgetResponse) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ScopedBudgetResponse) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetScopeType

`func (o *ScopedBudgetResponse) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *ScopedBudgetResponse) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *ScopedBudgetResponse) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetTokenLimit

`func (o *ScopedBudgetResponse) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *ScopedBudgetResponse) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *ScopedBudgetResponse) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.


### SetTokenLimitNil

`func (o *ScopedBudgetResponse) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *ScopedBudgetResponse) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil
### GetUpdatedAt

`func (o *ScopedBudgetResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScopedBudgetResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScopedBudgetResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


