# MRBetaAdvisorTool20260301

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**Model1**](Model1.md) |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**Caching** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMRBetaAdvisorTool20260301

`func NewMRBetaAdvisorTool20260301(model Model1, name string, type_ string, ) *MRBetaAdvisorTool20260301`

NewMRBetaAdvisorTool20260301 instantiates a new MRBetaAdvisorTool20260301 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaAdvisorTool20260301WithDefaults

`func NewMRBetaAdvisorTool20260301WithDefaults() *MRBetaAdvisorTool20260301`

NewMRBetaAdvisorTool20260301WithDefaults instantiates a new MRBetaAdvisorTool20260301 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *MRBetaAdvisorTool20260301) GetModel() Model1`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MRBetaAdvisorTool20260301) GetModelOk() (*Model1, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MRBetaAdvisorTool20260301) SetModel(v Model1)`

SetModel sets Model field to given value.


### GetName

`func (o *MRBetaAdvisorTool20260301) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaAdvisorTool20260301) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaAdvisorTool20260301) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaAdvisorTool20260301) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaAdvisorTool20260301) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaAdvisorTool20260301) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaAdvisorTool20260301) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaAdvisorTool20260301) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaAdvisorTool20260301) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaAdvisorTool20260301) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaAdvisorTool20260301) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaAdvisorTool20260301) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetCacheControl

`func (o *MRBetaAdvisorTool20260301) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaAdvisorTool20260301) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaAdvisorTool20260301) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaAdvisorTool20260301) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaAdvisorTool20260301) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaAdvisorTool20260301) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetCaching

`func (o *MRBetaAdvisorTool20260301) GetCaching() MRBetaCacheControlEphemeral`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *MRBetaAdvisorTool20260301) GetCachingOk() (*MRBetaCacheControlEphemeral, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *MRBetaAdvisorTool20260301) SetCaching(v MRBetaCacheControlEphemeral)`

SetCaching sets Caching field to given value.

### HasCaching

`func (o *MRBetaAdvisorTool20260301) HasCaching() bool`

HasCaching returns a boolean if a field has been set.

### SetCachingNil

`func (o *MRBetaAdvisorTool20260301) SetCachingNil(b bool)`

 SetCachingNil sets the value for Caching to be an explicit nil

### UnsetCaching
`func (o *MRBetaAdvisorTool20260301) UnsetCaching()`

UnsetCaching ensures that no value is present for Caching, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaAdvisorTool20260301) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaAdvisorTool20260301) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaAdvisorTool20260301) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaAdvisorTool20260301) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaAdvisorTool20260301) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaAdvisorTool20260301) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetMaxTokens

`func (o *MRBetaAdvisorTool20260301) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *MRBetaAdvisorTool20260301) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *MRBetaAdvisorTool20260301) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *MRBetaAdvisorTool20260301) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *MRBetaAdvisorTool20260301) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *MRBetaAdvisorTool20260301) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetMaxUses

`func (o *MRBetaAdvisorTool20260301) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *MRBetaAdvisorTool20260301) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *MRBetaAdvisorTool20260301) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *MRBetaAdvisorTool20260301) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### SetMaxUsesNil

`func (o *MRBetaAdvisorTool20260301) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *MRBetaAdvisorTool20260301) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
### GetStrict

`func (o *MRBetaAdvisorTool20260301) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaAdvisorTool20260301) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaAdvisorTool20260301) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaAdvisorTool20260301) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaAdvisorTool20260301) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaAdvisorTool20260301) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


