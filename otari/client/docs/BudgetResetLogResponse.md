# BudgetResetLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** |  | 
**Id** | **int32** |  | 
**NextResetAt** | **NullableString** |  | 
**PreviousSpend** | **float32** |  | 
**ResetAt** | **string** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewBudgetResetLogResponse

`func NewBudgetResetLogResponse(budgetId string, id int32, nextResetAt NullableString, previousSpend float32, resetAt string, userId NullableString, ) *BudgetResetLogResponse`

NewBudgetResetLogResponse instantiates a new BudgetResetLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetResetLogResponseWithDefaults

`func NewBudgetResetLogResponseWithDefaults() *BudgetResetLogResponse`

NewBudgetResetLogResponseWithDefaults instantiates a new BudgetResetLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *BudgetResetLogResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetResetLogResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetResetLogResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetId

`func (o *BudgetResetLogResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetResetLogResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetResetLogResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetNextResetAt

`func (o *BudgetResetLogResponse) GetNextResetAt() string`

GetNextResetAt returns the NextResetAt field if non-nil, zero value otherwise.

### GetNextResetAtOk

`func (o *BudgetResetLogResponse) GetNextResetAtOk() (*string, bool)`

GetNextResetAtOk returns a tuple with the NextResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextResetAt

`func (o *BudgetResetLogResponse) SetNextResetAt(v string)`

SetNextResetAt sets NextResetAt field to given value.


### SetNextResetAtNil

`func (o *BudgetResetLogResponse) SetNextResetAtNil(b bool)`

 SetNextResetAtNil sets the value for NextResetAt to be an explicit nil

### UnsetNextResetAt
`func (o *BudgetResetLogResponse) UnsetNextResetAt()`

UnsetNextResetAt ensures that no value is present for NextResetAt, not even an explicit nil
### GetPreviousSpend

`func (o *BudgetResetLogResponse) GetPreviousSpend() float32`

GetPreviousSpend returns the PreviousSpend field if non-nil, zero value otherwise.

### GetPreviousSpendOk

`func (o *BudgetResetLogResponse) GetPreviousSpendOk() (*float32, bool)`

GetPreviousSpendOk returns a tuple with the PreviousSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSpend

`func (o *BudgetResetLogResponse) SetPreviousSpend(v float32)`

SetPreviousSpend sets PreviousSpend field to given value.


### GetResetAt

`func (o *BudgetResetLogResponse) GetResetAt() string`

GetResetAt returns the ResetAt field if non-nil, zero value otherwise.

### GetResetAtOk

`func (o *BudgetResetLogResponse) GetResetAtOk() (*string, bool)`

GetResetAtOk returns a tuple with the ResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAt

`func (o *BudgetResetLogResponse) SetResetAt(v string)`

SetResetAt sets ResetAt field to given value.


### GetUserId

`func (o *BudgetResetLogResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BudgetResetLogResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BudgetResetLogResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *BudgetResetLogResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BudgetResetLogResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


