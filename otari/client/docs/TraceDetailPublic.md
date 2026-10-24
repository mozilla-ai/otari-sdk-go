# TraceDetailPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gates** | [**[]GateSummaryPublic**](GateSummaryPublic.md) | Each agent guardrail gate a coding agent&#39;s hook reported, the ones that did something first. | 
**Spans** | [**[]SpanPublic**](SpanPublic.md) |  | 
**State** | **string** |  | 
**Summary** | [**TraceSummaryPublic**](TraceSummaryPublic.md) |  | 
**ToolCosts** | [**[]ToolCostPublic**](ToolCostPublic.md) | An estimate for each tool call between two priced model requests, in the order they were made. | 
**Truncated** | **bool** | True when the session holds more spans than one read returns. | 
**Turns** | [**[]TurnPublic**](TurnPublic.md) |  | 

## Methods

### NewTraceDetailPublic

`func NewTraceDetailPublic(gates []GateSummaryPublic, spans []SpanPublic, state string, summary TraceSummaryPublic, toolCosts []ToolCostPublic, truncated bool, turns []TurnPublic, ) *TraceDetailPublic`

NewTraceDetailPublic instantiates a new TraceDetailPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceDetailPublicWithDefaults

`func NewTraceDetailPublicWithDefaults() *TraceDetailPublic`

NewTraceDetailPublicWithDefaults instantiates a new TraceDetailPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGates

`func (o *TraceDetailPublic) GetGates() []GateSummaryPublic`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *TraceDetailPublic) GetGatesOk() (*[]GateSummaryPublic, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *TraceDetailPublic) SetGates(v []GateSummaryPublic)`

SetGates sets Gates field to given value.


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


### GetToolCosts

`func (o *TraceDetailPublic) GetToolCosts() []ToolCostPublic`

GetToolCosts returns the ToolCosts field if non-nil, zero value otherwise.

### GetToolCostsOk

`func (o *TraceDetailPublic) GetToolCostsOk() (*[]ToolCostPublic, bool)`

GetToolCostsOk returns a tuple with the ToolCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCosts

`func (o *TraceDetailPublic) SetToolCosts(v []ToolCostPublic)`

SetToolCosts sets ToolCosts field to given value.


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


