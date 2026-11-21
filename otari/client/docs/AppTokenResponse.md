# AppTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**ManageUrl** | **string** | The absolute URL of the page where the person sees this key, and revokes it. | 
**Scope** | **string** | The scopes the key holds, space delimited. Fewer than the registration names where the person approved it before an operator narrowed it. | 
**TokenType** | Pointer to **string** |  | [optional] [default to "Bearer"]

## Methods

### NewAppTokenResponse

`func NewAppTokenResponse(accessToken string, manageUrl string, scope string, ) *AppTokenResponse`

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


