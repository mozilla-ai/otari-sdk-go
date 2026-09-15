# CatalogModelDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlsoAvailableFrom** | [**[]CatalogElsewhere**](CatalogElsewhere.md) |  | 
**Capabilities** | [**CatalogCapabilities**](CatalogCapabilities.md) |  | 
**ContextWindow** | Pointer to **NullableInt32** | The largest any offering serves. | [optional] 
**DefaultPricing** | **bool** | Whether an unpriced model is metered at the genai-prices default. | 
**Deprecated** | Pointer to **bool** | True only when every offering with metadata says so. | [optional] [default to false]
**Description** | Pointer to **NullableString** | models.dev&#39;s, from the offering that named the model. | [optional] 
**Discovered** | **bool** | Whether any offering was discovered from its provider. | 
**Family** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | The catalog id, vendor-qualified where the vendor is known: &#x60;z-ai/glm-5.3&#x60;, else the bare slug. | 
**InputModalities** | **[]string** |  | 
**KnowledgeCutoff** | Pointer to **NullableString** |  | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** | The largest any offering serves. | [optional] 
**MinInputPricePerMillion** | Pointer to **NullableFloat32** | The cheapest offering&#39;s, at the comparison context where one was asked for. | [optional] 
**MinOutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Name** | **string** |  | 
**OfferingCount** | **int32** |  | 
**Offerings** | [**[]CatalogOffering**](CatalogOffering.md) |  | 
**OpenWeights** | Pointer to **bool** |  | [optional] [default to false]
**OutputModalities** | **[]string** |  | 
**PriceSources** | **[]string** | Which price lists the priced offerings came from, distinct and sorted. | 
**ProviderCount** | **int32** |  | 
**Providers** | **[]string** | The provider instances offering it, sorted. | 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**ResolvesTo** | Pointer to **NullableString** | The offering &#x60;selector&#x60; resolves to. | [optional] 
**Selector** | Pointer to **NullableString** | The id as a selector: send it as &#x60;model&#x60; and the model&#39;s cheapest offering answers. Null until the gateway has indexed the catalog. | [optional] 
**Selectors** | **[]string** | Every offering&#39;s selector, so the list can be searched by one. | 
**UnpricedCount** | **int32** | How many offerings carry no price for this caller. | 
**Vendor** | **NullableString** |  | 

## Methods

### NewCatalogModelDetail

`func NewCatalogModelDetail(alsoAvailableFrom []CatalogElsewhere, capabilities CatalogCapabilities, defaultPricing bool, discovered bool, id string, inputModalities []string, name string, offeringCount int32, offerings []CatalogOffering, outputModalities []string, priceSources []string, providerCount int32, providers []string, selectors []string, unpricedCount int32, vendor NullableString, ) *CatalogModelDetail`

NewCatalogModelDetail instantiates a new CatalogModelDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogModelDetailWithDefaults

`func NewCatalogModelDetailWithDefaults() *CatalogModelDetail`

NewCatalogModelDetailWithDefaults instantiates a new CatalogModelDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlsoAvailableFrom

`func (o *CatalogModelDetail) GetAlsoAvailableFrom() []CatalogElsewhere`

GetAlsoAvailableFrom returns the AlsoAvailableFrom field if non-nil, zero value otherwise.

### GetAlsoAvailableFromOk

`func (o *CatalogModelDetail) GetAlsoAvailableFromOk() (*[]CatalogElsewhere, bool)`

GetAlsoAvailableFromOk returns a tuple with the AlsoAvailableFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoAvailableFrom

`func (o *CatalogModelDetail) SetAlsoAvailableFrom(v []CatalogElsewhere)`

SetAlsoAvailableFrom sets AlsoAvailableFrom field to given value.


### GetCapabilities

`func (o *CatalogModelDetail) GetCapabilities() CatalogCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CatalogModelDetail) GetCapabilitiesOk() (*CatalogCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CatalogModelDetail) SetCapabilities(v CatalogCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetContextWindow

`func (o *CatalogModelDetail) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *CatalogModelDetail) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *CatalogModelDetail) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *CatalogModelDetail) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *CatalogModelDetail) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *CatalogModelDetail) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetDefaultPricing

`func (o *CatalogModelDetail) GetDefaultPricing() bool`

GetDefaultPricing returns the DefaultPricing field if non-nil, zero value otherwise.

### GetDefaultPricingOk

`func (o *CatalogModelDetail) GetDefaultPricingOk() (*bool, bool)`

GetDefaultPricingOk returns a tuple with the DefaultPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPricing

`func (o *CatalogModelDetail) SetDefaultPricing(v bool)`

SetDefaultPricing sets DefaultPricing field to given value.


### GetDeprecated

`func (o *CatalogModelDetail) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CatalogModelDetail) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CatalogModelDetail) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *CatalogModelDetail) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogModelDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogModelDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogModelDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogModelDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogModelDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogModelDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiscovered

`func (o *CatalogModelDetail) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *CatalogModelDetail) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *CatalogModelDetail) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.


### GetFamily

`func (o *CatalogModelDetail) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *CatalogModelDetail) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *CatalogModelDetail) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *CatalogModelDetail) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *CatalogModelDetail) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *CatalogModelDetail) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetId

`func (o *CatalogModelDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogModelDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogModelDetail) SetId(v string)`

SetId sets Id field to given value.


### GetInputModalities

`func (o *CatalogModelDetail) GetInputModalities() []string`

GetInputModalities returns the InputModalities field if non-nil, zero value otherwise.

### GetInputModalitiesOk

`func (o *CatalogModelDetail) GetInputModalitiesOk() (*[]string, bool)`

GetInputModalitiesOk returns a tuple with the InputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputModalities

`func (o *CatalogModelDetail) SetInputModalities(v []string)`

SetInputModalities sets InputModalities field to given value.


### GetKnowledgeCutoff

`func (o *CatalogModelDetail) GetKnowledgeCutoff() string`

GetKnowledgeCutoff returns the KnowledgeCutoff field if non-nil, zero value otherwise.

### GetKnowledgeCutoffOk

`func (o *CatalogModelDetail) GetKnowledgeCutoffOk() (*string, bool)`

GetKnowledgeCutoffOk returns a tuple with the KnowledgeCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCutoff

`func (o *CatalogModelDetail) SetKnowledgeCutoff(v string)`

SetKnowledgeCutoff sets KnowledgeCutoff field to given value.

### HasKnowledgeCutoff

`func (o *CatalogModelDetail) HasKnowledgeCutoff() bool`

HasKnowledgeCutoff returns a boolean if a field has been set.

### SetKnowledgeCutoffNil

`func (o *CatalogModelDetail) SetKnowledgeCutoffNil(b bool)`

 SetKnowledgeCutoffNil sets the value for KnowledgeCutoff to be an explicit nil

### UnsetKnowledgeCutoff
`func (o *CatalogModelDetail) UnsetKnowledgeCutoff()`

UnsetKnowledgeCutoff ensures that no value is present for KnowledgeCutoff, not even an explicit nil
### GetMaxOutputTokens

`func (o *CatalogModelDetail) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *CatalogModelDetail) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *CatalogModelDetail) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *CatalogModelDetail) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *CatalogModelDetail) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *CatalogModelDetail) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetMinInputPricePerMillion

`func (o *CatalogModelDetail) GetMinInputPricePerMillion() float32`

GetMinInputPricePerMillion returns the MinInputPricePerMillion field if non-nil, zero value otherwise.

### GetMinInputPricePerMillionOk

`func (o *CatalogModelDetail) GetMinInputPricePerMillionOk() (*float32, bool)`

GetMinInputPricePerMillionOk returns a tuple with the MinInputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInputPricePerMillion

`func (o *CatalogModelDetail) SetMinInputPricePerMillion(v float32)`

SetMinInputPricePerMillion sets MinInputPricePerMillion field to given value.

### HasMinInputPricePerMillion

`func (o *CatalogModelDetail) HasMinInputPricePerMillion() bool`

HasMinInputPricePerMillion returns a boolean if a field has been set.

### SetMinInputPricePerMillionNil

`func (o *CatalogModelDetail) SetMinInputPricePerMillionNil(b bool)`

 SetMinInputPricePerMillionNil sets the value for MinInputPricePerMillion to be an explicit nil

### UnsetMinInputPricePerMillion
`func (o *CatalogModelDetail) UnsetMinInputPricePerMillion()`

UnsetMinInputPricePerMillion ensures that no value is present for MinInputPricePerMillion, not even an explicit nil
### GetMinOutputPricePerMillion

`func (o *CatalogModelDetail) GetMinOutputPricePerMillion() float32`

GetMinOutputPricePerMillion returns the MinOutputPricePerMillion field if non-nil, zero value otherwise.

### GetMinOutputPricePerMillionOk

`func (o *CatalogModelDetail) GetMinOutputPricePerMillionOk() (*float32, bool)`

GetMinOutputPricePerMillionOk returns a tuple with the MinOutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOutputPricePerMillion

`func (o *CatalogModelDetail) SetMinOutputPricePerMillion(v float32)`

SetMinOutputPricePerMillion sets MinOutputPricePerMillion field to given value.

### HasMinOutputPricePerMillion

`func (o *CatalogModelDetail) HasMinOutputPricePerMillion() bool`

HasMinOutputPricePerMillion returns a boolean if a field has been set.

### SetMinOutputPricePerMillionNil

`func (o *CatalogModelDetail) SetMinOutputPricePerMillionNil(b bool)`

 SetMinOutputPricePerMillionNil sets the value for MinOutputPricePerMillion to be an explicit nil

### UnsetMinOutputPricePerMillion
`func (o *CatalogModelDetail) UnsetMinOutputPricePerMillion()`

UnsetMinOutputPricePerMillion ensures that no value is present for MinOutputPricePerMillion, not even an explicit nil
### GetName

`func (o *CatalogModelDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogModelDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogModelDetail) SetName(v string)`

SetName sets Name field to given value.


### GetOfferingCount

`func (o *CatalogModelDetail) GetOfferingCount() int32`

GetOfferingCount returns the OfferingCount field if non-nil, zero value otherwise.

### GetOfferingCountOk

`func (o *CatalogModelDetail) GetOfferingCountOk() (*int32, bool)`

GetOfferingCountOk returns a tuple with the OfferingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingCount

`func (o *CatalogModelDetail) SetOfferingCount(v int32)`

SetOfferingCount sets OfferingCount field to given value.


### GetOfferings

`func (o *CatalogModelDetail) GetOfferings() []CatalogOffering`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *CatalogModelDetail) GetOfferingsOk() (*[]CatalogOffering, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *CatalogModelDetail) SetOfferings(v []CatalogOffering)`

SetOfferings sets Offerings field to given value.


### GetOpenWeights

`func (o *CatalogModelDetail) GetOpenWeights() bool`

GetOpenWeights returns the OpenWeights field if non-nil, zero value otherwise.

### GetOpenWeightsOk

`func (o *CatalogModelDetail) GetOpenWeightsOk() (*bool, bool)`

GetOpenWeightsOk returns a tuple with the OpenWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWeights

`func (o *CatalogModelDetail) SetOpenWeights(v bool)`

SetOpenWeights sets OpenWeights field to given value.

### HasOpenWeights

`func (o *CatalogModelDetail) HasOpenWeights() bool`

HasOpenWeights returns a boolean if a field has been set.

### GetOutputModalities

`func (o *CatalogModelDetail) GetOutputModalities() []string`

GetOutputModalities returns the OutputModalities field if non-nil, zero value otherwise.

### GetOutputModalitiesOk

`func (o *CatalogModelDetail) GetOutputModalitiesOk() (*[]string, bool)`

GetOutputModalitiesOk returns a tuple with the OutputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputModalities

`func (o *CatalogModelDetail) SetOutputModalities(v []string)`

SetOutputModalities sets OutputModalities field to given value.


### GetPriceSources

`func (o *CatalogModelDetail) GetPriceSources() []string`

GetPriceSources returns the PriceSources field if non-nil, zero value otherwise.

### GetPriceSourcesOk

`func (o *CatalogModelDetail) GetPriceSourcesOk() (*[]string, bool)`

GetPriceSourcesOk returns a tuple with the PriceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSources

`func (o *CatalogModelDetail) SetPriceSources(v []string)`

SetPriceSources sets PriceSources field to given value.


### GetProviderCount

`func (o *CatalogModelDetail) GetProviderCount() int32`

GetProviderCount returns the ProviderCount field if non-nil, zero value otherwise.

### GetProviderCountOk

`func (o *CatalogModelDetail) GetProviderCountOk() (*int32, bool)`

GetProviderCountOk returns a tuple with the ProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCount

`func (o *CatalogModelDetail) SetProviderCount(v int32)`

SetProviderCount sets ProviderCount field to given value.


### GetProviders

`func (o *CatalogModelDetail) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CatalogModelDetail) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CatalogModelDetail) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetReleaseDate

`func (o *CatalogModelDetail) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CatalogModelDetail) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CatalogModelDetail) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *CatalogModelDetail) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *CatalogModelDetail) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *CatalogModelDetail) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetResolvesTo

`func (o *CatalogModelDetail) GetResolvesTo() string`

GetResolvesTo returns the ResolvesTo field if non-nil, zero value otherwise.

### GetResolvesToOk

`func (o *CatalogModelDetail) GetResolvesToOk() (*string, bool)`

GetResolvesToOk returns a tuple with the ResolvesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvesTo

`func (o *CatalogModelDetail) SetResolvesTo(v string)`

SetResolvesTo sets ResolvesTo field to given value.

### HasResolvesTo

`func (o *CatalogModelDetail) HasResolvesTo() bool`

HasResolvesTo returns a boolean if a field has been set.

### SetResolvesToNil

`func (o *CatalogModelDetail) SetResolvesToNil(b bool)`

 SetResolvesToNil sets the value for ResolvesTo to be an explicit nil

### UnsetResolvesTo
`func (o *CatalogModelDetail) UnsetResolvesTo()`

UnsetResolvesTo ensures that no value is present for ResolvesTo, not even an explicit nil
### GetSelector

`func (o *CatalogModelDetail) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CatalogModelDetail) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CatalogModelDetail) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *CatalogModelDetail) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *CatalogModelDetail) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *CatalogModelDetail) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetSelectors

`func (o *CatalogModelDetail) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *CatalogModelDetail) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *CatalogModelDetail) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetUnpricedCount

`func (o *CatalogModelDetail) GetUnpricedCount() int32`

GetUnpricedCount returns the UnpricedCount field if non-nil, zero value otherwise.

### GetUnpricedCountOk

`func (o *CatalogModelDetail) GetUnpricedCountOk() (*int32, bool)`

GetUnpricedCountOk returns a tuple with the UnpricedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpricedCount

`func (o *CatalogModelDetail) SetUnpricedCount(v int32)`

SetUnpricedCount sets UnpricedCount field to given value.


### GetVendor

`func (o *CatalogModelDetail) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CatalogModelDetail) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CatalogModelDetail) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### SetVendorNil

`func (o *CatalogModelDetail) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *CatalogModelDetail) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


