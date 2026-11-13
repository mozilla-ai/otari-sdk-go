# SsoCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The authorization code from the provider&#39;s redirect. | 
**Iss** | Pointer to **NullableString** | The &#39;iss&#39; from the provider&#39;s redirect (RFC 9207), where it sent one. Checked against this connection&#39;s issuer when its discovery document advertises the parameter; a provider that does not advertise it sends none. | [optional] 
**State** | **string** | The &#39;state&#39; from the provider&#39;s redirect, as issued by /authorize. | 
**TermsAccepted** | Pointer to **bool** | Whether the person accepted this deployment&#39;s terms before going to the provider. Required where the bootstrap reports a terms_url and the sign-in is the first way in to an account nobody has claimed yet; ignored otherwise. | [optional] [default to false]

## Methods

### NewSsoCallbackRequest

`func NewSsoCallbackRequest(code string, state string, ) *SsoCallbackRequest`

NewSsoCallbackRequest instantiates a new SsoCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoCallbackRequestWithDefaults

`func NewSsoCallbackRequestWithDefaults() *SsoCallbackRequest`

NewSsoCallbackRequestWithDefaults instantiates a new SsoCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SsoCallbackRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SsoCallbackRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SsoCallbackRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetIss

`func (o *SsoCallbackRequest) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SsoCallbackRequest) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SsoCallbackRequest) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *SsoCallbackRequest) HasIss() bool`

HasIss returns a boolean if a field has been set.

### SetIssNil

`func (o *SsoCallbackRequest) SetIssNil(b bool)`

 SetIssNil sets the value for Iss to be an explicit nil

### UnsetIss
`func (o *SsoCallbackRequest) UnsetIss()`

UnsetIss ensures that no value is present for Iss, not even an explicit nil
### GetState

`func (o *SsoCallbackRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SsoCallbackRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SsoCallbackRequest) SetState(v string)`

SetState sets State field to given value.


### GetTermsAccepted

`func (o *SsoCallbackRequest) GetTermsAccepted() bool`

GetTermsAccepted returns the TermsAccepted field if non-nil, zero value otherwise.

### GetTermsAcceptedOk

`func (o *SsoCallbackRequest) GetTermsAcceptedOk() (*bool, bool)`

GetTermsAcceptedOk returns a tuple with the TermsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAccepted

`func (o *SsoCallbackRequest) SetTermsAccepted(v bool)`

SetTermsAccepted sets TermsAccepted field to given value.

### HasTermsAccepted

`func (o *SsoCallbackRequest) HasTermsAccepted() bool`

HasTermsAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


