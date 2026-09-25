# ResponseFilesListFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnthropicFileMetadata**](AnthropicFileMetadata.md) |  | 
**FirstId** | **string** |  | 
**HasMore** | **bool** |  | 
**LastId** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**NextPage** | **string** |  | 

## Methods

### NewResponseFilesListFiles

`func NewResponseFilesListFiles(data []AnthropicFileMetadata, firstId string, hasMore bool, lastId string, nextPage string, ) *ResponseFilesListFiles`

NewResponseFilesListFiles instantiates a new ResponseFilesListFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFilesListFilesWithDefaults

`func NewResponseFilesListFilesWithDefaults() *ResponseFilesListFiles`

NewResponseFilesListFilesWithDefaults instantiates a new ResponseFilesListFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseFilesListFiles) GetData() []AnthropicFileMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseFilesListFiles) GetDataOk() (*[]AnthropicFileMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseFilesListFiles) SetData(v []AnthropicFileMetadata)`

SetData sets Data field to given value.


### GetFirstId

`func (o *ResponseFilesListFiles) GetFirstId() string`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *ResponseFilesListFiles) GetFirstIdOk() (*string, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *ResponseFilesListFiles) SetFirstId(v string)`

SetFirstId sets FirstId field to given value.


### GetHasMore

`func (o *ResponseFilesListFiles) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ResponseFilesListFiles) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ResponseFilesListFiles) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLastId

`func (o *ResponseFilesListFiles) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *ResponseFilesListFiles) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *ResponseFilesListFiles) SetLastId(v string)`

SetLastId sets LastId field to given value.


### GetObject

`func (o *ResponseFilesListFiles) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseFilesListFiles) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseFilesListFiles) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseFilesListFiles) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetNextPage

`func (o *ResponseFilesListFiles) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ResponseFilesListFiles) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ResponseFilesListFiles) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


