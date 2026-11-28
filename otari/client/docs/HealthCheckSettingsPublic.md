# HealthCheckSettingsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmails** | **[]string** |  | 
**IntervalMinutes** | **int32** |  | 
**MaxTokens** | **int32** |  | 
**NextRunAt** | **NullableTime** |  | 
**Prompt** | **string** |  | 
**ReasoningEveryRuns** | **int32** |  | 
**ReasoningMaxTokens** | **int32** |  | 
**RunRequested** | **bool** |  | 
**ScheduleEnabled** | **bool** |  | 

## Methods

### NewHealthCheckSettingsPublic

`func NewHealthCheckSettingsPublic(alertEmails []string, intervalMinutes int32, maxTokens int32, nextRunAt NullableTime, prompt string, reasoningEveryRuns int32, reasoningMaxTokens int32, runRequested bool, scheduleEnabled bool, ) *HealthCheckSettingsPublic`

NewHealthCheckSettingsPublic instantiates a new HealthCheckSettingsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckSettingsPublicWithDefaults

`func NewHealthCheckSettingsPublicWithDefaults() *HealthCheckSettingsPublic`

NewHealthCheckSettingsPublicWithDefaults instantiates a new HealthCheckSettingsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmails

`func (o *HealthCheckSettingsPublic) GetAlertEmails() []string`

GetAlertEmails returns the AlertEmails field if non-nil, zero value otherwise.

### GetAlertEmailsOk

`func (o *HealthCheckSettingsPublic) GetAlertEmailsOk() (*[]string, bool)`

GetAlertEmailsOk returns a tuple with the AlertEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmails

`func (o *HealthCheckSettingsPublic) SetAlertEmails(v []string)`

SetAlertEmails sets AlertEmails field to given value.


### GetIntervalMinutes

`func (o *HealthCheckSettingsPublic) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckSettingsPublic) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckSettingsPublic) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.


### GetMaxTokens

`func (o *HealthCheckSettingsPublic) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *HealthCheckSettingsPublic) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *HealthCheckSettingsPublic) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.


### GetNextRunAt

`func (o *HealthCheckSettingsPublic) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *HealthCheckSettingsPublic) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *HealthCheckSettingsPublic) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.


### SetNextRunAtNil

`func (o *HealthCheckSettingsPublic) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *HealthCheckSettingsPublic) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetPrompt

`func (o *HealthCheckSettingsPublic) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *HealthCheckSettingsPublic) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *HealthCheckSettingsPublic) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetReasoningEveryRuns

`func (o *HealthCheckSettingsPublic) GetReasoningEveryRuns() int32`

GetReasoningEveryRuns returns the ReasoningEveryRuns field if non-nil, zero value otherwise.

### GetReasoningEveryRunsOk

`func (o *HealthCheckSettingsPublic) GetReasoningEveryRunsOk() (*int32, bool)`

GetReasoningEveryRunsOk returns a tuple with the ReasoningEveryRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEveryRuns

`func (o *HealthCheckSettingsPublic) SetReasoningEveryRuns(v int32)`

SetReasoningEveryRuns sets ReasoningEveryRuns field to given value.


### GetReasoningMaxTokens

`func (o *HealthCheckSettingsPublic) GetReasoningMaxTokens() int32`

GetReasoningMaxTokens returns the ReasoningMaxTokens field if non-nil, zero value otherwise.

### GetReasoningMaxTokensOk

`func (o *HealthCheckSettingsPublic) GetReasoningMaxTokensOk() (*int32, bool)`

GetReasoningMaxTokensOk returns a tuple with the ReasoningMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningMaxTokens

`func (o *HealthCheckSettingsPublic) SetReasoningMaxTokens(v int32)`

SetReasoningMaxTokens sets ReasoningMaxTokens field to given value.


### GetRunRequested

`func (o *HealthCheckSettingsPublic) GetRunRequested() bool`

GetRunRequested returns the RunRequested field if non-nil, zero value otherwise.

### GetRunRequestedOk

`func (o *HealthCheckSettingsPublic) GetRunRequestedOk() (*bool, bool)`

GetRunRequestedOk returns a tuple with the RunRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunRequested

`func (o *HealthCheckSettingsPublic) SetRunRequested(v bool)`

SetRunRequested sets RunRequested field to given value.


### GetScheduleEnabled

`func (o *HealthCheckSettingsPublic) GetScheduleEnabled() bool`

GetScheduleEnabled returns the ScheduleEnabled field if non-nil, zero value otherwise.

### GetScheduleEnabledOk

`func (o *HealthCheckSettingsPublic) GetScheduleEnabledOk() (*bool, bool)`

GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnabled

`func (o *HealthCheckSettingsPublic) SetScheduleEnabled(v bool)`

SetScheduleEnabled sets ScheduleEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


