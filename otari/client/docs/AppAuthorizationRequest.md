# AppAuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**CodeChallenge** | **string** |  | 
**CodeChallengeMethod** | **string** |  | 
**RedirectUri** | **string** |  | 
**ResponseType** | Pointer to **string** |  | [optional] [default to "code"]
**State** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewAppAuthorizationRequest

`func NewAppAuthorizationRequest(clientId string, codeChallenge string, codeChallengeMethod string, redirectUri string, state string, workspaceId string, ) *AppAuthorizationRequest`

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


