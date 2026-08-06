# UpdateStoredProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** | New API key. Omit to keep the existing one. Never returned. | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpectedUpdatedAt** | Pointer to **NullableString** | Optimistic concurrency: if set, the update 412s unless it matches the stored updated_at. | [optional] 
**ProviderType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateStoredProviderRequest

`func NewUpdateStoredProviderRequest() *UpdateStoredProviderRequest`

NewUpdateStoredProviderRequest instantiates a new UpdateStoredProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoredProviderRequestWithDefaults

`func NewUpdateStoredProviderRequestWithDefaults() *UpdateStoredProviderRequest`

NewUpdateStoredProviderRequestWithDefaults instantiates a new UpdateStoredProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *UpdateStoredProviderRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *UpdateStoredProviderRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *UpdateStoredProviderRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *UpdateStoredProviderRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *UpdateStoredProviderRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *UpdateStoredProviderRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *UpdateStoredProviderRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateStoredProviderRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateStoredProviderRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateStoredProviderRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *UpdateStoredProviderRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *UpdateStoredProviderRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetClientArgs

`func (o *UpdateStoredProviderRequest) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *UpdateStoredProviderRequest) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *UpdateStoredProviderRequest) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *UpdateStoredProviderRequest) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *UpdateStoredProviderRequest) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *UpdateStoredProviderRequest) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetExpectedUpdatedAt

`func (o *UpdateStoredProviderRequest) GetExpectedUpdatedAt() string`

GetExpectedUpdatedAt returns the ExpectedUpdatedAt field if non-nil, zero value otherwise.

### GetExpectedUpdatedAtOk

`func (o *UpdateStoredProviderRequest) GetExpectedUpdatedAtOk() (*string, bool)`

GetExpectedUpdatedAtOk returns a tuple with the ExpectedUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUpdatedAt

`func (o *UpdateStoredProviderRequest) SetExpectedUpdatedAt(v string)`

SetExpectedUpdatedAt sets ExpectedUpdatedAt field to given value.

### HasExpectedUpdatedAt

`func (o *UpdateStoredProviderRequest) HasExpectedUpdatedAt() bool`

HasExpectedUpdatedAt returns a boolean if a field has been set.

### SetExpectedUpdatedAtNil

`func (o *UpdateStoredProviderRequest) SetExpectedUpdatedAtNil(b bool)`

 SetExpectedUpdatedAtNil sets the value for ExpectedUpdatedAt to be an explicit nil

### UnsetExpectedUpdatedAt
`func (o *UpdateStoredProviderRequest) UnsetExpectedUpdatedAt()`

UnsetExpectedUpdatedAt ensures that no value is present for ExpectedUpdatedAt, not even an explicit nil
### GetProviderType

`func (o *UpdateStoredProviderRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *UpdateStoredProviderRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *UpdateStoredProviderRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *UpdateStoredProviderRequest) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### SetProviderTypeNil

`func (o *UpdateStoredProviderRequest) SetProviderTypeNil(b bool)`

 SetProviderTypeNil sets the value for ProviderType to be an explicit nil

### UnsetProviderType
`func (o *UpdateStoredProviderRequest) UnsetProviderType()`

UnsetProviderType ensures that no value is present for ProviderType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


