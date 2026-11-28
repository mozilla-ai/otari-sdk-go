# VerifyEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password chosen at signup. The link verifies the address only together with it, so a link that reaches someone who did not sign up verifies nothing. | 
**Token** | **string** | The token from the verification link. | 

## Methods

### NewVerifyEmailRequest

`func NewVerifyEmailRequest(password string, token string, ) *VerifyEmailRequest`

NewVerifyEmailRequest instantiates a new VerifyEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEmailRequestWithDefaults

`func NewVerifyEmailRequestWithDefaults() *VerifyEmailRequest`

NewVerifyEmailRequestWithDefaults instantiates a new VerifyEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *VerifyEmailRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VerifyEmailRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VerifyEmailRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *VerifyEmailRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VerifyEmailRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VerifyEmailRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


