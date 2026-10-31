# WorkspaceWebSearchKeysPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Every live key of the organization, not only this page. | 
**Data** | [**[]WorkspaceWebSearchKeyPublic**](WorkspaceWebSearchKeyPublic.md) |  | 
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


