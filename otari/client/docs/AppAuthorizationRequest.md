# AppAuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**CodeChallenge** | **string** |  | 
**CodeChallengeMethod** | **string** |  | 
**Expiration** | Pointer to **string** | How long the key lives once it is issued. | [optional] [default to "never"]
**Limit** | Pointer to [**NullableLimit**](Limit.md) |  | [optional] 
**RedirectUri** | **string** |  | 
**ResponseType** | Pointer to **string** |  | [optional] [default to "code"]
**Scopes** | [**[]KeyScope**](KeyScope.md) | The scopes the consent screen showed. The key holds these and what the registration still names, so one added since the screen was drawn is left out and one removed since takes effect. | 
**State** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewAppAuthorizationRequest

`func NewAppAuthorizationRequest(clientId string, codeChallenge string, codeChallengeMethod string, redirectUri string, scopes []KeyScope, state string, workspaceId string, ) *AppAuthorizationRequest`

NewAppAuthorizationRequest instantiates a new AppAuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAuthorizationRequestWithDefaults

`func NewAppAuthorizationRequestWithDefaults() *AppAuthorizationRequest`

NewAppAuthorizationRequestWithDefaults instantiates a new AppAuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AppAuthorizationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AppAuthorizationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AppAuthorizationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCodeChallenge

`func (o *AppAuthorizationRequest) GetCodeChallenge() string`

GetCodeChallenge returns the CodeChallenge field if non-nil, zero value otherwise.

### GetCodeChallengeOk

`func (o *AppAuthorizationRequest) GetCodeChallengeOk() (*string, bool)`

GetCodeChallengeOk returns a tuple with the CodeChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallenge

`func (o *AppAuthorizationRequest) SetCodeChallenge(v string)`

SetCodeChallenge sets CodeChallenge field to given value.


### GetCodeChallengeMethod

`func (o *AppAuthorizationRequest) GetCodeChallengeMethod() string`

GetCodeChallengeMethod returns the CodeChallengeMethod field if non-nil, zero value otherwise.

### GetCodeChallengeMethodOk

`func (o *AppAuthorizationRequest) GetCodeChallengeMethodOk() (*string, bool)`

GetCodeChallengeMethodOk returns a tuple with the CodeChallengeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethod

`func (o *AppAuthorizationRequest) SetCodeChallengeMethod(v string)`

SetCodeChallengeMethod sets CodeChallengeMethod field to given value.


### GetExpiration

`func (o *AppAuthorizationRequest) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AppAuthorizationRequest) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AppAuthorizationRequest) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AppAuthorizationRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetLimit

`func (o *AppAuthorizationRequest) GetLimit() Limit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AppAuthorizationRequest) GetLimitOk() (*Limit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AppAuthorizationRequest) SetLimit(v Limit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AppAuthorizationRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *AppAuthorizationRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AppAuthorizationRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetRedirectUri

`func (o *AppAuthorizationRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AppAuthorizationRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AppAuthorizationRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetResponseType

`func (o *AppAuthorizationRequest) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *AppAuthorizationRequest) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *AppAuthorizationRequest) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *AppAuthorizationRequest) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetScopes

`func (o *AppAuthorizationRequest) GetScopes() []KeyScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppAuthorizationRequest) GetScopesOk() (*[]KeyScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppAuthorizationRequest) SetScopes(v []KeyScope)`

SetScopes sets Scopes field to given value.


### GetState

`func (o *AppAuthorizationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppAuthorizationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppAuthorizationRequest) SetState(v string)`

SetState sets State field to given value.


### GetWorkspaceId

`func (o *AppAuthorizationRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AppAuthorizationRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AppAuthorizationRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


