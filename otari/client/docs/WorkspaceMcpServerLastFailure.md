# WorkspaceMcpServerLastFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**CheckFailure**](CheckFailure.md) |  | 
**FailedAt** | **string** |  | 
**HttpStatus** | Pointer to **NullableInt32** | The HTTP status the server refused with, if any | [optional] 

## Methods

### NewWorkspaceMcpServerLastFailure

`func NewWorkspaceMcpServerLastFailure(error_ CheckFailure, failedAt string, ) *WorkspaceMcpServerLastFailure`

NewWorkspaceMcpServerLastFailure instantiates a new WorkspaceMcpServerLastFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerLastFailureWithDefaults

`func NewWorkspaceMcpServerLastFailureWithDefaults() *WorkspaceMcpServerLastFailure`

NewWorkspaceMcpServerLastFailureWithDefaults instantiates a new WorkspaceMcpServerLastFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WorkspaceMcpServerLastFailure) GetError() CheckFailure`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkspaceMcpServerLastFailure) GetErrorOk() (*CheckFailure, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkspaceMcpServerLastFailure) SetError(v CheckFailure)`

SetError sets Error field to given value.


### GetFailedAt

`func (o *WorkspaceMcpServerLastFailure) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *WorkspaceMcpServerLastFailure) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *WorkspaceMcpServerLastFailure) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.


### GetHttpStatus

`func (o *WorkspaceMcpServerLastFailure) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *WorkspaceMcpServerLastFailure) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *WorkspaceMcpServerLastFailure) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *WorkspaceMcpServerLastFailure) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### SetHttpStatusNil

`func (o *WorkspaceMcpServerLastFailure) SetHttpStatusNil(b bool)`

 SetHttpStatusNil sets the value for HttpStatus to be an explicit nil

### UnsetHttpStatus
`func (o *WorkspaceMcpServerLastFailure) UnsetHttpStatus()`

UnsetHttpStatus ensures that no value is present for HttpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


