# DeploymentAdminAccessPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granted** | **bool** |  | 
**IsBootstrapOperator** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDeploymentAdminAccessPublic

`func NewDeploymentAdminAccessPublic(granted bool, ) *DeploymentAdminAccessPublic`

NewDeploymentAdminAccessPublic instantiates a new DeploymentAdminAccessPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentAdminAccessPublicWithDefaults

`func NewDeploymentAdminAccessPublicWithDefaults() *DeploymentAdminAccessPublic`

NewDeploymentAdminAccessPublicWithDefaults instantiates a new DeploymentAdminAccessPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranted

`func (o *DeploymentAdminAccessPublic) GetGranted() bool`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *DeploymentAdminAccessPublic) GetGrantedOk() (*bool, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *DeploymentAdminAccessPublic) SetGranted(v bool)`

SetGranted sets Granted field to given value.


### GetIsBootstrapOperator

`func (o *DeploymentAdminAccessPublic) GetIsBootstrapOperator() bool`

GetIsBootstrapOperator returns the IsBootstrapOperator field if non-nil, zero value otherwise.

### GetIsBootstrapOperatorOk

`func (o *DeploymentAdminAccessPublic) GetIsBootstrapOperatorOk() (*bool, bool)`

GetIsBootstrapOperatorOk returns a tuple with the IsBootstrapOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootstrapOperator

`func (o *DeploymentAdminAccessPublic) SetIsBootstrapOperator(v bool)`

SetIsBootstrapOperator sets IsBootstrapOperator field to given value.

### HasIsBootstrapOperator

`func (o *DeploymentAdminAccessPublic) HasIsBootstrapOperator() bool`

HasIsBootstrapOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


