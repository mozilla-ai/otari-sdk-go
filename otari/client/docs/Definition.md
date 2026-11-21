# Definition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSchema** | [**MRBetaResponseToolInputSchema**](MRBetaResponseToolInputSchema.md) |  | 
**Name** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**DeferLoading** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EagerInputStreaming** | Pointer to **bool** | Only end users that are, or are not, blocked. | [optional] 
**InputExamples** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Strict** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 
**CacheControl** | Pointer to [**MRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**Configs** | Pointer to [**map[string]MRBetaMCPToolConfig**](MRBetaMCPToolConfig.md) |  | [optional] 
**DisplayHeightPx** | **int32** |  | 
**DisplayWidthPx** | **int32** |  | 
**DisplayNumber** | Pointer to **int32** |  | [optional] 
**EnableZoom** | Pointer to **bool** | Only end users that are, or are not, blocked. | [optional] 
**MaxCharacters** | Pointer to **int32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**AllowedDomains** | Pointer to **[]string** |  | [optional] 
**BlockedDomains** | Pointer to **[]string** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**UserLocation** | Pointer to [**MRBetaUserLocation**](MRBetaUserLocation.md) |  | [optional] 
**Citations** | Pointer to [**MRBetaCitationsConfigParam**](MRBetaCitationsConfigParam.md) |  | [optional] 
**MaxContentTokens** | Pointer to **int32** |  | [optional] 
**UrlSources** | Pointer to [**MRBetaWebFetchURLSources**](MRBetaWebFetchURLSources.md) |  | [optional] 
**UseCache** | Pointer to **bool** |  | [optional] 
**ResponseInclusion** | Pointer to **string** |  | [optional] 
**Model** | [**Model1**](Model1.md) |  | 
**Caching** | Pointer to [**MRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**McpServerName** | **string** |  | 
**DefaultConfig** | Pointer to [**MRBetaMCPToolDefaultConfig**](MRBetaMCPToolDefaultConfig.md) |  | [optional] 
**Tools** | Pointer to [**[]MRBetaMCPToolParam**](MRBetaMCPToolParam.md) |  | [optional] 

## Methods

### NewDefinition

`func NewDefinition(inputSchema MRBetaResponseToolInputSchema, name string, type_ string, displayHeightPx int32, displayWidthPx int32, model Model1, mcpServerName string, ) *Definition`

NewDefinition instantiates a new Definition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionWithDefaults

`func NewDefinitionWithDefaults() *Definition`

NewDefinitionWithDefaults instantiates a new Definition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSchema

`func (o *Definition) GetInputSchema() MRBetaResponseToolInputSchema`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *Definition) GetInputSchemaOk() (*MRBetaResponseToolInputSchema, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *Definition) SetInputSchema(v MRBetaResponseToolInputSchema)`

SetInputSchema sets InputSchema field to given value.


### GetName

`func (o *Definition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Definition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Definition) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedCallers

`func (o *Definition) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *Definition) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *Definition) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *Definition) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### GetDeferLoading

`func (o *Definition) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *Definition) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *Definition) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *Definition) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### GetDescription

`func (o *Definition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Definition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Definition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Definition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEagerInputStreaming

`func (o *Definition) GetEagerInputStreaming() bool`

GetEagerInputStreaming returns the EagerInputStreaming field if non-nil, zero value otherwise.

### GetEagerInputStreamingOk

`func (o *Definition) GetEagerInputStreamingOk() (*bool, bool)`

GetEagerInputStreamingOk returns a tuple with the EagerInputStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEagerInputStreaming

`func (o *Definition) SetEagerInputStreaming(v bool)`

SetEagerInputStreaming sets EagerInputStreaming field to given value.

### HasEagerInputStreaming

`func (o *Definition) HasEagerInputStreaming() bool`

HasEagerInputStreaming returns a boolean if a field has been set.

### GetInputExamples

`func (o *Definition) GetInputExamples() []map[string]interface{}`

GetInputExamples returns the InputExamples field if non-nil, zero value otherwise.

### GetInputExamplesOk

`func (o *Definition) GetInputExamplesOk() (*[]map[string]interface{}, bool)`

GetInputExamplesOk returns a tuple with the InputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputExamples

`func (o *Definition) SetInputExamples(v []map[string]interface{})`

SetInputExamples sets InputExamples field to given value.

### HasInputExamples

`func (o *Definition) HasInputExamples() bool`

HasInputExamples returns a boolean if a field has been set.

### GetStrict

`func (o *Definition) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *Definition) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *Definition) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *Definition) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetType

`func (o *Definition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Definition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Definition) SetType(v string)`

SetType sets Type field to given value.


### GetCacheControl

`func (o *Definition) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *Definition) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *Definition) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *Definition) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### GetConfigs

`func (o *Definition) GetConfigs() map[string]MRBetaMCPToolConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *Definition) GetConfigsOk() (*map[string]MRBetaMCPToolConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *Definition) SetConfigs(v map[string]MRBetaMCPToolConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *Definition) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDisplayHeightPx

`func (o *Definition) GetDisplayHeightPx() int32`

GetDisplayHeightPx returns the DisplayHeightPx field if non-nil, zero value otherwise.

### GetDisplayHeightPxOk

`func (o *Definition) GetDisplayHeightPxOk() (*int32, bool)`

GetDisplayHeightPxOk returns a tuple with the DisplayHeightPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayHeightPx

`func (o *Definition) SetDisplayHeightPx(v int32)`

SetDisplayHeightPx sets DisplayHeightPx field to given value.


### GetDisplayWidthPx

`func (o *Definition) GetDisplayWidthPx() int32`

GetDisplayWidthPx returns the DisplayWidthPx field if non-nil, zero value otherwise.

### GetDisplayWidthPxOk

`func (o *Definition) GetDisplayWidthPxOk() (*int32, bool)`

GetDisplayWidthPxOk returns a tuple with the DisplayWidthPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayWidthPx

`func (o *Definition) SetDisplayWidthPx(v int32)`

SetDisplayWidthPx sets DisplayWidthPx field to given value.


### GetDisplayNumber

`func (o *Definition) GetDisplayNumber() int32`

GetDisplayNumber returns the DisplayNumber field if non-nil, zero value otherwise.

### GetDisplayNumberOk

`func (o *Definition) GetDisplayNumberOk() (*int32, bool)`

GetDisplayNumberOk returns a tuple with the DisplayNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNumber

`func (o *Definition) SetDisplayNumber(v int32)`

SetDisplayNumber sets DisplayNumber field to given value.

### HasDisplayNumber

`func (o *Definition) HasDisplayNumber() bool`

HasDisplayNumber returns a boolean if a field has been set.

### GetEnableZoom

`func (o *Definition) GetEnableZoom() bool`

GetEnableZoom returns the EnableZoom field if non-nil, zero value otherwise.

### GetEnableZoomOk

`func (o *Definition) GetEnableZoomOk() (*bool, bool)`

GetEnableZoomOk returns a tuple with the EnableZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoom

`func (o *Definition) SetEnableZoom(v bool)`

SetEnableZoom sets EnableZoom field to given value.

### HasEnableZoom

`func (o *Definition) HasEnableZoom() bool`

HasEnableZoom returns a boolean if a field has been set.

### GetMaxCharacters

`func (o *Definition) GetMaxCharacters() int32`

GetMaxCharacters returns the MaxCharacters field if non-nil, zero value otherwise.

### GetMaxCharactersOk

`func (o *Definition) GetMaxCharactersOk() (*int32, bool)`

GetMaxCharactersOk returns a tuple with the MaxCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharacters

`func (o *Definition) SetMaxCharacters(v int32)`

SetMaxCharacters sets MaxCharacters field to given value.

### HasMaxCharacters

`func (o *Definition) HasMaxCharacters() bool`

HasMaxCharacters returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *Definition) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *Definition) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *Definition) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *Definition) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetBlockedDomains

`func (o *Definition) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *Definition) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *Definition) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.

### HasBlockedDomains

`func (o *Definition) HasBlockedDomains() bool`

HasBlockedDomains returns a boolean if a field has been set.

### GetMaxUses

`func (o *Definition) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *Definition) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *Definition) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *Definition) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetUserLocation

`func (o *Definition) GetUserLocation() MRBetaUserLocation`

GetUserLocation returns the UserLocation field if non-nil, zero value otherwise.

### GetUserLocationOk

`func (o *Definition) GetUserLocationOk() (*MRBetaUserLocation, bool)`

GetUserLocationOk returns a tuple with the UserLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocation

`func (o *Definition) SetUserLocation(v MRBetaUserLocation)`

SetUserLocation sets UserLocation field to given value.

### HasUserLocation

`func (o *Definition) HasUserLocation() bool`

HasUserLocation returns a boolean if a field has been set.

### GetCitations

`func (o *Definition) GetCitations() MRBetaCitationsConfigParam`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *Definition) GetCitationsOk() (*MRBetaCitationsConfigParam, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *Definition) SetCitations(v MRBetaCitationsConfigParam)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *Definition) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetMaxContentTokens

`func (o *Definition) GetMaxContentTokens() int32`

GetMaxContentTokens returns the MaxContentTokens field if non-nil, zero value otherwise.

### GetMaxContentTokensOk

`func (o *Definition) GetMaxContentTokensOk() (*int32, bool)`

GetMaxContentTokensOk returns a tuple with the MaxContentTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContentTokens

`func (o *Definition) SetMaxContentTokens(v int32)`

SetMaxContentTokens sets MaxContentTokens field to given value.

### HasMaxContentTokens

`func (o *Definition) HasMaxContentTokens() bool`

HasMaxContentTokens returns a boolean if a field has been set.

### GetUrlSources

`func (o *Definition) GetUrlSources() MRBetaWebFetchURLSources`

GetUrlSources returns the UrlSources field if non-nil, zero value otherwise.

### GetUrlSourcesOk

`func (o *Definition) GetUrlSourcesOk() (*MRBetaWebFetchURLSources, bool)`

GetUrlSourcesOk returns a tuple with the UrlSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSources

`func (o *Definition) SetUrlSources(v MRBetaWebFetchURLSources)`

SetUrlSources sets UrlSources field to given value.

### HasUrlSources

`func (o *Definition) HasUrlSources() bool`

HasUrlSources returns a boolean if a field has been set.

### GetUseCache

`func (o *Definition) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *Definition) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *Definition) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *Definition) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### GetResponseInclusion

`func (o *Definition) GetResponseInclusion() string`

GetResponseInclusion returns the ResponseInclusion field if non-nil, zero value otherwise.

### GetResponseInclusionOk

`func (o *Definition) GetResponseInclusionOk() (*string, bool)`

GetResponseInclusionOk returns a tuple with the ResponseInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseInclusion

`func (o *Definition) SetResponseInclusion(v string)`

SetResponseInclusion sets ResponseInclusion field to given value.

### HasResponseInclusion

`func (o *Definition) HasResponseInclusion() bool`

HasResponseInclusion returns a boolean if a field has been set.

### GetModel

`func (o *Definition) GetModel() Model1`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Definition) GetModelOk() (*Model1, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Definition) SetModel(v Model1)`

SetModel sets Model field to given value.


### GetCaching

`func (o *Definition) GetCaching() MRBetaCacheControlEphemeral`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *Definition) GetCachingOk() (*MRBetaCacheControlEphemeral, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *Definition) SetCaching(v MRBetaCacheControlEphemeral)`

SetCaching sets Caching field to given value.

### HasCaching

`func (o *Definition) HasCaching() bool`

HasCaching returns a boolean if a field has been set.

### GetMaxTokens

`func (o *Definition) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *Definition) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *Definition) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *Definition) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetMcpServerName

`func (o *Definition) GetMcpServerName() string`

GetMcpServerName returns the McpServerName field if non-nil, zero value otherwise.

### GetMcpServerNameOk

`func (o *Definition) GetMcpServerNameOk() (*string, bool)`

GetMcpServerNameOk returns a tuple with the McpServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerName

`func (o *Definition) SetMcpServerName(v string)`

SetMcpServerName sets McpServerName field to given value.


### GetDefaultConfig

`func (o *Definition) GetDefaultConfig() MRBetaMCPToolDefaultConfig`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *Definition) GetDefaultConfigOk() (*MRBetaMCPToolDefaultConfig, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *Definition) SetDefaultConfig(v MRBetaMCPToolDefaultConfig)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *Definition) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### GetTools

`func (o *Definition) GetTools() []MRBetaMCPToolParam`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Definition) GetToolsOk() (*[]MRBetaMCPToolParam, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Definition) SetTools(v []MRBetaMCPToolParam)`

SetTools sets Tools field to given value.

### HasTools

`func (o *Definition) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


