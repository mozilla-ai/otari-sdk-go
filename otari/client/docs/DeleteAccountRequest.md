# DeleteAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmEmail** | **string** | The account&#39;s email address, typed to confirm the deletion. | 
**CurrentPassword** | Pointer to **NullableString** | The account&#39;s password. Required when the account has one. | [optional] 

## Methods

### NewDeleteAccountRequest

`func NewDeleteAccountRequest(confirmEmail string, ) *DeleteAccountRequest`

NewDeleteAccountRequest instantiates a new DeleteAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAccountRequestWithDefaults

`func NewDeleteAccountRequestWithDefaults() *DeleteAccountRequest`

NewDeleteAccountRequestWithDefaults instantiates a new DeleteAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmEmail

`func (o *DeleteAccountRequest) GetConfirmEmail() string`

GetConfirmEmail returns the ConfirmEmail field if non-nil, zero value otherwise.

### GetConfirmEmailOk

`func (o *DeleteAccountRequest) GetConfirmEmailOk() (*string, bool)`

GetConfirmEmailOk returns a tuple with the ConfirmEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmEmail

`func (o *DeleteAccountRequest) SetConfirmEmail(v string)`

SetConfirmEmail sets ConfirmEmail field to given value.


### GetCurrentPassword

`func (o *DeleteAccountRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *DeleteAccountRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *DeleteAccountRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *DeleteAccountRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *DeleteAccountRequest) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *DeleteAccountRequest) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


