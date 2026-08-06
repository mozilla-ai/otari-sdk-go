# MRBetaCompactionIterationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**NullableMRBetaCacheCreation**](MRBetaCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | **int32** |  | 
**CacheReadInputTokens** | **int32** |  | 
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaCompactionIterationUsage

`func NewMRBetaCompactionIterationUsage(cacheCreationInputTokens int32, cacheReadInputTokens int32, inputTokens int32, outputTokens int32, type_ string, ) *MRBetaCompactionIterationUsage`

NewMRBetaCompactionIterationUsage instantiates a new MRBetaCompactionIterationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCompactionIterationUsageWithDefaults

`func NewMRBetaCompactionIterationUsageWithDefaults() *MRBetaCompactionIterationUsage`

NewMRBetaCompactionIterationUsageWithDefaults instantiates a new MRBetaCompactionIterationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRBetaCompactionIterationUsage) GetCacheCreation() MRBetaCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRBetaCompactionIterationUsage) GetCacheCreationOk() (*MRBetaCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRBetaCompactionIterationUsage) SetCacheCreation(v MRBetaCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRBetaCompactionIterationUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### SetCacheCreationNil

`func (o *MRBetaCompactionIterationUsage) SetCacheCreationNil(b bool)`

 SetCacheCreationNil sets the value for CacheCreation to be an explicit nil

### UnsetCacheCreation
`func (o *MRBetaCompactionIterationUsage) UnsetCacheCreation()`

UnsetCacheCreation ensures that no value is present for CacheCreation, not even an explicit nil
### GetCacheCreationInputTokens

`func (o *MRBetaCompactionIterationUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRBetaCompactionIterationUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRBetaCompactionIterationUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.


### GetCacheReadInputTokens

`func (o *MRBetaCompactionIterationUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRBetaCompactionIterationUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRBetaCompactionIterationUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.


### GetInputTokens

`func (o *MRBetaCompactionIterationUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRBetaCompactionIterationUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRBetaCompactionIterationUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *MRBetaCompactionIterationUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRBetaCompactionIterationUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRBetaCompactionIterationUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetType

`func (o *MRBetaCompactionIterationUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCompactionIterationUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCompactionIterationUsage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


