# Content12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**Type** | **string** |  | 
**Content** | [**[]MRCodeExecutionOutputBlock**](MRCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 
**EncryptedStdout** | **string** |  | 

## Methods

### NewContent12

`func NewContent12(errorCode string, type_ string, content []MRCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, encryptedStdout string, ) *Content12`

NewContent12 instantiates a new Content12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent12WithDefaults

`func NewContent12WithDefaults() *Content12`

NewContent12WithDefaults instantiates a new Content12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content12) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content12) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content12) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetType

`func (o *Content12) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content12) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content12) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *Content12) GetContent() []MRCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content12) GetContentOk() (*[]MRCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content12) SetContent(v []MRCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *Content12) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *Content12) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *Content12) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *Content12) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Content12) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Content12) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *Content12) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Content12) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Content12) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetEncryptedStdout

`func (o *Content12) GetEncryptedStdout() string`

GetEncryptedStdout returns the EncryptedStdout field if non-nil, zero value otherwise.

### GetEncryptedStdoutOk

`func (o *Content12) GetEncryptedStdoutOk() (*string, bool)`

GetEncryptedStdoutOk returns a tuple with the EncryptedStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedStdout

`func (o *Content12) SetEncryptedStdout(v string)`

SetEncryptedStdout sets EncryptedStdout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


