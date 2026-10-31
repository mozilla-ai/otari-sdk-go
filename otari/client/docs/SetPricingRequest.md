# SetPricingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cached-input tokens | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** | Price per 1M cache-write (creation) tokens | [optional] 
**EffectiveAt** | Pointer to **NullableTime** | ISO 8601 datetime from which this price applies. Defaults to now if omitted. | [optional] 
**InputPricePerMillion** | Pointer to **float32** | Price per 1M of the rate&#39;s unit: input tokens, requests or images | [optional] 
**Kind** | Pointer to **NullableString** | What the row prices: &#39;model&#39; (the default), or &#39;search&#39; or &#39;fetch&#39; for a search or fetch instance, keyed by the instance&#39;s name and priced per call. Only the deployment operator sets a search or fetch rate. | [optional] 
**ModelKey** | **string** | Model identifier in format &#39;provider:model&#39;, or a search or fetch instance&#39;s name for those kinds | 
**OutputPricePerMillion** | Pointer to **float32** | Price per 1M output tokens; unused by a rate in requests or images | [optional] 
**PricePerRequest** | Pointer to **NullableFloat32** | USD for each request, whatever it uses: a completion model billed per call, such as Exa&#39;s or Liner&#39;s answers. Shorthand for unit &#39;requests&#39; with input_price_per_million at a million times this; not combined with token prices. | [optional] 
**PricingTiers** | Pointer to [**[]PricingTier**](PricingTier.md) | Whole-request context thresholds. Fields omitted by a tier inherit the base rate. | [optional] 
**Unit** | Pointer to **string** | What the rates are per: &#39;tokens&#39; for a model, &#39;requests&#39; for a gateway-run tool, a moderation call, or a completion model billed per call (USD per million requests), &#39;images&#39; for image generation (USD per million images). | [optional] [default to "tokens"]

## Methods

### NewSetPricingRequest

`func NewSetPricingRequest(modelKey string, ) *SetPricingRequest`

NewSetPricingRequest instantiates a new SetPricingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPricingRequestWithDefaults

`func NewSetPricingRequestWithDefaults() *SetPricingRequest`

NewSetPricingRequestWithDefaults instantiates a new SetPricingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *SetPricingRequest) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *SetPricingRequest) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *SetPricingRequest) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *SetPricingRequest) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *SetPricingRequest) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *SetPricingRequest) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *SetPricingRequest) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *SetPricingRequest) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *SetPricingRequest) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *SetPricingRequest) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *SetPricingRequest) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *SetPricingRequest) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetEffectiveAt

`func (o *SetPricingRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *SetPricingRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *SetPricingRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *SetPricingRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### SetEffectiveAtNil

`func (o *SetPricingRequest) SetEffectiveAtNil(b bool)`

 SetEffectiveAtNil sets the value for EffectiveAt to be an explicit nil

### UnsetEffectiveAt
`func (o *SetPricingRequest) UnsetEffectiveAt()`

UnsetEffectiveAt ensures that no value is present for EffectiveAt, not even an explicit nil
### GetInputPricePerMillion

`func (o *SetPricingRequest) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *SetPricingRequest) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *SetPricingRequest) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *SetPricingRequest) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### GetKind

`func (o *SetPricingRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SetPricingRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SetPricingRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SetPricingRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *SetPricingRequest) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *SetPricingRequest) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetModelKey

`func (o *SetPricingRequest) GetModelKey() string`

GetModelKey returns the ModelKey field if non-nil, zero value otherwise.

### GetModelKeyOk

`func (o *SetPricingRequest) GetModelKeyOk() (*string, bool)`

GetModelKeyOk returns a tuple with the ModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelKey

`func (o *SetPricingRequest) SetModelKey(v string)`

SetModelKey sets ModelKey field to given value.


### GetOutputPricePerMillion

`func (o *SetPricingRequest) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *SetPricingRequest) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *SetPricingRequest) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *SetPricingRequest) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### GetPricePerRequest

`func (o *SetPricingRequest) GetPricePerRequest() float32`

GetPricePerRequest returns the PricePerRequest field if non-nil, zero value otherwise.

### GetPricePerRequestOk

`func (o *SetPricingRequest) GetPricePerRequestOk() (*float32, bool)`

GetPricePerRequestOk returns a tuple with the PricePerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerRequest

`func (o *SetPricingRequest) SetPricePerRequest(v float32)`

SetPricePerRequest sets PricePerRequest field to given value.

### HasPricePerRequest

`func (o *SetPricingRequest) HasPricePerRequest() bool`

HasPricePerRequest returns a boolean if a field has been set.

### SetPricePerRequestNil

`func (o *SetPricingRequest) SetPricePerRequestNil(b bool)`

 SetPricePerRequestNil sets the value for PricePerRequest to be an explicit nil

### UnsetPricePerRequest
`func (o *SetPricingRequest) UnsetPricePerRequest()`

UnsetPricePerRequest ensures that no value is present for PricePerRequest, not even an explicit nil
### GetPricingTiers

`func (o *SetPricingRequest) GetPricingTiers() []PricingTier`

GetPricingTiers returns the PricingTiers field if non-nil, zero value otherwise.

### GetPricingTiersOk

`func (o *SetPricingRequest) GetPricingTiersOk() (*[]PricingTier, bool)`

GetPricingTiersOk returns a tuple with the PricingTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTiers

`func (o *SetPricingRequest) SetPricingTiers(v []PricingTier)`

SetPricingTiers sets PricingTiers field to given value.

### HasPricingTiers

`func (o *SetPricingRequest) HasPricingTiers() bool`

HasPricingTiers returns a boolean if a field has been set.

### SetPricingTiersNil

`func (o *SetPricingRequest) SetPricingTiersNil(b bool)`

 SetPricingTiersNil sets the value for PricingTiers to be an explicit nil

### UnsetPricingTiers
`func (o *SetPricingRequest) UnsetPricingTiers()`

UnsetPricingTiers ensures that no value is present for PricingTiers, not even an explicit nil
### GetUnit

`func (o *SetPricingRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *SetPricingRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *SetPricingRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *SetPricingRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


