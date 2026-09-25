# OrgProviderKeyModelPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWrite1hPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**InputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Model** | **string** |  | 
**OrgProviderKeyId** | **string** |  | 
**OutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**PriceSource** | Pointer to **NullableString** |  | [optional] 
**PricingId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrgProviderKeyModelPublic

`func NewOrgProviderKeyModelPublic(createdAt time.Time, enabled bool, id string, model string, orgProviderKeyId string, ) *OrgProviderKeyModelPublic`

NewOrgProviderKeyModelPublic instantiates a new OrgProviderKeyModelPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgProviderKeyModelPublicWithDefaults

`func NewOrgProviderKeyModelPublicWithDefaults() *OrgProviderKeyModelPublic`

NewOrgProviderKeyModelPublicWithDefaults instantiates a new OrgProviderKeyModelPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *OrgProviderKeyModelPublic) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *OrgProviderKeyModelPublic) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *OrgProviderKeyModelPublic) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *OrgProviderKeyModelPublic) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *OrgProviderKeyModelPublic) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *OrgProviderKeyModelPublic) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWrite1hPricePerMillion

`func (o *OrgProviderKeyModelPublic) GetCacheWrite1hPricePerMillion() float32`

GetCacheWrite1hPricePerMillion returns the CacheWrite1hPricePerMillion field if non-nil, zero value otherwise.

### GetCacheWrite1hPricePerMillionOk

`func (o *OrgProviderKeyModelPublic) GetCacheWrite1hPricePerMillionOk() (*float32, bool)`

GetCacheWrite1hPricePerMillionOk returns a tuple with the CacheWrite1hPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hPricePerMillion

`func (o *OrgProviderKeyModelPublic) SetCacheWrite1hPricePerMillion(v float32)`

SetCacheWrite1hPricePerMillion sets CacheWrite1hPricePerMillion field to given value.

### HasCacheWrite1hPricePerMillion

`func (o *OrgProviderKeyModelPublic) HasCacheWrite1hPricePerMillion() bool`

HasCacheWrite1hPricePerMillion returns a boolean if a field has been set.

### SetCacheWrite1hPricePerMillionNil

`func (o *OrgProviderKeyModelPublic) SetCacheWrite1hPricePerMillionNil(b bool)`

 SetCacheWrite1hPricePerMillionNil sets the value for CacheWrite1hPricePerMillion to be an explicit nil

### UnsetCacheWrite1hPricePerMillion
`func (o *OrgProviderKeyModelPublic) UnsetCacheWrite1hPricePerMillion()`

UnsetCacheWrite1hPricePerMillion ensures that no value is present for CacheWrite1hPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *OrgProviderKeyModelPublic) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *OrgProviderKeyModelPublic) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *OrgProviderKeyModelPublic) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *OrgProviderKeyModelPublic) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *OrgProviderKeyModelPublic) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *OrgProviderKeyModelPublic) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetCreatedAt

`func (o *OrgProviderKeyModelPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgProviderKeyModelPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgProviderKeyModelPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *OrgProviderKeyModelPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrgProviderKeyModelPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrgProviderKeyModelPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *OrgProviderKeyModelPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgProviderKeyModelPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgProviderKeyModelPublic) SetId(v string)`

SetId sets Id field to given value.


### GetInputPricePerMillion

`func (o *OrgProviderKeyModelPublic) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *OrgProviderKeyModelPublic) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *OrgProviderKeyModelPublic) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *OrgProviderKeyModelPublic) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### SetInputPricePerMillionNil

`func (o *OrgProviderKeyModelPublic) SetInputPricePerMillionNil(b bool)`

 SetInputPricePerMillionNil sets the value for InputPricePerMillion to be an explicit nil

### UnsetInputPricePerMillion
`func (o *OrgProviderKeyModelPublic) UnsetInputPricePerMillion()`

UnsetInputPricePerMillion ensures that no value is present for InputPricePerMillion, not even an explicit nil
### GetModel

`func (o *OrgProviderKeyModelPublic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OrgProviderKeyModelPublic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OrgProviderKeyModelPublic) SetModel(v string)`

SetModel sets Model field to given value.


### GetOrgProviderKeyId

`func (o *OrgProviderKeyModelPublic) GetOrgProviderKeyId() string`

GetOrgProviderKeyId returns the OrgProviderKeyId field if non-nil, zero value otherwise.

### GetOrgProviderKeyIdOk

`func (o *OrgProviderKeyModelPublic) GetOrgProviderKeyIdOk() (*string, bool)`

GetOrgProviderKeyIdOk returns a tuple with the OrgProviderKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgProviderKeyId

`func (o *OrgProviderKeyModelPublic) SetOrgProviderKeyId(v string)`

SetOrgProviderKeyId sets OrgProviderKeyId field to given value.


### GetOutputPricePerMillion

`func (o *OrgProviderKeyModelPublic) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *OrgProviderKeyModelPublic) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *OrgProviderKeyModelPublic) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *OrgProviderKeyModelPublic) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### SetOutputPricePerMillionNil

`func (o *OrgProviderKeyModelPublic) SetOutputPricePerMillionNil(b bool)`

 SetOutputPricePerMillionNil sets the value for OutputPricePerMillion to be an explicit nil

### UnsetOutputPricePerMillion
`func (o *OrgProviderKeyModelPublic) UnsetOutputPricePerMillion()`

UnsetOutputPricePerMillion ensures that no value is present for OutputPricePerMillion, not even an explicit nil
### GetPriceSource

`func (o *OrgProviderKeyModelPublic) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *OrgProviderKeyModelPublic) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *OrgProviderKeyModelPublic) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.

### HasPriceSource

`func (o *OrgProviderKeyModelPublic) HasPriceSource() bool`

HasPriceSource returns a boolean if a field has been set.

### SetPriceSourceNil

`func (o *OrgProviderKeyModelPublic) SetPriceSourceNil(b bool)`

 SetPriceSourceNil sets the value for PriceSource to be an explicit nil

### UnsetPriceSource
`func (o *OrgProviderKeyModelPublic) UnsetPriceSource()`

UnsetPriceSource ensures that no value is present for PriceSource, not even an explicit nil
### GetPricingId

`func (o *OrgProviderKeyModelPublic) GetPricingId() string`

GetPricingId returns the PricingId field if non-nil, zero value otherwise.

### GetPricingIdOk

`func (o *OrgProviderKeyModelPublic) GetPricingIdOk() (*string, bool)`

GetPricingIdOk returns a tuple with the PricingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingId

`func (o *OrgProviderKeyModelPublic) SetPricingId(v string)`

SetPricingId sets PricingId field to given value.

### HasPricingId

`func (o *OrgProviderKeyModelPublic) HasPricingId() bool`

HasPricingId returns a boolean if a field has been set.

### SetPricingIdNil

`func (o *OrgProviderKeyModelPublic) SetPricingIdNil(b bool)`

 SetPricingIdNil sets the value for PricingId to be an explicit nil

### UnsetPricingId
`func (o *OrgProviderKeyModelPublic) UnsetPricingId()`

UnsetPricingId ensures that no value is present for PricingId, not even an explicit nil
### GetUpdatedAt

`func (o *OrgProviderKeyModelPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgProviderKeyModelPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgProviderKeyModelPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgProviderKeyModelPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OrgProviderKeyModelPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OrgProviderKeyModelPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


