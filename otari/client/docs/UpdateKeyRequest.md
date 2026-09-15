# UpdateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** |  | [optional] 
**CaptureAgentTelemetry** | Pointer to **NullableBool** |  | [optional] 
**ExcludeFromBudget** | Pointer to **NullableBool** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**RejectUserMismatch** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateKeyRequest

`func NewUpdateKeyRequest() *UpdateKeyRequest`

NewUpdateKeyRequest instantiates a new UpdateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeyRequestWithDefaults

`func NewUpdateKeyRequestWithDefaults() *UpdateKeyRequest`

NewUpdateKeyRequestWithDefaults instantiates a new UpdateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *UpdateKeyRequest) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *UpdateKeyRequest) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *UpdateKeyRequest) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *UpdateKeyRequest) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *UpdateKeyRequest) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *UpdateKeyRequest) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCaptureAgentTelemetry

`func (o *UpdateKeyRequest) GetCaptureAgentTelemetry() bool`

GetCaptureAgentTelemetry returns the CaptureAgentTelemetry field if non-nil, zero value otherwise.

### GetCaptureAgentTelemetryOk

`func (o *UpdateKeyRequest) GetCaptureAgentTelemetryOk() (*bool, bool)`

GetCaptureAgentTelemetryOk returns a tuple with the CaptureAgentTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAgentTelemetry

`func (o *UpdateKeyRequest) SetCaptureAgentTelemetry(v bool)`

SetCaptureAgentTelemetry sets CaptureAgentTelemetry field to given value.

### HasCaptureAgentTelemetry

`func (o *UpdateKeyRequest) HasCaptureAgentTelemetry() bool`

HasCaptureAgentTelemetry returns a boolean if a field has been set.

### SetCaptureAgentTelemetryNil

`func (o *UpdateKeyRequest) SetCaptureAgentTelemetryNil(b bool)`

 SetCaptureAgentTelemetryNil sets the value for CaptureAgentTelemetry to be an explicit nil

### UnsetCaptureAgentTelemetry
`func (o *UpdateKeyRequest) UnsetCaptureAgentTelemetry()`

UnsetCaptureAgentTelemetry ensures that no value is present for CaptureAgentTelemetry, not even an explicit nil
### GetExcludeFromBudget

`func (o *UpdateKeyRequest) GetExcludeFromBudget() bool`

GetExcludeFromBudget returns the ExcludeFromBudget field if non-nil, zero value otherwise.

### GetExcludeFromBudgetOk

`func (o *UpdateKeyRequest) GetExcludeFromBudgetOk() (*bool, bool)`

GetExcludeFromBudgetOk returns a tuple with the ExcludeFromBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBudget

`func (o *UpdateKeyRequest) SetExcludeFromBudget(v bool)`

SetExcludeFromBudget sets ExcludeFromBudget field to given value.

### HasExcludeFromBudget

`func (o *UpdateKeyRequest) HasExcludeFromBudget() bool`

HasExcludeFromBudget returns a boolean if a field has been set.

### SetExcludeFromBudgetNil

`func (o *UpdateKeyRequest) SetExcludeFromBudgetNil(b bool)`

 SetExcludeFromBudgetNil sets the value for ExcludeFromBudget to be an explicit nil

### UnsetExcludeFromBudget
`func (o *UpdateKeyRequest) UnsetExcludeFromBudget()`

UnsetExcludeFromBudget ensures that no value is present for ExcludeFromBudget, not even an explicit nil
### GetExpiresAt

`func (o *UpdateKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *UpdateKeyRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *UpdateKeyRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsActive

`func (o *UpdateKeyRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateKeyRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateKeyRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateKeyRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateKeyRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateKeyRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetKeyName

`func (o *UpdateKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *UpdateKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *UpdateKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *UpdateKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *UpdateKeyRequest) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *UpdateKeyRequest) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetMetadata

`func (o *UpdateKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateKeyRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateKeyRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRejectUserMismatch

`func (o *UpdateKeyRequest) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *UpdateKeyRequest) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *UpdateKeyRequest) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.

### HasRejectUserMismatch

`func (o *UpdateKeyRequest) HasRejectUserMismatch() bool`

HasRejectUserMismatch returns a boolean if a field has been set.

### SetRejectUserMismatchNil

`func (o *UpdateKeyRequest) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *UpdateKeyRequest) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


