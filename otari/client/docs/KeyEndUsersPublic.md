# KeyEndUsersPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The live end users the cap counts: those on the budgets the key assigns, and those on none when the key has a list and no default. A key that assigns no budget counts every live end user of its user | 
**MaxEndUsers** | **NullableInt32** | The key&#39;s cap; null when it has none. A request naming a new end user is refused with end_user_limit_reached once count reaches it | 

## Methods

### NewKeyEndUsersPublic

`func NewKeyEndUsersPublic(count int32, maxEndUsers NullableInt32, ) *KeyEndUsersPublic`

NewKeyEndUsersPublic instantiates a new KeyEndUsersPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyEndUsersPublicWithDefaults

`func NewKeyEndUsersPublicWithDefaults() *KeyEndUsersPublic`

NewKeyEndUsersPublicWithDefaults instantiates a new KeyEndUsersPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KeyEndUsersPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KeyEndUsersPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KeyEndUsersPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetMaxEndUsers

`func (o *KeyEndUsersPublic) GetMaxEndUsers() int32`

GetMaxEndUsers returns the MaxEndUsers field if non-nil, zero value otherwise.

### GetMaxEndUsersOk

`func (o *KeyEndUsersPublic) GetMaxEndUsersOk() (*int32, bool)`

GetMaxEndUsersOk returns a tuple with the MaxEndUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndUsers

`func (o *KeyEndUsersPublic) SetMaxEndUsers(v int32)`

SetMaxEndUsers sets MaxEndUsers field to given value.


### SetMaxEndUsersNil

`func (o *KeyEndUsersPublic) SetMaxEndUsersNil(b bool)`

 SetMaxEndUsersNil sets the value for MaxEndUsers to be an explicit nil

### UnsetMaxEndUsers
`func (o *KeyEndUsersPublic) UnsetMaxEndUsers()`

UnsetMaxEndUsers ensures that no value is present for MaxEndUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


