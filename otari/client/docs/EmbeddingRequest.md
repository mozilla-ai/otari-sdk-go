# EmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to **NullableInt32** |  | [optional] 
**EncodingFormat** | Pointer to **NullableString** |  | [optional] 
**Input** | [**Input**](Input.md) |  | 
**Model** | **string** |  | 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmbeddingRequest

`func NewEmbeddingRequest(input Input, model string, ) *EmbeddingRequest`

NewEmbeddingRequest instantiates a new EmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingRequestWithDefaults

`func NewEmbeddingRequestWithDefaults() *EmbeddingRequest`

NewEmbeddingRequestWithDefaults instantiates a new EmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *EmbeddingRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *EmbeddingRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *EmbeddingRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *EmbeddingRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### SetDimensionsNil

`func (o *EmbeddingRequest) SetDimensionsNil(b bool)`

 SetDimensionsNil sets the value for Dimensions to be an explicit nil

### UnsetDimensions
`func (o *EmbeddingRequest) UnsetDimensions()`

UnsetDimensions ensures that no value is present for Dimensions, not even an explicit nil
### GetEncodingFormat

`func (o *EmbeddingRequest) GetEncodingFormat() string`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *EmbeddingRequest) GetEncodingFormatOk() (*string, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *EmbeddingRequest) SetEncodingFormat(v string)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *EmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.

### SetEncodingFormatNil

`func (o *EmbeddingRequest) SetEncodingFormatNil(b bool)`

 SetEncodingFormatNil sets the value for EncodingFormat to be an explicit nil

### UnsetEncodingFormat
`func (o *EmbeddingRequest) UnsetEncodingFormat()`

UnsetEncodingFormat ensures that no value is present for EncodingFormat, not even an explicit nil
### GetInput

`func (o *EmbeddingRequest) GetInput() Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EmbeddingRequest) GetInputOk() (*Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EmbeddingRequest) SetInput(v Input)`

SetInput sets Input field to given value.


### GetModel

`func (o *EmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetUser

`func (o *EmbeddingRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EmbeddingRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EmbeddingRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *EmbeddingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *EmbeddingRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *EmbeddingRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


