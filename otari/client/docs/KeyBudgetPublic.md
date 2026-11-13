# KeyBudgetPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetsEnforced** | **bool** | False when budgets are off for this key: it is budget-exempt, or the deployment runs with budget_strategy disabled. True otherwise, even when no cap applies | 
**Limits** | [**[]KeyLimitPublic**](KeyLimitPublic.md) | The key&#39;s user&#39;s budget first, then each spend ceiling, the most specific first | 
**RemainingUsd** | **NullableFloat32** | The fewest dollars any dollar cap on every provider leaves; null when none applies. A token, request or provider-narrowed cap in limits can refuse a request sooner | 

## Methods

### NewKeyBudgetPublic

`func NewKeyBudgetPublic(budgetsEnforced bool, limits []KeyLimitPublic, remainingUsd NullableFloat32, ) *KeyBudgetPublic`

NewKeyBudgetPublic instantiates a new KeyBudgetPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyBudgetPublicWithDefaults

`func NewKeyBudgetPublicWithDefaults() *KeyBudgetPublic`

NewKeyBudgetPublicWithDefaults instantiates a new KeyBudgetPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetsEnforced

`func (o *KeyBudgetPublic) GetBudgetsEnforced() bool`

GetBudgetsEnforced returns the BudgetsEnforced field if non-nil, zero value otherwise.

### GetBudgetsEnforcedOk

`func (o *KeyBudgetPublic) GetBudgetsEnforcedOk() (*bool, bool)`

GetBudgetsEnforcedOk returns a tuple with the BudgetsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetsEnforced

`func (o *KeyBudgetPublic) SetBudgetsEnforced(v bool)`

SetBudgetsEnforced sets BudgetsEnforced field to given value.


### GetLimits

`func (o *KeyBudgetPublic) GetLimits() []KeyLimitPublic`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KeyBudgetPublic) GetLimitsOk() (*[]KeyLimitPublic, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KeyBudgetPublic) SetLimits(v []KeyLimitPublic)`

SetLimits sets Limits field to given value.


### GetRemainingUsd

`func (o *KeyBudgetPublic) GetRemainingUsd() float32`

GetRemainingUsd returns the RemainingUsd field if non-nil, zero value otherwise.

### GetRemainingUsdOk

`func (o *KeyBudgetPublic) GetRemainingUsdOk() (*float32, bool)`

GetRemainingUsdOk returns a tuple with the RemainingUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingUsd

`func (o *KeyBudgetPublic) SetRemainingUsd(v float32)`

SetRemainingUsd sets RemainingUsd field to given value.


### SetRemainingUsdNil

`func (o *KeyBudgetPublic) SetRemainingUsdNil(b bool)`

 SetRemainingUsdNil sets the value for RemainingUsd to be an explicit nil

### UnsetRemainingUsd
`func (o *KeyBudgetPublic) UnsetRemainingUsd()`

UnsetRemainingUsd ensures that no value is present for RemainingUsd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


