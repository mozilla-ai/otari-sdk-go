# TraceDetailPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spans** | [**[]SpanPublic**](SpanPublic.md) |  | 
**State** | **string** |  | 
**Summary** | [**TraceSummaryPublic**](TraceSummaryPublic.md) |  | 
**Truncated** | **bool** | True when the session holds more spans than one read returns. | 
**Turns** | [**[]TurnPublic**](TurnPublic.md) |  | 

## Methods

### NewTraceDetailPublic

`func NewTraceDetailPublic(spans []SpanPublic, state string, summary TraceSummaryPublic, truncated bool, turns []TurnPublic, ) *TraceDetailPublic`

NewTraceDetailPublic instantiates a new TraceDetailPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceDetailPublicWithDefaults

`func NewTraceDetailPublicWithDefaults() *TraceDetailPublic`

NewTraceDetailPublicWithDefaults instantiates a new TraceDetailPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpans

`func (o *TraceDetailPublic) GetSpans() []SpanPublic`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *TraceDetailPublic) GetSpansOk() (*[]SpanPublic, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *TraceDetailPublic) SetSpans(v []SpanPublic)`

SetSpans sets Spans field to given value.


### GetState

`func (o *TraceDetailPublic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TraceDetailPublic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TraceDetailPublic) SetState(v string)`

SetState sets State field to given value.


### GetSummary

`func (o *TraceDetailPublic) GetSummary() TraceSummaryPublic`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TraceDetailPublic) GetSummaryOk() (*TraceSummaryPublic, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TraceDetailPublic) SetSummary(v TraceSummaryPublic)`

SetSummary sets Summary field to given value.


### GetTruncated

`func (o *TraceDetailPublic) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *TraceDetailPublic) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *TraceDetailPublic) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetTurns

`func (o *TraceDetailPublic) GetTurns() []TurnPublic`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *TraceDetailPublic) GetTurnsOk() (*[]TurnPublic, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *TraceDetailPublic) SetTurns(v []TurnPublic)`

SetTurns sets Turns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


