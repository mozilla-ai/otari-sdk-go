# NoulQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**NullableNoulCriteria**](NoulCriteria.md) |  | [optional] 
**Instructions** | [**Instructions1**](Instructions1.md) |  | 
**Type** | **string** |  | 

## Methods

### NewNoulQuestion

`func NewNoulQuestion(instructions Instructions1, type_ string, ) *NoulQuestion`

NewNoulQuestion instantiates a new NoulQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoulQuestionWithDefaults

`func NewNoulQuestionWithDefaults() *NoulQuestion`

NewNoulQuestionWithDefaults instantiates a new NoulQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *NoulQuestion) GetCriteria() NoulCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *NoulQuestion) GetCriteriaOk() (*NoulCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *NoulQuestion) SetCriteria(v NoulCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *NoulQuestion) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *NoulQuestion) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *NoulQuestion) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil
### GetInstructions

`func (o *NoulQuestion) GetInstructions() Instructions1`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *NoulQuestion) GetInstructionsOk() (*Instructions1, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *NoulQuestion) SetInstructions(v Instructions1)`

SetInstructions sets Instructions field to given value.


### GetType

`func (o *NoulQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoulQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoulQuestion) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


