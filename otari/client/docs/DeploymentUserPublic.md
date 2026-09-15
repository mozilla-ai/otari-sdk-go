# DeploymentUserPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Email** | **NullableString** |  | 
**FullName** | **NullableString** |  | 
**Id** | **string** |  | 
**IsActive** | **bool** |  | 
**IsBootstrapOperator** | **bool** |  | 
**IsSelf** | **bool** |  | 
**IsSuperuser** | **bool** |  | 
**LastSignInAt** | **NullableTime** |  | 
**Organizations** | [**[]DeploymentUserOrganizationPublic**](DeploymentUserOrganizationPublic.md) |  | 

## Methods

### NewDeploymentUserPublic

`func NewDeploymentUserPublic(createdAt time.Time, email NullableString, fullName NullableString, id string, isActive bool, isBootstrapOperator bool, isSelf bool, isSuperuser bool, lastSignInAt NullableTime, organizations []DeploymentUserOrganizationPublic, ) *DeploymentUserPublic`

NewDeploymentUserPublic instantiates a new DeploymentUserPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentUserPublicWithDefaults

`func NewDeploymentUserPublicWithDefaults() *DeploymentUserPublic`

NewDeploymentUserPublicWithDefaults instantiates a new DeploymentUserPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DeploymentUserPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentUserPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentUserPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *DeploymentUserPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeploymentUserPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeploymentUserPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *DeploymentUserPublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *DeploymentUserPublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *DeploymentUserPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DeploymentUserPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DeploymentUserPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.


### SetFullNameNil

`func (o *DeploymentUserPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *DeploymentUserPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetId

`func (o *DeploymentUserPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentUserPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentUserPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *DeploymentUserPublic) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DeploymentUserPublic) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DeploymentUserPublic) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsBootstrapOperator

`func (o *DeploymentUserPublic) GetIsBootstrapOperator() bool`

GetIsBootstrapOperator returns the IsBootstrapOperator field if non-nil, zero value otherwise.

### GetIsBootstrapOperatorOk

`func (o *DeploymentUserPublic) GetIsBootstrapOperatorOk() (*bool, bool)`

GetIsBootstrapOperatorOk returns a tuple with the IsBootstrapOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootstrapOperator

`func (o *DeploymentUserPublic) SetIsBootstrapOperator(v bool)`

SetIsBootstrapOperator sets IsBootstrapOperator field to given value.


### GetIsSelf

`func (o *DeploymentUserPublic) GetIsSelf() bool`

GetIsSelf returns the IsSelf field if non-nil, zero value otherwise.

### GetIsSelfOk

`func (o *DeploymentUserPublic) GetIsSelfOk() (*bool, bool)`

GetIsSelfOk returns a tuple with the IsSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelf

`func (o *DeploymentUserPublic) SetIsSelf(v bool)`

SetIsSelf sets IsSelf field to given value.


### GetIsSuperuser

`func (o *DeploymentUserPublic) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *DeploymentUserPublic) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *DeploymentUserPublic) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.


### GetLastSignInAt

`func (o *DeploymentUserPublic) GetLastSignInAt() time.Time`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *DeploymentUserPublic) GetLastSignInAtOk() (*time.Time, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *DeploymentUserPublic) SetLastSignInAt(v time.Time)`

SetLastSignInAt sets LastSignInAt field to given value.


### SetLastSignInAtNil

`func (o *DeploymentUserPublic) SetLastSignInAtNil(b bool)`

 SetLastSignInAtNil sets the value for LastSignInAt to be an explicit nil

### UnsetLastSignInAt
`func (o *DeploymentUserPublic) UnsetLastSignInAt()`

UnsetLastSignInAt ensures that no value is present for LastSignInAt, not even an explicit nil
### GetOrganizations

`func (o *DeploymentUserPublic) GetOrganizations() []DeploymentUserOrganizationPublic`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *DeploymentUserPublic) GetOrganizationsOk() (*[]DeploymentUserOrganizationPublic, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *DeploymentUserPublic) SetOrganizations(v []DeploymentUserOrganizationPublic)`

SetOrganizations sets Organizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


