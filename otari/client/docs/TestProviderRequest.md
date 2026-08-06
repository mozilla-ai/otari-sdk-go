# TestProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**Instance** | Pointer to **NullableString** | Provider/instance name; the impl when no provider_type. | [optional] 
**ProviderType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTestProviderRequest

`func NewTestProviderRequest() *TestProviderRequest`

NewTestProviderRequest instantiates a new TestProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestProviderRequestWithDefaults

`func NewTestProviderRequestWithDefaults() *TestProviderRequest`

NewTestProviderRequestWithDefaults instantiates a new TestProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *TestProviderRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *TestProviderRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *TestProviderRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *TestProviderRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *TestProviderRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *TestProviderRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *TestProviderRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TestProviderRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TestProviderRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TestProviderRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *TestProviderRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *TestProviderRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetClientArgs

`func (o *TestProviderRequest) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *TestProviderRequest) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *TestProviderRequest) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *TestProviderRequest) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *TestProviderRequest) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *TestProviderRequest) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetInstance

`func (o *TestProviderRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *TestProviderRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *TestProviderRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *TestProviderRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *TestProviderRequest) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *TestProviderRequest) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetProviderType

`func (o *TestProviderRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *TestProviderRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *TestProviderRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *TestProviderRequest) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### SetProviderTypeNil

`func (o *TestProviderRequest) SetProviderTypeNil(b bool)`

 SetProviderTypeNil sets the value for ProviderType to be an explicit nil

### UnsetProviderType
`func (o *TestProviderRequest) UnsetProviderType()`

UnsetProviderType ensures that no value is present for ProviderType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


