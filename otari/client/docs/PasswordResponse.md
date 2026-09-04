# PasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The address this identity signs in with. | 
**MasterKeySignInRetired** | **bool** | Whether POST /v1/auth/session has stopped accepting the master key as a dashboard login. True once the operator identity has a password, which is what claiming the deployment means; a member setting their own password leaves an unclaimed deployment on the master key. Either way the master key stays the credential for the management API. | 

## Methods

### NewPasswordResponse

`func NewPasswordResponse(email string, masterKeySignInRetired bool, ) *PasswordResponse`

NewPasswordResponse instantiates a new PasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordResponseWithDefaults

`func NewPasswordResponseWithDefaults() *PasswordResponse`

NewPasswordResponseWithDefaults instantiates a new PasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PasswordResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PasswordResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PasswordResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetMasterKeySignInRetired

`func (o *PasswordResponse) GetMasterKeySignInRetired() bool`

GetMasterKeySignInRetired returns the MasterKeySignInRetired field if non-nil, zero value otherwise.

### GetMasterKeySignInRetiredOk

`func (o *PasswordResponse) GetMasterKeySignInRetiredOk() (*bool, bool)`

GetMasterKeySignInRetiredOk returns a tuple with the MasterKeySignInRetired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeySignInRetired

`func (o *PasswordResponse) SetMasterKeySignInRetired(v bool)`

SetMasterKeySignInRetired sets MasterKeySignInRetired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


