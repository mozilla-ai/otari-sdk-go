# ResponseFilesDeleteFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** |  | [optional] [default to true]
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "file"]
**Type** | Pointer to **string** |  | [optional] [default to "file_deleted"]

## Methods

### NewResponseFilesDeleteFile

`func NewResponseFilesDeleteFile(id string, ) *ResponseFilesDeleteFile`

NewResponseFilesDeleteFile instantiates a new ResponseFilesDeleteFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFilesDeleteFileWithDefaults

`func NewResponseFilesDeleteFileWithDefaults() *ResponseFilesDeleteFile`

NewResponseFilesDeleteFileWithDefaults instantiates a new ResponseFilesDeleteFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *ResponseFilesDeleteFile) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ResponseFilesDeleteFile) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ResponseFilesDeleteFile) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ResponseFilesDeleteFile) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *ResponseFilesDeleteFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseFilesDeleteFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseFilesDeleteFile) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ResponseFilesDeleteFile) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseFilesDeleteFile) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseFilesDeleteFile) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseFilesDeleteFile) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *ResponseFilesDeleteFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseFilesDeleteFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseFilesDeleteFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseFilesDeleteFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


