# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "search"]
**Results** | [**[]SearchResultItem**](SearchResultItem.md) |  | 
**SearchTool** | **string** | The configured search tool that served the request | 

## Methods

### NewSearchResponse

`func NewSearchResponse(results []SearchResultItem, searchTool string, ) *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SearchResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SearchResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SearchResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SearchResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetResults

`func (o *SearchResponse) GetResults() []SearchResultItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResponse) GetResultsOk() (*[]SearchResultItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResponse) SetResults(v []SearchResultItem)`

SetResults sets Results field to given value.


### GetSearchTool

`func (o *SearchResponse) GetSearchTool() string`

GetSearchTool returns the SearchTool field if non-nil, zero value otherwise.

### GetSearchToolOk

`func (o *SearchResponse) GetSearchToolOk() (*string, bool)`

GetSearchToolOk returns a tuple with the SearchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTool

`func (o *SearchResponse) SetSearchTool(v string)`

SetSearchTool sets SearchTool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


