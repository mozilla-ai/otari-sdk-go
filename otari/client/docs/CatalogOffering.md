# CatalogOffering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextWindow** | Pointer to **NullableInt32** |  | [optional] 
**Credential** | [**CatalogCredential**](CatalogCredential.md) | Who may price it: &#x60;deployment&#x60; for a &#x60;providers:&#x60; instance the operator configured, &#x60;hosted&#x60; for a provider the deployment pays for in any workspace of the viewer&#39;s organization, &#x60;organization&#x60; for one the viewer&#39;s organization may set its own rate for. A workspace can still call a &#x60;hosted&#x60; provider with the organization&#39;s own key. | 
**Discovered** | **bool** | Whether the provider itself reported this model. | 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**MetadataInputPricePerMillion** | Pointer to **NullableFloat32** | What models.dev lists this provider charging, for a cross-check. Not billed from: two independent datasets disagreeing is the cheapest stale-price detector there is. | [optional] 
**MetadataOutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**PriceReference** | Pointer to **NullableString** | For a default, the genai-prices &#x60;provider:model&#x60; entry that matched; the selector otherwise. | [optional] 
**PriceSource** | Pointer to **NullableString** | Which price list &#x60;pricing&#x60; came from, for this viewer: the organization&#39;s own override, the deployment&#39;s stored row, or the genai-prices defaults. Null when nothing prices it. | [optional] 
**Pricing** | Pointer to [**NullableModelPricingInfo**](ModelPricingInfo.md) |  | [optional] 
**Provider** | **string** | The provider instance the selector names. | 
**ProviderType** | **string** | The any-llm implementation behind the instance. | 
**Quantization** | Pointer to **NullableString** | From the provider&#39;s id, when it names one. | [optional] 
**Selector** | **string** | What to send as &#x60;model&#x60;, in &#x60;instance:model&#x60; form. | 
**ShortSelector** | Pointer to **NullableString** | The pinned spelling the gateway also accepts for this offering: the instance with the model&#39;s catalog id (&#x60;fireworks:openai/gpt-oss-120b&#x60;), which pins the instance and reaches the model&#39;s cheapest offering on it. Null for a dearer sibling on the same instance, or until the gateway has indexed the catalog. | [optional] 
**Usage30d** | Pointer to [**NullableOfferingUsage**](OfferingUsage.md) |  | [optional] 

## Methods

### NewCatalogOffering

`func NewCatalogOffering(credential CatalogCredential, discovered bool, provider string, providerType string, selector string, ) *CatalogOffering`

NewCatalogOffering instantiates a new CatalogOffering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOfferingWithDefaults

`func NewCatalogOfferingWithDefaults() *CatalogOffering`

NewCatalogOfferingWithDefaults instantiates a new CatalogOffering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextWindow

`func (o *CatalogOffering) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *CatalogOffering) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *CatalogOffering) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *CatalogOffering) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *CatalogOffering) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *CatalogOffering) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetCredential

`func (o *CatalogOffering) GetCredential() CatalogCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CatalogOffering) GetCredentialOk() (*CatalogCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CatalogOffering) SetCredential(v CatalogCredential)`

SetCredential sets Credential field to given value.


### GetDiscovered

`func (o *CatalogOffering) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *CatalogOffering) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *CatalogOffering) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.


### GetMaxOutputTokens

`func (o *CatalogOffering) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *CatalogOffering) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *CatalogOffering) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *CatalogOffering) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *CatalogOffering) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *CatalogOffering) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetMetadataInputPricePerMillion

`func (o *CatalogOffering) GetMetadataInputPricePerMillion() float32`

GetMetadataInputPricePerMillion returns the MetadataInputPricePerMillion field if non-nil, zero value otherwise.

### GetMetadataInputPricePerMillionOk

`func (o *CatalogOffering) GetMetadataInputPricePerMillionOk() (*float32, bool)`

GetMetadataInputPricePerMillionOk returns a tuple with the MetadataInputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInputPricePerMillion

`func (o *CatalogOffering) SetMetadataInputPricePerMillion(v float32)`

SetMetadataInputPricePerMillion sets MetadataInputPricePerMillion field to given value.

### HasMetadataInputPricePerMillion

`func (o *CatalogOffering) HasMetadataInputPricePerMillion() bool`

HasMetadataInputPricePerMillion returns a boolean if a field has been set.

### SetMetadataInputPricePerMillionNil

`func (o *CatalogOffering) SetMetadataInputPricePerMillionNil(b bool)`

 SetMetadataInputPricePerMillionNil sets the value for MetadataInputPricePerMillion to be an explicit nil

### UnsetMetadataInputPricePerMillion
`func (o *CatalogOffering) UnsetMetadataInputPricePerMillion()`

UnsetMetadataInputPricePerMillion ensures that no value is present for MetadataInputPricePerMillion, not even an explicit nil
### GetMetadataOutputPricePerMillion

`func (o *CatalogOffering) GetMetadataOutputPricePerMillion() float32`

GetMetadataOutputPricePerMillion returns the MetadataOutputPricePerMillion field if non-nil, zero value otherwise.

### GetMetadataOutputPricePerMillionOk

`func (o *CatalogOffering) GetMetadataOutputPricePerMillionOk() (*float32, bool)`

GetMetadataOutputPricePerMillionOk returns a tuple with the MetadataOutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataOutputPricePerMillion

`func (o *CatalogOffering) SetMetadataOutputPricePerMillion(v float32)`

SetMetadataOutputPricePerMillion sets MetadataOutputPricePerMillion field to given value.

### HasMetadataOutputPricePerMillion

`func (o *CatalogOffering) HasMetadataOutputPricePerMillion() bool`

HasMetadataOutputPricePerMillion returns a boolean if a field has been set.

### SetMetadataOutputPricePerMillionNil

`func (o *CatalogOffering) SetMetadataOutputPricePerMillionNil(b bool)`

 SetMetadataOutputPricePerMillionNil sets the value for MetadataOutputPricePerMillion to be an explicit nil

### UnsetMetadataOutputPricePerMillion
`func (o *CatalogOffering) UnsetMetadataOutputPricePerMillion()`

UnsetMetadataOutputPricePerMillion ensures that no value is present for MetadataOutputPricePerMillion, not even an explicit nil
### GetPriceReference

`func (o *CatalogOffering) GetPriceReference() string`

GetPriceReference returns the PriceReference field if non-nil, zero value otherwise.

### GetPriceReferenceOk

`func (o *CatalogOffering) GetPriceReferenceOk() (*string, bool)`

GetPriceReferenceOk returns a tuple with the PriceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceReference

`func (o *CatalogOffering) SetPriceReference(v string)`

SetPriceReference sets PriceReference field to given value.

### HasPriceReference

`func (o *CatalogOffering) HasPriceReference() bool`

HasPriceReference returns a boolean if a field has been set.

### SetPriceReferenceNil

`func (o *CatalogOffering) SetPriceReferenceNil(b bool)`

 SetPriceReferenceNil sets the value for PriceReference to be an explicit nil

### UnsetPriceReference
`func (o *CatalogOffering) UnsetPriceReference()`

UnsetPriceReference ensures that no value is present for PriceReference, not even an explicit nil
### GetPriceSource

`func (o *CatalogOffering) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *CatalogOffering) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *CatalogOffering) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.

### HasPriceSource

`func (o *CatalogOffering) HasPriceSource() bool`

HasPriceSource returns a boolean if a field has been set.

### SetPriceSourceNil

`func (o *CatalogOffering) SetPriceSourceNil(b bool)`

 SetPriceSourceNil sets the value for PriceSource to be an explicit nil

### UnsetPriceSource
`func (o *CatalogOffering) UnsetPriceSource()`

UnsetPriceSource ensures that no value is present for PriceSource, not even an explicit nil
### GetPricing

`func (o *CatalogOffering) GetPricing() ModelPricingInfo`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CatalogOffering) GetPricingOk() (*ModelPricingInfo, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CatalogOffering) SetPricing(v ModelPricingInfo)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *CatalogOffering) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *CatalogOffering) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *CatalogOffering) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetProvider

`func (o *CatalogOffering) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CatalogOffering) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CatalogOffering) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderType

`func (o *CatalogOffering) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CatalogOffering) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CatalogOffering) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetQuantization

`func (o *CatalogOffering) GetQuantization() string`

GetQuantization returns the Quantization field if non-nil, zero value otherwise.

### GetQuantizationOk

`func (o *CatalogOffering) GetQuantizationOk() (*string, bool)`

GetQuantizationOk returns a tuple with the Quantization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantization

`func (o *CatalogOffering) SetQuantization(v string)`

SetQuantization sets Quantization field to given value.

### HasQuantization

`func (o *CatalogOffering) HasQuantization() bool`

HasQuantization returns a boolean if a field has been set.

### SetQuantizationNil

`func (o *CatalogOffering) SetQuantizationNil(b bool)`

 SetQuantizationNil sets the value for Quantization to be an explicit nil

### UnsetQuantization
`func (o *CatalogOffering) UnsetQuantization()`

UnsetQuantization ensures that no value is present for Quantization, not even an explicit nil
### GetSelector

`func (o *CatalogOffering) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CatalogOffering) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CatalogOffering) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetShortSelector

`func (o *CatalogOffering) GetShortSelector() string`

GetShortSelector returns the ShortSelector field if non-nil, zero value otherwise.

### GetShortSelectorOk

`func (o *CatalogOffering) GetShortSelectorOk() (*string, bool)`

GetShortSelectorOk returns a tuple with the ShortSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortSelector

`func (o *CatalogOffering) SetShortSelector(v string)`

SetShortSelector sets ShortSelector field to given value.

### HasShortSelector

`func (o *CatalogOffering) HasShortSelector() bool`

HasShortSelector returns a boolean if a field has been set.

### SetShortSelectorNil

`func (o *CatalogOffering) SetShortSelectorNil(b bool)`

 SetShortSelectorNil sets the value for ShortSelector to be an explicit nil

### UnsetShortSelector
`func (o *CatalogOffering) UnsetShortSelector()`

UnsetShortSelector ensures that no value is present for ShortSelector, not even an explicit nil
### GetUsage30d

`func (o *CatalogOffering) GetUsage30d() OfferingUsage`

GetUsage30d returns the Usage30d field if non-nil, zero value otherwise.

### GetUsage30dOk

`func (o *CatalogOffering) GetUsage30dOk() (*OfferingUsage, bool)`

GetUsage30dOk returns a tuple with the Usage30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage30d

`func (o *CatalogOffering) SetUsage30d(v OfferingUsage)`

SetUsage30d sets Usage30d field to given value.

### HasUsage30d

`func (o *CatalogOffering) HasUsage30d() bool`

HasUsage30d returns a boolean if a field has been set.

### SetUsage30dNil

`func (o *CatalogOffering) SetUsage30dNil(b bool)`

 SetUsage30dNil sets the value for Usage30d to be an explicit nil

### UnsetUsage30d
`func (o *CatalogOffering) UnsetUsage30d()`

UnsetUsage30d ensures that no value is present for Usage30d, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


