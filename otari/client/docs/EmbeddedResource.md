# EmbeddedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to [**NullableAnnotations**](Annotations.md) |  | [optional] 
**Resource** | [**Resource**](Resource.md) |  | 
**Type** | **string** |  | 

## Methods

### NewEmbeddedResource

`func NewEmbeddedResource(resource Resource, type_ string, ) *EmbeddedResource`

NewEmbeddedResource instantiates a new EmbeddedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedResourceWithDefaults

`func NewEmbeddedResourceWithDefaults() *EmbeddedResource`

NewEmbeddedResourceWithDefaults instantiates a new EmbeddedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *EmbeddedResource) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EmbeddedResource) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EmbeddedResource) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EmbeddedResource) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *EmbeddedResource) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *EmbeddedResource) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAnnotations

`func (o *EmbeddedResource) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EmbeddedResource) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EmbeddedResource) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EmbeddedResource) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *EmbeddedResource) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *EmbeddedResource) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetResource

`func (o *EmbeddedResource) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EmbeddedResource) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EmbeddedResource) SetResource(v Resource)`

SetResource sets Resource field to given value.


### GetType

`func (o *EmbeddedResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmbeddedResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmbeddedResource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


