# OpenAIFileDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** |  | [optional] [default to true]
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "file"]

## Methods

### NewOpenAIFileDeleted

`func NewOpenAIFileDeleted(id string, ) *OpenAIFileDeleted`

NewOpenAIFileDeleted instantiates a new OpenAIFileDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAIFileDeletedWithDefaults

`func NewOpenAIFileDeletedWithDefaults() *OpenAIFileDeleted`

NewOpenAIFileDeletedWithDefaults instantiates a new OpenAIFileDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *OpenAIFileDeleted) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OpenAIFileDeleted) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OpenAIFileDeleted) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OpenAIFileDeleted) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *OpenAIFileDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenAIFileDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenAIFileDeleted) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *OpenAIFileDeleted) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenAIFileDeleted) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenAIFileDeleted) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenAIFileDeleted) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


