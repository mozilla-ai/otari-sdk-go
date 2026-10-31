# HostedCatalogUpdatePreviewPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | [**HostedCatalogRefreshPublic**](HostedCatalogRefreshPublic.md) |  | 
**PriceError** | Pointer to **NullableString** |  | [optional] 
**Prices** | [**NullableHostedCatalogPriceUpdatePublic**](HostedCatalogPriceUpdatePublic.md) |  | 

## Methods

### NewHostedCatalogUpdatePreviewPublic

`func NewHostedCatalogUpdatePreviewPublic(catalog HostedCatalogRefreshPublic, prices NullableHostedCatalogPriceUpdatePublic, ) *HostedCatalogUpdatePreviewPublic`

NewHostedCatalogUpdatePreviewPublic instantiates a new HostedCatalogUpdatePreviewPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCatalogUpdatePreviewPublicWithDefaults

`func NewHostedCatalogUpdatePreviewPublicWithDefaults() *HostedCatalogUpdatePreviewPublic`

NewHostedCatalogUpdatePreviewPublicWithDefaults instantiates a new HostedCatalogUpdatePreviewPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *HostedCatalogUpdatePreviewPublic) GetCatalog() HostedCatalogRefreshPublic`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *HostedCatalogUpdatePreviewPublic) GetCatalogOk() (*HostedCatalogRefreshPublic, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *HostedCatalogUpdatePreviewPublic) SetCatalog(v HostedCatalogRefreshPublic)`

SetCatalog sets Catalog field to given value.


### GetPriceError

`func (o *HostedCatalogUpdatePreviewPublic) GetPriceError() string`

GetPriceError returns the PriceError field if non-nil, zero value otherwise.

### GetPriceErrorOk

`func (o *HostedCatalogUpdatePreviewPublic) GetPriceErrorOk() (*string, bool)`

GetPriceErrorOk returns a tuple with the PriceError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceError

`func (o *HostedCatalogUpdatePreviewPublic) SetPriceError(v string)`

SetPriceError sets PriceError field to given value.

### HasPriceError

`func (o *HostedCatalogUpdatePreviewPublic) HasPriceError() bool`

HasPriceError returns a boolean if a field has been set.

### SetPriceErrorNil

`func (o *HostedCatalogUpdatePreviewPublic) SetPriceErrorNil(b bool)`

 SetPriceErrorNil sets the value for PriceError to be an explicit nil

### UnsetPriceError
`func (o *HostedCatalogUpdatePreviewPublic) UnsetPriceError()`

UnsetPriceError ensures that no value is present for PriceError, not even an explicit nil
### GetPrices

`func (o *HostedCatalogUpdatePreviewPublic) GetPrices() HostedCatalogPriceUpdatePublic`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *HostedCatalogUpdatePreviewPublic) GetPricesOk() (*HostedCatalogPriceUpdatePublic, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *HostedCatalogUpdatePreviewPublic) SetPrices(v HostedCatalogPriceUpdatePublic)`

SetPrices sets Prices field to given value.


### SetPricesNil

`func (o *HostedCatalogUpdatePreviewPublic) SetPricesNil(b bool)`

 SetPricesNil sets the value for Prices to be an explicit nil

### UnsetPrices
`func (o *HostedCatalogUpdatePreviewPublic) UnsetPrices()`

UnsetPrices ensures that no value is present for Prices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


