# McpConnectionCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checked** | Pointer to **bool** | The sign-in was made only to check the server: its health was recorded and nothing was connected | [optional] [default to false]
**ReturnTo** | **NullableString** | The dashboard path the sign-in was started with, if any | 
**ServerId** | **string** |  | 
**ServerName** | **string** |  | 
**Shared** | Pointer to **bool** | The sign-in was a manager&#39;s for everyone: it is now the shared server&#39;s account | [optional] [default to false]
**WorkspaceId** | **string** |  | 

## Methods

### NewMcpConnectionCompleted

`func NewMcpConnectionCompleted(returnTo NullableString, serverId string, serverName string, workspaceId string, ) *McpConnectionCompleted`

NewMcpConnectionCompleted instantiates a new McpConnectionCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionCompletedWithDefaults

`func NewMcpConnectionCompletedWithDefaults() *McpConnectionCompleted`

NewMcpConnectionCompletedWithDefaults instantiates a new McpConnectionCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecked

`func (o *McpConnectionCompleted) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *McpConnectionCompleted) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *McpConnectionCompleted) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *McpConnectionCompleted) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetReturnTo

`func (o *McpConnectionCompleted) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *McpConnectionCompleted) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *McpConnectionCompleted) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.


### SetReturnToNil

`func (o *McpConnectionCompleted) SetReturnToNil(b bool)`

 SetReturnToNil sets the value for ReturnTo to be an explicit nil

### UnsetReturnTo
`func (o *McpConnectionCompleted) UnsetReturnTo()`

UnsetReturnTo ensures that no value is present for ReturnTo, not even an explicit nil
### GetServerId

`func (o *McpConnectionCompleted) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpConnectionCompleted) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpConnectionCompleted) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetServerName

`func (o *McpConnectionCompleted) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *McpConnectionCompleted) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *McpConnectionCompleted) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetShared

`func (o *McpConnectionCompleted) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *McpConnectionCompleted) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *McpConnectionCompleted) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *McpConnectionCompleted) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *McpConnectionCompleted) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *McpConnectionCompleted) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *McpConnectionCompleted) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


