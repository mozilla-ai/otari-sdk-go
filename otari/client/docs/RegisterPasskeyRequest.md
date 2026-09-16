# RegisterPasskeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | **map[string]interface{}** | The browser&#39;s PublicKeyCredential, serialized. | 
**Name** | Pointer to **NullableString** | What to call this passkey in the credential list. Optional: an unnamed one is numbered rather than refused, so a browser that offers no prompt still works. | [optional] 

## Methods

### NewRegisterPasskeyRequest

`func NewRegisterPasskeyRequest(credential map[string]interface{}, ) *RegisterPasskeyRequest`

NewRegisterPasskeyRequest instantiates a new RegisterPasskeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPasskeyRequestWithDefaults

`func NewRegisterPasskeyRequestWithDefaults() *RegisterPasskeyRequest`

NewRegisterPasskeyRequestWithDefaults instantiates a new RegisterPasskeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *RegisterPasskeyRequest) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *RegisterPasskeyRequest) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *RegisterPasskeyRequest) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.


### GetName

`func (o *RegisterPasskeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterPasskeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterPasskeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterPasskeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegisterPasskeyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegisterPasskeyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


