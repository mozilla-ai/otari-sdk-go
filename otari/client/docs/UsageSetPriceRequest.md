# UsageSetPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to [**NullableApiKeyId**](ApiKeyId.md) |  | [optional] 
**ByFilter** | Pointer to **bool** |  | [optional] [default to false]
**CacheReadPricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**CacheWritePricePerMillion** | Pointer to **NullableFloat32** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Endpoint** | Pointer to **NullableString** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**InputPricePerMillion** | **float32** |  | 
**Model** | Pointer to [**NullableModel**](Model.md) |  | [optional] 
**OutputPricePerMillion** | **float32** |  | 
**Priced** | Pointer to **NullableBool** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceLabel** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Tool** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to [**NullableUserId**](UserId.md) |  | [optional] 
**WorkspaceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUsageSetPriceRequest

`func NewUsageSetPriceRequest(inputPricePerMillion float32, outputPricePerMillion float32, ) *UsageSetPriceRequest`

NewUsageSetPriceRequest instantiates a new UsageSetPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSetPriceRequestWithDefaults

`func NewUsageSetPriceRequestWithDefaults() *UsageSetPriceRequest`

NewUsageSetPriceRequestWithDefaults instantiates a new UsageSetPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *UsageSetPriceRequest) GetApiKeyId() ApiKeyId`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageSetPriceRequest) GetApiKeyIdOk() (*ApiKeyId, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageSetPriceRequest) SetApiKeyId(v ApiKeyId)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *UsageSetPriceRequest) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *UsageSetPriceRequest) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageSetPriceRequest) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetByFilter

`func (o *UsageSetPriceRequest) GetByFilter() bool`

GetByFilter returns the ByFilter field if non-nil, zero value otherwise.

### GetByFilterOk

`func (o *UsageSetPriceRequest) GetByFilterOk() (*bool, bool)`

GetByFilterOk returns a tuple with the ByFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFilter

`func (o *UsageSetPriceRequest) SetByFilter(v bool)`

SetByFilter sets ByFilter field to given value.

### HasByFilter

`func (o *UsageSetPriceRequest) HasByFilter() bool`

HasByFilter returns a boolean if a field has been set.

### GetCacheReadPricePerMillion

`func (o *UsageSetPriceRequest) GetCacheReadPricePerMillion() float32`

GetCacheReadPricePerMillion returns the CacheReadPricePerMillion field if non-nil, zero value otherwise.

### GetCacheReadPricePerMillionOk

`func (o *UsageSetPriceRequest) GetCacheReadPricePerMillionOk() (*float32, bool)`

GetCacheReadPricePerMillionOk returns a tuple with the CacheReadPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadPricePerMillion

`func (o *UsageSetPriceRequest) SetCacheReadPricePerMillion(v float32)`

SetCacheReadPricePerMillion sets CacheReadPricePerMillion field to given value.

### HasCacheReadPricePerMillion

`func (o *UsageSetPriceRequest) HasCacheReadPricePerMillion() bool`

HasCacheReadPricePerMillion returns a boolean if a field has been set.

### SetCacheReadPricePerMillionNil

`func (o *UsageSetPriceRequest) SetCacheReadPricePerMillionNil(b bool)`

 SetCacheReadPricePerMillionNil sets the value for CacheReadPricePerMillion to be an explicit nil

### UnsetCacheReadPricePerMillion
`func (o *UsageSetPriceRequest) UnsetCacheReadPricePerMillion()`

UnsetCacheReadPricePerMillion ensures that no value is present for CacheReadPricePerMillion, not even an explicit nil
### GetCacheWritePricePerMillion

`func (o *UsageSetPriceRequest) GetCacheWritePricePerMillion() float32`

GetCacheWritePricePerMillion returns the CacheWritePricePerMillion field if non-nil, zero value otherwise.

### GetCacheWritePricePerMillionOk

`func (o *UsageSetPriceRequest) GetCacheWritePricePerMillionOk() (*float32, bool)`

GetCacheWritePricePerMillionOk returns a tuple with the CacheWritePricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWritePricePerMillion

`func (o *UsageSetPriceRequest) SetCacheWritePricePerMillion(v float32)`

SetCacheWritePricePerMillion sets CacheWritePricePerMillion field to given value.

### HasCacheWritePricePerMillion

`func (o *UsageSetPriceRequest) HasCacheWritePricePerMillion() bool`

HasCacheWritePricePerMillion returns a boolean if a field has been set.

### SetCacheWritePricePerMillionNil

`func (o *UsageSetPriceRequest) SetCacheWritePricePerMillionNil(b bool)`

 SetCacheWritePricePerMillionNil sets the value for CacheWritePricePerMillion to be an explicit nil

### UnsetCacheWritePricePerMillion
`func (o *UsageSetPriceRequest) UnsetCacheWritePricePerMillion()`

UnsetCacheWritePricePerMillion ensures that no value is present for CacheWritePricePerMillion, not even an explicit nil
### GetEndDate

`func (o *UsageSetPriceRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageSetPriceRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageSetPriceRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageSetPriceRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UsageSetPriceRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UsageSetPriceRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetEndpoint

`func (o *UsageSetPriceRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageSetPriceRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageSetPriceRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UsageSetPriceRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *UsageSetPriceRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *UsageSetPriceRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIds

`func (o *UsageSetPriceRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UsageSetPriceRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UsageSetPriceRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UsageSetPriceRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *UsageSetPriceRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UsageSetPriceRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetInputPricePerMillion

`func (o *UsageSetPriceRequest) GetInputPricePerMillion() float32`

GetInputPricePerMillion returns the InputPricePerMillion field if non-nil, zero value otherwise.

### GetInputPricePerMillionOk

`func (o *UsageSetPriceRequest) GetInputPricePerMillionOk() (*float32, bool)`

GetInputPricePerMillionOk returns a tuple with the InputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerMillion

`func (o *UsageSetPriceRequest) SetInputPricePerMillion(v float32)`

SetInputPricePerMillion sets InputPricePerMillion field to given value.


### GetModel

`func (o *UsageSetPriceRequest) GetModel() Model`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UsageSetPriceRequest) GetModelOk() (*Model, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UsageSetPriceRequest) SetModel(v Model)`

SetModel sets Model field to given value.

### HasModel

`func (o *UsageSetPriceRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *UsageSetPriceRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *UsageSetPriceRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetOutputPricePerMillion

`func (o *UsageSetPriceRequest) GetOutputPricePerMillion() float32`

GetOutputPricePerMillion returns the OutputPricePerMillion field if non-nil, zero value otherwise.

### GetOutputPricePerMillionOk

`func (o *UsageSetPriceRequest) GetOutputPricePerMillionOk() (*float32, bool)`

GetOutputPricePerMillionOk returns a tuple with the OutputPricePerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerMillion

`func (o *UsageSetPriceRequest) SetOutputPricePerMillion(v float32)`

SetOutputPricePerMillion sets OutputPricePerMillion field to given value.


### GetPriced

`func (o *UsageSetPriceRequest) GetPriced() bool`

GetPriced returns the Priced field if non-nil, zero value otherwise.

### GetPricedOk

`func (o *UsageSetPriceRequest) GetPricedOk() (*bool, bool)`

GetPricedOk returns a tuple with the Priced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriced

`func (o *UsageSetPriceRequest) SetPriced(v bool)`

SetPriced sets Priced field to given value.

### HasPriced

`func (o *UsageSetPriceRequest) HasPriced() bool`

HasPriced returns a boolean if a field has been set.

### SetPricedNil

`func (o *UsageSetPriceRequest) SetPricedNil(b bool)`

 SetPricedNil sets the value for Priced to be an explicit nil

### UnsetPriced
`func (o *UsageSetPriceRequest) UnsetPriced()`

UnsetPriced ensures that no value is present for Priced, not even an explicit nil
### GetProvider

`func (o *UsageSetPriceRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageSetPriceRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageSetPriceRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageSetPriceRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UsageSetPriceRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UsageSetPriceRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetSource

`func (o *UsageSetPriceRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageSetPriceRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageSetPriceRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageSetPriceRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *UsageSetPriceRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *UsageSetPriceRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceLabel

`func (o *UsageSetPriceRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *UsageSetPriceRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *UsageSetPriceRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *UsageSetPriceRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### SetSourceLabelNil

`func (o *UsageSetPriceRequest) SetSourceLabelNil(b bool)`

 SetSourceLabelNil sets the value for SourceLabel to be an explicit nil

### UnsetSourceLabel
`func (o *UsageSetPriceRequest) UnsetSourceLabel()`

UnsetSourceLabel ensures that no value is present for SourceLabel, not even an explicit nil
### GetStartDate

`func (o *UsageSetPriceRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageSetPriceRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageSetPriceRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageSetPriceRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UsageSetPriceRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UsageSetPriceRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStatus

`func (o *UsageSetPriceRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageSetPriceRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageSetPriceRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsageSetPriceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UsageSetPriceRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UsageSetPriceRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTool

`func (o *UsageSetPriceRequest) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *UsageSetPriceRequest) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *UsageSetPriceRequest) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *UsageSetPriceRequest) HasTool() bool`

HasTool returns a boolean if a field has been set.

### SetToolNil

`func (o *UsageSetPriceRequest) SetToolNil(b bool)`

 SetToolNil sets the value for Tool to be an explicit nil

### UnsetTool
`func (o *UsageSetPriceRequest) UnsetTool()`

UnsetTool ensures that no value is present for Tool, not even an explicit nil
### GetUserId

`func (o *UsageSetPriceRequest) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageSetPriceRequest) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageSetPriceRequest) SetUserId(v UserId)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UsageSetPriceRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UsageSetPriceRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageSetPriceRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *UsageSetPriceRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UsageSetPriceRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UsageSetPriceRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UsageSetPriceRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *UsageSetPriceRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *UsageSetPriceRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


