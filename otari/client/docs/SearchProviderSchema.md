# SearchProviderSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultApiBase** | Pointer to **NullableString** | The endpoint a tool on this provider uses when it declares no api_base. Null means nothing supplies one, so an api_base is required. One that comes from the deployment&#39;s own settings, such as the web_search_url a searxng tool inherits, is shown only to a caller who operates the deployment: nothing else inherits it, an organization&#39;s key included. | [optional] 
**DocUrl** | Pointer to **NullableString** | The provider&#39;s API documentation. | [optional] 
**Formats** | Pointer to **[]string** | Fetch: the formats the page text comes back in. | [optional] 
**Id** | **string** | Value to send as &#39;provider&#39;. | 
**Instances** | Pointer to **[]string** | The names of this deployment&#39;s configured and stored tools on this provider, shown only to a caller who operates the deployment. | [optional] 
**KeyInUrl** | Pointer to **NullableBool** | Search: true when the API key travels in the request URL. | [optional] 
**Kind** | **string** | Whether this is a search provider or a fetch provider. | 
**MaxResults** | Pointer to **NullableInt32** | Search: the most results one call can ask for. | [optional] 
**MaxUrlsPerCall** | Pointer to **NullableInt32** | Fetch: the most pages one call can fetch. | [optional] 
**Options** | Pointer to [**[]SearchProviderOptionSchema**](SearchProviderOptionSchema.md) | The native options a tool may set, under the provider&#39;s own names. Null when there is no schema for the provider yet, so a tool&#39;s options are passed unchecked. | [optional] 
**QueryInUrl** | Pointer to **NullableBool** | Search: true when the query travels in the request URL. | [optional] 
**RendersJavascript** | Pointer to **NullableBool** | Fetch: true when the provider runs a page&#39;s JavaScript before reading it. | [optional] 
**RequiresApiBase** | **bool** | True when this provider has no endpoint of its own, so the tool must say where the backend is. | 
**RequiresApiKey** | **bool** | True when a tool on this provider must carry an API key. | 
**Tier** | Pointer to **NullableString** | The library&#39;s tier for the provider. | [optional] 

## Methods

### NewSearchProviderSchema

`func NewSearchProviderSchema(id string, kind string, requiresApiBase bool, requiresApiKey bool, ) *SearchProviderSchema`

NewSearchProviderSchema instantiates a new SearchProviderSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProviderSchemaWithDefaults

`func NewSearchProviderSchemaWithDefaults() *SearchProviderSchema`

NewSearchProviderSchemaWithDefaults instantiates a new SearchProviderSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultApiBase

`func (o *SearchProviderSchema) GetDefaultApiBase() string`

GetDefaultApiBase returns the DefaultApiBase field if non-nil, zero value otherwise.

### GetDefaultApiBaseOk

`func (o *SearchProviderSchema) GetDefaultApiBaseOk() (*string, bool)`

GetDefaultApiBaseOk returns a tuple with the DefaultApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiBase

`func (o *SearchProviderSchema) SetDefaultApiBase(v string)`

SetDefaultApiBase sets DefaultApiBase field to given value.

### HasDefaultApiBase

`func (o *SearchProviderSchema) HasDefaultApiBase() bool`

HasDefaultApiBase returns a boolean if a field has been set.

### SetDefaultApiBaseNil

`func (o *SearchProviderSchema) SetDefaultApiBaseNil(b bool)`

 SetDefaultApiBaseNil sets the value for DefaultApiBase to be an explicit nil

### UnsetDefaultApiBase
`func (o *SearchProviderSchema) UnsetDefaultApiBase()`

UnsetDefaultApiBase ensures that no value is present for DefaultApiBase, not even an explicit nil
### GetDocUrl

`func (o *SearchProviderSchema) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *SearchProviderSchema) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *SearchProviderSchema) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *SearchProviderSchema) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.

### SetDocUrlNil

`func (o *SearchProviderSchema) SetDocUrlNil(b bool)`

 SetDocUrlNil sets the value for DocUrl to be an explicit nil

### UnsetDocUrl
`func (o *SearchProviderSchema) UnsetDocUrl()`

UnsetDocUrl ensures that no value is present for DocUrl, not even an explicit nil
### GetFormats

`func (o *SearchProviderSchema) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *SearchProviderSchema) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *SearchProviderSchema) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *SearchProviderSchema) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### SetFormatsNil

`func (o *SearchProviderSchema) SetFormatsNil(b bool)`

 SetFormatsNil sets the value for Formats to be an explicit nil

### UnsetFormats
`func (o *SearchProviderSchema) UnsetFormats()`

UnsetFormats ensures that no value is present for Formats, not even an explicit nil
### GetId

`func (o *SearchProviderSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchProviderSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchProviderSchema) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *SearchProviderSchema) GetInstances() []string`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *SearchProviderSchema) GetInstancesOk() (*[]string, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *SearchProviderSchema) SetInstances(v []string)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *SearchProviderSchema) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetKeyInUrl

`func (o *SearchProviderSchema) GetKeyInUrl() bool`

GetKeyInUrl returns the KeyInUrl field if non-nil, zero value otherwise.

### GetKeyInUrlOk

`func (o *SearchProviderSchema) GetKeyInUrlOk() (*bool, bool)`

GetKeyInUrlOk returns a tuple with the KeyInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInUrl

`func (o *SearchProviderSchema) SetKeyInUrl(v bool)`

SetKeyInUrl sets KeyInUrl field to given value.

### HasKeyInUrl

`func (o *SearchProviderSchema) HasKeyInUrl() bool`

HasKeyInUrl returns a boolean if a field has been set.

### SetKeyInUrlNil

`func (o *SearchProviderSchema) SetKeyInUrlNil(b bool)`

 SetKeyInUrlNil sets the value for KeyInUrl to be an explicit nil

### UnsetKeyInUrl
`func (o *SearchProviderSchema) UnsetKeyInUrl()`

UnsetKeyInUrl ensures that no value is present for KeyInUrl, not even an explicit nil
### GetKind

`func (o *SearchProviderSchema) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SearchProviderSchema) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SearchProviderSchema) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaxResults

`func (o *SearchProviderSchema) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchProviderSchema) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchProviderSchema) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchProviderSchema) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### SetMaxResultsNil

`func (o *SearchProviderSchema) SetMaxResultsNil(b bool)`

 SetMaxResultsNil sets the value for MaxResults to be an explicit nil

### UnsetMaxResults
`func (o *SearchProviderSchema) UnsetMaxResults()`

UnsetMaxResults ensures that no value is present for MaxResults, not even an explicit nil
### GetMaxUrlsPerCall

`func (o *SearchProviderSchema) GetMaxUrlsPerCall() int32`

GetMaxUrlsPerCall returns the MaxUrlsPerCall field if non-nil, zero value otherwise.

### GetMaxUrlsPerCallOk

`func (o *SearchProviderSchema) GetMaxUrlsPerCallOk() (*int32, bool)`

GetMaxUrlsPerCallOk returns a tuple with the MaxUrlsPerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUrlsPerCall

`func (o *SearchProviderSchema) SetMaxUrlsPerCall(v int32)`

SetMaxUrlsPerCall sets MaxUrlsPerCall field to given value.

### HasMaxUrlsPerCall

`func (o *SearchProviderSchema) HasMaxUrlsPerCall() bool`

HasMaxUrlsPerCall returns a boolean if a field has been set.

### SetMaxUrlsPerCallNil

`func (o *SearchProviderSchema) SetMaxUrlsPerCallNil(b bool)`

 SetMaxUrlsPerCallNil sets the value for MaxUrlsPerCall to be an explicit nil

### UnsetMaxUrlsPerCall
`func (o *SearchProviderSchema) UnsetMaxUrlsPerCall()`

UnsetMaxUrlsPerCall ensures that no value is present for MaxUrlsPerCall, not even an explicit nil
### GetOptions

`func (o *SearchProviderSchema) GetOptions() []SearchProviderOptionSchema`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SearchProviderSchema) GetOptionsOk() (*[]SearchProviderOptionSchema, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SearchProviderSchema) SetOptions(v []SearchProviderOptionSchema)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SearchProviderSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SearchProviderSchema) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SearchProviderSchema) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetQueryInUrl

`func (o *SearchProviderSchema) GetQueryInUrl() bool`

GetQueryInUrl returns the QueryInUrl field if non-nil, zero value otherwise.

### GetQueryInUrlOk

`func (o *SearchProviderSchema) GetQueryInUrlOk() (*bool, bool)`

GetQueryInUrlOk returns a tuple with the QueryInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInUrl

`func (o *SearchProviderSchema) SetQueryInUrl(v bool)`

SetQueryInUrl sets QueryInUrl field to given value.

### HasQueryInUrl

`func (o *SearchProviderSchema) HasQueryInUrl() bool`

HasQueryInUrl returns a boolean if a field has been set.

### SetQueryInUrlNil

`func (o *SearchProviderSchema) SetQueryInUrlNil(b bool)`

 SetQueryInUrlNil sets the value for QueryInUrl to be an explicit nil

### UnsetQueryInUrl
`func (o *SearchProviderSchema) UnsetQueryInUrl()`

UnsetQueryInUrl ensures that no value is present for QueryInUrl, not even an explicit nil
### GetRendersJavascript

`func (o *SearchProviderSchema) GetRendersJavascript() bool`

GetRendersJavascript returns the RendersJavascript field if non-nil, zero value otherwise.

### GetRendersJavascriptOk

`func (o *SearchProviderSchema) GetRendersJavascriptOk() (*bool, bool)`

GetRendersJavascriptOk returns a tuple with the RendersJavascript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendersJavascript

`func (o *SearchProviderSchema) SetRendersJavascript(v bool)`

SetRendersJavascript sets RendersJavascript field to given value.

### HasRendersJavascript

`func (o *SearchProviderSchema) HasRendersJavascript() bool`

HasRendersJavascript returns a boolean if a field has been set.

### SetRendersJavascriptNil

`func (o *SearchProviderSchema) SetRendersJavascriptNil(b bool)`

 SetRendersJavascriptNil sets the value for RendersJavascript to be an explicit nil

### UnsetRendersJavascript
`func (o *SearchProviderSchema) UnsetRendersJavascript()`

UnsetRendersJavascript ensures that no value is present for RendersJavascript, not even an explicit nil
### GetRequiresApiBase

`func (o *SearchProviderSchema) GetRequiresApiBase() bool`

GetRequiresApiBase returns the RequiresApiBase field if non-nil, zero value otherwise.

### GetRequiresApiBaseOk

`func (o *SearchProviderSchema) GetRequiresApiBaseOk() (*bool, bool)`

GetRequiresApiBaseOk returns a tuple with the RequiresApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiBase

`func (o *SearchProviderSchema) SetRequiresApiBase(v bool)`

SetRequiresApiBase sets RequiresApiBase field to given value.


### GetRequiresApiKey

`func (o *SearchProviderSchema) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *SearchProviderSchema) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *SearchProviderSchema) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.


### GetTier

`func (o *SearchProviderSchema) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SearchProviderSchema) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SearchProviderSchema) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SearchProviderSchema) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *SearchProviderSchema) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *SearchProviderSchema) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


