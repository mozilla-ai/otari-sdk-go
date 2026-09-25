# OpenAIFileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OpenAIFileObject**](OpenAIFileObject.md) |  | 
**FirstId** | **NullableString** |  | 
**HasMore** | **bool** |  | 
**LastId** | **NullableString** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]

## Methods

### NewOpenAIFileList

`func NewOpenAIFileList(data []OpenAIFileObject, firstId NullableString, hasMore bool, lastId NullableString, ) *OpenAIFileList`

NewOpenAIFileList instantiates a new OpenAIFileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAIFileListWithDefaults

`func NewOpenAIFileListWithDefaults() *OpenAIFileList`

NewOpenAIFileListWithDefaults instantiates a new OpenAIFileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OpenAIFileList) GetData() []OpenAIFileObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenAIFileList) GetDataOk() (*[]OpenAIFileObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenAIFileList) SetData(v []OpenAIFileObject)`

SetData sets Data field to given value.


### GetFirstId

`func (o *OpenAIFileList) GetFirstId() string`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *OpenAIFileList) GetFirstIdOk() (*string, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *OpenAIFileList) SetFirstId(v string)`

SetFirstId sets FirstId field to given value.


### SetFirstIdNil

`func (o *OpenAIFileList) SetFirstIdNil(b bool)`

 SetFirstIdNil sets the value for FirstId to be an explicit nil

### UnsetFirstId
`func (o *OpenAIFileList) UnsetFirstId()`

UnsetFirstId ensures that no value is present for FirstId, not even an explicit nil
### GetHasMore

`func (o *OpenAIFileList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OpenAIFileList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OpenAIFileList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLastId

`func (o *OpenAIFileList) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *OpenAIFileList) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *OpenAIFileList) SetLastId(v string)`

SetLastId sets LastId field to given value.


### SetLastIdNil

`func (o *OpenAIFileList) SetLastIdNil(b bool)`

 SetLastIdNil sets the value for LastId to be an explicit nil

### UnsetLastId
`func (o *OpenAIFileList) UnsetLastId()`

UnsetLastId ensures that no value is present for LastId, not even an explicit nil
### GetObject

`func (o *OpenAIFileList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenAIFileList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenAIFileList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenAIFileList) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


