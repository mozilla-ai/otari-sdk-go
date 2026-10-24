# ToolCostPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calling** | **float32** | Its share of the output of the request that called it. | 
**CallsSharingRequest** | **int32** | Tool calls the request that called it made, sharing its output. | 
**CallsSharingResult** | **int32** | Tool calls whose results arrived together, sharing that growth. | 
**LaterRequests** | **int32** | The later requests that read its result again. | 
**Rereading** | **float32** | Reading its result again in later requests, until the context was compacted. | 
**ResultTokens** | **int32** | About how many tokens its result added to the context. | 
**SpanId** | **string** |  | 
**Subagent** | **NullableFloat32** | What a subagent the call ran cost, exactly. | 
**Total** | **float32** |  | 
**Writing** | **float32** | Writing its result into the context, on the next request. | 

## Methods

### NewToolCostPublic

`func NewToolCostPublic(calling float32, callsSharingRequest int32, callsSharingResult int32, laterRequests int32, rereading float32, resultTokens int32, spanId string, subagent NullableFloat32, total float32, writing float32, ) *ToolCostPublic`

NewToolCostPublic instantiates a new ToolCostPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCostPublicWithDefaults

`func NewToolCostPublicWithDefaults() *ToolCostPublic`

NewToolCostPublicWithDefaults instantiates a new ToolCostPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalling

`func (o *ToolCostPublic) GetCalling() float32`

GetCalling returns the Calling field if non-nil, zero value otherwise.

### GetCallingOk

`func (o *ToolCostPublic) GetCallingOk() (*float32, bool)`

GetCallingOk returns a tuple with the Calling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalling

`func (o *ToolCostPublic) SetCalling(v float32)`

SetCalling sets Calling field to given value.


### GetCallsSharingRequest

`func (o *ToolCostPublic) GetCallsSharingRequest() int32`

GetCallsSharingRequest returns the CallsSharingRequest field if non-nil, zero value otherwise.

### GetCallsSharingRequestOk

`func (o *ToolCostPublic) GetCallsSharingRequestOk() (*int32, bool)`

GetCallsSharingRequestOk returns a tuple with the CallsSharingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsSharingRequest

`func (o *ToolCostPublic) SetCallsSharingRequest(v int32)`

SetCallsSharingRequest sets CallsSharingRequest field to given value.


### GetCallsSharingResult

`func (o *ToolCostPublic) GetCallsSharingResult() int32`

GetCallsSharingResult returns the CallsSharingResult field if non-nil, zero value otherwise.

### GetCallsSharingResultOk

`func (o *ToolCostPublic) GetCallsSharingResultOk() (*int32, bool)`

GetCallsSharingResultOk returns a tuple with the CallsSharingResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsSharingResult

`func (o *ToolCostPublic) SetCallsSharingResult(v int32)`

SetCallsSharingResult sets CallsSharingResult field to given value.


### GetLaterRequests

`func (o *ToolCostPublic) GetLaterRequests() int32`

GetLaterRequests returns the LaterRequests field if non-nil, zero value otherwise.

### GetLaterRequestsOk

`func (o *ToolCostPublic) GetLaterRequestsOk() (*int32, bool)`

GetLaterRequestsOk returns a tuple with the LaterRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaterRequests

`func (o *ToolCostPublic) SetLaterRequests(v int32)`

SetLaterRequests sets LaterRequests field to given value.


### GetRereading

`func (o *ToolCostPublic) GetRereading() float32`

GetRereading returns the Rereading field if non-nil, zero value otherwise.

### GetRereadingOk

`func (o *ToolCostPublic) GetRereadingOk() (*float32, bool)`

GetRereadingOk returns a tuple with the Rereading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRereading

`func (o *ToolCostPublic) SetRereading(v float32)`

SetRereading sets Rereading field to given value.


### GetResultTokens

`func (o *ToolCostPublic) GetResultTokens() int32`

GetResultTokens returns the ResultTokens field if non-nil, zero value otherwise.

### GetResultTokensOk

`func (o *ToolCostPublic) GetResultTokensOk() (*int32, bool)`

GetResultTokensOk returns a tuple with the ResultTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTokens

`func (o *ToolCostPublic) SetResultTokens(v int32)`

SetResultTokens sets ResultTokens field to given value.


### GetSpanId

`func (o *ToolCostPublic) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *ToolCostPublic) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *ToolCostPublic) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSubagent

`func (o *ToolCostPublic) GetSubagent() float32`

GetSubagent returns the Subagent field if non-nil, zero value otherwise.

### GetSubagentOk

`func (o *ToolCostPublic) GetSubagentOk() (*float32, bool)`

GetSubagentOk returns a tuple with the Subagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubagent

`func (o *ToolCostPublic) SetSubagent(v float32)`

SetSubagent sets Subagent field to given value.


### SetSubagentNil

`func (o *ToolCostPublic) SetSubagentNil(b bool)`

 SetSubagentNil sets the value for Subagent to be an explicit nil

### UnsetSubagent
`func (o *ToolCostPublic) UnsetSubagent()`

UnsetSubagent ensures that no value is present for Subagent, not even an explicit nil
### GetTotal

`func (o *ToolCostPublic) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ToolCostPublic) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ToolCostPublic) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetWriting

`func (o *ToolCostPublic) GetWriting() float32`

GetWriting returns the Writing field if non-nil, zero value otherwise.

### GetWritingOk

`func (o *ToolCostPublic) GetWritingOk() (*float32, bool)`

GetWritingOk returns a tuple with the Writing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriting

`func (o *ToolCostPublic) SetWriting(v float32)`

SetWriting sets Writing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


