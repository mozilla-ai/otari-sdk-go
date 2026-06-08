# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **NullableString** |  | 
**Blocked** | **bool** |  | 
**BudgetId** | **NullableString** |  | 
**BudgetStartedAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**NextBudgetResetAt** | **NullableString** |  | 
**Reserved** | **float32** |  | 
**Spend** | **float32** |  | 
**UpdatedAt** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewUserResponse

`func NewUserResponse(alias NullableString, blocked bool, budgetId NullableString, budgetStartedAt NullableString, createdAt string, metadata map[string]interface{}, nextBudgetResetAt NullableString, reserved float32, spend float32, updatedAt string, userId string, ) *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *UserResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *UserResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *UserResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *UserResponse) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *UserResponse) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetBlocked

`func (o *UserResponse) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *UserResponse) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *UserResponse) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetBudgetId

`func (o *UserResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *UserResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *UserResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### SetBudgetIdNil

`func (o *UserResponse) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *UserResponse) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetBudgetStartedAt

`func (o *UserResponse) GetBudgetStartedAt() string`

GetBudgetStartedAt returns the BudgetStartedAt field if non-nil, zero value otherwise.

### GetBudgetStartedAtOk

`func (o *UserResponse) GetBudgetStartedAtOk() (*string, bool)`

GetBudgetStartedAtOk returns a tuple with the BudgetStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetStartedAt

`func (o *UserResponse) SetBudgetStartedAt(v string)`

SetBudgetStartedAt sets BudgetStartedAt field to given value.


### SetBudgetStartedAtNil

`func (o *UserResponse) SetBudgetStartedAtNil(b bool)`

 SetBudgetStartedAtNil sets the value for BudgetStartedAt to be an explicit nil

### UnsetBudgetStartedAt
`func (o *UserResponse) UnsetBudgetStartedAt()`

UnsetBudgetStartedAt ensures that no value is present for BudgetStartedAt, not even an explicit nil
### GetCreatedAt

`func (o *UserResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetadata

`func (o *UserResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetNextBudgetResetAt

`func (o *UserResponse) GetNextBudgetResetAt() string`

GetNextBudgetResetAt returns the NextBudgetResetAt field if non-nil, zero value otherwise.

### GetNextBudgetResetAtOk

`func (o *UserResponse) GetNextBudgetResetAtOk() (*string, bool)`

GetNextBudgetResetAtOk returns a tuple with the NextBudgetResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBudgetResetAt

`func (o *UserResponse) SetNextBudgetResetAt(v string)`

SetNextBudgetResetAt sets NextBudgetResetAt field to given value.


### SetNextBudgetResetAtNil

`func (o *UserResponse) SetNextBudgetResetAtNil(b bool)`

 SetNextBudgetResetAtNil sets the value for NextBudgetResetAt to be an explicit nil

### UnsetNextBudgetResetAt
`func (o *UserResponse) UnsetNextBudgetResetAt()`

UnsetNextBudgetResetAt ensures that no value is present for NextBudgetResetAt, not even an explicit nil
### GetReserved

`func (o *UserResponse) GetReserved() float32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *UserResponse) GetReservedOk() (*float32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *UserResponse) SetReserved(v float32)`

SetReserved sets Reserved field to given value.


### GetSpend

`func (o *UserResponse) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *UserResponse) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *UserResponse) SetSpend(v float32)`

SetSpend sets Spend field to given value.


### GetUpdatedAt

`func (o *UserResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *UserResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


