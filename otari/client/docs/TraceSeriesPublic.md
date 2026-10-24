# TraceSeriesPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Points** | [**[]TraceBucketPublic**](TraceBucketPublic.md) |  | 

## Methods

### NewTraceSeriesPublic

`func NewTraceSeriesPublic(bucket string, points []TraceBucketPublic, ) *TraceSeriesPublic`

NewTraceSeriesPublic instantiates a new TraceSeriesPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSeriesPublicWithDefaults

`func NewTraceSeriesPublicWithDefaults() *TraceSeriesPublic`

NewTraceSeriesPublicWithDefaults instantiates a new TraceSeriesPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *TraceSeriesPublic) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *TraceSeriesPublic) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *TraceSeriesPublic) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPoints

`func (o *TraceSeriesPublic) GetPoints() []TraceBucketPublic`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *TraceSeriesPublic) GetPointsOk() (*[]TraceBucketPublic, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *TraceSeriesPublic) SetPoints(v []TraceBucketPublic)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


