# ResponseFilesCreateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** |  | 
**CreatedAt** | **string** |  | 
**ExpiresAt** | **string** |  | 
**Filename** | **string** |  | 
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "file"]
**Purpose** | **string** |  | 
**Downloadable** | Pointer to **bool** |  | [optional] [default to true]
**MimeType** | **string** |  | 
**SizeBytes** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "file"]

## Methods

### NewResponseFilesCreateFile

`func NewResponseFilesCreateFile(bytes int32, createdAt string, expiresAt string, filename string, id string, purpose string, mimeType string, sizeBytes int32, ) *ResponseFilesCreateFile`

NewResponseFilesCreateFile instantiates a new ResponseFilesCreateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFilesCreateFileWithDefaults

`func NewResponseFilesCreateFileWithDefaults() *ResponseFilesCreateFile`

NewResponseFilesCreateFileWithDefaults instantiates a new ResponseFilesCreateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *ResponseFilesCreateFile) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ResponseFilesCreateFile) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ResponseFilesCreateFile) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *ResponseFilesCreateFile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseFilesCreateFile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseFilesCreateFile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *ResponseFilesCreateFile) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ResponseFilesCreateFile) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ResponseFilesCreateFile) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFilename

`func (o *ResponseFilesCreateFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ResponseFilesCreateFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ResponseFilesCreateFile) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetId

`func (o *ResponseFilesCreateFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseFilesCreateFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseFilesCreateFile) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ResponseFilesCreateFile) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseFilesCreateFile) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseFilesCreateFile) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseFilesCreateFile) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPurpose

`func (o *ResponseFilesCreateFile) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ResponseFilesCreateFile) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ResponseFilesCreateFile) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetDownloadable

`func (o *ResponseFilesCreateFile) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ResponseFilesCreateFile) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ResponseFilesCreateFile) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ResponseFilesCreateFile) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseFilesCreateFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseFilesCreateFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseFilesCreateFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetSizeBytes

`func (o *ResponseFilesCreateFile) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ResponseFilesCreateFile) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ResponseFilesCreateFile) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetType

`func (o *ResponseFilesCreateFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseFilesCreateFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseFilesCreateFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseFilesCreateFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


