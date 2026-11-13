# EmbeddingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedding** | [**Embedding**](Embedding.md) |  | 
**Index** | **int32** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "embedding"]

## Methods

### NewEmbeddingData

`func NewEmbeddingData(embedding Embedding, index int32, ) *EmbeddingData`

NewEmbeddingData instantiates a new EmbeddingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingDataWithDefaults

`func NewEmbeddingDataWithDefaults() *EmbeddingData`

NewEmbeddingDataWithDefaults instantiates a new EmbeddingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedding

`func (o *EmbeddingData) GetEmbedding() Embedding`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *EmbeddingData) GetEmbeddingOk() (*Embedding, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *EmbeddingData) SetEmbedding(v Embedding)`

SetEmbedding sets Embedding field to given value.


### GetIndex

`func (o *EmbeddingData) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *EmbeddingData) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *EmbeddingData) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetObject

`func (o *EmbeddingData) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EmbeddingData) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EmbeddingData) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *EmbeddingData) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


