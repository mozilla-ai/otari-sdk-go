# OAuthCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The authorization code from the provider&#39;s redirect. | 
**State** | **string** | The &#39;state&#39; from the provider&#39;s redirect, as issued by /authorize. | 

## Methods

### NewOAuthCallbackRequest

`func NewOAuthCallbackRequest(code string, state string, ) *OAuthCallbackRequest`

NewOAuthCallbackRequest instantiates a new OAuthCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCallbackRequestWithDefaults

`func NewOAuthCallbackRequestWithDefaults() *OAuthCallbackRequest`

NewOAuthCallbackRequestWithDefaults instantiates a new OAuthCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OAuthCallbackRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthCallbackRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthCallbackRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetState

`func (o *OAuthCallbackRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OAuthCallbackRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OAuthCallbackRequest) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


