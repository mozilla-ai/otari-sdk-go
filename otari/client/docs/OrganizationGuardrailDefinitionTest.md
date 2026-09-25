# OrganizationGuardrailDefinitionTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The input to check, as a request&#39;s user text would reach it | 
**ValidateKwargs** | Pointer to **map[string]interface{}** | Per-check arguments, as a mandate&#39;s validate_kwargs would hand them to this guardrail | [optional] 

## Methods

### NewOrganizationGuardrailDefinitionTest

`func NewOrganizationGuardrailDefinitionTest(text string, ) *OrganizationGuardrailDefinitionTest`

NewOrganizationGuardrailDefinitionTest instantiates a new OrganizationGuardrailDefinitionTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailDefinitionTestWithDefaults

`func NewOrganizationGuardrailDefinitionTestWithDefaults() *OrganizationGuardrailDefinitionTest`

NewOrganizationGuardrailDefinitionTestWithDefaults instantiates a new OrganizationGuardrailDefinitionTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *OrganizationGuardrailDefinitionTest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OrganizationGuardrailDefinitionTest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OrganizationGuardrailDefinitionTest) SetText(v string)`

SetText sets Text field to given value.


### GetValidateKwargs

`func (o *OrganizationGuardrailDefinitionTest) GetValidateKwargs() map[string]interface{}`

GetValidateKwargs returns the ValidateKwargs field if non-nil, zero value otherwise.

### GetValidateKwargsOk

`func (o *OrganizationGuardrailDefinitionTest) GetValidateKwargsOk() (*map[string]interface{}, bool)`

GetValidateKwargsOk returns a tuple with the ValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateKwargs

`func (o *OrganizationGuardrailDefinitionTest) SetValidateKwargs(v map[string]interface{})`

SetValidateKwargs sets ValidateKwargs field to given value.

### HasValidateKwargs

`func (o *OrganizationGuardrailDefinitionTest) HasValidateKwargs() bool`

HasValidateKwargs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


