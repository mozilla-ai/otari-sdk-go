# RequirementGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**EnvVars** | Pointer to **[]string** |  | [optional] [default to {}]
**Parameters** | **[]string** |  | 

## Methods

### NewRequirementGroup

`func NewRequirementGroup(description string, parameters []string, ) *RequirementGroup`

NewRequirementGroup instantiates a new RequirementGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementGroupWithDefaults

`func NewRequirementGroupWithDefaults() *RequirementGroup`

NewRequirementGroupWithDefaults instantiates a new RequirementGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RequirementGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequirementGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequirementGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvVars

`func (o *RequirementGroup) GetEnvVars() []string`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *RequirementGroup) GetEnvVarsOk() (*[]string, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *RequirementGroup) SetEnvVars(v []string)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *RequirementGroup) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetParameters

`func (o *RequirementGroup) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RequirementGroup) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RequirementGroup) SetParameters(v []string)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


