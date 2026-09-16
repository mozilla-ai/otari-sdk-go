# CreateScopedBudgetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** | The budget this ceiling enforces; its limit and period are read through it | 
**Name** | Pointer to **NullableString** | Admin-facing label for this ceiling | [optional] 
**ProviderKeyId** | Pointer to **NullableString** | Narrow the cap to one provider instance; omit or null to cap spend across every provider. A blank value would store a ceiling that never binds, so it is refused; this does not check that the value names a configured provider instance | [optional] 
**ScopeId** | **string** | Id of the capped identity: an organization, workspace, membership row, or API key | 
**ScopeType** | **string** | Which kind of identity this ceiling caps | 

## Methods

### NewCreateScopedBudgetRequest

`func NewCreateScopedBudgetRequest(budgetId string, scopeId string, scopeType string, ) *CreateScopedBudgetRequest`

NewCreateScopedBudgetRequest instantiates a new CreateScopedBudgetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScopedBudgetRequestWithDefaults

`func NewCreateScopedBudgetRequestWithDefaults() *CreateScopedBudgetRequest`

NewCreateScopedBudgetRequestWithDefaults instantiates a new CreateScopedBudgetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *CreateScopedBudgetRequest) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *CreateScopedBudgetRequest) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *CreateScopedBudgetRequest) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetName

`func (o *CreateScopedBudgetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateScopedBudgetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateScopedBudgetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateScopedBudgetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateScopedBudgetRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateScopedBudgetRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderKeyId

`func (o *CreateScopedBudgetRequest) GetProviderKeyId() string`

GetProviderKeyId returns the ProviderKeyId field if non-nil, zero value otherwise.

### GetProviderKeyIdOk

`func (o *CreateScopedBudgetRequest) GetProviderKeyIdOk() (*string, bool)`

GetProviderKeyIdOk returns a tuple with the ProviderKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKeyId

`func (o *CreateScopedBudgetRequest) SetProviderKeyId(v string)`

SetProviderKeyId sets ProviderKeyId field to given value.

### HasProviderKeyId

`func (o *CreateScopedBudgetRequest) HasProviderKeyId() bool`

HasProviderKeyId returns a boolean if a field has been set.

### SetProviderKeyIdNil

`func (o *CreateScopedBudgetRequest) SetProviderKeyIdNil(b bool)`

 SetProviderKeyIdNil sets the value for ProviderKeyId to be an explicit nil

### UnsetProviderKeyId
`func (o *CreateScopedBudgetRequest) UnsetProviderKeyId()`

UnsetProviderKeyId ensures that no value is present for ProviderKeyId, not even an explicit nil
### GetScopeId

`func (o *CreateScopedBudgetRequest) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateScopedBudgetRequest) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateScopedBudgetRequest) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetScopeType

`func (o *CreateScopedBudgetRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *CreateScopedBudgetRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *CreateScopedBudgetRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


