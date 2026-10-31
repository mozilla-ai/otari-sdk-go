# ProviderHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | **int32** |  | 
**History** | [**[]ProviderHistoryPoint**](ProviderHistoryPoint.md) | Every finished run in the last 24 hours, oldest first. | 
**Provider** | **string** |  | 
**Status** | **string** |  | 
**TopError** | **NullableString** | The most common failure, with the model id taken out. | 
**TopErrorCount** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewProviderHealthSummary

`func NewProviderHealthSummary(answered int32, history []ProviderHistoryPoint, provider string, status string, topError NullableString, topErrorCount int32, total int32, ) *ProviderHealthSummary`

NewProviderHealthSummary instantiates a new ProviderHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderHealthSummaryWithDefaults

`func NewProviderHealthSummaryWithDefaults() *ProviderHealthSummary`

NewProviderHealthSummaryWithDefaults instantiates a new ProviderHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *ProviderHealthSummary) GetAnswered() int32`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *ProviderHealthSummary) GetAnsweredOk() (*int32, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *ProviderHealthSummary) SetAnswered(v int32)`

SetAnswered sets Answered field to given value.


### GetHistory

`func (o *ProviderHealthSummary) GetHistory() []ProviderHistoryPoint`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *ProviderHealthSummary) GetHistoryOk() (*[]ProviderHistoryPoint, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *ProviderHealthSummary) SetHistory(v []ProviderHistoryPoint)`

SetHistory sets History field to given value.


### GetProvider

`func (o *ProviderHealthSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderHealthSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderHealthSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetStatus

`func (o *ProviderHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProviderHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProviderHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTopError

`func (o *ProviderHealthSummary) GetTopError() string`

GetTopError returns the TopError field if non-nil, zero value otherwise.

### GetTopErrorOk

`func (o *ProviderHealthSummary) GetTopErrorOk() (*string, bool)`

GetTopErrorOk returns a tuple with the TopError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopError

`func (o *ProviderHealthSummary) SetTopError(v string)`

SetTopError sets TopError field to given value.


### SetTopErrorNil

`func (o *ProviderHealthSummary) SetTopErrorNil(b bool)`

 SetTopErrorNil sets the value for TopError to be an explicit nil

### UnsetTopError
`func (o *ProviderHealthSummary) UnsetTopError()`

UnsetTopError ensures that no value is present for TopError, not even an explicit nil
### GetTopErrorCount

`func (o *ProviderHealthSummary) GetTopErrorCount() int32`

GetTopErrorCount returns the TopErrorCount field if non-nil, zero value otherwise.

### GetTopErrorCountOk

`func (o *ProviderHealthSummary) GetTopErrorCountOk() (*int32, bool)`

GetTopErrorCountOk returns a tuple with the TopErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopErrorCount

`func (o *ProviderHealthSummary) SetTopErrorCount(v int32)`

SetTopErrorCount sets TopErrorCount field to given value.


### GetTotal

`func (o *ProviderHealthSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderHealthSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderHealthSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


