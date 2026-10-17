# HostedProviderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | **string** |  | 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Provider** | **string** |  | 
**RestrictedToOrganizationId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHostedProviderCreateRequest

`func NewHostedProviderCreateRequest(apiKey string, provider string, ) *HostedProviderCreateRequest`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


