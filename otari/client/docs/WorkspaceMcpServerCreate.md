# WorkspaceMcpServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | Pointer to **[]string** | Allow-list of tool names; null exposes every tool the server offers | [optional] 
**AttachByDefault** | Pointer to **bool** | Add this server to every model request from the workspace&#39;s keys, as if the request named it in mcp_server_ids. A key the server&#39;s access rules grant nothing is left without it | [optional] [default to false]
**AuthorizationToken** | Pointer to **NullableString** | Bearer token for the server; requires an https URL. Encrypted at rest, never returned | [optional] 
**CredentialMode** | Pointer to [**McpCredentialMode**](McpCredentialMode.md) | Whose credential a call carries: &#39;shared&#39; sends one credential for every caller, authorization_token or the account a manager signs in to for everyone once the server is stored; &#39;per_user&#39; sends the calling user&#39;s own token, which each user gets by signing in to the server, and takes no authorization_token | [optional] [default to SHARED]
**Enabled** | Pointer to **bool** | Whether a request naming this server actually reaches it | [optional] [default to true]
**EveryoneTools** | Pointer to **[]string** | What every key of the workspace may call on this server: tool names, &#39;*&#39; for every tool, or &#39;prefix*&#39;, or &#39;@read&#39; for read_tools; [] denies everyone. A grant for a key or its owner overrides it | [optional] 
**FoldServerIds** | Pointer to **[]string** | On an organization&#39;s server: workspaces&#39; own servers at the same URL to fold into it. Each one&#39;s rules become its workspace&#39;s settings on the new server, its id keeps working there, and it is deleted. A personal sign-in is kept where the new server is personal at its very URL and can renew it: the sign-in holds a refresh token and was made through the new server&#39;s client. The rest end | [optional] 
**Name** | **string** | Label for the server, unique within the workspace | 
**OauthClientId** | Pointer to **NullableString** | For a server whose authorization server does not register clients on its own: the OAuth client ID a sign-in goes through. Omit it to let Otari register itself | [optional] 
**OauthClientSecret** | Pointer to **NullableString** | The secret of oauth_client_id, when it has one. Encrypted at rest, never returned | [optional] 
**PurposeHint** | Pointer to **NullableString** | Hint prepended to the system message to help the model choose | [optional] 
**ReadTools** | Pointer to **[]string** | The tool names a manager labeled read, which &#39;@read&#39; in an access rule stands for. Omit it, or send null, and the labels follow the tools the server itself marks readOnlyHint | [optional] 
**Url** | **string** | Streamable HTTP MCP endpoint | 

## Methods

### NewWorkspaceMcpServerCreate

`func NewWorkspaceMcpServerCreate(name string, url string, ) *WorkspaceMcpServerCreate`

NewWorkspaceMcpServerCreate instantiates a new WorkspaceMcpServerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerCreateWithDefaults

`func NewWorkspaceMcpServerCreateWithDefaults() *WorkspaceMcpServerCreate`

NewWorkspaceMcpServerCreateWithDefaults instantiates a new WorkspaceMcpServerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerCreate) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerCreate) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerCreate) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *WorkspaceMcpServerCreate) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### SetAllowedToolsNil

`func (o *WorkspaceMcpServerCreate) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerCreate) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetAttachByDefault

`func (o *WorkspaceMcpServerCreate) GetAttachByDefault() bool`

GetAttachByDefault returns the AttachByDefault field if non-nil, zero value otherwise.

### GetAttachByDefaultOk

`func (o *WorkspaceMcpServerCreate) GetAttachByDefaultOk() (*bool, bool)`

GetAttachByDefaultOk returns a tuple with the AttachByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachByDefault

`func (o *WorkspaceMcpServerCreate) SetAttachByDefault(v bool)`

SetAttachByDefault sets AttachByDefault field to given value.

### HasAttachByDefault

`func (o *WorkspaceMcpServerCreate) HasAttachByDefault() bool`

HasAttachByDefault returns a boolean if a field has been set.

### GetAuthorizationToken

`func (o *WorkspaceMcpServerCreate) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *WorkspaceMcpServerCreate) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *WorkspaceMcpServerCreate) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *WorkspaceMcpServerCreate) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### SetAuthorizationTokenNil

`func (o *WorkspaceMcpServerCreate) SetAuthorizationTokenNil(b bool)`

 SetAuthorizationTokenNil sets the value for AuthorizationToken to be an explicit nil

### UnsetAuthorizationToken
`func (o *WorkspaceMcpServerCreate) UnsetAuthorizationToken()`

UnsetAuthorizationToken ensures that no value is present for AuthorizationToken, not even an explicit nil
### GetCredentialMode

`func (o *WorkspaceMcpServerCreate) GetCredentialMode() McpCredentialMode`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *WorkspaceMcpServerCreate) GetCredentialModeOk() (*McpCredentialMode, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *WorkspaceMcpServerCreate) SetCredentialMode(v McpCredentialMode)`

SetCredentialMode sets CredentialMode field to given value.

### HasCredentialMode

`func (o *WorkspaceMcpServerCreate) HasCredentialMode() bool`

HasCredentialMode returns a boolean if a field has been set.

### GetEnabled

`func (o *WorkspaceMcpServerCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkspaceMcpServerCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEveryoneTools

`func (o *WorkspaceMcpServerCreate) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *WorkspaceMcpServerCreate) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *WorkspaceMcpServerCreate) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.

### HasEveryoneTools

`func (o *WorkspaceMcpServerCreate) HasEveryoneTools() bool`

HasEveryoneTools returns a boolean if a field has been set.

### GetFoldServerIds

`func (o *WorkspaceMcpServerCreate) GetFoldServerIds() []string`

GetFoldServerIds returns the FoldServerIds field if non-nil, zero value otherwise.

### GetFoldServerIdsOk

`func (o *WorkspaceMcpServerCreate) GetFoldServerIdsOk() (*[]string, bool)`

GetFoldServerIdsOk returns a tuple with the FoldServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldServerIds

`func (o *WorkspaceMcpServerCreate) SetFoldServerIds(v []string)`

SetFoldServerIds sets FoldServerIds field to given value.

### HasFoldServerIds

`func (o *WorkspaceMcpServerCreate) HasFoldServerIds() bool`

HasFoldServerIds returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceMcpServerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceMcpServerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceMcpServerCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOauthClientId

`func (o *WorkspaceMcpServerCreate) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *WorkspaceMcpServerCreate) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *WorkspaceMcpServerCreate) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *WorkspaceMcpServerCreate) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *WorkspaceMcpServerCreate) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *WorkspaceMcpServerCreate) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthClientSecret

`func (o *WorkspaceMcpServerCreate) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *WorkspaceMcpServerCreate) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *WorkspaceMcpServerCreate) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *WorkspaceMcpServerCreate) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *WorkspaceMcpServerCreate) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *WorkspaceMcpServerCreate) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetPurposeHint

`func (o *WorkspaceMcpServerCreate) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceMcpServerCreate) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceMcpServerCreate) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.

### HasPurposeHint

`func (o *WorkspaceMcpServerCreate) HasPurposeHint() bool`

HasPurposeHint returns a boolean if a field has been set.

### SetPurposeHintNil

`func (o *WorkspaceMcpServerCreate) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceMcpServerCreate) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetReadTools

`func (o *WorkspaceMcpServerCreate) GetReadTools() []string`

GetReadTools returns the ReadTools field if non-nil, zero value otherwise.

### GetReadToolsOk

`func (o *WorkspaceMcpServerCreate) GetReadToolsOk() (*[]string, bool)`

GetReadToolsOk returns a tuple with the ReadTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTools

`func (o *WorkspaceMcpServerCreate) SetReadTools(v []string)`

SetReadTools sets ReadTools field to given value.

### HasReadTools

`func (o *WorkspaceMcpServerCreate) HasReadTools() bool`

HasReadTools returns a boolean if a field has been set.

### SetReadToolsNil

`func (o *WorkspaceMcpServerCreate) SetReadToolsNil(b bool)`

 SetReadToolsNil sets the value for ReadTools to be an explicit nil

### UnsetReadTools
`func (o *WorkspaceMcpServerCreate) UnsetReadTools()`

UnsetReadTools ensures that no value is present for ReadTools, not even an explicit nil
### GetUrl

`func (o *WorkspaceMcpServerCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkspaceMcpServerCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkspaceMcpServerCreate) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


