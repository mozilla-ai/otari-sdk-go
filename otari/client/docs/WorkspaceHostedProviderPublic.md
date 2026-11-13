# WorkspaceHostedProviderPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customized** | **bool** | Whether this workspace stores a departure from full inheritance. | 
**Enabled** | **bool** |  | 
**Models** | **[]string** | This workspace&#39;s model list, or null for every served model. | 
**Provider** | **string** | The hosted provider&#39;s name, as a request&#39;s selector prefix spells it. | 
**ServedByOwnKey** | **bool** | Whether the workspace calls this provider on the organization&#39;s own key, so this row is not used. | 
**ServedModels** | **[]string** | The models the deployment serves on the provider for this organization, or null for any priced one. | 
**ZeroDataRetention** | Pointer to **bool** | Whether the deployment declares this provider keeps no prompt or completion data. | [optional] [default to false]

## Methods

### NewWorkspaceHostedProviderPublic

`func NewWorkspaceHostedProviderPublic(customized bool, enabled bool, models []string, provider string, servedByOwnKey bool, servedModels []string, ) *WorkspaceHostedProviderPublic`

NewWorkspaceHostedProviderPublic instantiates a new WorkspaceHostedProviderPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceHostedProviderPublicWithDefaults

`func NewWorkspaceHostedProviderPublicWithDefaults() *WorkspaceHostedProviderPublic`

NewWorkspaceHostedProviderPublicWithDefaults instantiates a new WorkspaceHostedProviderPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomized

`func (o *WorkspaceHostedProviderPublic) GetCustomized() bool`

GetCustomized returns the Customized field if non-nil, zero value otherwise.

### GetCustomizedOk

`func (o *WorkspaceHostedProviderPublic) GetCustomizedOk() (*bool, bool)`

GetCustomizedOk returns a tuple with the Customized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomized

`func (o *WorkspaceHostedProviderPublic) SetCustomized(v bool)`

SetCustomized sets Customized field to given value.


### GetEnabled

`func (o *WorkspaceHostedProviderPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceHostedProviderPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceHostedProviderPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetModels

`func (o *WorkspaceHostedProviderPublic) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *WorkspaceHostedProviderPublic) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *WorkspaceHostedProviderPublic) SetModels(v []string)`

SetModels sets Models field to given value.


### SetModelsNil

`func (o *WorkspaceHostedProviderPublic) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *WorkspaceHostedProviderPublic) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetProvider

`func (o *WorkspaceHostedProviderPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WorkspaceHostedProviderPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WorkspaceHostedProviderPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetServedByOwnKey

`func (o *WorkspaceHostedProviderPublic) GetServedByOwnKey() bool`

GetServedByOwnKey returns the ServedByOwnKey field if non-nil, zero value otherwise.

### GetServedByOwnKeyOk

`func (o *WorkspaceHostedProviderPublic) GetServedByOwnKeyOk() (*bool, bool)`

GetServedByOwnKeyOk returns a tuple with the ServedByOwnKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedByOwnKey

`func (o *WorkspaceHostedProviderPublic) SetServedByOwnKey(v bool)`

SetServedByOwnKey sets ServedByOwnKey field to given value.


### GetServedModels

`func (o *WorkspaceHostedProviderPublic) GetServedModels() []string`

GetServedModels returns the ServedModels field if non-nil, zero value otherwise.

### GetServedModelsOk

`func (o *WorkspaceHostedProviderPublic) GetServedModelsOk() (*[]string, bool)`

GetServedModelsOk returns a tuple with the ServedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedModels

`func (o *WorkspaceHostedProviderPublic) SetServedModels(v []string)`

SetServedModels sets ServedModels field to given value.


### SetServedModelsNil

`func (o *WorkspaceHostedProviderPublic) SetServedModelsNil(b bool)`

 SetServedModelsNil sets the value for ServedModels to be an explicit nil

### UnsetServedModels
`func (o *WorkspaceHostedProviderPublic) UnsetServedModels()`

UnsetServedModels ensures that no value is present for ServedModels, not even an explicit nil
### GetZeroDataRetention

`func (o *WorkspaceHostedProviderPublic) GetZeroDataRetention() bool`

GetZeroDataRetention returns the ZeroDataRetention field if non-nil, zero value otherwise.

### GetZeroDataRetentionOk

`func (o *WorkspaceHostedProviderPublic) GetZeroDataRetentionOk() (*bool, bool)`

GetZeroDataRetentionOk returns a tuple with the ZeroDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroDataRetention

`func (o *WorkspaceHostedProviderPublic) SetZeroDataRetention(v bool)`

SetZeroDataRetention sets ZeroDataRetention field to given value.

### HasZeroDataRetention

`func (o *WorkspaceHostedProviderPublic) HasZeroDataRetention() bool`

HasZeroDataRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


