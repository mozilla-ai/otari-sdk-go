# RankResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]RecordedPool**](RecordedPool.md) |  | 
**Recorded** | **int32** |  | 
**SeedCount** | **int32** |  | 

## Methods

### NewRankResponse

`func NewRankResponse(pools []RecordedPool, recorded int32, seedCount int32, ) *RankResponse`

NewRankResponse instantiates a new RankResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankResponseWithDefaults

`func NewRankResponseWithDefaults() *RankResponse`

NewRankResponseWithDefaults instantiates a new RankResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *RankResponse) GetPools() []RecordedPool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *RankResponse) GetPoolsOk() (*[]RecordedPool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *RankResponse) SetPools(v []RecordedPool)`

SetPools sets Pools field to given value.


### GetRecorded

`func (o *RankResponse) GetRecorded() int32`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *RankResponse) GetRecordedOk() (*int32, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *RankResponse) SetRecorded(v int32)`

SetRecorded sets Recorded field to given value.


### GetSeedCount

`func (o *RankResponse) GetSeedCount() int32`

GetSeedCount returns the SeedCount field if non-nil, zero value otherwise.

### GetSeedCountOk

`func (o *RankResponse) GetSeedCountOk() (*int32, bool)`

GetSeedCountOk returns a tuple with the SeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedCount

`func (o *RankResponse) SetSeedCount(v int32)`

SetSeedCount sets SeedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


