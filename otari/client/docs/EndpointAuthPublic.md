# EndpointAuthPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** |  | 
**HeaderName** | **string** |  | 
**Last4** | **string** | The last four characters of the stored value. | 
**Method** | [**EndpointAuthMethod**](EndpointAuthMethod.md) | How the credential is produced. | 
**State** | [**EndpointAuthState**](EndpointAuthState.md) | How the credential stands against its own expiry. Reported so every reader draws the line in the same place. No value of it refuses a request: the proxy answers that on every request. | 

## Methods

### NewEndpointAuthPublic

`func NewEndpointAuthPublic(expiresAt time.Time, headerName string, last4 string, method EndpointAuthMethod, state EndpointAuthState, ) *EndpointAuthPublic`

NewEndpointAuthPublic instantiates a new EndpointAuthPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAuthPublicWithDefaults

`func NewEndpointAuthPublicWithDefaults() *EndpointAuthPublic`

NewEndpointAuthPublicWithDefaults instantiates a new EndpointAuthPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *EndpointAuthPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EndpointAuthPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EndpointAuthPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHeaderName

`func (o *EndpointAuthPublic) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *EndpointAuthPublic) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *EndpointAuthPublic) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### GetLast4

`func (o *EndpointAuthPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *EndpointAuthPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *EndpointAuthPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetMethod

`func (o *EndpointAuthPublic) GetMethod() EndpointAuthMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EndpointAuthPublic) GetMethodOk() (*EndpointAuthMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EndpointAuthPublic) SetMethod(v EndpointAuthMethod)`

SetMethod sets Method field to given value.


### GetState

`func (o *EndpointAuthPublic) GetState() EndpointAuthState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EndpointAuthPublic) GetStateOk() (*EndpointAuthState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EndpointAuthPublic) SetState(v EndpointAuthState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


