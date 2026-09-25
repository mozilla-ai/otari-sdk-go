# OrganizationGuardrailTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explanation** | **NullableString** | The service&#39;s reason, when it gives one | 
**Score** | **NullableFloat32** | The service&#39;s score, when it gives one | 
**Valid** | **NullableBool** | False when the guardrail flagged the text, null when it gave no verdict | 

## Methods

### NewOrganizationGuardrailTestResult

`func NewOrganizationGuardrailTestResult(explanation NullableString, score NullableFloat32, valid NullableBool, ) *OrganizationGuardrailTestResult`

NewOrganizationGuardrailTestResult instantiates a new OrganizationGuardrailTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailTestResultWithDefaults

`func NewOrganizationGuardrailTestResultWithDefaults() *OrganizationGuardrailTestResult`

NewOrganizationGuardrailTestResultWithDefaults instantiates a new OrganizationGuardrailTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplanation

`func (o *OrganizationGuardrailTestResult) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *OrganizationGuardrailTestResult) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *OrganizationGuardrailTestResult) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### SetExplanationNil

`func (o *OrganizationGuardrailTestResult) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *OrganizationGuardrailTestResult) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil
### GetScore

`func (o *OrganizationGuardrailTestResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OrganizationGuardrailTestResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OrganizationGuardrailTestResult) SetScore(v float32)`

SetScore sets Score field to given value.


### SetScoreNil

`func (o *OrganizationGuardrailTestResult) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *OrganizationGuardrailTestResult) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetValid

`func (o *OrganizationGuardrailTestResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *OrganizationGuardrailTestResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *OrganizationGuardrailTestResult) SetValid(v bool)`

SetValid sets Valid field to given value.


### SetValidNil

`func (o *OrganizationGuardrailTestResult) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *OrganizationGuardrailTestResult) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


