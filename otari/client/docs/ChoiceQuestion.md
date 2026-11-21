# ChoiceQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**map[string]CriteriaValue**](CriteriaValue.md) | Option name to what it means | 
**Instructions** | [**Instructions**](Instructions.md) |  | 
**Type** | **string** |  | 

## Methods

### NewChoiceQuestion

`func NewChoiceQuestion(criteria map[string]CriteriaValue, instructions Instructions, type_ string, ) *ChoiceQuestion`

NewChoiceQuestion instantiates a new ChoiceQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceQuestionWithDefaults

`func NewChoiceQuestionWithDefaults() *ChoiceQuestion`

NewChoiceQuestionWithDefaults instantiates a new ChoiceQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *ChoiceQuestion) GetCriteria() map[string]CriteriaValue`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ChoiceQuestion) GetCriteriaOk() (*map[string]CriteriaValue, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ChoiceQuestion) SetCriteria(v map[string]CriteriaValue)`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *ChoiceQuestion) GetInstructions() Instructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ChoiceQuestion) GetInstructionsOk() (*Instructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ChoiceQuestion) SetInstructions(v Instructions)`

SetInstructions sets Instructions field to given value.


### GetType

`func (o *ChoiceQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChoiceQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChoiceQuestion) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


