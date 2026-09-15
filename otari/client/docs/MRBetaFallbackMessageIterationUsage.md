# MRBetaFallbackMessageIterationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**NullableMRBetaCacheCreation**](MRBetaCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | **int32** |  | 
**CacheReadInputTokens** | **int32** |  | 
**InputTokens** | **int32** |  | 
**Model** | [**Model1**](Model1.md) |  | 
**OutputTokens** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaFallbackMessageIterationUsage

`func NewMRBetaFallbackMessageIterationUsage(cacheCreationInputTokens int32, cacheReadInputTokens int32, inputTokens int32, model Model1, outputTokens int32, type_ string, ) *MRBetaFallbackMessageIterationUsage`

NewMRBetaFallbackMessageIterationUsage instantiates a new MRBetaFallbackMessageIterationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaFallbackMessageIterationUsageWithDefaults

`func NewMRBetaFallbackMessageIterationUsageWithDefaults() *MRBetaFallbackMessageIterationUsage`

NewMRBetaFallbackMessageIterationUsageWithDefaults instantiates a new MRBetaFallbackMessageIterationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheCreation() MRBetaCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheCreationOk() (*MRBetaCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRBetaFallbackMessageIterationUsage) SetCacheCreation(v MRBetaCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRBetaFallbackMessageIterationUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### SetCacheCreationNil

`func (o *MRBetaFallbackMessageIterationUsage) SetCacheCreationNil(b bool)`

 SetCacheCreationNil sets the value for CacheCreation to be an explicit nil

### UnsetCacheCreation
`func (o *MRBetaFallbackMessageIterationUsage) UnsetCacheCreation()`

UnsetCacheCreation ensures that no value is present for CacheCreation, not even an explicit nil
### GetCacheCreationInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.


### GetCacheReadInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRBetaFallbackMessageIterationUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.


### GetInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRBetaFallbackMessageIterationUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRBetaFallbackMessageIterationUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetModel

`func (o *MRBetaFallbackMessageIterationUsage) GetModel() Model1`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MRBetaFallbackMessageIterationUsage) GetModelOk() (*Model1, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MRBetaFallbackMessageIterationUsage) SetModel(v Model1)`

SetModel sets Model field to given value.


### GetOutputTokens

`func (o *MRBetaFallbackMessageIterationUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRBetaFallbackMessageIterationUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRBetaFallbackMessageIterationUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetType

`func (o *MRBetaFallbackMessageIterationUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaFallbackMessageIterationUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaFallbackMessageIterationUsage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


