# UsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**ByApiKey** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**ByEndpoint** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**ByModel** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**ByProvider** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**BySource** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**BySourceLabel** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**ByTool** | Pointer to [**[]UsageToolRow**](UsageToolRow.md) |  | [optional] [default to {}]
**ByUser** | [**[]UsageGroupRow**](UsageGroupRow.md) |  | 
**EndDate** | **string** |  | 
**ErrorsByStatusCode** | [**[]UsageErrorCodeRow**](UsageErrorCodeRow.md) |  | 
**Series** | [**[]UsageSeriesPoint**](UsageSeriesPoint.md) |  | 
**StartDate** | **string** |  | 
**Totals** | [**UsageTotals**](UsageTotals.md) |  | 

## Methods

### NewUsageSummary

`func NewUsageSummary(bucket string, byApiKey []UsageGroupRow, byEndpoint []UsageGroupRow, byModel []UsageGroupRow, byProvider []UsageGroupRow, bySource []UsageGroupRow, bySourceLabel []UsageGroupRow, byUser []UsageGroupRow, endDate string, errorsByStatusCode []UsageErrorCodeRow, series []UsageSeriesPoint, startDate string, totals UsageTotals, ) *UsageSummary`

NewUsageSummary instantiates a new UsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSummaryWithDefaults

`func NewUsageSummaryWithDefaults() *UsageSummary`

NewUsageSummaryWithDefaults instantiates a new UsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *UsageSummary) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UsageSummary) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UsageSummary) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetByApiKey

`func (o *UsageSummary) GetByApiKey() []UsageGroupRow`

GetByApiKey returns the ByApiKey field if non-nil, zero value otherwise.

### GetByApiKeyOk

`func (o *UsageSummary) GetByApiKeyOk() (*[]UsageGroupRow, bool)`

GetByApiKeyOk returns a tuple with the ByApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByApiKey

`func (o *UsageSummary) SetByApiKey(v []UsageGroupRow)`

SetByApiKey sets ByApiKey field to given value.


### GetByEndpoint

`func (o *UsageSummary) GetByEndpoint() []UsageGroupRow`

GetByEndpoint returns the ByEndpoint field if non-nil, zero value otherwise.

### GetByEndpointOk

`func (o *UsageSummary) GetByEndpointOk() (*[]UsageGroupRow, bool)`

GetByEndpointOk returns a tuple with the ByEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByEndpoint

`func (o *UsageSummary) SetByEndpoint(v []UsageGroupRow)`

SetByEndpoint sets ByEndpoint field to given value.


### GetByModel

`func (o *UsageSummary) GetByModel() []UsageGroupRow`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *UsageSummary) GetByModelOk() (*[]UsageGroupRow, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *UsageSummary) SetByModel(v []UsageGroupRow)`

SetByModel sets ByModel field to given value.


### GetByProvider

`func (o *UsageSummary) GetByProvider() []UsageGroupRow`

GetByProvider returns the ByProvider field if non-nil, zero value otherwise.

### GetByProviderOk

`func (o *UsageSummary) GetByProviderOk() (*[]UsageGroupRow, bool)`

GetByProviderOk returns a tuple with the ByProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProvider

`func (o *UsageSummary) SetByProvider(v []UsageGroupRow)`

SetByProvider sets ByProvider field to given value.


### GetBySource

`func (o *UsageSummary) GetBySource() []UsageGroupRow`

GetBySource returns the BySource field if non-nil, zero value otherwise.

### GetBySourceOk

`func (o *UsageSummary) GetBySourceOk() (*[]UsageGroupRow, bool)`

GetBySourceOk returns a tuple with the BySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySource

`func (o *UsageSummary) SetBySource(v []UsageGroupRow)`

SetBySource sets BySource field to given value.


### GetBySourceLabel

`func (o *UsageSummary) GetBySourceLabel() []UsageGroupRow`

GetBySourceLabel returns the BySourceLabel field if non-nil, zero value otherwise.

### GetBySourceLabelOk

`func (o *UsageSummary) GetBySourceLabelOk() (*[]UsageGroupRow, bool)`

GetBySourceLabelOk returns a tuple with the BySourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySourceLabel

`func (o *UsageSummary) SetBySourceLabel(v []UsageGroupRow)`

SetBySourceLabel sets BySourceLabel field to given value.


### GetByTool

`func (o *UsageSummary) GetByTool() []UsageToolRow`

GetByTool returns the ByTool field if non-nil, zero value otherwise.

### GetByToolOk

`func (o *UsageSummary) GetByToolOk() (*[]UsageToolRow, bool)`

GetByToolOk returns a tuple with the ByTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTool

`func (o *UsageSummary) SetByTool(v []UsageToolRow)`

SetByTool sets ByTool field to given value.

### HasByTool

`func (o *UsageSummary) HasByTool() bool`

HasByTool returns a boolean if a field has been set.

### GetByUser

`func (o *UsageSummary) GetByUser() []UsageGroupRow`

GetByUser returns the ByUser field if non-nil, zero value otherwise.

### GetByUserOk

`func (o *UsageSummary) GetByUserOk() (*[]UsageGroupRow, bool)`

GetByUserOk returns a tuple with the ByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUser

`func (o *UsageSummary) SetByUser(v []UsageGroupRow)`

SetByUser sets ByUser field to given value.


### GetEndDate

`func (o *UsageSummary) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageSummary) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageSummary) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetErrorsByStatusCode

`func (o *UsageSummary) GetErrorsByStatusCode() []UsageErrorCodeRow`

GetErrorsByStatusCode returns the ErrorsByStatusCode field if non-nil, zero value otherwise.

### GetErrorsByStatusCodeOk

`func (o *UsageSummary) GetErrorsByStatusCodeOk() (*[]UsageErrorCodeRow, bool)`

GetErrorsByStatusCodeOk returns a tuple with the ErrorsByStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsByStatusCode

`func (o *UsageSummary) SetErrorsByStatusCode(v []UsageErrorCodeRow)`

SetErrorsByStatusCode sets ErrorsByStatusCode field to given value.


### GetSeries

`func (o *UsageSummary) GetSeries() []UsageSeriesPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *UsageSummary) GetSeriesOk() (*[]UsageSeriesPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *UsageSummary) SetSeries(v []UsageSeriesPoint)`

SetSeries sets Series field to given value.


### GetStartDate

`func (o *UsageSummary) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageSummary) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageSummary) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetTotals

`func (o *UsageSummary) GetTotals() UsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *UsageSummary) GetTotalsOk() (*UsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *UsageSummary) SetTotals(v UsageTotals)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


