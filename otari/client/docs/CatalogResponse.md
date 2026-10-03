# CatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Models matching all filters before paging. | 
**DefaultPricing** | **bool** | Whether an unpriced model is metered at the genai-prices default. | 
**DefaultsAsOf** | **NullableTime** | When the accepted genai-prices snapshot was taken. Null while the bundled dataset serves. | 
**Facets** | Pointer to [**NullableCatalogFacets**](CatalogFacets.md) | Present when include_facets is requested. | [optional] 
**MetadataAvailable** | **bool** | False when models.dev could not be read; descriptions are then absent. | 
**Models** | [**[]CatalogModelSummary**](CatalogModelSummary.md) |  | 

## Methods

### NewCatalogResponse

`func NewCatalogResponse(count int32, defaultPricing bool, defaultsAsOf NullableTime, metadataAvailable bool, models []CatalogModelSummary, ) *CatalogResponse`

NewCatalogResponse instantiates a new CatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogResponseWithDefaults

`func NewCatalogResponseWithDefaults() *CatalogResponse`

NewCatalogResponseWithDefaults instantiates a new CatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CatalogResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CatalogResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CatalogResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDefaultPricing

`func (o *CatalogResponse) GetDefaultPricing() bool`

GetDefaultPricing returns the DefaultPricing field if non-nil, zero value otherwise.

### GetDefaultPricingOk

`func (o *CatalogResponse) GetDefaultPricingOk() (*bool, bool)`

GetDefaultPricingOk returns a tuple with the DefaultPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPricing

`func (o *CatalogResponse) SetDefaultPricing(v bool)`

SetDefaultPricing sets DefaultPricing field to given value.


### GetDefaultsAsOf

`func (o *CatalogResponse) GetDefaultsAsOf() time.Time`

GetDefaultsAsOf returns the DefaultsAsOf field if non-nil, zero value otherwise.

### GetDefaultsAsOfOk

`func (o *CatalogResponse) GetDefaultsAsOfOk() (*time.Time, bool)`

GetDefaultsAsOfOk returns a tuple with the DefaultsAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsAsOf

`func (o *CatalogResponse) SetDefaultsAsOf(v time.Time)`

SetDefaultsAsOf sets DefaultsAsOf field to given value.


### SetDefaultsAsOfNil

`func (o *CatalogResponse) SetDefaultsAsOfNil(b bool)`

 SetDefaultsAsOfNil sets the value for DefaultsAsOf to be an explicit nil

### UnsetDefaultsAsOf
`func (o *CatalogResponse) UnsetDefaultsAsOf()`

UnsetDefaultsAsOf ensures that no value is present for DefaultsAsOf, not even an explicit nil
### GetFacets

`func (o *CatalogResponse) GetFacets() CatalogFacets`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *CatalogResponse) GetFacetsOk() (*CatalogFacets, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *CatalogResponse) SetFacets(v CatalogFacets)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *CatalogResponse) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### SetFacetsNil

`func (o *CatalogResponse) SetFacetsNil(b bool)`

 SetFacetsNil sets the value for Facets to be an explicit nil

### UnsetFacets
`func (o *CatalogResponse) UnsetFacets()`

UnsetFacets ensures that no value is present for Facets, not even an explicit nil
### GetMetadataAvailable

`func (o *CatalogResponse) GetMetadataAvailable() bool`

GetMetadataAvailable returns the MetadataAvailable field if non-nil, zero value otherwise.

### GetMetadataAvailableOk

`func (o *CatalogResponse) GetMetadataAvailableOk() (*bool, bool)`

GetMetadataAvailableOk returns a tuple with the MetadataAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAvailable

`func (o *CatalogResponse) SetMetadataAvailable(v bool)`

SetMetadataAvailable sets MetadataAvailable field to given value.


### GetModels

`func (o *CatalogResponse) GetModels() []CatalogModelSummary`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CatalogResponse) GetModelsOk() (*[]CatalogModelSummary, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CatalogResponse) SetModels(v []CatalogModelSummary)`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


