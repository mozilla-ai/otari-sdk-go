# CreateOwnKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** | Model allow-list: null &#x3D; any model your user default allows, [] &#x3D; deny all, or canonical instance:model entries. A key can only narrow your own model access, never broaden it. | [optional] 
**CaptureAgentTelemetry** | Pointer to **NullableBool** | Per-key override of the deployment-wide capture_agent_telemetry setting: null (default) inherits it, true always stores this key&#39;s coding-agent telemetry, false always discards it. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Optional expiration timestamp | [optional] 
**KeyName** | Pointer to **NullableString** | Optional name for the key | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional metadata | [optional] 
**RejectUserMismatch** | Pointer to **NullableBool** | Per-key override of the deployment-wide reject_user_mismatch setting: null (default) inherits it, true always rejects a request naming a different &#39;user&#39;, false always accepts it. Spend binds to your own user either way. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Workspace this key belongs to, which must be one you may see in your active organization. Omitted means that organization&#39;s default workspace, refused when you are not a member of it. | [optional] 

## Methods

### NewCreateOwnKeyRequest

`func NewCreateOwnKeyRequest() *CreateOwnKeyRequest`

NewCreateOwnKeyRequest instantiates a new CreateOwnKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOwnKeyRequestWithDefaults

`func NewCreateOwnKeyRequestWithDefaults() *CreateOwnKeyRequest`

NewCreateOwnKeyRequestWithDefaults instantiates a new CreateOwnKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *CreateOwnKeyRequest) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *CreateOwnKeyRequest) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *CreateOwnKeyRequest) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *CreateOwnKeyRequest) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *CreateOwnKeyRequest) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *CreateOwnKeyRequest) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCaptureAgentTelemetry

`func (o *CreateOwnKeyRequest) GetCaptureAgentTelemetry() bool`

GetCaptureAgentTelemetry returns the CaptureAgentTelemetry field if non-nil, zero value otherwise.

### GetCaptureAgentTelemetryOk

`func (o *CreateOwnKeyRequest) GetCaptureAgentTelemetryOk() (*bool, bool)`

GetCaptureAgentTelemetryOk returns a tuple with the CaptureAgentTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAgentTelemetry

`func (o *CreateOwnKeyRequest) SetCaptureAgentTelemetry(v bool)`

SetCaptureAgentTelemetry sets CaptureAgentTelemetry field to given value.

### HasCaptureAgentTelemetry

`func (o *CreateOwnKeyRequest) HasCaptureAgentTelemetry() bool`

HasCaptureAgentTelemetry returns a boolean if a field has been set.

### SetCaptureAgentTelemetryNil

`func (o *CreateOwnKeyRequest) SetCaptureAgentTelemetryNil(b bool)`

 SetCaptureAgentTelemetryNil sets the value for CaptureAgentTelemetry to be an explicit nil

### UnsetCaptureAgentTelemetry
`func (o *CreateOwnKeyRequest) UnsetCaptureAgentTelemetry()`

UnsetCaptureAgentTelemetry ensures that no value is present for CaptureAgentTelemetry, not even an explicit nil
### GetExpiresAt

`func (o *CreateOwnKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateOwnKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateOwnKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateOwnKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateOwnKeyRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateOwnKeyRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetKeyName

`func (o *CreateOwnKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateOwnKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateOwnKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *CreateOwnKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *CreateOwnKeyRequest) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *CreateOwnKeyRequest) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetMetadata

`func (o *CreateOwnKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateOwnKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateOwnKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateOwnKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRejectUserMismatch

`func (o *CreateOwnKeyRequest) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *CreateOwnKeyRequest) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *CreateOwnKeyRequest) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.

### HasRejectUserMismatch

`func (o *CreateOwnKeyRequest) HasRejectUserMismatch() bool`

HasRejectUserMismatch returns a boolean if a field has been set.

### SetRejectUserMismatchNil

`func (o *CreateOwnKeyRequest) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *CreateOwnKeyRequest) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil
### GetWorkspaceId

`func (o *CreateOwnKeyRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CreateOwnKeyRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CreateOwnKeyRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *CreateOwnKeyRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *CreateOwnKeyRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *CreateOwnKeyRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


