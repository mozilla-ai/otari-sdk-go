# MRBetaEncryptedCodeExecutionResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]MRBetaCodeExecutionOutputBlock**](MRBetaCodeExecutionOutputBlock.md) |  | 
**EncryptedStdout** | **string** |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaEncryptedCodeExecutionResultBlock

`func NewMRBetaEncryptedCodeExecutionResultBlock(content []MRBetaCodeExecutionOutputBlock, encryptedStdout string, returnCode int32, stderr string, type_ string, ) *MRBetaEncryptedCodeExecutionResultBlock`

NewMRBetaEncryptedCodeExecutionResultBlock instantiates a new MRBetaEncryptedCodeExecutionResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaEncryptedCodeExecutionResultBlockWithDefaults

`func NewMRBetaEncryptedCodeExecutionResultBlockWithDefaults() *MRBetaEncryptedCodeExecutionResultBlock`

NewMRBetaEncryptedCodeExecutionResultBlockWithDefaults instantiates a new MRBetaEncryptedCodeExecutionResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetContent() []MRBetaCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetContentOk() (*[]MRBetaCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaEncryptedCodeExecutionResultBlock) SetContent(v []MRBetaCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetEncryptedStdout

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetEncryptedStdout() string`

GetEncryptedStdout returns the EncryptedStdout field if non-nil, zero value otherwise.

### GetEncryptedStdoutOk

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetEncryptedStdoutOk() (*string, bool)`

GetEncryptedStdoutOk returns a tuple with the EncryptedStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedStdout

`func (o *MRBetaEncryptedCodeExecutionResultBlock) SetEncryptedStdout(v string)`

SetEncryptedStdout sets EncryptedStdout field to given value.


### GetReturnCode

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *MRBetaEncryptedCodeExecutionResultBlock) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *MRBetaEncryptedCodeExecutionResultBlock) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetType

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaEncryptedCodeExecutionResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaEncryptedCodeExecutionResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


