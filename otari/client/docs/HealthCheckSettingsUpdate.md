# HealthCheckSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEmails** | Pointer to **interface{}** |  | [optional] 
**IntervalMinutes** | Pointer to **NullableInt32** | Minutes between scheduled runs. | [optional] 
**MaxTokens** | Pointer to **NullableInt32** | Output cap for a model that does not reason. | [optional] 
**Prompt** | Pointer to **NullableString** | The user message every model is sent. A check passes on any answer, whatever it says. | [optional] 
**ReasoningEveryRuns** | Pointer to **NullableInt32** | Include reasoning models every Nth run; 0 leaves them out. | [optional] 
**ReasoningMaxTokens** | Pointer to **NullableInt32** | Output cap for a reasoning model. | [optional] 
**ScheduleEnabled** | Pointer to **NullableBool** | Run the checks on a schedule. | [optional] 

## Methods

### NewHealthCheckSettingsUpdate

`func NewHealthCheckSettingsUpdate() *HealthCheckSettingsUpdate`

NewHealthCheckSettingsUpdate instantiates a new HealthCheckSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckSettingsUpdateWithDefaults

`func NewHealthCheckSettingsUpdateWithDefaults() *HealthCheckSettingsUpdate`

NewHealthCheckSettingsUpdateWithDefaults instantiates a new HealthCheckSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEmails

`func (o *HealthCheckSettingsUpdate) GetAlertEmails() interface{}`

GetAlertEmails returns the AlertEmails field if non-nil, zero value otherwise.

### GetAlertEmailsOk

`func (o *HealthCheckSettingsUpdate) GetAlertEmailsOk() (*interface{}, bool)`

GetAlertEmailsOk returns a tuple with the AlertEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmails

`func (o *HealthCheckSettingsUpdate) SetAlertEmails(v interface{})`

SetAlertEmails sets AlertEmails field to given value.

### HasAlertEmails

`func (o *HealthCheckSettingsUpdate) HasAlertEmails() bool`

HasAlertEmails returns a boolean if a field has been set.

### SetAlertEmailsNil

`func (o *HealthCheckSettingsUpdate) SetAlertEmailsNil(b bool)`

 SetAlertEmailsNil sets the value for AlertEmails to be an explicit nil

### UnsetAlertEmails
`func (o *HealthCheckSettingsUpdate) UnsetAlertEmails()`

UnsetAlertEmails ensures that no value is present for AlertEmails, not even an explicit nil
### GetIntervalMinutes

`func (o *HealthCheckSettingsUpdate) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckSettingsUpdate) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckSettingsUpdate) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.

### HasIntervalMinutes

`func (o *HealthCheckSettingsUpdate) HasIntervalMinutes() bool`

HasIntervalMinutes returns a boolean if a field has been set.

### SetIntervalMinutesNil

`func (o *HealthCheckSettingsUpdate) SetIntervalMinutesNil(b bool)`

 SetIntervalMinutesNil sets the value for IntervalMinutes to be an explicit nil

### UnsetIntervalMinutes
`func (o *HealthCheckSettingsUpdate) UnsetIntervalMinutes()`

UnsetIntervalMinutes ensures that no value is present for IntervalMinutes, not even an explicit nil
### GetMaxTokens

`func (o *HealthCheckSettingsUpdate) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *HealthCheckSettingsUpdate) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *HealthCheckSettingsUpdate) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *HealthCheckSettingsUpdate) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *HealthCheckSettingsUpdate) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *HealthCheckSettingsUpdate) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetPrompt

`func (o *HealthCheckSettingsUpdate) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *HealthCheckSettingsUpdate) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *HealthCheckSettingsUpdate) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *HealthCheckSettingsUpdate) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *HealthCheckSettingsUpdate) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *HealthCheckSettingsUpdate) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetReasoningEveryRuns

`func (o *HealthCheckSettingsUpdate) GetReasoningEveryRuns() int32`

GetReasoningEveryRuns returns the ReasoningEveryRuns field if non-nil, zero value otherwise.

### GetReasoningEveryRunsOk

`func (o *HealthCheckSettingsUpdate) GetReasoningEveryRunsOk() (*int32, bool)`

GetReasoningEveryRunsOk returns a tuple with the ReasoningEveryRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEveryRuns

`func (o *HealthCheckSettingsUpdate) SetReasoningEveryRuns(v int32)`

SetReasoningEveryRuns sets ReasoningEveryRuns field to given value.

### HasReasoningEveryRuns

`func (o *HealthCheckSettingsUpdate) HasReasoningEveryRuns() bool`

HasReasoningEveryRuns returns a boolean if a field has been set.

### SetReasoningEveryRunsNil

`func (o *HealthCheckSettingsUpdate) SetReasoningEveryRunsNil(b bool)`

 SetReasoningEveryRunsNil sets the value for ReasoningEveryRuns to be an explicit nil

### UnsetReasoningEveryRuns
`func (o *HealthCheckSettingsUpdate) UnsetReasoningEveryRuns()`

UnsetReasoningEveryRuns ensures that no value is present for ReasoningEveryRuns, not even an explicit nil
### GetReasoningMaxTokens

`func (o *HealthCheckSettingsUpdate) GetReasoningMaxTokens() int32`

GetReasoningMaxTokens returns the ReasoningMaxTokens field if non-nil, zero value otherwise.

### GetReasoningMaxTokensOk

`func (o *HealthCheckSettingsUpdate) GetReasoningMaxTokensOk() (*int32, bool)`

GetReasoningMaxTokensOk returns a tuple with the ReasoningMaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningMaxTokens

`func (o *HealthCheckSettingsUpdate) SetReasoningMaxTokens(v int32)`

SetReasoningMaxTokens sets ReasoningMaxTokens field to given value.

### HasReasoningMaxTokens

`func (o *HealthCheckSettingsUpdate) HasReasoningMaxTokens() bool`

HasReasoningMaxTokens returns a boolean if a field has been set.

### SetReasoningMaxTokensNil

`func (o *HealthCheckSettingsUpdate) SetReasoningMaxTokensNil(b bool)`

 SetReasoningMaxTokensNil sets the value for ReasoningMaxTokens to be an explicit nil

### UnsetReasoningMaxTokens
`func (o *HealthCheckSettingsUpdate) UnsetReasoningMaxTokens()`

UnsetReasoningMaxTokens ensures that no value is present for ReasoningMaxTokens, not even an explicit nil
### GetScheduleEnabled

`func (o *HealthCheckSettingsUpdate) GetScheduleEnabled() bool`

GetScheduleEnabled returns the ScheduleEnabled field if non-nil, zero value otherwise.

### GetScheduleEnabledOk

`func (o *HealthCheckSettingsUpdate) GetScheduleEnabledOk() (*bool, bool)`

GetScheduleEnabledOk returns a tuple with the ScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnabled

`func (o *HealthCheckSettingsUpdate) SetScheduleEnabled(v bool)`

SetScheduleEnabled sets ScheduleEnabled field to given value.

### HasScheduleEnabled

`func (o *HealthCheckSettingsUpdate) HasScheduleEnabled() bool`

HasScheduleEnabled returns a boolean if a field has been set.

### SetScheduleEnabledNil

`func (o *HealthCheckSettingsUpdate) SetScheduleEnabledNil(b bool)`

 SetScheduleEnabledNil sets the value for ScheduleEnabled to be an explicit nil

### UnsetScheduleEnabled
`func (o *HealthCheckSettingsUpdate) UnsetScheduleEnabled()`

UnsetScheduleEnabled ensures that no value is present for ScheduleEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


