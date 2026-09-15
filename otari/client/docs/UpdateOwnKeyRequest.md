# UpdateOwnKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** |  | [optional] 
**CaptureAgentTelemetry** | Pointer to **NullableBool** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**RejectUserMismatch** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateOwnKeyRequest

`func NewUpdateOwnKeyRequest() *UpdateOwnKeyRequest`

NewUpdateOwnKeyRequest instantiates a new UpdateOwnKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOwnKeyRequestWithDefaults

`func NewUpdateOwnKeyRequestWithDefaults() *UpdateOwnKeyRequest`

NewUpdateOwnKeyRequestWithDefaults instantiates a new UpdateOwnKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *UpdateOwnKeyRequest) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *UpdateOwnKeyRequest) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *UpdateOwnKeyRequest) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *UpdateOwnKeyRequest) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *UpdateOwnKeyRequest) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *UpdateOwnKeyRequest) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCaptureAgentTelemetry

`func (o *UpdateOwnKeyRequest) GetCaptureAgentTelemetry() bool`

GetCaptureAgentTelemetry returns the CaptureAgentTelemetry field if non-nil, zero value otherwise.

### GetCaptureAgentTelemetryOk

`func (o *UpdateOwnKeyRequest) GetCaptureAgentTelemetryOk() (*bool, bool)`

GetCaptureAgentTelemetryOk returns a tuple with the CaptureAgentTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAgentTelemetry

`func (o *UpdateOwnKeyRequest) SetCaptureAgentTelemetry(v bool)`

SetCaptureAgentTelemetry sets CaptureAgentTelemetry field to given value.

### HasCaptureAgentTelemetry

`func (o *UpdateOwnKeyRequest) HasCaptureAgentTelemetry() bool`

HasCaptureAgentTelemetry returns a boolean if a field has been set.

### SetCaptureAgentTelemetryNil

`func (o *UpdateOwnKeyRequest) SetCaptureAgentTelemetryNil(b bool)`

 SetCaptureAgentTelemetryNil sets the value for CaptureAgentTelemetry to be an explicit nil

### UnsetCaptureAgentTelemetry
`func (o *UpdateOwnKeyRequest) UnsetCaptureAgentTelemetry()`

UnsetCaptureAgentTelemetry ensures that no value is present for CaptureAgentTelemetry, not even an explicit nil
### GetExpiresAt

`func (o *UpdateOwnKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateOwnKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateOwnKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateOwnKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *UpdateOwnKeyRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *UpdateOwnKeyRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsActive

`func (o *UpdateOwnKeyRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateOwnKeyRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateOwnKeyRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateOwnKeyRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateOwnKeyRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateOwnKeyRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetKeyName

`func (o *UpdateOwnKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *UpdateOwnKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *UpdateOwnKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *UpdateOwnKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *UpdateOwnKeyRequest) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *UpdateOwnKeyRequest) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetMetadata

`func (o *UpdateOwnKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateOwnKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateOwnKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateOwnKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateOwnKeyRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateOwnKeyRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetRejectUserMismatch

`func (o *UpdateOwnKeyRequest) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *UpdateOwnKeyRequest) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *UpdateOwnKeyRequest) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.

### HasRejectUserMismatch

`func (o *UpdateOwnKeyRequest) HasRejectUserMismatch() bool`

HasRejectUserMismatch returns a boolean if a field has been set.

### SetRejectUserMismatchNil

`func (o *UpdateOwnKeyRequest) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *UpdateOwnKeyRequest) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


