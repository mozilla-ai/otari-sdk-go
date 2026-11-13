# CodeExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]CodeExecutionOutputFile**](CodeExecutionOutputFile.md) | The files the code wrote. Empty when it wrote none, or where the Files API is off. | 
**Object** | Pointer to **string** |  | [optional] [default to "code_execution.result"]
**ReturnCode** | **NullableInt32** | The program&#39;s exit status; null when the sandbox did not report one. | 
**Stderr** | **string** |  | 
**Stdout** | **string** |  | 

## Methods

### NewCodeExecuteResponse

`func NewCodeExecuteResponse(files []CodeExecutionOutputFile, returnCode NullableInt32, stderr string, stdout string, ) *CodeExecuteResponse`

NewCodeExecuteResponse instantiates a new CodeExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeExecuteResponseWithDefaults

`func NewCodeExecuteResponseWithDefaults() *CodeExecuteResponse`

NewCodeExecuteResponseWithDefaults instantiates a new CodeExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CodeExecuteResponse) GetFiles() []CodeExecutionOutputFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CodeExecuteResponse) GetFilesOk() (*[]CodeExecutionOutputFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CodeExecuteResponse) SetFiles(v []CodeExecutionOutputFile)`

SetFiles sets Files field to given value.


### GetObject

`func (o *CodeExecuteResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CodeExecuteResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CodeExecuteResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CodeExecuteResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetReturnCode

`func (o *CodeExecuteResponse) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CodeExecuteResponse) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CodeExecuteResponse) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.


### SetReturnCodeNil

`func (o *CodeExecuteResponse) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *CodeExecuteResponse) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetStderr

`func (o *CodeExecuteResponse) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CodeExecuteResponse) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CodeExecuteResponse) SetStderr(v string)`

SetStderr sets Stderr field to given value.


### GetStdout

`func (o *CodeExecuteResponse) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CodeExecuteResponse) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CodeExecuteResponse) SetStdout(v string)`

SetStdout sets Stdout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


