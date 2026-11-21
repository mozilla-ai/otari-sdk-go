# SearchToolCatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SearchToolCatalogEntry**](SearchToolCatalogEntry.md) |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]

## Methods

### NewSearchToolCatalogResponse

`func NewSearchToolCatalogResponse(data []SearchToolCatalogEntry, ) *SearchToolCatalogResponse`

NewSearchToolCatalogResponse instantiates a new SearchToolCatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchToolCatalogResponseWithDefaults

`func NewSearchToolCatalogResponseWithDefaults() *SearchToolCatalogResponse`

NewSearchToolCatalogResponseWithDefaults instantiates a new SearchToolCatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchToolCatalogResponse) GetData() []SearchToolCatalogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchToolCatalogResponse) GetDataOk() (*[]SearchToolCatalogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchToolCatalogResponse) SetData(v []SearchToolCatalogEntry)`

SetData sets Data field to given value.


### GetObject

`func (o *SearchToolCatalogResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SearchToolCatalogResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SearchToolCatalogResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SearchToolCatalogResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


