# McpConnectionPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAt** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** | When the current access token expires; null when the server stated no lifetime | 
**Renewable** | **bool** | Whether Otari can renew the token on its own; false means signing in again once it expires | 
**ServerId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewMcpConnectionPublic

`func NewMcpConnectionPublic(connectedAt time.Time, expiresAt NullableTime, renewable bool, serverId string, updatedAt time.Time, ) *McpConnectionPublic`

NewMcpConnectionPublic instantiates a new McpConnectionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionPublicWithDefaults

`func NewMcpConnectionPublicWithDefaults() *McpConnectionPublic`

NewMcpConnectionPublicWithDefaults instantiates a new McpConnectionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedAt

`func (o *McpConnectionPublic) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *McpConnectionPublic) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *McpConnectionPublic) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.


### GetExpiresAt

`func (o *McpConnectionPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *McpConnectionPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *McpConnectionPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *McpConnectionPublic) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *McpConnectionPublic) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetRenewable

`func (o *McpConnectionPublic) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *McpConnectionPublic) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *McpConnectionPublic) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.


### GetServerId

`func (o *McpConnectionPublic) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpConnectionPublic) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpConnectionPublic) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetUpdatedAt

`func (o *McpConnectionPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *McpConnectionPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *McpConnectionPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


