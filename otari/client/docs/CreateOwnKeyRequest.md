# CreateOwnKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** | Model allow-list: null &#x3D; any model your user default allows, [] &#x3D; deny all, or canonical instance:model entries. A key can only narrow your own model access, never broaden it. | [optional] 
**AllowedTools** | Pointer to **[]string** | Search tool allow-list for POST /api/v1/search: null &#x3D; inherit your own, [] &#x3D; no tool, or search tool names (&#39;*&#39; for every search tool). A key can only narrow your own tool access. | [optional] 
**CaptureAgentTelemetry** | Pointer to **NullableBool** | Per-key override of the deployment-wide capture_agent_telemetry setting: null (default) inherits it, true always stores this key&#39;s coding-agent telemetry, false always discards it. | [optional] 
**EndUserBudgetId** | Pointer to **NullableString** | Budget of this organization each end user this key creates is capped at, unless the request names another with Otari-End-User-Budget. Null leaves end users capped only by this key&#39;s own ceiling. | [optional] 
**EndUserBudgetIds** | Pointer to **[]string** | Budgets of this organization a request may start a new end user on by naming one in Otari-End-User-Budget. Null allows end_user_budget_id alone. When both are set, end_user_budget_id must be on the list. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Optional expiration timestamp | [optional] 
**IsServiceKey** | Pointer to **bool** | When true, a request may name an end user in its &#39;user&#39; field. Each end user is created on first use and billed to its own budget, while this key&#39;s own ceiling caps all of them together. Only an organization owner or admin may create one. It belongs to the organization rather than to you: every owner or admin manages it, and its end users are reachable from no other key. Fixed once the key exists. | [optional] [default to false]
**KeyName** | Pointer to **NullableString** | Optional name for the key | [optional] 
**MaxEndUsers** | Pointer to **NullableInt32** | Most live end users this service key may have. A request that would create one more is refused with 403 and end_user_limit_reached; deleting end users frees room. Null is 10000: an organization&#39;s service key always has a cap. | [optional] 
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
### GetAllowedTools

`func (o *CreateOwnKeyRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *CreateOwnKeyRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *CreateOwnKeyRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *CreateOwnKeyRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### SetAllowedToolsNil

`func (o *CreateOwnKeyRequest) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *CreateOwnKeyRequest) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
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
### GetEndUserBudgetId

`func (o *CreateOwnKeyRequest) GetEndUserBudgetId() string`

GetEndUserBudgetId returns the EndUserBudgetId field if non-nil, zero value otherwise.

### GetEndUserBudgetIdOk

`func (o *CreateOwnKeyRequest) GetEndUserBudgetIdOk() (*string, bool)`

GetEndUserBudgetIdOk returns a tuple with the EndUserBudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserBudgetId

`func (o *CreateOwnKeyRequest) SetEndUserBudgetId(v string)`

SetEndUserBudgetId sets EndUserBudgetId field to given value.

### HasEndUserBudgetId

`func (o *CreateOwnKeyRequest) HasEndUserBudgetId() bool`

HasEndUserBudgetId returns a boolean if a field has been set.

### SetEndUserBudgetIdNil

`func (o *CreateOwnKeyRequest) SetEndUserBudgetIdNil(b bool)`

 SetEndUserBudgetIdNil sets the value for EndUserBudgetId to be an explicit nil

### UnsetEndUserBudgetId
`func (o *CreateOwnKeyRequest) UnsetEndUserBudgetId()`

UnsetEndUserBudgetId ensures that no value is present for EndUserBudgetId, not even an explicit nil
### GetEndUserBudgetIds

`func (o *CreateOwnKeyRequest) GetEndUserBudgetIds() []string`

GetEndUserBudgetIds returns the EndUserBudgetIds field if non-nil, zero value otherwise.

### GetEndUserBudgetIdsOk

`func (o *CreateOwnKeyRequest) GetEndUserBudgetIdsOk() (*[]string, bool)`

GetEndUserBudgetIdsOk returns a tuple with the EndUserBudgetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserBudgetIds

`func (o *CreateOwnKeyRequest) SetEndUserBudgetIds(v []string)`

SetEndUserBudgetIds sets EndUserBudgetIds field to given value.

### HasEndUserBudgetIds

`func (o *CreateOwnKeyRequest) HasEndUserBudgetIds() bool`

HasEndUserBudgetIds returns a boolean if a field has been set.

### SetEndUserBudgetIdsNil

`func (o *CreateOwnKeyRequest) SetEndUserBudgetIdsNil(b bool)`

 SetEndUserBudgetIdsNil sets the value for EndUserBudgetIds to be an explicit nil

### UnsetEndUserBudgetIds
`func (o *CreateOwnKeyRequest) UnsetEndUserBudgetIds()`

UnsetEndUserBudgetIds ensures that no value is present for EndUserBudgetIds, not even an explicit nil
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
### GetIsServiceKey

`func (o *CreateOwnKeyRequest) GetIsServiceKey() bool`

GetIsServiceKey returns the IsServiceKey field if non-nil, zero value otherwise.

### GetIsServiceKeyOk

`func (o *CreateOwnKeyRequest) GetIsServiceKeyOk() (*bool, bool)`

GetIsServiceKeyOk returns a tuple with the IsServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceKey

`func (o *CreateOwnKeyRequest) SetIsServiceKey(v bool)`

SetIsServiceKey sets IsServiceKey field to given value.

### HasIsServiceKey

`func (o *CreateOwnKeyRequest) HasIsServiceKey() bool`

HasIsServiceKey returns a boolean if a field has been set.

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
### GetMaxEndUsers

`func (o *CreateOwnKeyRequest) GetMaxEndUsers() int32`

GetMaxEndUsers returns the MaxEndUsers field if non-nil, zero value otherwise.

### GetMaxEndUsersOk

`func (o *CreateOwnKeyRequest) GetMaxEndUsersOk() (*int32, bool)`

GetMaxEndUsersOk returns a tuple with the MaxEndUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndUsers

`func (o *CreateOwnKeyRequest) SetMaxEndUsers(v int32)`

SetMaxEndUsers sets MaxEndUsers field to given value.

### HasMaxEndUsers

`func (o *CreateOwnKeyRequest) HasMaxEndUsers() bool`

HasMaxEndUsers returns a boolean if a field has been set.

### SetMaxEndUsersNil

`func (o *CreateOwnKeyRequest) SetMaxEndUsersNil(b bool)`

 SetMaxEndUsersNil sets the value for MaxEndUsers to be an explicit nil

### UnsetMaxEndUsers
`func (o *CreateOwnKeyRequest) UnsetMaxEndUsers()`

UnsetMaxEndUsers ensures that no value is present for MaxEndUsers, not even an explicit nil
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


