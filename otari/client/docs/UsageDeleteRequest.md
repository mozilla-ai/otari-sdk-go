# UsageDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **NullableString** |  | [optional] 
**ByFilter** | Pointer to **bool** |  | [optional] [default to false]
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Endpoint** | Pointer to **NullableString** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Priced** | Pointer to **NullableBool** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceLabel** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Tool** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUsageDeleteRequest

`func NewUsageDeleteRequest() *UsageDeleteRequest`

NewUsageDeleteRequest instantiates a new UsageDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDeleteRequestWithDefaults

`func NewUsageDeleteRequestWithDefaults() *UsageDeleteRequest`

NewUsageDeleteRequestWithDefaults instantiates a new UsageDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *UsageDeleteRequest) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageDeleteRequest) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageDeleteRequest) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *UsageDeleteRequest) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *UsageDeleteRequest) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageDeleteRequest) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetByFilter

`func (o *UsageDeleteRequest) GetByFilter() bool`

GetByFilter returns the ByFilter field if non-nil, zero value otherwise.

### GetByFilterOk

`func (o *UsageDeleteRequest) GetByFilterOk() (*bool, bool)`

GetByFilterOk returns a tuple with the ByFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFilter

`func (o *UsageDeleteRequest) SetByFilter(v bool)`

SetByFilter sets ByFilter field to given value.

### HasByFilter

`func (o *UsageDeleteRequest) HasByFilter() bool`

HasByFilter returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageDeleteRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageDeleteRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageDeleteRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageDeleteRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *UsageDeleteRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *UsageDeleteRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetEndpoint

`func (o *UsageDeleteRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageDeleteRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageDeleteRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UsageDeleteRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *UsageDeleteRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *UsageDeleteRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIds

`func (o *UsageDeleteRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UsageDeleteRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UsageDeleteRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UsageDeleteRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *UsageDeleteRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UsageDeleteRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetModel

`func (o *UsageDeleteRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UsageDeleteRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UsageDeleteRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UsageDeleteRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *UsageDeleteRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *UsageDeleteRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetPriced

`func (o *UsageDeleteRequest) GetPriced() bool`

GetPriced returns the Priced field if non-nil, zero value otherwise.

### GetPricedOk

`func (o *UsageDeleteRequest) GetPricedOk() (*bool, bool)`

GetPricedOk returns a tuple with the Priced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriced

`func (o *UsageDeleteRequest) SetPriced(v bool)`

SetPriced sets Priced field to given value.

### HasPriced

`func (o *UsageDeleteRequest) HasPriced() bool`

HasPriced returns a boolean if a field has been set.

### SetPricedNil

`func (o *UsageDeleteRequest) SetPricedNil(b bool)`

 SetPricedNil sets the value for Priced to be an explicit nil

### UnsetPriced
`func (o *UsageDeleteRequest) UnsetPriced()`

UnsetPriced ensures that no value is present for Priced, not even an explicit nil
### GetProvider

`func (o *UsageDeleteRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageDeleteRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageDeleteRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UsageDeleteRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UsageDeleteRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UsageDeleteRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetSource

`func (o *UsageDeleteRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageDeleteRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageDeleteRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageDeleteRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *UsageDeleteRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *UsageDeleteRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceLabel

`func (o *UsageDeleteRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *UsageDeleteRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *UsageDeleteRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *UsageDeleteRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### SetSourceLabelNil

`func (o *UsageDeleteRequest) SetSourceLabelNil(b bool)`

 SetSourceLabelNil sets the value for SourceLabel to be an explicit nil

### UnsetSourceLabel
`func (o *UsageDeleteRequest) UnsetSourceLabel()`

UnsetSourceLabel ensures that no value is present for SourceLabel, not even an explicit nil
### GetStartDate

`func (o *UsageDeleteRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageDeleteRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageDeleteRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageDeleteRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *UsageDeleteRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *UsageDeleteRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStatus

`func (o *UsageDeleteRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageDeleteRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageDeleteRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsageDeleteRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UsageDeleteRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UsageDeleteRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTool

`func (o *UsageDeleteRequest) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *UsageDeleteRequest) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *UsageDeleteRequest) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *UsageDeleteRequest) HasTool() bool`

HasTool returns a boolean if a field has been set.

### SetToolNil

`func (o *UsageDeleteRequest) SetToolNil(b bool)`

 SetToolNil sets the value for Tool to be an explicit nil

### UnsetTool
`func (o *UsageDeleteRequest) UnsetTool()`

UnsetTool ensures that no value is present for Tool, not even an explicit nil
### GetUserId

`func (o *UsageDeleteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageDeleteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageDeleteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UsageDeleteRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UsageDeleteRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageDeleteRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


