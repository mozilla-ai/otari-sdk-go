# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **NullableString** | Optional admin-facing alias | [optional] 
**Blocked** | Pointer to **bool** | Whether user is blocked | [optional] [default to false]
**BudgetId** | Pointer to **NullableString** | Optional budget ID | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional metadata | [optional] 
**UserId** | **string** | Unique user identifier | 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(userId string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *CreateUserRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateUserRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateUserRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CreateUserRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CreateUserRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CreateUserRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetBlocked

`func (o *CreateUserRequest) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *CreateUserRequest) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *CreateUserRequest) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *CreateUserRequest) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBudgetId

`func (o *CreateUserRequest) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *CreateUserRequest) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *CreateUserRequest) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *CreateUserRequest) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *CreateUserRequest) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *CreateUserRequest) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetMetadata

`func (o *CreateUserRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateUserRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateUserRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateUserRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUserId

`func (o *CreateUserRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateUserRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateUserRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


