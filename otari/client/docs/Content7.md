# Content7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Content** | **string** |  | 
**FileType** | **string** |  | 
**NumLines** | Pointer to **int32** |  | [optional] 
**StartLine** | Pointer to **int32** |  | [optional] 
**TotalLines** | Pointer to **int32** |  | [optional] 
**IsFileUpdate** | **bool** |  | 
**Lines** | Pointer to **[]string** |  | [optional] 
**NewLines** | Pointer to **int32** |  | [optional] 
**NewStart** | Pointer to **int32** |  | [optional] 
**OldLines** | Pointer to **int32** |  | [optional] 
**OldStart** | Pointer to **int32** |  | [optional] 

## Methods

### NewContent7

`func NewContent7(errorCode string, type_ string, content string, fileType string, isFileUpdate bool, ) *Content7`

NewContent7 instantiates a new Content7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent7WithDefaults

`func NewContent7WithDefaults() *Content7`

NewContent7WithDefaults instantiates a new Content7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content7) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content7) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content7) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *Content7) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Content7) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Content7) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Content7) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetType

`func (o *Content7) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content7) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content7) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *Content7) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content7) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content7) SetContent(v string)`

SetContent sets Content field to given value.


### GetFileType

`func (o *Content7) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *Content7) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *Content7) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetNumLines

`func (o *Content7) GetNumLines() int32`

GetNumLines returns the NumLines field if non-nil, zero value otherwise.

### GetNumLinesOk

`func (o *Content7) GetNumLinesOk() (*int32, bool)`

GetNumLinesOk returns a tuple with the NumLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLines

`func (o *Content7) SetNumLines(v int32)`

SetNumLines sets NumLines field to given value.

### HasNumLines

`func (o *Content7) HasNumLines() bool`

HasNumLines returns a boolean if a field has been set.

### GetStartLine

`func (o *Content7) GetStartLine() int32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *Content7) GetStartLineOk() (*int32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *Content7) SetStartLine(v int32)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *Content7) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.

### GetTotalLines

`func (o *Content7) GetTotalLines() int32`

GetTotalLines returns the TotalLines field if non-nil, zero value otherwise.

### GetTotalLinesOk

`func (o *Content7) GetTotalLinesOk() (*int32, bool)`

GetTotalLinesOk returns a tuple with the TotalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLines

`func (o *Content7) SetTotalLines(v int32)`

SetTotalLines sets TotalLines field to given value.

### HasTotalLines

`func (o *Content7) HasTotalLines() bool`

HasTotalLines returns a boolean if a field has been set.

### GetIsFileUpdate

`func (o *Content7) GetIsFileUpdate() bool`

GetIsFileUpdate returns the IsFileUpdate field if non-nil, zero value otherwise.

### GetIsFileUpdateOk

`func (o *Content7) GetIsFileUpdateOk() (*bool, bool)`

GetIsFileUpdateOk returns a tuple with the IsFileUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFileUpdate

`func (o *Content7) SetIsFileUpdate(v bool)`

SetIsFileUpdate sets IsFileUpdate field to given value.


### GetLines

`func (o *Content7) GetLines() []string`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Content7) GetLinesOk() (*[]string, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Content7) SetLines(v []string)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Content7) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNewLines

`func (o *Content7) GetNewLines() int32`

GetNewLines returns the NewLines field if non-nil, zero value otherwise.

### GetNewLinesOk

`func (o *Content7) GetNewLinesOk() (*int32, bool)`

GetNewLinesOk returns a tuple with the NewLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLines

`func (o *Content7) SetNewLines(v int32)`

SetNewLines sets NewLines field to given value.

### HasNewLines

`func (o *Content7) HasNewLines() bool`

HasNewLines returns a boolean if a field has been set.

### GetNewStart

`func (o *Content7) GetNewStart() int32`

GetNewStart returns the NewStart field if non-nil, zero value otherwise.

### GetNewStartOk

`func (o *Content7) GetNewStartOk() (*int32, bool)`

GetNewStartOk returns a tuple with the NewStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStart

`func (o *Content7) SetNewStart(v int32)`

SetNewStart sets NewStart field to given value.

### HasNewStart

`func (o *Content7) HasNewStart() bool`

HasNewStart returns a boolean if a field has been set.

### GetOldLines

`func (o *Content7) GetOldLines() int32`

GetOldLines returns the OldLines field if non-nil, zero value otherwise.

### GetOldLinesOk

`func (o *Content7) GetOldLinesOk() (*int32, bool)`

GetOldLinesOk returns a tuple with the OldLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldLines

`func (o *Content7) SetOldLines(v int32)`

SetOldLines sets OldLines field to given value.

### HasOldLines

`func (o *Content7) HasOldLines() bool`

HasOldLines returns a boolean if a field has been set.

### GetOldStart

`func (o *Content7) GetOldStart() int32`

GetOldStart returns the OldStart field if non-nil, zero value otherwise.

### GetOldStartOk

`func (o *Content7) GetOldStartOk() (*int32, bool)`

GetOldStartOk returns a tuple with the OldStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldStart

`func (o *Content7) SetOldStart(v int32)`

SetOldStart sets OldStart field to given value.

### HasOldStart

`func (o *Content7) HasOldStart() bool`

HasOldStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


