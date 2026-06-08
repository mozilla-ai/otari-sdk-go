# McpServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AuthorizationToken** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**PurposeHint** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewMcpServerConfig

`func NewMcpServerConfig(name string, url string, ) *McpServerConfig`

NewMcpServerConfig instantiates a new McpServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerConfigWithDefaults

`func NewMcpServerConfigWithDefaults() *McpServerConfig`

NewMcpServerConfigWithDefaults instantiates a new McpServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *McpServerConfig) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *McpServerConfig) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *McpServerConfig) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *McpServerConfig) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### SetAllowedToolsNil

`func (o *McpServerConfig) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *McpServerConfig) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetAuthorizationToken

`func (o *McpServerConfig) GetAuthorizationToken() string`

GetAuthorizationToken returns the AuthorizationToken field if non-nil, zero value otherwise.

### GetAuthorizationTokenOk

`func (o *McpServerConfig) GetAuthorizationTokenOk() (*string, bool)`

GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationToken

`func (o *McpServerConfig) SetAuthorizationToken(v string)`

SetAuthorizationToken sets AuthorizationToken field to given value.

### HasAuthorizationToken

`func (o *McpServerConfig) HasAuthorizationToken() bool`

HasAuthorizationToken returns a boolean if a field has been set.

### SetAuthorizationTokenNil

`func (o *McpServerConfig) SetAuthorizationTokenNil(b bool)`

 SetAuthorizationTokenNil sets the value for AuthorizationToken to be an explicit nil

### UnsetAuthorizationToken
`func (o *McpServerConfig) UnsetAuthorizationToken()`

UnsetAuthorizationToken ensures that no value is present for AuthorizationToken, not even an explicit nil
### GetName

`func (o *McpServerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetPurposeHint

`func (o *McpServerConfig) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *McpServerConfig) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *McpServerConfig) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.

### HasPurposeHint

`func (o *McpServerConfig) HasPurposeHint() bool`

HasPurposeHint returns a boolean if a field has been set.

### SetPurposeHintNil

`func (o *McpServerConfig) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *McpServerConfig) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetUrl

`func (o *McpServerConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerConfig) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


