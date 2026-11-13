# OrganizationSsoPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [**NullableOIDCConnectionPublic**](OIDCConnectionPublic.md) | The organization&#39;s own connection, or null until one is saved. | 
**Enforcement** | [**OrganizationSsoEnforcementPublic**](OrganizationSsoEnforcementPublic.md) | The organization&#39;s current sign-in enforcement. Independent of whether a connection exists. | 
**RedirectUri** | **NullableString** | The absolute callback address to register at the identity provider, whether or not a connection is saved yet. Null where this deployment has no public address of its own configured, because there is no address to register until it has one. | 

## Methods

### NewOrganizationSsoPublic

`func NewOrganizationSsoPublic(connection NullableOIDCConnectionPublic, enforcement OrganizationSsoEnforcementPublic, redirectUri NullableString, ) *OrganizationSsoPublic`

NewOrganizationSsoPublic instantiates a new OrganizationSsoPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSsoPublicWithDefaults

`func NewOrganizationSsoPublicWithDefaults() *OrganizationSsoPublic`

NewOrganizationSsoPublicWithDefaults instantiates a new OrganizationSsoPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *OrganizationSsoPublic) GetConnection() OIDCConnectionPublic`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *OrganizationSsoPublic) GetConnectionOk() (*OIDCConnectionPublic, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *OrganizationSsoPublic) SetConnection(v OIDCConnectionPublic)`

SetConnection sets Connection field to given value.


### SetConnectionNil

`func (o *OrganizationSsoPublic) SetConnectionNil(b bool)`

 SetConnectionNil sets the value for Connection to be an explicit nil

### UnsetConnection
`func (o *OrganizationSsoPublic) UnsetConnection()`

UnsetConnection ensures that no value is present for Connection, not even an explicit nil
### GetEnforcement

`func (o *OrganizationSsoPublic) GetEnforcement() OrganizationSsoEnforcementPublic`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *OrganizationSsoPublic) GetEnforcementOk() (*OrganizationSsoEnforcementPublic, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *OrganizationSsoPublic) SetEnforcement(v OrganizationSsoEnforcementPublic)`

SetEnforcement sets Enforcement field to given value.


### GetRedirectUri

`func (o *OrganizationSsoPublic) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OrganizationSsoPublic) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OrganizationSsoPublic) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### SetRedirectUriNil

`func (o *OrganizationSsoPublic) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *OrganizationSsoPublic) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


