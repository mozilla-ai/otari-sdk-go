# MRMessageIterationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**NullableMRBetaCacheCreation**](MRBetaCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | **int32** |  | 
**CacheReadInputTokens** | **int32** |  | 
**InputTokens** | **int32** |  | 
**Model** | Pointer to **NullableString** |  | [optional] 
**OutputTokens** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRMessageIterationUsage

`func NewMRMessageIterationUsage(cacheCreationInputTokens int32, cacheReadInputTokens int32, inputTokens int32, outputTokens int32, type_ string, ) *MRMessageIterationUsage`

NewMRMessageIterationUsage instantiates a new MRMessageIterationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRMessageIterationUsageWithDefaults

`func NewMRMessageIterationUsageWithDefaults() *MRMessageIterationUsage`

NewMRMessageIterationUsageWithDefaults instantiates a new MRMessageIterationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRMessageIterationUsage) GetCacheCreation() MRBetaCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRMessageIterationUsage) GetCacheCreationOk() (*MRBetaCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRMessageIterationUsage) SetCacheCreation(v MRBetaCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRMessageIterationUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### SetCacheCreationNil

`func (o *MRMessageIterationUsage) SetCacheCreationNil(b bool)`

 SetCacheCreationNil sets the value for CacheCreation to be an explicit nil

### UnsetCacheCreation
`func (o *MRMessageIterationUsage) UnsetCacheCreation()`

UnsetCacheCreation ensures that no value is present for CacheCreation, not even an explicit nil
### GetCacheCreationInputTokens

`func (o *MRMessageIterationUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRMessageIterationUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRMessageIterationUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.


### GetCacheReadInputTokens

`func (o *MRMessageIterationUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRMessageIterationUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRMessageIterationUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.


### GetInputTokens

`func (o *MRMessageIterationUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRMessageIterationUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRMessageIterationUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetModel

`func (o *MRMessageIterationUsage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MRMessageIterationUsage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MRMessageIterationUsage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MRMessageIterationUsage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MRMessageIterationUsage) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MRMessageIterationUsage) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetOutputTokens

`func (o *MRMessageIterationUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRMessageIterationUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRMessageIterationUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetType

`func (o *MRMessageIterationUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRMessageIterationUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRMessageIterationUsage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


