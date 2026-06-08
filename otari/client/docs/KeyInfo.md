# KeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**ExpiresAt** | **NullableString** |  | 
**Id** | **string** |  | 
**IsActive** | **bool** |  | 
**KeyName** | **NullableString** |  | 
**LastUsedAt** | **NullableString** |  | 
**Metadata** | **map[string]interface{}** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewKeyInfo

`func NewKeyInfo(createdAt string, expiresAt NullableString, id string, isActive bool, keyName NullableString, lastUsedAt NullableString, metadata map[string]interface{}, userId NullableString, ) *KeyInfo`

NewKeyInfo instantiates a new KeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoWithDefaults

`func NewKeyInfoWithDefaults() *KeyInfo`

NewKeyInfoWithDefaults instantiates a new KeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KeyInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *KeyInfo) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *KeyInfo) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *KeyInfo) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *KeyInfo) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *KeyInfo) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *KeyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *KeyInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *KeyInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *KeyInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetKeyName

`func (o *KeyInfo) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *KeyInfo) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *KeyInfo) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### SetKeyNameNil

`func (o *KeyInfo) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *KeyInfo) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetLastUsedAt

`func (o *KeyInfo) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *KeyInfo) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *KeyInfo) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *KeyInfo) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *KeyInfo) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetMetadata

`func (o *KeyInfo) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KeyInfo) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KeyInfo) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetUserId

`func (o *KeyInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KeyInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KeyInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *KeyInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *KeyInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


