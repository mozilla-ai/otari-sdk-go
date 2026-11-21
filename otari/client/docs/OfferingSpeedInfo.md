# OfferingSpeedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutputTokensPerSecondP50** | **NullableFloat32** | Median output tokens per second over the whole request, time to first token included, as OpenRouter measures it. Null on fewer than 20 measured requests. | 
**OutputTokensPerSecondP95** | **NullableFloat32** | The 5th percentile of output tokens per second: the slow tail. | 
**ThroughputSamples** | **int32** | Requests the throughput is read from: those with 20 or more output tokens. | 
**TtftMsP50** | **NullableInt32** | Median milliseconds to the first token. Null on fewer than 20 streamed requests. | 
**TtftMsP95** | **NullableInt32** | The 95th percentile of milliseconds to the first token, rounded: the slow tail. | 
**TtftSamples** | **int32** | Streamed requests the time to first token is read from. | 

## Methods

### NewOfferingSpeedInfo

`func NewOfferingSpeedInfo(outputTokensPerSecondP50 NullableFloat32, outputTokensPerSecondP95 NullableFloat32, throughputSamples int32, ttftMsP50 NullableInt32, ttftMsP95 NullableInt32, ttftSamples int32, ) *OfferingSpeedInfo`

NewOfferingSpeedInfo instantiates a new OfferingSpeedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferingSpeedInfoWithDefaults

`func NewOfferingSpeedInfoWithDefaults() *OfferingSpeedInfo`

NewOfferingSpeedInfoWithDefaults instantiates a new OfferingSpeedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputTokensPerSecondP50

`func (o *OfferingSpeedInfo) GetOutputTokensPerSecondP50() float32`

GetOutputTokensPerSecondP50 returns the OutputTokensPerSecondP50 field if non-nil, zero value otherwise.

### GetOutputTokensPerSecondP50Ok

`func (o *OfferingSpeedInfo) GetOutputTokensPerSecondP50Ok() (*float32, bool)`

GetOutputTokensPerSecondP50Ok returns a tuple with the OutputTokensPerSecondP50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokensPerSecondP50

`func (o *OfferingSpeedInfo) SetOutputTokensPerSecondP50(v float32)`

SetOutputTokensPerSecondP50 sets OutputTokensPerSecondP50 field to given value.


### SetOutputTokensPerSecondP50Nil

`func (o *OfferingSpeedInfo) SetOutputTokensPerSecondP50Nil(b bool)`

 SetOutputTokensPerSecondP50Nil sets the value for OutputTokensPerSecondP50 to be an explicit nil

### UnsetOutputTokensPerSecondP50
`func (o *OfferingSpeedInfo) UnsetOutputTokensPerSecondP50()`

UnsetOutputTokensPerSecondP50 ensures that no value is present for OutputTokensPerSecondP50, not even an explicit nil
### GetOutputTokensPerSecondP95

`func (o *OfferingSpeedInfo) GetOutputTokensPerSecondP95() float32`

GetOutputTokensPerSecondP95 returns the OutputTokensPerSecondP95 field if non-nil, zero value otherwise.

### GetOutputTokensPerSecondP95Ok

`func (o *OfferingSpeedInfo) GetOutputTokensPerSecondP95Ok() (*float32, bool)`

GetOutputTokensPerSecondP95Ok returns a tuple with the OutputTokensPerSecondP95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokensPerSecondP95

`func (o *OfferingSpeedInfo) SetOutputTokensPerSecondP95(v float32)`

SetOutputTokensPerSecondP95 sets OutputTokensPerSecondP95 field to given value.


### SetOutputTokensPerSecondP95Nil

`func (o *OfferingSpeedInfo) SetOutputTokensPerSecondP95Nil(b bool)`

 SetOutputTokensPerSecondP95Nil sets the value for OutputTokensPerSecondP95 to be an explicit nil

### UnsetOutputTokensPerSecondP95
`func (o *OfferingSpeedInfo) UnsetOutputTokensPerSecondP95()`

UnsetOutputTokensPerSecondP95 ensures that no value is present for OutputTokensPerSecondP95, not even an explicit nil
### GetThroughputSamples

`func (o *OfferingSpeedInfo) GetThroughputSamples() int32`

GetThroughputSamples returns the ThroughputSamples field if non-nil, zero value otherwise.

### GetThroughputSamplesOk

`func (o *OfferingSpeedInfo) GetThroughputSamplesOk() (*int32, bool)`

GetThroughputSamplesOk returns a tuple with the ThroughputSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputSamples

`func (o *OfferingSpeedInfo) SetThroughputSamples(v int32)`

SetThroughputSamples sets ThroughputSamples field to given value.


### GetTtftMsP50

`func (o *OfferingSpeedInfo) GetTtftMsP50() int32`

GetTtftMsP50 returns the TtftMsP50 field if non-nil, zero value otherwise.

### GetTtftMsP50Ok

`func (o *OfferingSpeedInfo) GetTtftMsP50Ok() (*int32, bool)`

GetTtftMsP50Ok returns a tuple with the TtftMsP50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtftMsP50

`func (o *OfferingSpeedInfo) SetTtftMsP50(v int32)`

SetTtftMsP50 sets TtftMsP50 field to given value.


### SetTtftMsP50Nil

`func (o *OfferingSpeedInfo) SetTtftMsP50Nil(b bool)`

 SetTtftMsP50Nil sets the value for TtftMsP50 to be an explicit nil

### UnsetTtftMsP50
`func (o *OfferingSpeedInfo) UnsetTtftMsP50()`

UnsetTtftMsP50 ensures that no value is present for TtftMsP50, not even an explicit nil
### GetTtftMsP95

`func (o *OfferingSpeedInfo) GetTtftMsP95() int32`

GetTtftMsP95 returns the TtftMsP95 field if non-nil, zero value otherwise.

### GetTtftMsP95Ok

`func (o *OfferingSpeedInfo) GetTtftMsP95Ok() (*int32, bool)`

GetTtftMsP95Ok returns a tuple with the TtftMsP95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtftMsP95

`func (o *OfferingSpeedInfo) SetTtftMsP95(v int32)`

SetTtftMsP95 sets TtftMsP95 field to given value.


### SetTtftMsP95Nil

`func (o *OfferingSpeedInfo) SetTtftMsP95Nil(b bool)`

 SetTtftMsP95Nil sets the value for TtftMsP95 to be an explicit nil

### UnsetTtftMsP95
`func (o *OfferingSpeedInfo) UnsetTtftMsP95()`

UnsetTtftMsP95 ensures that no value is present for TtftMsP95, not even an explicit nil
### GetTtftSamples

`func (o *OfferingSpeedInfo) GetTtftSamples() int32`

GetTtftSamples returns the TtftSamples field if non-nil, zero value otherwise.

### GetTtftSamplesOk

`func (o *OfferingSpeedInfo) GetTtftSamplesOk() (*int32, bool)`

GetTtftSamplesOk returns a tuple with the TtftSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtftSamples

`func (o *OfferingSpeedInfo) SetTtftSamples(v int32)`

SetTtftSamples sets TtftSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


