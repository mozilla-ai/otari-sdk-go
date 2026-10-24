# OAuthErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | invalid_request, unsupported_grant_type, invalid_grant or temporarily_unavailable | 
**ErrorDescription** | **string** |  | 

## Methods

### NewOAuthErrorResponse

`func NewOAuthErrorResponse(error_ string, errorDescription string, ) *OAuthErrorResponse`

NewOAuthErrorResponse instantiates a new OAuthErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthErrorResponseWithDefaults

`func NewOAuthErrorResponseWithDefaults() *OAuthErrorResponse`

NewOAuthErrorResponseWithDefaults instantiates a new OAuthErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OAuthErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuthErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuthErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *OAuthErrorResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuthErrorResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuthErrorResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


