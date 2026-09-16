# Content14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**ToolReferences** | [**[]MRToolReferenceBlock**](MRToolReferenceBlock.md) |  | 

## Methods

### NewContent14

`func NewContent14(errorCode string, type_ string, toolReferences []MRToolReferenceBlock, ) *Content14`

NewContent14 instantiates a new Content14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent14WithDefaults

`func NewContent14WithDefaults() *Content14`

NewContent14WithDefaults instantiates a new Content14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content14) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content14) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content14) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *Content14) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Content14) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Content14) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Content14) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetType

`func (o *Content14) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content14) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content14) SetType(v string)`

SetType sets Type field to given value.


### GetToolReferences

`func (o *Content14) GetToolReferences() []MRToolReferenceBlock`

GetToolReferences returns the ToolReferences field if non-nil, zero value otherwise.

### GetToolReferencesOk

`func (o *Content14) GetToolReferencesOk() (*[]MRToolReferenceBlock, bool)`

GetToolReferencesOk returns a tuple with the ToolReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolReferences

`func (o *Content14) SetToolReferences(v []MRToolReferenceBlock)`

SetToolReferences sets ToolReferences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


