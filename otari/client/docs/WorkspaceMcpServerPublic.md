# WorkspaceMcpServerPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | **[]string** |  | 
**AttachByDefault** | **bool** | Whether every model request from the workspace&#39;s keys gets this server without naming it | 
**CallerAccess** | Pointer to **NullableString** | On a read of one server or of the list: what the access rules give the caller&#39;s own user. &#39;all&#39; every tool, &#39;read&#39; the read tools only, &#39;some&#39; a named subset, &#39;none&#39; nothing. A key with a rule of its own can differ | [optional] 
**ConnectedUsers** | Pointer to **NullableInt32** | On a per_user server, for a caller who may manage the workspace: how many users have connected their own account. Null on a shared server, and for a caller who may not manage | [optional] 
**CreatedAt** | **string** |  | 
**CredentialMode** | [**McpCredentialMode**](McpCredentialMode.md) | &#39;shared&#39;: one workspace credential for every caller. &#39;per_user&#39;: each user&#39;s own connection | 
**Enabled** | **bool** |  | 
**EveryoneTools** | **[]string** | What every key of the workspace may call; [] denies everyone | 
**HasOauthClientSecret** | **bool** |  | 
**HasToken** | **bool** |  | 
**Health** | [**NullableWorkspaceMcpServerHealth**](WorkspaceMcpServerHealth.md) | What the gateway last learned by talking to the server, from a check or from a real request; null until either has happened | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NeedsSetup** | Pointer to **bool** | A shared server with no token and no shared account, whose server asked for a credential when last reached. Requests skip it, as they skip a disabled server, until it has one | [optional] [default to false]
**OauthClientId** | **NullableString** | The OAuth client ID a manager supplied for sign-ins; null when Otari registers itself | 
**PurposeHint** | **NullableString** |  | 
**ReadTools** | **[]string** | The tool names &#39;@read&#39; stands for: a manager&#39;s labels, or until one edits them the tools the server marked readOnlyHint when its tools were first listed. A later listing never adds to them | 
**ReadToolsFromServer** | **bool** | True while no manager has edited read_tools, so they are the readOnlyHint of the server&#39;s first listing | 
**SharedAccount** | Pointer to [**NullableMcpSharedAccountPublic**](McpSharedAccountPublic.md) | On a shared server: the account a manager signed in to for everyone, which every call carries in place of a token; null when it has none | [optional] 
**UpdatedAt** | **string** |  | 
**Url** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceMcpServerPublic

`func NewWorkspaceMcpServerPublic(allowedTools []string, attachByDefault bool, createdAt string, credentialMode McpCredentialMode, enabled bool, everyoneTools []string, hasOauthClientSecret bool, hasToken bool, health NullableWorkspaceMcpServerHealth, id string, name string, oauthClientId NullableString, purposeHint NullableString, readTools []string, readToolsFromServer bool, updatedAt string, url string, workspaceId string, ) *WorkspaceMcpServerPublic`

NewWorkspaceMcpServerPublic instantiates a new WorkspaceMcpServerPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerPublicWithDefaults

`func NewWorkspaceMcpServerPublicWithDefaults() *WorkspaceMcpServerPublic`

NewWorkspaceMcpServerPublicWithDefaults instantiates a new WorkspaceMcpServerPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerPublic) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerPublic) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerPublic) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.


### SetAllowedToolsNil

`func (o *WorkspaceMcpServerPublic) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerPublic) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetAttachByDefault

`func (o *WorkspaceMcpServerPublic) GetAttachByDefault() bool`

GetAttachByDefault returns the AttachByDefault field if non-nil, zero value otherwise.

### GetAttachByDefaultOk

`func (o *WorkspaceMcpServerPublic) GetAttachByDefaultOk() (*bool, bool)`

GetAttachByDefaultOk returns a tuple with the AttachByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachByDefault

`func (o *WorkspaceMcpServerPublic) SetAttachByDefault(v bool)`

SetAttachByDefault sets AttachByDefault field to given value.


### GetCallerAccess

`func (o *WorkspaceMcpServerPublic) GetCallerAccess() string`

GetCallerAccess returns the CallerAccess field if non-nil, zero value otherwise.

### GetCallerAccessOk

`func (o *WorkspaceMcpServerPublic) GetCallerAccessOk() (*string, bool)`

GetCallerAccessOk returns a tuple with the CallerAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerAccess

`func (o *WorkspaceMcpServerPublic) SetCallerAccess(v string)`

SetCallerAccess sets CallerAccess field to given value.

### HasCallerAccess

`func (o *WorkspaceMcpServerPublic) HasCallerAccess() bool`

HasCallerAccess returns a boolean if a field has been set.

### SetCallerAccessNil

`func (o *WorkspaceMcpServerPublic) SetCallerAccessNil(b bool)`

 SetCallerAccessNil sets the value for CallerAccess to be an explicit nil

### UnsetCallerAccess
`func (o *WorkspaceMcpServerPublic) UnsetCallerAccess()`

UnsetCallerAccess ensures that no value is present for CallerAccess, not even an explicit nil
### GetConnectedUsers

`func (o *WorkspaceMcpServerPublic) GetConnectedUsers() int32`

GetConnectedUsers returns the ConnectedUsers field if non-nil, zero value otherwise.

### GetConnectedUsersOk

`func (o *WorkspaceMcpServerPublic) GetConnectedUsersOk() (*int32, bool)`

GetConnectedUsersOk returns a tuple with the ConnectedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedUsers

`func (o *WorkspaceMcpServerPublic) SetConnectedUsers(v int32)`

SetConnectedUsers sets ConnectedUsers field to given value.

### HasConnectedUsers

`func (o *WorkspaceMcpServerPublic) HasConnectedUsers() bool`

HasConnectedUsers returns a boolean if a field has been set.

### SetConnectedUsersNil

`func (o *WorkspaceMcpServerPublic) SetConnectedUsersNil(b bool)`

 SetConnectedUsersNil sets the value for ConnectedUsers to be an explicit nil

### UnsetConnectedUsers
`func (o *WorkspaceMcpServerPublic) UnsetConnectedUsers()`

UnsetConnectedUsers ensures that no value is present for ConnectedUsers, not even an explicit nil
### GetCreatedAt

`func (o *WorkspaceMcpServerPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMcpServerPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMcpServerPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCredentialMode

`func (o *WorkspaceMcpServerPublic) GetCredentialMode() McpCredentialMode`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *WorkspaceMcpServerPublic) GetCredentialModeOk() (*McpCredentialMode, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *WorkspaceMcpServerPublic) SetCredentialMode(v McpCredentialMode)`

SetCredentialMode sets CredentialMode field to given value.


### GetEnabled

`func (o *WorkspaceMcpServerPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEveryoneTools

`func (o *WorkspaceMcpServerPublic) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *WorkspaceMcpServerPublic) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *WorkspaceMcpServerPublic) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.


### GetHasOauthClientSecret

`func (o *WorkspaceMcpServerPublic) GetHasOauthClientSecret() bool`

GetHasOauthClientSecret returns the HasOauthClientSecret field if non-nil, zero value otherwise.

### GetHasOauthClientSecretOk

`func (o *WorkspaceMcpServerPublic) GetHasOauthClientSecretOk() (*bool, bool)`

GetHasOauthClientSecretOk returns a tuple with the HasOauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOauthClientSecret

`func (o *WorkspaceMcpServerPublic) SetHasOauthClientSecret(v bool)`

SetHasOauthClientSecret sets HasOauthClientSecret field to given value.


### GetHasToken

`func (o *WorkspaceMcpServerPublic) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *WorkspaceMcpServerPublic) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *WorkspaceMcpServerPublic) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.


### GetHealth

`func (o *WorkspaceMcpServerPublic) GetHealth() WorkspaceMcpServerHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *WorkspaceMcpServerPublic) GetHealthOk() (*WorkspaceMcpServerHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *WorkspaceMcpServerPublic) SetHealth(v WorkspaceMcpServerHealth)`

SetHealth sets Health field to given value.


### SetHealthNil

`func (o *WorkspaceMcpServerPublic) SetHealthNil(b bool)`

 SetHealthNil sets the value for Health to be an explicit nil

### UnsetHealth
`func (o *WorkspaceMcpServerPublic) UnsetHealth()`

UnsetHealth ensures that no value is present for Health, not even an explicit nil
### GetId

`func (o *WorkspaceMcpServerPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMcpServerPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMcpServerPublic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkspaceMcpServerPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceMcpServerPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceMcpServerPublic) SetName(v string)`

SetName sets Name field to given value.


### GetNeedsSetup

`func (o *WorkspaceMcpServerPublic) GetNeedsSetup() bool`

GetNeedsSetup returns the NeedsSetup field if non-nil, zero value otherwise.

### GetNeedsSetupOk

`func (o *WorkspaceMcpServerPublic) GetNeedsSetupOk() (*bool, bool)`

GetNeedsSetupOk returns a tuple with the NeedsSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsSetup

`func (o *WorkspaceMcpServerPublic) SetNeedsSetup(v bool)`

SetNeedsSetup sets NeedsSetup field to given value.

### HasNeedsSetup

`func (o *WorkspaceMcpServerPublic) HasNeedsSetup() bool`

HasNeedsSetup returns a boolean if a field has been set.

### GetOauthClientId

`func (o *WorkspaceMcpServerPublic) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *WorkspaceMcpServerPublic) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *WorkspaceMcpServerPublic) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### SetOauthClientIdNil

`func (o *WorkspaceMcpServerPublic) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *WorkspaceMcpServerPublic) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetPurposeHint

`func (o *WorkspaceMcpServerPublic) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceMcpServerPublic) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceMcpServerPublic) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.


### SetPurposeHintNil

`func (o *WorkspaceMcpServerPublic) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceMcpServerPublic) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetReadTools

`func (o *WorkspaceMcpServerPublic) GetReadTools() []string`

GetReadTools returns the ReadTools field if non-nil, zero value otherwise.

### GetReadToolsOk

`func (o *WorkspaceMcpServerPublic) GetReadToolsOk() (*[]string, bool)`

GetReadToolsOk returns a tuple with the ReadTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTools

`func (o *WorkspaceMcpServerPublic) SetReadTools(v []string)`

SetReadTools sets ReadTools field to given value.


### GetReadToolsFromServer

`func (o *WorkspaceMcpServerPublic) GetReadToolsFromServer() bool`

GetReadToolsFromServer returns the ReadToolsFromServer field if non-nil, zero value otherwise.

### GetReadToolsFromServerOk

`func (o *WorkspaceMcpServerPublic) GetReadToolsFromServerOk() (*bool, bool)`

GetReadToolsFromServerOk returns a tuple with the ReadToolsFromServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToolsFromServer

`func (o *WorkspaceMcpServerPublic) SetReadToolsFromServer(v bool)`

SetReadToolsFromServer sets ReadToolsFromServer field to given value.


### GetSharedAccount

`func (o *WorkspaceMcpServerPublic) GetSharedAccount() McpSharedAccountPublic`

GetSharedAccount returns the SharedAccount field if non-nil, zero value otherwise.

### GetSharedAccountOk

`func (o *WorkspaceMcpServerPublic) GetSharedAccountOk() (*McpSharedAccountPublic, bool)`

GetSharedAccountOk returns a tuple with the SharedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccount

`func (o *WorkspaceMcpServerPublic) SetSharedAccount(v McpSharedAccountPublic)`

SetSharedAccount sets SharedAccount field to given value.

### HasSharedAccount

`func (o *WorkspaceMcpServerPublic) HasSharedAccount() bool`

HasSharedAccount returns a boolean if a field has been set.

### SetSharedAccountNil

`func (o *WorkspaceMcpServerPublic) SetSharedAccountNil(b bool)`

 SetSharedAccountNil sets the value for SharedAccount to be an explicit nil

### UnsetSharedAccount
`func (o *WorkspaceMcpServerPublic) UnsetSharedAccount()`

UnsetSharedAccount ensures that no value is present for SharedAccount, not even an explicit nil
### GetUpdatedAt

`func (o *WorkspaceMcpServerPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceMcpServerPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceMcpServerPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *WorkspaceMcpServerPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkspaceMcpServerPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkspaceMcpServerPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWorkspaceId

`func (o *WorkspaceMcpServerPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceMcpServerPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceMcpServerPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


