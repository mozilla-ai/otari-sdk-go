# ListedSearchToolSchema

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
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**PricePerCall** | **NullableFloat32** | USD for one call on the deployment&#39;s own account when the provider reports no cost of its own, from the deployment&#39;s price list. An override the operator set for one organization prices that organization&#39;s calls instead. Null when no row prices the tool: with require_pricing on, its calls are refused. | 
**PriceSource** | **string** | Which row price_per_call comes from: &#39;own&#39;, the tool&#39;s own row (its kind and name as model_key); &#39;config&#39;, that row while the tool&#39;s config.yml entry sets price_per_call, which each start writes to it again when the rate in force differs, so a rate set through the API lasts until the next start; &#39;old_key&#39;, the tool&#39;s old &lt;provider&gt;:&lt;name&gt; row, which prices it while it has no row of its own; &#39;none&#39;, no row. | 
**Provider** | **string** |  | 
**ShadowsConfig** | Pointer to **bool** | True when a config-file search tool of the same name exists; the stored one is in effect. | [optional] [default to false]
**Timeout** | Pointer to **NullableFloat32** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListedSearchToolSchema

`func NewListedSearchToolSchema(name string, pricePerCall NullableFloat32, priceSource string, provider string, ) *ListedSearchToolSchema`

NewListedSearchToolSchema instantiates a new ListedSearchToolSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedSearchToolSchemaWithDefaults

`func NewListedSearchToolSchemaWithDefaults() *ListedSearchToolSchema`

NewListedSearchToolSchemaWithDefaults instantiates a new ListedSearchToolSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ListedSearchToolSchema) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ListedSearchToolSchema) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ListedSearchToolSchema) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *ListedSearchToolSchema) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *ListedSearchToolSchema) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *ListedSearchToolSchema) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetCreatedAt

`func (o *ListedSearchToolSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListedSearchToolSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListedSearchToolSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListedSearchToolSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ListedSearchToolSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ListedSearchToolSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDecryptable

`func (o *ListedSearchToolSchema) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *ListedSearchToolSchema) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *ListedSearchToolSchema) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *ListedSearchToolSchema) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetFetchTool

`func (o *ListedSearchToolSchema) GetFetchTool() string`

GetFetchTool returns the FetchTool field if non-nil, zero value otherwise.

### GetFetchToolOk

`func (o *ListedSearchToolSchema) GetFetchToolOk() (*string, bool)`

GetFetchToolOk returns a tuple with the FetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTool

`func (o *ListedSearchToolSchema) SetFetchTool(v string)`

SetFetchTool sets FetchTool field to given value.

### HasFetchTool

`func (o *ListedSearchToolSchema) HasFetchTool() bool`

HasFetchTool returns a boolean if a field has been set.

### SetFetchToolNil

`func (o *ListedSearchToolSchema) SetFetchToolNil(b bool)`

 SetFetchToolNil sets the value for FetchTool to be an explicit nil

### UnsetFetchTool
`func (o *ListedSearchToolSchema) UnsetFetchTool()`

UnsetFetchTool ensures that no value is present for FetchTool, not even an explicit nil
### GetKind

`func (o *ListedSearchToolSchema) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListedSearchToolSchema) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListedSearchToolSchema) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ListedSearchToolSchema) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLast4

`func (o *ListedSearchToolSchema) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *ListedSearchToolSchema) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *ListedSearchToolSchema) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *ListedSearchToolSchema) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *ListedSearchToolSchema) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *ListedSearchToolSchema) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *ListedSearchToolSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListedSearchToolSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListedSearchToolSchema) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *ListedSearchToolSchema) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ListedSearchToolSchema) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ListedSearchToolSchema) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ListedSearchToolSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPricePerCall

`func (o *ListedSearchToolSchema) GetPricePerCall() float32`

GetPricePerCall returns the PricePerCall field if non-nil, zero value otherwise.

### GetPricePerCallOk

`func (o *ListedSearchToolSchema) GetPricePerCallOk() (*float32, bool)`

GetPricePerCallOk returns a tuple with the PricePerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCall

`func (o *ListedSearchToolSchema) SetPricePerCall(v float32)`

SetPricePerCall sets PricePerCall field to given value.


### SetPricePerCallNil

`func (o *ListedSearchToolSchema) SetPricePerCallNil(b bool)`

 SetPricePerCallNil sets the value for PricePerCall to be an explicit nil

### UnsetPricePerCall
`func (o *ListedSearchToolSchema) UnsetPricePerCall()`

UnsetPricePerCall ensures that no value is present for PricePerCall, not even an explicit nil
### GetPriceSource

`func (o *ListedSearchToolSchema) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *ListedSearchToolSchema) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *ListedSearchToolSchema) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.


### GetProvider

`func (o *ListedSearchToolSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListedSearchToolSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListedSearchToolSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetShadowsConfig

`func (o *ListedSearchToolSchema) GetShadowsConfig() bool`

GetShadowsConfig returns the ShadowsConfig field if non-nil, zero value otherwise.

### GetShadowsConfigOk

`func (o *ListedSearchToolSchema) GetShadowsConfigOk() (*bool, bool)`

GetShadowsConfigOk returns a tuple with the ShadowsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowsConfig

`func (o *ListedSearchToolSchema) SetShadowsConfig(v bool)`

SetShadowsConfig sets ShadowsConfig field to given value.

### HasShadowsConfig

`func (o *ListedSearchToolSchema) HasShadowsConfig() bool`

HasShadowsConfig returns a boolean if a field has been set.

### GetTimeout

`func (o *ListedSearchToolSchema) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ListedSearchToolSchema) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ListedSearchToolSchema) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ListedSearchToolSchema) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *ListedSearchToolSchema) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *ListedSearchToolSchema) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUpdatedAt

`func (o *ListedSearchToolSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListedSearchToolSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListedSearchToolSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListedSearchToolSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ListedSearchToolSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ListedSearchToolSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


