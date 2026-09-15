# PoolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | **int32** |  | 
**Warm** | **bool** |  | 

## Methods

### NewPoolStatus

`func NewPoolStatus(records int32, warm bool, ) *PoolStatus`

NewPoolStatus instantiates a new PoolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolStatusWithDefaults

`func NewPoolStatusWithDefaults() *PoolStatus`

NewPoolStatusWithDefaults instantiates a new PoolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *PoolStatus) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PoolStatus) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PoolStatus) SetRecords(v int32)`

SetRecords sets Records field to given value.


### GetWarm

`func (o *PoolStatus) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *PoolStatus) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *PoolStatus) SetWarm(v bool)`

SetWarm sets Warm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


