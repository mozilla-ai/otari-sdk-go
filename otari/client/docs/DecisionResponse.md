# DecisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | [**map[string]DecisionAnswer**](DecisionAnswer.md) |  | 
**Model** | **string** |  | 
**Usage** | Pointer to [**NullableDecisionUsage**](DecisionUsage.md) |  | [optional] 

## Methods

### NewDecisionResponse

`func NewDecisionResponse(answers map[string]DecisionAnswer, model string, ) *DecisionResponse`

NewDecisionResponse instantiates a new DecisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionResponseWithDefaults

`func NewDecisionResponseWithDefaults() *DecisionResponse`

NewDecisionResponseWithDefaults instantiates a new DecisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *DecisionResponse) GetAnswers() map[string]DecisionAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *DecisionResponse) GetAnswersOk() (*map[string]DecisionAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *DecisionResponse) SetAnswers(v map[string]DecisionAnswer)`

SetAnswers sets Answers field to given value.


### GetModel

`func (o *DecisionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetUsage

`func (o *DecisionResponse) GetUsage() DecisionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DecisionResponse) GetUsageOk() (*DecisionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DecisionResponse) SetUsage(v DecisionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DecisionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *DecisionResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *DecisionResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


