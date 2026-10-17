# HostedModelUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**InputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**OutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHostedModelUpdateRequest

`func NewHostedModelUpdateRequest() *HostedModelUpdateRequest`

NewHostedModelUpdateRequest instantiates a new HostedModelUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedModelUpdateRequestWithDefaults

`func NewHostedModelUpdateRequestWithDefaults() *HostedModelUpdateRequest`

NewHostedModelUpdateRequestWithDefaults instantiates a new HostedModelUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *HostedModelUpdateRequest) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *HostedModelUpdateRequest) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *HostedModelUpdateRequest) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *HostedModelUpdateRequest) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *HostedModelUpdateRequest) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *HostedModelUpdateRequest) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *HostedModelUpdateRequest) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *HostedModelUpdateRequest) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *HostedModelUpdateRequest) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *HostedModelUpdateRequest) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *HostedModelUpdateRequest) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *HostedModelUpdateRequest) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetEnabled

`func (o *HostedModelUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedModelUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedModelUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HostedModelUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *HostedModelUpdateRequest) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *HostedModelUpdateRequest) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetInputPricePerMillion

`func (o *HostedModelUpdateRequest) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *HostedModelUpdateRequest) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *HostedModelUpdateRequest) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *HostedModelUpdateRequest) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### SetInputPricePerMillionNil

`func (o *HostedModelUpdateRequest) SetInputPricePerMillionNil(b bool)`

 SetInputPricePerMillionNil sets the value for InputPricePerMillion to be an explicit nil

### UnsetInputPricePerMillion
`func (o *HostedModelUpdateRequest) UnsetInputPricePerMillion()`

UnsetInputPricePerMillion ensures that no value is present for InputPricePerMillion, not even an explicit nil
### GetOutputPricePerMillion

`func (o *HostedModelUpdateRequest) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *HostedModelUpdateRequest) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *HostedModelUpdateRequest) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *HostedModelUpdateRequest) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### SetOutputPricePerMillionNil

`func (o *HostedModelUpdateRequest) SetOutputPricePerMillionNil(b bool)`

 SetOutputPricePerMillionNil sets the value for OutputPricePerMillion to be an explicit nil

### UnsetOutputPricePerMillion
`func (o *HostedModelUpdateRequest) UnsetOutputPricePerMillion()`

UnsetOutputPricePerMillion ensures that no value is present for OutputPricePerMillion, not even an explicit nil
### GetUnit

`func (o *HostedModelUpdateRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostedModelUpdateRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostedModelUpdateRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HostedModelUpdateRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HostedModelUpdateRequest) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HostedModelUpdateRequest) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


