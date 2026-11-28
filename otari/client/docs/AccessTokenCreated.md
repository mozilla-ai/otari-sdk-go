# AccessTokenCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** | When it ends. For an MCP client&#39;s sign-in, when renewing its access token stops working. | 
**Id** | **string** |  | 
**LastUsedAt** | **NullableTime** | Accurate to about a minute. Null until first used. | 
**Name** | **string** |  | 
**OrganizationId** | **NullableString** | The organization it reads, whichever one the dashboard has switched to since. | 
**Token** | **string** | The token itself. Shown once: only its hash is stored. | 
**TokenSuffix** | **string** | The token&#39;s last characters, to tell it apart from the account&#39;s others. | 

## Methods

### NewAccessTokenCreated

`func NewAccessTokenCreated(createdAt time.Time, expiresAt NullableTime, id string, lastUsedAt NullableTime, name string, organizationId NullableString, token string, tokenSuffix string, ) *AccessTokenCreated`

NewAccessTokenCreated instantiates a new AccessTokenCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCreatedWithDefaults

`func NewAccessTokenCreatedWithDefaults() *AccessTokenCreated`

NewAccessTokenCreatedWithDefaults instantiates a new AccessTokenCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AccessTokenCreated) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessTokenCreated) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessTokenCreated) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *AccessTokenCreated) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccessTokenCreated) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccessTokenCreated) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *AccessTokenCreated) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AccessTokenCreated) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetId

`func (o *AccessTokenCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenCreated) SetId(v string)`

SetId sets Id field to given value.


### GetLastUsedAt

`func (o *AccessTokenCreated) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AccessTokenCreated) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AccessTokenCreated) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *AccessTokenCreated) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *AccessTokenCreated) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetName

`func (o *AccessTokenCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenCreated) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *AccessTokenCreated) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AccessTokenCreated) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AccessTokenCreated) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *AccessTokenCreated) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AccessTokenCreated) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetToken

`func (o *AccessTokenCreated) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessTokenCreated) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessTokenCreated) SetToken(v string)`

SetToken sets Token field to given value.


### GetTokenSuffix

`func (o *AccessTokenCreated) GetTokenSuffix() string`

GetTokenSuffix returns the TokenSuffix field if non-nil, zero value otherwise.

### GetTokenSuffixOk

`func (o *AccessTokenCreated) GetTokenSuffixOk() (*string, bool)`

GetTokenSuffixOk returns a tuple with the TokenSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSuffix

`func (o *AccessTokenCreated) SetTokenSuffix(v string)`

SetTokenSuffix sets TokenSuffix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


