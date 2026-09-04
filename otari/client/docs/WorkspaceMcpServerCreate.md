# WorkspaceMcpServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | Pointer to **[]string** | Allow-list of tool names; null exposes every tool the server offers | [optional] 
**AuthorizationToken** | Pointer to **NullableString** | Bearer token for the server; requires an https URL. Encrypted at rest, never returned | [optional] 
**Enabled** | Pointer to **bool** | Whether a request naming this server actually reaches it | [optional] [default to true]
**Name** | **string** | Label for the server, unique within the workspace | 
**PurposeHint** | Pointer to **NullableString** | Hint prepended to the system message to help the model choose | [optional] 
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


