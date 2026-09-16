# StoredSearchToolSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Decryptable** | Pointer to **bool** |  | [optional] [default to true]
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Provider** | **string** |  | 
**ShadowsConfig** | Pointer to **bool** | True when a config-file search tool of the same name exists; the stored one is in effect. | [optional] [default to false]
**Timeout** | Pointer to **NullableFloat32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStoredSearchToolSchema

`func NewStoredSearchToolSchema(name string, provider string, ) *StoredSearchToolSchema`

NewStoredSearchToolSchema instantiates a new StoredSearchToolSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredSearchToolSchemaWithDefaults

`func NewStoredSearchToolSchemaWithDefaults() *StoredSearchToolSchema`

NewStoredSearchToolSchemaWithDefaults instantiates a new StoredSearchToolSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *StoredSearchToolSchema) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *StoredSearchToolSchema) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *StoredSearchToolSchema) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *StoredSearchToolSchema) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *StoredSearchToolSchema) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *StoredSearchToolSchema) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetCreatedAt

`func (o *StoredSearchToolSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoredSearchToolSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoredSearchToolSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoredSearchToolSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *StoredSearchToolSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StoredSearchToolSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDecryptable

`func (o *StoredSearchToolSchema) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *StoredSearchToolSchema) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *StoredSearchToolSchema) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *StoredSearchToolSchema) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetLast4

`func (o *StoredSearchToolSchema) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StoredSearchToolSchema) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StoredSearchToolSchema) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StoredSearchToolSchema) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *StoredSearchToolSchema) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *StoredSearchToolSchema) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *StoredSearchToolSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoredSearchToolSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoredSearchToolSchema) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *StoredSearchToolSchema) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StoredSearchToolSchema) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StoredSearchToolSchema) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StoredSearchToolSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProvider

`func (o *StoredSearchToolSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *StoredSearchToolSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *StoredSearchToolSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetShadowsConfig

`func (o *StoredSearchToolSchema) GetShadowsConfig() bool`

GetShadowsConfig returns the ShadowsConfig field if non-nil, zero value otherwise.

### GetShadowsConfigOk

`func (o *StoredSearchToolSchema) GetShadowsConfigOk() (*bool, bool)`

GetShadowsConfigOk returns a tuple with the ShadowsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowsConfig

`func (o *StoredSearchToolSchema) SetShadowsConfig(v bool)`

SetShadowsConfig sets ShadowsConfig field to given value.

### HasShadowsConfig

`func (o *StoredSearchToolSchema) HasShadowsConfig() bool`

HasShadowsConfig returns a boolean if a field has been set.

### GetTimeout

`func (o *StoredSearchToolSchema) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StoredSearchToolSchema) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StoredSearchToolSchema) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StoredSearchToolSchema) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *StoredSearchToolSchema) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *StoredSearchToolSchema) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUpdatedAt

`func (o *StoredSearchToolSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoredSearchToolSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoredSearchToolSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StoredSearchToolSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *StoredSearchToolSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StoredSearchToolSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


