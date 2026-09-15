# AgentTelemetryDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to [**NullableApiKeyId**](ApiKeyId.md) |  | [optional] 
**ByFilter** | Pointer to **bool** |  | [optional] [default to false]
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**UserId** | Pointer to [**NullableUserId**](UserId.md) |  | [optional] 

## Methods

### NewAgentTelemetryDeleteRequest

`func NewAgentTelemetryDeleteRequest() *AgentTelemetryDeleteRequest`

NewAgentTelemetryDeleteRequest instantiates a new AgentTelemetryDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryDeleteRequestWithDefaults

`func NewAgentTelemetryDeleteRequestWithDefaults() *AgentTelemetryDeleteRequest`

NewAgentTelemetryDeleteRequestWithDefaults instantiates a new AgentTelemetryDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *AgentTelemetryDeleteRequest) GetApiKeyId() ApiKeyId`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *AgentTelemetryDeleteRequest) GetApiKeyIdOk() (*ApiKeyId, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *AgentTelemetryDeleteRequest) SetApiKeyId(v ApiKeyId)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *AgentTelemetryDeleteRequest) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *AgentTelemetryDeleteRequest) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *AgentTelemetryDeleteRequest) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetByFilter

`func (o *AgentTelemetryDeleteRequest) GetByFilter() bool`

GetByFilter returns the ByFilter field if non-nil, zero value otherwise.

### GetByFilterOk

`func (o *AgentTelemetryDeleteRequest) GetByFilterOk() (*bool, bool)`

GetByFilterOk returns a tuple with the ByFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByFilter

`func (o *AgentTelemetryDeleteRequest) SetByFilter(v bool)`

SetByFilter sets ByFilter field to given value.

### HasByFilter

`func (o *AgentTelemetryDeleteRequest) HasByFilter() bool`

HasByFilter returns a boolean if a field has been set.

### GetEndDate

`func (o *AgentTelemetryDeleteRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AgentTelemetryDeleteRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AgentTelemetryDeleteRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AgentTelemetryDeleteRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *AgentTelemetryDeleteRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *AgentTelemetryDeleteRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetIds

`func (o *AgentTelemetryDeleteRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AgentTelemetryDeleteRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AgentTelemetryDeleteRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AgentTelemetryDeleteRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *AgentTelemetryDeleteRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *AgentTelemetryDeleteRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetName

`func (o *AgentTelemetryDeleteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentTelemetryDeleteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentTelemetryDeleteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentTelemetryDeleteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentTelemetryDeleteRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentTelemetryDeleteRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDate

`func (o *AgentTelemetryDeleteRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AgentTelemetryDeleteRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AgentTelemetryDeleteRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AgentTelemetryDeleteRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *AgentTelemetryDeleteRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *AgentTelemetryDeleteRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetUserId

`func (o *AgentTelemetryDeleteRequest) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AgentTelemetryDeleteRequest) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AgentTelemetryDeleteRequest) SetUserId(v UserId)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AgentTelemetryDeleteRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AgentTelemetryDeleteRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AgentTelemetryDeleteRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


