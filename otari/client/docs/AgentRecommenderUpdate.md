# AgentRecommenderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**InputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**OutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentRecommenderUpdate

`func NewAgentRecommenderUpdate() *AgentRecommenderUpdate`

NewAgentRecommenderUpdate instantiates a new AgentRecommenderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRecommenderUpdateWithDefaults

`func NewAgentRecommenderUpdateWithDefaults() *AgentRecommenderUpdate`

NewAgentRecommenderUpdateWithDefaults instantiates a new AgentRecommenderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *AgentRecommenderUpdate) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *AgentRecommenderUpdate) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *AgentRecommenderUpdate) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *AgentRecommenderUpdate) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *AgentRecommenderUpdate) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *AgentRecommenderUpdate) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *AgentRecommenderUpdate) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AgentRecommenderUpdate) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AgentRecommenderUpdate) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AgentRecommenderUpdate) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *AgentRecommenderUpdate) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *AgentRecommenderUpdate) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetEnabled

`func (o *AgentRecommenderUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentRecommenderUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentRecommenderUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentRecommenderUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AgentRecommenderUpdate) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AgentRecommenderUpdate) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInputPricePerMillion

`func (o *AgentRecommenderUpdate) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *AgentRecommenderUpdate) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *AgentRecommenderUpdate) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *AgentRecommenderUpdate) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### SetInputPricePerMillionNil

`func (o *AgentRecommenderUpdate) SetInputPricePerMillionNil(b bool)`

 SetInputPricePerMillionNil sets the value for InputPricePerMillion to be an explicit nil

### UnsetInputPricePerMillion
`func (o *AgentRecommenderUpdate) UnsetInputPricePerMillion()`

UnsetInputPricePerMillion ensures that no value is present for InputPricePerMillion, not even an explicit nil
### GetModel

`func (o *AgentRecommenderUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentRecommenderUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentRecommenderUpdate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentRecommenderUpdate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *AgentRecommenderUpdate) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AgentRecommenderUpdate) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetOutputPricePerMillion

`func (o *AgentRecommenderUpdate) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *AgentRecommenderUpdate) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *AgentRecommenderUpdate) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *AgentRecommenderUpdate) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### SetOutputPricePerMillionNil

`func (o *AgentRecommenderUpdate) SetOutputPricePerMillionNil(b bool)`

 SetOutputPricePerMillionNil sets the value for OutputPricePerMillion to be an explicit nil

### UnsetOutputPricePerMillion
`func (o *AgentRecommenderUpdate) UnsetOutputPricePerMillion()`

UnsetOutputPricePerMillion ensures that no value is present for OutputPricePerMillion, not even an explicit nil
### GetProvider

`func (o *AgentRecommenderUpdate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentRecommenderUpdate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentRecommenderUpdate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AgentRecommenderUpdate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AgentRecommenderUpdate) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AgentRecommenderUpdate) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


