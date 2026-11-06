# OrgWebSearchKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** | The provider&#39;s endpoint, when not its own. Must be https when the key carries an API key. | [optional] 
**ApiKey** | Pointer to **NullableString** | Required unless the provider takes no key. | [optional] 
**Kind** | Pointer to **string** | Whether the key serves the search tool or the fetch tool, and enrichment. Fixed once created. | [optional] [default to "search"]
**Name** | **string** | How the organization tells this key apart from its others. | 
**Options** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Provider** | **string** | The provider the key is for: one the provider catalog, GET /api/v1/search-tools/providers, lists for the key&#39;s kind, other than searxng, or tavily or brave for a search key. | 

## Methods

### NewOrgWebSearchKeyCreateRequest

`func NewOrgWebSearchKeyCreateRequest(name string, provider string, ) *OrgWebSearchKeyCreateRequest`

NewOrgWebSearchKeyCreateRequest instantiates a new OrgWebSearchKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWebSearchKeyCreateRequestWithDefaults

`func NewOrgWebSearchKeyCreateRequestWithDefaults() *OrgWebSearchKeyCreateRequest`

NewOrgWebSearchKeyCreateRequestWithDefaults instantiates a new OrgWebSearchKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *OrgWebSearchKeyCreateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *OrgWebSearchKeyCreateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *OrgWebSearchKeyCreateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *OrgWebSearchKeyCreateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *OrgWebSearchKeyCreateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *OrgWebSearchKeyCreateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *OrgWebSearchKeyCreateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OrgWebSearchKeyCreateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OrgWebSearchKeyCreateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OrgWebSearchKeyCreateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *OrgWebSearchKeyCreateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *OrgWebSearchKeyCreateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetKind

`func (o *OrgWebSearchKeyCreateRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgWebSearchKeyCreateRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgWebSearchKeyCreateRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrgWebSearchKeyCreateRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *OrgWebSearchKeyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgWebSearchKeyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgWebSearchKeyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *OrgWebSearchKeyCreateRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrgWebSearchKeyCreateRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrgWebSearchKeyCreateRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OrgWebSearchKeyCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *OrgWebSearchKeyCreateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *OrgWebSearchKeyCreateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProvider

`func (o *OrgWebSearchKeyCreateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrgWebSearchKeyCreateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrgWebSearchKeyCreateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


