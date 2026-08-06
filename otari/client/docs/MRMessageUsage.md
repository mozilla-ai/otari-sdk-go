# MRMessageUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheCreation** | Pointer to [**NullableMRCacheCreation**](MRCacheCreation.md) |  | [optional] 
**CacheCreationInputTokens** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**CacheReadInputTokens** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**InferenceGeo** | Pointer to **NullableString** | Delete the alias scoped to this user. Omit to delete the global alias of that name. | [optional] 
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**ServerToolUse** | Pointer to [**NullableMRServerToolUsage**](MRServerToolUsage.md) |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 
**Iterations** | Pointer to [**[]MRMessageUsageIterationsInner**](MRMessageUsageIterationsInner.md) |  | [optional] 
**Speed** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMRMessageUsage

`func NewMRMessageUsage(inputTokens int32, outputTokens int32, ) *MRMessageUsage`

NewMRMessageUsage instantiates a new MRMessageUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRMessageUsageWithDefaults

`func NewMRMessageUsageWithDefaults() *MRMessageUsage`

NewMRMessageUsageWithDefaults instantiates a new MRMessageUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheCreation

`func (o *MRMessageUsage) GetCacheCreation() MRCacheCreation`

GetCacheCreation returns the CacheCreation field if non-nil, zero value otherwise.

### GetCacheCreationOk

`func (o *MRMessageUsage) GetCacheCreationOk() (*MRCacheCreation, bool)`

GetCacheCreationOk returns a tuple with the CacheCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreation

`func (o *MRMessageUsage) SetCacheCreation(v MRCacheCreation)`

SetCacheCreation sets CacheCreation field to given value.

### HasCacheCreation

`func (o *MRMessageUsage) HasCacheCreation() bool`

HasCacheCreation returns a boolean if a field has been set.

### SetCacheCreationNil

`func (o *MRMessageUsage) SetCacheCreationNil(b bool)`

 SetCacheCreationNil sets the value for CacheCreation to be an explicit nil

### UnsetCacheCreation
`func (o *MRMessageUsage) UnsetCacheCreation()`

UnsetCacheCreation ensures that no value is present for CacheCreation, not even an explicit nil
### GetCacheCreationInputTokens

`func (o *MRMessageUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *MRMessageUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *MRMessageUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.

### HasCacheCreationInputTokens

`func (o *MRMessageUsage) HasCacheCreationInputTokens() bool`

HasCacheCreationInputTokens returns a boolean if a field has been set.

### SetCacheCreationInputTokensNil

`func (o *MRMessageUsage) SetCacheCreationInputTokensNil(b bool)`

 SetCacheCreationInputTokensNil sets the value for CacheCreationInputTokens to be an explicit nil

### UnsetCacheCreationInputTokens
`func (o *MRMessageUsage) UnsetCacheCreationInputTokens()`

UnsetCacheCreationInputTokens ensures that no value is present for CacheCreationInputTokens, not even an explicit nil
### GetCacheReadInputTokens

`func (o *MRMessageUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *MRMessageUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *MRMessageUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.

### HasCacheReadInputTokens

`func (o *MRMessageUsage) HasCacheReadInputTokens() bool`

HasCacheReadInputTokens returns a boolean if a field has been set.

### SetCacheReadInputTokensNil

`func (o *MRMessageUsage) SetCacheReadInputTokensNil(b bool)`

 SetCacheReadInputTokensNil sets the value for CacheReadInputTokens to be an explicit nil

### UnsetCacheReadInputTokens
`func (o *MRMessageUsage) UnsetCacheReadInputTokens()`

UnsetCacheReadInputTokens ensures that no value is present for CacheReadInputTokens, not even an explicit nil
### GetInferenceGeo

`func (o *MRMessageUsage) GetInferenceGeo() string`

GetInferenceGeo returns the InferenceGeo field if non-nil, zero value otherwise.

### GetInferenceGeoOk

`func (o *MRMessageUsage) GetInferenceGeoOk() (*string, bool)`

GetInferenceGeoOk returns a tuple with the InferenceGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceGeo

`func (o *MRMessageUsage) SetInferenceGeo(v string)`

SetInferenceGeo sets InferenceGeo field to given value.

### HasInferenceGeo

`func (o *MRMessageUsage) HasInferenceGeo() bool`

HasInferenceGeo returns a boolean if a field has been set.

### SetInferenceGeoNil

`func (o *MRMessageUsage) SetInferenceGeoNil(b bool)`

 SetInferenceGeoNil sets the value for InferenceGeo to be an explicit nil

### UnsetInferenceGeo
`func (o *MRMessageUsage) UnsetInferenceGeo()`

UnsetInferenceGeo ensures that no value is present for InferenceGeo, not even an explicit nil
### GetInputTokens

`func (o *MRMessageUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *MRMessageUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *MRMessageUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *MRMessageUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *MRMessageUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *MRMessageUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetServerToolUse

`func (o *MRMessageUsage) GetServerToolUse() MRServerToolUsage`

GetServerToolUse returns the ServerToolUse field if non-nil, zero value otherwise.

### GetServerToolUseOk

`func (o *MRMessageUsage) GetServerToolUseOk() (*MRServerToolUsage, bool)`

GetServerToolUseOk returns a tuple with the ServerToolUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerToolUse

`func (o *MRMessageUsage) SetServerToolUse(v MRServerToolUsage)`

SetServerToolUse sets ServerToolUse field to given value.

### HasServerToolUse

`func (o *MRMessageUsage) HasServerToolUse() bool`

HasServerToolUse returns a boolean if a field has been set.

### SetServerToolUseNil

`func (o *MRMessageUsage) SetServerToolUseNil(b bool)`

 SetServerToolUseNil sets the value for ServerToolUse to be an explicit nil

### UnsetServerToolUse
`func (o *MRMessageUsage) UnsetServerToolUse()`

UnsetServerToolUse ensures that no value is present for ServerToolUse, not even an explicit nil
### GetServiceTier

`func (o *MRMessageUsage) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *MRMessageUsage) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *MRMessageUsage) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *MRMessageUsage) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *MRMessageUsage) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *MRMessageUsage) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetIterations

`func (o *MRMessageUsage) GetIterations() []MRMessageUsageIterationsInner`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *MRMessageUsage) GetIterationsOk() (*[]MRMessageUsageIterationsInner, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *MRMessageUsage) SetIterations(v []MRMessageUsageIterationsInner)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *MRMessageUsage) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *MRMessageUsage) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *MRMessageUsage) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetSpeed

`func (o *MRMessageUsage) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MRMessageUsage) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MRMessageUsage) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MRMessageUsage) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *MRMessageUsage) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *MRMessageUsage) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


