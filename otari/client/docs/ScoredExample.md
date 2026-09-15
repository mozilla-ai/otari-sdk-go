# ScoredExample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSource** | Pointer to **string** | Provenance of the scores: &#39;human&#39; or &#39;judge&#39;. | [optional] [default to "human"]
**Prompt** | **string** | The prompt that was tried. | 
**Scores** | **map[string]float32** | Selector -&gt; quality in [0.0, 1.0], where 1.0 is a great answer. Ties are fine and meaningful: two models that both answered well is exactly the case where the router should take the cheaper one. | 
**TaskId** | Pointer to **NullableString** | Partition this example belongs to, matching the Otari-Router-Task header requests send. Omit to file it in the user&#39;s default pool. | [optional] 

## Methods

### NewScoredExample

`func NewScoredExample(prompt string, scores map[string]float32, ) *ScoredExample`

NewScoredExample instantiates a new ScoredExample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoredExampleWithDefaults

`func NewScoredExampleWithDefaults() *ScoredExample`

NewScoredExampleWithDefaults instantiates a new ScoredExample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSource

`func (o *ScoredExample) GetLabelSource() string`

GetLabelSource returns the LabelSource field if non-nil, zero value otherwise.

### GetLabelSourceOk

`func (o *ScoredExample) GetLabelSourceOk() (*string, bool)`

GetLabelSourceOk returns a tuple with the LabelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSource

`func (o *ScoredExample) SetLabelSource(v string)`

SetLabelSource sets LabelSource field to given value.

### HasLabelSource

`func (o *ScoredExample) HasLabelSource() bool`

HasLabelSource returns a boolean if a field has been set.

### GetPrompt

`func (o *ScoredExample) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ScoredExample) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ScoredExample) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetScores

`func (o *ScoredExample) GetScores() map[string]float32`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ScoredExample) GetScoresOk() (*map[string]float32, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ScoredExample) SetScores(v map[string]float32)`

SetScores sets Scores field to given value.


### GetTaskId

`func (o *ScoredExample) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ScoredExample) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ScoredExample) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ScoredExample) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *ScoredExample) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *ScoredExample) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


