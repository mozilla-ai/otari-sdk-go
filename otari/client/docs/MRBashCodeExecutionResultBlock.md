# MRBashCodeExecutionResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]MRBashCodeExecutionOutputBlock**](MRBashCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBashCodeExecutionResultBlock

`func NewMRBashCodeExecutionResultBlock(content []MRBashCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, type_ string, ) *MRBashCodeExecutionResultBlock`

NewMRBashCodeExecutionResultBlock instantiates a new MRBashCodeExecutionResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBashCodeExecutionResultBlockWithDefaults

`func NewMRBashCodeExecutionResultBlockWithDefaults() *MRBashCodeExecutionResultBlock`

NewMRBashCodeExecutionResultBlockWithDefaults instantiates a new MRBashCodeExecutionResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBashCodeExecutionResultBlock) GetContent() []MRBashCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBashCodeExecutionResultBlock) GetContentOk() (*[]MRBashCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBashCodeExecutionResultBlock) SetContent(v []MRBashCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *MRBashCodeExecutionResultBlock) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *MRBashCodeExecutionResultBlock) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *MRBashCodeExecutionResultBlock) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *MRBashCodeExecutionResultBlock) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *MRBashCodeExecutionResultBlock) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *MRBashCodeExecutionResultBlock) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *MRBashCodeExecutionResultBlock) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *MRBashCodeExecutionResultBlock) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *MRBashCodeExecutionResultBlock) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetType

`func (o *MRBashCodeExecutionResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBashCodeExecutionResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBashCodeExecutionResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


