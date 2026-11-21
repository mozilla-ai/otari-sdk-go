# ImagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **int32** |  | 
**Background** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to [**[]IMGImage**](IMGImage.md) |  | [optional] 
**OutputFormat** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to **NullableString** |  | [optional] 
**SafetyIdentifier** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to [**NullableSize**](Size.md) |  | [optional] 
**SvgUsage** | Pointer to [**NullableIMGSvgUsage**](IMGSvgUsage.md) |  | [optional] 
**Usage** | Pointer to [**NullableIMGUsage**](IMGUsage.md) |  | [optional] 

## Methods

### NewImagesResponse

`func NewImagesResponse(created int32, ) *ImagesResponse`

NewImagesResponse instantiates a new ImagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesResponseWithDefaults

`func NewImagesResponseWithDefaults() *ImagesResponse`

NewImagesResponseWithDefaults instantiates a new ImagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ImagesResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImagesResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImagesResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetBackground

`func (o *ImagesResponse) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *ImagesResponse) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *ImagesResponse) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *ImagesResponse) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *ImagesResponse) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *ImagesResponse) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetData

`func (o *ImagesResponse) GetData() []IMGImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImagesResponse) GetDataOk() (*[]IMGImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImagesResponse) SetData(v []IMGImage)`

SetData sets Data field to given value.

### HasData

`func (o *ImagesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ImagesResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ImagesResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetOutputFormat

`func (o *ImagesResponse) GetOutputFormat() string`

GetOutputFormat returns the OutputFormat field if non-nil, zero value otherwise.

### GetOutputFormatOk

`func (o *ImagesResponse) GetOutputFormatOk() (*string, bool)`

GetOutputFormatOk returns a tuple with the OutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFormat

`func (o *ImagesResponse) SetOutputFormat(v string)`

SetOutputFormat sets OutputFormat field to given value.

### HasOutputFormat

`func (o *ImagesResponse) HasOutputFormat() bool`

HasOutputFormat returns a boolean if a field has been set.

### SetOutputFormatNil

`func (o *ImagesResponse) SetOutputFormatNil(b bool)`

 SetOutputFormatNil sets the value for OutputFormat to be an explicit nil

### UnsetOutputFormat
`func (o *ImagesResponse) UnsetOutputFormat()`

UnsetOutputFormat ensures that no value is present for OutputFormat, not even an explicit nil
### GetQuality

`func (o *ImagesResponse) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ImagesResponse) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ImagesResponse) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ImagesResponse) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *ImagesResponse) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *ImagesResponse) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetSafetyIdentifier

`func (o *ImagesResponse) GetSafetyIdentifier() string`

GetSafetyIdentifier returns the SafetyIdentifier field if non-nil, zero value otherwise.

### GetSafetyIdentifierOk

`func (o *ImagesResponse) GetSafetyIdentifierOk() (*string, bool)`

GetSafetyIdentifierOk returns a tuple with the SafetyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyIdentifier

`func (o *ImagesResponse) SetSafetyIdentifier(v string)`

SetSafetyIdentifier sets SafetyIdentifier field to given value.

### HasSafetyIdentifier

`func (o *ImagesResponse) HasSafetyIdentifier() bool`

HasSafetyIdentifier returns a boolean if a field has been set.

### SetSafetyIdentifierNil

`func (o *ImagesResponse) SetSafetyIdentifierNil(b bool)`

 SetSafetyIdentifierNil sets the value for SafetyIdentifier to be an explicit nil

### UnsetSafetyIdentifier
`func (o *ImagesResponse) UnsetSafetyIdentifier()`

UnsetSafetyIdentifier ensures that no value is present for SafetyIdentifier, not even an explicit nil
### GetSize

`func (o *ImagesResponse) GetSize() Size`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImagesResponse) GetSizeOk() (*Size, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImagesResponse) SetSize(v Size)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImagesResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ImagesResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ImagesResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSvgUsage

`func (o *ImagesResponse) GetSvgUsage() IMGSvgUsage`

GetSvgUsage returns the SvgUsage field if non-nil, zero value otherwise.

### GetSvgUsageOk

`func (o *ImagesResponse) GetSvgUsageOk() (*IMGSvgUsage, bool)`

GetSvgUsageOk returns a tuple with the SvgUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvgUsage

`func (o *ImagesResponse) SetSvgUsage(v IMGSvgUsage)`

SetSvgUsage sets SvgUsage field to given value.

### HasSvgUsage

`func (o *ImagesResponse) HasSvgUsage() bool`

HasSvgUsage returns a boolean if a field has been set.

### SetSvgUsageNil

`func (o *ImagesResponse) SetSvgUsageNil(b bool)`

 SetSvgUsageNil sets the value for SvgUsage to be an explicit nil

### UnsetSvgUsage
`func (o *ImagesResponse) UnsetSvgUsage()`

UnsetSvgUsage ensures that no value is present for SvgUsage, not even an explicit nil
### GetUsage

`func (o *ImagesResponse) GetUsage() IMGUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ImagesResponse) GetUsageOk() (*IMGUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ImagesResponse) SetUsage(v IMGUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ImagesResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *ImagesResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *ImagesResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


