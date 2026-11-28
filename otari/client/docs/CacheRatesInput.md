# CacheRatesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explicit** | Pointer to [**NullableCacheRatesInput**](CacheRatesInput.md) | The prices of explicit caching, where apart. | [optional] 
**Read** | Pointer to **NullableFloat32** | Price of a cached read. | [optional] 
**Storage** | Pointer to **NullableFloat32** | Price of cache storage. | [optional] 
**Write** | Pointer to [**NullableWrite**](Write.md) |  | [optional] 

## Methods

### NewCacheRatesInput

`func NewCacheRatesInput() *CacheRatesInput`

NewCacheRatesInput instantiates a new CacheRatesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheRatesInputWithDefaults

`func NewCacheRatesInputWithDefaults() *CacheRatesInput`

NewCacheRatesInputWithDefaults instantiates a new CacheRatesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplicit

`func (o *CacheRatesInput) GetExplicit() CacheRatesInput`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *CacheRatesInput) GetExplicitOk() (*CacheRatesInput, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *CacheRatesInput) SetExplicit(v CacheRatesInput)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *CacheRatesInput) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### SetExplicitNil

`func (o *CacheRatesInput) SetExplicitNil(b bool)`

 SetExplicitNil sets the value for Explicit to be an explicit nil

### UnsetExplicit
`func (o *CacheRatesInput) UnsetExplicit()`

UnsetExplicit ensures that no value is present for Explicit, not even an explicit nil
### GetRead

`func (o *CacheRatesInput) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *CacheRatesInput) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *CacheRatesInput) SetRead(v float32)`

SetRead sets Read field to given value.

### HasRead

`func (o *CacheRatesInput) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *CacheRatesInput) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *CacheRatesInput) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetStorage

`func (o *CacheRatesInput) GetStorage() float32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CacheRatesInput) GetStorageOk() (*float32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CacheRatesInput) SetStorage(v float32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CacheRatesInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CacheRatesInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CacheRatesInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetWrite

`func (o *CacheRatesInput) GetWrite() Write`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *CacheRatesInput) GetWriteOk() (*Write, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *CacheRatesInput) SetWrite(v Write)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *CacheRatesInput) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### SetWriteNil

`func (o *CacheRatesInput) SetWriteNil(b bool)`

 SetWriteNil sets the value for Write to be an explicit nil

### UnsetWrite
`func (o *CacheRatesInput) UnsetWrite()`

UnsetWrite ensures that no value is present for Write, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


