# OrganizationGuardrailDefinitionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateKwargs** | Pointer to **map[string]interface{}** | Constructor arguments for the guardrail. Every argument the catalog marks secret is encrypted at rest and never returned; the rest are stored and returned as sent | [optional] 
**Enabled** | Pointer to **bool** | False stops the guardrail everywhere it is mandated | [optional] [default to true]
**GuardrailName** | **string** | The any-guardrail class to build, as the built-in guardrail catalog names it | 
**Name** | **string** | The organization&#39;s own label for this definition, unique within the organization | 

## Methods

### NewOrganizationGuardrailDefinitionCreate

`func NewOrganizationGuardrailDefinitionCreate(guardrailName string, name string, ) *OrganizationGuardrailDefinitionCreate`

NewOrganizationGuardrailDefinitionCreate instantiates a new OrganizationGuardrailDefinitionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailDefinitionCreateWithDefaults

`func NewOrganizationGuardrailDefinitionCreateWithDefaults() *OrganizationGuardrailDefinitionCreate`

NewOrganizationGuardrailDefinitionCreateWithDefaults instantiates a new OrganizationGuardrailDefinitionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateKwargs

`func (o *OrganizationGuardrailDefinitionCreate) GetCreateKwargs() map[string]interface{}`

GetCreateKwargs returns the CreateKwargs field if non-nil, zero value otherwise.

### GetCreateKwargsOk

`func (o *OrganizationGuardrailDefinitionCreate) GetCreateKwargsOk() (*map[string]interface{}, bool)`

GetCreateKwargsOk returns a tuple with the CreateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateKwargs

`func (o *OrganizationGuardrailDefinitionCreate) SetCreateKwargs(v map[string]interface{})`

SetCreateKwargs sets CreateKwargs field to given value.

### HasCreateKwargs

`func (o *OrganizationGuardrailDefinitionCreate) HasCreateKwargs() bool`

HasCreateKwargs returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationGuardrailDefinitionCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationGuardrailDefinitionCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationGuardrailDefinitionCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationGuardrailDefinitionCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGuardrailName

`func (o *OrganizationGuardrailDefinitionCreate) GetGuardrailName() string`

GetGuardrailName returns the GuardrailName field if non-nil, zero value otherwise.

### GetGuardrailNameOk

`func (o *OrganizationGuardrailDefinitionCreate) GetGuardrailNameOk() (*string, bool)`

GetGuardrailNameOk returns a tuple with the GuardrailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailName

`func (o *OrganizationGuardrailDefinitionCreate) SetGuardrailName(v string)`

SetGuardrailName sets GuardrailName field to given value.


### GetName

`func (o *OrganizationGuardrailDefinitionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationGuardrailDefinitionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationGuardrailDefinitionCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


