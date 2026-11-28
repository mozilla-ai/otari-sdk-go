# McpErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Detail** | **string** |  | 
**ExecutionState** | [**ExecutionState**](ExecutionState.md) |  | 
**RequestId** | **string** |  | 
**SignInUrl** | Pointer to **NullableString** | With mcp_sign_in_required: a link that starts this key&#39;s user&#39;s own sign-in to the server. It works for 30 minutes and stores nothing until it is opened | [optional] 

## Methods

### NewMcpErrorBody

`func NewMcpErrorBody(code string, detail string, executionState ExecutionState, requestId string, ) *McpErrorBody`

NewMcpErrorBody instantiates a new McpErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpErrorBodyWithDefaults

`func NewMcpErrorBodyWithDefaults() *McpErrorBody`

NewMcpErrorBodyWithDefaults instantiates a new McpErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *McpErrorBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *McpErrorBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *McpErrorBody) SetCode(v string)`

SetCode sets Code field to given value.


### GetDetail

`func (o *McpErrorBody) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *McpErrorBody) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *McpErrorBody) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetExecutionState

`func (o *McpErrorBody) GetExecutionState() ExecutionState`

GetExecutionState returns the ExecutionState field if non-nil, zero value otherwise.

### GetExecutionStateOk

`func (o *McpErrorBody) GetExecutionStateOk() (*ExecutionState, bool)`

GetExecutionStateOk returns a tuple with the ExecutionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionState

`func (o *McpErrorBody) SetExecutionState(v ExecutionState)`

SetExecutionState sets ExecutionState field to given value.


### GetRequestId

`func (o *McpErrorBody) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *McpErrorBody) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *McpErrorBody) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSignInUrl

`func (o *McpErrorBody) GetSignInUrl() string`

GetSignInUrl returns the SignInUrl field if non-nil, zero value otherwise.

### GetSignInUrlOk

`func (o *McpErrorBody) GetSignInUrlOk() (*string, bool)`

GetSignInUrlOk returns a tuple with the SignInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInUrl

`func (o *McpErrorBody) SetSignInUrl(v string)`

SetSignInUrl sets SignInUrl field to given value.

### HasSignInUrl

`func (o *McpErrorBody) HasSignInUrl() bool`

HasSignInUrl returns a boolean if a field has been set.

### SetSignInUrlNil

`func (o *McpErrorBody) SetSignInUrlNil(b bool)`

 SetSignInUrlNil sets the value for SignInUrl to be an explicit nil

### UnsetSignInUrl
`func (o *McpErrorBody) UnsetSignInUrl()`

UnsetSignInUrl ensures that no value is present for SignInUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


