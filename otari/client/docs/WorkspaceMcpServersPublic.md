# WorkspaceMcpServersPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Data** | [**[]WorkspaceMcpServerPublic**](WorkspaceMcpServerPublic.md) |  | 
**SignInUnavailable** | Pointer to **NullableString** | Why this deployment cannot sign anyone in to an MCP server yet, in words for a workspace admin: a server in personal-accounts mode cannot be added, and nobody can sign in for everyone. Null when sign-ins work. | [optional] 

## Methods

### NewWorkspaceMcpServersPublic

`func NewWorkspaceMcpServersPublic(count int32, data []WorkspaceMcpServerPublic, ) *WorkspaceMcpServersPublic`

NewWorkspaceMcpServersPublic instantiates a new WorkspaceMcpServersPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServersPublicWithDefaults

`func NewWorkspaceMcpServersPublicWithDefaults() *WorkspaceMcpServersPublic`

NewWorkspaceMcpServersPublicWithDefaults instantiates a new WorkspaceMcpServersPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WorkspaceMcpServersPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkspaceMcpServersPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkspaceMcpServersPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *WorkspaceMcpServersPublic) GetData() []WorkspaceMcpServerPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WorkspaceMcpServersPublic) GetDataOk() (*[]WorkspaceMcpServerPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WorkspaceMcpServersPublic) SetData(v []WorkspaceMcpServerPublic)`

SetData sets Data field to given value.


### GetSignInUnavailable

`func (o *WorkspaceMcpServersPublic) GetSignInUnavailable() string`

GetSignInUnavailable returns the SignInUnavailable field if non-nil, zero value otherwise.

### GetSignInUnavailableOk

`func (o *WorkspaceMcpServersPublic) GetSignInUnavailableOk() (*string, bool)`

GetSignInUnavailableOk returns a tuple with the SignInUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInUnavailable

`func (o *WorkspaceMcpServersPublic) SetSignInUnavailable(v string)`

SetSignInUnavailable sets SignInUnavailable field to given value.

### HasSignInUnavailable

`func (o *WorkspaceMcpServersPublic) HasSignInUnavailable() bool`

HasSignInUnavailable returns a boolean if a field has been set.

### SetSignInUnavailableNil

`func (o *WorkspaceMcpServersPublic) SetSignInUnavailableNil(b bool)`

 SetSignInUnavailableNil sets the value for SignInUnavailable to be an explicit nil

### UnsetSignInUnavailable
`func (o *WorkspaceMcpServersPublic) UnsetSignInUnavailable()`

UnsetSignInUnavailable ensures that no value is present for SignInUnavailable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


