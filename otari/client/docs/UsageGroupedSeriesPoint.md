# UsageGroupedSeriesPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketStart** | **string** |  | 
**Cost** | **float32** |  | 
**IsOther** | Pointer to **bool** |  | [optional] [default to false]
**Key** | **NullableString** |  | 
**Requests** | **int32** |  | 
**Tokens** | **int32** |  | 

## Methods

### NewUsageGroupedSeriesPoint

`func NewUsageGroupedSeriesPoint(bucketStart string, cost float32, key NullableString, requests int32, tokens int32, ) *UsageGroupedSeriesPoint`

NewUsageGroupedSeriesPoint instantiates a new UsageGroupedSeriesPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageGroupedSeriesPointWithDefaults

`func NewUsageGroupedSeriesPointWithDefaults() *UsageGroupedSeriesPoint`

NewUsageGroupedSeriesPointWithDefaults instantiates a new UsageGroupedSeriesPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketStart

`func (o *UsageGroupedSeriesPoint) GetBucketStart() string`

GetBucketStart returns the BucketStart field if non-nil, zero value otherwise.

### GetBucketStartOk

`func (o *UsageGroupedSeriesPoint) GetBucketStartOk() (*string, bool)`

GetBucketStartOk returns a tuple with the BucketStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketStart

`func (o *UsageGroupedSeriesPoint) SetBucketStart(v string)`

SetBucketStart sets BucketStart field to given value.


### GetCost

`func (o *UsageGroupedSeriesPoint) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageGroupedSeriesPoint) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageGroupedSeriesPoint) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetIsOther

`func (o *UsageGroupedSeriesPoint) GetIsOther() bool`

GetIsOther returns the IsOther field if non-nil, zero value otherwise.

### GetIsOtherOk

`func (o *UsageGroupedSeriesPoint) GetIsOtherOk() (*bool, bool)`

GetIsOtherOk returns a tuple with the IsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOther

`func (o *UsageGroupedSeriesPoint) SetIsOther(v bool)`

SetIsOther sets IsOther field to given value.

### HasIsOther

`func (o *UsageGroupedSeriesPoint) HasIsOther() bool`

HasIsOther returns a boolean if a field has been set.

### GetKey

`func (o *UsageGroupedSeriesPoint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UsageGroupedSeriesPoint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UsageGroupedSeriesPoint) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *UsageGroupedSeriesPoint) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UsageGroupedSeriesPoint) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetRequests

`func (o *UsageGroupedSeriesPoint) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageGroupedSeriesPoint) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageGroupedSeriesPoint) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetTokens

`func (o *UsageGroupedSeriesPoint) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageGroupedSeriesPoint) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageGroupedSeriesPoint) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


