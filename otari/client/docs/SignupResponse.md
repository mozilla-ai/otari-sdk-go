# SignupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | What the caller should do next. | 
**Verified** | Pointer to **bool** | The account is verified already, through its invitation, so the caller can sign in now. | [optional] [default to false]

## Methods

### NewSignupResponse

`func NewSignupResponse(message string, ) *SignupResponse`

NewSignupResponse instantiates a new SignupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupResponseWithDefaults

`func NewSignupResponseWithDefaults() *SignupResponse`

NewSignupResponseWithDefaults instantiates a new SignupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SignupResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignupResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignupResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetVerified

`func (o *SignupResponse) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SignupResponse) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SignupResponse) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SignupResponse) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


