# CatalogModelSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**CatalogCapabilities**](CatalogCapabilities.md) |  | 
**ContextWindow** | Pointer to **NullableInt32** | The largest any offering serves. | [optional] 
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

### NewCatalogModelSummary

`func NewCatalogModelSummary(capabilities CatalogCapabilities, discovered bool, id string, inputModalities []string, name string, offeringCount int32, outputModalities []string, priceSources []string, providerCount int32, providers []string, selectors []string, unpricedCount int32, vendor NullableString, ) *CatalogModelSummary`

NewCatalogModelSummary instantiates a new CatalogModelSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogModelSummaryWithDefaults

`func NewCatalogModelSummaryWithDefaults() *CatalogModelSummary`

NewCatalogModelSummaryWithDefaults instantiates a new CatalogModelSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CatalogModelSummary) GetCapabilities() CatalogCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CatalogModelSummary) GetCapabilitiesOk() (*CatalogCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CatalogModelSummary) SetCapabilities(v CatalogCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetContextWindow

`func (o *CatalogModelSummary) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *CatalogModelSummary) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *CatalogModelSummary) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *CatalogModelSummary) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *CatalogModelSummary) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *CatalogModelSummary) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetDeprecated

`func (o *CatalogModelSummary) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CatalogModelSummary) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CatalogModelSummary) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *CatalogModelSummary) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogModelSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogModelSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogModelSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogModelSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogModelSummary) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogModelSummary) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiscovered

`func (o *CatalogModelSummary) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *CatalogModelSummary) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *CatalogModelSummary) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.


### GetFamily

`func (o *CatalogModelSummary) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *CatalogModelSummary) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *CatalogModelSummary) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *CatalogModelSummary) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *CatalogModelSummary) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *CatalogModelSummary) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetId

`func (o *CatalogModelSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogModelSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogModelSummary) SetId(v string)`

SetId sets Id field to given value.


### GetInputModalities

`func (o *CatalogModelSummary) GetInputModalities() []string`

GetInputModalities returns the InputModalities field if non-nil, zero value otherwise.

### GetInputModalitiesOk

`func (o *CatalogModelSummary) GetInputModalitiesOk() (*[]string, bool)`

GetInputModalitiesOk returns a tuple with the InputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputModalities

`func (o *CatalogModelSummary) SetInputModalities(v []string)`

SetInputModalities sets InputModalities field to given value.


### GetKnowledgeCutoff

`func (o *CatalogModelSummary) GetKnowledgeCutoff() string`

GetKnowledgeCutoff returns the KnowledgeCutoff field if non-nil, zero value otherwise.

### GetKnowledgeCutoffOk

`func (o *CatalogModelSummary) GetKnowledgeCutoffOk() (*string, bool)`

GetKnowledgeCutoffOk returns a tuple with the KnowledgeCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCutoff

`func (o *CatalogModelSummary) SetKnowledgeCutoff(v string)`

SetKnowledgeCutoff sets KnowledgeCutoff field to given value.

### HasKnowledgeCutoff

`func (o *CatalogModelSummary) HasKnowledgeCutoff() bool`

HasKnowledgeCutoff returns a boolean if a field has been set.

### SetKnowledgeCutoffNil

`func (o *CatalogModelSummary) SetKnowledgeCutoffNil(b bool)`

 SetKnowledgeCutoffNil sets the value for KnowledgeCutoff to be an explicit nil

### UnsetKnowledgeCutoff
`func (o *CatalogModelSummary) UnsetKnowledgeCutoff()`

UnsetKnowledgeCutoff ensures that no value is present for KnowledgeCutoff, not even an explicit nil
### GetMaxOutputTokens

`func (o *CatalogModelSummary) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *CatalogModelSummary) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *CatalogModelSummary) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *CatalogModelSummary) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *CatalogModelSummary) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *CatalogModelSummary) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetMinInputPricePerMillion

`func (o *CatalogModelSummary) GetMinInputPricePerMillion() float32`

GetMinInputPricePerMillion returns the MinInputPricePerMillion field if non-nil, zero value otherwise.

### GetMinInputPricePerMillionOk

`func (o *CatalogModelSummary) GetMinInputPricePerMillionOk() (*float32, bool)`

GetMinInputPricePerMillionOk returns a tuple with the MinInputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInputPricePerMillion

`func (o *CatalogModelSummary) SetMinInputPricePerMillion(v float32)`

SetMinInputPricePerMillion sets MinInputPricePerMillion field to given value.

### HasMinInputPricePerMillion

`func (o *CatalogModelSummary) HasMinInputPricePerMillion() bool`

HasMinInputPricePerMillion returns a boolean if a field has been set.

### SetMinInputPricePerMillionNil

`func (o *CatalogModelSummary) SetMinInputPricePerMillionNil(b bool)`

 SetMinInputPricePerMillionNil sets the value for MinInputPricePerMillion to be an explicit nil

### UnsetMinInputPricePerMillion
`func (o *CatalogModelSummary) UnsetMinInputPricePerMillion()`

UnsetMinInputPricePerMillion ensures that no value is present for MinInputPricePerMillion, not even an explicit nil
### GetMinOutputPricePerMillion

`func (o *CatalogModelSummary) GetMinOutputPricePerMillion() float32`

GetMinOutputPricePerMillion returns the MinOutputPricePerMillion field if non-nil, zero value otherwise.

### GetMinOutputPricePerMillionOk

`func (o *CatalogModelSummary) GetMinOutputPricePerMillionOk() (*float32, bool)`

GetMinOutputPricePerMillionOk returns a tuple with the MinOutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOutputPricePerMillion

`func (o *CatalogModelSummary) SetMinOutputPricePerMillion(v float32)`

SetMinOutputPricePerMillion sets MinOutputPricePerMillion field to given value.

### HasMinOutputPricePerMillion

`func (o *CatalogModelSummary) HasMinOutputPricePerMillion() bool`

HasMinOutputPricePerMillion returns a boolean if a field has been set.

### SetMinOutputPricePerMillionNil

`func (o *CatalogModelSummary) SetMinOutputPricePerMillionNil(b bool)`

 SetMinOutputPricePerMillionNil sets the value for MinOutputPricePerMillion to be an explicit nil

### UnsetMinOutputPricePerMillion
`func (o *CatalogModelSummary) UnsetMinOutputPricePerMillion()`

UnsetMinOutputPricePerMillion ensures that no value is present for MinOutputPricePerMillion, not even an explicit nil
### GetName

`func (o *CatalogModelSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogModelSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogModelSummary) SetName(v string)`

SetName sets Name field to given value.


### GetOfferingCount

`func (o *CatalogModelSummary) GetOfferingCount() int32`

GetOfferingCount returns the OfferingCount field if non-nil, zero value otherwise.

### GetOfferingCountOk

`func (o *CatalogModelSummary) GetOfferingCountOk() (*int32, bool)`

GetOfferingCountOk returns a tuple with the OfferingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferingCount

`func (o *CatalogModelSummary) SetOfferingCount(v int32)`

SetOfferingCount sets OfferingCount field to given value.


### GetOpenWeights

`func (o *CatalogModelSummary) GetOpenWeights() bool`

GetOpenWeights returns the OpenWeights field if non-nil, zero value otherwise.

### GetOpenWeightsOk

`func (o *CatalogModelSummary) GetOpenWeightsOk() (*bool, bool)`

GetOpenWeightsOk returns a tuple with the OpenWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWeights

`func (o *CatalogModelSummary) SetOpenWeights(v bool)`

SetOpenWeights sets OpenWeights field to given value.

### HasOpenWeights

`func (o *CatalogModelSummary) HasOpenWeights() bool`

HasOpenWeights returns a boolean if a field has been set.

### GetOutputModalities

`func (o *CatalogModelSummary) GetOutputModalities() []string`

GetOutputModalities returns the OutputModalities field if non-nil, zero value otherwise.

### GetOutputModalitiesOk

`func (o *CatalogModelSummary) GetOutputModalitiesOk() (*[]string, bool)`

GetOutputModalitiesOk returns a tuple with the OutputModalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputModalities

`func (o *CatalogModelSummary) SetOutputModalities(v []string)`

SetOutputModalities sets OutputModalities field to given value.


### GetPriceSources

`func (o *CatalogModelSummary) GetPriceSources() []string`

GetPriceSources returns the PriceSources field if non-nil, zero value otherwise.

### GetPriceSourcesOk

`func (o *CatalogModelSummary) GetPriceSourcesOk() (*[]string, bool)`

GetPriceSourcesOk returns a tuple with the PriceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSources

`func (o *CatalogModelSummary) SetPriceSources(v []string)`

SetPriceSources sets PriceSources field to given value.


### GetProviderCount

`func (o *CatalogModelSummary) GetProviderCount() int32`

GetProviderCount returns the ProviderCount field if non-nil, zero value otherwise.

### GetProviderCountOk

`func (o *CatalogModelSummary) GetProviderCountOk() (*int32, bool)`

GetProviderCountOk returns a tuple with the ProviderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCount

`func (o *CatalogModelSummary) SetProviderCount(v int32)`

SetProviderCount sets ProviderCount field to given value.


### GetProviders

`func (o *CatalogModelSummary) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CatalogModelSummary) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CatalogModelSummary) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### GetReleaseDate

`func (o *CatalogModelSummary) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CatalogModelSummary) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CatalogModelSummary) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *CatalogModelSummary) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *CatalogModelSummary) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *CatalogModelSummary) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetResolvesTo

`func (o *CatalogModelSummary) GetResolvesTo() string`

GetResolvesTo returns the ResolvesTo field if non-nil, zero value otherwise.

### GetResolvesToOk

`func (o *CatalogModelSummary) GetResolvesToOk() (*string, bool)`

GetResolvesToOk returns a tuple with the ResolvesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvesTo

`func (o *CatalogModelSummary) SetResolvesTo(v string)`

SetResolvesTo sets ResolvesTo field to given value.

### HasResolvesTo

`func (o *CatalogModelSummary) HasResolvesTo() bool`

HasResolvesTo returns a boolean if a field has been set.

### SetResolvesToNil

`func (o *CatalogModelSummary) SetResolvesToNil(b bool)`

 SetResolvesToNil sets the value for ResolvesTo to be an explicit nil

### UnsetResolvesTo
`func (o *CatalogModelSummary) UnsetResolvesTo()`

UnsetResolvesTo ensures that no value is present for ResolvesTo, not even an explicit nil
### GetSelector

`func (o *CatalogModelSummary) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CatalogModelSummary) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CatalogModelSummary) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *CatalogModelSummary) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *CatalogModelSummary) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *CatalogModelSummary) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetSelectors

`func (o *CatalogModelSummary) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *CatalogModelSummary) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *CatalogModelSummary) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetUnpricedCount

`func (o *CatalogModelSummary) GetUnpricedCount() int32`

GetUnpricedCount returns the UnpricedCount field if non-nil, zero value otherwise.

### GetUnpricedCountOk

`func (o *CatalogModelSummary) GetUnpricedCountOk() (*int32, bool)`

GetUnpricedCountOk returns a tuple with the UnpricedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpricedCount

`func (o *CatalogModelSummary) SetUnpricedCount(v int32)`

SetUnpricedCount sets UnpricedCount field to given value.


### GetVendor

`func (o *CatalogModelSummary) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CatalogModelSummary) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CatalogModelSummary) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### SetVendorNil

`func (o *CatalogModelSummary) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *CatalogModelSummary) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


