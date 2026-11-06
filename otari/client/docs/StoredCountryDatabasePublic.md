# StoredCountryDatabasePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltAt** | **time.Time** | When the file was built, from its metadata. | 
**DatabaseType** | **string** | The type the file&#39;s metadata names, such as &#x60;GeoLite2-Country&#x60; or &#x60;DBIP-Country-Lite&#x60;. | 
**Sha256** | **string** | The file&#39;s SHA-256, in hex. | 
**SizeBytes** | **int32** |  | 
**UploadedAt** | **time.Time** |  | 

## Methods

### NewStoredCountryDatabasePublic

`func NewStoredCountryDatabasePublic(builtAt time.Time, databaseType string, sha256 string, sizeBytes int32, uploadedAt time.Time, ) *StoredCountryDatabasePublic`

NewStoredCountryDatabasePublic instantiates a new StoredCountryDatabasePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredCountryDatabasePublicWithDefaults

`func NewStoredCountryDatabasePublicWithDefaults() *StoredCountryDatabasePublic`

NewStoredCountryDatabasePublicWithDefaults instantiates a new StoredCountryDatabasePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltAt

`func (o *StoredCountryDatabasePublic) GetBuiltAt() time.Time`

GetBuiltAt returns the BuiltAt field if non-nil, zero value otherwise.

### GetBuiltAtOk

`func (o *StoredCountryDatabasePublic) GetBuiltAtOk() (*time.Time, bool)`

GetBuiltAtOk returns a tuple with the BuiltAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltAt

`func (o *StoredCountryDatabasePublic) SetBuiltAt(v time.Time)`

SetBuiltAt sets BuiltAt field to given value.


### GetDatabaseType

`func (o *StoredCountryDatabasePublic) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *StoredCountryDatabasePublic) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *StoredCountryDatabasePublic) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.


### GetSha256

`func (o *StoredCountryDatabasePublic) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *StoredCountryDatabasePublic) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *StoredCountryDatabasePublic) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.


### GetSizeBytes

`func (o *StoredCountryDatabasePublic) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *StoredCountryDatabasePublic) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *StoredCountryDatabasePublic) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.


### GetUploadedAt

`func (o *StoredCountryDatabasePublic) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *StoredCountryDatabasePublic) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *StoredCountryDatabasePublic) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


