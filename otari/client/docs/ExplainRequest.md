# ExplainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** | Simulate an API key&#39;s allow-list. Omit for unrestricted. | [optional] 
**BudgetRemainingUsd** | Pointer to **NullableFloat32** | Simulated budget remaining, USD. | [optional] 
**BudgetUsedPct** | Pointer to **NullableFloat32** | Simulated budget usage percentage. | [optional] 
**KeyId** | Pointer to **NullableString** | Evaluate conditions as this API key id. | [optional] 
**Name** | Pointer to **NullableString** | An existing policy to explain. | [optional] 
**Spec** | Pointer to **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | [optional] 
**UserId** | Pointer to **NullableString** | Evaluate conditions as this user. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Resolve &#x60;name&#x60; and the policy&#39;s candidate selectors in this workspace. Omit for the deployment&#39;s default workspace. | [optional] 

## Methods

### NewExplainRequest

`func NewExplainRequest() *ExplainRequest`

NewExplainRequest instantiates a new ExplainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainRequestWithDefaults

`func NewExplainRequestWithDefaults() *ExplainRequest`

NewExplainRequestWithDefaults instantiates a new ExplainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *ExplainRequest) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *ExplainRequest) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *ExplainRequest) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *ExplainRequest) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *ExplainRequest) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *ExplainRequest) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetBudgetRemainingUsd

`func (o *ExplainRequest) GetBudgetRemainingUsd() float32`

GetBudgetRemainingUsd returns the BudgetRemainingUsd field if non-nil, zero value otherwise.

### GetBudgetRemainingUsdOk

`func (o *ExplainRequest) GetBudgetRemainingUsdOk() (*float32, bool)`

GetBudgetRemainingUsdOk returns a tuple with the BudgetRemainingUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetRemainingUsd

`func (o *ExplainRequest) SetBudgetRemainingUsd(v float32)`

SetBudgetRemainingUsd sets BudgetRemainingUsd field to given value.

### HasBudgetRemainingUsd

`func (o *ExplainRequest) HasBudgetRemainingUsd() bool`

HasBudgetRemainingUsd returns a boolean if a field has been set.

### SetBudgetRemainingUsdNil

`func (o *ExplainRequest) SetBudgetRemainingUsdNil(b bool)`

 SetBudgetRemainingUsdNil sets the value for BudgetRemainingUsd to be an explicit nil

### UnsetBudgetRemainingUsd
`func (o *ExplainRequest) UnsetBudgetRemainingUsd()`

UnsetBudgetRemainingUsd ensures that no value is present for BudgetRemainingUsd, not even an explicit nil
### GetBudgetUsedPct

`func (o *ExplainRequest) GetBudgetUsedPct() float32`

GetBudgetUsedPct returns the BudgetUsedPct field if non-nil, zero value otherwise.

### GetBudgetUsedPctOk

`func (o *ExplainRequest) GetBudgetUsedPctOk() (*float32, bool)`

GetBudgetUsedPctOk returns a tuple with the BudgetUsedPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetUsedPct

`func (o *ExplainRequest) SetBudgetUsedPct(v float32)`

SetBudgetUsedPct sets BudgetUsedPct field to given value.

### HasBudgetUsedPct

`func (o *ExplainRequest) HasBudgetUsedPct() bool`

HasBudgetUsedPct returns a boolean if a field has been set.

### SetBudgetUsedPctNil

`func (o *ExplainRequest) SetBudgetUsedPctNil(b bool)`

 SetBudgetUsedPctNil sets the value for BudgetUsedPct to be an explicit nil

### UnsetBudgetUsedPct
`func (o *ExplainRequest) UnsetBudgetUsedPct()`

UnsetBudgetUsedPct ensures that no value is present for BudgetUsedPct, not even an explicit nil
### GetKeyId

`func (o *ExplainRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ExplainRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ExplainRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ExplainRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### SetKeyIdNil

`func (o *ExplainRequest) SetKeyIdNil(b bool)`

 SetKeyIdNil sets the value for KeyId to be an explicit nil

### UnsetKeyId
`func (o *ExplainRequest) UnsetKeyId()`

UnsetKeyId ensures that no value is present for KeyId, not even an explicit nil
### GetName

`func (o *ExplainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplainRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExplainRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExplainRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExplainRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSpec

`func (o *ExplainRequest) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ExplainRequest) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ExplainRequest) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ExplainRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *ExplainRequest) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *ExplainRequest) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetUserId

`func (o *ExplainRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExplainRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExplainRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExplainRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExplainRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExplainRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *ExplainRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ExplainRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ExplainRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ExplainRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *ExplainRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *ExplainRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


