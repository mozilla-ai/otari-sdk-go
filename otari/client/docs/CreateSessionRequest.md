# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The identity&#39;s sign-in address. | [optional] 
**MasterKey** | Pointer to **NullableString** | The gateway master key; verified once and never stored by the browser. Accepted only while the operator identity has no password, which is to say while nobody has claimed this deployment (see GET /v1/bootstrap). | [optional] 
**Password** | Pointer to **NullableString** | The identity&#39;s password. | [optional] 

## Methods

### NewCreateSessionRequest

`func NewCreateSessionRequest() *CreateSessionRequest`

NewCreateSessionRequest instantiates a new CreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionRequestWithDefaults

`func NewCreateSessionRequestWithDefaults() *CreateSessionRequest`

NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateSessionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateSessionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateSessionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateSessionRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateSessionRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateSessionRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMasterKey

`func (o *CreateSessionRequest) GetMasterKey() string`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *CreateSessionRequest) GetMasterKeyOk() (*string, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *CreateSessionRequest) SetMasterKey(v string)`

SetMasterKey sets MasterKey field to given value.

### HasMasterKey

`func (o *CreateSessionRequest) HasMasterKey() bool`

HasMasterKey returns a boolean if a field has been set.

### SetMasterKeyNil

`func (o *CreateSessionRequest) SetMasterKeyNil(b bool)`

 SetMasterKeyNil sets the value for MasterKey to be an explicit nil

### UnsetMasterKey
`func (o *CreateSessionRequest) UnsetMasterKey()`

UnsetMasterKey ensures that no value is present for MasterKey, not even an explicit nil
### GetPassword

`func (o *CreateSessionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSessionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSessionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSessionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateSessionRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateSessionRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


