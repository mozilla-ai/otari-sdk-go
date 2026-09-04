# OrganizationModelPricingPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | **NullableFloat32** |  | 
**CacheWrite1hPricePerMillion** | **NullableFloat32** |  | 
**CacheWritePricePerMillion** | **NullableFloat32** |  | 
**CreatedAt** | **time.Time** |  | 
**EffectiveFrom** | **time.Time** |  | 
**EffectiveTo** | **NullableTime** |  | 
**Id** | **string** |  | 
**InputPricePerMillion** | **float32** |  | 
**ModelKey** | **string** |  | 
**OrganizationId** | **string** |  | 
**OutputPricePerMillion** | **float32** |  | 
**PricingTiers** | [**[]PricingTier**](PricingTier.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewOrganizationModelPricingPublic

`func NewOrganizationModelPricingPublic(cacheReadPricePerMillion NullableFloat32, cacheWrite1hPricePerMillion NullableFloat32, cacheWritePricePerMillion NullableFloat32, createdAt time.Time, effectiveFrom time.Time, effectiveTo NullableTime, id string, inputPricePerMillion float32, modelKey string, organizationId string, outputPricePerMillion float32, pricingTiers []PricingTier, updatedAt time.Time, ) *OrganizationModelPricingPublic`

NewOrganizationModelPricingPublic instantiates a new OrganizationModelPricingPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationModelPricingPublicWithDefaults

`func NewOrganizationModelPricingPublicWithDefaults() *OrganizationModelPricingPublic`

NewOrganizationModelPricingPublicWithDefaults instantiates a new OrganizationModelPricingPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *OrganizationModelPricingPublic) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *OrganizationModelPricingPublic) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *OrganizationModelPricingPublic) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.


### SetCacheReadPricePerMillionNil

`func (o *OrganizationModelPricingPublic) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *OrganizationModelPricingPublic) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWrite1hPricePerMillion

`func (o *OrganizationModelPricingPublic) GetCacheWrite1hPricePerMillion() float32`

GetCacheWrite1hPricePerMillion returns the CacheWrite1hPricePerMillion field if non-nil, zero value otherwise.

### GetCacheWrite1hPricePerMillionOk

`func (o *OrganizationModelPricingPublic) GetCacheWrite1hPricePerMillionOk() (*float32, bool)`

GetCacheWrite1hPricePerMillionOk returns a tuple with the CacheWrite1hPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hPricePerMillion

`func (o *OrganizationModelPricingPublic) SetCacheWrite1hPricePerMillion(v float32)`

SetCacheWrite1hPricePerMillion sets CacheWrite1hPricePerMillion field to given value.


### SetCacheWrite1hPricePerMillionNil

`func (o *OrganizationModelPricingPublic) SetCacheWrite1hPricePerMillionNil(b bool)`

 SetCacheWrite1hPricePerMillionNil sets the value for CacheWrite1hPricePerMillion to be an explicit nil

### UnsetCacheWrite1hPricePerMillion
`func (o *OrganizationModelPricingPublic) UnsetCacheWrite1hPricePerMillion()`

UnsetCacheWrite1hPricePerMillion ensures that no value is present for CacheWrite1hPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *OrganizationModelPricingPublic) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *OrganizationModelPricingPublic) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *OrganizationModelPricingPublic) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.


### SetCacheWritePricePerMillionNil

`func (o *OrganizationModelPricingPublic) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *OrganizationModelPricingPublic) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetCreatedAt

`func (o *OrganizationModelPricingPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationModelPricingPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationModelPricingPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEffectiveFrom

`func (o *OrganizationModelPricingPublic) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *OrganizationModelPricingPublic) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *OrganizationModelPricingPublic) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *OrganizationModelPricingPublic) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *OrganizationModelPricingPublic) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *OrganizationModelPricingPublic) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.


### SetEffectiveToNil

`func (o *OrganizationModelPricingPublic) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *OrganizationModelPricingPublic) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil
### GetId

`func (o *OrganizationModelPricingPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationModelPricingPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationModelPricingPublic) SetId(v string)`

SetId sets Id field to given value.


### GetInputPricePerMillion

`func (o *OrganizationModelPricingPublic) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *OrganizationModelPricingPublic) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *OrganizationModelPricingPublic) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetModelKey

`func (o *OrganizationModelPricingPublic) GetModelKey() string`

GetModelKey returns the ModelKey field if non-nil, zero value otherwise.

### GetModelKeyOk

`func (o *OrganizationModelPricingPublic) GetModelKeyOk() (*string, bool)`

GetModelKeyOk returns a tuple with the ModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelKey

`func (o *OrganizationModelPricingPublic) SetModelKey(v string)`

SetModelKey sets ModelKey field to given value.


### GetOrganizationId

`func (o *OrganizationModelPricingPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationModelPricingPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationModelPricingPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetOutputPricePerMillion

`func (o *OrganizationModelPricingPublic) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *OrganizationModelPricingPublic) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *OrganizationModelPricingPublic) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.


### GetPricingTiers

`func (o *OrganizationModelPricingPublic) GetPricingTiers() []PricingTier`

GetPricingTiers returns the PricingTiers field if non-nil, zero value otherwise.

### GetPricingTiersOk

`func (o *OrganizationModelPricingPublic) GetPricingTiersOk() (*[]PricingTier, bool)`

GetPricingTiersOk returns a tuple with the PricingTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTiers

`func (o *OrganizationModelPricingPublic) SetPricingTiers(v []PricingTier)`

SetPricingTiers sets PricingTiers field to given value.


### GetUpdatedAt

`func (o *OrganizationModelPricingPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationModelPricingPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationModelPricingPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


