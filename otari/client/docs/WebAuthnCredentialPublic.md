# WebAuthnCredentialPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackedUp** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**CredentialId** | **string** |  | 
**Id** | **string** |  | 
**IsUsable** | **bool** |  | 
**LastUsedAt** | **NullableTime** |  | 
**Name** | **string** |  | 
**RpId** | **string** |  | 
**Transports** | **[]string** |  | 

## Methods

### NewWebAuthnCredentialPublic

`func NewWebAuthnCredentialPublic(backedUp bool, createdAt time.Time, credentialId string, id string, isUsable bool, lastUsedAt NullableTime, name string, rpId string, transports []string, ) *WebAuthnCredentialPublic`

NewWebAuthnCredentialPublic instantiates a new WebAuthnCredentialPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnCredentialPublicWithDefaults

`func NewWebAuthnCredentialPublicWithDefaults() *WebAuthnCredentialPublic`

NewWebAuthnCredentialPublicWithDefaults instantiates a new WebAuthnCredentialPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackedUp

`func (o *WebAuthnCredentialPublic) GetBackedUp() bool`

GetBackedUp returns the BackedUp field if non-nil, zero value otherwise.

### GetBackedUpOk

`func (o *WebAuthnCredentialPublic) GetBackedUpOk() (*bool, bool)`

GetBackedUpOk returns a tuple with the BackedUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackedUp

`func (o *WebAuthnCredentialPublic) SetBackedUp(v bool)`

SetBackedUp sets BackedUp field to given value.


### GetCreatedAt

`func (o *WebAuthnCredentialPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebAuthnCredentialPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebAuthnCredentialPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCredentialId

`func (o *WebAuthnCredentialPublic) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *WebAuthnCredentialPublic) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *WebAuthnCredentialPublic) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetId

`func (o *WebAuthnCredentialPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebAuthnCredentialPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebAuthnCredentialPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIsUsable

`func (o *WebAuthnCredentialPublic) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *WebAuthnCredentialPublic) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *WebAuthnCredentialPublic) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.


### GetLastUsedAt

`func (o *WebAuthnCredentialPublic) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *WebAuthnCredentialPublic) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *WebAuthnCredentialPublic) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *WebAuthnCredentialPublic) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *WebAuthnCredentialPublic) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetName

`func (o *WebAuthnCredentialPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebAuthnCredentialPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebAuthnCredentialPublic) SetName(v string)`

SetName sets Name field to given value.


### GetRpId

`func (o *WebAuthnCredentialPublic) GetRpId() string`

GetRpId returns the RpId field if non-nil, zero value otherwise.

### GetRpIdOk

`func (o *WebAuthnCredentialPublic) GetRpIdOk() (*string, bool)`

GetRpIdOk returns a tuple with the RpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpId

`func (o *WebAuthnCredentialPublic) SetRpId(v string)`

SetRpId sets RpId field to given value.


### GetTransports

`func (o *WebAuthnCredentialPublic) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *WebAuthnCredentialPublic) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *WebAuthnCredentialPublic) SetTransports(v []string)`

SetTransports sets Transports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


