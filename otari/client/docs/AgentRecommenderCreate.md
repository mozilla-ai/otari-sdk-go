# AgentRecommenderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** | Root URL of the provider; blank for its own. | [optional] 
**ApiKey** | Pointer to **NullableString** | Stored encrypted, never returned. | [optional] 
**InputPricePerMillion** | Pointer to **float32** |  | [optional] [default to 0.0]
**Model** | **string** | The decision model, as the provider names it. | 
**OutputPricePerMillion** | Pointer to **float32** |  | [optional] [default to 0.0]
**Provider** | **string** | The decision provider: typesafe, openrouter or llamacpp. | 

## Methods

### NewAgentRecommenderCreate

`func NewAgentRecommenderCreate(model string, provider string, ) *AgentRecommenderCreate`

NewAgentRecommenderCreate instantiates a new AgentRecommenderCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRecommenderCreateWithDefaults

`func NewAgentRecommenderCreateWithDefaults() *AgentRecommenderCreate`

NewAgentRecommenderCreateWithDefaults instantiates a new AgentRecommenderCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *AgentRecommenderCreate) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *AgentRecommenderCreate) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *AgentRecommenderCreate) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *AgentRecommenderCreate) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *AgentRecommenderCreate) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *AgentRecommenderCreate) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *AgentRecommenderCreate) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AgentRecommenderCreate) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AgentRecommenderCreate) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AgentRecommenderCreate) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *AgentRecommenderCreate) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *AgentRecommenderCreate) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetInputPricePerMillion

`func (o *AgentRecommenderCreate) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *AgentRecommenderCreate) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *AgentRecommenderCreate) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *AgentRecommenderCreate) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### GetModel

`func (o *AgentRecommenderCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentRecommenderCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentRecommenderCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputPricePerMillion

`func (o *AgentRecommenderCreate) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *AgentRecommenderCreate) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *AgentRecommenderCreate) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *AgentRecommenderCreate) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### GetProvider

`func (o *AgentRecommenderCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentRecommenderCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentRecommenderCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


