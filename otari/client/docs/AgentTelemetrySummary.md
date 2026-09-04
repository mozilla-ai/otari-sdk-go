# AgentTelemetrySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | [**AgentTelemetryBehavior**](AgentTelemetryBehavior.md) |  | 
**Bucket** | **string** |  | 
**EndDate** | **string** |  | 
**Measures** | [**AgentTelemetryMeasures**](AgentTelemetryMeasures.md) |  | 
**Outcomes** | [**AgentTelemetryOutcomes**](AgentTelemetryOutcomes.md) |  | 
**Series** | [**[]AgentTelemetrySeriesPoint**](AgentTelemetrySeriesPoint.md) |  | 
**StartDate** | **string** |  | 
**Usage** | [**AgentTelemetryUsage**](AgentTelemetryUsage.md) |  | 

## Methods

### NewAgentTelemetrySummary

`func NewAgentTelemetrySummary(behavior AgentTelemetryBehavior, bucket string, endDate string, measures AgentTelemetryMeasures, outcomes AgentTelemetryOutcomes, series []AgentTelemetrySeriesPoint, startDate string, usage AgentTelemetryUsage, ) *AgentTelemetrySummary`

NewAgentTelemetrySummary instantiates a new AgentTelemetrySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetrySummaryWithDefaults

`func NewAgentTelemetrySummaryWithDefaults() *AgentTelemetrySummary`

NewAgentTelemetrySummaryWithDefaults instantiates a new AgentTelemetrySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *AgentTelemetrySummary) GetBehavior() AgentTelemetryBehavior`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *AgentTelemetrySummary) GetBehaviorOk() (*AgentTelemetryBehavior, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *AgentTelemetrySummary) SetBehavior(v AgentTelemetryBehavior)`

SetBehavior sets Behavior field to given value.


### GetBucket

`func (o *AgentTelemetrySummary) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AgentTelemetrySummary) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AgentTelemetrySummary) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetEndDate

`func (o *AgentTelemetrySummary) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AgentTelemetrySummary) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AgentTelemetrySummary) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetMeasures

`func (o *AgentTelemetrySummary) GetMeasures() AgentTelemetryMeasures`

GetMeasures returns the Measures field if non-nil, zero value otherwise.

### GetMeasuresOk

`func (o *AgentTelemetrySummary) GetMeasuresOk() (*AgentTelemetryMeasures, bool)`

GetMeasuresOk returns a tuple with the Measures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasures

`func (o *AgentTelemetrySummary) SetMeasures(v AgentTelemetryMeasures)`

SetMeasures sets Measures field to given value.


### GetOutcomes

`func (o *AgentTelemetrySummary) GetOutcomes() AgentTelemetryOutcomes`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *AgentTelemetrySummary) GetOutcomesOk() (*AgentTelemetryOutcomes, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *AgentTelemetrySummary) SetOutcomes(v AgentTelemetryOutcomes)`

SetOutcomes sets Outcomes field to given value.


### GetSeries

`func (o *AgentTelemetrySummary) GetSeries() []AgentTelemetrySeriesPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AgentTelemetrySummary) GetSeriesOk() (*[]AgentTelemetrySeriesPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AgentTelemetrySummary) SetSeries(v []AgentTelemetrySeriesPoint)`

SetSeries sets Series field to given value.


### GetStartDate

`func (o *AgentTelemetrySummary) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AgentTelemetrySummary) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AgentTelemetrySummary) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetUsage

`func (o *AgentTelemetrySummary) GetUsage() AgentTelemetryUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AgentTelemetrySummary) GetUsageOk() (*AgentTelemetryUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AgentTelemetrySummary) SetUsage(v AgentTelemetryUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


