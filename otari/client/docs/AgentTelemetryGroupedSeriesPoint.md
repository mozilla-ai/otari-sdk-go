# AgentTelemetryGroupedSeriesPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketStart** | **string** |  | 
**IsOther** | Pointer to **bool** |  | [optional] [default to false]
**Key** | **NullableString** |  | 
**Rows** | **int32** |  | 

## Methods

### NewAgentTelemetryGroupedSeriesPoint

`func NewAgentTelemetryGroupedSeriesPoint(bucketStart string, key NullableString, rows int32, ) *AgentTelemetryGroupedSeriesPoint`

NewAgentTelemetryGroupedSeriesPoint instantiates a new AgentTelemetryGroupedSeriesPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryGroupedSeriesPointWithDefaults

`func NewAgentTelemetryGroupedSeriesPointWithDefaults() *AgentTelemetryGroupedSeriesPoint`

NewAgentTelemetryGroupedSeriesPointWithDefaults instantiates a new AgentTelemetryGroupedSeriesPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketStart

`func (o *AgentTelemetryGroupedSeriesPoint) GetBucketStart() string`

GetBucketStart returns the BucketStart field if non-nil, zero value otherwise.

### GetBucketStartOk

`func (o *AgentTelemetryGroupedSeriesPoint) GetBucketStartOk() (*string, bool)`

GetBucketStartOk returns a tuple with the BucketStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketStart

`func (o *AgentTelemetryGroupedSeriesPoint) SetBucketStart(v string)`

SetBucketStart sets BucketStart field to given value.


### GetIsOther

`func (o *AgentTelemetryGroupedSeriesPoint) GetIsOther() bool`

GetIsOther returns the IsOther field if non-nil, zero value otherwise.

### GetIsOtherOk

`func (o *AgentTelemetryGroupedSeriesPoint) GetIsOtherOk() (*bool, bool)`

GetIsOtherOk returns a tuple with the IsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOther

`func (o *AgentTelemetryGroupedSeriesPoint) SetIsOther(v bool)`

SetIsOther sets IsOther field to given value.

### HasIsOther

`func (o *AgentTelemetryGroupedSeriesPoint) HasIsOther() bool`

HasIsOther returns a boolean if a field has been set.

### GetKey

`func (o *AgentTelemetryGroupedSeriesPoint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentTelemetryGroupedSeriesPoint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentTelemetryGroupedSeriesPoint) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *AgentTelemetryGroupedSeriesPoint) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AgentTelemetryGroupedSeriesPoint) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetRows

`func (o *AgentTelemetryGroupedSeriesPoint) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AgentTelemetryGroupedSeriesPoint) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AgentTelemetryGroupedSeriesPoint) SetRows(v int32)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


