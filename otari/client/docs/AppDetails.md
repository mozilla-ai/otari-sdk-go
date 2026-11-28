# AppDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to **string** | What the app is given: an API key billed to a workspace the person picks, or, for a client that registered itself such as an MCP client, the person&#39;s own read-only access token. | [optional] [default to "api_key"]
**Issuer** | **string** |  | 
**Name** | **string** |  | 
**RedirectUri** | **string** |  | 
**Scopes** | [**[]KeyScope**](KeyScope.md) | The scopes each key this app is issued will hold, from its registration. | 
**SpendLimits** | **bool** | Whether a spending limit the person sets binds the key. It does not where the deployment has turned budget enforcement off, so the consent screen does not offer one. | 
**Verified** | **bool** | Whether the operator registered this app. An unregistered app is named after its site, which nothing has checked beyond its callback being on it. | 

## Methods

### NewAppDetails

`func NewAppDetails(issuer string, name string, redirectUri string, scopes []KeyScope, spendLimits bool, verified bool, ) *AppDetails`

NewAppDetails instantiates a new AppDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDetailsWithDefaults

`func NewAppDetailsWithDefaults() *AppDetails`

NewAppDetailsWithDefaults instantiates a new AppDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AppDetails) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AppDetails) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AppDetails) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AppDetails) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetIssuer

`func (o *AppDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AppDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AppDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetName

`func (o *AppDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDetails) SetName(v string)`

SetName sets Name field to given value.


### GetRedirectUri

`func (o *AppDetails) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AppDetails) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AppDetails) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetScopes

`func (o *AppDetails) GetScopes() []KeyScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppDetails) GetScopesOk() (*[]KeyScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppDetails) SetScopes(v []KeyScope)`

SetScopes sets Scopes field to given value.


### GetSpendLimits

`func (o *AppDetails) GetSpendLimits() bool`

GetSpendLimits returns the SpendLimits field if non-nil, zero value otherwise.

### GetSpendLimitsOk

`func (o *AppDetails) GetSpendLimitsOk() (*bool, bool)`

GetSpendLimitsOk returns a tuple with the SpendLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimits

`func (o *AppDetails) SetSpendLimits(v bool)`

SetSpendLimits sets SpendLimits field to given value.


### GetVerified

`func (o *AppDetails) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AppDetails) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AppDetails) SetVerified(v bool)`

SetVerified sets Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


