# AuthenticatePasskeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | **map[string]interface{}** | The browser&#39;s PublicKeyCredential assertion, serialized. | 

## Methods

### NewAuthenticatePasskeyRequest

`func NewAuthenticatePasskeyRequest(credential map[string]interface{}, ) *AuthenticatePasskeyRequest`

NewAuthenticatePasskeyRequest instantiates a new AuthenticatePasskeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatePasskeyRequestWithDefaults

`func NewAuthenticatePasskeyRequestWithDefaults() *AuthenticatePasskeyRequest`

NewAuthenticatePasskeyRequestWithDefaults instantiates a new AuthenticatePasskeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *AuthenticatePasskeyRequest) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AuthenticatePasskeyRequest) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AuthenticatePasskeyRequest) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


