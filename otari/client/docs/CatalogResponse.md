# CatalogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPricing** | **bool** | Whether an unpriced model is metered at the genai-prices default. | 
**DefaultsAsOf** | **NullableTime** | When the accepted genai-prices snapshot was taken. Null while the bundled dataset serves. | 
**MetadataAvailable** | **bool** | False when models.dev could not be read; descriptions are then absent. | 
**Models** | [**[]CatalogModelSummary**](CatalogModelSummary.md) |  | 

## Methods

### NewCatalogResponse

`func NewCatalogResponse(defaultPricing bool, defaultsAsOf NullableTime, metadataAvailable bool, models []CatalogModelSummary, ) *CatalogResponse`

NewCatalogResponse instantiates a new CatalogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogResponseWithDefaults

`func NewCatalogResponseWithDefaults() *CatalogResponse`

NewCatalogResponseWithDefaults instantiates a new CatalogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


