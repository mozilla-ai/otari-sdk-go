# CatalogProviderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextLength** | Pointer to **NullableInt32** | The largest context window this instance serves. | [optional] 
**Id** | **string** | The catalog id, the one &#x60;GET /catalog/models/{id}&#x60; takes. | 
**InputPrice** | Pointer to **NullableFloat32** | This instance&#39;s input rate for the caller, in USD per million of &#x60;price_unit&#x60;, as the catalog&#39;s &#x60;min_input_price_per_million&#x60; is. The cheapest where the instance offers the model more than once. | [optional] 
**Name** | **string** |  | 
**OutputPrice** | Pointer to **NullableFloat32** | This instance&#39;s output rate for the caller, in the same unit as &#x60;input_price&#x60;. | [optional] 
**PriceUnit** | Pointer to **NullableString** | What the prices are per. Null where nothing on this instance prices the model. | [optional] 
**Vendor** | **NullableString** |  | 

## Methods

### NewCatalogProviderModel

`func NewCatalogProviderModel(id string, name string, vendor NullableString, ) *CatalogProviderModel`

NewCatalogProviderModel instantiates a new CatalogProviderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogProviderModelWithDefaults

`func NewCatalogProviderModelWithDefaults() *CatalogProviderModel`

NewCatalogProviderModelWithDefaults instantiates a new CatalogProviderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextLength

`func (o *CatalogProviderModel) GetContextLength() int32`

GetContextLength returns the ContextLength field if non-nil, zero value otherwise.

### GetContextLengthOk

`func (o *CatalogProviderModel) GetContextLengthOk() (*int32, bool)`

GetContextLengthOk returns a tuple with the ContextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLength

`func (o *CatalogProviderModel) SetContextLength(v int32)`

SetContextLength sets ContextLength field to given value.

### HasContextLength

`func (o *CatalogProviderModel) HasContextLength() bool`

HasContextLength returns a boolean if a field has been set.

### SetContextLengthNil

`func (o *CatalogProviderModel) SetContextLengthNil(b bool)`

 SetContextLengthNil sets the value for ContextLength to be an explicit nil

### UnsetContextLength
`func (o *CatalogProviderModel) UnsetContextLength()`

UnsetContextLength ensures that no value is present for ContextLength, not even an explicit nil
### GetId

`func (o *CatalogProviderModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogProviderModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogProviderModel) SetId(v string)`

SetId sets Id field to given value.


### GetInputPrice

`func (o *CatalogProviderModel) GetInputPrice() float32`

GetInputPrice returns the InputPrice field if non-nil, zero value otherwise.

### GetInputPriceOk

`func (o *CatalogProviderModel) GetInputPriceOk() (*float32, bool)`

GetInputPriceOk returns a tuple with the InputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPrice

`func (o *CatalogProviderModel) SetInputPrice(v float32)`

SetInputPrice sets InputPrice field to given value.

### HasInputPrice

`func (o *CatalogProviderModel) HasInputPrice() bool`

HasInputPrice returns a boolean if a field has been set.

### SetInputPriceNil

`func (o *CatalogProviderModel) SetInputPriceNil(b bool)`

 SetInputPriceNil sets the value for InputPrice to be an explicit nil

### UnsetInputPrice
`func (o *CatalogProviderModel) UnsetInputPrice()`

UnsetInputPrice ensures that no value is present for InputPrice, not even an explicit nil
### GetName

`func (o *CatalogProviderModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogProviderModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogProviderModel) SetName(v string)`

SetName sets Name field to given value.


### GetOutputPrice

`func (o *CatalogProviderModel) GetOutputPrice() float32`

GetOutputPrice returns the OutputPrice field if non-nil, zero value otherwise.

### GetOutputPriceOk

`func (o *CatalogProviderModel) GetOutputPriceOk() (*float32, bool)`

GetOutputPriceOk returns a tuple with the OutputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPrice

`func (o *CatalogProviderModel) SetOutputPrice(v float32)`

SetOutputPrice sets OutputPrice field to given value.

### HasOutputPrice

`func (o *CatalogProviderModel) HasOutputPrice() bool`

HasOutputPrice returns a boolean if a field has been set.

### SetOutputPriceNil

`func (o *CatalogProviderModel) SetOutputPriceNil(b bool)`

 SetOutputPriceNil sets the value for OutputPrice to be an explicit nil

### UnsetOutputPrice
`func (o *CatalogProviderModel) UnsetOutputPrice()`

UnsetOutputPrice ensures that no value is present for OutputPrice, not even an explicit nil
### GetPriceUnit

`func (o *CatalogProviderModel) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *CatalogProviderModel) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *CatalogProviderModel) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *CatalogProviderModel) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *CatalogProviderModel) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *CatalogProviderModel) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetVendor

`func (o *CatalogProviderModel) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CatalogProviderModel) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CatalogProviderModel) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### SetVendorNil

`func (o *CatalogProviderModel) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *CatalogProviderModel) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


