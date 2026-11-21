# WorkspaceWebSearchKeysPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Every live key of the organization, not only this page. | 
**Data** | [**[]WorkspaceWebSearchKeyPublic**](WorkspaceWebSearchKeyPublic.md) |  | 
**DeploymentFetchTool** | Pointer to **NullableString** | The name of the deployment&#39;s fetch tool, which the workspace&#39;s fetches run on when no key serves them, whether or not one does now: builtin_fetch for the built-in fetcher. | [optional] 
**DeploymentSearchTool** | Pointer to **NullableString** | The name of the deployment&#39;s search tool, which the workspace&#39;s searches run on when no key serves them, whether or not one does now. Null when the deployment has none. | [optional] 
**FetchProvider** | Pointer to **NullableString** | The provider the workspace&#39;s fetch tool runs on, and the enrichment of its search hits unless the search&#39;s instance names a fetch instance of its own: its effective fetch key&#39;s, else the deployment&#39;s default fetch instance&#39;s, which is builtin for the built-in fetcher. | [optional] 
**SearchProvider** | Pointer to **NullableString** | The provider the workspace&#39;s searches run on: its effective search key&#39;s, else the deployment&#39;s default search&#39;s. Null when nothing searches for it. | [optional] 

## Methods

### NewWorkspaceWebSearchKeysPublic

`func NewWorkspaceWebSearchKeysPublic(count int32, data []WorkspaceWebSearchKeyPublic, ) *WorkspaceWebSearchKeysPublic`

NewWorkspaceWebSearchKeysPublic instantiates a new WorkspaceWebSearchKeysPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWebSearchKeysPublicWithDefaults

`func NewWorkspaceWebSearchKeysPublicWithDefaults() *WorkspaceWebSearchKeysPublic`

NewWorkspaceWebSearchKeysPublicWithDefaults instantiates a new WorkspaceWebSearchKeysPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WorkspaceWebSearchKeysPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WorkspaceWebSearchKeysPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WorkspaceWebSearchKeysPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *WorkspaceWebSearchKeysPublic) GetData() []WorkspaceWebSearchKeyPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WorkspaceWebSearchKeysPublic) GetDataOk() (*[]WorkspaceWebSearchKeyPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WorkspaceWebSearchKeysPublic) SetData(v []WorkspaceWebSearchKeyPublic)`

SetData sets Data field to given value.


### GetDeploymentFetchTool

`func (o *WorkspaceWebSearchKeysPublic) GetDeploymentFetchTool() string`

GetDeploymentFetchTool returns the DeploymentFetchTool field if non-nil, zero value otherwise.

### GetDeploymentFetchToolOk

`func (o *WorkspaceWebSearchKeysPublic) GetDeploymentFetchToolOk() (*string, bool)`

GetDeploymentFetchToolOk returns a tuple with the DeploymentFetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentFetchTool

`func (o *WorkspaceWebSearchKeysPublic) SetDeploymentFetchTool(v string)`

SetDeploymentFetchTool sets DeploymentFetchTool field to given value.

### HasDeploymentFetchTool

`func (o *WorkspaceWebSearchKeysPublic) HasDeploymentFetchTool() bool`

HasDeploymentFetchTool returns a boolean if a field has been set.

### SetDeploymentFetchToolNil

`func (o *WorkspaceWebSearchKeysPublic) SetDeploymentFetchToolNil(b bool)`

 SetDeploymentFetchToolNil sets the value for DeploymentFetchTool to be an explicit nil

### UnsetDeploymentFetchTool
`func (o *WorkspaceWebSearchKeysPublic) UnsetDeploymentFetchTool()`

UnsetDeploymentFetchTool ensures that no value is present for DeploymentFetchTool, not even an explicit nil
### GetDeploymentSearchTool

`func (o *WorkspaceWebSearchKeysPublic) GetDeploymentSearchTool() string`

GetDeploymentSearchTool returns the DeploymentSearchTool field if non-nil, zero value otherwise.

### GetDeploymentSearchToolOk

`func (o *WorkspaceWebSearchKeysPublic) GetDeploymentSearchToolOk() (*string, bool)`

GetDeploymentSearchToolOk returns a tuple with the DeploymentSearchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSearchTool

`func (o *WorkspaceWebSearchKeysPublic) SetDeploymentSearchTool(v string)`

SetDeploymentSearchTool sets DeploymentSearchTool field to given value.

### HasDeploymentSearchTool

`func (o *WorkspaceWebSearchKeysPublic) HasDeploymentSearchTool() bool`

HasDeploymentSearchTool returns a boolean if a field has been set.

### SetDeploymentSearchToolNil

`func (o *WorkspaceWebSearchKeysPublic) SetDeploymentSearchToolNil(b bool)`

 SetDeploymentSearchToolNil sets the value for DeploymentSearchTool to be an explicit nil

### UnsetDeploymentSearchTool
`func (o *WorkspaceWebSearchKeysPublic) UnsetDeploymentSearchTool()`

UnsetDeploymentSearchTool ensures that no value is present for DeploymentSearchTool, not even an explicit nil
### GetFetchProvider

`func (o *WorkspaceWebSearchKeysPublic) GetFetchProvider() string`

GetFetchProvider returns the FetchProvider field if non-nil, zero value otherwise.

### GetFetchProviderOk

`func (o *WorkspaceWebSearchKeysPublic) GetFetchProviderOk() (*string, bool)`

GetFetchProviderOk returns a tuple with the FetchProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchProvider

`func (o *WorkspaceWebSearchKeysPublic) SetFetchProvider(v string)`

SetFetchProvider sets FetchProvider field to given value.

### HasFetchProvider

`func (o *WorkspaceWebSearchKeysPublic) HasFetchProvider() bool`

HasFetchProvider returns a boolean if a field has been set.

### SetFetchProviderNil

`func (o *WorkspaceWebSearchKeysPublic) SetFetchProviderNil(b bool)`

 SetFetchProviderNil sets the value for FetchProvider to be an explicit nil

### UnsetFetchProvider
`func (o *WorkspaceWebSearchKeysPublic) UnsetFetchProvider()`

UnsetFetchProvider ensures that no value is present for FetchProvider, not even an explicit nil
### GetSearchProvider

`func (o *WorkspaceWebSearchKeysPublic) GetSearchProvider() string`

GetSearchProvider returns the SearchProvider field if non-nil, zero value otherwise.

### GetSearchProviderOk

`func (o *WorkspaceWebSearchKeysPublic) GetSearchProviderOk() (*string, bool)`

GetSearchProviderOk returns a tuple with the SearchProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProvider

`func (o *WorkspaceWebSearchKeysPublic) SetSearchProvider(v string)`

SetSearchProvider sets SearchProvider field to given value.

### HasSearchProvider

`func (o *WorkspaceWebSearchKeysPublic) HasSearchProvider() bool`

HasSearchProvider returns a boolean if a field has been set.

### SetSearchProviderNil

`func (o *WorkspaceWebSearchKeysPublic) SetSearchProviderNil(b bool)`

 SetSearchProviderNil sets the value for SearchProvider to be an explicit nil

### UnsetSearchProvider
`func (o *WorkspaceWebSearchKeysPublic) UnsetSearchProvider()`

UnsetSearchProvider ensures that no value is present for SearchProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


