# MRMessageUsageIterationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**MRBetaCacheCreation**](MRBetaCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | **int32** |  | 
**CacheReadInputTokens** | **int32** |  | 
**InputTokens** | **int32** |  | 
**Model** | [**Model1**](Model1.md) |  | 
**OutputTokens** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRMessageUsageIterationsInner

`func NewMRMessageUsageIterationsInner(cacheCreationInputTokens int32, cacheReadInputTokens int32, inputTokens int32, model Model1, outputTokens int32, type_ string, ) *MRMessageUsageIterationsInner`

NewMRMessageUsageIterationsInner instantiates a new MRMessageUsageIterationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRMessageUsageIterationsInnerWithDefaults

`func NewMRMessageUsageIterationsInnerWithDefaults() *MRMessageUsageIterationsInner`

NewMRMessageUsageIterationsInnerWithDefaults instantiates a new MRMessageUsageIterationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRMessageUsageIterationsInner) GetCacheCreation() MRBetaCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRMessageUsageIterationsInner) GetCacheCreationOk() (*MRBetaCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRMessageUsageIterationsInner) SetCacheCreation(v MRBetaCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRMessageUsageIterationsInner) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### GetCacheCreationInputTokens

`func (o *MRMessageUsageIterationsInner) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRMessageUsageIterationsInner) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRMessageUsageIterationsInner) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.


### GetCacheReadInputTokens

`func (o *MRMessageUsageIterationsInner) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRMessageUsageIterationsInner) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRMessageUsageIterationsInner) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.


### GetInputTokens

`func (o *MRMessageUsageIterationsInner) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRMessageUsageIterationsInner) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRMessageUsageIterationsInner) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetModel

`func (o *MRMessageUsageIterationsInner) GetModel() Model1`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MRMessageUsageIterationsInner) GetModelOk() (*Model1, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MRMessageUsageIterationsInner) SetModel(v Model1)`

SetModel sets Model field to given value.


### GetOutputTokens

`func (o *MRMessageUsageIterationsInner) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRMessageUsageIterationsInner) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRMessageUsageIterationsInner) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetType

`func (o *MRMessageUsageIterationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRMessageUsageIterationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRMessageUsageIterationsInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


