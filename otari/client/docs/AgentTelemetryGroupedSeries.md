# AgentTelemetryGroupedSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**EndDate** | **string** |  | 
**GroupBy** | **string** |  | 
**Groups** | [**[]AgentTelemetryGroupRow**](AgentTelemetryGroupRow.md) |  | 
**Points** | [**[]AgentTelemetryGroupedSeriesPoint**](AgentTelemetryGroupedSeriesPoint.md) |  | 
**StartDate** | **string** |  | 

## Methods

### NewAgentTelemetryGroupedSeries

`func NewAgentTelemetryGroupedSeries(bucket string, endDate string, groupBy string, groups []AgentTelemetryGroupRow, points []AgentTelemetryGroupedSeriesPoint, startDate string, ) *AgentTelemetryGroupedSeries`

NewAgentTelemetryGroupedSeries instantiates a new AgentTelemetryGroupedSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryGroupedSeriesWithDefaults

`func NewAgentTelemetryGroupedSeriesWithDefaults() *AgentTelemetryGroupedSeries`

NewAgentTelemetryGroupedSeriesWithDefaults instantiates a new AgentTelemetryGroupedSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *AgentTelemetryGroupedSeries) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AgentTelemetryGroupedSeries) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AgentTelemetryGroupedSeries) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetEndDate

`func (o *AgentTelemetryGroupedSeries) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AgentTelemetryGroupedSeries) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AgentTelemetryGroupedSeries) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetGroupBy

`func (o *AgentTelemetryGroupedSeries) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *AgentTelemetryGroupedSeries) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *AgentTelemetryGroupedSeries) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.


### GetGroups

`func (o *AgentTelemetryGroupedSeries) GetGroups() []AgentTelemetryGroupRow`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AgentTelemetryGroupedSeries) GetGroupsOk() (*[]AgentTelemetryGroupRow, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AgentTelemetryGroupedSeries) SetGroups(v []AgentTelemetryGroupRow)`

SetGroups sets Groups field to given value.


### GetPoints

`func (o *AgentTelemetryGroupedSeries) GetPoints() []AgentTelemetryGroupedSeriesPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *AgentTelemetryGroupedSeries) GetPointsOk() (*[]AgentTelemetryGroupedSeriesPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *AgentTelemetryGroupedSeries) SetPoints(v []AgentTelemetryGroupedSeriesPoint)`

SetPoints sets Points field to given value.


### GetStartDate

`func (o *AgentTelemetryGroupedSeries) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AgentTelemetryGroupedSeries) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AgentTelemetryGroupedSeries) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


