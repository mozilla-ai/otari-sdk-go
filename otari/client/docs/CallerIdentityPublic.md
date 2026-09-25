# CallerIdentityPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimsDeployment** | **bool** | Whether setting this identity&#39;s password claims the deployment, which stops the master key signing in to the dashboard. True for the deployment&#39;s operator until it holds a password, whether or not it already has an address; false for everybody else. | 
**Email** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**HasPassword** | **bool** | Whether this identity holds a dashboard password. False for one that signs in only through an OAuth provider or a passkey, and for a roster entry nobody has claimed yet. PUT /api/v1/auth/password requires current_password from a cookie-authenticated caller exactly while this is true. | 
**UserId** | **string** |  | 

## Methods

### NewCallerIdentityPublic

`func NewCallerIdentityPublic(claimsDeployment bool, hasPassword bool, userId string, ) *CallerIdentityPublic`

NewCallerIdentityPublic instantiates a new CallerIdentityPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerIdentityPublicWithDefaults

`func NewCallerIdentityPublicWithDefaults() *CallerIdentityPublic`

NewCallerIdentityPublicWithDefaults instantiates a new CallerIdentityPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimsDeployment

`func (o *CallerIdentityPublic) GetClaimsDeployment() bool`

GetClaimsDeployment returns the ClaimsDeployment field if non-nil, zero value otherwise.

### GetClaimsDeploymentOk

`func (o *CallerIdentityPublic) GetClaimsDeploymentOk() (*bool, bool)`

GetClaimsDeploymentOk returns a tuple with the ClaimsDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsDeployment

`func (o *CallerIdentityPublic) SetClaimsDeployment(v bool)`

SetClaimsDeployment sets ClaimsDeployment field to given value.


### GetEmail

`func (o *CallerIdentityPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CallerIdentityPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CallerIdentityPublic) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CallerIdentityPublic) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CallerIdentityPublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CallerIdentityPublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *CallerIdentityPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CallerIdentityPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CallerIdentityPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CallerIdentityPublic) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *CallerIdentityPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *CallerIdentityPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetHasPassword

`func (o *CallerIdentityPublic) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *CallerIdentityPublic) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *CallerIdentityPublic) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetUserId

`func (o *CallerIdentityPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CallerIdentityPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CallerIdentityPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


