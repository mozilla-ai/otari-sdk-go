# ExternalUsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CacheTokensInPrompt** | Pointer to **bool** |  | [optional] [default to false]
**CacheWrite1hTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CacheWriteTokens** | Pointer to **int32** |  | [optional] [default to 0]
**DurationMs** | Pointer to **NullableInt32** |  | [optional] 
**InputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**Model** | **string** |  | 
**OutputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**Provider** | **string** |  | 
**SessionLabel** | Pointer to **NullableString** |  | [optional] 
**SourceEventId** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] [default to "success"]
**Timestamp** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExternalUsageEvent

`func NewExternalUsageEvent(model string, provider string, sourceEventId string, timestamp time.Time, ) *ExternalUsageEvent`

NewExternalUsageEvent instantiates a new ExternalUsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUsageEventWithDefaults

`func NewExternalUsageEventWithDefaults() *ExternalUsageEvent`

NewExternalUsageEventWithDefaults instantiates a new ExternalUsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadTokens

`func (o *ExternalUsageEvent) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *ExternalUsageEvent) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *ExternalUsageEvent) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.

### HasCacheReadTokens

`func (o *ExternalUsageEvent) HasCacheReadTokens() bool`

HasCacheReadTokens returns a boolean if a field has been set.

### GetCacheTokensInPrompt

`func (o *ExternalUsageEvent) GetCacheTokensInPrompt() bool`

GetCacheTokensInPrompt returns the CacheTokensInPrompt field if non-nil, zero value otherwise.

### GetCacheTokensInPromptOk

`func (o *ExternalUsageEvent) GetCacheTokensInPromptOk() (*bool, bool)`

GetCacheTokensInPromptOk returns a tuple with the CacheTokensInPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTokensInPrompt

`func (o *ExternalUsageEvent) SetCacheTokensInPrompt(v bool)`

SetCacheTokensInPrompt sets CacheTokensInPrompt field to given value.

### HasCacheTokensInPrompt

`func (o *ExternalUsageEvent) HasCacheTokensInPrompt() bool`

HasCacheTokensInPrompt returns a boolean if a field has been set.

### GetCacheWrite1hTokens

`func (o *ExternalUsageEvent) GetCacheWrite1hTokens() int32`

GetCacheWrite1hTokens returns the CacheWrite1hTokens field if non-nil, zero value otherwise.

### GetCacheWrite1hTokensOk

`func (o *ExternalUsageEvent) GetCacheWrite1hTokensOk() (*int32, bool)`

GetCacheWrite1hTokensOk returns a tuple with the CacheWrite1hTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hTokens

`func (o *ExternalUsageEvent) SetCacheWrite1hTokens(v int32)`

SetCacheWrite1hTokens sets CacheWrite1hTokens field to given value.

### HasCacheWrite1hTokens

`func (o *ExternalUsageEvent) HasCacheWrite1hTokens() bool`

HasCacheWrite1hTokens returns a boolean if a field has been set.

### GetCacheWriteTokens

`func (o *ExternalUsageEvent) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *ExternalUsageEvent) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *ExternalUsageEvent) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *ExternalUsageEvent) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### GetDurationMs

`func (o *ExternalUsageEvent) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ExternalUsageEvent) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ExternalUsageEvent) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ExternalUsageEvent) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *ExternalUsageEvent) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *ExternalUsageEvent) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetInputTokens

`func (o *ExternalUsageEvent) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *ExternalUsageEvent) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *ExternalUsageEvent) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *ExternalUsageEvent) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetModel

`func (o *ExternalUsageEvent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ExternalUsageEvent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ExternalUsageEvent) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputTokens

`func (o *ExternalUsageEvent) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *ExternalUsageEvent) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *ExternalUsageEvent) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *ExternalUsageEvent) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetProvider

`func (o *ExternalUsageEvent) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ExternalUsageEvent) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ExternalUsageEvent) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSessionLabel

`func (o *ExternalUsageEvent) GetSessionLabel() string`

GetSessionLabel returns the SessionLabel field if non-nil, zero value otherwise.

### GetSessionLabelOk

`func (o *ExternalUsageEvent) GetSessionLabelOk() (*string, bool)`

GetSessionLabelOk returns a tuple with the SessionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLabel

`func (o *ExternalUsageEvent) SetSessionLabel(v string)`

SetSessionLabel sets SessionLabel field to given value.

### HasSessionLabel

`func (o *ExternalUsageEvent) HasSessionLabel() bool`

HasSessionLabel returns a boolean if a field has been set.

### SetSessionLabelNil

`func (o *ExternalUsageEvent) SetSessionLabelNil(b bool)`

 SetSessionLabelNil sets the value for SessionLabel to be an explicit nil

### UnsetSessionLabel
`func (o *ExternalUsageEvent) UnsetSessionLabel()`

UnsetSessionLabel ensures that no value is present for SessionLabel, not even an explicit nil
### GetSourceEventId

`func (o *ExternalUsageEvent) GetSourceEventId() string`

GetSourceEventId returns the SourceEventId field if non-nil, zero value otherwise.

### GetSourceEventIdOk

`func (o *ExternalUsageEvent) GetSourceEventIdOk() (*string, bool)`

GetSourceEventIdOk returns a tuple with the SourceEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEventId

`func (o *ExternalUsageEvent) SetSourceEventId(v string)`

SetSourceEventId sets SourceEventId field to given value.


### GetStatus

`func (o *ExternalUsageEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalUsageEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalUsageEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalUsageEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExternalUsageEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExternalUsageEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExternalUsageEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUserId

`func (o *ExternalUsageEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalUsageEvent) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalUsageEvent) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExternalUsageEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExternalUsageEvent) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExternalUsageEvent) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


