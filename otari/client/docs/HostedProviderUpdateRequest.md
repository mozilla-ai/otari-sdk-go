# HostedProviderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**RestrictedToOrganizationId** | Pointer to **NullableString** |  | [optional] 
**ServesUnlistedModels** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewHostedProviderUpdateRequest

`func NewHostedProviderUpdateRequest() *HostedProviderUpdateRequest`

NewHostedProviderUpdateRequest instantiates a new HostedProviderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedProviderUpdateRequestWithDefaults

`func NewHostedProviderUpdateRequestWithDefaults() *HostedProviderUpdateRequest`

NewHostedProviderUpdateRequestWithDefaults instantiates a new HostedProviderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *HostedProviderUpdateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *HostedProviderUpdateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *HostedProviderUpdateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *HostedProviderUpdateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *HostedProviderUpdateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *HostedProviderUpdateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *HostedProviderUpdateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HostedProviderUpdateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HostedProviderUpdateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HostedProviderUpdateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *HostedProviderUpdateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HostedProviderUpdateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetClientArgs

`func (o *HostedProviderUpdateRequest) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *HostedProviderUpdateRequest) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *HostedProviderUpdateRequest) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *HostedProviderUpdateRequest) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *HostedProviderUpdateRequest) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *HostedProviderUpdateRequest) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetEnabled

`func (o *HostedProviderUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedProviderUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedProviderUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HostedProviderUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *HostedProviderUpdateRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *HostedProviderUpdateRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRestrictedToOrganizationId

`func (o *HostedProviderUpdateRequest) GetRestrictedToOrganizationId() string`

GetRestrictedToOrganizationId returns the RestrictedToOrganizationId field if non-nil, zero value otherwise.

### GetRestrictedToOrganizationIdOk

`func (o *HostedProviderUpdateRequest) GetRestrictedToOrganizationIdOk() (*string, bool)`

GetRestrictedToOrganizationIdOk returns a tuple with the RestrictedToOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToOrganizationId

`func (o *HostedProviderUpdateRequest) SetRestrictedToOrganizationId(v string)`

SetRestrictedToOrganizationId sets RestrictedToOrganizationId field to given value.

### HasRestrictedToOrganizationId

`func (o *HostedProviderUpdateRequest) HasRestrictedToOrganizationId() bool`

HasRestrictedToOrganizationId returns a boolean if a field has been set.

### SetRestrictedToOrganizationIdNil

`func (o *HostedProviderUpdateRequest) SetRestrictedToOrganizationIdNil(b bool)`

 SetRestrictedToOrganizationIdNil sets the value for RestrictedToOrganizationId to be an explicit nil

### UnsetRestrictedToOrganizationId
`func (o *HostedProviderUpdateRequest) UnsetRestrictedToOrganizationId()`

UnsetRestrictedToOrganizationId ensures that no value is present for RestrictedToOrganizationId, not even an explicit nil
### GetServesUnlistedModels

`func (o *HostedProviderUpdateRequest) GetServesUnlistedModels() bool`

GetServesUnlistedModels returns the ServesUnlistedModels field if non-nil, zero value otherwise.

### GetServesUnlistedModelsOk

`func (o *HostedProviderUpdateRequest) GetServesUnlistedModelsOk() (*bool, bool)`

GetServesUnlistedModelsOk returns a tuple with the ServesUnlistedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServesUnlistedModels

`func (o *HostedProviderUpdateRequest) SetServesUnlistedModels(v bool)`

SetServesUnlistedModels sets ServesUnlistedModels field to given value.

### HasServesUnlistedModels

`func (o *HostedProviderUpdateRequest) HasServesUnlistedModels() bool`

HasServesUnlistedModels returns a boolean if a field has been set.

### SetServesUnlistedModelsNil

`func (o *HostedProviderUpdateRequest) SetServesUnlistedModelsNil(b bool)`

 SetServesUnlistedModelsNil sets the value for ServesUnlistedModels to be an explicit nil

### UnsetServesUnlistedModels
`func (o *HostedProviderUpdateRequest) UnsetServesUnlistedModels()`

UnsetServesUnlistedModels ensures that no value is present for ServesUnlistedModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


