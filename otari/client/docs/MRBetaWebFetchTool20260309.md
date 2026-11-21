# MRBetaWebFetchTool20260309

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**AllowedDomains** | Pointer to **[]string** |  | [optional] 
**BlockedDomains** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**Citations** | Pointer to [**NullableMRBetaCitationsConfigParam**](MRBetaCitationsConfigParam.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**MaxContentTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 
**UrlSources** | Pointer to [**NullableMRBetaWebFetchURLSources**](MRBetaWebFetchURLSources.md) |  | [optional] 
**UseCache** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMRBetaWebFetchTool20260309

`func NewMRBetaWebFetchTool20260309(name string, type_ string, ) *MRBetaWebFetchTool20260309`

NewMRBetaWebFetchTool20260309 instantiates a new MRBetaWebFetchTool20260309 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaWebFetchTool20260309WithDefaults

`func NewMRBetaWebFetchTool20260309WithDefaults() *MRBetaWebFetchTool20260309`

NewMRBetaWebFetchTool20260309WithDefaults instantiates a new MRBetaWebFetchTool20260309 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MRBetaWebFetchTool20260309) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaWebFetchTool20260309) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaWebFetchTool20260309) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaWebFetchTool20260309) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaWebFetchTool20260309) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaWebFetchTool20260309) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaWebFetchTool20260309) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaWebFetchTool20260309) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaWebFetchTool20260309) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaWebFetchTool20260309) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaWebFetchTool20260309) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaWebFetchTool20260309) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetAllowedDomains

`func (o *MRBetaWebFetchTool20260309) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *MRBetaWebFetchTool20260309) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *MRBetaWebFetchTool20260309) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *MRBetaWebFetchTool20260309) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### SetAllowedDomainsNil

`func (o *MRBetaWebFetchTool20260309) SetAllowedDomainsNil(b bool)`

 SetAllowedDomainsNil sets the value for AllowedDomains to be an explicit nil

### UnsetAllowedDomains
`func (o *MRBetaWebFetchTool20260309) UnsetAllowedDomains()`

UnsetAllowedDomains ensures that no value is present for AllowedDomains, not even an explicit nil
### GetBlockedDomains

`func (o *MRBetaWebFetchTool20260309) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *MRBetaWebFetchTool20260309) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *MRBetaWebFetchTool20260309) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.

### HasBlockedDomains

`func (o *MRBetaWebFetchTool20260309) HasBlockedDomains() bool`

HasBlockedDomains returns a boolean if a field has been set.

### SetBlockedDomainsNil

`func (o *MRBetaWebFetchTool20260309) SetBlockedDomainsNil(b bool)`

 SetBlockedDomainsNil sets the value for BlockedDomains to be an explicit nil

### UnsetBlockedDomains
`func (o *MRBetaWebFetchTool20260309) UnsetBlockedDomains()`

UnsetBlockedDomains ensures that no value is present for BlockedDomains, not even an explicit nil
### GetCacheControl

`func (o *MRBetaWebFetchTool20260309) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaWebFetchTool20260309) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaWebFetchTool20260309) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaWebFetchTool20260309) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaWebFetchTool20260309) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaWebFetchTool20260309) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetCitations

`func (o *MRBetaWebFetchTool20260309) GetCitations() MRBetaCitationsConfigParam`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *MRBetaWebFetchTool20260309) GetCitationsOk() (*MRBetaCitationsConfigParam, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *MRBetaWebFetchTool20260309) SetCitations(v MRBetaCitationsConfigParam)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *MRBetaWebFetchTool20260309) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### SetCitationsNil

`func (o *MRBetaWebFetchTool20260309) SetCitationsNil(b bool)`

 SetCitationsNil sets the value for Citations to be an explicit nil

### UnsetCitations
`func (o *MRBetaWebFetchTool20260309) UnsetCitations()`

UnsetCitations ensures that no value is present for Citations, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaWebFetchTool20260309) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaWebFetchTool20260309) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaWebFetchTool20260309) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaWebFetchTool20260309) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaWebFetchTool20260309) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaWebFetchTool20260309) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetMaxContentTokens

`func (o *MRBetaWebFetchTool20260309) GetMaxContentTokens() int32`

GetMaxContentTokens returns the MaxContentTokens field if non-nil, zero value otherwise.

### GetMaxContentTokensOk

`func (o *MRBetaWebFetchTool20260309) GetMaxContentTokensOk() (*int32, bool)`

GetMaxContentTokensOk returns a tuple with the MaxContentTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContentTokens

`func (o *MRBetaWebFetchTool20260309) SetMaxContentTokens(v int32)`

SetMaxContentTokens sets MaxContentTokens field to given value.

### HasMaxContentTokens

`func (o *MRBetaWebFetchTool20260309) HasMaxContentTokens() bool`

HasMaxContentTokens returns a boolean if a field has been set.

### SetMaxContentTokensNil

`func (o *MRBetaWebFetchTool20260309) SetMaxContentTokensNil(b bool)`

 SetMaxContentTokensNil sets the value for MaxContentTokens to be an explicit nil

### UnsetMaxContentTokens
`func (o *MRBetaWebFetchTool20260309) UnsetMaxContentTokens()`

UnsetMaxContentTokens ensures that no value is present for MaxContentTokens, not even an explicit nil
### GetMaxUses

`func (o *MRBetaWebFetchTool20260309) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *MRBetaWebFetchTool20260309) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *MRBetaWebFetchTool20260309) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *MRBetaWebFetchTool20260309) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### SetMaxUsesNil

`func (o *MRBetaWebFetchTool20260309) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *MRBetaWebFetchTool20260309) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
### GetStrict

`func (o *MRBetaWebFetchTool20260309) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaWebFetchTool20260309) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaWebFetchTool20260309) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaWebFetchTool20260309) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaWebFetchTool20260309) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaWebFetchTool20260309) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil
### GetUrlSources

`func (o *MRBetaWebFetchTool20260309) GetUrlSources() MRBetaWebFetchURLSources`

GetUrlSources returns the UrlSources field if non-nil, zero value otherwise.

### GetUrlSourcesOk

`func (o *MRBetaWebFetchTool20260309) GetUrlSourcesOk() (*MRBetaWebFetchURLSources, bool)`

GetUrlSourcesOk returns a tuple with the UrlSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSources

`func (o *MRBetaWebFetchTool20260309) SetUrlSources(v MRBetaWebFetchURLSources)`

SetUrlSources sets UrlSources field to given value.

### HasUrlSources

`func (o *MRBetaWebFetchTool20260309) HasUrlSources() bool`

HasUrlSources returns a boolean if a field has been set.

### SetUrlSourcesNil

`func (o *MRBetaWebFetchTool20260309) SetUrlSourcesNil(b bool)`

 SetUrlSourcesNil sets the value for UrlSources to be an explicit nil

### UnsetUrlSources
`func (o *MRBetaWebFetchTool20260309) UnsetUrlSources()`

UnsetUrlSources ensures that no value is present for UrlSources, not even an explicit nil
### GetUseCache

`func (o *MRBetaWebFetchTool20260309) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *MRBetaWebFetchTool20260309) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *MRBetaWebFetchTool20260309) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *MRBetaWebFetchTool20260309) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### SetUseCacheNil

`func (o *MRBetaWebFetchTool20260309) SetUseCacheNil(b bool)`

 SetUseCacheNil sets the value for UseCache to be an explicit nil

### UnsetUseCache
`func (o *MRBetaWebFetchTool20260309) UnsetUseCache()`

UnsetUseCache ensures that no value is present for UseCache, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


