# GuardrailProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guardrail** | **string** | The any-guardrail class the profile is built from | 
**ModelId** | Pointer to **NullableString** | The model the operator pinned, when they pinned one | [optional] 
**Parameters** | Pointer to [**[]GuardrailParameterSpec**](GuardrailParameterSpec.md) | The validate_kwargs this profile accepts | [optional] 
**ParametersKnown** | **bool** | False when this gateway&#39;s any-guardrail is older than the service&#39;s and has no schema for that class. The profile is still selectable; only its typed fields are missing | 
**Profile** | **string** | The name a guardrail entry puts in its profile field | 

## Methods

### NewGuardrailProfileSpec

`func NewGuardrailProfileSpec(guardrail string, parametersKnown bool, profile string, ) *GuardrailProfileSpec`

NewGuardrailProfileSpec instantiates a new GuardrailProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailProfileSpecWithDefaults

`func NewGuardrailProfileSpecWithDefaults() *GuardrailProfileSpec`

NewGuardrailProfileSpecWithDefaults instantiates a new GuardrailProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuardrail

`func (o *GuardrailProfileSpec) GetGuardrail() string`

GetGuardrail returns the Guardrail field if non-nil, zero value otherwise.

### GetGuardrailOk

`func (o *GuardrailProfileSpec) GetGuardrailOk() (*string, bool)`

GetGuardrailOk returns a tuple with the Guardrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrail

`func (o *GuardrailProfileSpec) SetGuardrail(v string)`

SetGuardrail sets Guardrail field to given value.


### GetModelId

`func (o *GuardrailProfileSpec) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *GuardrailProfileSpec) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *GuardrailProfileSpec) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *GuardrailProfileSpec) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *GuardrailProfileSpec) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *GuardrailProfileSpec) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetParameters

`func (o *GuardrailProfileSpec) GetParameters() []GuardrailParameterSpec`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GuardrailProfileSpec) GetParametersOk() (*[]GuardrailParameterSpec, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GuardrailProfileSpec) SetParameters(v []GuardrailParameterSpec)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GuardrailProfileSpec) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetParametersKnown

`func (o *GuardrailProfileSpec) GetParametersKnown() bool`

GetParametersKnown returns the ParametersKnown field if non-nil, zero value otherwise.

### GetParametersKnownOk

`func (o *GuardrailProfileSpec) GetParametersKnownOk() (*bool, bool)`

GetParametersKnownOk returns a tuple with the ParametersKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParametersKnown

`func (o *GuardrailProfileSpec) SetParametersKnown(v bool)`

SetParametersKnown sets ParametersKnown field to given value.


### GetProfile

`func (o *GuardrailProfileSpec) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GuardrailProfileSpec) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GuardrailProfileSpec) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


