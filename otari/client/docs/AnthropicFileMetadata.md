# AnthropicFileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**Downloadable** | Pointer to **bool** |  | [optional] [default to true]
**ExpiresAt** | **NullableString** |  | 
**Filename** | **string** |  | 
**Id** | **string** |  | 
**MimeType** | **string** |  | 
**SizeBytes** | **int32** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "file"]

## Methods

### NewAnthropicFileMetadata

`func NewAnthropicFileMetadata(createdAt string, expiresAt NullableString, filename string, id string, mimeType string, sizeBytes int32, ) *AnthropicFileMetadata`

NewAnthropicFileMetadata instantiates a new AnthropicFileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicFileMetadataWithDefaults

`func NewAnthropicFileMetadataWithDefaults() *AnthropicFileMetadata`

NewAnthropicFileMetadataWithDefaults instantiates a new AnthropicFileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AnthropicFileMetadata) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnthropicFileMetadata) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnthropicFileMetadata) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownloadable

`func (o *AnthropicFileMetadata) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *AnthropicFileMetadata) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *AnthropicFileMetadata) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *AnthropicFileMetadata) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AnthropicFileMetadata) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AnthropicFileMetadata) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AnthropicFileMetadata) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *AnthropicFileMetadata) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *AnthropicFileMetadata) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFilename

`func (o *AnthropicFileMetadata) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AnthropicFileMetadata) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AnthropicFileMetadata) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetId

`func (o *AnthropicFileMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnthropicFileMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnthropicFileMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetMimeType

`func (o *AnthropicFileMetadata) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AnthropicFileMetadata) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AnthropicFileMetadata) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetSizeBytes

`func (o *AnthropicFileMetadata) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *AnthropicFileMetadata) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *AnthropicFileMetadata) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetType

`func (o *AnthropicFileMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnthropicFileMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnthropicFileMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnthropicFileMetadata) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


