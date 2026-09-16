# AgentTelemetrySeriesPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTime** | Pointer to **float32** |  | [optional] [default to 0.0]
**ApiErrors** | Pointer to **int32** |  | [optional] [default to 0]
**BucketStart** | **string** |  | 
**Commits** | Pointer to **float32** |  | [optional] [default to 0.0]
**Cost** | Pointer to **float32** |  | [optional] [default to 0.0]
**LinesOfCode** | Pointer to **float32** |  | [optional] [default to 0.0]
**PullRequests** | Pointer to **float32** |  | [optional] [default to 0.0]
**ToolCalls** | Pointer to **int32** |  | [optional] [default to 0]
**Turns** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewAgentTelemetrySeriesPoint

`func NewAgentTelemetrySeriesPoint(bucketStart string, ) *AgentTelemetrySeriesPoint`

NewAgentTelemetrySeriesPoint instantiates a new AgentTelemetrySeriesPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetrySeriesPointWithDefaults

`func NewAgentTelemetrySeriesPointWithDefaults() *AgentTelemetrySeriesPoint`

NewAgentTelemetrySeriesPointWithDefaults instantiates a new AgentTelemetrySeriesPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTime

`func (o *AgentTelemetrySeriesPoint) GetActiveTime() float32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *AgentTelemetrySeriesPoint) GetActiveTimeOk() (*float32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *AgentTelemetrySeriesPoint) SetActiveTime(v float32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *AgentTelemetrySeriesPoint) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### GetApiErrors

`func (o *AgentTelemetrySeriesPoint) GetApiErrors() int32`

GetApiErrors returns the ApiErrors field if non-nil, zero value otherwise.

### GetApiErrorsOk

`func (o *AgentTelemetrySeriesPoint) GetApiErrorsOk() (*int32, bool)`

GetApiErrorsOk returns a tuple with the ApiErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiErrors

`func (o *AgentTelemetrySeriesPoint) SetApiErrors(v int32)`

SetApiErrors sets ApiErrors field to given value.

### HasApiErrors

`func (o *AgentTelemetrySeriesPoint) HasApiErrors() bool`

HasApiErrors returns a boolean if a field has been set.

### GetBucketStart

`func (o *AgentTelemetrySeriesPoint) GetBucketStart() string`

GetBucketStart returns the BucketStart field if non-nil, zero value otherwise.

### GetBucketStartOk

`func (o *AgentTelemetrySeriesPoint) GetBucketStartOk() (*string, bool)`

GetBucketStartOk returns a tuple with the BucketStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketStart

`func (o *AgentTelemetrySeriesPoint) SetBucketStart(v string)`

SetBucketStart sets BucketStart field to given value.


### GetCommits

`func (o *AgentTelemetrySeriesPoint) GetCommits() float32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *AgentTelemetrySeriesPoint) GetCommitsOk() (*float32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *AgentTelemetrySeriesPoint) SetCommits(v float32)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *AgentTelemetrySeriesPoint) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetCost

`func (o *AgentTelemetrySeriesPoint) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AgentTelemetrySeriesPoint) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AgentTelemetrySeriesPoint) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AgentTelemetrySeriesPoint) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetLinesOfCode

`func (o *AgentTelemetrySeriesPoint) GetLinesOfCode() float32`

GetLinesOfCode returns the LinesOfCode field if non-nil, zero value otherwise.

### GetLinesOfCodeOk

`func (o *AgentTelemetrySeriesPoint) GetLinesOfCodeOk() (*float32, bool)`

GetLinesOfCodeOk returns a tuple with the LinesOfCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesOfCode

`func (o *AgentTelemetrySeriesPoint) SetLinesOfCode(v float32)`

SetLinesOfCode sets LinesOfCode field to given value.

### HasLinesOfCode

`func (o *AgentTelemetrySeriesPoint) HasLinesOfCode() bool`

HasLinesOfCode returns a boolean if a field has been set.

### GetPullRequests

`func (o *AgentTelemetrySeriesPoint) GetPullRequests() float32`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *AgentTelemetrySeriesPoint) GetPullRequestsOk() (*float32, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *AgentTelemetrySeriesPoint) SetPullRequests(v float32)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *AgentTelemetrySeriesPoint) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetToolCalls

`func (o *AgentTelemetrySeriesPoint) GetToolCalls() int32`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AgentTelemetrySeriesPoint) GetToolCallsOk() (*int32, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AgentTelemetrySeriesPoint) SetToolCalls(v int32)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AgentTelemetrySeriesPoint) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetTurns

`func (o *AgentTelemetrySeriesPoint) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *AgentTelemetrySeriesPoint) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *AgentTelemetrySeriesPoint) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *AgentTelemetrySeriesPoint) HasTurns() bool`

HasTurns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


