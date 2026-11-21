# DeploymentProviderMigrationRowPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Instance** | **string** |  | 
**Models** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **[]string** |  | [optional] 
**Outcome** | **string** |  | 
**Provider** | Pointer to **NullableString** |  | [optional] 
**SourceRemoved** | Pointer to **bool** |  | [optional] [default to false]
**UpdatedAt** | Pointer to **NullableTime** | The deployment provider&#39;s version, which an apply names to move this row as shown. | [optional] 

## Methods

### NewDeploymentProviderMigrationRowPublic

`func NewDeploymentProviderMigrationRowPublic(instance string, outcome string, ) *DeploymentProviderMigrationRowPublic`

NewDeploymentProviderMigrationRowPublic instantiates a new DeploymentProviderMigrationRowPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProviderMigrationRowPublicWithDefaults

`func NewDeploymentProviderMigrationRowPublicWithDefaults() *DeploymentProviderMigrationRowPublic`

NewDeploymentProviderMigrationRowPublicWithDefaults instantiates a new DeploymentProviderMigrationRowPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *DeploymentProviderMigrationRowPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *DeploymentProviderMigrationRowPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *DeploymentProviderMigrationRowPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *DeploymentProviderMigrationRowPublic) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *DeploymentProviderMigrationRowPublic) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *DeploymentProviderMigrationRowPublic) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetDetail

`func (o *DeploymentProviderMigrationRowPublic) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DeploymentProviderMigrationRowPublic) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DeploymentProviderMigrationRowPublic) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DeploymentProviderMigrationRowPublic) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *DeploymentProviderMigrationRowPublic) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *DeploymentProviderMigrationRowPublic) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetInstance

`func (o *DeploymentProviderMigrationRowPublic) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *DeploymentProviderMigrationRowPublic) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *DeploymentProviderMigrationRowPublic) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetModels

`func (o *DeploymentProviderMigrationRowPublic) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *DeploymentProviderMigrationRowPublic) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *DeploymentProviderMigrationRowPublic) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *DeploymentProviderMigrationRowPublic) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *DeploymentProviderMigrationRowPublic) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *DeploymentProviderMigrationRowPublic) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetNotes

`func (o *DeploymentProviderMigrationRowPublic) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DeploymentProviderMigrationRowPublic) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DeploymentProviderMigrationRowPublic) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DeploymentProviderMigrationRowPublic) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOutcome

`func (o *DeploymentProviderMigrationRowPublic) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *DeploymentProviderMigrationRowPublic) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *DeploymentProviderMigrationRowPublic) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetProvider

`func (o *DeploymentProviderMigrationRowPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DeploymentProviderMigrationRowPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DeploymentProviderMigrationRowPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DeploymentProviderMigrationRowPublic) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *DeploymentProviderMigrationRowPublic) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *DeploymentProviderMigrationRowPublic) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetSourceRemoved

`func (o *DeploymentProviderMigrationRowPublic) GetSourceRemoved() bool`

GetSourceRemoved returns the SourceRemoved field if non-nil, zero value otherwise.

### GetSourceRemovedOk

`func (o *DeploymentProviderMigrationRowPublic) GetSourceRemovedOk() (*bool, bool)`

GetSourceRemovedOk returns a tuple with the SourceRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRemoved

`func (o *DeploymentProviderMigrationRowPublic) SetSourceRemoved(v bool)`

SetSourceRemoved sets SourceRemoved field to given value.

### HasSourceRemoved

`func (o *DeploymentProviderMigrationRowPublic) HasSourceRemoved() bool`

HasSourceRemoved returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeploymentProviderMigrationRowPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentProviderMigrationRowPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentProviderMigrationRowPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentProviderMigrationRowPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DeploymentProviderMigrationRowPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DeploymentProviderMigrationRowPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


