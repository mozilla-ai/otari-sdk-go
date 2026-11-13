# EmbeddingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EmbeddingData**](EmbeddingData.md) |  | 
**Model** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Usage** | Pointer to [**NullableEmbeddingUsage**](EmbeddingUsage.md) | Absent when the provider reported no usage | [optional] 

## Methods

### NewEmbeddingResponse

`func NewEmbeddingResponse(data []EmbeddingData, model string, ) *EmbeddingResponse`

NewEmbeddingResponse instantiates a new EmbeddingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingResponseWithDefaults

`func NewEmbeddingResponseWithDefaults() *EmbeddingResponse`

NewEmbeddingResponseWithDefaults instantiates a new EmbeddingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmbeddingResponse) GetData() []EmbeddingData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmbeddingResponse) GetDataOk() (*[]EmbeddingData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmbeddingResponse) SetData(v []EmbeddingData)`

SetData sets Data field to given value.


### GetModel

`func (o *EmbeddingResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *EmbeddingResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EmbeddingResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EmbeddingResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *EmbeddingResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUsage

`func (o *EmbeddingResponse) GetUsage() EmbeddingUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EmbeddingResponse) GetUsageOk() (*EmbeddingUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EmbeddingResponse) SetUsage(v EmbeddingUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EmbeddingResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *EmbeddingResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *EmbeddingResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


