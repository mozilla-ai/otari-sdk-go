# CatalogFacets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | **[]string** | Every provider instance offering an authorized model, sorted. | 
**TotalCount** | **int32** | Authorized models before any of the request&#39;s filters. | 
**Vendors** | [**[]CatalogVendorFacet**](CatalogVendorFacet.md) |  | 

## Methods

### NewCatalogFacets

`func NewCatalogFacets(providers []string, totalCount int32, vendors []CatalogVendorFacet, ) *CatalogFacets`

NewCatalogFacets instantiates a new CatalogFacets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogFacetsWithDefaults

`func NewCatalogFacetsWithDefaults() *CatalogFacets`

NewCatalogFacetsWithDefaults instantiates a new CatalogFacets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *CatalogFacets) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CatalogFacets) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CatalogFacets) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetTotalCount

`func (o *CatalogFacets) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CatalogFacets) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CatalogFacets) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetVendors

`func (o *CatalogFacets) GetVendors() []CatalogVendorFacet`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *CatalogFacets) GetVendorsOk() (*[]CatalogVendorFacet, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *CatalogFacets) SetVendors(v []CatalogVendorFacet)`

SetVendors sets Vendors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


