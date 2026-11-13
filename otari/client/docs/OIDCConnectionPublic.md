# OIDCConnectionPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecretLast4** | **string** |  | 
**DefaultRole** | **string** |  | 
**DiscoveredAt** | **time.Time** | When the provider&#39;s configuration was last read. | 
**DisplayName** | **string** |  | 
**Domains** | [**[]OIDCConnectionDomain**](OIDCConnectionDomain.md) | Every one of the organization&#39;s email domains, each with its DNS proof&#39;s expiry and whether its claim is enabled. Only a domain that is both enabled and has an unexpired proof admits a sign-in. | 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**Issuer** | **string** | The organization&#39;s own claim. | 
**ResolvedIssuer** | **string** | What discovery confirmed the issuer to be. | 

## Methods

### NewOIDCConnectionPublic

`func NewOIDCConnectionPublic(clientId string, clientSecretLast4 string, defaultRole string, discoveredAt time.Time, displayName string, domains []OIDCConnectionDomain, enabled bool, id string, issuer string, resolvedIssuer string, ) *OIDCConnectionPublic`

NewOIDCConnectionPublic instantiates a new OIDCConnectionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCConnectionPublicWithDefaults

`func NewOIDCConnectionPublicWithDefaults() *OIDCConnectionPublic`

NewOIDCConnectionPublicWithDefaults instantiates a new OIDCConnectionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OIDCConnectionPublic) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCConnectionPublic) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCConnectionPublic) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecretLast4

`func (o *OIDCConnectionPublic) GetClientSecretLast4() string`

GetClientSecretLast4 returns the ClientSecretLast4 field if non-nil, zero value otherwise.

### GetClientSecretLast4Ok

`func (o *OIDCConnectionPublic) GetClientSecretLast4Ok() (*string, bool)`

GetClientSecretLast4Ok returns a tuple with the ClientSecretLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretLast4

`func (o *OIDCConnectionPublic) SetClientSecretLast4(v string)`

SetClientSecretLast4 sets ClientSecretLast4 field to given value.


### GetDefaultRole

`func (o *OIDCConnectionPublic) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *OIDCConnectionPublic) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *OIDCConnectionPublic) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.


### GetDiscoveredAt

`func (o *OIDCConnectionPublic) GetDiscoveredAt() time.Time`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *OIDCConnectionPublic) GetDiscoveredAtOk() (*time.Time, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *OIDCConnectionPublic) SetDiscoveredAt(v time.Time)`

SetDiscoveredAt sets DiscoveredAt field to given value.


### GetDisplayName

`func (o *OIDCConnectionPublic) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OIDCConnectionPublic) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OIDCConnectionPublic) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDomains

`func (o *OIDCConnectionPublic) GetDomains() []OIDCConnectionDomain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OIDCConnectionPublic) GetDomainsOk() (*[]OIDCConnectionDomain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OIDCConnectionPublic) SetDomains(v []OIDCConnectionDomain)`

SetDomains sets Domains field to given value.


### GetEnabled

`func (o *OIDCConnectionPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OIDCConnectionPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OIDCConnectionPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *OIDCConnectionPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OIDCConnectionPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OIDCConnectionPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIssuer

`func (o *OIDCConnectionPublic) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCConnectionPublic) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCConnectionPublic) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetResolvedIssuer

`func (o *OIDCConnectionPublic) GetResolvedIssuer() string`

GetResolvedIssuer returns the ResolvedIssuer field if non-nil, zero value otherwise.

### GetResolvedIssuerOk

`func (o *OIDCConnectionPublic) GetResolvedIssuerOk() (*string, bool)`

GetResolvedIssuerOk returns a tuple with the ResolvedIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedIssuer

`func (o *OIDCConnectionPublic) SetResolvedIssuer(v string)`

SetResolvedIssuer sets ResolvedIssuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


