# OrganizationGuardrailDefinitionTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explanation** | **NullableString** | The vendor&#39;s reason, when it gives one | 
**Score** | **NullableFloat32** | The vendor&#39;s score, when it gives one | 
**Valid** | **bool** | False when the guardrail flagged the text | 

## Methods

### NewOrganizationGuardrailDefinitionTestResult

`func NewOrganizationGuardrailDefinitionTestResult(explanation NullableString, score NullableFloat32, valid bool, ) *OrganizationGuardrailDefinitionTestResult`

NewOrganizationGuardrailDefinitionTestResult instantiates a new OrganizationGuardrailDefinitionTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailDefinitionTestResultWithDefaults

`func NewOrganizationGuardrailDefinitionTestResultWithDefaults() *OrganizationGuardrailDefinitionTestResult`

NewOrganizationGuardrailDefinitionTestResultWithDefaults instantiates a new OrganizationGuardrailDefinitionTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplanation

`func (o *OrganizationGuardrailDefinitionTestResult) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *OrganizationGuardrailDefinitionTestResult) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *OrganizationGuardrailDefinitionTestResult) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### SetExplanationNil

`func (o *OrganizationGuardrailDefinitionTestResult) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *OrganizationGuardrailDefinitionTestResult) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetScore

`func (o *OrganizationGuardrailDefinitionTestResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OrganizationGuardrailDefinitionTestResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OrganizationGuardrailDefinitionTestResult) SetScore(v float32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *OrganizationGuardrailDefinitionTestResult) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *OrganizationGuardrailDefinitionTestResult) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetValid

`func (o *OrganizationGuardrailDefinitionTestResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *OrganizationGuardrailDefinitionTestResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *OrganizationGuardrailDefinitionTestResult) SetValid(v bool)`

SetValid sets Valid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


