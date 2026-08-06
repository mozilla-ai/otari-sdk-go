# Content5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**Type** | **string** |  | 
**Content** | [**[]MRBetaCodeExecutionOutputBlock**](MRBetaCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 
**EncryptedStdout** | **string** |  | 

## Methods

### NewContent5

`func NewContent5(errorCode string, type_ string, content []MRBetaCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, encryptedStdout string, ) *Content5`

NewContent5 instantiates a new Content5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent5WithDefaults

`func NewContent5WithDefaults() *Content5`

NewContent5WithDefaults instantiates a new Content5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content5) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content5) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content5) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetType

`func (o *Content5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content5) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *Content5) GetContent() []MRBetaCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content5) GetContentOk() (*[]MRBetaCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content5) SetContent(v []MRBetaCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *Content5) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *Content5) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *Content5) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *Content5) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *Content5) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *Content5) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *Content5) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *Content5) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *Content5) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetEncryptedStdout

`func (o *Content5) GetEncryptedStdout() string`

GetEncryptedStdout returns the EncryptedStdout field if non-nil, zero value otherwise.

### GetEncryptedStdoutOk

`func (o *Content5) GetEncryptedStdoutOk() (*string, bool)`

GetEncryptedStdoutOk returns a tuple with the EncryptedStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedStdout

`func (o *Content5) SetEncryptedStdout(v string)`

SetEncryptedStdout sets EncryptedStdout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


