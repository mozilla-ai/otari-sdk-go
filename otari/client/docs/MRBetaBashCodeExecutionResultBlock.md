# MRBetaBashCodeExecutionResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]MRBetaBashCodeExecutionOutputBlock**](MRBetaBashCodeExecutionOutputBlock.md) |  | 
**ReturnCode** | **int32** |  | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaBashCodeExecutionResultBlock

`func NewMRBetaBashCodeExecutionResultBlock(content []MRBetaBashCodeExecutionOutputBlock, returnCode int32, stderr string, stdout string, type_ string, ) *MRBetaBashCodeExecutionResultBlock`

NewMRBetaBashCodeExecutionResultBlock instantiates a new MRBetaBashCodeExecutionResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaBashCodeExecutionResultBlockWithDefaults

`func NewMRBetaBashCodeExecutionResultBlockWithDefaults() *MRBetaBashCodeExecutionResultBlock`

NewMRBetaBashCodeExecutionResultBlockWithDefaults instantiates a new MRBetaBashCodeExecutionResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaBashCodeExecutionResultBlock) GetContent() []MRBetaBashCodeExecutionOutputBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaBashCodeExecutionResultBlock) GetContentOk() (*[]MRBetaBashCodeExecutionOutputBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaBashCodeExecutionResultBlock) SetContent(v []MRBetaBashCodeExecutionOutputBlock)`

SetContent sets Content field to given value.


### GetReturnCode

`func (o *MRBetaBashCodeExecutionResultBlock) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *MRBetaBashCodeExecutionResultBlock) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *MRBetaBashCodeExecutionResultBlock) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### GetStderr

`func (o *MRBetaBashCodeExecutionResultBlock) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *MRBetaBashCodeExecutionResultBlock) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *MRBetaBashCodeExecutionResultBlock) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *MRBetaBashCodeExecutionResultBlock) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *MRBetaBashCodeExecutionResultBlock) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *MRBetaBashCodeExecutionResultBlock) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetType

`func (o *MRBetaBashCodeExecutionResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaBashCodeExecutionResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaBashCodeExecutionResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


