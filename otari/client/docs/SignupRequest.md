# SignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The address an admin added or invited. | 
**FullName** | Pointer to **NullableString** | Filled in only if not already set. | [optional] 
**Password** | **string** | The password to sign in with once verified. At least 8 characters, at most 72 bytes. | 
**TermsAccepted** | Pointer to **bool** | Whether the caller accepted this deployment&#39;s terms. | [optional] [default to false]

## Methods

### NewSignupRequest

`func NewSignupRequest(email string, password string, ) *SignupRequest`

NewSignupRequest instantiates a new SignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupRequestWithDefaults

`func NewSignupRequestWithDefaults() *SignupRequest`

NewSignupRequestWithDefaults instantiates a new SignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *SignupRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SignupRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SignupRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SignupRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *SignupRequest) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *SignupRequest) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPassword

`func (o *SignupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTermsAccepted

`func (o *SignupRequest) GetTermsAccepted() bool`

GetTermsAccepted returns the TermsAccepted field if non-nil, zero value otherwise.

### GetTermsAcceptedOk

`func (o *SignupRequest) GetTermsAcceptedOk() (*bool, bool)`

GetTermsAcceptedOk returns a tuple with the TermsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAccepted

`func (o *SignupRequest) SetTermsAccepted(v bool)`

SetTermsAccepted sets TermsAccepted field to given value.

### HasTermsAccepted

`func (o *SignupRequest) HasTermsAccepted() bool`

HasTermsAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


