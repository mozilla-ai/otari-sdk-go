# ContentInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to [**Annotations**](Annotations.md) |  | [optional] 
**Text** | **string** |  | 
**Type** | **string** |  | 
**Data** | **string** |  | 
**MimeType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Icons** | Pointer to [**[]Icon**](Icon.md) |  | [optional] 
**Name** | **string** |  | 
**Size** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Uri** | **string** |  | 
**Resource** | [**Resource**](Resource.md) |  | 

## Methods

### NewContentInner

`func NewContentInner(text string, type_ string, data string, mimeType string, name string, uri string, resource Resource, ) *ContentInner`

NewContentInner instantiates a new ContentInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentInnerWithDefaults

`func NewContentInnerWithDefaults() *ContentInner`

NewContentInnerWithDefaults instantiates a new ContentInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ContentInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContentInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContentInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ContentInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAnnotations

`func (o *ContentInner) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ContentInner) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ContentInner) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ContentInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetText

`func (o *ContentInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ContentInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ContentInner) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ContentInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentInner) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *ContentInner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentInner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentInner) SetData(v string)`

SetData sets Data field to given value.


### GetMimeType

`func (o *ContentInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ContentInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ContentInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetDescription

`func (o *ContentInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcons

`func (o *ContentInner) GetIcons() []Icon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *ContentInner) GetIconsOk() (*[]Icon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *ContentInner) SetIcons(v []Icon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *ContentInner) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### GetName

`func (o *ContentInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentInner) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ContentInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTitle

`func (o *ContentInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUri

`func (o *ContentInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ContentInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ContentInner) SetUri(v string)`

SetUri sets Uri field to given value.


### GetResource

`func (o *ContentInner) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ContentInner) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ContentInner) SetResource(v Resource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


