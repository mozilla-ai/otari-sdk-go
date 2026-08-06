# UsageSeriesPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketStart** | **string** |  | 
**CacheReadTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CacheWriteTokens** | Pointer to **int32** |  | [optional] [default to 0]
**Cost** | **float32** |  | 
**Errors** | Pointer to **int32** |  | [optional] [default to 0]
**InputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**OutputTokens** | Pointer to **int32** |  | [optional] [default to 0]
**Requests** | **int32** |  | 
**Tokens** | **int32** |  | 

## Methods

### NewUsageSeriesPoint

`func NewUsageSeriesPoint(bucketStart string, cost float32, requests int32, tokens int32, ) *UsageSeriesPoint`

NewUsageSeriesPoint instantiates a new UsageSeriesPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSeriesPointWithDefaults

`func NewUsageSeriesPointWithDefaults() *UsageSeriesPoint`

NewUsageSeriesPointWithDefaults instantiates a new UsageSeriesPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketStart

`func (o *UsageSeriesPoint) GetBucketStart() string`

GetBucketStart returns the BucketStart field if non-nil, zero value otherwise.

### GetBucketStartOk

`func (o *UsageSeriesPoint) GetBucketStartOk() (*string, bool)`

GetBucketStartOk returns a tuple with the BucketStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketStart

`func (o *UsageSeriesPoint) SetBucketStart(v string)`

SetBucketStart sets BucketStart field to given value.


### GetCacheReadTokens

`func (o *UsageSeriesPoint) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *UsageSeriesPoint) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *UsageSeriesPoint) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.

### HasCacheReadTokens

`func (o *UsageSeriesPoint) HasCacheReadTokens() bool`

HasCacheReadTokens returns a boolean if a field has been set.

### GetCacheWriteTokens

`func (o *UsageSeriesPoint) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *UsageSeriesPoint) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *UsageSeriesPoint) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *UsageSeriesPoint) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### GetCost

`func (o *UsageSeriesPoint) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageSeriesPoint) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageSeriesPoint) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetErrors

`func (o *UsageSeriesPoint) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsageSeriesPoint) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsageSeriesPoint) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UsageSeriesPoint) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInputTokens

`func (o *UsageSeriesPoint) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *UsageSeriesPoint) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *UsageSeriesPoint) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *UsageSeriesPoint) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *UsageSeriesPoint) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *UsageSeriesPoint) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *UsageSeriesPoint) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *UsageSeriesPoint) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetRequests

`func (o *UsageSeriesPoint) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageSeriesPoint) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageSeriesPoint) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetTokens

`func (o *UsageSeriesPoint) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageSeriesPoint) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageSeriesPoint) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


