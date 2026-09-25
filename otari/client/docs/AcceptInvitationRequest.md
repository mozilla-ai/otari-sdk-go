# AcceptInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **NullableString** | Filled in only if not already set. | [optional] 
**Password** | Pointer to **NullableString** | Sets the invited identity&#39;s password in the same step, when the preview reported needs_password. Needs no mail: the link is the proof, whether it was emailed or an admin handed it over. | [optional] 
**TermsAccepted** | Pointer to **bool** | Whether the caller accepted this deployment&#39;s terms. | [optional] [default to false]
**Token** | **string** |  | 

## Methods

### NewAcceptInvitationRequest

`func NewAcceptInvitationRequest(token string, ) *AcceptInvitationRequest`

NewAcceptInvitationRequest instantiates a new AcceptInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptInvitationRequestWithDefaults

`func NewAcceptInvitationRequestWithDefaults() *AcceptInvitationRequest`

NewAcceptInvitationRequestWithDefaults instantiates a new AcceptInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *AcceptInvitationRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AcceptInvitationRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AcceptInvitationRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *AcceptInvitationRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *AcceptInvitationRequest) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *AcceptInvitationRequest) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPassword

`func (o *AcceptInvitationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AcceptInvitationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AcceptInvitationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AcceptInvitationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AcceptInvitationRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AcceptInvitationRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTermsAccepted

`func (o *AcceptInvitationRequest) GetTermsAccepted() bool`

GetTermsAccepted returns the TermsAccepted field if non-nil, zero value otherwise.

### GetTermsAcceptedOk

`func (o *AcceptInvitationRequest) GetTermsAcceptedOk() (*bool, bool)`

GetTermsAcceptedOk returns a tuple with the TermsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAccepted

`func (o *AcceptInvitationRequest) SetTermsAccepted(v bool)`

SetTermsAccepted sets TermsAccepted field to given value.

### HasTermsAccepted

`func (o *AcceptInvitationRequest) HasTermsAccepted() bool`

HasTermsAccepted returns a boolean if a field has been set.

### GetToken

`func (o *AcceptInvitationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AcceptInvitationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AcceptInvitationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


