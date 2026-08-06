# MRBetaCodeExecutionResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]MRBetaCodeExecutionOutputBlock**](MRBetaCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaCodeExecutionResultBlock

`func NewMRBetaCodeExecutionResultBlock(content []MRBetaCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, type_ string, ) *MRBetaCodeExecutionResultBlock`

NewMRBetaCodeExecutionResultBlock instantiates a new MRBetaCodeExecutionResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCodeExecutionResultBlockWithDefaults

`func NewMRBetaCodeExecutionResultBlockWithDefaults() *MRBetaCodeExecutionResultBlock`

NewMRBetaCodeExecutionResultBlockWithDefaults instantiates a new MRBetaCodeExecutionResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaCodeExecutionResultBlock) GetContent() []MRBetaCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaCodeExecutionResultBlock) GetContentOk() (*[]MRBetaCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaCodeExecutionResultBlock) SetContent(v []MRBetaCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *MRBetaCodeExecutionResultBlock) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *MRBetaCodeExecutionResultBlock) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *MRBetaCodeExecutionResultBlock) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *MRBetaCodeExecutionResultBlock) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *MRBetaCodeExecutionResultBlock) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *MRBetaCodeExecutionResultBlock) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *MRBetaCodeExecutionResultBlock) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *MRBetaCodeExecutionResultBlock) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *MRBetaCodeExecutionResultBlock) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetType

`func (o *MRBetaCodeExecutionResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCodeExecutionResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCodeExecutionResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


