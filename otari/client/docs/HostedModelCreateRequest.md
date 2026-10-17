# HostedModelCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**InputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Model** | **string** |  | 
**OutputPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHostedModelCreateRequest

`func NewHostedModelCreateRequest(model string, ) *HostedModelCreateRequest`

NewHostedModelCreateRequest instantiates a new HostedModelCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedModelCreateRequestWithDefaults

`func NewHostedModelCreateRequestWithDefaults() *HostedModelCreateRequest`

NewHostedModelCreateRequestWithDefaults instantiates a new HostedModelCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadPricePerMillion

`func (o *HostedModelCreateRequest) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *HostedModelCreateRequest) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *HostedModelCreateRequest) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *HostedModelCreateRequest) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *HostedModelCreateRequest) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *HostedModelCreateRequest) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *HostedModelCreateRequest) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *HostedModelCreateRequest) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *HostedModelCreateRequest) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *HostedModelCreateRequest) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *HostedModelCreateRequest) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *HostedModelCreateRequest) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetInputPricePerMillion

`func (o *HostedModelCreateRequest) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *HostedModelCreateRequest) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *HostedModelCreateRequest) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.

### HasInputPricePerMillion

`func (o *HostedModelCreateRequest) HasInputPricePerMillion() bool`

HasInputPricePerMillion returns a boolean if a field has been set.

### SetInputPricePerMillionNil

`func (o *HostedModelCreateRequest) SetInputPricePerMillionNil(b bool)`

 SetInputPricePerMillionNil sets the value for InputPricePerMillion to be an explicit nil

### UnsetInputPricePerMillion
`func (o *HostedModelCreateRequest) UnsetInputPricePerMillion()`

UnsetInputPricePerMillion ensures that no value is present for InputPricePerMillion, not even an explicit nil
### GetModel

`func (o *HostedModelCreateRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HostedModelCreateRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HostedModelCreateRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputPricePerMillion

`func (o *HostedModelCreateRequest) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *HostedModelCreateRequest) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *HostedModelCreateRequest) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.

### HasOutputPricePerMillion

`func (o *HostedModelCreateRequest) HasOutputPricePerMillion() bool`

HasOutputPricePerMillion returns a boolean if a field has been set.

### SetOutputPricePerMillionNil

`func (o *HostedModelCreateRequest) SetOutputPricePerMillionNil(b bool)`

 SetOutputPricePerMillionNil sets the value for OutputPricePerMillion to be an explicit nil

### UnsetOutputPricePerMillion
`func (o *HostedModelCreateRequest) UnsetOutputPricePerMillion()`

UnsetOutputPricePerMillion ensures that no value is present for OutputPricePerMillion, not even an explicit nil
### GetUnit

`func (o *HostedModelCreateRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HostedModelCreateRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HostedModelCreateRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HostedModelCreateRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HostedModelCreateRequest) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HostedModelCreateRequest) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


