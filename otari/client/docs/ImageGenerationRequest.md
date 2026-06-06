# ImageGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**N** | Pointer to **NullableInt32** |  | [optional] 
**Prompt** | **string** |  | 
**Quality** | Pointer to **NullableString** |  | [optional] 
**ResponseFormat** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImageGenerationRequest

`func NewImageGenerationRequest(model string, prompt string, ) *ImageGenerationRequest`

NewImageGenerationRequest instantiates a new ImageGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestWithDefaults

`func NewImageGenerationRequestWithDefaults() *ImageGenerationRequest`

NewImageGenerationRequestWithDefaults instantiates a new ImageGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ImageGenerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ImageGenerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ImageGenerationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetN

`func (o *ImageGenerationRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ImageGenerationRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ImageGenerationRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ImageGenerationRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *ImageGenerationRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *ImageGenerationRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetPrompt

`func (o *ImageGenerationRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ImageGenerationRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ImageGenerationRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetQuality

`func (o *ImageGenerationRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ImageGenerationRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ImageGenerationRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ImageGenerationRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *ImageGenerationRequest) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *ImageGenerationRequest) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetResponseFormat

`func (o *ImageGenerationRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ImageGenerationRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ImageGenerationRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ImageGenerationRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### SetResponseFormatNil

`func (o *ImageGenerationRequest) SetResponseFormatNil(b bool)`

 SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil

### UnsetResponseFormat
`func (o *ImageGenerationRequest) UnsetResponseFormat()`

UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
### GetSize

`func (o *ImageGenerationRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageGenerationRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageGenerationRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageGenerationRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ImageGenerationRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ImageGenerationRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStyle

`func (o *ImageGenerationRequest) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ImageGenerationRequest) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ImageGenerationRequest) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ImageGenerationRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *ImageGenerationRequest) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *ImageGenerationRequest) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetUser

`func (o *ImageGenerationRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ImageGenerationRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ImageGenerationRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ImageGenerationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ImageGenerationRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ImageGenerationRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


