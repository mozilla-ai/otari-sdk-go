# CreateKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**ExcludeFromBudget** | **bool** |  | 
**ExpiresAt** | **NullableString** |  | 
**Id** | **string** |  | 
**IsActive** | **bool** |  | 
**Key** | **string** |  | 
**KeyName** | **NullableString** |  | 
**KeyPrefix** | **NullableString** |  | 
**Metadata** | **map[string]interface{}** |  | 
**RejectUserMismatch** | **NullableBool** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewCreateKeyResponse

`func NewCreateKeyResponse(allowedModels []string, createdAt string, excludeFromBudget bool, expiresAt NullableString, id string, isActive bool, key string, keyName NullableString, keyPrefix NullableString, metadata map[string]interface{}, rejectUserMismatch NullableBool, userId NullableString, ) *CreateKeyResponse`

NewCreateKeyResponse instantiates a new CreateKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyResponseWithDefaults

`func NewCreateKeyResponseWithDefaults() *CreateKeyResponse`

NewCreateKeyResponseWithDefaults instantiates a new CreateKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *CreateKeyResponse) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *CreateKeyResponse) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *CreateKeyResponse) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.


### SetAllowedModelsNil

`func (o *CreateKeyResponse) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *CreateKeyResponse) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCreatedAt

`func (o *CreateKeyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateKeyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateKeyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExcludeFromBudget

`func (o *CreateKeyResponse) GetExcludeFromBudget() bool`

GetExcludeFromBudget returns the ExcludeFromBudget field if non-nil, zero value otherwise.

### GetExcludeFromBudgetOk

`func (o *CreateKeyResponse) GetExcludeFromBudgetOk() (*bool, bool)`

GetExcludeFromBudgetOk returns a tuple with the ExcludeFromBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBudget

`func (o *CreateKeyResponse) SetExcludeFromBudget(v bool)`

SetExcludeFromBudget sets ExcludeFromBudget field to given value.


### GetExpiresAt

`func (o *CreateKeyResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateKeyResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateKeyResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CreateKeyResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateKeyResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *CreateKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateKeyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *CreateKeyResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateKeyResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateKeyResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetKey

`func (o *CreateKeyResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateKeyResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateKeyResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetKeyName

`func (o *CreateKeyResponse) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateKeyResponse) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateKeyResponse) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### SetKeyNameNil

`func (o *CreateKeyResponse) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *CreateKeyResponse) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetKeyPrefix

`func (o *CreateKeyResponse) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *CreateKeyResponse) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *CreateKeyResponse) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### SetKeyPrefixNil

`func (o *CreateKeyResponse) SetKeyPrefixNil(b bool)`

 SetKeyPrefixNil sets the value for KeyPrefix to be an explicit nil

### UnsetKeyPrefix
`func (o *CreateKeyResponse) UnsetKeyPrefix()`

UnsetKeyPrefix ensures that no value is present for KeyPrefix, not even an explicit nil
### GetMetadata

`func (o *CreateKeyResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateKeyResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateKeyResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetRejectUserMismatch

`func (o *CreateKeyResponse) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *CreateKeyResponse) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *CreateKeyResponse) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.


### SetRejectUserMismatchNil

`func (o *CreateKeyResponse) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *CreateKeyResponse) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil
### GetUserId

`func (o *CreateKeyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateKeyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateKeyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *CreateKeyResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateKeyResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


