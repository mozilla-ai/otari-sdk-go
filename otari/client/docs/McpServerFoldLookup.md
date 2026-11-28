# McpServerFoldLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialMode** | Pointer to [**McpCredentialMode**](McpCredentialMode.md) | How the new server signs in, when no server_id is given | [optional] [default to SHARED]
**ServerId** | Pointer to **NullableString** | The server about to be made the organization&#39;s, whose own URL is matched, if any | [optional] 
**Url** | Pointer to **NullableString** | The URL of the server to add; not sent with server_id | [optional] 
**WithToken** | Pointer to **bool** | The new server is added with a token of its own, when no server_id is given | [optional] [default to false]

## Methods

### NewMcpServerFoldLookup

`func NewMcpServerFoldLookup() *McpServerFoldLookup`

NewMcpServerFoldLookup instantiates a new McpServerFoldLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerFoldLookupWithDefaults

`func NewMcpServerFoldLookupWithDefaults() *McpServerFoldLookup`

NewMcpServerFoldLookupWithDefaults instantiates a new McpServerFoldLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialMode

`func (o *McpServerFoldLookup) GetCredentialMode() McpCredentialMode`

GetCredentialMode returns the CredentialMode field if non-nil, zero value otherwise.

### GetCredentialModeOk

`func (o *McpServerFoldLookup) GetCredentialModeOk() (*McpCredentialMode, bool)`

GetCredentialModeOk returns a tuple with the CredentialMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialMode

`func (o *McpServerFoldLookup) SetCredentialMode(v McpCredentialMode)`

SetCredentialMode sets CredentialMode field to given value.

### HasCredentialMode

`func (o *McpServerFoldLookup) HasCredentialMode() bool`

HasCredentialMode returns a boolean if a field has been set.

### GetServerId

`func (o *McpServerFoldLookup) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpServerFoldLookup) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpServerFoldLookup) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *McpServerFoldLookup) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *McpServerFoldLookup) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *McpServerFoldLookup) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUrl

`func (o *McpServerFoldLookup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerFoldLookup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerFoldLookup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *McpServerFoldLookup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *McpServerFoldLookup) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *McpServerFoldLookup) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWithToken

`func (o *McpServerFoldLookup) GetWithToken() bool`

GetWithToken returns the WithToken field if non-nil, zero value otherwise.

### GetWithTokenOk

`func (o *McpServerFoldLookup) GetWithTokenOk() (*bool, bool)`

GetWithTokenOk returns a tuple with the WithToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithToken

`func (o *McpServerFoldLookup) SetWithToken(v bool)`

SetWithToken sets WithToken field to given value.

### HasWithToken

`func (o *McpServerFoldLookup) HasWithToken() bool`

HasWithToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


