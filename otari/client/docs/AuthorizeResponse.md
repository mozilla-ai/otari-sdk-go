# AuthorizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUrl** | **string** | The provider consent screen to navigate to. | 
**State** | **string** | An opaque CSRF value to keep for the length of the redirect, compare against the &#39;state&#39; the provider returns, and send back with the authorization code. A callback whose state does not match the one held by the browser that started the flow should be abandoned by the client rather than sent here; one that does is checked again against this deployment&#39;s own record of it. The response also sets an HttpOnly cookie that the callback requires, so the exchange can only be completed from the browser this call was made from. | 

## Methods

### NewAuthorizeResponse

`func NewAuthorizeResponse(authorizationUrl string, state string, ) *AuthorizeResponse`

NewAuthorizeResponse instantiates a new AuthorizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeResponseWithDefaults

`func NewAuthorizeResponseWithDefaults() *AuthorizeResponse`

NewAuthorizeResponseWithDefaults instantiates a new AuthorizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUrl

`func (o *AuthorizeResponse) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *AuthorizeResponse) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *AuthorizeResponse) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.


### GetState

`func (o *AuthorizeResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AuthorizeResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AuthorizeResponse) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


