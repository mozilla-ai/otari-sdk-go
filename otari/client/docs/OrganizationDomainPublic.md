# OrganizationDomainPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**DefaultRole** | Pointer to **string** |  | [optional] [default to "member"]
**Domain** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Id** | **string** |  | 
**OrganizationId** | **string** |  | 
**ProofExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**VerificationRecord** | **string** |  | 
**VerifiedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrganizationDomainPublic

`func NewOrganizationDomainPublic(createdAt time.Time, domain string, id string, organizationId string, verificationRecord string, ) *OrganizationDomainPublic`

NewOrganizationDomainPublic instantiates a new OrganizationDomainPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDomainPublicWithDefaults

`func NewOrganizationDomainPublicWithDefaults() *OrganizationDomainPublic`

NewOrganizationDomainPublicWithDefaults instantiates a new OrganizationDomainPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrganizationDomainPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDomainPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDomainPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultRole

`func (o *OrganizationDomainPublic) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *OrganizationDomainPublic) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *OrganizationDomainPublic) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *OrganizationDomainPublic) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationDomainPublic) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationDomainPublic) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationDomainPublic) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnabled

`func (o *OrganizationDomainPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationDomainPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationDomainPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationDomainPublic) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *OrganizationDomainPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDomainPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDomainPublic) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *OrganizationDomainPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationDomainPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationDomainPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProofExpiresAt

`func (o *OrganizationDomainPublic) GetProofExpiresAt() time.Time`

GetProofExpiresAt returns the ProofExpiresAt field if non-nil, zero value otherwise.

### GetProofExpiresAtOk

`func (o *OrganizationDomainPublic) GetProofExpiresAtOk() (*time.Time, bool)`

GetProofExpiresAtOk returns a tuple with the ProofExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofExpiresAt

`func (o *OrganizationDomainPublic) SetProofExpiresAt(v time.Time)`

SetProofExpiresAt sets ProofExpiresAt field to given value.

### HasProofExpiresAt

`func (o *OrganizationDomainPublic) HasProofExpiresAt() bool`

HasProofExpiresAt returns a boolean if a field has been set.

### SetProofExpiresAtNil

`func (o *OrganizationDomainPublic) SetProofExpiresAtNil(b bool)`

 SetProofExpiresAtNil sets the value for ProofExpiresAt to be an explicit nil

### UnsetProofExpiresAt
`func (o *OrganizationDomainPublic) UnsetProofExpiresAt()`

UnsetProofExpiresAt ensures that no value is present for ProofExpiresAt, not even an explicit nil
### GetUpdatedAt

`func (o *OrganizationDomainPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationDomainPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationDomainPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationDomainPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OrganizationDomainPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OrganizationDomainPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetVerificationRecord

`func (o *OrganizationDomainPublic) GetVerificationRecord() string`

GetVerificationRecord returns the VerificationRecord field if non-nil, zero value otherwise.

### GetVerificationRecordOk

`func (o *OrganizationDomainPublic) GetVerificationRecordOk() (*string, bool)`

GetVerificationRecordOk returns a tuple with the VerificationRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRecord

`func (o *OrganizationDomainPublic) SetVerificationRecord(v string)`

SetVerificationRecord sets VerificationRecord field to given value.


### GetVerifiedAt

`func (o *OrganizationDomainPublic) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *OrganizationDomainPublic) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *OrganizationDomainPublic) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *OrganizationDomainPublic) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *OrganizationDomainPublic) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *OrganizationDomainPublic) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


