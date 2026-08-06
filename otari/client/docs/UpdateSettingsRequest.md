# UpdateSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetEstimateDefaultOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**DefaultPricing** | Pointer to **NullableBool** |  | [optional] 
**FileUnderstandingEnabled** | Pointer to **NullableBool** |  | [optional] 
**ModelCacheTtlSeconds** | Pointer to **NullableInt32** |  | [optional] 
**ModelDiscovery** | Pointer to **NullableBool** |  | [optional] 
**ModelDiscoveryNegativeTtlSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**ModelDiscoveryTimeoutSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**ModelsDevCacheTtlSeconds** | Pointer to **NullableInt32** |  | [optional] 
**ModelsDevMetadata** | Pointer to **NullableBool** |  | [optional] 
**RejectUserMismatch** | Pointer to **NullableBool** |  | [optional] 
**RequirePricing** | Pointer to **NullableBool** |  | [optional] 
**StreamMissingUsagePolicy** | Pointer to **NullableString** |  | [optional] 
**VisionDescribeMaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**VisionDescribeModel** | Pointer to **NullableString** |  | [optional] 
**VisionStrategy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateSettingsRequest

`func NewUpdateSettingsRequest() *UpdateSettingsRequest`

NewUpdateSettingsRequest instantiates a new UpdateSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingsRequestWithDefaults

`func NewUpdateSettingsRequestWithDefaults() *UpdateSettingsRequest`

NewUpdateSettingsRequestWithDefaults instantiates a new UpdateSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetEstimateDefaultOutputTokens

`func (o *UpdateSettingsRequest) GetBudgetEstimateDefaultOutputTokens() int32`

GetBudgetEstimateDefaultOutputTokens returns the BudgetEstimateDefaultOutputTokens field if non-nil, zero value otherwise.

### GetBudgetEstimateDefaultOutputTokensOk

`func (o *UpdateSettingsRequest) GetBudgetEstimateDefaultOutputTokensOk() (*int32, bool)`

GetBudgetEstimateDefaultOutputTokensOk returns a tuple with the BudgetEstimateDefaultOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetEstimateDefaultOutputTokens

`func (o *UpdateSettingsRequest) SetBudgetEstimateDefaultOutputTokens(v int32)`

SetBudgetEstimateDefaultOutputTokens sets BudgetEstimateDefaultOutputTokens field to given value.

### HasBudgetEstimateDefaultOutputTokens

`func (o *UpdateSettingsRequest) HasBudgetEstimateDefaultOutputTokens() bool`

HasBudgetEstimateDefaultOutputTokens returns a boolean if a field has been set.

### SetBudgetEstimateDefaultOutputTokensNil

`func (o *UpdateSettingsRequest) SetBudgetEstimateDefaultOutputTokensNil(b bool)`

 SetBudgetEstimateDefaultOutputTokensNil sets the value for BudgetEstimateDefaultOutputTokens to be an explicit nil

### UnsetBudgetEstimateDefaultOutputTokens
`func (o *UpdateSettingsRequest) UnsetBudgetEstimateDefaultOutputTokens()`

UnsetBudgetEstimateDefaultOutputTokens ensures that no value is present for BudgetEstimateDefaultOutputTokens, not even an explicit nil
### GetDefaultPricing

`func (o *UpdateSettingsRequest) GetDefaultPricing() bool`

GetDefaultPricing returns the DefaultPricing field if non-nil, zero value otherwise.

### GetDefaultPricingOk

`func (o *UpdateSettingsRequest) GetDefaultPricingOk() (*bool, bool)`

GetDefaultPricingOk returns a tuple with the DefaultPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPricing

`func (o *UpdateSettingsRequest) SetDefaultPricing(v bool)`

SetDefaultPricing sets DefaultPricing field to given value.

### HasDefaultPricing

`func (o *UpdateSettingsRequest) HasDefaultPricing() bool`

HasDefaultPricing returns a boolean if a field has been set.

### SetDefaultPricingNil

`func (o *UpdateSettingsRequest) SetDefaultPricingNil(b bool)`

 SetDefaultPricingNil sets the value for DefaultPricing to be an explicit nil

### UnsetDefaultPricing
`func (o *UpdateSettingsRequest) UnsetDefaultPricing()`

UnsetDefaultPricing ensures that no value is present for DefaultPricing, not even an explicit nil
### GetFileUnderstandingEnabled

`func (o *UpdateSettingsRequest) GetFileUnderstandingEnabled() bool`

GetFileUnderstandingEnabled returns the FileUnderstandingEnabled field if non-nil, zero value otherwise.

### GetFileUnderstandingEnabledOk

`func (o *UpdateSettingsRequest) GetFileUnderstandingEnabledOk() (*bool, bool)`

GetFileUnderstandingEnabledOk returns a tuple with the FileUnderstandingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUnderstandingEnabled

`func (o *UpdateSettingsRequest) SetFileUnderstandingEnabled(v bool)`

SetFileUnderstandingEnabled sets FileUnderstandingEnabled field to given value.

### HasFileUnderstandingEnabled

`func (o *UpdateSettingsRequest) HasFileUnderstandingEnabled() bool`

HasFileUnderstandingEnabled returns a boolean if a field has been set.

### SetFileUnderstandingEnabledNil

`func (o *UpdateSettingsRequest) SetFileUnderstandingEnabledNil(b bool)`

 SetFileUnderstandingEnabledNil sets the value for FileUnderstandingEnabled to be an explicit nil

### UnsetFileUnderstandingEnabled
`func (o *UpdateSettingsRequest) UnsetFileUnderstandingEnabled()`

UnsetFileUnderstandingEnabled ensures that no value is present for FileUnderstandingEnabled, not even an explicit nil
### GetModelCacheTtlSeconds

`func (o *UpdateSettingsRequest) GetModelCacheTtlSeconds() int32`

GetModelCacheTtlSeconds returns the ModelCacheTtlSeconds field if non-nil, zero value otherwise.

### GetModelCacheTtlSecondsOk

`func (o *UpdateSettingsRequest) GetModelCacheTtlSecondsOk() (*int32, bool)`

GetModelCacheTtlSecondsOk returns a tuple with the ModelCacheTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCacheTtlSeconds

`func (o *UpdateSettingsRequest) SetModelCacheTtlSeconds(v int32)`

SetModelCacheTtlSeconds sets ModelCacheTtlSeconds field to given value.

### HasModelCacheTtlSeconds

`func (o *UpdateSettingsRequest) HasModelCacheTtlSeconds() bool`

HasModelCacheTtlSeconds returns a boolean if a field has been set.

### SetModelCacheTtlSecondsNil

`func (o *UpdateSettingsRequest) SetModelCacheTtlSecondsNil(b bool)`

 SetModelCacheTtlSecondsNil sets the value for ModelCacheTtlSeconds to be an explicit nil

### UnsetModelCacheTtlSeconds
`func (o *UpdateSettingsRequest) UnsetModelCacheTtlSeconds()`

UnsetModelCacheTtlSeconds ensures that no value is present for ModelCacheTtlSeconds, not even an explicit nil
### GetModelDiscovery

`func (o *UpdateSettingsRequest) GetModelDiscovery() bool`

GetModelDiscovery returns the ModelDiscovery field if non-nil, zero value otherwise.

### GetModelDiscoveryOk

`func (o *UpdateSettingsRequest) GetModelDiscoveryOk() (*bool, bool)`

GetModelDiscoveryOk returns a tuple with the ModelDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDiscovery

`func (o *UpdateSettingsRequest) SetModelDiscovery(v bool)`

SetModelDiscovery sets ModelDiscovery field to given value.

### HasModelDiscovery

`func (o *UpdateSettingsRequest) HasModelDiscovery() bool`

HasModelDiscovery returns a boolean if a field has been set.

### SetModelDiscoveryNil

`func (o *UpdateSettingsRequest) SetModelDiscoveryNil(b bool)`

 SetModelDiscoveryNil sets the value for ModelDiscovery to be an explicit nil

### UnsetModelDiscovery
`func (o *UpdateSettingsRequest) UnsetModelDiscovery()`

UnsetModelDiscovery ensures that no value is present for ModelDiscovery, not even an explicit nil
### GetModelDiscoveryNegativeTtlSeconds

`func (o *UpdateSettingsRequest) GetModelDiscoveryNegativeTtlSeconds() float32`

GetModelDiscoveryNegativeTtlSeconds returns the ModelDiscoveryNegativeTtlSeconds field if non-nil, zero value otherwise.

### GetModelDiscoveryNegativeTtlSecondsOk

`func (o *UpdateSettingsRequest) GetModelDiscoveryNegativeTtlSecondsOk() (*float32, bool)`

GetModelDiscoveryNegativeTtlSecondsOk returns a tuple with the ModelDiscoveryNegativeTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDiscoveryNegativeTtlSeconds

`func (o *UpdateSettingsRequest) SetModelDiscoveryNegativeTtlSeconds(v float32)`

SetModelDiscoveryNegativeTtlSeconds sets ModelDiscoveryNegativeTtlSeconds field to given value.

### HasModelDiscoveryNegativeTtlSeconds

`func (o *UpdateSettingsRequest) HasModelDiscoveryNegativeTtlSeconds() bool`

HasModelDiscoveryNegativeTtlSeconds returns a boolean if a field has been set.

### SetModelDiscoveryNegativeTtlSecondsNil

`func (o *UpdateSettingsRequest) SetModelDiscoveryNegativeTtlSecondsNil(b bool)`

 SetModelDiscoveryNegativeTtlSecondsNil sets the value for ModelDiscoveryNegativeTtlSeconds to be an explicit nil

### UnsetModelDiscoveryNegativeTtlSeconds
`func (o *UpdateSettingsRequest) UnsetModelDiscoveryNegativeTtlSeconds()`

UnsetModelDiscoveryNegativeTtlSeconds ensures that no value is present for ModelDiscoveryNegativeTtlSeconds, not even an explicit nil
### GetModelDiscoveryTimeoutSeconds

`func (o *UpdateSettingsRequest) GetModelDiscoveryTimeoutSeconds() float32`

GetModelDiscoveryTimeoutSeconds returns the ModelDiscoveryTimeoutSeconds field if non-nil, zero value otherwise.

### GetModelDiscoveryTimeoutSecondsOk

`func (o *UpdateSettingsRequest) GetModelDiscoveryTimeoutSecondsOk() (*float32, bool)`

GetModelDiscoveryTimeoutSecondsOk returns a tuple with the ModelDiscoveryTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDiscoveryTimeoutSeconds

`func (o *UpdateSettingsRequest) SetModelDiscoveryTimeoutSeconds(v float32)`

SetModelDiscoveryTimeoutSeconds sets ModelDiscoveryTimeoutSeconds field to given value.

### HasModelDiscoveryTimeoutSeconds

`func (o *UpdateSettingsRequest) HasModelDiscoveryTimeoutSeconds() bool`

HasModelDiscoveryTimeoutSeconds returns a boolean if a field has been set.

### SetModelDiscoveryTimeoutSecondsNil

`func (o *UpdateSettingsRequest) SetModelDiscoveryTimeoutSecondsNil(b bool)`

 SetModelDiscoveryTimeoutSecondsNil sets the value for ModelDiscoveryTimeoutSeconds to be an explicit nil

### UnsetModelDiscoveryTimeoutSeconds
`func (o *UpdateSettingsRequest) UnsetModelDiscoveryTimeoutSeconds()`

UnsetModelDiscoveryTimeoutSeconds ensures that no value is present for ModelDiscoveryTimeoutSeconds, not even an explicit nil
### GetModelsDevCacheTtlSeconds

`func (o *UpdateSettingsRequest) GetModelsDevCacheTtlSeconds() int32`

GetModelsDevCacheTtlSeconds returns the ModelsDevCacheTtlSeconds field if non-nil, zero value otherwise.

### GetModelsDevCacheTtlSecondsOk

`func (o *UpdateSettingsRequest) GetModelsDevCacheTtlSecondsOk() (*int32, bool)`

GetModelsDevCacheTtlSecondsOk returns a tuple with the ModelsDevCacheTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelsDevCacheTtlSeconds

`func (o *UpdateSettingsRequest) SetModelsDevCacheTtlSeconds(v int32)`

SetModelsDevCacheTtlSeconds sets ModelsDevCacheTtlSeconds field to given value.

### HasModelsDevCacheTtlSeconds

`func (o *UpdateSettingsRequest) HasModelsDevCacheTtlSeconds() bool`

HasModelsDevCacheTtlSeconds returns a boolean if a field has been set.

### SetModelsDevCacheTtlSecondsNil

`func (o *UpdateSettingsRequest) SetModelsDevCacheTtlSecondsNil(b bool)`

 SetModelsDevCacheTtlSecondsNil sets the value for ModelsDevCacheTtlSeconds to be an explicit nil

### UnsetModelsDevCacheTtlSeconds
`func (o *UpdateSettingsRequest) UnsetModelsDevCacheTtlSeconds()`

UnsetModelsDevCacheTtlSeconds ensures that no value is present for ModelsDevCacheTtlSeconds, not even an explicit nil
### GetModelsDevMetadata

`func (o *UpdateSettingsRequest) GetModelsDevMetadata() bool`

GetModelsDevMetadata returns the ModelsDevMetadata field if non-nil, zero value otherwise.

### GetModelsDevMetadataOk

`func (o *UpdateSettingsRequest) GetModelsDevMetadataOk() (*bool, bool)`

GetModelsDevMetadataOk returns a tuple with the ModelsDevMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelsDevMetadata

`func (o *UpdateSettingsRequest) SetModelsDevMetadata(v bool)`

SetModelsDevMetadata sets ModelsDevMetadata field to given value.

### HasModelsDevMetadata

`func (o *UpdateSettingsRequest) HasModelsDevMetadata() bool`

HasModelsDevMetadata returns a boolean if a field has been set.

### SetModelsDevMetadataNil

`func (o *UpdateSettingsRequest) SetModelsDevMetadataNil(b bool)`

 SetModelsDevMetadataNil sets the value for ModelsDevMetadata to be an explicit nil

### UnsetModelsDevMetadata
`func (o *UpdateSettingsRequest) UnsetModelsDevMetadata()`

UnsetModelsDevMetadata ensures that no value is present for ModelsDevMetadata, not even an explicit nil
### GetRejectUserMismatch

`func (o *UpdateSettingsRequest) GetRejectUserMismatch() bool`

GetRejectUserMismatch returns the RejectUserMismatch field if non-nil, zero value otherwise.

### GetRejectUserMismatchOk

`func (o *UpdateSettingsRequest) GetRejectUserMismatchOk() (*bool, bool)`

GetRejectUserMismatchOk returns a tuple with the RejectUserMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectUserMismatch

`func (o *UpdateSettingsRequest) SetRejectUserMismatch(v bool)`

SetRejectUserMismatch sets RejectUserMismatch field to given value.

### HasRejectUserMismatch

`func (o *UpdateSettingsRequest) HasRejectUserMismatch() bool`

HasRejectUserMismatch returns a boolean if a field has been set.

### SetRejectUserMismatchNil

`func (o *UpdateSettingsRequest) SetRejectUserMismatchNil(b bool)`

 SetRejectUserMismatchNil sets the value for RejectUserMismatch to be an explicit nil

### UnsetRejectUserMismatch
`func (o *UpdateSettingsRequest) UnsetRejectUserMismatch()`

UnsetRejectUserMismatch ensures that no value is present for RejectUserMismatch, not even an explicit nil
### GetRequirePricing

`func (o *UpdateSettingsRequest) GetRequirePricing() bool`

GetRequirePricing returns the RequirePricing field if non-nil, zero value otherwise.

### GetRequirePricingOk

`func (o *UpdateSettingsRequest) GetRequirePricingOk() (*bool, bool)`

GetRequirePricingOk returns a tuple with the RequirePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePricing

`func (o *UpdateSettingsRequest) SetRequirePricing(v bool)`

SetRequirePricing sets RequirePricing field to given value.

### HasRequirePricing

`func (o *UpdateSettingsRequest) HasRequirePricing() bool`

HasRequirePricing returns a boolean if a field has been set.

### SetRequirePricingNil

`func (o *UpdateSettingsRequest) SetRequirePricingNil(b bool)`

 SetRequirePricingNil sets the value for RequirePricing to be an explicit nil

### UnsetRequirePricing
`func (o *UpdateSettingsRequest) UnsetRequirePricing()`

UnsetRequirePricing ensures that no value is present for RequirePricing, not even an explicit nil
### GetStreamMissingUsagePolicy

`func (o *UpdateSettingsRequest) GetStreamMissingUsagePolicy() string`

GetStreamMissingUsagePolicy returns the StreamMissingUsagePolicy field if non-nil, zero value otherwise.

### GetStreamMissingUsagePolicyOk

`func (o *UpdateSettingsRequest) GetStreamMissingUsagePolicyOk() (*string, bool)`

GetStreamMissingUsagePolicyOk returns a tuple with the StreamMissingUsagePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamMissingUsagePolicy

`func (o *UpdateSettingsRequest) SetStreamMissingUsagePolicy(v string)`

SetStreamMissingUsagePolicy sets StreamMissingUsagePolicy field to given value.

### HasStreamMissingUsagePolicy

`func (o *UpdateSettingsRequest) HasStreamMissingUsagePolicy() bool`

HasStreamMissingUsagePolicy returns a boolean if a field has been set.

### SetStreamMissingUsagePolicyNil

`func (o *UpdateSettingsRequest) SetStreamMissingUsagePolicyNil(b bool)`

 SetStreamMissingUsagePolicyNil sets the value for StreamMissingUsagePolicy to be an explicit nil

### UnsetStreamMissingUsagePolicy
`func (o *UpdateSettingsRequest) UnsetStreamMissingUsagePolicy()`

UnsetStreamMissingUsagePolicy ensures that no value is present for StreamMissingUsagePolicy, not even an explicit nil
### GetVisionDescribeMaxTokens

`func (o *UpdateSettingsRequest) GetVisionDescribeMaxTokens() int32`

GetVisionDescribeMaxTokens returns the VisionDescribeMaxTokens field if non-nil, zero value otherwise.

### GetVisionDescribeMaxTokensOk

`func (o *UpdateSettingsRequest) GetVisionDescribeMaxTokensOk() (*int32, bool)`

GetVisionDescribeMaxTokensOk returns a tuple with the VisionDescribeMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionDescribeMaxTokens

`func (o *UpdateSettingsRequest) SetVisionDescribeMaxTokens(v int32)`

SetVisionDescribeMaxTokens sets VisionDescribeMaxTokens field to given value.

### HasVisionDescribeMaxTokens

`func (o *UpdateSettingsRequest) HasVisionDescribeMaxTokens() bool`

HasVisionDescribeMaxTokens returns a boolean if a field has been set.

### SetVisionDescribeMaxTokensNil

`func (o *UpdateSettingsRequest) SetVisionDescribeMaxTokensNil(b bool)`

 SetVisionDescribeMaxTokensNil sets the value for VisionDescribeMaxTokens to be an explicit nil

### UnsetVisionDescribeMaxTokens
`func (o *UpdateSettingsRequest) UnsetVisionDescribeMaxTokens()`

UnsetVisionDescribeMaxTokens ensures that no value is present for VisionDescribeMaxTokens, not even an explicit nil
### GetVisionDescribeModel

`func (o *UpdateSettingsRequest) GetVisionDescribeModel() string`

GetVisionDescribeModel returns the VisionDescribeModel field if non-nil, zero value otherwise.

### GetVisionDescribeModelOk

`func (o *UpdateSettingsRequest) GetVisionDescribeModelOk() (*string, bool)`

GetVisionDescribeModelOk returns a tuple with the VisionDescribeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionDescribeModel

`func (o *UpdateSettingsRequest) SetVisionDescribeModel(v string)`

SetVisionDescribeModel sets VisionDescribeModel field to given value.

### HasVisionDescribeModel

`func (o *UpdateSettingsRequest) HasVisionDescribeModel() bool`

HasVisionDescribeModel returns a boolean if a field has been set.

### SetVisionDescribeModelNil

`func (o *UpdateSettingsRequest) SetVisionDescribeModelNil(b bool)`

 SetVisionDescribeModelNil sets the value for VisionDescribeModel to be an explicit nil

### UnsetVisionDescribeModel
`func (o *UpdateSettingsRequest) UnsetVisionDescribeModel()`

UnsetVisionDescribeModel ensures that no value is present for VisionDescribeModel, not even an explicit nil
### GetVisionStrategy

`func (o *UpdateSettingsRequest) GetVisionStrategy() string`

GetVisionStrategy returns the VisionStrategy field if non-nil, zero value otherwise.

### GetVisionStrategyOk

`func (o *UpdateSettingsRequest) GetVisionStrategyOk() (*string, bool)`

GetVisionStrategyOk returns a tuple with the VisionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionStrategy

`func (o *UpdateSettingsRequest) SetVisionStrategy(v string)`

SetVisionStrategy sets VisionStrategy field to given value.

### HasVisionStrategy

`func (o *UpdateSettingsRequest) HasVisionStrategy() bool`

HasVisionStrategy returns a boolean if a field has been set.

### SetVisionStrategyNil

`func (o *UpdateSettingsRequest) SetVisionStrategyNil(b bool)`

 SetVisionStrategyNil sets the value for VisionStrategy to be an explicit nil

### UnsetVisionStrategy
`func (o *UpdateSettingsRequest) UnsetVisionStrategy()`

UnsetVisionStrategy ensures that no value is present for VisionStrategy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


