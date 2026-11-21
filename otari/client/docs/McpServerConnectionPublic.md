# McpServerConnectionPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAt** | **time.Time** |  | 
**Email** | **NullableString** | The dashboard account&#39;s email, when the user is one&#39;s | 
**FullName** | **NullableString** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** | The request-plane user the connection belongs to | 

## Methods

### NewMcpServerConnectionPublic

`func NewMcpServerConnectionPublic(connectedAt time.Time, email NullableString, fullName NullableString, updatedAt time.Time, userId string, ) *McpServerConnectionPublic`

NewMcpServerConnectionPublic instantiates a new McpServerConnectionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerConnectionPublicWithDefaults

`func NewMcpServerConnectionPublicWithDefaults() *McpServerConnectionPublic`

NewMcpServerConnectionPublicWithDefaults instantiates a new McpServerConnectionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedAt

`func (o *McpServerConnectionPublic) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *McpServerConnectionPublic) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *McpServerConnectionPublic) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.


### GetEmail

`func (o *McpServerConnectionPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *McpServerConnectionPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *McpServerConnectionPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *McpServerConnectionPublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *McpServerConnectionPublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *McpServerConnectionPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *McpServerConnectionPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *McpServerConnectionPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.


### SetFullNameNil

`func (o *McpServerConnectionPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *McpServerConnectionPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetUpdatedAt

`func (o *McpServerConnectionPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *McpServerConnectionPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *McpServerConnectionPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *McpServerConnectionPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *McpServerConnectionPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *McpServerConnectionPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


