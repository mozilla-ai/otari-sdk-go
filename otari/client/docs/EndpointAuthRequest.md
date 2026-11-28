# EndpointAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** | When the organization expects to rotate this credential. Must carry a UTC offset and be in the future. It is a rotation date, not a switch: a request is never refused because it passed, since the proxy answers that itself on every request. | 
**HeaderName** | **string** | The header the value is sent in. Letters, digits and hyphens only. A header a provider sends its own credential in, one reserved for the connection, &#x60;Host&#x60;, and Otari&#39;s own namespace are all refused. | 
**Method** | [**EndpointAuthMethod**](EndpointAuthMethod.md) | How the credential is produced. &#x60;static_header&#x60; sends one stored value as one header. | 
**Secret** | Pointer to **NullableString** | The value to send, at least 8 characters. Write-only: responses report only its last four characters. Required when the credential is first set, and on any change of base URL. Omit it to keep the value already stored; an empty string is refused rather than read as keeping it. | [optional] 

## Methods

### NewEndpointAuthRequest

`func NewEndpointAuthRequest(expiresAt time.Time, headerName string, method EndpointAuthMethod, ) *EndpointAuthRequest`

NewEndpointAuthRequest instantiates a new EndpointAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAuthRequestWithDefaults

`func NewEndpointAuthRequestWithDefaults() *EndpointAuthRequest`

NewEndpointAuthRequestWithDefaults instantiates a new EndpointAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *EndpointAuthRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EndpointAuthRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EndpointAuthRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHeaderName

`func (o *EndpointAuthRequest) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *EndpointAuthRequest) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *EndpointAuthRequest) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetMethod

`func (o *EndpointAuthRequest) GetMethod() EndpointAuthMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EndpointAuthRequest) GetMethodOk() (*EndpointAuthMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EndpointAuthRequest) SetMethod(v EndpointAuthMethod)`

SetMethod sets Method field to given value.


### GetSecret

`func (o *EndpointAuthRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EndpointAuthRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EndpointAuthRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *EndpointAuthRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *EndpointAuthRequest) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *EndpointAuthRequest) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


