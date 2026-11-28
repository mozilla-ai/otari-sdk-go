# McpServerFoldCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialMode** | [**McpCredentialMode**](McpCredentialMode.md) |  | 
**Enabled** | Pointer to **bool** | Whether it is on. An organization server that is off cannot take the place of one that is on | [optional] [default to true]
**KeepSignIn** | **int32** | Of those, how many keep their sign-in on the organization&#39;s server | 
**Name** | **string** |  | 
**ReplacesSharedCredential** | **bool** | It holds a token or a shared account, which the organization server&#39;s credential replaces | 
**SameCredential** | **bool** | It signs in the way the organization&#39;s server does and holds no token or shared account of its own, so folding it changes no account its calls use. A fold step ticks these by default, unless someone would have to sign in again or go on through another sign-in | 
**ServerId** | **string** |  | 
**SignInAgain** | **int32** | Of those, how many must sign in again | 
**SignedIn** | **int32** | People signed in to it with their own account | 
**SwitchAccount** | Pointer to **int32** | Of those, how many already hold a sign-in on the organization&#39;s server and go on through it, which may be another account upstream than the one they used here | [optional] [default to 0]
**WorkspaceId** | **string** |  | 
**WorkspaceName** | **string** |  | 

## Methods

### NewMcpServerFoldCandidate

`func NewMcpServerFoldCandidate(credentialMode McpCredentialMode, keepSignIn int32, name string, replacesSharedCredential bool, sameCredential bool, serverId string, signInAgain int32, signedIn int32, workspaceId string, workspaceName string, ) *McpServerFoldCandidate`

NewMcpServerFoldCandidate instantiates a new McpServerFoldCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerFoldCandidateWithDefaults

`func NewMcpServerFoldCandidateWithDefaults() *McpServerFoldCandidate`

NewMcpServerFoldCandidateWithDefaults instantiates a new McpServerFoldCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialMode

`func (o *McpServerFoldCandidate) GetCredentialMode() McpCredentialMode`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *McpServerFoldCandidate) GetCredentialModeOk() (*McpCredentialMode, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *McpServerFoldCandidate) SetCredentialMode(v McpCredentialMode)`

SetCredentialMode sets CredentialMode field to given value.


### GetEnabled

`func (o *McpServerFoldCandidate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpServerFoldCandidate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpServerFoldCandidate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpServerFoldCandidate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeepSignIn

`func (o *McpServerFoldCandidate) GetKeepSignIn() int32`

GetKeepSignIn returns the KeepSignIn field if non-nil, zero value otherwise.

### GetKeepSignInOk

`func (o *McpServerFoldCandidate) GetKeepSignInOk() (*int32, bool)`

GetKeepSignInOk returns a tuple with the KeepSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepSignIn

`func (o *McpServerFoldCandidate) SetKeepSignIn(v int32)`

SetKeepSignIn sets KeepSignIn field to given value.


### GetName

`func (o *McpServerFoldCandidate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerFoldCandidate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerFoldCandidate) SetName(v string)`

SetName sets Name field to given value.


### GetReplacesSharedCredential

`func (o *McpServerFoldCandidate) GetReplacesSharedCredential() bool`

GetReplacesSharedCredential returns the ReplacesSharedCredential field if non-nil, zero value otherwise.

### GetReplacesSharedCredentialOk

`func (o *McpServerFoldCandidate) GetReplacesSharedCredentialOk() (*bool, bool)`

GetReplacesSharedCredentialOk returns a tuple with the ReplacesSharedCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacesSharedCredential

`func (o *McpServerFoldCandidate) SetReplacesSharedCredential(v bool)`

SetReplacesSharedCredential sets ReplacesSharedCredential field to given value.


### GetSameCredential

`func (o *McpServerFoldCandidate) GetSameCredential() bool`

GetSameCredential returns the SameCredential field if non-nil, zero value otherwise.

### GetSameCredentialOk

`func (o *McpServerFoldCandidate) GetSameCredentialOk() (*bool, bool)`

GetSameCredentialOk returns a tuple with the SameCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameCredential

`func (o *McpServerFoldCandidate) SetSameCredential(v bool)`

SetSameCredential sets SameCredential field to given value.


### GetServerId

`func (o *McpServerFoldCandidate) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpServerFoldCandidate) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpServerFoldCandidate) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetSignInAgain

`func (o *McpServerFoldCandidate) GetSignInAgain() int32`

GetSignInAgain returns the SignInAgain field if non-nil, zero value otherwise.

### GetSignInAgainOk

`func (o *McpServerFoldCandidate) GetSignInAgainOk() (*int32, bool)`

GetSignInAgainOk returns a tuple with the SignInAgain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInAgain

`func (o *McpServerFoldCandidate) SetSignInAgain(v int32)`

SetSignInAgain sets SignInAgain field to given value.


### GetSignedIn

`func (o *McpServerFoldCandidate) GetSignedIn() int32`

GetSignedIn returns the SignedIn field if non-nil, zero value otherwise.

### GetSignedInOk

`func (o *McpServerFoldCandidate) GetSignedInOk() (*int32, bool)`

GetSignedInOk returns a tuple with the SignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedIn

`func (o *McpServerFoldCandidate) SetSignedIn(v int32)`

SetSignedIn sets SignedIn field to given value.


### GetSwitchAccount

`func (o *McpServerFoldCandidate) GetSwitchAccount() int32`

GetSwitchAccount returns the SwitchAccount field if non-nil, zero value otherwise.

### GetSwitchAccountOk

`func (o *McpServerFoldCandidate) GetSwitchAccountOk() (*int32, bool)`

GetSwitchAccountOk returns a tuple with the SwitchAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchAccount

`func (o *McpServerFoldCandidate) SetSwitchAccount(v int32)`

SetSwitchAccount sets SwitchAccount field to given value.

### HasSwitchAccount

`func (o *McpServerFoldCandidate) HasSwitchAccount() bool`

HasSwitchAccount returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *McpServerFoldCandidate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *McpServerFoldCandidate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *McpServerFoldCandidate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetWorkspaceName

`func (o *McpServerFoldCandidate) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *McpServerFoldCandidate) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *McpServerFoldCandidate) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


