# QuestionsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**[]CriteriaInner**](CriteriaInner.md) | Level descriptions, lowest first | 
**Instructions** | [**Instructions1**](Instructions1.md) |  | 
**Type** | **string** |  | 

## Methods

### NewQuestionsValue

`func NewQuestionsValue(criteria []CriteriaInner, instructions Instructions1, type_ string, ) *QuestionsValue`

NewQuestionsValue instantiates a new QuestionsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionsValueWithDefaults

`func NewQuestionsValueWithDefaults() *QuestionsValue`

NewQuestionsValueWithDefaults instantiates a new QuestionsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *QuestionsValue) GetCriteria() []CriteriaInner`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *QuestionsValue) GetCriteriaOk() (*[]CriteriaInner, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *QuestionsValue) SetCriteria(v []CriteriaInner)`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *QuestionsValue) GetInstructions() Instructions1`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *QuestionsValue) GetInstructionsOk() (*Instructions1, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *QuestionsValue) SetInstructions(v Instructions1)`

SetInstructions sets Instructions field to given value.


### GetType

`func (o *QuestionsValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuestionsValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuestionsValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


