# ResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** | The password to sign in with from now on. At least 8 characters, at most 72 bytes. | 
**TermsAccepted** | Pointer to **bool** | Whether the caller accepted this deployment&#39;s terms. Required where the bootstrap reports a terms_url and the account has not accepted them yet. | [optional] [default to false]
**Token** | **string** | The token from the reset link. | 

## Methods

### NewResetPasswordRequest

`func NewResetPasswordRequest(newPassword string, token string, ) *ResetPasswordRequest`

NewResetPasswordRequest instantiates a new ResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordRequestWithDefaults

`func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest`

NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *ResetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetTermsAccepted

`func (o *ResetPasswordRequest) GetTermsAccepted() bool`

GetTermsAccepted returns the TermsAccepted field if non-nil, zero value otherwise.

### GetTermsAcceptedOk

`func (o *ResetPasswordRequest) GetTermsAcceptedOk() (*bool, bool)`

GetTermsAcceptedOk returns a tuple with the TermsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAccepted

`func (o *ResetPasswordRequest) SetTermsAccepted(v bool)`

SetTermsAccepted sets TermsAccepted field to given value.

### HasTermsAccepted

`func (o *ResetPasswordRequest) HasTermsAccepted() bool`

HasTermsAccepted returns a boolean if a field has been set.

### GetToken

`func (o *ResetPasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetPasswordRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetPasswordRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


