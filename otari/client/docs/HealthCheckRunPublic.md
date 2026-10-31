# HealthCheckRunPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | **int32** |  | 
**FinishedAt** | **NullableTime** |  | 
**Id** | **string** |  | 
**IncludedReasoning** | **bool** |  | 
**Passed** | **int32** |  | 
**StartedAt** | **time.Time** |  | 
**Total** | **int32** |  | 

## Methods

### NewHealthCheckRunPublic

`func NewHealthCheckRunPublic(failed int32, finishedAt NullableTime, id string, includedReasoning bool, passed int32, startedAt time.Time, total int32, ) *HealthCheckRunPublic`

NewHealthCheckRunPublic instantiates a new HealthCheckRunPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckRunPublicWithDefaults

`func NewHealthCheckRunPublicWithDefaults() *HealthCheckRunPublic`

NewHealthCheckRunPublicWithDefaults instantiates a new HealthCheckRunPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *HealthCheckRunPublic) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *HealthCheckRunPublic) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *HealthCheckRunPublic) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetFinishedAt

`func (o *HealthCheckRunPublic) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *HealthCheckRunPublic) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *HealthCheckRunPublic) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *HealthCheckRunPublic) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *HealthCheckRunPublic) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetId

`func (o *HealthCheckRunPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheckRunPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheckRunPublic) SetId(v string)`

SetId sets Id field to given value.


### GetIncludedReasoning

`func (o *HealthCheckRunPublic) GetIncludedReasoning() bool`

GetIncludedReasoning returns the IncludedReasoning field if non-nil, zero value otherwise.

### GetIncludedReasoningOk

`func (o *HealthCheckRunPublic) GetIncludedReasoningOk() (*bool, bool)`

GetIncludedReasoningOk returns a tuple with the IncludedReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedReasoning

`func (o *HealthCheckRunPublic) SetIncludedReasoning(v bool)`

SetIncludedReasoning sets IncludedReasoning field to given value.


### GetPassed

`func (o *HealthCheckRunPublic) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *HealthCheckRunPublic) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *HealthCheckRunPublic) SetPassed(v int32)`

SetPassed sets Passed field to given value.


### GetStartedAt

`func (o *HealthCheckRunPublic) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *HealthCheckRunPublic) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *HealthCheckRunPublic) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTotal

`func (o *HealthCheckRunPublic) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HealthCheckRunPublic) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HealthCheckRunPublic) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


