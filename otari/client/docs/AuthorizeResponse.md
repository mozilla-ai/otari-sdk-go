# AuthorizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUrl** | **string** | The provider consent screen to navigate to. | 
**State** | **string** | An opaque CSRF value to keep for the length of the redirect and compare against the &#39;state&#39; the provider returns. It is not stored on this deployment, so a callback whose state does not match the one held by the browser that started the flow must be abandoned by the client rather than sent here. | 

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


