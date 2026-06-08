# Content3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**Type** | **string** |  | 
**Content** | [**[]MRBashCodeExecutionOutputBlock**](MRBashCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 

## Methods

### NewContent3

`func NewContent3(errorCode string, type_ string, content []MRBashCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, ) *Content3`

NewContent3 instantiates a new Content3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent3WithDefaults

`func NewContent3WithDefaults() *Content3`

NewContent3WithDefaults instantiates a new Content3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content3) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content3) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content3) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetType

`func (o *Content3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content3) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *Content3) GetContent() []MRBashCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content3) GetContentOk() (*[]MRBashCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content3) SetContent(v []MRBashCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *Content3) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *Content3) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *Content3) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *Content3) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Content3) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Content3) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *Content3) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Content3) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Content3) SetStdout(v string)`

SetStdout sets Stdout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


