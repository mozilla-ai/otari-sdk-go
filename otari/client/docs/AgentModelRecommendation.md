# AgentModelRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | A model alias or id the harness can start the subagent on. | 
**Probabilities** | Pointer to **map[string]float32** | The decision model&#39;s probability for each candidate, when it reports them. | [optional] 
**Reason** | Pointer to **NullableString** | Why, in one short sentence the harness may show. | [optional] 

## Methods

### NewAgentModelRecommendation

`func NewAgentModelRecommendation(model string, ) *AgentModelRecommendation`

NewAgentModelRecommendation instantiates a new AgentModelRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentModelRecommendationWithDefaults

`func NewAgentModelRecommendationWithDefaults() *AgentModelRecommendation`

NewAgentModelRecommendationWithDefaults instantiates a new AgentModelRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AgentModelRecommendation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentModelRecommendation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentModelRecommendation) SetModel(v string)`

SetModel sets Model field to given value.


### GetProbabilities

`func (o *AgentModelRecommendation) GetProbabilities() map[string]float32`

GetProbabilities returns the Probabilities field if non-nil, zero value otherwise.

### GetProbabilitiesOk

`func (o *AgentModelRecommendation) GetProbabilitiesOk() (*map[string]float32, bool)`

GetProbabilitiesOk returns a tuple with the Probabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilities

`func (o *AgentModelRecommendation) SetProbabilities(v map[string]float32)`

SetProbabilities sets Probabilities field to given value.

### HasProbabilities

`func (o *AgentModelRecommendation) HasProbabilities() bool`

HasProbabilities returns a boolean if a field has been set.

### SetProbabilitiesNil

`func (o *AgentModelRecommendation) SetProbabilitiesNil(b bool)`

 SetProbabilitiesNil sets the value for Probabilities to be an explicit nil

### UnsetProbabilities
`func (o *AgentModelRecommendation) UnsetProbabilities()`

UnsetProbabilities ensures that no value is present for Probabilities, not even an explicit nil
### GetReason

`func (o *AgentModelRecommendation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AgentModelRecommendation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AgentModelRecommendation) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AgentModelRecommendation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *AgentModelRecommendation) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AgentModelRecommendation) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


