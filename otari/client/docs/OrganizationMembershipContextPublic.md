# OrganizationMembershipContextPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByoProviderKeysAllowed** | Pointer to **bool** |  | [optional] [default to false]
**Caller** | Pointer to [**NullableCallerIdentityPublic**](CallerIdentityPublic.md) |  | [optional] 
**DeploymentOperator** | Pointer to **bool** |  | [optional] [default to false]
**Organization** | [**OrganizationPublic**](OrganizationPublic.md) |  | 
**OrganizationMemberId** | **string** |  | 
**ProviderKeyEncryptionAvailable** | Pointer to **bool** |  | [optional] [default to false]
**Role** | **string** |  | 
**Status** | **string** |  | 
**WorkspaceMemberships** | Pointer to [**[]CallerWorkspaceMembershipPublic**](CallerWorkspaceMembershipPublic.md) |  | [optional] 

## Methods

### NewOrganizationMembershipContextPublic

`func NewOrganizationMembershipContextPublic(organization OrganizationPublic, organizationMemberId string, role string, status string, ) *OrganizationMembershipContextPublic`

NewOrganizationMembershipContextPublic instantiates a new OrganizationMembershipContextPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipContextPublicWithDefaults

`func NewOrganizationMembershipContextPublicWithDefaults() *OrganizationMembershipContextPublic`

NewOrganizationMembershipContextPublicWithDefaults instantiates a new OrganizationMembershipContextPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByoProviderKeysAllowed

`func (o *OrganizationMembershipContextPublic) GetByoProviderKeysAllowed() bool`

GetByoProviderKeysAllowed returns the ByoProviderKeysAllowed field if non-nil, zero value otherwise.

### GetByoProviderKeysAllowedOk

`func (o *OrganizationMembershipContextPublic) GetByoProviderKeysAllowedOk() (*bool, bool)`

GetByoProviderKeysAllowedOk returns a tuple with the ByoProviderKeysAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoProviderKeysAllowed

`func (o *OrganizationMembershipContextPublic) SetByoProviderKeysAllowed(v bool)`

SetByoProviderKeysAllowed sets ByoProviderKeysAllowed field to given value.

### HasByoProviderKeysAllowed

`func (o *OrganizationMembershipContextPublic) HasByoProviderKeysAllowed() bool`

HasByoProviderKeysAllowed returns a boolean if a field has been set.

### GetCaller

`func (o *OrganizationMembershipContextPublic) GetCaller() CallerIdentityPublic`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *OrganizationMembershipContextPublic) GetCallerOk() (*CallerIdentityPublic, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *OrganizationMembershipContextPublic) SetCaller(v CallerIdentityPublic)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *OrganizationMembershipContextPublic) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *OrganizationMembershipContextPublic) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *OrganizationMembershipContextPublic) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetDeploymentOperator

`func (o *OrganizationMembershipContextPublic) GetDeploymentOperator() bool`

GetDeploymentOperator returns the DeploymentOperator field if non-nil, zero value otherwise.

### GetDeploymentOperatorOk

`func (o *OrganizationMembershipContextPublic) GetDeploymentOperatorOk() (*bool, bool)`

GetDeploymentOperatorOk returns a tuple with the DeploymentOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOperator

`func (o *OrganizationMembershipContextPublic) SetDeploymentOperator(v bool)`

SetDeploymentOperator sets DeploymentOperator field to given value.

### HasDeploymentOperator

`func (o *OrganizationMembershipContextPublic) HasDeploymentOperator() bool`

HasDeploymentOperator returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationMembershipContextPublic) GetOrganization() OrganizationPublic`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationMembershipContextPublic) GetOrganizationOk() (*OrganizationPublic, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationMembershipContextPublic) SetOrganization(v OrganizationPublic)`

SetOrganization sets Organization field to given value.


### GetOrganizationMemberId

`func (o *OrganizationMembershipContextPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *OrganizationMembershipContextPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *OrganizationMembershipContextPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.


### GetProviderKeyEncryptionAvailable

`func (o *OrganizationMembershipContextPublic) GetProviderKeyEncryptionAvailable() bool`

GetProviderKeyEncryptionAvailable returns the ProviderKeyEncryptionAvailable field if non-nil, zero value otherwise.

### GetProviderKeyEncryptionAvailableOk

`func (o *OrganizationMembershipContextPublic) GetProviderKeyEncryptionAvailableOk() (*bool, bool)`

GetProviderKeyEncryptionAvailableOk returns a tuple with the ProviderKeyEncryptionAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKeyEncryptionAvailable

`func (o *OrganizationMembershipContextPublic) SetProviderKeyEncryptionAvailable(v bool)`

SetProviderKeyEncryptionAvailable sets ProviderKeyEncryptionAvailable field to given value.

### HasProviderKeyEncryptionAvailable

`func (o *OrganizationMembershipContextPublic) HasProviderKeyEncryptionAvailable() bool`

HasProviderKeyEncryptionAvailable returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMembershipContextPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMembershipContextPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMembershipContextPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *OrganizationMembershipContextPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationMembershipContextPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationMembershipContextPublic) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWorkspaceMemberships

`func (o *OrganizationMembershipContextPublic) GetWorkspaceMemberships() []CallerWorkspaceMembershipPublic`

GetWorkspaceMemberships returns the WorkspaceMemberships field if non-nil, zero value otherwise.

### GetWorkspaceMembershipsOk

`func (o *OrganizationMembershipContextPublic) GetWorkspaceMembershipsOk() (*[]CallerWorkspaceMembershipPublic, bool)`

GetWorkspaceMembershipsOk returns a tuple with the WorkspaceMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberships

`func (o *OrganizationMembershipContextPublic) SetWorkspaceMemberships(v []CallerWorkspaceMembershipPublic)`

SetWorkspaceMemberships sets WorkspaceMemberships field to given value.

### HasWorkspaceMemberships

`func (o *OrganizationMembershipContextPublic) HasWorkspaceMemberships() bool`

HasWorkspaceMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


