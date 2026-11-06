# CodeExecutionOutputFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **string** | Download it from GET /api/v1/files/{file_id}/content | 
**Filename** | **string** | The name the code gave it in the working directory | 

## Methods

### NewCodeExecutionOutputFile

`func NewCodeExecutionOutputFile(fileId string, filename string, ) *CodeExecutionOutputFile`

NewCodeExecutionOutputFile instantiates a new CodeExecutionOutputFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeExecutionOutputFileWithDefaults

`func NewCodeExecutionOutputFileWithDefaults() *CodeExecutionOutputFile`

NewCodeExecutionOutputFileWithDefaults instantiates a new CodeExecutionOutputFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *CodeExecutionOutputFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *CodeExecutionOutputFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *CodeExecutionOutputFile) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetFilename

`func (o *CodeExecutionOutputFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CodeExecutionOutputFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CodeExecutionOutputFile) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


