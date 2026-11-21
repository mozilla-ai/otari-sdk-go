# SsoAuthorizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUrl** | **string** | The identity provider&#39;s consent screen to navigate to. | 
**State** | **string** | An opaque CSRF value to keep for the length of the redirect and send back with the authorization code. The response also sets an HttpOnly cookie that the callback requires, so the exchange can only be completed from the browser this call was made from. | 

## Methods

### NewSsoAuthorizeResponse

`func NewSsoAuthorizeResponse(authorizationUrl string, state string, ) *SsoAuthorizeResponse`

NewSsoAuthorizeResponse instantiates a new SsoAuthorizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoAuthorizeResponseWithDefaults

`func NewSsoAuthorizeResponseWithDefaults() *SsoAuthorizeResponse`

NewSsoAuthorizeResponseWithDefaults instantiates a new SsoAuthorizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUrl

`func (o *SsoAuthorizeResponse) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *SsoAuthorizeResponse) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *SsoAuthorizeResponse) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.


### GetState

`func (o *SsoAuthorizeResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SsoAuthorizeResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SsoAuthorizeResponse) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


