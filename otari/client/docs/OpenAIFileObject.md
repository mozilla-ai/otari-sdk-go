# OpenAIFileObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** |  | 
**CreatedAt** | **int32** |  | 
**ExpiresAt** | **NullableInt32** |  | 
**Filename** | **string** |  | 
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "file"]
**Purpose** | **string** |  | 

## Methods

### NewOpenAIFileObject

`func NewOpenAIFileObject(bytes int32, createdAt int32, expiresAt NullableInt32, filename string, id string, purpose string, ) *OpenAIFileObject`

NewOpenAIFileObject instantiates a new OpenAIFileObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenAIFileObjectWithDefaults

`func NewOpenAIFileObjectWithDefaults() *OpenAIFileObject`

NewOpenAIFileObjectWithDefaults instantiates a new OpenAIFileObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *OpenAIFileObject) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *OpenAIFileObject) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *OpenAIFileObject) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *OpenAIFileObject) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenAIFileObject) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenAIFileObject) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *OpenAIFileObject) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OpenAIFileObject) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OpenAIFileObject) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *OpenAIFileObject) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OpenAIFileObject) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFilename

`func (o *OpenAIFileObject) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *OpenAIFileObject) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *OpenAIFileObject) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetId

`func (o *OpenAIFileObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenAIFileObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenAIFileObject) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *OpenAIFileObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenAIFileObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenAIFileObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenAIFileObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPurpose

`func (o *OpenAIFileObject) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *OpenAIFileObject) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *OpenAIFileObject) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


