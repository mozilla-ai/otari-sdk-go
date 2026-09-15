# GuardrailParameterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to **[]string** | Allowed values for an enum parameter | [optional] 
**Default** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** | One-line help text from the guardrail&#39;s docstring | [optional] 
**EnvVar** | Pointer to **NullableString** | The environment variable that supplies this parameter when no value is stored, so a form can offer that instead of demanding a credential the deployment already has | [optional] 
**Name** | **string** | The keyword argument&#39;s name, as it is sent in validate_kwargs | 
**Required** | **bool** | Whether a value must be supplied for the guardrail to run. Folds together the signature having no default and upstream&#39;s effectively-required flag, which covers a parameter that defaults to a value the guardrail then refuses to run without | 
**Secret** | Pointer to **bool** | Whether the value is a credential, so a form masks it and never echoes it back | [optional] [default to false]
**Storable** | Pointer to **bool** | Whether a saved value can stand in for this parameter. False for a secret whose type is json, which upstream uses for a live object (an authenticated SDK client or session) that cannot be written down. A form offers no field for one | [optional] [default to true]
**Type** | **string** | Value shape, so a form can render the matching control | 

## Methods

### NewGuardrailParameterSpec

`func NewGuardrailParameterSpec(name string, required bool, type_ string, ) *GuardrailParameterSpec`

NewGuardrailParameterSpec instantiates a new GuardrailParameterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailParameterSpecWithDefaults

`func NewGuardrailParameterSpecWithDefaults() *GuardrailParameterSpec`

NewGuardrailParameterSpecWithDefaults instantiates a new GuardrailParameterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *GuardrailParameterSpec) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *GuardrailParameterSpec) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *GuardrailParameterSpec) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *GuardrailParameterSpec) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *GuardrailParameterSpec) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *GuardrailParameterSpec) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetDefault

`func (o *GuardrailParameterSpec) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GuardrailParameterSpec) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GuardrailParameterSpec) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *GuardrailParameterSpec) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *GuardrailParameterSpec) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *GuardrailParameterSpec) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDescription

`func (o *GuardrailParameterSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuardrailParameterSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuardrailParameterSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuardrailParameterSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuardrailParameterSpec) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuardrailParameterSpec) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvVar

`func (o *GuardrailParameterSpec) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *GuardrailParameterSpec) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *GuardrailParameterSpec) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *GuardrailParameterSpec) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### SetEnvVarNil

`func (o *GuardrailParameterSpec) SetEnvVarNil(b bool)`

 SetEnvVarNil sets the value for EnvVar to be an explicit nil

### UnsetEnvVar
`func (o *GuardrailParameterSpec) UnsetEnvVar()`

UnsetEnvVar ensures that no value is present for EnvVar, not even an explicit nil
### GetName

`func (o *GuardrailParameterSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuardrailParameterSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuardrailParameterSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *GuardrailParameterSpec) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GuardrailParameterSpec) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GuardrailParameterSpec) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSecret

`func (o *GuardrailParameterSpec) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GuardrailParameterSpec) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GuardrailParameterSpec) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GuardrailParameterSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStorable

`func (o *GuardrailParameterSpec) GetStorable() bool`

GetStorable returns the Storable field if non-nil, zero value otherwise.

### GetStorableOk

`func (o *GuardrailParameterSpec) GetStorableOk() (*bool, bool)`

GetStorableOk returns a tuple with the Storable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorable

`func (o *GuardrailParameterSpec) SetStorable(v bool)`

SetStorable sets Storable field to given value.

### HasStorable

`func (o *GuardrailParameterSpec) HasStorable() bool`

HasStorable returns a boolean if a field has been set.

### GetType

`func (o *GuardrailParameterSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuardrailParameterSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuardrailParameterSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


