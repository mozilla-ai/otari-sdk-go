# SetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **NullableString** | The password being replaced. Required when the identity already has one and the request is authenticated by the session cookie; ignored when the master key is sent in a header, which needs no proof of the old password (it still needs &#x60;email&#x60; when the identity has no sign-in address yet). | [optional] 
**Email** | Pointer to **NullableString** | The address to sign in with. Required when the identity has none, which is the state first boot leaves the operator in, including when the master key is what authenticates the call. Resubmitting the address the identity already holds is accepted and ignored; only a *different* address is refused, because changing one is not supported yet. | [optional] 
**NewPassword** | **string** | The password to sign in with from now on. At least 8 characters, and at most 72 bytes, which is bcrypt&#39;s ceiling. | 

## Methods

### NewSetPasswordRequest

`func NewSetPasswordRequest(newPassword string, ) *SetPasswordRequest`

NewSetPasswordRequest instantiates a new SetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPasswordRequestWithDefaults

`func NewSetPasswordRequestWithDefaults() *SetPasswordRequest`

NewSetPasswordRequestWithDefaults instantiates a new SetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *SetPasswordRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *SetPasswordRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *SetPasswordRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *SetPasswordRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *SetPasswordRequest) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *SetPasswordRequest) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetEmail

`func (o *SetPasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SetPasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SetPasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SetPasswordRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SetPasswordRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SetPasswordRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNewPassword

`func (o *SetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *SetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *SetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


