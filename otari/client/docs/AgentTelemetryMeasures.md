# AgentTelemetryMeasures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostPerCommit** | Pointer to **NullableFloat32** |  | [optional] 
**CostPerLine** | Pointer to **NullableFloat32** |  | [optional] 
**CostPerPullRequest** | Pointer to **NullableFloat32** |  | [optional] 
**EditAcceptanceRate** | Pointer to **NullableFloat32** |  | [optional] 
**ErrorRate** | Pointer to **NullableFloat32** |  | [optional] 
**SpendPerActiveHour** | Pointer to **NullableFloat32** |  | [optional] 
**ToolAcceptanceRate** | Pointer to **NullableFloat32** |  | [optional] 
**TurnsPerSession** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewAgentTelemetryMeasures

`func NewAgentTelemetryMeasures() *AgentTelemetryMeasures`

NewAgentTelemetryMeasures instantiates a new AgentTelemetryMeasures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryMeasuresWithDefaults

`func NewAgentTelemetryMeasuresWithDefaults() *AgentTelemetryMeasures`

NewAgentTelemetryMeasuresWithDefaults instantiates a new AgentTelemetryMeasures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostPerCommit

`func (o *AgentTelemetryMeasures) GetCostPerCommit() float32`

GetCostPerCommit returns the CostPerCommit field if non-nil, zero value otherwise.

### GetCostPerCommitOk

`func (o *AgentTelemetryMeasures) GetCostPerCommitOk() (*float32, bool)`

GetCostPerCommitOk returns a tuple with the CostPerCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCommit

`func (o *AgentTelemetryMeasures) SetCostPerCommit(v float32)`

SetCostPerCommit sets CostPerCommit field to given value.

### HasCostPerCommit

`func (o *AgentTelemetryMeasures) HasCostPerCommit() bool`

HasCostPerCommit returns a boolean if a field has been set.

### SetCostPerCommitNil

`func (o *AgentTelemetryMeasures) SetCostPerCommitNil(b bool)`

 SetCostPerCommitNil sets the value for CostPerCommit to be an explicit nil

### UnsetCostPerCommit
`func (o *AgentTelemetryMeasures) UnsetCostPerCommit()`

UnsetCostPerCommit ensures that no value is present for CostPerCommit, not even an explicit nil
### GetCostPerLine

`func (o *AgentTelemetryMeasures) GetCostPerLine() float32`

GetCostPerLine returns the CostPerLine field if non-nil, zero value otherwise.

### GetCostPerLineOk

`func (o *AgentTelemetryMeasures) GetCostPerLineOk() (*float32, bool)`

GetCostPerLineOk returns a tuple with the CostPerLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerLine

`func (o *AgentTelemetryMeasures) SetCostPerLine(v float32)`

SetCostPerLine sets CostPerLine field to given value.

### HasCostPerLine

`func (o *AgentTelemetryMeasures) HasCostPerLine() bool`

HasCostPerLine returns a boolean if a field has been set.

### SetCostPerLineNil

`func (o *AgentTelemetryMeasures) SetCostPerLineNil(b bool)`

 SetCostPerLineNil sets the value for CostPerLine to be an explicit nil

### UnsetCostPerLine
`func (o *AgentTelemetryMeasures) UnsetCostPerLine()`

UnsetCostPerLine ensures that no value is present for CostPerLine, not even an explicit nil
### GetCostPerPullRequest

`func (o *AgentTelemetryMeasures) GetCostPerPullRequest() float32`

GetCostPerPullRequest returns the CostPerPullRequest field if non-nil, zero value otherwise.

### GetCostPerPullRequestOk

`func (o *AgentTelemetryMeasures) GetCostPerPullRequestOk() (*float32, bool)`

GetCostPerPullRequestOk returns a tuple with the CostPerPullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerPullRequest

`func (o *AgentTelemetryMeasures) SetCostPerPullRequest(v float32)`

SetCostPerPullRequest sets CostPerPullRequest field to given value.

### HasCostPerPullRequest

`func (o *AgentTelemetryMeasures) HasCostPerPullRequest() bool`

HasCostPerPullRequest returns a boolean if a field has been set.

### SetCostPerPullRequestNil

`func (o *AgentTelemetryMeasures) SetCostPerPullRequestNil(b bool)`

 SetCostPerPullRequestNil sets the value for CostPerPullRequest to be an explicit nil

### UnsetCostPerPullRequest
`func (o *AgentTelemetryMeasures) UnsetCostPerPullRequest()`

UnsetCostPerPullRequest ensures that no value is present for CostPerPullRequest, not even an explicit nil
### GetEditAcceptanceRate

`func (o *AgentTelemetryMeasures) GetEditAcceptanceRate() float32`

GetEditAcceptanceRate returns the EditAcceptanceRate field if non-nil, zero value otherwise.

### GetEditAcceptanceRateOk

`func (o *AgentTelemetryMeasures) GetEditAcceptanceRateOk() (*float32, bool)`

GetEditAcceptanceRateOk returns a tuple with the EditAcceptanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditAcceptanceRate

`func (o *AgentTelemetryMeasures) SetEditAcceptanceRate(v float32)`

SetEditAcceptanceRate sets EditAcceptanceRate field to given value.

### HasEditAcceptanceRate

`func (o *AgentTelemetryMeasures) HasEditAcceptanceRate() bool`

HasEditAcceptanceRate returns a boolean if a field has been set.

### SetEditAcceptanceRateNil

`func (o *AgentTelemetryMeasures) SetEditAcceptanceRateNil(b bool)`

 SetEditAcceptanceRateNil sets the value for EditAcceptanceRate to be an explicit nil

### UnsetEditAcceptanceRate
`func (o *AgentTelemetryMeasures) UnsetEditAcceptanceRate()`

UnsetEditAcceptanceRate ensures that no value is present for EditAcceptanceRate, not even an explicit nil
### GetErrorRate

`func (o *AgentTelemetryMeasures) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *AgentTelemetryMeasures) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *AgentTelemetryMeasures) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *AgentTelemetryMeasures) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### SetErrorRateNil

`func (o *AgentTelemetryMeasures) SetErrorRateNil(b bool)`

 SetErrorRateNil sets the value for ErrorRate to be an explicit nil

### UnsetErrorRate
`func (o *AgentTelemetryMeasures) UnsetErrorRate()`

UnsetErrorRate ensures that no value is present for ErrorRate, not even an explicit nil
### GetSpendPerActiveHour

`func (o *AgentTelemetryMeasures) GetSpendPerActiveHour() float32`

GetSpendPerActiveHour returns the SpendPerActiveHour field if non-nil, zero value otherwise.

### GetSpendPerActiveHourOk

`func (o *AgentTelemetryMeasures) GetSpendPerActiveHourOk() (*float32, bool)`

GetSpendPerActiveHourOk returns a tuple with the SpendPerActiveHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendPerActiveHour

`func (o *AgentTelemetryMeasures) SetSpendPerActiveHour(v float32)`

SetSpendPerActiveHour sets SpendPerActiveHour field to given value.

### HasSpendPerActiveHour

`func (o *AgentTelemetryMeasures) HasSpendPerActiveHour() bool`

HasSpendPerActiveHour returns a boolean if a field has been set.

### SetSpendPerActiveHourNil

`func (o *AgentTelemetryMeasures) SetSpendPerActiveHourNil(b bool)`

 SetSpendPerActiveHourNil sets the value for SpendPerActiveHour to be an explicit nil

### UnsetSpendPerActiveHour
`func (o *AgentTelemetryMeasures) UnsetSpendPerActiveHour()`

UnsetSpendPerActiveHour ensures that no value is present for SpendPerActiveHour, not even an explicit nil
### GetToolAcceptanceRate

`func (o *AgentTelemetryMeasures) GetToolAcceptanceRate() float32`

GetToolAcceptanceRate returns the ToolAcceptanceRate field if non-nil, zero value otherwise.

### GetToolAcceptanceRateOk

`func (o *AgentTelemetryMeasures) GetToolAcceptanceRateOk() (*float32, bool)`

GetToolAcceptanceRateOk returns a tuple with the ToolAcceptanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolAcceptanceRate

`func (o *AgentTelemetryMeasures) SetToolAcceptanceRate(v float32)`

SetToolAcceptanceRate sets ToolAcceptanceRate field to given value.

### HasToolAcceptanceRate

`func (o *AgentTelemetryMeasures) HasToolAcceptanceRate() bool`

HasToolAcceptanceRate returns a boolean if a field has been set.

### SetToolAcceptanceRateNil

`func (o *AgentTelemetryMeasures) SetToolAcceptanceRateNil(b bool)`

 SetToolAcceptanceRateNil sets the value for ToolAcceptanceRate to be an explicit nil

### UnsetToolAcceptanceRate
`func (o *AgentTelemetryMeasures) UnsetToolAcceptanceRate()`

UnsetToolAcceptanceRate ensures that no value is present for ToolAcceptanceRate, not even an explicit nil
### GetTurnsPerSession

`func (o *AgentTelemetryMeasures) GetTurnsPerSession() float32`

GetTurnsPerSession returns the TurnsPerSession field if non-nil, zero value otherwise.

### GetTurnsPerSessionOk

`func (o *AgentTelemetryMeasures) GetTurnsPerSessionOk() (*float32, bool)`

GetTurnsPerSessionOk returns a tuple with the TurnsPerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnsPerSession

`func (o *AgentTelemetryMeasures) SetTurnsPerSession(v float32)`

SetTurnsPerSession sets TurnsPerSession field to given value.

### HasTurnsPerSession

`func (o *AgentTelemetryMeasures) HasTurnsPerSession() bool`

HasTurnsPerSession returns a boolean if a field has been set.

### SetTurnsPerSessionNil

`func (o *AgentTelemetryMeasures) SetTurnsPerSessionNil(b bool)`

 SetTurnsPerSessionNil sets the value for TurnsPerSession to be an explicit nil

### UnsetTurnsPerSession
`func (o *AgentTelemetryMeasures) UnsetTurnsPerSession()`

UnsetTurnsPerSession ensures that no value is present for TurnsPerSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


