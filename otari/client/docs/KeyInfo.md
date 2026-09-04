# KeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | **[]string** |  | 
**CaptureAgentTelemetry** | **NullableBool** |  | 
**CreatedAt** | **string** |  | 
**ExcludeFromBudget** | **bool** |  | 
**ExpiresAt** | **NullableString** |  | 
**Id** | **string** |  | 
**IsActive** | **bool** |  | 
**KeyName** | **NullableString** |  | 
**KeyPrefix** | **NullableString** |  | 
**LastUsedAt** | **NullableString** |  | 
**Metadata** | **map[string]interface{}** |  | 
**RejectUserMismatch** | **NullableBool** |  | 
**UserId** | **NullableString** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewKeyInfo

`func NewKeyInfo(allowedModels []string, captureAgentTelemetry NullableBool, createdAt string, excludeFromBudget bool, expiresAt NullableString, id string, isActive bool, keyName NullableString, keyPrefix NullableString, lastUsedAt NullableString, metadata map[string]interface{}, rejectUserMismatch NullableBool, userId NullableString, workspaceId string, ) *KeyInfo`

NewKeyInfo instantiates a new KeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoWithDefaults

`func NewKeyInfoWithDefaults() *KeyInfo`

NewKeyInfoWithDefaults instantiates a new KeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *KeyInfo) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *KeyInfo) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *KeyInfo) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.


### SetAllowedModelsNil

`func (o *KeyInfo) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *KeyInfo) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCaptureAgentTelemetry

`func (o *KeyInfo) GetCaptureAgentTelemetry() bool`

GetCaptureAgentTelemetry returns the CaptureAgentTelemetry field if non-nil, zero value otherwise.

### GetCaptureAgentTelemetryOk

`func (o *KeyInfo) GetCaptureAgentTelemetryOk() (*bool, bool)`

GetCaptureAgentTelemetryOk returns a tuple with the CaptureAgentTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAgentTelemetry

`func (o *KeyInfo) SetCaptureAgentTelemetry(v bool)`

SetCaptureAgentTelemetry sets CaptureAgentTelemetry field to given value.


### SetCaptureAgentTelemetryNil

`func (o *KeyInfo) SetCaptureAgentTelemetryNil(b bool)`

 SetCaptureAgentTelemetryNil sets the value for CaptureAgentTelemetry to be an explicit nil

### UnsetCaptureAgentTelemetry
`func (o *KeyInfo) UnsetCaptureAgentTelemetry()`

UnsetCaptureAgentTelemetry ensures that no value is present for CaptureAgentTelemetry, not even an explicit nil
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


### GetExcludeFromBudget

`func (o *KeyInfo) GetExcludeFromBudget() bool`

GetExcludeFromBudget returns the ExcludeFromBudget field if non-nil, zero value otherwise.

### GetExcludeFromBudgetOk

`func (o *KeyInfo) GetExcludeFromBudgetOk() (*bool, bool)`

GetExcludeFromBudgetOk returns a tuple with the ExcludeFromBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBudget

`func (o *KeyInfo) SetExcludeFromBudget(v bool)`

SetExcludeFromBudget sets ExcludeFromBudget field to given value.


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
### GetKeyPrefix

`func (o *KeyInfo) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *KeyInfo) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *KeyInfo) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.


### SetKeyPrefixNil

`func (o *KeyInfo) SetKeyPrefixNil(b bool)`

 SetKeyPrefixNil sets the value for KeyPrefix to be an explicit nil

### UnsetKeyPrefix
`func (o *KeyInfo) UnsetKeyPrefix()`

UnsetKeyPrefix ensures that no value is present for KeyPrefix, not even an explicit nil
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


### GetRejectUserMismatch

`func (o *KeyInfo) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *KeyInfo) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *KeyInfo) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.


### SetRejectUserMismatchNil

`func (o *KeyInfo) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *KeyInfo) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil
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
### GetWorkspaceId

`func (o *KeyInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KeyInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KeyInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


