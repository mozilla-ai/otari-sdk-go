# OrgProviderKeyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ArchivedAt** | Pointer to **NullableTime** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**IsOrgDefault** | **bool** |  | 
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**Provider** | **string** |  | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrgProviderKeyPublic

`func NewOrgProviderKeyPublic(createdAt time.Time, id string, isOrgDefault bool, name string, organizationId string, provider string, ) *OrgProviderKeyPublic`

NewOrgProviderKeyPublic instantiates a new OrgProviderKeyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgProviderKeyPublicWithDefaults

`func NewOrgProviderKeyPublicWithDefaults() *OrgProviderKeyPublic`

NewOrgProviderKeyPublicWithDefaults instantiates a new OrgProviderKeyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *OrgProviderKeyPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *OrgProviderKeyPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *OrgProviderKeyPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *OrgProviderKeyPublic) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *OrgProviderKeyPublic) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *OrgProviderKeyPublic) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetArchivedAt

`func (o *OrgProviderKeyPublic) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *OrgProviderKeyPublic) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *OrgProviderKeyPublic) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *OrgProviderKeyPublic) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### SetArchivedAtNil

`func (o *OrgProviderKeyPublic) SetArchivedAtNil(b bool)`

 SetArchivedAtNil sets the value for ArchivedAt to be an explicit nil

### UnsetArchivedAt
`func (o *OrgProviderKeyPublic) UnsetArchivedAt()`

UnsetArchivedAt ensures that no value is present for ArchivedAt, not even an explicit nil
### GetClientArgs

`func (o *OrgProviderKeyPublic) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *OrgProviderKeyPublic) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *OrgProviderKeyPublic) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *OrgProviderKeyPublic) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *OrgProviderKeyPublic) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *OrgProviderKeyPublic) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetCreatedAt

`func (o *OrgProviderKeyPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgProviderKeyPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgProviderKeyPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *OrgProviderKeyPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgProviderKeyPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgProviderKeyPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIsOrgDefault

`func (o *OrgProviderKeyPublic) GetIsOrgDefault() bool`

GetIsOrgDefault returns the IsOrgDefault field if non-nil, zero value otherwise.

### GetIsOrgDefaultOk

`func (o *OrgProviderKeyPublic) GetIsOrgDefaultOk() (*bool, bool)`

GetIsOrgDefaultOk returns a tuple with the IsOrgDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgDefault

`func (o *OrgProviderKeyPublic) SetIsOrgDefault(v bool)`

SetIsOrgDefault sets IsOrgDefault field to given value.


### GetLast4

`func (o *OrgProviderKeyPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *OrgProviderKeyPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *OrgProviderKeyPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *OrgProviderKeyPublic) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *OrgProviderKeyPublic) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *OrgProviderKeyPublic) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *OrgProviderKeyPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgProviderKeyPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgProviderKeyPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *OrgProviderKeyPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrgProviderKeyPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrgProviderKeyPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProvider

`func (o *OrgProviderKeyPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrgProviderKeyPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrgProviderKeyPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUpdatedAt

`func (o *OrgProviderKeyPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgProviderKeyPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgProviderKeyPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgProviderKeyPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OrgProviderKeyPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OrgProviderKeyPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


