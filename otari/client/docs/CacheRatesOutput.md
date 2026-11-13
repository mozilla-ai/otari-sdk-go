# CacheRatesOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explicit** | Pointer to [**NullableCacheRatesOutput**](CacheRatesOutput.md) | The prices of explicit caching, where apart. | [optional] 
**Read** | Pointer to **NullableFloat32** | Price of a cached read. | [optional] 
**Storage** | Pointer to **NullableFloat32** | Price of cache storage. | [optional] 
**Write** | Pointer to [**NullableWrite**](Write.md) |  | [optional] 

## Methods

### NewCacheRatesOutput

`func NewCacheRatesOutput() *CacheRatesOutput`

NewCacheRatesOutput instantiates a new CacheRatesOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheRatesOutputWithDefaults

`func NewCacheRatesOutputWithDefaults() *CacheRatesOutput`

NewCacheRatesOutputWithDefaults instantiates a new CacheRatesOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplicit

`func (o *CacheRatesOutput) GetExplicit() CacheRatesOutput`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *CacheRatesOutput) GetExplicitOk() (*CacheRatesOutput, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *CacheRatesOutput) SetExplicit(v CacheRatesOutput)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *CacheRatesOutput) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### SetExplicitNil

`func (o *CacheRatesOutput) SetExplicitNil(b bool)`

 SetExplicitNil sets the value for Explicit to be an explicit nil

### UnsetExplicit
`func (o *CacheRatesOutput) UnsetExplicit()`

UnsetExplicit ensures that no value is present for Explicit, not even an explicit nil
### GetRead

`func (o *CacheRatesOutput) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *CacheRatesOutput) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *CacheRatesOutput) SetRead(v float32)`

SetRead sets Read field to given value.

### HasRead

`func (o *CacheRatesOutput) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *CacheRatesOutput) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *CacheRatesOutput) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetStorage

`func (o *CacheRatesOutput) GetStorage() float32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CacheRatesOutput) GetStorageOk() (*float32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CacheRatesOutput) SetStorage(v float32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CacheRatesOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CacheRatesOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CacheRatesOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetWrite

`func (o *CacheRatesOutput) GetWrite() Write`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *CacheRatesOutput) GetWriteOk() (*Write, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *CacheRatesOutput) SetWrite(v Write)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *CacheRatesOutput) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### SetWriteNil

`func (o *CacheRatesOutput) SetWriteNil(b bool)`

 SetWriteNil sets the value for Write to be an explicit nil

### UnsetWrite
`func (o *CacheRatesOutput) UnsetWrite()`

UnsetWrite ensures that no value is present for Write, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


