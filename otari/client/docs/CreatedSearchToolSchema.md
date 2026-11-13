# CreatedSearchToolSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Decryptable** | Pointer to **bool** |  | [optional] [default to true]
**FetchTool** | Pointer to **NullableString** | A search instance&#39;s enrichment fetch instance. Null means the fetch default enriches it. | [optional] 
**Kind** | Pointer to **string** | Whether this is a search or a fetch instance. | [optional] [default to "search"]
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Notice** | Pointer to **NullableString** | What else the create changed, for the operator to read. | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**PinnedWebSearchDefaultTool** | Pointer to **NullableString** | Set when this create also stored web_search_default_tool, naming the search instance that was the in-loop default because it was the only one, so that adding a second does not turn in-loop search off. This runtime value wins over the configuration file until it is cleared. | [optional] 
**Provider** | **string** |  | 
**ShadowsConfig** | Pointer to **bool** | True when a config-file search tool of the same name exists; the stored one is in effect. | [optional] [default to false]
**Timeout** | Pointer to **NullableFloat32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreatedSearchToolSchema

`func NewCreatedSearchToolSchema(name string, provider string, ) *CreatedSearchToolSchema`

NewCreatedSearchToolSchema instantiates a new CreatedSearchToolSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedSearchToolSchemaWithDefaults

`func NewCreatedSearchToolSchemaWithDefaults() *CreatedSearchToolSchema`

NewCreatedSearchToolSchemaWithDefaults instantiates a new CreatedSearchToolSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *CreatedSearchToolSchema) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *CreatedSearchToolSchema) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *CreatedSearchToolSchema) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *CreatedSearchToolSchema) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *CreatedSearchToolSchema) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *CreatedSearchToolSchema) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetCreatedAt

`func (o *CreatedSearchToolSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatedSearchToolSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatedSearchToolSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreatedSearchToolSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CreatedSearchToolSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CreatedSearchToolSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDecryptable

`func (o *CreatedSearchToolSchema) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *CreatedSearchToolSchema) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *CreatedSearchToolSchema) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *CreatedSearchToolSchema) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetFetchTool

`func (o *CreatedSearchToolSchema) GetFetchTool() string`

GetFetchTool returns the FetchTool field if non-nil, zero value otherwise.

### GetFetchToolOk

`func (o *CreatedSearchToolSchema) GetFetchToolOk() (*string, bool)`

GetFetchToolOk returns a tuple with the FetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTool

`func (o *CreatedSearchToolSchema) SetFetchTool(v string)`

SetFetchTool sets FetchTool field to given value.

### HasFetchTool

`func (o *CreatedSearchToolSchema) HasFetchTool() bool`

HasFetchTool returns a boolean if a field has been set.

### SetFetchToolNil

`func (o *CreatedSearchToolSchema) SetFetchToolNil(b bool)`

 SetFetchToolNil sets the value for FetchTool to be an explicit nil

### UnsetFetchTool
`func (o *CreatedSearchToolSchema) UnsetFetchTool()`

UnsetFetchTool ensures that no value is present for FetchTool, not even an explicit nil
### GetKind

`func (o *CreatedSearchToolSchema) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreatedSearchToolSchema) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreatedSearchToolSchema) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreatedSearchToolSchema) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLast4

`func (o *CreatedSearchToolSchema) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *CreatedSearchToolSchema) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *CreatedSearchToolSchema) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *CreatedSearchToolSchema) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *CreatedSearchToolSchema) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *CreatedSearchToolSchema) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *CreatedSearchToolSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedSearchToolSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedSearchToolSchema) SetName(v string)`

SetName sets Name field to given value.


### GetNotice

`func (o *CreatedSearchToolSchema) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *CreatedSearchToolSchema) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *CreatedSearchToolSchema) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *CreatedSearchToolSchema) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### SetNoticeNil

`func (o *CreatedSearchToolSchema) SetNoticeNil(b bool)`

 SetNoticeNil sets the value for Notice to be an explicit nil

### UnsetNotice
`func (o *CreatedSearchToolSchema) UnsetNotice()`

UnsetNotice ensures that no value is present for Notice, not even an explicit nil
### GetOptions

`func (o *CreatedSearchToolSchema) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreatedSearchToolSchema) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreatedSearchToolSchema) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreatedSearchToolSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPinnedWebSearchDefaultTool

`func (o *CreatedSearchToolSchema) GetPinnedWebSearchDefaultTool() string`

GetPinnedWebSearchDefaultTool returns the PinnedWebSearchDefaultTool field if non-nil, zero value otherwise.

### GetPinnedWebSearchDefaultToolOk

`func (o *CreatedSearchToolSchema) GetPinnedWebSearchDefaultToolOk() (*string, bool)`

GetPinnedWebSearchDefaultToolOk returns a tuple with the PinnedWebSearchDefaultTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedWebSearchDefaultTool

`func (o *CreatedSearchToolSchema) SetPinnedWebSearchDefaultTool(v string)`

SetPinnedWebSearchDefaultTool sets PinnedWebSearchDefaultTool field to given value.

### HasPinnedWebSearchDefaultTool

`func (o *CreatedSearchToolSchema) HasPinnedWebSearchDefaultTool() bool`

HasPinnedWebSearchDefaultTool returns a boolean if a field has been set.

### SetPinnedWebSearchDefaultToolNil

`func (o *CreatedSearchToolSchema) SetPinnedWebSearchDefaultToolNil(b bool)`

 SetPinnedWebSearchDefaultToolNil sets the value for PinnedWebSearchDefaultTool to be an explicit nil

### UnsetPinnedWebSearchDefaultTool
`func (o *CreatedSearchToolSchema) UnsetPinnedWebSearchDefaultTool()`

UnsetPinnedWebSearchDefaultTool ensures that no value is present for PinnedWebSearchDefaultTool, not even an explicit nil
### GetProvider

`func (o *CreatedSearchToolSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreatedSearchToolSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreatedSearchToolSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetShadowsConfig

`func (o *CreatedSearchToolSchema) GetShadowsConfig() bool`

GetShadowsConfig returns the ShadowsConfig field if non-nil, zero value otherwise.

### GetShadowsConfigOk

`func (o *CreatedSearchToolSchema) GetShadowsConfigOk() (*bool, bool)`

GetShadowsConfigOk returns a tuple with the ShadowsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowsConfig

`func (o *CreatedSearchToolSchema) SetShadowsConfig(v bool)`

SetShadowsConfig sets ShadowsConfig field to given value.

### HasShadowsConfig

`func (o *CreatedSearchToolSchema) HasShadowsConfig() bool`

HasShadowsConfig returns a boolean if a field has been set.

### GetTimeout

`func (o *CreatedSearchToolSchema) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreatedSearchToolSchema) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreatedSearchToolSchema) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CreatedSearchToolSchema) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CreatedSearchToolSchema) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CreatedSearchToolSchema) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUpdatedAt

`func (o *CreatedSearchToolSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreatedSearchToolSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreatedSearchToolSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreatedSearchToolSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CreatedSearchToolSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CreatedSearchToolSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


