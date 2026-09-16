# AcceptedSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | **time.Time** |  | 
**AcceptedBy** | **string** | &#x60;operator&#x60; for a dashboard confirm, &#x60;schedule&#x60; for the auto policy. | 
**Id** | **string** |  | 
**ModelCount** | **int32** |  | 

## Methods

### NewAcceptedSnapshotResponse

`func NewAcceptedSnapshotResponse(acceptedAt time.Time, acceptedBy string, id string, modelCount int32, ) *AcceptedSnapshotResponse`

NewAcceptedSnapshotResponse instantiates a new AcceptedSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedSnapshotResponseWithDefaults

`func NewAcceptedSnapshotResponseWithDefaults() *AcceptedSnapshotResponse`

NewAcceptedSnapshotResponseWithDefaults instantiates a new AcceptedSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *AcceptedSnapshotResponse) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *AcceptedSnapshotResponse) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *AcceptedSnapshotResponse) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.


### GetAcceptedBy

`func (o *AcceptedSnapshotResponse) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *AcceptedSnapshotResponse) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *AcceptedSnapshotResponse) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.


### GetId

`func (o *AcceptedSnapshotResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcceptedSnapshotResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcceptedSnapshotResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModelCount

`func (o *AcceptedSnapshotResponse) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *AcceptedSnapshotResponse) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *AcceptedSnapshotResponse) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


