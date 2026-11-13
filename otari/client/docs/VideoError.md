# VideoError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewVideoError

`func NewVideoError(message string, ) *VideoError`

NewVideoError instantiates a new VideoError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoErrorWithDefaults

`func NewVideoErrorWithDefaults() *VideoError`

NewVideoErrorWithDefaults instantiates a new VideoError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VideoError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VideoError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VideoError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VideoError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VideoError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VideoError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *VideoError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VideoError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VideoError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


