# RerankResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Zero-based index into the request&#39;s documents | 
**RelevanceScore** | **float32** | Relevance score, higher is more relevant | 

## Methods

### NewRerankResult

`func NewRerankResult(index int32, relevanceScore float32, ) *RerankResult`

NewRerankResult instantiates a new RerankResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankResultWithDefaults

`func NewRerankResultWithDefaults() *RerankResult`

NewRerankResultWithDefaults instantiates a new RerankResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RerankResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RerankResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RerankResult) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetRelevanceScore

`func (o *RerankResult) GetRelevanceScore() float32`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *RerankResult) GetRelevanceScoreOk() (*float32, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *RerankResult) SetRelevanceScore(v float32)`

SetRelevanceScore sets RelevanceScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


