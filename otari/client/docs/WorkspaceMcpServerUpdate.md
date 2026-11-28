# WorkspaceMcpServerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AttachByDefault** | Pointer to **bool** |  | [optional] 
**AuthorizationToken** | Pointer to **NullableString** |  | [optional] 
**CredentialMode** | Pointer to [**McpCredentialMode**](McpCredentialMode.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EveryoneTools** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OauthClientId** | Pointer to **NullableString** |  | [optional] 
**OauthClientSecret** | Pointer to **NullableString** |  | [optional] 
**PurposeHint** | Pointer to **NullableString** |  | [optional] 
**ReadTools** | Pointer to **[]string** | The tool names a manager labeled read. An explicit null goes back to following the tools the server marks readOnlyHint; omit it to leave the labels as they are | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkspaceMcpServerUpdate

`func NewWorkspaceMcpServerUpdate() *WorkspaceMcpServerUpdate`

NewWorkspaceMcpServerUpdate instantiates a new WorkspaceMcpServerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerUpdateWithDefaults

`func NewWorkspaceMcpServerUpdateWithDefaults() *WorkspaceMcpServerUpdate`

NewWorkspaceMcpServerUpdateWithDefaults instantiates a new WorkspaceMcpServerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerUpdate) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerUpdate) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerUpdate) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *WorkspaceMcpServerUpdate) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### SetAllowedToolsNil

`func (o *WorkspaceMcpServerUpdate) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerUpdate) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetAttachByDefault

`func (o *WorkspaceMcpServerUpdate) GetAttachByDefault() bool`

GetAttachByDefault returns the AttachByDefault field if non-nil, zero value otherwise.

### GetAttachByDefaultOk

`func (o *WorkspaceMcpServerUpdate) GetAttachByDefaultOk() (*bool, bool)`

GetAttachByDefaultOk returns a tuple with the AttachByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachByDefault

`func (o *WorkspaceMcpServerUpdate) SetAttachByDefault(v bool)`

SetAttachByDefault sets AttachByDefault field to given value.

### HasAttachByDefault

`func (o *WorkspaceMcpServerUpdate) HasAttachByDefault() bool`

HasAttachByDefault returns a boolean if a field has been set.

### GetAuthorizationToken

`func (o *WorkspaceMcpServerUpdate) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *WorkspaceMcpServerUpdate) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *WorkspaceMcpServerUpdate) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *WorkspaceMcpServerUpdate) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### SetAuthorizationTokenNil

`func (o *WorkspaceMcpServerUpdate) SetAuthorizationTokenNil(b bool)`

 SetAuthorizationTokenNil sets the value for AuthorizationToken to be an explicit nil

### UnsetAuthorizationToken
`func (o *WorkspaceMcpServerUpdate) UnsetAuthorizationToken()`

UnsetAuthorizationToken ensures that no value is present for AuthorizationToken, not even an explicit nil
### GetCredentialMode

`func (o *WorkspaceMcpServerUpdate) GetCredentialMode() McpCredentialMode`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *WorkspaceMcpServerUpdate) GetCredentialModeOk() (*McpCredentialMode, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *WorkspaceMcpServerUpdate) SetCredentialMode(v McpCredentialMode)`

SetCredentialMode sets CredentialMode field to given value.

### HasCredentialMode

`func (o *WorkspaceMcpServerUpdate) HasCredentialMode() bool`

HasCredentialMode returns a boolean if a field has been set.

### GetEnabled

`func (o *WorkspaceMcpServerUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkspaceMcpServerUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEveryoneTools

`func (o *WorkspaceMcpServerUpdate) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *WorkspaceMcpServerUpdate) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *WorkspaceMcpServerUpdate) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.

### HasEveryoneTools

`func (o *WorkspaceMcpServerUpdate) HasEveryoneTools() bool`

HasEveryoneTools returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceMcpServerUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceMcpServerUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceMcpServerUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceMcpServerUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOauthClientId

`func (o *WorkspaceMcpServerUpdate) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *WorkspaceMcpServerUpdate) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *WorkspaceMcpServerUpdate) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *WorkspaceMcpServerUpdate) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *WorkspaceMcpServerUpdate) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *WorkspaceMcpServerUpdate) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthClientSecret

`func (o *WorkspaceMcpServerUpdate) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *WorkspaceMcpServerUpdate) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *WorkspaceMcpServerUpdate) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *WorkspaceMcpServerUpdate) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *WorkspaceMcpServerUpdate) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *WorkspaceMcpServerUpdate) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetPurposeHint

`func (o *WorkspaceMcpServerUpdate) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceMcpServerUpdate) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceMcpServerUpdate) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.

### HasPurposeHint

`func (o *WorkspaceMcpServerUpdate) HasPurposeHint() bool`

HasPurposeHint returns a boolean if a field has been set.

### SetPurposeHintNil

`func (o *WorkspaceMcpServerUpdate) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceMcpServerUpdate) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetReadTools

`func (o *WorkspaceMcpServerUpdate) GetReadTools() []string`

GetReadTools returns the ReadTools field if non-nil, zero value otherwise.

### GetReadToolsOk

`func (o *WorkspaceMcpServerUpdate) GetReadToolsOk() (*[]string, bool)`

GetReadToolsOk returns a tuple with the ReadTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTools

`func (o *WorkspaceMcpServerUpdate) SetReadTools(v []string)`

SetReadTools sets ReadTools field to given value.

### HasReadTools

`func (o *WorkspaceMcpServerUpdate) HasReadTools() bool`

HasReadTools returns a boolean if a field has been set.

### SetReadToolsNil

`func (o *WorkspaceMcpServerUpdate) SetReadToolsNil(b bool)`

 SetReadToolsNil sets the value for ReadTools to be an explicit nil

### UnsetReadTools
`func (o *WorkspaceMcpServerUpdate) UnsetReadTools()`

UnsetReadTools ensures that no value is present for ReadTools, not even an explicit nil
### GetUrl

`func (o *WorkspaceMcpServerUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkspaceMcpServerUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkspaceMcpServerUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WorkspaceMcpServerUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


