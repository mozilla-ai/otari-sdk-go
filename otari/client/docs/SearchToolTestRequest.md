# SearchToolTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** | Backend endpoint. Omit to inherit the provider&#39;s default (searxng inherits web_search_url). | [optional] 
**ApiKey** | Pointer to **NullableString** | Provider API key. Stored encrypted; never returned. | [optional] 
**FetchTool** | Pointer to **NullableString** | For a search instance: the fetch instance that enriches its results, a configured or stored one or builtin_fetch. Omit it for the fetch default. | [optional] 
**Kind** | Pointer to **string** | &#39;search&#39; or &#39;fetch&#39;. It cannot change once created. | [optional] [default to "search"]
**Name** | **string** | Name callers pass as &#39;search_tool_name&#39; or in /api/v1/search/{tool}, or that names a fetch instance. It contains no &#39;/&#39; or &#39;:&#39;, is not builtin_fetch or none in any case, and is unique across search and fetch instances. | 
**Options** | Pointer to **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | [optional] 
**Provider** | **string** | Provider id. GET /api/v1/search-tools/providers lists the search providers, and with ?kind&#x3D;fetch the fetch providers. | 
**Query** | Pointer to **NullableString** | For a search instance: the query to run. | [optional] 
**Timeout** | Pointer to **NullableFloat32** | Per-request timeout in seconds. | [optional] 
**Url** | Pointer to **NullableString** | For a fetch instance: the page to fetch. | [optional] 

## Methods

### NewSearchToolTestRequest

`func NewSearchToolTestRequest(name string, provider string, ) *SearchToolTestRequest`

NewSearchToolTestRequest instantiates a new SearchToolTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchToolTestRequestWithDefaults

`func NewSearchToolTestRequestWithDefaults() *SearchToolTestRequest`

NewSearchToolTestRequestWithDefaults instantiates a new SearchToolTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *SearchToolTestRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *SearchToolTestRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *SearchToolTestRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *SearchToolTestRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *SearchToolTestRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *SearchToolTestRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *SearchToolTestRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SearchToolTestRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SearchToolTestRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *SearchToolTestRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *SearchToolTestRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *SearchToolTestRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetFetchTool

`func (o *SearchToolTestRequest) GetFetchTool() string`

GetFetchTool returns the FetchTool field if non-nil, zero value otherwise.

### GetFetchToolOk

`func (o *SearchToolTestRequest) GetFetchToolOk() (*string, bool)`

GetFetchToolOk returns a tuple with the FetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTool

`func (o *SearchToolTestRequest) SetFetchTool(v string)`

SetFetchTool sets FetchTool field to given value.

### HasFetchTool

`func (o *SearchToolTestRequest) HasFetchTool() bool`

HasFetchTool returns a boolean if a field has been set.

### SetFetchToolNil

`func (o *SearchToolTestRequest) SetFetchToolNil(b bool)`

 SetFetchToolNil sets the value for FetchTool to be an explicit nil

### UnsetFetchTool
`func (o *SearchToolTestRequest) UnsetFetchTool()`

UnsetFetchTool ensures that no value is present for FetchTool, not even an explicit nil
### GetKind

`func (o *SearchToolTestRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SearchToolTestRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SearchToolTestRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SearchToolTestRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *SearchToolTestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchToolTestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchToolTestRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *SearchToolTestRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SearchToolTestRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SearchToolTestRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SearchToolTestRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *SearchToolTestRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *SearchToolTestRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProvider

`func (o *SearchToolTestRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SearchToolTestRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SearchToolTestRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetQuery

`func (o *SearchToolTestRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchToolTestRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchToolTestRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchToolTestRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *SearchToolTestRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *SearchToolTestRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetTimeout

`func (o *SearchToolTestRequest) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SearchToolTestRequest) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SearchToolTestRequest) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SearchToolTestRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *SearchToolTestRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *SearchToolTestRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUrl

`func (o *SearchToolTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchToolTestRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchToolTestRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SearchToolTestRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SearchToolTestRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SearchToolTestRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


