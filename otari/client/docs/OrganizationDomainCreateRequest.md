# OrganizationDomainCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **string** |  | [optional] [default to "member"]
**Domain** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewOrganizationDomainCreateRequest

`func NewOrganizationDomainCreateRequest(domain string, ) *OrganizationDomainCreateRequest`

NewOrganizationDomainCreateRequest instantiates a new OrganizationDomainCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDomainCreateRequestWithDefaults

`func NewOrganizationDomainCreateRequestWithDefaults() *OrganizationDomainCreateRequest`

NewOrganizationDomainCreateRequestWithDefaults instantiates a new OrganizationDomainCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *OrganizationDomainCreateRequest) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *OrganizationDomainCreateRequest) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *OrganizationDomainCreateRequest) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *OrganizationDomainCreateRequest) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationDomainCreateRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationDomainCreateRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationDomainCreateRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnabled

`func (o *OrganizationDomainCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationDomainCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationDomainCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationDomainCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


