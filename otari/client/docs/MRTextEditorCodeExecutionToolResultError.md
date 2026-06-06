# MRTextEditorCodeExecutionToolResultError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewMRTextEditorCodeExecutionToolResultError

`func NewMRTextEditorCodeExecutionToolResultError(errorCode string, type_ string, ) *MRTextEditorCodeExecutionToolResultError`

NewMRTextEditorCodeExecutionToolResultError instantiates a new MRTextEditorCodeExecutionToolResultError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRTextEditorCodeExecutionToolResultErrorWithDefaults

`func NewMRTextEditorCodeExecutionToolResultErrorWithDefaults() *MRTextEditorCodeExecutionToolResultError`

NewMRTextEditorCodeExecutionToolResultErrorWithDefaults instantiates a new MRTextEditorCodeExecutionToolResultError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *MRTextEditorCodeExecutionToolResultError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MRTextEditorCodeExecutionToolResultError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MRTextEditorCodeExecutionToolResultError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *MRTextEditorCodeExecutionToolResultError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MRTextEditorCodeExecutionToolResultError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MRTextEditorCodeExecutionToolResultError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MRTextEditorCodeExecutionToolResultError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *MRTextEditorCodeExecutionToolResultError) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *MRTextEditorCodeExecutionToolResultError) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetType

`func (o *MRTextEditorCodeExecutionToolResultError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRTextEditorCodeExecutionToolResultError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRTextEditorCodeExecutionToolResultError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


