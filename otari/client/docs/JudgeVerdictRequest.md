# JudgeVerdictRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GateId** | **string** |  | 
**Outcome** | **string** |  | 
**Reasoning** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewJudgeVerdictRequest

`func NewJudgeVerdictRequest(gateId string, outcome string, ) *JudgeVerdictRequest`

NewJudgeVerdictRequest instantiates a new JudgeVerdictRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeVerdictRequestWithDefaults

`func NewJudgeVerdictRequestWithDefaults() *JudgeVerdictRequest`

NewJudgeVerdictRequestWithDefaults instantiates a new JudgeVerdictRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateId

`func (o *JudgeVerdictRequest) GetGateId() string`

GetGateId returns the GateId field if non-nil, zero value otherwise.

### GetGateIdOk

`func (o *JudgeVerdictRequest) GetGateIdOk() (*string, bool)`

GetGateIdOk returns a tuple with the GateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateId

`func (o *JudgeVerdictRequest) SetGateId(v string)`

SetGateId sets GateId field to given value.


### GetOutcome

`func (o *JudgeVerdictRequest) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *JudgeVerdictRequest) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *JudgeVerdictRequest) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetReasoning

`func (o *JudgeVerdictRequest) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *JudgeVerdictRequest) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *JudgeVerdictRequest) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *JudgeVerdictRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


