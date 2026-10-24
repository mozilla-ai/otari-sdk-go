# DeploymentProviderMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]DeploymentProviderMigrationReviewedRow**](DeploymentProviderMigrationReviewedRow.md) |  | [optional] 

## Methods

### NewDeploymentProviderMigrationRequest

`func NewDeploymentProviderMigrationRequest() *DeploymentProviderMigrationRequest`

NewDeploymentProviderMigrationRequest instantiates a new DeploymentProviderMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentProviderMigrationRequestWithDefaults

`func NewDeploymentProviderMigrationRequestWithDefaults() *DeploymentProviderMigrationRequest`

NewDeploymentProviderMigrationRequestWithDefaults instantiates a new DeploymentProviderMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *DeploymentProviderMigrationRequest) GetRows() []DeploymentProviderMigrationReviewedRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DeploymentProviderMigrationRequest) GetRowsOk() (*[]DeploymentProviderMigrationReviewedRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DeploymentProviderMigrationRequest) SetRows(v []DeploymentProviderMigrationReviewedRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *DeploymentProviderMigrationRequest) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRowsNil

`func (o *DeploymentProviderMigrationRequest) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *DeploymentProviderMigrationRequest) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


