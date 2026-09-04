# WorkspaceMcpServerPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTools** | **[]string** |  | 
**CreatedAt** | **string** |  | 
**Enabled** | **bool** |  | 
**HasToken** | **bool** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**PurposeHint** | **NullableString** |  | 
**UpdatedAt** | **string** |  | 
**Url** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceMcpServerPublic

`func NewWorkspaceMcpServerPublic(allowedTools []string, createdAt string, enabled bool, hasToken bool, id string, name string, purposeHint NullableString, updatedAt string, url string, workspaceId string, ) *WorkspaceMcpServerPublic`

NewWorkspaceMcpServerPublic instantiates a new WorkspaceMcpServerPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerPublicWithDefaults

`func NewWorkspaceMcpServerPublicWithDefaults() *WorkspaceMcpServerPublic`

NewWorkspaceMcpServerPublicWithDefaults instantiates a new WorkspaceMcpServerPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTools

`func (o *WorkspaceMcpServerPublic) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *WorkspaceMcpServerPublic) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *WorkspaceMcpServerPublic) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.


### SetAllowedToolsNil

`func (o *WorkspaceMcpServerPublic) SetAllowedToolsNil(b bool)`

 SetAllowedToolsNil sets the value for AllowedTools to be an explicit nil

### UnsetAllowedTools
`func (o *WorkspaceMcpServerPublic) UnsetAllowedTools()`

UnsetAllowedTools ensures that no value is present for AllowedTools, not even an explicit nil
### GetCreatedAt

`func (o *WorkspaceMcpServerPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMcpServerPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMcpServerPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *WorkspaceMcpServerPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceMcpServerPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceMcpServerPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasToken

`func (o *WorkspaceMcpServerPublic) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *WorkspaceMcpServerPublic) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *WorkspaceMcpServerPublic) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.


### GetId

`func (o *WorkspaceMcpServerPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMcpServerPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMcpServerPublic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkspaceMcpServerPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceMcpServerPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceMcpServerPublic) SetName(v string)`

SetName sets Name field to given value.


### GetPurposeHint

`func (o *WorkspaceMcpServerPublic) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceMcpServerPublic) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceMcpServerPublic) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.


### SetPurposeHintNil

`func (o *WorkspaceMcpServerPublic) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceMcpServerPublic) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetUpdatedAt

`func (o *WorkspaceMcpServerPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceMcpServerPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceMcpServerPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *WorkspaceMcpServerPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkspaceMcpServerPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkspaceMcpServerPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWorkspaceId

`func (o *WorkspaceMcpServerPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceMcpServerPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceMcpServerPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


