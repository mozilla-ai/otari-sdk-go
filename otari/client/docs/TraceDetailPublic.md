# TraceDetailPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentAccess** | Pointer to **string** | How this read reached the session&#39;s content: &#39;owner&#39; for its own user, &#39;share&#39; through a share its user made, &#39;break_glass&#39; for a recorded opening without one, &#39;none&#39; otherwise. | [optional] [default to "none"]
**ContentShare** | Pointer to [**NullableContentSharePublic**](ContentSharePublic.md) | The share the content was read through, with its note. Its reads are not listed. | [optional] 
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

### GetContentAccess

`func (o *TraceDetailPublic) GetContentAccess() string`

GetContentAccess returns the ContentAccess field if non-nil, zero value otherwise.

### GetContentAccessOk

`func (o *TraceDetailPublic) GetContentAccessOk() (*string, bool)`

GetContentAccessOk returns a tuple with the ContentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccess

`func (o *TraceDetailPublic) SetContentAccess(v string)`

SetContentAccess sets ContentAccess field to given value.

### HasContentAccess

`func (o *TraceDetailPublic) HasContentAccess() bool`

HasContentAccess returns a boolean if a field has been set.

### GetContentShare

`func (o *TraceDetailPublic) GetContentShare() ContentSharePublic`

GetContentShare returns the ContentShare field if non-nil, zero value otherwise.

### GetContentShareOk

`func (o *TraceDetailPublic) GetContentShareOk() (*ContentSharePublic, bool)`

GetContentShareOk returns a tuple with the ContentShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentShare

`func (o *TraceDetailPublic) SetContentShare(v ContentSharePublic)`

SetContentShare sets ContentShare field to given value.

### HasContentShare

`func (o *TraceDetailPublic) HasContentShare() bool`

HasContentShare returns a boolean if a field has been set.

### SetContentShareNil

`func (o *TraceDetailPublic) SetContentShareNil(b bool)`

 SetContentShareNil sets the value for ContentShare to be an explicit nil

### UnsetContentShare
`func (o *TraceDetailPublic) UnsetContentShare()`

UnsetContentShare ensures that no value is present for ContentShare, not even an explicit nil
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


