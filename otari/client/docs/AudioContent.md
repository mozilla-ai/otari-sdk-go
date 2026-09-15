# AudioContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to [**NullableAnnotations**](Annotations.md) |  | [optional] 
**Data** | **string** |  | 
**MimeType** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewAudioContent

`func NewAudioContent(data string, mimeType string, type_ string, ) *AudioContent`

NewAudioContent instantiates a new AudioContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioContentWithDefaults

`func NewAudioContentWithDefaults() *AudioContent`

NewAudioContentWithDefaults instantiates a new AudioContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AudioContent) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AudioContent) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AudioContent) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AudioContent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *AudioContent) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *AudioContent) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAnnotations

`func (o *AudioContent) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AudioContent) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AudioContent) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AudioContent) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *AudioContent) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *AudioContent) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetData

`func (o *AudioContent) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AudioContent) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AudioContent) SetData(v string)`

SetData sets Data field to given value.


### GetMimeType

`func (o *AudioContent) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AudioContent) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AudioContent) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetType

`func (o *AudioContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudioContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudioContent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


