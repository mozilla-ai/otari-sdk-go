# WorkspaceHostedProviderAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | False switches the provider off for this workspace. | 
**Models** | Pointer to **[]string** | The models this workspace may call on the provider. Null for every model the deployment serves on it, including ones added later. A list must name at least one model. | [optional] 

## Methods

### NewWorkspaceHostedProviderAccessRequest

`func NewWorkspaceHostedProviderAccessRequest(enabled bool, ) *WorkspaceHostedProviderAccessRequest`

NewWorkspaceHostedProviderAccessRequest instantiates a new WorkspaceHostedProviderAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceHostedProviderAccessRequestWithDefaults

`func NewWorkspaceHostedProviderAccessRequestWithDefaults() *WorkspaceHostedProviderAccessRequest`

NewWorkspaceHostedProviderAccessRequestWithDefaults instantiates a new WorkspaceHostedProviderAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WorkspaceHostedProviderAccessRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceHostedProviderAccessRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceHostedProviderAccessRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetModels

`func (o *WorkspaceHostedProviderAccessRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *WorkspaceHostedProviderAccessRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *WorkspaceHostedProviderAccessRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *WorkspaceHostedProviderAccessRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *WorkspaceHostedProviderAccessRequest) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *WorkspaceHostedProviderAccessRequest) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


