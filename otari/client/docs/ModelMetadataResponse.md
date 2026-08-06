# ModelMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | False when metadata could not be loaded (enrichment disabled or models.dev unreachable). | 
**Models** | Pointer to [**map[string]ModelMetadata**](ModelMetadata.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] [default to "models.dev"]

## Methods

### NewModelMetadataResponse

`func NewModelMetadataResponse(available bool, ) *ModelMetadataResponse`

NewModelMetadataResponse instantiates a new ModelMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataResponseWithDefaults

`func NewModelMetadataResponseWithDefaults() *ModelMetadataResponse`

NewModelMetadataResponseWithDefaults instantiates a new ModelMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ModelMetadataResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ModelMetadataResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ModelMetadataResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetModels

`func (o *ModelMetadataResponse) GetModels() map[string]ModelMetadata`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ModelMetadataResponse) GetModelsOk() (*map[string]ModelMetadata, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ModelMetadataResponse) SetModels(v map[string]ModelMetadata)`

SetModels sets Models field to given value.

### HasModels

`func (o *ModelMetadataResponse) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetSource

`func (o *ModelMetadataResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ModelMetadataResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ModelMetadataResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ModelMetadataResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


