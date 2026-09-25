# AllocationHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CappedCount** | **int32** | Rows with a finite cap, which are the ones that can be judged. | 
**NearCount** | **int32** | Rows at 80% of their allowance or more, but not past it. | 
**OverCount** | **int32** | Rows at or past their allowance. | 
**TotalCount** | **int32** | Rows of any kind, so a caller can tell &#39;none configured&#39; from &#39;none caps spend&#39;. | 
**Worst** | [**NullableWorstAllocationResponse**](WorstAllocationResponse.md) |  | 

## Methods

### NewAllocationHealthResponse

`func NewAllocationHealthResponse(cappedCount int32, nearCount int32, overCount int32, totalCount int32, worst NullableWorstAllocationResponse, ) *AllocationHealthResponse`

NewAllocationHealthResponse instantiates a new AllocationHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationHealthResponseWithDefaults

`func NewAllocationHealthResponseWithDefaults() *AllocationHealthResponse`

NewAllocationHealthResponseWithDefaults instantiates a new AllocationHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCappedCount

`func (o *AllocationHealthResponse) GetCappedCount() int32`

GetCappedCount returns the CappedCount field if non-nil, zero value otherwise.

### GetCappedCountOk

`func (o *AllocationHealthResponse) GetCappedCountOk() (*int32, bool)`

GetCappedCountOk returns a tuple with the CappedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCappedCount

`func (o *AllocationHealthResponse) SetCappedCount(v int32)`

SetCappedCount sets CappedCount field to given value.


### GetNearCount

`func (o *AllocationHealthResponse) GetNearCount() int32`

GetNearCount returns the NearCount field if non-nil, zero value otherwise.

### GetNearCountOk

`func (o *AllocationHealthResponse) GetNearCountOk() (*int32, bool)`

GetNearCountOk returns a tuple with the NearCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearCount

`func (o *AllocationHealthResponse) SetNearCount(v int32)`

SetNearCount sets NearCount field to given value.


### GetOverCount

`func (o *AllocationHealthResponse) GetOverCount() int32`

GetOverCount returns the OverCount field if non-nil, zero value otherwise.

### GetOverCountOk

`func (o *AllocationHealthResponse) GetOverCountOk() (*int32, bool)`

GetOverCountOk returns a tuple with the OverCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCount

`func (o *AllocationHealthResponse) SetOverCount(v int32)`

SetOverCount sets OverCount field to given value.


### GetTotalCount

`func (o *AllocationHealthResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AllocationHealthResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AllocationHealthResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetWorst

`func (o *AllocationHealthResponse) GetWorst() WorstAllocationResponse`

GetWorst returns the Worst field if non-nil, zero value otherwise.

### GetWorstOk

`func (o *AllocationHealthResponse) GetWorstOk() (*WorstAllocationResponse, bool)`

GetWorstOk returns a tuple with the Worst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorst

`func (o *AllocationHealthResponse) SetWorst(v WorstAllocationResponse)`

SetWorst sets Worst field to given value.


### SetWorstNil

`func (o *AllocationHealthResponse) SetWorstNil(b bool)`

 SetWorstNil sets the value for Worst to be an explicit nil

### UnsetWorst
`func (o *AllocationHealthResponse) UnsetWorst()`

UnsetWorst ensures that no value is present for Worst, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


