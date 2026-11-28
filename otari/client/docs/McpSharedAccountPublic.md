# McpSharedAccountPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Renewable** | **bool** | Whether Otari can renew the token on its own; false means signing in again once it expires | 
**SignedInAt** | **string** |  | 
**SignedInBy** | **NullableString** | The name, or else the email, of the dashboard account that signed in; null once it is gone | 

## Methods

### NewMcpSharedAccountPublic

`func NewMcpSharedAccountPublic(renewable bool, signedInAt string, signedInBy NullableString, ) *McpSharedAccountPublic`

NewMcpSharedAccountPublic instantiates a new McpSharedAccountPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpSharedAccountPublicWithDefaults

`func NewMcpSharedAccountPublicWithDefaults() *McpSharedAccountPublic`

NewMcpSharedAccountPublicWithDefaults instantiates a new McpSharedAccountPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenewable

`func (o *McpSharedAccountPublic) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *McpSharedAccountPublic) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *McpSharedAccountPublic) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### GetSignedInAt

`func (o *McpSharedAccountPublic) GetSignedInAt() string`

GetSignedInAt returns the SignedInAt field if non-nil, zero value otherwise.

### GetSignedInAtOk

`func (o *McpSharedAccountPublic) GetSignedInAtOk() (*string, bool)`

GetSignedInAtOk returns a tuple with the SignedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedInAt

`func (o *McpSharedAccountPublic) SetSignedInAt(v string)`

SetSignedInAt sets SignedInAt field to given value.


### GetSignedInBy

`func (o *McpSharedAccountPublic) GetSignedInBy() string`

GetSignedInBy returns the SignedInBy field if non-nil, zero value otherwise.

### GetSignedInByOk

`func (o *McpSharedAccountPublic) GetSignedInByOk() (*string, bool)`

GetSignedInByOk returns a tuple with the SignedInBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedInBy

`func (o *McpSharedAccountPublic) SetSignedInBy(v string)`

SetSignedInBy sets SignedInBy field to given value.


### SetSignedInByNil

`func (o *McpSharedAccountPublic) SetSignedInByNil(b bool)`

 SetSignedInByNil sets the value for SignedInBy to be an explicit nil

### UnsetSignedInBy
`func (o *McpSharedAccountPublic) UnsetSignedInBy()`

UnsetSignedInBy ensures that no value is present for SignedInBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


