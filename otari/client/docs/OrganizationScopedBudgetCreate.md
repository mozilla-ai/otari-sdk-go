# OrganizationScopedBudgetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | **string** | The budget this ceiling enforces, which must be one this organization owns | 
**Name** | Pointer to **NullableString** | Admin-facing label for this ceiling | [optional] 
**ProviderKeyId** | Pointer to **NullableString** | Narrow the cap to one provider instance; omit or null to cap spend across every provider. Must name a real instance: a blank value would store a ceiling that never binds | [optional] 
**ScopeId** | **string** | Id of the capped identity: this organization, one of its workspaces, a membership in either, or an API key in one | 
**ScopeType** | **string** | Which kind of identity this ceiling caps | 

## Methods

### NewOrganizationScopedBudgetCreate

`func NewOrganizationScopedBudgetCreate(budgetId string, scopeId string, scopeType string, ) *OrganizationScopedBudgetCreate`

NewOrganizationScopedBudgetCreate instantiates a new OrganizationScopedBudgetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationScopedBudgetCreateWithDefaults

`func NewOrganizationScopedBudgetCreateWithDefaults() *OrganizationScopedBudgetCreate`

NewOrganizationScopedBudgetCreateWithDefaults instantiates a new OrganizationScopedBudgetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *OrganizationScopedBudgetCreate) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *OrganizationScopedBudgetCreate) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *OrganizationScopedBudgetCreate) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetName

`func (o *OrganizationScopedBudgetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationScopedBudgetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationScopedBudgetCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationScopedBudgetCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationScopedBudgetCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationScopedBudgetCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderKeyId

`func (o *OrganizationScopedBudgetCreate) GetProviderKeyId() string`

GetProviderKeyId returns the ProviderKeyId field if non-nil, zero value otherwise.

### GetProviderKeyIdOk

`func (o *OrganizationScopedBudgetCreate) GetProviderKeyIdOk() (*string, bool)`

GetProviderKeyIdOk returns a tuple with the ProviderKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKeyId

`func (o *OrganizationScopedBudgetCreate) SetProviderKeyId(v string)`

SetProviderKeyId sets ProviderKeyId field to given value.

### HasProviderKeyId

`func (o *OrganizationScopedBudgetCreate) HasProviderKeyId() bool`

HasProviderKeyId returns a boolean if a field has been set.

### SetProviderKeyIdNil

`func (o *OrganizationScopedBudgetCreate) SetProviderKeyIdNil(b bool)`

 SetProviderKeyIdNil sets the value for ProviderKeyId to be an explicit nil

### UnsetProviderKeyId
`func (o *OrganizationScopedBudgetCreate) UnsetProviderKeyId()`

UnsetProviderKeyId ensures that no value is present for ProviderKeyId, not even an explicit nil
### GetScopeId

`func (o *OrganizationScopedBudgetCreate) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *OrganizationScopedBudgetCreate) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *OrganizationScopedBudgetCreate) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetScopeType

`func (o *OrganizationScopedBudgetCreate) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *OrganizationScopedBudgetCreate) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *OrganizationScopedBudgetCreate) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


