# DecisionAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | Pointer to **NullableString** | The chosen option, for a choice question | [optional] 
**Confidence** | Pointer to **NullableFloat32** |  | [optional] 
**Legend** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Noul** | Pointer to **NullableFloat32** | Probability of yes, for a noul question | [optional] 
**Probabilities** | Pointer to **map[string]float32** |  | [optional] 
**Score** | Pointer to **NullableFloat32** | The level, for a score question | [optional] 
**Type** | **string** |  | 

## Methods

### NewDecisionAnswer

`func NewDecisionAnswer(type_ string, ) *DecisionAnswer`

NewDecisionAnswer instantiates a new DecisionAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionAnswerWithDefaults

`func NewDecisionAnswerWithDefaults() *DecisionAnswer`

NewDecisionAnswerWithDefaults instantiates a new DecisionAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *DecisionAnswer) GetChoice() string`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *DecisionAnswer) GetChoiceOk() (*string, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *DecisionAnswer) SetChoice(v string)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *DecisionAnswer) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### SetChoiceNil

`func (o *DecisionAnswer) SetChoiceNil(b bool)`

 SetChoiceNil sets the value for Choice to be an explicit nil

### UnsetChoice
`func (o *DecisionAnswer) UnsetChoice()`

UnsetChoice ensures that no value is present for Choice, not even an explicit nil
### GetConfidence

`func (o *DecisionAnswer) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DecisionAnswer) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DecisionAnswer) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *DecisionAnswer) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *DecisionAnswer) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *DecisionAnswer) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetLegend

`func (o *DecisionAnswer) GetLegend() map[string]interface{}`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *DecisionAnswer) GetLegendOk() (*map[string]interface{}, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *DecisionAnswer) SetLegend(v map[string]interface{})`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *DecisionAnswer) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### SetLegendNil

`func (o *DecisionAnswer) SetLegendNil(b bool)`

 SetLegendNil sets the value for Legend to be an explicit nil

### UnsetLegend
`func (o *DecisionAnswer) UnsetLegend()`

UnsetLegend ensures that no value is present for Legend, not even an explicit nil
### GetNoul

`func (o *DecisionAnswer) GetNoul() float32`

GetNoul returns the Noul field if non-nil, zero value otherwise.

### GetNoulOk

`func (o *DecisionAnswer) GetNoulOk() (*float32, bool)`

GetNoulOk returns a tuple with the Noul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoul

`func (o *DecisionAnswer) SetNoul(v float32)`

SetNoul sets Noul field to given value.

### HasNoul

`func (o *DecisionAnswer) HasNoul() bool`

HasNoul returns a boolean if a field has been set.

### SetNoulNil

`func (o *DecisionAnswer) SetNoulNil(b bool)`

 SetNoulNil sets the value for Noul to be an explicit nil

### UnsetNoul
`func (o *DecisionAnswer) UnsetNoul()`

UnsetNoul ensures that no value is present for Noul, not even an explicit nil
### GetProbabilities

`func (o *DecisionAnswer) GetProbabilities() map[string]float32`

GetProbabilities returns the Probabilities field if non-nil, zero value otherwise.

### GetProbabilitiesOk

`func (o *DecisionAnswer) GetProbabilitiesOk() (*map[string]float32, bool)`

GetProbabilitiesOk returns a tuple with the Probabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilities

`func (o *DecisionAnswer) SetProbabilities(v map[string]float32)`

SetProbabilities sets Probabilities field to given value.

### HasProbabilities

`func (o *DecisionAnswer) HasProbabilities() bool`

HasProbabilities returns a boolean if a field has been set.

### SetProbabilitiesNil

`func (o *DecisionAnswer) SetProbabilitiesNil(b bool)`

 SetProbabilitiesNil sets the value for Probabilities to be an explicit nil

### UnsetProbabilities
`func (o *DecisionAnswer) UnsetProbabilities()`

UnsetProbabilities ensures that no value is present for Probabilities, not even an explicit nil
### GetScore

`func (o *DecisionAnswer) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DecisionAnswer) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DecisionAnswer) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DecisionAnswer) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *DecisionAnswer) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *DecisionAnswer) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetType

`func (o *DecisionAnswer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DecisionAnswer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DecisionAnswer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


