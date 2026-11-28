# CodeExecutionEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | Whether a run that names no environment runs in this one. | 
**Description** | **NullableString** | What it contains and how big it is, when the operator said. | 
**Label** | **string** | A display name; the name itself when the operator gave none. | 
**Name** | **string** | What POST /api/v1/code/execute&#39;s environment field takes. | 
**Object** | Pointer to **string** |  | [optional] [default to "code_execution.environment"]
**PricePerRun** | **NullableFloat32** | USD per run for the caller&#39;s organization: the environment&#39;s own price, else code execution&#39;s. Null when unpriced. | 

## Methods

### NewCodeExecutionEnvironment

`func NewCodeExecutionEnvironment(default_ bool, description NullableString, label string, name string, pricePerRun NullableFloat32, ) *CodeExecutionEnvironment`

NewCodeExecutionEnvironment instantiates a new CodeExecutionEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeExecutionEnvironmentWithDefaults

`func NewCodeExecutionEnvironmentWithDefaults() *CodeExecutionEnvironment`

NewCodeExecutionEnvironmentWithDefaults instantiates a new CodeExecutionEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *CodeExecutionEnvironment) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CodeExecutionEnvironment) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CodeExecutionEnvironment) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetDescription

`func (o *CodeExecutionEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CodeExecutionEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CodeExecutionEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CodeExecutionEnvironment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CodeExecutionEnvironment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabel

`func (o *CodeExecutionEnvironment) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CodeExecutionEnvironment) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CodeExecutionEnvironment) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CodeExecutionEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeExecutionEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeExecutionEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetObject

`func (o *CodeExecutionEnvironment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CodeExecutionEnvironment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CodeExecutionEnvironment) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CodeExecutionEnvironment) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPricePerRun

`func (o *CodeExecutionEnvironment) GetPricePerRun() float32`

GetPricePerRun returns the PricePerRun field if non-nil, zero value otherwise.

### GetPricePerRunOk

`func (o *CodeExecutionEnvironment) GetPricePerRunOk() (*float32, bool)`

GetPricePerRunOk returns a tuple with the PricePerRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerRun

`func (o *CodeExecutionEnvironment) SetPricePerRun(v float32)`

SetPricePerRun sets PricePerRun field to given value.


### SetPricePerRunNil

`func (o *CodeExecutionEnvironment) SetPricePerRunNil(b bool)`

 SetPricePerRunNil sets the value for PricePerRun to be an explicit nil

### UnsetPricePerRun
`func (o *CodeExecutionEnvironment) UnsetPricePerRun()`

UnsetPricePerRun ensures that no value is present for PricePerRun, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


