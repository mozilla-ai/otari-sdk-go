# ProviderEndpointUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** | Base URL of the endpoint. Refused when it resolves to a private, loopback, link-local or reserved address. | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**DefaultParams** | Pointer to **map[string]interface{}** | Fields added to every request body sent to this endpoint, beneath the caller&#39;s own: a field the caller sets wins. For fields the gateway does not model, such as vLLM&#39;s &#39;chat_template_kwargs&#39;. Credential and transport fields are refused. | [optional] 
**Name** | Pointer to **NullableString** | What callers put before the colon to reach this endpoint, as &#39;&lt;name&gt;:&lt;model&gt;&#39;. Letters, digits, &#39;.&#39;, &#39;_&#39; and &#39;-&#39;, starting with a letter or digit. It may not be a provider&#39;s name or a configured instance&#39;s. | [optional] 
**Provider** | Pointer to **NullableString** | The implementation that speaks to the endpoint: &#39;openai&#39; for an OpenAI-compatible server, &#39;anthropic&#39; for an Anthropic-compatible one. Whichever it is, callers may use any of the chat, responses and messages routes. | [optional] 

## Methods

### NewProviderEndpointUpdateRequest

`func NewProviderEndpointUpdateRequest() *ProviderEndpointUpdateRequest`

NewProviderEndpointUpdateRequest instantiates a new ProviderEndpointUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderEndpointUpdateRequestWithDefaults

`func NewProviderEndpointUpdateRequestWithDefaults() *ProviderEndpointUpdateRequest`

NewProviderEndpointUpdateRequestWithDefaults instantiates a new ProviderEndpointUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ProviderEndpointUpdateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ProviderEndpointUpdateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ProviderEndpointUpdateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *ProviderEndpointUpdateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *ProviderEndpointUpdateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *ProviderEndpointUpdateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *ProviderEndpointUpdateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ProviderEndpointUpdateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ProviderEndpointUpdateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ProviderEndpointUpdateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *ProviderEndpointUpdateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *ProviderEndpointUpdateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetDefaultParams

`func (o *ProviderEndpointUpdateRequest) GetDefaultParams() map[string]interface{}`

GetDefaultParams returns the DefaultParams field if non-nil, zero value otherwise.

### GetDefaultParamsOk

`func (o *ProviderEndpointUpdateRequest) GetDefaultParamsOk() (*map[string]interface{}, bool)`

GetDefaultParamsOk returns a tuple with the DefaultParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParams

`func (o *ProviderEndpointUpdateRequest) SetDefaultParams(v map[string]interface{})`

SetDefaultParams sets DefaultParams field to given value.

### HasDefaultParams

`func (o *ProviderEndpointUpdateRequest) HasDefaultParams() bool`

HasDefaultParams returns a boolean if a field has been set.

### SetDefaultParamsNil

`func (o *ProviderEndpointUpdateRequest) SetDefaultParamsNil(b bool)`

 SetDefaultParamsNil sets the value for DefaultParams to be an explicit nil

### UnsetDefaultParams
`func (o *ProviderEndpointUpdateRequest) UnsetDefaultParams()`

UnsetDefaultParams ensures that no value is present for DefaultParams, not even an explicit nil
### GetName

`func (o *ProviderEndpointUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderEndpointUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderEndpointUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderEndpointUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProviderEndpointUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProviderEndpointUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvider

`func (o *ProviderEndpointUpdateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderEndpointUpdateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderEndpointUpdateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderEndpointUpdateRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ProviderEndpointUpdateRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ProviderEndpointUpdateRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


