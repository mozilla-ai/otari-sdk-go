# OIDCConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The OAuth client ID this deployment authenticates with. | 
**ClientSecret** | Pointer to **NullableString** | The OAuth client secret. Stored encrypted and never served back. Leave it out to keep the one already stored, which is refused unless the issuer and the client ID are both unchanged. | [optional] 
**DefaultRole** | Pointer to **string** | The role given to a person&#39;s first sign-in through this connection. | [optional] [default to "member"]
**DisplayName** | **string** | The identity provider&#39;s name, as shown to people signing in. | 
**Issuer** | **string** | The identity provider&#39;s issuer URL, as the organization&#39;s own claim. | 

## Methods

### NewOIDCConnectionRequest

`func NewOIDCConnectionRequest(clientId string, displayName string, issuer string, ) *OIDCConnectionRequest`

NewOIDCConnectionRequest instantiates a new OIDCConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCConnectionRequestWithDefaults

`func NewOIDCConnectionRequestWithDefaults() *OIDCConnectionRequest`

NewOIDCConnectionRequestWithDefaults instantiates a new OIDCConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OIDCConnectionRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCConnectionRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCConnectionRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OIDCConnectionRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OIDCConnectionRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OIDCConnectionRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OIDCConnectionRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OIDCConnectionRequest) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OIDCConnectionRequest) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetDefaultRole

`func (o *OIDCConnectionRequest) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *OIDCConnectionRequest) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *OIDCConnectionRequest) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *OIDCConnectionRequest) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetDisplayName

`func (o *OIDCConnectionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OIDCConnectionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OIDCConnectionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIssuer

`func (o *OIDCConnectionRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OIDCConnectionRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OIDCConnectionRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


