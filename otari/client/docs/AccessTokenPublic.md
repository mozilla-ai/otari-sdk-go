# AccessTokenPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** | When it ends. For an MCP client&#39;s sign-in, when renewing its access token stops working. | 
**Id** | **string** |  | 
**LastUsedAt** | **NullableTime** | Accurate to about a minute. Null until first used. | 
**Name** | **string** |  | 
**OrganizationId** | **NullableString** | The organization it reads, whichever one the dashboard has switched to since. | 
**TokenSuffix** | **string** | The token&#39;s last characters, to tell it apart from the account&#39;s others. | 

## Methods

### NewAccessTokenPublic

`func NewAccessTokenPublic(createdAt time.Time, expiresAt NullableTime, id string, lastUsedAt NullableTime, name string, organizationId NullableString, tokenSuffix string, ) *AccessTokenPublic`

NewAccessTokenPublic instantiates a new AccessTokenPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenPublicWithDefaults

`func NewAccessTokenPublicWithDefaults() *AccessTokenPublic`

NewAccessTokenPublicWithDefaults instantiates a new AccessTokenPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AccessTokenPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessTokenPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessTokenPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *AccessTokenPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccessTokenPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccessTokenPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *AccessTokenPublic) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AccessTokenPublic) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *AccessTokenPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenPublic) SetId(v string)`

SetId sets Id field to given value.


### GetLastUsedAt

`func (o *AccessTokenPublic) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AccessTokenPublic) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AccessTokenPublic) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *AccessTokenPublic) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *AccessTokenPublic) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetName

`func (o *AccessTokenPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *AccessTokenPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccessTokenPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccessTokenPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *AccessTokenPublic) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AccessTokenPublic) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetTokenSuffix

`func (o *AccessTokenPublic) GetTokenSuffix() string`

GetTokenSuffix returns the TokenSuffix field if non-nil, zero value otherwise.

### GetTokenSuffixOk

`func (o *AccessTokenPublic) GetTokenSuffixOk() (*string, bool)`

GetTokenSuffixOk returns a tuple with the TokenSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSuffix

`func (o *AccessTokenPublic) SetTokenSuffix(v string)`

SetTokenSuffix sets TokenSuffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


