# DecisionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **NullableFloat32** | The provider&#39;s own charge in USD, when it reports one | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**OutputTokens** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewDecisionUsage

`func NewDecisionUsage() *DecisionUsage`

NewDecisionUsage instantiates a new DecisionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionUsageWithDefaults

`func NewDecisionUsageWithDefaults() *DecisionUsage`

NewDecisionUsageWithDefaults instantiates a new DecisionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *DecisionUsage) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DecisionUsage) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DecisionUsage) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DecisionUsage) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *DecisionUsage) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *DecisionUsage) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetInputTokens

`func (o *DecisionUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *DecisionUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *DecisionUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *DecisionUsage) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *DecisionUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *DecisionUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *DecisionUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *DecisionUsage) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


