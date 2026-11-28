# AgentRecommenderPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | **NullableString** |  | 
**ApiKeyLast4** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**Enabled** | **bool** |  | 
**HasApiKey** | **bool** |  | 
**Id** | **string** |  | 
**InputPricePerMillion** | **float32** |  | 
**Model** | **string** |  | 
**OutputPricePerMillion** | **float32** |  | 
**Provider** | **string** |  | 
**UpdatedAt** | **NullableTime** |  | 

## Methods

### NewAgentRecommenderPublic

`func NewAgentRecommenderPublic(apiBase NullableString, apiKeyLast4 NullableString, createdAt time.Time, enabled bool, hasApiKey bool, id string, inputPricePerMillion float32, model string, outputPricePerMillion float32, provider string, updatedAt NullableTime, ) *AgentRecommenderPublic`

NewAgentRecommenderPublic instantiates a new AgentRecommenderPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRecommenderPublicWithDefaults

`func NewAgentRecommenderPublicWithDefaults() *AgentRecommenderPublic`

NewAgentRecommenderPublicWithDefaults instantiates a new AgentRecommenderPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *AgentRecommenderPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *AgentRecommenderPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *AgentRecommenderPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.


### SetApiBaseNil

`func (o *AgentRecommenderPublic) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *AgentRecommenderPublic) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKeyLast4

`func (o *AgentRecommenderPublic) GetApiKeyLast4() string`

GetApiKeyLast4 returns the ApiKeyLast4 field if non-nil, zero value otherwise.

### GetApiKeyLast4Ok

`func (o *AgentRecommenderPublic) GetApiKeyLast4Ok() (*string, bool)`

GetApiKeyLast4Ok returns a tuple with the ApiKeyLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyLast4

`func (o *AgentRecommenderPublic) SetApiKeyLast4(v string)`

SetApiKeyLast4 sets ApiKeyLast4 field to given value.


### SetApiKeyLast4Nil

`func (o *AgentRecommenderPublic) SetApiKeyLast4Nil(b bool)`

 SetApiKeyLast4Nil sets the value for ApiKeyLast4 to be an explicit nil

### UnsetApiKeyLast4
`func (o *AgentRecommenderPublic) UnsetApiKeyLast4()`

UnsetApiKeyLast4 ensures that no value is present for ApiKeyLast4, not even an explicit nil
### GetCreatedAt

`func (o *AgentRecommenderPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentRecommenderPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentRecommenderPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *AgentRecommenderPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentRecommenderPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentRecommenderPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasApiKey

`func (o *AgentRecommenderPublic) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *AgentRecommenderPublic) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *AgentRecommenderPublic) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.


### GetId

`func (o *AgentRecommenderPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRecommenderPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRecommenderPublic) SetId(v string)`

SetId sets Id field to given value.


### GetInputPricePerMillion

`func (o *AgentRecommenderPublic) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *AgentRecommenderPublic) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *AgentRecommenderPublic) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetModel

`func (o *AgentRecommenderPublic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentRecommenderPublic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentRecommenderPublic) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputPricePerMillion

`func (o *AgentRecommenderPublic) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *AgentRecommenderPublic) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *AgentRecommenderPublic) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.


### GetProvider

`func (o *AgentRecommenderPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AgentRecommenderPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AgentRecommenderPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUpdatedAt

`func (o *AgentRecommenderPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentRecommenderPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentRecommenderPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *AgentRecommenderPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AgentRecommenderPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


