# WorkspaceMcpServerHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | **string** |  | 
**Error** | Pointer to [**NullableCheckFailure**](CheckFailure.md) | Why the check failed; null when healthy | [optional] 
**HttpStatus** | Pointer to **NullableInt32** | The HTTP status the server refused with, if any | [optional] 
**LastFailure** | Pointer to [**NullableWorkspaceMcpServerLastFailure**](WorkspaceMcpServerLastFailure.md) | The most recent contact that got no answer, however long ago; null if none is known | [optional] 
**MissingAllowedTools** | Pointer to **[]string** | Allow-listed names the server did not offer | [optional] 
**ReadOnlyHintedTools** | Pointer to **[]string** | Those of tools the server annotates readOnlyHint in this listing: untrusted, and never what &#39;@read&#39; stands for after the server&#39;s first listing | [optional] 
**Status** | **string** | failing: the last contact got no answer. degraded: the last contact got one, but another failed within the last 60 minutes. healthy: neither | 
**ToolCount** | Pointer to **NullableInt32** | Live tools the allow-list admits; null when the check failed, or when the server answered a request that did not list its tools and no listing is on record | [optional] 
**ToolDescriptions** | Pointer to **map[string]string** | The first line of each offered tool&#39;s own description, where it has one: untrusted text the server wrote, for a person to read | [optional] 
**Tools** | Pointer to **[]string** | Every tool name the server offers, allowed or not, up to 200; empty on a failure | [optional] 

## Methods

### NewWorkspaceMcpServerHealth

`func NewWorkspaceMcpServerHealth(checkedAt string, status string, ) *WorkspaceMcpServerHealth`

NewWorkspaceMcpServerHealth instantiates a new WorkspaceMcpServerHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerHealthWithDefaults

`func NewWorkspaceMcpServerHealthWithDefaults() *WorkspaceMcpServerHealth`

NewWorkspaceMcpServerHealthWithDefaults instantiates a new WorkspaceMcpServerHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *WorkspaceMcpServerHealth) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *WorkspaceMcpServerHealth) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *WorkspaceMcpServerHealth) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.


### GetError

`func (o *WorkspaceMcpServerHealth) GetError() CheckFailure`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkspaceMcpServerHealth) GetErrorOk() (*CheckFailure, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkspaceMcpServerHealth) SetError(v CheckFailure)`

SetError sets Error field to given value.

### HasError

`func (o *WorkspaceMcpServerHealth) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WorkspaceMcpServerHealth) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WorkspaceMcpServerHealth) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHttpStatus

`func (o *WorkspaceMcpServerHealth) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *WorkspaceMcpServerHealth) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *WorkspaceMcpServerHealth) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *WorkspaceMcpServerHealth) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### SetHttpStatusNil

`func (o *WorkspaceMcpServerHealth) SetHttpStatusNil(b bool)`

 SetHttpStatusNil sets the value for HttpStatus to be an explicit nil

### UnsetHttpStatus
`func (o *WorkspaceMcpServerHealth) UnsetHttpStatus()`

UnsetHttpStatus ensures that no value is present for HttpStatus, not even an explicit nil
### GetLastFailure

`func (o *WorkspaceMcpServerHealth) GetLastFailure() WorkspaceMcpServerLastFailure`

GetLastFailure returns the LastFailure field if non-nil, zero value otherwise.

### GetLastFailureOk

`func (o *WorkspaceMcpServerHealth) GetLastFailureOk() (*WorkspaceMcpServerLastFailure, bool)`

GetLastFailureOk returns a tuple with the LastFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailure

`func (o *WorkspaceMcpServerHealth) SetLastFailure(v WorkspaceMcpServerLastFailure)`

SetLastFailure sets LastFailure field to given value.

### HasLastFailure

`func (o *WorkspaceMcpServerHealth) HasLastFailure() bool`

HasLastFailure returns a boolean if a field has been set.

### SetLastFailureNil

`func (o *WorkspaceMcpServerHealth) SetLastFailureNil(b bool)`

 SetLastFailureNil sets the value for LastFailure to be an explicit nil

### UnsetLastFailure
`func (o *WorkspaceMcpServerHealth) UnsetLastFailure()`

UnsetLastFailure ensures that no value is present for LastFailure, not even an explicit nil
### GetMissingAllowedTools

`func (o *WorkspaceMcpServerHealth) GetMissingAllowedTools() []string`

GetMissingAllowedTools returns the MissingAllowedTools field if non-nil, zero value otherwise.

### GetMissingAllowedToolsOk

`func (o *WorkspaceMcpServerHealth) GetMissingAllowedToolsOk() (*[]string, bool)`

GetMissingAllowedToolsOk returns a tuple with the MissingAllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAllowedTools

`func (o *WorkspaceMcpServerHealth) SetMissingAllowedTools(v []string)`

SetMissingAllowedTools sets MissingAllowedTools field to given value.

### HasMissingAllowedTools

`func (o *WorkspaceMcpServerHealth) HasMissingAllowedTools() bool`

HasMissingAllowedTools returns a boolean if a field has been set.

### GetReadOnlyHintedTools

`func (o *WorkspaceMcpServerHealth) GetReadOnlyHintedTools() []string`

GetReadOnlyHintedTools returns the ReadOnlyHintedTools field if non-nil, zero value otherwise.

### GetReadOnlyHintedToolsOk

`func (o *WorkspaceMcpServerHealth) GetReadOnlyHintedToolsOk() (*[]string, bool)`

GetReadOnlyHintedToolsOk returns a tuple with the ReadOnlyHintedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyHintedTools

`func (o *WorkspaceMcpServerHealth) SetReadOnlyHintedTools(v []string)`

SetReadOnlyHintedTools sets ReadOnlyHintedTools field to given value.

### HasReadOnlyHintedTools

`func (o *WorkspaceMcpServerHealth) HasReadOnlyHintedTools() bool`

HasReadOnlyHintedTools returns a boolean if a field has been set.

### GetStatus

`func (o *WorkspaceMcpServerHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkspaceMcpServerHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkspaceMcpServerHealth) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToolCount

`func (o *WorkspaceMcpServerHealth) GetToolCount() int32`

GetToolCount returns the ToolCount field if non-nil, zero value otherwise.

### GetToolCountOk

`func (o *WorkspaceMcpServerHealth) GetToolCountOk() (*int32, bool)`

GetToolCountOk returns a tuple with the ToolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCount

`func (o *WorkspaceMcpServerHealth) SetToolCount(v int32)`

SetToolCount sets ToolCount field to given value.

### HasToolCount

`func (o *WorkspaceMcpServerHealth) HasToolCount() bool`

HasToolCount returns a boolean if a field has been set.

### SetToolCountNil

`func (o *WorkspaceMcpServerHealth) SetToolCountNil(b bool)`

 SetToolCountNil sets the value for ToolCount to be an explicit nil

### UnsetToolCount
`func (o *WorkspaceMcpServerHealth) UnsetToolCount()`

UnsetToolCount ensures that no value is present for ToolCount, not even an explicit nil
### GetToolDescriptions

`func (o *WorkspaceMcpServerHealth) GetToolDescriptions() map[string]string`

GetToolDescriptions returns the ToolDescriptions field if non-nil, zero value otherwise.

### GetToolDescriptionsOk

`func (o *WorkspaceMcpServerHealth) GetToolDescriptionsOk() (*map[string]string, bool)`

GetToolDescriptionsOk returns a tuple with the ToolDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolDescriptions

`func (o *WorkspaceMcpServerHealth) SetToolDescriptions(v map[string]string)`

SetToolDescriptions sets ToolDescriptions field to given value.

### HasToolDescriptions

`func (o *WorkspaceMcpServerHealth) HasToolDescriptions() bool`

HasToolDescriptions returns a boolean if a field has been set.

### GetTools

`func (o *WorkspaceMcpServerHealth) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerHealth) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerHealth) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *WorkspaceMcpServerHealth) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


