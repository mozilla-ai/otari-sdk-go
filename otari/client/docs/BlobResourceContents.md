# BlobResourceContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Blob** | **string** |  | 
**MimeType** | Pointer to **NullableString** |  | [optional] 
**Uri** | **string** |  | 

## Methods

### NewBlobResourceContents

`func NewBlobResourceContents(blob string, uri string, ) *BlobResourceContents`

NewBlobResourceContents instantiates a new BlobResourceContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobResourceContentsWithDefaults

`func NewBlobResourceContentsWithDefaults() *BlobResourceContents`

NewBlobResourceContentsWithDefaults instantiates a new BlobResourceContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BlobResourceContents) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BlobResourceContents) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BlobResourceContents) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BlobResourceContents) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *BlobResourceContents) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *BlobResourceContents) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetBlob

`func (o *BlobResourceContents) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *BlobResourceContents) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *BlobResourceContents) SetBlob(v string)`

SetBlob sets Blob field to given value.


### GetMimeType

`func (o *BlobResourceContents) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BlobResourceContents) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BlobResourceContents) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BlobResourceContents) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *BlobResourceContents) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *BlobResourceContents) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetUri

`func (o *BlobResourceContents) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BlobResourceContents) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BlobResourceContents) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


