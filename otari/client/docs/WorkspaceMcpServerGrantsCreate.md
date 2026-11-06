# WorkspaceMcpServerGrantsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyIds** | Pointer to **[]string** |  | [optional] 
**Tools** | **[]string** | Tool patterns: names, &#39;*&#39;, &#39;prefix*&#39; or &#39;@read&#39; for the server&#39;s read_tools; [] denies | 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWorkspaceMcpServerGrantsCreate

`func NewWorkspaceMcpServerGrantsCreate(tools []string, ) *WorkspaceMcpServerGrantsCreate`

NewWorkspaceMcpServerGrantsCreate instantiates a new WorkspaceMcpServerGrantsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerGrantsCreateWithDefaults

`func NewWorkspaceMcpServerGrantsCreateWithDefaults() *WorkspaceMcpServerGrantsCreate`

NewWorkspaceMcpServerGrantsCreateWithDefaults instantiates a new WorkspaceMcpServerGrantsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyIds

`func (o *WorkspaceMcpServerGrantsCreate) GetApiKeyIds() []string`

GetApiKeyIds returns the ApiKeyIds field if non-nil, zero value otherwise.

### GetApiKeyIdsOk

`func (o *WorkspaceMcpServerGrantsCreate) GetApiKeyIdsOk() (*[]string, bool)`

GetApiKeyIdsOk returns a tuple with the ApiKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyIds

`func (o *WorkspaceMcpServerGrantsCreate) SetApiKeyIds(v []string)`

SetApiKeyIds sets ApiKeyIds field to given value.

### HasApiKeyIds

`func (o *WorkspaceMcpServerGrantsCreate) HasApiKeyIds() bool`

HasApiKeyIds returns a boolean if a field has been set.

### GetTools

`func (o *WorkspaceMcpServerGrantsCreate) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerGrantsCreate) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerGrantsCreate) SetTools(v []string)`

SetTools sets Tools field to given value.


### GetUserIds

`func (o *WorkspaceMcpServerGrantsCreate) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *WorkspaceMcpServerGrantsCreate) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *WorkspaceMcpServerGrantsCreate) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *WorkspaceMcpServerGrantsCreate) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


