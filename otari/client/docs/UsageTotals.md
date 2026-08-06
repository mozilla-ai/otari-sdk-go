# UsageTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgLatencyMs** | **NullableFloat32** |  | 
**BilledInputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**BilledOutputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CacheReadTokens** | **int32** |  | 
**CacheWrite1hTokens** | **int32** |  | 
**CacheWriteTokens** | **int32** |  | 
**CompletionTokens** | **int32** |  | 
**Cost** | **float32** |  | 
**ErrorCount** | **int32** |  | 
**PromptTokens** | **int32** |  | 
**RequestCount** | **int32** |  | 
**TotalTokens** | **int32** |  | 
**UnpricedRequests** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewUsageTotals

`func NewUsageTotals(avgLatencyMs NullableFloat32, cacheReadTokens int32, cacheWrite1hTokens int32, cacheWriteTokens int32, completionTokens int32, cost float32, errorCount int32, promptTokens int32, requestCount int32, totalTokens int32, ) *UsageTotals`

NewUsageTotals instantiates a new UsageTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTotalsWithDefaults

`func NewUsageTotalsWithDefaults() *UsageTotals`

NewUsageTotalsWithDefaults instantiates a new UsageTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgLatencyMs

`func (o *UsageTotals) GetAvgLatencyMs() float32`

GetAvgLatencyMs returns the AvgLatencyMs field if non-nil, zero value otherwise.

### GetAvgLatencyMsOk

`func (o *UsageTotals) GetAvgLatencyMsOk() (*float32, bool)`

GetAvgLatencyMsOk returns a tuple with the AvgLatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyMs

`func (o *UsageTotals) SetAvgLatencyMs(v float32)`

SetAvgLatencyMs sets AvgLatencyMs field to given value.


### SetAvgLatencyMsNil

`func (o *UsageTotals) SetAvgLatencyMsNil(b bool)`

 SetAvgLatencyMsNil sets the value for AvgLatencyMs to be an explicit nil

### UnsetAvgLatencyMs
`func (o *UsageTotals) UnsetAvgLatencyMs()`

UnsetAvgLatencyMs ensures that no value is present for AvgLatencyMs, not even an explicit nil
### GetBilledInputTokens

`func (o *UsageTotals) GetBilledInputTokens() int32`

GetBilledInputTokens returns the BilledInputTokens field if non-nil, zero value otherwise.

### GetBilledInputTokensOk

`func (o *UsageTotals) GetBilledInputTokensOk() (*int32, bool)`

GetBilledInputTokensOk returns a tuple with the BilledInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledInputTokens

`func (o *UsageTotals) SetBilledInputTokens(v int32)`

SetBilledInputTokens sets BilledInputTokens field to given value.

### HasBilledInputTokens

`func (o *UsageTotals) HasBilledInputTokens() bool`

HasBilledInputTokens returns a boolean if a field has been set.

### GetBilledOutputTokens

`func (o *UsageTotals) GetBilledOutputTokens() int32`

GetBilledOutputTokens returns the BilledOutputTokens field if non-nil, zero value otherwise.

### GetBilledOutputTokensOk

`func (o *UsageTotals) GetBilledOutputTokensOk() (*int32, bool)`

GetBilledOutputTokensOk returns a tuple with the BilledOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledOutputTokens

`func (o *UsageTotals) SetBilledOutputTokens(v int32)`

SetBilledOutputTokens sets BilledOutputTokens field to given value.

### HasBilledOutputTokens

`func (o *UsageTotals) HasBilledOutputTokens() bool`

HasBilledOutputTokens returns a boolean if a field has been set.

### GetCacheReadTokens

`func (o *UsageTotals) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *UsageTotals) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *UsageTotals) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### GetCacheWrite1hTokens

`func (o *UsageTotals) GetCacheWrite1hTokens() int32`

GetCacheWrite1hTokens returns the CacheWrite1hTokens field if non-nil, zero value otherwise.

### GetCacheWrite1hTokensOk

`func (o *UsageTotals) GetCacheWrite1hTokensOk() (*int32, bool)`

GetCacheWrite1hTokensOk returns a tuple with the CacheWrite1hTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hTokens

`func (o *UsageTotals) SetCacheWrite1hTokens(v int32)`

SetCacheWrite1hTokens sets CacheWrite1hTokens field to given value.


### GetCacheWriteTokens

`func (o *UsageTotals) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *UsageTotals) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *UsageTotals) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.


### GetCompletionTokens

`func (o *UsageTotals) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageTotals) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageTotals) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetCost

`func (o *UsageTotals) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageTotals) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageTotals) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetErrorCount

`func (o *UsageTotals) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *UsageTotals) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *UsageTotals) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetPromptTokens

`func (o *UsageTotals) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageTotals) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageTotals) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetRequestCount

`func (o *UsageTotals) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *UsageTotals) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *UsageTotals) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetTotalTokens

`func (o *UsageTotals) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageTotals) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageTotals) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetUnpricedRequests

`func (o *UsageTotals) GetUnpricedRequests() int32`

GetUnpricedRequests returns the UnpricedRequests field if non-nil, zero value otherwise.

### GetUnpricedRequestsOk

`func (o *UsageTotals) GetUnpricedRequestsOk() (*int32, bool)`

GetUnpricedRequestsOk returns a tuple with the UnpricedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpricedRequests

`func (o *UsageTotals) SetUnpricedRequests(v int32)`

SetUnpricedRequests sets UnpricedRequests field to given value.

### HasUnpricedRequests

`func (o *UsageTotals) HasUnpricedRequests() bool`

HasUnpricedRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


