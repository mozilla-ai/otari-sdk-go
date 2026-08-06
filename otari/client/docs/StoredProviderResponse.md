# StoredProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Decryptable** | Pointer to **bool** |  | [optional] [default to true]
**Instance** | **string** |  | 
**Last4** | Pointer to **NullableString** |  | [optional] 
**ProviderType** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStoredProviderResponse

`func NewStoredProviderResponse(instance string, ) *StoredProviderResponse`

NewStoredProviderResponse instantiates a new StoredProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredProviderResponseWithDefaults

`func NewStoredProviderResponseWithDefaults() *StoredProviderResponse`

NewStoredProviderResponseWithDefaults instantiates a new StoredProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *StoredProviderResponse) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *StoredProviderResponse) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *StoredProviderResponse) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *StoredProviderResponse) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *StoredProviderResponse) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *StoredProviderResponse) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetClientArgs

`func (o *StoredProviderResponse) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *StoredProviderResponse) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *StoredProviderResponse) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *StoredProviderResponse) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StoredProviderResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoredProviderResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoredProviderResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoredProviderResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *StoredProviderResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StoredProviderResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDecryptable

`func (o *StoredProviderResponse) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *StoredProviderResponse) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *StoredProviderResponse) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *StoredProviderResponse) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetInstance

`func (o *StoredProviderResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *StoredProviderResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *StoredProviderResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetLast4

`func (o *StoredProviderResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StoredProviderResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StoredProviderResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StoredProviderResponse) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *StoredProviderResponse) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *StoredProviderResponse) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetProviderType

`func (o *StoredProviderResponse) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *StoredProviderResponse) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *StoredProviderResponse) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *StoredProviderResponse) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### SetProviderTypeNil

`func (o *StoredProviderResponse) SetProviderTypeNil(b bool)`

 SetProviderTypeNil sets the value for ProviderType to be an explicit nil

### UnsetProviderType
`func (o *StoredProviderResponse) UnsetProviderType()`

UnsetProviderType ensures that no value is present for ProviderType, not even an explicit nil
### GetUpdatedAt

`func (o *StoredProviderResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoredProviderResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoredProviderResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StoredProviderResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *StoredProviderResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StoredProviderResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


