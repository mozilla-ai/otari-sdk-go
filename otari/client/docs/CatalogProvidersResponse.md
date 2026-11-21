# CatalogProvidersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Providers the caller may see, before paging. | 
**Providers** | [**[]CatalogProviderSummary**](CatalogProviderSummary.md) | The requested page, sorted by display name. | 

## Methods

### NewCatalogProvidersResponse

`func NewCatalogProvidersResponse(count int32, providers []CatalogProviderSummary, ) *CatalogProvidersResponse`

NewCatalogProvidersResponse instantiates a new CatalogProvidersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogProvidersResponseWithDefaults

`func NewCatalogProvidersResponseWithDefaults() *CatalogProvidersResponse`

NewCatalogProvidersResponseWithDefaults instantiates a new CatalogProvidersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CatalogProvidersResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CatalogProvidersResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CatalogProvidersResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetProviders

`func (o *CatalogProvidersResponse) GetProviders() []CatalogProviderSummary`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CatalogProvidersResponse) GetProvidersOk() (*[]CatalogProviderSummary, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CatalogProvidersResponse) SetProviders(v []CatalogProviderSummary)`

SetProviders sets Providers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


