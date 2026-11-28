# OIDCConnectionDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Enabled** | **bool** | Whether the domain&#39;s claim is turned on. | 
**ProofExpiresAt** | **NullableTime** | When the domain&#39;s DNS proof stops being acted on, or None when it has never been verified. | 

## Methods

### NewOIDCConnectionDomain

`func NewOIDCConnectionDomain(domain string, enabled bool, proofExpiresAt NullableTime, ) *OIDCConnectionDomain`

NewOIDCConnectionDomain instantiates a new OIDCConnectionDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCConnectionDomainWithDefaults

`func NewOIDCConnectionDomainWithDefaults() *OIDCConnectionDomain`

NewOIDCConnectionDomainWithDefaults instantiates a new OIDCConnectionDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *OIDCConnectionDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OIDCConnectionDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OIDCConnectionDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetEnabled

`func (o *OIDCConnectionDomain) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OIDCConnectionDomain) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OIDCConnectionDomain) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProofExpiresAt

`func (o *OIDCConnectionDomain) GetProofExpiresAt() time.Time`

GetProofExpiresAt returns the ProofExpiresAt field if non-nil, zero value otherwise.

### GetProofExpiresAtOk

`func (o *OIDCConnectionDomain) GetProofExpiresAtOk() (*time.Time, bool)`

GetProofExpiresAtOk returns a tuple with the ProofExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofExpiresAt

`func (o *OIDCConnectionDomain) SetProofExpiresAt(v time.Time)`

SetProofExpiresAt sets ProofExpiresAt field to given value.


### SetProofExpiresAtNil

`func (o *OIDCConnectionDomain) SetProofExpiresAtNil(b bool)`

 SetProofExpiresAtNil sets the value for ProofExpiresAt to be an explicit nil

### UnsetProofExpiresAt
`func (o *OIDCConnectionDomain) UnsetProofExpiresAt()`

UnsetProofExpiresAt ensures that no value is present for ProofExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


