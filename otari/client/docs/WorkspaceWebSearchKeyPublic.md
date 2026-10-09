# WorkspaceWebSearchKeyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** | The workspace turned this key off. | 
**IsDefault** | **bool** | The workspace pinned this key as its own. | 
**IsEffective** | **bool** | This is the key the workspace&#39;s searches use. | 
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OrgWebSearchKeyId** | **string** |  | 
**Provider** | **string** |  | 
**Usable** | **bool** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceWebSearchKeyPublic

`func NewWorkspaceWebSearchKeyPublic(disabled bool, isDefault bool, isEffective bool, name string, orgWebSearchKeyId string, provider string, usable bool, workspaceId string, ) *WorkspaceWebSearchKeyPublic`

NewWorkspaceWebSearchKeyPublic instantiates a new WorkspaceWebSearchKeyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWebSearchKeyPublicWithDefaults

`func NewWorkspaceWebSearchKeyPublicWithDefaults() *WorkspaceWebSearchKeyPublic`

NewWorkspaceWebSearchKeyPublicWithDefaults instantiates a new WorkspaceWebSearchKeyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *WorkspaceWebSearchKeyPublic) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WorkspaceWebSearchKeyPublic) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WorkspaceWebSearchKeyPublic) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetIsDefault

`func (o *WorkspaceWebSearchKeyPublic) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *WorkspaceWebSearchKeyPublic) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *WorkspaceWebSearchKeyPublic) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsEffective

`func (o *WorkspaceWebSearchKeyPublic) GetIsEffective() bool`

GetIsEffective returns the IsEffective field if non-nil, zero value otherwise.

### GetIsEffectiveOk

`func (o *WorkspaceWebSearchKeyPublic) GetIsEffectiveOk() (*bool, bool)`

GetIsEffectiveOk returns a tuple with the IsEffective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEffective

`func (o *WorkspaceWebSearchKeyPublic) SetIsEffective(v bool)`

SetIsEffective sets IsEffective field to given value.


### GetLast4

`func (o *WorkspaceWebSearchKeyPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *WorkspaceWebSearchKeyPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *WorkspaceWebSearchKeyPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *WorkspaceWebSearchKeyPublic) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *WorkspaceWebSearchKeyPublic) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *WorkspaceWebSearchKeyPublic) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *WorkspaceWebSearchKeyPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceWebSearchKeyPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceWebSearchKeyPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOrgWebSearchKeyId

`func (o *WorkspaceWebSearchKeyPublic) GetOrgWebSearchKeyId() string`

GetOrgWebSearchKeyId returns the OrgWebSearchKeyId field if non-nil, zero value otherwise.

### GetOrgWebSearchKeyIdOk

`func (o *WorkspaceWebSearchKeyPublic) GetOrgWebSearchKeyIdOk() (*string, bool)`

GetOrgWebSearchKeyIdOk returns a tuple with the OrgWebSearchKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgWebSearchKeyId

`func (o *WorkspaceWebSearchKeyPublic) SetOrgWebSearchKeyId(v string)`

SetOrgWebSearchKeyId sets OrgWebSearchKeyId field to given value.


### GetProvider

`func (o *WorkspaceWebSearchKeyPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WorkspaceWebSearchKeyPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WorkspaceWebSearchKeyPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUsable

`func (o *WorkspaceWebSearchKeyPublic) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *WorkspaceWebSearchKeyPublic) GetUsableOk() (*bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *WorkspaceWebSearchKeyPublic) SetUsable(v bool)`

SetUsable sets Usable field to given value.


### GetWorkspaceId

`func (o *WorkspaceWebSearchKeyPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceWebSearchKeyPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceWebSearchKeyPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


