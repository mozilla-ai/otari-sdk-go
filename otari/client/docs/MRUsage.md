# MRUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**NullableMRCacheCreation**](MRCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | Pointer to **NullableInt32** |  | [optional] 
**CacheReadInputTokens** | Pointer to **NullableInt32** |  | [optional] 
**InferenceGeo** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**ServerToolUse** | Pointer to [**NullableMRServerToolUsage**](MRServerToolUsage.md) |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMRUsage

`func NewMRUsage(inputTokens int32, outputTokens int32, ) *MRUsage`

NewMRUsage instantiates a new MRUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRUsageWithDefaults

`func NewMRUsageWithDefaults() *MRUsage`

NewMRUsageWithDefaults instantiates a new MRUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRUsage) GetCacheCreation() MRCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRUsage) GetCacheCreationOk() (*MRCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRUsage) SetCacheCreation(v MRCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### SetCacheCreationNil

`func (o *MRUsage) SetCacheCreationNil(b bool)`

 SetCacheCreationNil sets the value for CacheCreation to be an explicit nil

### UnsetCacheCreation
`func (o *MRUsage) UnsetCacheCreation()`

UnsetCacheCreation ensures that no value is present for CacheCreation, not even an explicit nil
### GetCacheCreationInputTokens

`func (o *MRUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.

### HasCacheCreationInputTokens

`func (o *MRUsage) HasCacheCreationInputTokens() bool`

HasCacheCreationInputTokens returns a boolean if a field has been set.

### SetCacheCreationInputTokensNil

`func (o *MRUsage) SetCacheCreationInputTokensNil(b bool)`

 SetCacheCreationInputTokensNil sets the value for CacheCreationInputTokens to be an explicit nil

### UnsetCacheCreationInputTokens
`func (o *MRUsage) UnsetCacheCreationInputTokens()`

UnsetCacheCreationInputTokens ensures that no value is present for CacheCreationInputTokens, not even an explicit nil
### GetCacheReadInputTokens

`func (o *MRUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.

### HasCacheReadInputTokens

`func (o *MRUsage) HasCacheReadInputTokens() bool`

HasCacheReadInputTokens returns a boolean if a field has been set.

### SetCacheReadInputTokensNil

`func (o *MRUsage) SetCacheReadInputTokensNil(b bool)`

 SetCacheReadInputTokensNil sets the value for CacheReadInputTokens to be an explicit nil

### UnsetCacheReadInputTokens
`func (o *MRUsage) UnsetCacheReadInputTokens()`

UnsetCacheReadInputTokens ensures that no value is present for CacheReadInputTokens, not even an explicit nil
### GetInferenceGeo

`func (o *MRUsage) GetInferenceGeo() string`

GetInferenceGeo returns the InferenceGeo field if non-nil, zero value otherwise.

### GetInferenceGeoOk

`func (o *MRUsage) GetInferenceGeoOk() (*string, bool)`

GetInferenceGeoOk returns a tuple with the InferenceGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceGeo

`func (o *MRUsage) SetInferenceGeo(v string)`

SetInferenceGeo sets InferenceGeo field to given value.

### HasInferenceGeo

`func (o *MRUsage) HasInferenceGeo() bool`

HasInferenceGeo returns a boolean if a field has been set.

### SetInferenceGeoNil

`func (o *MRUsage) SetInferenceGeoNil(b bool)`

 SetInferenceGeoNil sets the value for InferenceGeo to be an explicit nil

### UnsetInferenceGeo
`func (o *MRUsage) UnsetInferenceGeo()`

UnsetInferenceGeo ensures that no value is present for InferenceGeo, not even an explicit nil
### GetInputTokens

`func (o *MRUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *MRUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetServerToolUse

`func (o *MRUsage) GetServerToolUse() MRServerToolUsage`

GetServerToolUse returns the ServerToolUse field if non-nil, zero value otherwise.

### GetServerToolUseOk

`func (o *MRUsage) GetServerToolUseOk() (*MRServerToolUsage, bool)`

GetServerToolUseOk returns a tuple with the ServerToolUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerToolUse

`func (o *MRUsage) SetServerToolUse(v MRServerToolUsage)`

SetServerToolUse sets ServerToolUse field to given value.

### HasServerToolUse

`func (o *MRUsage) HasServerToolUse() bool`

HasServerToolUse returns a boolean if a field has been set.

### SetServerToolUseNil

`func (o *MRUsage) SetServerToolUseNil(b bool)`

 SetServerToolUseNil sets the value for ServerToolUse to be an explicit nil

### UnsetServerToolUse
`func (o *MRUsage) UnsetServerToolUse()`

UnsetServerToolUse ensures that no value is present for ServerToolUse, not even an explicit nil
### GetServiceTier

`func (o *MRUsage) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *MRUsage) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *MRUsage) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *MRUsage) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *MRUsage) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *MRUsage) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


