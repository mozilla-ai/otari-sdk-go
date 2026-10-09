# CreateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** | Model allow-list: null &#x3D; any model, [] &#x3D; deny all, or canonical instance:model entries (with instance:* / instance:prefix* wildcards). | [optional] 
**CaptureAgentTelemetry** | Pointer to **NullableBool** | Per-key override of the deployment-wide capture_agent_telemetry setting: null (default) inherits it, true always stores this key&#39;s coding-agent telemetry, false always discards it. Covers both behavioral events (tool_result, tool_decision, user_prompt, api_error) from POST /otlp/v1/logs and outcome-metric data points (lines of code, commits, pull requests, active time) from POST /otlp/v1/metrics. Usage capture and billing are unaffected either way. | [optional] 
**EndUserBudgetId** | Pointer to **NullableString** | Budget each end user this key creates is capped at, unless the request names another with Otari-End-User-Budget. Null leaves end users capped only by this key&#39;s own ceiling. | [optional] 
**EndUserBudgetIds** | Pointer to **[]string** | Budgets a request may start a new end user on by naming one in Otari-End-User-Budget. Null allows end_user_budget_id alone. When both are set, end_user_budget_id must be on the list. | [optional] 
**ExcludeFromBudget** | Pointer to **bool** | When true, requests on this key are logged with cost but never reserved, reconciled into the user&#39;s spend, or gated by budget. | [optional] [default to false]
**ExpiresAt** | Pointer to **NullableTime** | Optional expiration timestamp | [optional] 
**IsServiceKey** | Pointer to **bool** | When true, a request may name an end user in its &#39;user&#39; field. Each end user is created on first use, owned by this key&#39;s user, and billed to its own budget, while this key&#39;s own ceiling caps all of them together. | [optional] [default to false]
**KeyName** | Pointer to **NullableString** | Optional name for the key | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional metadata | [optional] 
**RejectUserMismatch** | Pointer to **NullableBool** | Per-key override of the deployment-wide reject_user_mismatch setting: null (default) inherits it, true always rejects a request naming a different &#39;user&#39;, false always accepts it. Spend binds to this key&#39;s own user either way. | [optional] 
**UserId** | Pointer to **NullableString** | Optional user ID to associate with this key | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Workspace this key belongs to, which must be one in the caller&#39;s organization. Omitted means that organization&#39;s default workspace. A key belongs to exactly one workspace: requests on it are scoped and billed there, so the workspace is read off the key rather than off a request header. | [optional] 

## Methods

### NewCreateKeyRequest

`func NewCreateKeyRequest() *CreateKeyRequest`

NewCreateKeyRequest instantiates a new CreateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyRequestWithDefaults

`func NewCreateKeyRequestWithDefaults() *CreateKeyRequest`

NewCreateKeyRequestWithDefaults instantiates a new CreateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *CreateKeyRequest) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *CreateKeyRequest) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *CreateKeyRequest) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *CreateKeyRequest) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *CreateKeyRequest) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *CreateKeyRequest) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetCaptureAgentTelemetry

`func (o *CreateKeyRequest) GetCaptureAgentTelemetry() bool`

GetCaptureAgentTelemetry returns the CaptureAgentTelemetry field if non-nil, zero value otherwise.

### GetCaptureAgentTelemetryOk

`func (o *CreateKeyRequest) GetCaptureAgentTelemetryOk() (*bool, bool)`

GetCaptureAgentTelemetryOk returns a tuple with the CaptureAgentTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAgentTelemetry

`func (o *CreateKeyRequest) SetCaptureAgentTelemetry(v bool)`

SetCaptureAgentTelemetry sets CaptureAgentTelemetry field to given value.

### HasCaptureAgentTelemetry

`func (o *CreateKeyRequest) HasCaptureAgentTelemetry() bool`

HasCaptureAgentTelemetry returns a boolean if a field has been set.

### SetCaptureAgentTelemetryNil

`func (o *CreateKeyRequest) SetCaptureAgentTelemetryNil(b bool)`

 SetCaptureAgentTelemetryNil sets the value for CaptureAgentTelemetry to be an explicit nil

### UnsetCaptureAgentTelemetry
`func (o *CreateKeyRequest) UnsetCaptureAgentTelemetry()`

UnsetCaptureAgentTelemetry ensures that no value is present for CaptureAgentTelemetry, not even an explicit nil
### GetEndUserBudgetId

`func (o *CreateKeyRequest) GetEndUserBudgetId() string`

GetEndUserBudgetId returns the EndUserBudgetId field if non-nil, zero value otherwise.

### GetEndUserBudgetIdOk

`func (o *CreateKeyRequest) GetEndUserBudgetIdOk() (*string, bool)`

GetEndUserBudgetIdOk returns a tuple with the EndUserBudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserBudgetId

`func (o *CreateKeyRequest) SetEndUserBudgetId(v string)`

SetEndUserBudgetId sets EndUserBudgetId field to given value.

### HasEndUserBudgetId

`func (o *CreateKeyRequest) HasEndUserBudgetId() bool`

HasEndUserBudgetId returns a boolean if a field has been set.

### SetEndUserBudgetIdNil

`func (o *CreateKeyRequest) SetEndUserBudgetIdNil(b bool)`

 SetEndUserBudgetIdNil sets the value for EndUserBudgetId to be an explicit nil

### UnsetEndUserBudgetId
`func (o *CreateKeyRequest) UnsetEndUserBudgetId()`

UnsetEndUserBudgetId ensures that no value is present for EndUserBudgetId, not even an explicit nil
### GetEndUserBudgetIds

`func (o *CreateKeyRequest) GetEndUserBudgetIds() []string`

GetEndUserBudgetIds returns the EndUserBudgetIds field if non-nil, zero value otherwise.

### GetEndUserBudgetIdsOk

`func (o *CreateKeyRequest) GetEndUserBudgetIdsOk() (*[]string, bool)`

GetEndUserBudgetIdsOk returns a tuple with the EndUserBudgetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserBudgetIds

`func (o *CreateKeyRequest) SetEndUserBudgetIds(v []string)`

SetEndUserBudgetIds sets EndUserBudgetIds field to given value.

### HasEndUserBudgetIds

`func (o *CreateKeyRequest) HasEndUserBudgetIds() bool`

HasEndUserBudgetIds returns a boolean if a field has been set.

### SetEndUserBudgetIdsNil

`func (o *CreateKeyRequest) SetEndUserBudgetIdsNil(b bool)`

 SetEndUserBudgetIdsNil sets the value for EndUserBudgetIds to be an explicit nil

### UnsetEndUserBudgetIds
`func (o *CreateKeyRequest) UnsetEndUserBudgetIds()`

UnsetEndUserBudgetIds ensures that no value is present for EndUserBudgetIds, not even an explicit nil
### GetExcludeFromBudget

`func (o *CreateKeyRequest) GetExcludeFromBudget() bool`

GetExcludeFromBudget returns the ExcludeFromBudget field if non-nil, zero value otherwise.

### GetExcludeFromBudgetOk

`func (o *CreateKeyRequest) GetExcludeFromBudgetOk() (*bool, bool)`

GetExcludeFromBudgetOk returns a tuple with the ExcludeFromBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBudget

`func (o *CreateKeyRequest) SetExcludeFromBudget(v bool)`

SetExcludeFromBudget sets ExcludeFromBudget field to given value.

### HasExcludeFromBudget

`func (o *CreateKeyRequest) HasExcludeFromBudget() bool`

HasExcludeFromBudget returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateKeyRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateKeyRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateKeyRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CreateKeyRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CreateKeyRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsServiceKey

`func (o *CreateKeyRequest) GetIsServiceKey() bool`

GetIsServiceKey returns the IsServiceKey field if non-nil, zero value otherwise.

### GetIsServiceKeyOk

`func (o *CreateKeyRequest) GetIsServiceKeyOk() (*bool, bool)`

GetIsServiceKeyOk returns a tuple with the IsServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceKey

`func (o *CreateKeyRequest) SetIsServiceKey(v bool)`

SetIsServiceKey sets IsServiceKey field to given value.

### HasIsServiceKey

`func (o *CreateKeyRequest) HasIsServiceKey() bool`

HasIsServiceKey returns a boolean if a field has been set.

### GetKeyName

`func (o *CreateKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *CreateKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *CreateKeyRequest) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *CreateKeyRequest) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetMetadata

`func (o *CreateKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRejectUserMismatch

`func (o *CreateKeyRequest) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *CreateKeyRequest) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *CreateKeyRequest) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.

### HasRejectUserMismatch

`func (o *CreateKeyRequest) HasRejectUserMismatch() bool`

HasRejectUserMismatch returns a boolean if a field has been set.

### SetRejectUserMismatchNil

`func (o *CreateKeyRequest) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *CreateKeyRequest) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil
### GetUserId

`func (o *CreateKeyRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateKeyRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateKeyRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateKeyRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateKeyRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateKeyRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *CreateKeyRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CreateKeyRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CreateKeyRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *CreateKeyRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *CreateKeyRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *CreateKeyRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


