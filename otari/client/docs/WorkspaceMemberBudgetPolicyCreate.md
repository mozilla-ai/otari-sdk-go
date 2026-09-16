# WorkspaceMemberBudgetPolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** | The budget this workspace hands to every member | 
**ProviderKeyId** | Pointer to **NullableString** | Narrow the default to one provider instance; omit or null to apply to every provider. Must name a real instance: a blank value would materialize ceilings that never bind | [optional] 

## Methods

### NewWorkspaceMemberBudgetPolicyCreate

`func NewWorkspaceMemberBudgetPolicyCreate(budgetId string, ) *WorkspaceMemberBudgetPolicyCreate`

NewWorkspaceMemberBudgetPolicyCreate instantiates a new WorkspaceMemberBudgetPolicyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMemberBudgetPolicyCreateWithDefaults

`func NewWorkspaceMemberBudgetPolicyCreateWithDefaults() *WorkspaceMemberBudgetPolicyCreate`

NewWorkspaceMemberBudgetPolicyCreateWithDefaults instantiates a new WorkspaceMemberBudgetPolicyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *WorkspaceMemberBudgetPolicyCreate) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *WorkspaceMemberBudgetPolicyCreate) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *WorkspaceMemberBudgetPolicyCreate) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetProviderKeyId

`func (o *WorkspaceMemberBudgetPolicyCreate) GetProviderKeyId() string`

GetProviderKeyId returns the ProviderKeyId field if non-nil, zero value otherwise.

### GetProviderKeyIdOk

`func (o *WorkspaceMemberBudgetPolicyCreate) GetProviderKeyIdOk() (*string, bool)`

GetProviderKeyIdOk returns a tuple with the ProviderKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKeyId

`func (o *WorkspaceMemberBudgetPolicyCreate) SetProviderKeyId(v string)`

SetProviderKeyId sets ProviderKeyId field to given value.

### HasProviderKeyId

`func (o *WorkspaceMemberBudgetPolicyCreate) HasProviderKeyId() bool`

HasProviderKeyId returns a boolean if a field has been set.

### SetProviderKeyIdNil

`func (o *WorkspaceMemberBudgetPolicyCreate) SetProviderKeyIdNil(b bool)`

 SetProviderKeyIdNil sets the value for ProviderKeyId to be an explicit nil

### UnsetProviderKeyId
`func (o *WorkspaceMemberBudgetPolicyCreate) UnsetProviderKeyId()`

UnsetProviderKeyId ensures that no value is present for ProviderKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


