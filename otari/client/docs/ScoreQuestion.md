# ScoreQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**[]CriteriaInner**](CriteriaInner.md) | Level descriptions, lowest first | 
**Instructions** | [**Instructions1**](Instructions1.md) |  | 
**Type** | **string** |  | 

## Methods

### NewScoreQuestion

`func NewScoreQuestion(criteria []CriteriaInner, instructions Instructions1, type_ string, ) *ScoreQuestion`

NewScoreQuestion instantiates a new ScoreQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreQuestionWithDefaults

`func NewScoreQuestionWithDefaults() *ScoreQuestion`

NewScoreQuestionWithDefaults instantiates a new ScoreQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *ScoreQuestion) GetCriteria() []CriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ScoreQuestion) GetCriteriaOk() (*[]CriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ScoreQuestion) SetCriteria(v []CriteriaInner)`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *ScoreQuestion) GetInstructions() Instructions1`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ScoreQuestion) GetInstructionsOk() (*Instructions1, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ScoreQuestion) SetInstructions(v Instructions1)`

SetInstructions sets Instructions field to given value.


### GetType

`func (o *ScoreQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScoreQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScoreQuestion) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


