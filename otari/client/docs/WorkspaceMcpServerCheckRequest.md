# WorkspaceMcpServerCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | Pointer to **[]string** | The allow-list to report the catalog against | [optional] 
**AuthorizationToken** | Pointer to **NullableString** | Bearer token; omit to use the stored server&#39;s, if any | [optional] 
**ServerId** | Pointer to **NullableString** | A stored server whose token to use | [optional] 
**Url** | **string** | Streamable HTTP MCP endpoint | 

## Methods

### NewWorkspaceMcpServerCheckRequest

`func NewWorkspaceMcpServerCheckRequest(url string, ) *WorkspaceMcpServerCheckRequest`

NewWorkspaceMcpServerCheckRequest instantiates a new WorkspaceMcpServerCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerCheckRequestWithDefaults

`func NewWorkspaceMcpServerCheckRequestWithDefaults() *WorkspaceMcpServerCheckRequest`

NewWorkspaceMcpServerCheckRequestWithDefaults instantiates a new WorkspaceMcpServerCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerCheckRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerCheckRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerCheckRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *WorkspaceMcpServerCheckRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### SetAllowedToolsNil

`func (o *WorkspaceMcpServerCheckRequest) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerCheckRequest) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetAuthorizationToken

`func (o *WorkspaceMcpServerCheckRequest) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *WorkspaceMcpServerCheckRequest) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *WorkspaceMcpServerCheckRequest) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *WorkspaceMcpServerCheckRequest) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### SetAuthorizationTokenNil

`func (o *WorkspaceMcpServerCheckRequest) SetAuthorizationTokenNil(b bool)`

 SetAuthorizationTokenNil sets the value for AuthorizationToken to be an explicit nil

### UnsetAuthorizationToken
`func (o *WorkspaceMcpServerCheckRequest) UnsetAuthorizationToken()`

UnsetAuthorizationToken ensures that no value is present for AuthorizationToken, not even an explicit nil
### GetServerId

`func (o *WorkspaceMcpServerCheckRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *WorkspaceMcpServerCheckRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *WorkspaceMcpServerCheckRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *WorkspaceMcpServerCheckRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *WorkspaceMcpServerCheckRequest) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *WorkspaceMcpServerCheckRequest) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUrl

`func (o *WorkspaceMcpServerCheckRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkspaceMcpServerCheckRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkspaceMcpServerCheckRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


