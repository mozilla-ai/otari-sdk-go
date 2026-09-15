# ResourceLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to [**NullableAnnotations**](Annotations.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Icons** | Pointer to [**[]Icon**](Icon.md) |  | [optional] 
**MimeType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Uri** | **string** |  | 

## Methods

### NewResourceLink

`func NewResourceLink(name string, type_ string, uri string, ) *ResourceLink`

NewResourceLink instantiates a new ResourceLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLinkWithDefaults

`func NewResourceLinkWithDefaults() *ResourceLink`

NewResourceLinkWithDefaults instantiates a new ResourceLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ResourceLink) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResourceLink) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResourceLink) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ResourceLink) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ResourceLink) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ResourceLink) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAnnotations

`func (o *ResourceLink) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ResourceLink) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ResourceLink) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ResourceLink) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *ResourceLink) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *ResourceLink) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetDescription

`func (o *ResourceLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourceLink) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceLink) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcons

`func (o *ResourceLink) GetIcons() []Icon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *ResourceLink) GetIconsOk() (*[]Icon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *ResourceLink) SetIcons(v []Icon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *ResourceLink) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIconsNil

`func (o *ResourceLink) SetIconsNil(b bool)`

 SetIconsNil sets the value for Icons to be an explicit nil

### UnsetIcons
`func (o *ResourceLink) UnsetIcons()`

UnsetIcons ensures that no value is present for Icons, not even an explicit nil
### GetMimeType

`func (o *ResourceLink) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResourceLink) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResourceLink) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResourceLink) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *ResourceLink) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *ResourceLink) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetName

`func (o *ResourceLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceLink) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ResourceLink) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResourceLink) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResourceLink) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResourceLink) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ResourceLink) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ResourceLink) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetTitle

`func (o *ResourceLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResourceLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResourceLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResourceLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResourceLink) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResourceLink) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *ResourceLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceLink) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *ResourceLink) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ResourceLink) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ResourceLink) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


