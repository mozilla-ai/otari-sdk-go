# RRRerankResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Zero-based index into the original documents list | 
**RelevanceScore** | **float32** | Relevance score, higher is more relevant | 

## Methods

### NewRRRerankResult

`func NewRRRerankResult(index int32, relevanceScore float32, ) *RRRerankResult`

NewRRRerankResult instantiates a new RRRerankResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRRRerankResultWithDefaults

`func NewRRRerankResultWithDefaults() *RRRerankResult`

NewRRRerankResultWithDefaults instantiates a new RRRerankResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RRRerankResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RRRerankResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RRRerankResult) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetRelevanceScore

`func (o *RRRerankResult) GetRelevanceScore() float32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *RRRerankResult) GetRelevanceScoreOk() (*float32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *RRRerankResult) SetRelevanceScore(v float32)`

SetRelevanceScore sets RelevanceScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


