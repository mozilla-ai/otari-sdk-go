# WorstAllocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | **float32** |  | 
**BudgetId** | **string** |  | 
**Name** | **NullableString** | The row&#39;s own name, or null where nobody gave it one. | 
**ScopeId** | **NullableString** | The scope&#39;s id, so an unnamed ceiling can be named after it. | 
**ScopeType** | **NullableString** | What a spend ceiling caps (workspace, org_member, api_token, ...); null for a budget. | 
**Spent** | **float32** |  | 

## Methods

### NewWorstAllocationResponse

`func NewWorstAllocationResponse(allocated float32, budgetId string, name NullableString, scopeId NullableString, scopeType NullableString, spent float32, ) *WorstAllocationResponse`

NewWorstAllocationResponse instantiates a new WorstAllocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorstAllocationResponseWithDefaults

`func NewWorstAllocationResponseWithDefaults() *WorstAllocationResponse`

NewWorstAllocationResponseWithDefaults instantiates a new WorstAllocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *WorstAllocationResponse) GetAllocated() float32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *WorstAllocationResponse) GetAllocatedOk() (*float32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *WorstAllocationResponse) SetAllocated(v float32)`

SetAllocated sets Allocated field to given value.


### GetBudgetId

`func (o *WorstAllocationResponse) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *WorstAllocationResponse) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *WorstAllocationResponse) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetName

`func (o *WorstAllocationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorstAllocationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorstAllocationResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *WorstAllocationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorstAllocationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScopeId

`func (o *WorstAllocationResponse) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *WorstAllocationResponse) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *WorstAllocationResponse) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### SetScopeIdNil

`func (o *WorstAllocationResponse) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *WorstAllocationResponse) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetScopeType

`func (o *WorstAllocationResponse) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *WorstAllocationResponse) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *WorstAllocationResponse) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### SetScopeTypeNil

`func (o *WorstAllocationResponse) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *WorstAllocationResponse) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetSpent

`func (o *WorstAllocationResponse) GetSpent() float32`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *WorstAllocationResponse) GetSpentOk() (*float32, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *WorstAllocationResponse) SetSpent(v float32)`

SetSpent sets Spent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


