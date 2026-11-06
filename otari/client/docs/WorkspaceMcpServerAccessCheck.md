# WorkspaceMcpServerAccessCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** |  | 
**DecidedBy** | **string** | The key&#39;s own grant, its owner&#39;s, the server&#39;s everyone_tools, or the server itself (disabled, or its allowed_tools leaves the tool out) | 
**GrantId** | **NullableString** | The deciding grant, when a grant decided | 
**Tools** | **[]string** | The deciding rule&#39;s tool patterns; the server&#39;s allowed_tools when it decided | 

## Methods

### NewWorkspaceMcpServerAccessCheck

`func NewWorkspaceMcpServerAccessCheck(allowed bool, decidedBy string, grantId NullableString, tools []string, ) *WorkspaceMcpServerAccessCheck`

NewWorkspaceMcpServerAccessCheck instantiates a new WorkspaceMcpServerAccessCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerAccessCheckWithDefaults

`func NewWorkspaceMcpServerAccessCheckWithDefaults() *WorkspaceMcpServerAccessCheck`

NewWorkspaceMcpServerAccessCheckWithDefaults instantiates a new WorkspaceMcpServerAccessCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *WorkspaceMcpServerAccessCheck) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *WorkspaceMcpServerAccessCheck) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *WorkspaceMcpServerAccessCheck) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetDecidedBy

`func (o *WorkspaceMcpServerAccessCheck) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *WorkspaceMcpServerAccessCheck) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *WorkspaceMcpServerAccessCheck) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.


### GetGrantId

`func (o *WorkspaceMcpServerAccessCheck) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *WorkspaceMcpServerAccessCheck) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *WorkspaceMcpServerAccessCheck) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### SetGrantIdNil

`func (o *WorkspaceMcpServerAccessCheck) SetGrantIdNil(b bool)`

 SetGrantIdNil sets the value for GrantId to be an explicit nil

### UnsetGrantId
`func (o *WorkspaceMcpServerAccessCheck) UnsetGrantId()`

UnsetGrantId ensures that no value is present for GrantId, not even an explicit nil
### GetTools

`func (o *WorkspaceMcpServerAccessCheck) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerAccessCheck) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerAccessCheck) SetTools(v []string)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


