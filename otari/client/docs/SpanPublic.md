# SpanPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approximate** | **bool** | True when the start was inferred rather than measured. | 
**Attributes** | [**map[string]AttributesValue**](AttributesValue.md) |  | 
**Cost** | **NullableFloat32** |  | 
**DurationMs** | **NullableInt32** |  | 
**EndTime** | **NullableTime** |  | 
**ErrorClass** | **NullableString** |  | 
**InputTokens** | **NullableInt32** |  | 
**Kind** | **string** |  | 
**Model** | **NullableString** |  | 
**Name** | **string** |  | 
**OpensTurn** | **bool** |  | 
**Operation** | **NullableString** |  | 
**Origin** | **string** |  | 
**Outcome** | **string** |  | 
**OutputTokens** | **NullableInt32** |  | 
**ParentSpanId** | **NullableString** |  | 
**Provider** | **NullableString** |  | 
**Recovered** | **bool** |  | 
**RequestId** | **NullableString** |  | 
**SpanId** | **string** |  | 
**StartTime** | **NullableTime** |  | 
**ToolCallId** | **NullableString** |  | 
**ToolName** | **NullableString** |  | 
**ToolType** | **NullableString** |  | 

## Methods

### NewSpanPublic

`func NewSpanPublic(approximate bool, attributes map[string]AttributesValue, cost NullableFloat32, durationMs NullableInt32, endTime NullableTime, errorClass NullableString, inputTokens NullableInt32, kind string, model NullableString, name string, opensTurn bool, operation NullableString, origin string, outcome string, outputTokens NullableInt32, parentSpanId NullableString, provider NullableString, recovered bool, requestId NullableString, spanId string, startTime NullableTime, toolCallId NullableString, toolName NullableString, toolType NullableString, ) *SpanPublic`

NewSpanPublic instantiates a new SpanPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanPublicWithDefaults

`func NewSpanPublicWithDefaults() *SpanPublic`

NewSpanPublicWithDefaults instantiates a new SpanPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximate

`func (o *SpanPublic) GetApproximate() bool`

GetApproximate returns the Approximate field if non-nil, zero value otherwise.

### GetApproximateOk

`func (o *SpanPublic) GetApproximateOk() (*bool, bool)`

GetApproximateOk returns a tuple with the Approximate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximate

`func (o *SpanPublic) SetApproximate(v bool)`

SetApproximate sets Approximate field to given value.


### GetAttributes

`func (o *SpanPublic) GetAttributes() map[string]AttributesValue`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanPublic) GetAttributesOk() (*map[string]AttributesValue, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanPublic) SetAttributes(v map[string]AttributesValue)`

SetAttributes sets Attributes field to given value.


### GetCost

`func (o *SpanPublic) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SpanPublic) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SpanPublic) SetCost(v float32)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *SpanPublic) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *SpanPublic) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetDurationMs

`func (o *SpanPublic) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SpanPublic) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SpanPublic) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### SetDurationMsNil

`func (o *SpanPublic) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *SpanPublic) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetEndTime

`func (o *SpanPublic) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SpanPublic) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SpanPublic) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *SpanPublic) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *SpanPublic) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetErrorClass

`func (o *SpanPublic) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *SpanPublic) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *SpanPublic) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.


### SetErrorClassNil

`func (o *SpanPublic) SetErrorClassNil(b bool)`

 SetErrorClassNil sets the value for ErrorClass to be an explicit nil

### UnsetErrorClass
`func (o *SpanPublic) UnsetErrorClass()`

UnsetErrorClass ensures that no value is present for ErrorClass, not even an explicit nil
### GetInputTokens

`func (o *SpanPublic) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *SpanPublic) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *SpanPublic) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### SetInputTokensNil

`func (o *SpanPublic) SetInputTokensNil(b bool)`

 SetInputTokensNil sets the value for InputTokens to be an explicit nil

### UnsetInputTokens
`func (o *SpanPublic) UnsetInputTokens()`

UnsetInputTokens ensures that no value is present for InputTokens, not even an explicit nil
### GetKind

`func (o *SpanPublic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SpanPublic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SpanPublic) SetKind(v string)`

SetKind sets Kind field to given value.


### GetModel

`func (o *SpanPublic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SpanPublic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SpanPublic) SetModel(v string)`

SetModel sets Model field to given value.


### SetModelNil

`func (o *SpanPublic) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *SpanPublic) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetName

`func (o *SpanPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOpensTurn

`func (o *SpanPublic) GetOpensTurn() bool`

GetOpensTurn returns the OpensTurn field if non-nil, zero value otherwise.

### GetOpensTurnOk

`func (o *SpanPublic) GetOpensTurnOk() (*bool, bool)`

GetOpensTurnOk returns a tuple with the OpensTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensTurn

`func (o *SpanPublic) SetOpensTurn(v bool)`

SetOpensTurn sets OpensTurn field to given value.


### GetOperation

`func (o *SpanPublic) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SpanPublic) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SpanPublic) SetOperation(v string)`

SetOperation sets Operation field to given value.


### SetOperationNil

`func (o *SpanPublic) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *SpanPublic) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetOrigin

`func (o *SpanPublic) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SpanPublic) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SpanPublic) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetOutcome

`func (o *SpanPublic) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *SpanPublic) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *SpanPublic) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetOutputTokens

`func (o *SpanPublic) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *SpanPublic) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *SpanPublic) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### SetOutputTokensNil

`func (o *SpanPublic) SetOutputTokensNil(b bool)`

 SetOutputTokensNil sets the value for OutputTokens to be an explicit nil

### UnsetOutputTokens
`func (o *SpanPublic) UnsetOutputTokens()`

UnsetOutputTokens ensures that no value is present for OutputTokens, not even an explicit nil
### GetParentSpanId

`func (o *SpanPublic) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *SpanPublic) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *SpanPublic) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.


### SetParentSpanIdNil

`func (o *SpanPublic) SetParentSpanIdNil(b bool)`

 SetParentSpanIdNil sets the value for ParentSpanId to be an explicit nil

### UnsetParentSpanId
`func (o *SpanPublic) UnsetParentSpanId()`

UnsetParentSpanId ensures that no value is present for ParentSpanId, not even an explicit nil
### GetProvider

`func (o *SpanPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SpanPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SpanPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *SpanPublic) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *SpanPublic) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetRecovered

`func (o *SpanPublic) GetRecovered() bool`

GetRecovered returns the Recovered field if non-nil, zero value otherwise.

### GetRecoveredOk

`func (o *SpanPublic) GetRecoveredOk() (*bool, bool)`

GetRecoveredOk returns a tuple with the Recovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovered

`func (o *SpanPublic) SetRecovered(v bool)`

SetRecovered sets Recovered field to given value.


### GetRequestId

`func (o *SpanPublic) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SpanPublic) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SpanPublic) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### SetRequestIdNil

`func (o *SpanPublic) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *SpanPublic) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetSpanId

`func (o *SpanPublic) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanPublic) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanPublic) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetStartTime

`func (o *SpanPublic) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SpanPublic) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SpanPublic) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *SpanPublic) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *SpanPublic) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetToolCallId

`func (o *SpanPublic) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *SpanPublic) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *SpanPublic) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### SetToolCallIdNil

`func (o *SpanPublic) SetToolCallIdNil(b bool)`

 SetToolCallIdNil sets the value for ToolCallId to be an explicit nil

### UnsetToolCallId
`func (o *SpanPublic) UnsetToolCallId()`

UnsetToolCallId ensures that no value is present for ToolCallId, not even an explicit nil
### GetToolName

`func (o *SpanPublic) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *SpanPublic) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *SpanPublic) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### SetToolNameNil

`func (o *SpanPublic) SetToolNameNil(b bool)`

 SetToolNameNil sets the value for ToolName to be an explicit nil

### UnsetToolName
`func (o *SpanPublic) UnsetToolName()`

UnsetToolName ensures that no value is present for ToolName, not even an explicit nil
### GetToolType

`func (o *SpanPublic) GetToolType() string`

GetToolType returns the ToolType field if non-nil, zero value otherwise.

### GetToolTypeOk

`func (o *SpanPublic) GetToolTypeOk() (*string, bool)`

GetToolTypeOk returns a tuple with the ToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolType

`func (o *SpanPublic) SetToolType(v string)`

SetToolType sets ToolType field to given value.


### SetToolTypeNil

`func (o *SpanPublic) SetToolTypeNil(b bool)`

 SetToolTypeNil sets the value for ToolType to be an explicit nil

### UnsetToolType
`func (o *SpanPublic) UnsetToolType()`

UnsetToolType ensures that no value is present for ToolType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


