# RecordedPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | **int32** |  | 
**TaskId** | **NullableString** |  | 
**Warm** | **bool** |  | 

## Methods

### NewRecordedPool

`func NewRecordedPool(records int32, taskId NullableString, warm bool, ) *RecordedPool`

NewRecordedPool instantiates a new RecordedPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordedPoolWithDefaults

`func NewRecordedPoolWithDefaults() *RecordedPool`

NewRecordedPoolWithDefaults instantiates a new RecordedPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *RecordedPool) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecordedPool) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecordedPool) SetRecords(v int32)`

SetRecords sets Records field to given value.


### GetTaskId

`func (o *RecordedPool) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *RecordedPool) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *RecordedPool) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### SetTaskIdNil

`func (o *RecordedPool) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *RecordedPool) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetWarm

`func (o *RecordedPool) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *RecordedPool) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *RecordedPool) SetWarm(v bool)`

SetWarm sets Warm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


