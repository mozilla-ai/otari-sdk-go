# AppTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**ExpiresIn** | Pointer to **NullableInt32** | Seconds until an access token expires, where it does. | [optional] 
**ManageUrl** | **string** | The absolute URL of the page where the person sees this key, and revokes it. | 
**RefreshToken** | Pointer to **NullableString** | Renews an MCP client&#39;s access token at this endpoint, and is replaced on each use. Using one a second time ends the sign-in. | [optional] 
**Scope** | Pointer to **NullableString** | The scopes the key holds, space delimited. Fewer than the registration names where the person approved it before an operator narrowed it. Absent for an access token. | [optional] 
**TokenType** | Pointer to **string** |  | [optional] [default to "Bearer"]

## Methods

### NewAppTokenResponse

`func NewAppTokenResponse(accessToken string, manageUrl string, ) *AppTokenResponse`

NewAppTokenResponse instantiates a new AppTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTokenResponseWithDefaults

`func NewAppTokenResponseWithDefaults() *AppTokenResponse`

NewAppTokenResponseWithDefaults instantiates a new AppTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AppTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AppTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AppTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *AppTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AppTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AppTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AppTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### SetExpiresInNil

`func (o *AppTokenResponse) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *AppTokenResponse) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
### GetManageUrl

`func (o *AppTokenResponse) GetManageUrl() string`

GetManageUrl returns the ManageUrl field if non-nil, zero value otherwise.

### GetManageUrlOk

`func (o *AppTokenResponse) GetManageUrlOk() (*string, bool)`

GetManageUrlOk returns a tuple with the ManageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageUrl

`func (o *AppTokenResponse) SetManageUrl(v string)`

SetManageUrl sets ManageUrl field to given value.


### GetRefreshToken

`func (o *AppTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AppTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AppTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AppTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *AppTokenResponse) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *AppTokenResponse) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetScope

`func (o *AppTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AppTokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *AppTokenResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *AppTokenResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetTokenType

`func (o *AppTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AppTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AppTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AppTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


