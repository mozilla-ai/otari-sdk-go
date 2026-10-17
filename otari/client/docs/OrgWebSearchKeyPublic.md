# OrgWebSearchKeyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**IsOrgDefault** | **bool** |  | 
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**Provider** | **string** |  | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Usable** | **bool** | False when this deployment cannot decrypt the stored key, so no search uses it. It is still listed, because replacing or deleting it is what fixes it. | 

## Methods

### NewOrgWebSearchKeyPublic

`func NewOrgWebSearchKeyPublic(createdAt time.Time, id string, isOrgDefault bool, name string, organizationId string, provider string, usable bool, ) *OrgWebSearchKeyPublic`

NewOrgWebSearchKeyPublic instantiates a new OrgWebSearchKeyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWebSearchKeyPublicWithDefaults

`func NewOrgWebSearchKeyPublicWithDefaults() *OrgWebSearchKeyPublic`

NewOrgWebSearchKeyPublicWithDefaults instantiates a new OrgWebSearchKeyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedAt

`func (o *OrgWebSearchKeyPublic) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *OrgWebSearchKeyPublic) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *OrgWebSearchKeyPublic) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *OrgWebSearchKeyPublic) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *OrgWebSearchKeyPublic) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *OrgWebSearchKeyPublic) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetCreatedAt

`func (o *OrgWebSearchKeyPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgWebSearchKeyPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgWebSearchKeyPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *OrgWebSearchKeyPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgWebSearchKeyPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgWebSearchKeyPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIsOrgDefault

`func (o *OrgWebSearchKeyPublic) GetIsOrgDefault() bool`

GetIsOrgDefault returns the IsOrgDefault field if non-nil, zero value otherwise.

### GetIsOrgDefaultOk

`func (o *OrgWebSearchKeyPublic) GetIsOrgDefaultOk() (*bool, bool)`

GetIsOrgDefaultOk returns a tuple with the IsOrgDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgDefault

`func (o *OrgWebSearchKeyPublic) SetIsOrgDefault(v bool)`

SetIsOrgDefault sets IsOrgDefault field to given value.


### GetLast4

`func (o *OrgWebSearchKeyPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *OrgWebSearchKeyPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *OrgWebSearchKeyPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *OrgWebSearchKeyPublic) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *OrgWebSearchKeyPublic) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *OrgWebSearchKeyPublic) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *OrgWebSearchKeyPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgWebSearchKeyPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgWebSearchKeyPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *OrgWebSearchKeyPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrgWebSearchKeyPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrgWebSearchKeyPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProvider

`func (o *OrgWebSearchKeyPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrgWebSearchKeyPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrgWebSearchKeyPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUpdatedAt

`func (o *OrgWebSearchKeyPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgWebSearchKeyPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgWebSearchKeyPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgWebSearchKeyPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OrgWebSearchKeyPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OrgWebSearchKeyPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUsable

`func (o *OrgWebSearchKeyPublic) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *OrgWebSearchKeyPublic) GetUsableOk() (*bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *OrgWebSearchKeyPublic) SetUsable(v bool)`

SetUsable sets Usable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


