# WorkspaceMcpServerMyTools

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | &#39;all&#39; when the rule admits every tool the server offers, &#39;none&#39; when it admits nothing | 
**CheckedAt** | **NullableString** | When the tool list was read; null when it is not known | 
**DecidedBy** | **string** | The caller&#39;s own grant, the server&#39;s everyone_tools, or the server itself when it is disabled | 
**Known** | **bool** | Whether the server&#39;s tool list is known: from its last passing check, or listed just now with the caller&#39;s own connection to a per-user server | 
**ServerId** | **string** |  | 
**Tools** | [**[]McpToolSummary**](McpToolSummary.md) | The tools the caller may use, in the server&#39;s order; empty when not known | 

## Methods

### NewWorkspaceMcpServerMyTools

`func NewWorkspaceMcpServerMyTools(access string, checkedAt NullableString, decidedBy string, known bool, serverId string, tools []McpToolSummary, ) *WorkspaceMcpServerMyTools`

NewWorkspaceMcpServerMyTools instantiates a new WorkspaceMcpServerMyTools object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerMyToolsWithDefaults

`func NewWorkspaceMcpServerMyToolsWithDefaults() *WorkspaceMcpServerMyTools`

NewWorkspaceMcpServerMyToolsWithDefaults instantiates a new WorkspaceMcpServerMyTools object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *WorkspaceMcpServerMyTools) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *WorkspaceMcpServerMyTools) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *WorkspaceMcpServerMyTools) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetCheckedAt

`func (o *WorkspaceMcpServerMyTools) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *WorkspaceMcpServerMyTools) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *WorkspaceMcpServerMyTools) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.


### SetCheckedAtNil

`func (o *WorkspaceMcpServerMyTools) SetCheckedAtNil(b bool)`

 SetCheckedAtNil sets the value for CheckedAt to be an explicit nil

### UnsetCheckedAt
`func (o *WorkspaceMcpServerMyTools) UnsetCheckedAt()`

UnsetCheckedAt ensures that no value is present for CheckedAt, not even an explicit nil
### GetDecidedBy

`func (o *WorkspaceMcpServerMyTools) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *WorkspaceMcpServerMyTools) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *WorkspaceMcpServerMyTools) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.


### GetKnown

`func (o *WorkspaceMcpServerMyTools) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *WorkspaceMcpServerMyTools) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *WorkspaceMcpServerMyTools) SetKnown(v bool)`

SetKnown sets Known field to given value.


### GetServerId

`func (o *WorkspaceMcpServerMyTools) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *WorkspaceMcpServerMyTools) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *WorkspaceMcpServerMyTools) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetTools

`func (o *WorkspaceMcpServerMyTools) GetTools() []McpToolSummary`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerMyTools) GetToolsOk() (*[]McpToolSummary, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerMyTools) SetTools(v []McpToolSummary)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


