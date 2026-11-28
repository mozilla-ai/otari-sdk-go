# AccessTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** | How long the token lives from now, at most a year. | [optional] [default to "90-days"]
**Name** | **string** | What the token is for. | 

## Methods

### NewAccessTokenCreate

`func NewAccessTokenCreate(name string, ) *AccessTokenCreate`

NewAccessTokenCreate instantiates a new AccessTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCreateWithDefaults

`func NewAccessTokenCreateWithDefaults() *AccessTokenCreate`

NewAccessTokenCreateWithDefaults instantiates a new AccessTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *AccessTokenCreate) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AccessTokenCreate) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AccessTokenCreate) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AccessTokenCreate) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetName

`func (o *AccessTokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


