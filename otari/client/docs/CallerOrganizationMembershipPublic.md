# CallerOrganizationMembershipPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActiveOrganization** | Pointer to **bool** |  | [optional] [default to false]
**Organization** | [**OrganizationPublic**](OrganizationPublic.md) |  | 
**OrganizationMemberId** | **string** |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewCallerOrganizationMembershipPublic

`func NewCallerOrganizationMembershipPublic(organization OrganizationPublic, organizationMemberId string, role string, status string, ) *CallerOrganizationMembershipPublic`

NewCallerOrganizationMembershipPublic instantiates a new CallerOrganizationMembershipPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerOrganizationMembershipPublicWithDefaults

`func NewCallerOrganizationMembershipPublicWithDefaults() *CallerOrganizationMembershipPublic`

NewCallerOrganizationMembershipPublicWithDefaults instantiates a new CallerOrganizationMembershipPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActiveOrganization

`func (o *CallerOrganizationMembershipPublic) GetIsActiveOrganization() bool`

GetIsActiveOrganization returns the IsActiveOrganization field if non-nil, zero value otherwise.

### GetIsActiveOrganizationOk

`func (o *CallerOrganizationMembershipPublic) GetIsActiveOrganizationOk() (*bool, bool)`

GetIsActiveOrganizationOk returns a tuple with the IsActiveOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveOrganization

`func (o *CallerOrganizationMembershipPublic) SetIsActiveOrganization(v bool)`

SetIsActiveOrganization sets IsActiveOrganization field to given value.

### HasIsActiveOrganization

`func (o *CallerOrganizationMembershipPublic) HasIsActiveOrganization() bool`

HasIsActiveOrganization returns a boolean if a field has been set.

### GetOrganization

`func (o *CallerOrganizationMembershipPublic) GetOrganization() OrganizationPublic`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CallerOrganizationMembershipPublic) GetOrganizationOk() (*OrganizationPublic, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CallerOrganizationMembershipPublic) SetOrganization(v OrganizationPublic)`

SetOrganization sets Organization field to given value.


### GetOrganizationMemberId

`func (o *CallerOrganizationMembershipPublic) GetOrganizationMemberId() string`

GetOrganizationMemberId returns the OrganizationMemberId field if non-nil, zero value otherwise.

### GetOrganizationMemberIdOk

`func (o *CallerOrganizationMembershipPublic) GetOrganizationMemberIdOk() (*string, bool)`

GetOrganizationMemberIdOk returns a tuple with the OrganizationMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMemberId

`func (o *CallerOrganizationMembershipPublic) SetOrganizationMemberId(v string)`

SetOrganizationMemberId sets OrganizationMemberId field to given value.


### GetRole

`func (o *CallerOrganizationMembershipPublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CallerOrganizationMembershipPublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CallerOrganizationMembershipPublic) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *CallerOrganizationMembershipPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallerOrganizationMembershipPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallerOrganizationMembershipPublic) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


