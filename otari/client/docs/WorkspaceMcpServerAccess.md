# WorkspaceMcpServerAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EveryoneTools** | **[]string** |  | 
**Grants** | [**[]WorkspaceMcpServerGrantPublic**](WorkspaceMcpServerGrantPublic.md) |  | 

## Methods

### NewWorkspaceMcpServerAccess

`func NewWorkspaceMcpServerAccess(everyoneTools []string, grants []WorkspaceMcpServerGrantPublic, ) *WorkspaceMcpServerAccess`

NewWorkspaceMcpServerAccess instantiates a new WorkspaceMcpServerAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerAccessWithDefaults

`func NewWorkspaceMcpServerAccessWithDefaults() *WorkspaceMcpServerAccess`

NewWorkspaceMcpServerAccessWithDefaults instantiates a new WorkspaceMcpServerAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEveryoneTools

`func (o *WorkspaceMcpServerAccess) GetEveryoneTools() []string`

GetEveryoneTools returns the EveryoneTools field if non-nil, zero value otherwise.

### GetEveryoneToolsOk

`func (o *WorkspaceMcpServerAccess) GetEveryoneToolsOk() (*[]string, bool)`

GetEveryoneToolsOk returns a tuple with the EveryoneTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveryoneTools

`func (o *WorkspaceMcpServerAccess) SetEveryoneTools(v []string)`

SetEveryoneTools sets EveryoneTools field to given value.


### GetGrants

`func (o *WorkspaceMcpServerAccess) GetGrants() []WorkspaceMcpServerGrantPublic`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *WorkspaceMcpServerAccess) GetGrantsOk() (*[]WorkspaceMcpServerGrantPublic, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *WorkspaceMcpServerAccess) SetGrants(v []WorkspaceMcpServerGrantPublic)`

SetGrants sets Grants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


