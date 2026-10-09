# UpdateSearchToolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** | New API key. Omit to keep the existing one. Never returned. | [optional] 
**ExpectedUpdatedAt** | Pointer to **NullableString** | Optimistic concurrency: if set, the update 412s unless it matches the stored updated_at. | [optional] 
**FetchTool** | Pointer to **NullableString** | For a search instance: the fetch instance that enriches its results. Null clears it. | [optional] 
**Kind** | Pointer to **NullableString** | Accepted only when it matches the stored kind, which cannot change. | [optional] 
**Options** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewUpdateSearchToolRequest

`func NewUpdateSearchToolRequest() *UpdateSearchToolRequest`

NewUpdateSearchToolRequest instantiates a new UpdateSearchToolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSearchToolRequestWithDefaults

`func NewUpdateSearchToolRequestWithDefaults() *UpdateSearchToolRequest`

NewUpdateSearchToolRequestWithDefaults instantiates a new UpdateSearchToolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *UpdateSearchToolRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *UpdateSearchToolRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *UpdateSearchToolRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *UpdateSearchToolRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *UpdateSearchToolRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *UpdateSearchToolRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *UpdateSearchToolRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateSearchToolRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateSearchToolRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateSearchToolRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *UpdateSearchToolRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *UpdateSearchToolRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetExpectedUpdatedAt

`func (o *UpdateSearchToolRequest) GetExpectedUpdatedAt() string`

GetExpectedUpdatedAt returns the ExpectedUpdatedAt field if non-nil, zero value otherwise.

### GetExpectedUpdatedAtOk

`func (o *UpdateSearchToolRequest) GetExpectedUpdatedAtOk() (*string, bool)`

GetExpectedUpdatedAtOk returns a tuple with the ExpectedUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUpdatedAt

`func (o *UpdateSearchToolRequest) SetExpectedUpdatedAt(v string)`

SetExpectedUpdatedAt sets ExpectedUpdatedAt field to given value.

### HasExpectedUpdatedAt

`func (o *UpdateSearchToolRequest) HasExpectedUpdatedAt() bool`

HasExpectedUpdatedAt returns a boolean if a field has been set.

### SetExpectedUpdatedAtNil

`func (o *UpdateSearchToolRequest) SetExpectedUpdatedAtNil(b bool)`

 SetExpectedUpdatedAtNil sets the value for ExpectedUpdatedAt to be an explicit nil

### UnsetExpectedUpdatedAt
`func (o *UpdateSearchToolRequest) UnsetExpectedUpdatedAt()`

UnsetExpectedUpdatedAt ensures that no value is present for ExpectedUpdatedAt, not even an explicit nil
### GetFetchTool

`func (o *UpdateSearchToolRequest) GetFetchTool() string`

GetFetchTool returns the FetchTool field if non-nil, zero value otherwise.

### GetFetchToolOk

`func (o *UpdateSearchToolRequest) GetFetchToolOk() (*string, bool)`

GetFetchToolOk returns a tuple with the FetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTool

`func (o *UpdateSearchToolRequest) SetFetchTool(v string)`

SetFetchTool sets FetchTool field to given value.

### HasFetchTool

`func (o *UpdateSearchToolRequest) HasFetchTool() bool`

HasFetchTool returns a boolean if a field has been set.

### SetFetchToolNil

`func (o *UpdateSearchToolRequest) SetFetchToolNil(b bool)`

 SetFetchToolNil sets the value for FetchTool to be an explicit nil

### UnsetFetchTool
`func (o *UpdateSearchToolRequest) UnsetFetchTool()`

UnsetFetchTool ensures that no value is present for FetchTool, not even an explicit nil
### GetKind

`func (o *UpdateSearchToolRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateSearchToolRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateSearchToolRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateSearchToolRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *UpdateSearchToolRequest) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *UpdateSearchToolRequest) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetOptions

`func (o *UpdateSearchToolRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateSearchToolRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateSearchToolRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateSearchToolRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateSearchToolRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateSearchToolRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProvider

`func (o *UpdateSearchToolRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UpdateSearchToolRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UpdateSearchToolRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UpdateSearchToolRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UpdateSearchToolRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UpdateSearchToolRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetTimeout

`func (o *UpdateSearchToolRequest) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateSearchToolRequest) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateSearchToolRequest) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateSearchToolRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *UpdateSearchToolRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *UpdateSearchToolRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


