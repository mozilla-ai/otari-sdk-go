# RoutingPlanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Declined** | Pointer to **NullableString** | The router&#39;s own reason for not choosing, when it was asked and declined. | [optional] 
**Dropped** | Pointer to [**[]DroppedCandidateRecord**](DroppedCandidateRecord.md) | The candidates that never became attempts, each with the reason it was left out. | [optional] 

## Methods

### NewRoutingPlanRecord

`func NewRoutingPlanRecord() *RoutingPlanRecord`

NewRoutingPlanRecord instantiates a new RoutingPlanRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPlanRecordWithDefaults

`func NewRoutingPlanRecordWithDefaults() *RoutingPlanRecord`

NewRoutingPlanRecordWithDefaults instantiates a new RoutingPlanRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclined

`func (o *RoutingPlanRecord) GetDeclined() string`

GetDeclined returns the Declined field if non-nil, zero value otherwise.

### GetDeclinedOk

`func (o *RoutingPlanRecord) GetDeclinedOk() (*string, bool)`

GetDeclinedOk returns a tuple with the Declined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclined

`func (o *RoutingPlanRecord) SetDeclined(v string)`

SetDeclined sets Declined field to given value.

### HasDeclined

`func (o *RoutingPlanRecord) HasDeclined() bool`

HasDeclined returns a boolean if a field has been set.

### SetDeclinedNil

`func (o *RoutingPlanRecord) SetDeclinedNil(b bool)`

 SetDeclinedNil sets the value for Declined to be an explicit nil

### UnsetDeclined
`func (o *RoutingPlanRecord) UnsetDeclined()`

UnsetDeclined ensures that no value is present for Declined, not even an explicit nil
### GetDropped

`func (o *RoutingPlanRecord) GetDropped() []DroppedCandidateRecord`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *RoutingPlanRecord) GetDroppedOk() (*[]DroppedCandidateRecord, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *RoutingPlanRecord) SetDropped(v []DroppedCandidateRecord)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *RoutingPlanRecord) HasDropped() bool`

HasDropped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


