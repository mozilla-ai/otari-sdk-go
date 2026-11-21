# TraceBucketPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | Bucket start, ISO 8601 UTC. | 
**Failed** | **int32** | Sessions active in the bucket that hold an unrecovered error other than a tool&#39;s. A tool&#39;s error, an MCP tool&#39;s included, goes back to the model, so it alone does not fail a session. | 
**Succeeded** | **int32** | Sessions active in the bucket that did not fail. | 

## Methods

### NewTraceBucketPublic

`func NewTraceBucketPublic(bucket string, failed int32, succeeded int32, ) *TraceBucketPublic`

NewTraceBucketPublic instantiates a new TraceBucketPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceBucketPublicWithDefaults

`func NewTraceBucketPublicWithDefaults() *TraceBucketPublic`

NewTraceBucketPublicWithDefaults instantiates a new TraceBucketPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *TraceBucketPublic) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *TraceBucketPublic) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *TraceBucketPublic) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetFailed

`func (o *TraceBucketPublic) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TraceBucketPublic) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TraceBucketPublic) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetSucceeded

`func (o *TraceBucketPublic) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *TraceBucketPublic) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *TraceBucketPublic) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


