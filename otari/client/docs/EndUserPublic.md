# EndUserPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **bool** |  | 
**BudgetId** | **NullableString** |  | 
**BudgetStartedAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**CurrentRequests** | **int32** |  | 
**CurrentTokens** | **int32** |  | 
**ExternalId** | **string** |  | 
**NextBudgetResetAt** | **NullableString** |  | 
**OwnerUserId** | **string** |  | 
**Reserved** | **float32** |  | 
**Spend** | **float32** |  | 
**UserId** | **string** |  | 

## Methods

### NewEndUserPublic

`func NewEndUserPublic(blocked bool, budgetId NullableString, budgetStartedAt NullableString, createdAt string, currentRequests int32, currentTokens int32, externalId string, nextBudgetResetAt NullableString, ownerUserId string, reserved float32, spend float32, userId string, ) *EndUserPublic`

NewEndUserPublic instantiates a new EndUserPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserPublicWithDefaults

`func NewEndUserPublicWithDefaults() *EndUserPublic`

NewEndUserPublicWithDefaults instantiates a new EndUserPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *EndUserPublic) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *EndUserPublic) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *EndUserPublic) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetBudgetId

`func (o *EndUserPublic) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *EndUserPublic) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *EndUserPublic) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### SetBudgetIdNil

`func (o *EndUserPublic) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *EndUserPublic) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetBudgetStartedAt

`func (o *EndUserPublic) GetBudgetStartedAt() string`

GetBudgetStartedAt returns the BudgetStartedAt field if non-nil, zero value otherwise.

### GetBudgetStartedAtOk

`func (o *EndUserPublic) GetBudgetStartedAtOk() (*string, bool)`

GetBudgetStartedAtOk returns a tuple with the BudgetStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetStartedAt

`func (o *EndUserPublic) SetBudgetStartedAt(v string)`

SetBudgetStartedAt sets BudgetStartedAt field to given value.


### SetBudgetStartedAtNil

`func (o *EndUserPublic) SetBudgetStartedAtNil(b bool)`

 SetBudgetStartedAtNil sets the value for BudgetStartedAt to be an explicit nil

### UnsetBudgetStartedAt
`func (o *EndUserPublic) UnsetBudgetStartedAt()`

UnsetBudgetStartedAt ensures that no value is present for BudgetStartedAt, not even an explicit nil
### GetCreatedAt

`func (o *EndUserPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EndUserPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EndUserPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrentRequests

`func (o *EndUserPublic) GetCurrentRequests() int32`

GetCurrentRequests returns the CurrentRequests field if non-nil, zero value otherwise.

### GetCurrentRequestsOk

`func (o *EndUserPublic) GetCurrentRequestsOk() (*int32, bool)`

GetCurrentRequestsOk returns a tuple with the CurrentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRequests

`func (o *EndUserPublic) SetCurrentRequests(v int32)`

SetCurrentRequests sets CurrentRequests field to given value.


### GetCurrentTokens

`func (o *EndUserPublic) GetCurrentTokens() int32`

GetCurrentTokens returns the CurrentTokens field if non-nil, zero value otherwise.

### GetCurrentTokensOk

`func (o *EndUserPublic) GetCurrentTokensOk() (*int32, bool)`

GetCurrentTokensOk returns a tuple with the CurrentTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTokens

`func (o *EndUserPublic) SetCurrentTokens(v int32)`

SetCurrentTokens sets CurrentTokens field to given value.


### GetExternalId

`func (o *EndUserPublic) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EndUserPublic) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EndUserPublic) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetNextBudgetResetAt

`func (o *EndUserPublic) GetNextBudgetResetAt() string`

GetNextBudgetResetAt returns the NextBudgetResetAt field if non-nil, zero value otherwise.

### GetNextBudgetResetAtOk

`func (o *EndUserPublic) GetNextBudgetResetAtOk() (*string, bool)`

GetNextBudgetResetAtOk returns a tuple with the NextBudgetResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBudgetResetAt

`func (o *EndUserPublic) SetNextBudgetResetAt(v string)`

SetNextBudgetResetAt sets NextBudgetResetAt field to given value.


### SetNextBudgetResetAtNil

`func (o *EndUserPublic) SetNextBudgetResetAtNil(b bool)`

 SetNextBudgetResetAtNil sets the value for NextBudgetResetAt to be an explicit nil

### UnsetNextBudgetResetAt
`func (o *EndUserPublic) UnsetNextBudgetResetAt()`

UnsetNextBudgetResetAt ensures that no value is present for NextBudgetResetAt, not even an explicit nil
### GetOwnerUserId

`func (o *EndUserPublic) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *EndUserPublic) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *EndUserPublic) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetReserved

`func (o *EndUserPublic) GetReserved() float32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *EndUserPublic) GetReservedOk() (*float32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *EndUserPublic) SetReserved(v float32)`

SetReserved sets Reserved field to given value.


### GetSpend

`func (o *EndUserPublic) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *EndUserPublic) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *EndUserPublic) SetSpend(v float32)`

SetSpend sets Spend field to given value.


### GetUserId

`func (o *EndUserPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EndUserPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EndUserPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


