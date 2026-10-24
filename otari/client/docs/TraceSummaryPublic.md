# TraceSummaryPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** |  | 
**Cost** | **float32** | USD, a snapshot of the usage rows&#39; cost. Usage stays the billing record. | 
**ErrorCount** | **int32** | Spans that ended in an error no fallback recovered, including tools&#39; errors, which do not fail the session. | 
**Harness** | **NullableString** | The agent or client that sent the requests, from its User-Agent. | 
**InputTokens** | **int32** |  | 
**LastActivityAt** | **time.Time** |  | 
**Name** | **NullableString** |  | 
**OutputTokens** | **int32** |  | 
**SessionSource** | **string** | What grouped the session&#39;s requests: the client&#39;s own id, its agent harness&#39;s id, an OTLP conversation id, or nothing (a single request). | 
**SpanCount** | **int32** |  | 
**StartedAt** | **time.Time** |  | 
**StepCount** | **int32** | Requests the gateway served in the session. | 
**TraceId** | **string** |  | 
**UserId** | **NullableString** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewTraceSummaryPublic

`func NewTraceSummaryPublic(apiKeyId NullableString, cost float32, errorCount int32, harness NullableString, inputTokens int32, lastActivityAt time.Time, name NullableString, outputTokens int32, sessionSource string, spanCount int32, startedAt time.Time, stepCount int32, traceId string, userId NullableString, workspaceId string, ) *TraceSummaryPublic`

NewTraceSummaryPublic instantiates a new TraceSummaryPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSummaryPublicWithDefaults

`func NewTraceSummaryPublicWithDefaults() *TraceSummaryPublic`

NewTraceSummaryPublicWithDefaults instantiates a new TraceSummaryPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *TraceSummaryPublic) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *TraceSummaryPublic) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *TraceSummaryPublic) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *TraceSummaryPublic) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *TraceSummaryPublic) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetCost

`func (o *TraceSummaryPublic) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TraceSummaryPublic) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TraceSummaryPublic) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetErrorCount

`func (o *TraceSummaryPublic) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *TraceSummaryPublic) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *TraceSummaryPublic) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetHarness

`func (o *TraceSummaryPublic) GetHarness() string`

GetHarness returns the Harness field if non-nil, zero value otherwise.

### GetHarnessOk

`func (o *TraceSummaryPublic) GetHarnessOk() (*string, bool)`

GetHarnessOk returns a tuple with the Harness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarness

`func (o *TraceSummaryPublic) SetHarness(v string)`

SetHarness sets Harness field to given value.


### SetHarnessNil

`func (o *TraceSummaryPublic) SetHarnessNil(b bool)`

 SetHarnessNil sets the value for Harness to be an explicit nil

### UnsetHarness
`func (o *TraceSummaryPublic) UnsetHarness()`

UnsetHarness ensures that no value is present for Harness, not even an explicit nil
### GetInputTokens

`func (o *TraceSummaryPublic) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *TraceSummaryPublic) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *TraceSummaryPublic) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetLastActivityAt

`func (o *TraceSummaryPublic) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *TraceSummaryPublic) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *TraceSummaryPublic) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.


### GetName

`func (o *TraceSummaryPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TraceSummaryPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TraceSummaryPublic) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *TraceSummaryPublic) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TraceSummaryPublic) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOutputTokens

`func (o *TraceSummaryPublic) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *TraceSummaryPublic) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *TraceSummaryPublic) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetSessionSource

`func (o *TraceSummaryPublic) GetSessionSource() string`

GetSessionSource returns the SessionSource field if non-nil, zero value otherwise.

### GetSessionSourceOk

`func (o *TraceSummaryPublic) GetSessionSourceOk() (*string, bool)`

GetSessionSourceOk returns a tuple with the SessionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSource

`func (o *TraceSummaryPublic) SetSessionSource(v string)`

SetSessionSource sets SessionSource field to given value.


### GetSpanCount

`func (o *TraceSummaryPublic) GetSpanCount() int32`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *TraceSummaryPublic) GetSpanCountOk() (*int32, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *TraceSummaryPublic) SetSpanCount(v int32)`

SetSpanCount sets SpanCount field to given value.


### GetStartedAt

`func (o *TraceSummaryPublic) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TraceSummaryPublic) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TraceSummaryPublic) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStepCount

`func (o *TraceSummaryPublic) GetStepCount() int32`

GetStepCount returns the StepCount field if non-nil, zero value otherwise.

### GetStepCountOk

`func (o *TraceSummaryPublic) GetStepCountOk() (*int32, bool)`

GetStepCountOk returns a tuple with the StepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCount

`func (o *TraceSummaryPublic) SetStepCount(v int32)`

SetStepCount sets StepCount field to given value.


### GetTraceId

`func (o *TraceSummaryPublic) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TraceSummaryPublic) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TraceSummaryPublic) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetUserId

`func (o *TraceSummaryPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TraceSummaryPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TraceSummaryPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *TraceSummaryPublic) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TraceSummaryPublic) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *TraceSummaryPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *TraceSummaryPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *TraceSummaryPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


