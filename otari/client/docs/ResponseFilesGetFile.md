# ResponseFilesGetFile

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

### NewResponseFilesGetFile

`func NewResponseFilesGetFile(bytes int32, createdAt string, expiresAt string, filename string, id string, purpose string, mimeType string, sizeBytes int32, ) *ResponseFilesGetFile`

NewResponseFilesGetFile instantiates a new ResponseFilesGetFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFilesGetFileWithDefaults

`func NewResponseFilesGetFileWithDefaults() *ResponseFilesGetFile`

NewResponseFilesGetFileWithDefaults instantiates a new ResponseFilesGetFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *ResponseFilesGetFile) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ResponseFilesGetFile) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ResponseFilesGetFile) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetCreatedAt

`func (o *ResponseFilesGetFile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseFilesGetFile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseFilesGetFile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *ResponseFilesGetFile) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ResponseFilesGetFile) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ResponseFilesGetFile) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetFilename

`func (o *ResponseFilesGetFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ResponseFilesGetFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ResponseFilesGetFile) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetId

`func (o *ResponseFilesGetFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseFilesGetFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseFilesGetFile) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ResponseFilesGetFile) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ResponseFilesGetFile) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ResponseFilesGetFile) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ResponseFilesGetFile) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPurpose

`func (o *ResponseFilesGetFile) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ResponseFilesGetFile) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ResponseFilesGetFile) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetDownloadable

`func (o *ResponseFilesGetFile) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ResponseFilesGetFile) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ResponseFilesGetFile) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ResponseFilesGetFile) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetMimeType

`func (o *ResponseFilesGetFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseFilesGetFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseFilesGetFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetSizeBytes

`func (o *ResponseFilesGetFile) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ResponseFilesGetFile) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ResponseFilesGetFile) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetType

`func (o *ResponseFilesGetFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseFilesGetFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseFilesGetFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseFilesGetFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


