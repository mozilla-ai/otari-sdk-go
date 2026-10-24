# HostedCatalogRefreshPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kept** | [**[]HostedCatalogKeptGroupPublic**](HostedCatalogKeptGroupPublic.md) |  | 
**OrphanedRosters** | Pointer to **[]string** |  | [optional] 
**Providers** | [**[]HostedCatalogProviderRefreshPublic**](HostedCatalogProviderRefreshPublic.md) |  | 
**Removed** | **[]string** |  | 
**RemovedCount** | **int32** |  | 
**RemovedOverrideRows** | **int32** |  | 
**RemovedPriceRows** | **int32** |  | 

## Methods

### NewHostedCatalogRefreshPublic

`func NewHostedCatalogRefreshPublic(kept []HostedCatalogKeptGroupPublic, providers []HostedCatalogProviderRefreshPublic, removed []string, removedCount int32, removedOverrideRows int32, removedPriceRows int32, ) *HostedCatalogRefreshPublic`

NewHostedCatalogRefreshPublic instantiates a new HostedCatalogRefreshPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCatalogRefreshPublicWithDefaults

`func NewHostedCatalogRefreshPublicWithDefaults() *HostedCatalogRefreshPublic`

NewHostedCatalogRefreshPublicWithDefaults instantiates a new HostedCatalogRefreshPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKept

`func (o *HostedCatalogRefreshPublic) GetKept() []HostedCatalogKeptGroupPublic`

GetKept returns the Kept field if non-nil, zero value otherwise.

### GetKeptOk

`func (o *HostedCatalogRefreshPublic) GetKeptOk() (*[]HostedCatalogKeptGroupPublic, bool)`

GetKeptOk returns a tuple with the Kept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKept

`func (o *HostedCatalogRefreshPublic) SetKept(v []HostedCatalogKeptGroupPublic)`

SetKept sets Kept field to given value.


### GetOrphanedRosters

`func (o *HostedCatalogRefreshPublic) GetOrphanedRosters() []string`

GetOrphanedRosters returns the OrphanedRosters field if non-nil, zero value otherwise.

### GetOrphanedRostersOk

`func (o *HostedCatalogRefreshPublic) GetOrphanedRostersOk() (*[]string, bool)`

GetOrphanedRostersOk returns a tuple with the OrphanedRosters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedRosters

`func (o *HostedCatalogRefreshPublic) SetOrphanedRosters(v []string)`

SetOrphanedRosters sets OrphanedRosters field to given value.

### HasOrphanedRosters

`func (o *HostedCatalogRefreshPublic) HasOrphanedRosters() bool`

HasOrphanedRosters returns a boolean if a field has been set.

### GetProviders

`func (o *HostedCatalogRefreshPublic) GetProviders() []HostedCatalogProviderRefreshPublic`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *HostedCatalogRefreshPublic) GetProvidersOk() (*[]HostedCatalogProviderRefreshPublic, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *HostedCatalogRefreshPublic) SetProviders(v []HostedCatalogProviderRefreshPublic)`

SetProviders sets Providers field to given value.


### GetRemoved

`func (o *HostedCatalogRefreshPublic) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *HostedCatalogRefreshPublic) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *HostedCatalogRefreshPublic) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.


### GetRemovedCount

`func (o *HostedCatalogRefreshPublic) GetRemovedCount() int32`

GetRemovedCount returns the RemovedCount field if non-nil, zero value otherwise.

### GetRemovedCountOk

`func (o *HostedCatalogRefreshPublic) GetRemovedCountOk() (*int32, bool)`

GetRemovedCountOk returns a tuple with the RemovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCount

`func (o *HostedCatalogRefreshPublic) SetRemovedCount(v int32)`

SetRemovedCount sets RemovedCount field to given value.


### GetRemovedOverrideRows

`func (o *HostedCatalogRefreshPublic) GetRemovedOverrideRows() int32`

GetRemovedOverrideRows returns the RemovedOverrideRows field if non-nil, zero value otherwise.

### GetRemovedOverrideRowsOk

`func (o *HostedCatalogRefreshPublic) GetRemovedOverrideRowsOk() (*int32, bool)`

GetRemovedOverrideRowsOk returns a tuple with the RemovedOverrideRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedOverrideRows

`func (o *HostedCatalogRefreshPublic) SetRemovedOverrideRows(v int32)`

SetRemovedOverrideRows sets RemovedOverrideRows field to given value.


### GetRemovedPriceRows

`func (o *HostedCatalogRefreshPublic) GetRemovedPriceRows() int32`

GetRemovedPriceRows returns the RemovedPriceRows field if non-nil, zero value otherwise.

### GetRemovedPriceRowsOk

`func (o *HostedCatalogRefreshPublic) GetRemovedPriceRowsOk() (*int32, bool)`

GetRemovedPriceRowsOk returns a tuple with the RemovedPriceRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedPriceRows

`func (o *HostedCatalogRefreshPublic) SetRemovedPriceRows(v int32)`

SetRemovedPriceRows sets RemovedPriceRows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


