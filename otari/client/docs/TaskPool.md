# TaskPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | **int32** |  | 
**TaskId** | **string** |  | 
**Warm** | **bool** |  | 

## Methods

### NewTaskPool

`func NewTaskPool(records int32, taskId string, warm bool, ) *TaskPool`

NewTaskPool instantiates a new TaskPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPoolWithDefaults

`func NewTaskPoolWithDefaults() *TaskPool`

NewTaskPoolWithDefaults instantiates a new TaskPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *TaskPool) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *TaskPool) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *TaskPool) SetRecords(v int32)`

SetRecords sets Records field to given value.


### GetTaskId

`func (o *TaskPool) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskPool) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskPool) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetWarm

`func (o *TaskPool) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *TaskPool) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *TaskPool) SetWarm(v bool)`

SetWarm sets Warm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


