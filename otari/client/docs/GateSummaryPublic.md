# GateSummaryPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **int32** | Checks where the gate stopped the agent. | 
**Enforcement** | **NullableString** |  | 
**Flagged** | **int32** | Checks an advisory gate failed; the agent went ahead. | 
**GateId** | **string** |  | 
**GateType** | **NullableString** | From the guardrail&#39;s definition, when it was sent. | 
**Message** | **NullableString** | What the gate says when it fails, from the guardrail&#39;s definition. | 
**Origin** | **NullableString** | Whether the gate is the repository&#39;s (.otari/) or the user&#39;s own (~/.otari/). | 
**Passed** | **int32** |  | 
**Refused** | **int32** | Checks where the gate could not decide and held the agent anyway. | 
**Runs** | **[]string** | The moments the gate runs at (pre_tool_use.edit_target, stop.session, ...). | 
**SpanId** | **NullableString** | The check to open first: the earliest that did anything but pass. Null for a gate that never ran. | 
**Undecided** | **int32** | Checks where the gate could not decide; the agent went ahead unchecked. | 

## Methods

### NewGateSummaryPublic

`func NewGateSummaryPublic(blocked int32, enforcement NullableString, flagged int32, gateId string, gateType NullableString, message NullableString, origin NullableString, passed int32, refused int32, runs []string, spanId NullableString, undecided int32, ) *GateSummaryPublic`

NewGateSummaryPublic instantiates a new GateSummaryPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateSummaryPublicWithDefaults

`func NewGateSummaryPublicWithDefaults() *GateSummaryPublic`

NewGateSummaryPublicWithDefaults instantiates a new GateSummaryPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *GateSummaryPublic) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *GateSummaryPublic) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *GateSummaryPublic) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.


### GetEnforcement

`func (o *GateSummaryPublic) GetEnforcement() string`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *GateSummaryPublic) GetEnforcementOk() (*string, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *GateSummaryPublic) SetEnforcement(v string)`

SetEnforcement sets Enforcement field to given value.


### SetEnforcementNil

`func (o *GateSummaryPublic) SetEnforcementNil(b bool)`

 SetEnforcementNil sets the value for Enforcement to be an explicit nil

### UnsetEnforcement
`func (o *GateSummaryPublic) UnsetEnforcement()`

UnsetEnforcement ensures that no value is present for Enforcement, not even an explicit nil
### GetFlagged

`func (o *GateSummaryPublic) GetFlagged() int32`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *GateSummaryPublic) GetFlaggedOk() (*int32, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *GateSummaryPublic) SetFlagged(v int32)`

SetFlagged sets Flagged field to given value.


### GetGateId

`func (o *GateSummaryPublic) GetGateId() string`

GetGateId returns the GateId field if non-nil, zero value otherwise.

### GetGateIdOk

`func (o *GateSummaryPublic) GetGateIdOk() (*string, bool)`

GetGateIdOk returns a tuple with the GateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateId

`func (o *GateSummaryPublic) SetGateId(v string)`

SetGateId sets GateId field to given value.


### GetGateType

`func (o *GateSummaryPublic) GetGateType() string`

GetGateType returns the GateType field if non-nil, zero value otherwise.

### GetGateTypeOk

`func (o *GateSummaryPublic) GetGateTypeOk() (*string, bool)`

GetGateTypeOk returns a tuple with the GateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateType

`func (o *GateSummaryPublic) SetGateType(v string)`

SetGateType sets GateType field to given value.


### SetGateTypeNil

`func (o *GateSummaryPublic) SetGateTypeNil(b bool)`

 SetGateTypeNil sets the value for GateType to be an explicit nil

### UnsetGateType
`func (o *GateSummaryPublic) UnsetGateType()`

UnsetGateType ensures that no value is present for GateType, not even an explicit nil
### GetMessage

`func (o *GateSummaryPublic) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GateSummaryPublic) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GateSummaryPublic) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *GateSummaryPublic) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GateSummaryPublic) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOrigin

`func (o *GateSummaryPublic) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GateSummaryPublic) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GateSummaryPublic) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### SetOriginNil

`func (o *GateSummaryPublic) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *GateSummaryPublic) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPassed

`func (o *GateSummaryPublic) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *GateSummaryPublic) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *GateSummaryPublic) SetPassed(v int32)`

SetPassed sets Passed field to given value.


### GetRefused

`func (o *GateSummaryPublic) GetRefused() int32`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GateSummaryPublic) GetRefusedOk() (*int32, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GateSummaryPublic) SetRefused(v int32)`

SetRefused sets Refused field to given value.


### GetRuns

`func (o *GateSummaryPublic) GetRuns() []string`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *GateSummaryPublic) GetRunsOk() (*[]string, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *GateSummaryPublic) SetRuns(v []string)`

SetRuns sets Runs field to given value.


### GetSpanId

`func (o *GateSummaryPublic) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *GateSummaryPublic) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *GateSummaryPublic) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### SetSpanIdNil

`func (o *GateSummaryPublic) SetSpanIdNil(b bool)`

 SetSpanIdNil sets the value for SpanId to be an explicit nil

### UnsetSpanId
`func (o *GateSummaryPublic) UnsetSpanId()`

UnsetSpanId ensures that no value is present for SpanId, not even an explicit nil
### GetUndecided

`func (o *GateSummaryPublic) GetUndecided() int32`

GetUndecided returns the Undecided field if non-nil, zero value otherwise.

### GetUndecidedOk

`func (o *GateSummaryPublic) GetUndecidedOk() (*int32, bool)`

GetUndecidedOk returns a tuple with the Undecided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndecided

`func (o *GateSummaryPublic) SetUndecided(v int32)`

SetUndecided sets Undecided field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


