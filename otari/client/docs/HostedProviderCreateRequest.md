# HostedProviderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Name** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** |  | 
**RestrictedToOrganizationId** | Pointer to **NullableString** |  | [optional] 
**ServesUnlistedModels** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewHostedProviderCreateRequest

`func NewHostedProviderCreateRequest(provider string, ) *HostedProviderCreateRequest`

NewHostedProviderCreateRequest instantiates a new HostedProviderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedProviderCreateRequestWithDefaults

`func NewHostedProviderCreateRequestWithDefaults() *HostedProviderCreateRequest`

NewHostedProviderCreateRequestWithDefaults instantiates a new HostedProviderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *HostedProviderCreateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *HostedProviderCreateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *HostedProviderCreateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *HostedProviderCreateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *HostedProviderCreateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *HostedProviderCreateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *HostedProviderCreateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HostedProviderCreateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HostedProviderCreateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HostedProviderCreateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *HostedProviderCreateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HostedProviderCreateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetClientArgs

`func (o *HostedProviderCreateRequest) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *HostedProviderCreateRequest) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *HostedProviderCreateRequest) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *HostedProviderCreateRequest) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *HostedProviderCreateRequest) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *HostedProviderCreateRequest) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetEnabled

`func (o *HostedProviderCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedProviderCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedProviderCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HostedProviderCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *HostedProviderCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostedProviderCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostedProviderCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostedProviderCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HostedProviderCreateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HostedProviderCreateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvider

`func (o *HostedProviderCreateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HostedProviderCreateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HostedProviderCreateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRestrictedToOrganizationId

`func (o *HostedProviderCreateRequest) GetRestrictedToOrganizationId() string`

GetRestrictedToOrganizationId returns the RestrictedToOrganizationId field if non-nil, zero value otherwise.

### GetRestrictedToOrganizationIdOk

`func (o *HostedProviderCreateRequest) GetRestrictedToOrganizationIdOk() (*string, bool)`

GetRestrictedToOrganizationIdOk returns a tuple with the RestrictedToOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToOrganizationId

`func (o *HostedProviderCreateRequest) SetRestrictedToOrganizationId(v string)`

SetRestrictedToOrganizationId sets RestrictedToOrganizationId field to given value.

### HasRestrictedToOrganizationId

`func (o *HostedProviderCreateRequest) HasRestrictedToOrganizationId() bool`

HasRestrictedToOrganizationId returns a boolean if a field has been set.

### SetRestrictedToOrganizationIdNil

`func (o *HostedProviderCreateRequest) SetRestrictedToOrganizationIdNil(b bool)`

 SetRestrictedToOrganizationIdNil sets the value for RestrictedToOrganizationId to be an explicit nil

### UnsetRestrictedToOrganizationId
`func (o *HostedProviderCreateRequest) UnsetRestrictedToOrganizationId()`

UnsetRestrictedToOrganizationId ensures that no value is present for RestrictedToOrganizationId, not even an explicit nil
### GetServesUnlistedModels

`func (o *HostedProviderCreateRequest) GetServesUnlistedModels() bool`

GetServesUnlistedModels returns the ServesUnlistedModels field if non-nil, zero value otherwise.

### GetServesUnlistedModelsOk

`func (o *HostedProviderCreateRequest) GetServesUnlistedModelsOk() (*bool, bool)`

GetServesUnlistedModelsOk returns a tuple with the ServesUnlistedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServesUnlistedModels

`func (o *HostedProviderCreateRequest) SetServesUnlistedModels(v bool)`

SetServesUnlistedModels sets ServesUnlistedModels field to given value.

### HasServesUnlistedModels

`func (o *HostedProviderCreateRequest) HasServesUnlistedModels() bool`

HasServesUnlistedModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


