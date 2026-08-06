# UsageGroupedSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**EndDate** | **string** |  | 
**GroupBy** | **string** |  | 
**Groups** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**Points** | [**[]UsageGroupedSeriesPoint**](UsageGroupedSeriesPoint.md) |  | 
**StartDate** | **string** |  | 

## Methods

### NewUsageGroupedSeries

`func NewUsageGroupedSeries(bucket string, endDate string, groupBy string, groups []UsageGroupRow, points []UsageGroupedSeriesPoint, startDate string, ) *UsageGroupedSeries`

NewUsageGroupedSeries instantiates a new UsageGroupedSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageGroupedSeriesWithDefaults

`func NewUsageGroupedSeriesWithDefaults() *UsageGroupedSeries`

NewUsageGroupedSeriesWithDefaults instantiates a new UsageGroupedSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *UsageGroupedSeries) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UsageGroupedSeries) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UsageGroupedSeries) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetEndDate

`func (o *UsageGroupedSeries) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageGroupedSeries) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageGroupedSeries) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetGroupBy

`func (o *UsageGroupedSeries) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *UsageGroupedSeries) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *UsageGroupedSeries) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.


### GetGroups

`func (o *UsageGroupedSeries) GetGroups() []UsageGroupRow`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UsageGroupedSeries) GetGroupsOk() (*[]UsageGroupRow, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UsageGroupedSeries) SetGroups(v []UsageGroupRow)`

SetGroups sets Groups field to given value.


### GetPoints

`func (o *UsageGroupedSeries) GetPoints() []UsageGroupedSeriesPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *UsageGroupedSeries) GetPointsOk() (*[]UsageGroupedSeriesPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *UsageGroupedSeries) SetPoints(v []UsageGroupedSeriesPoint)`

SetPoints sets Points field to given value.


### GetStartDate

`func (o *UsageGroupedSeries) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageGroupedSeries) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageGroupedSeries) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


