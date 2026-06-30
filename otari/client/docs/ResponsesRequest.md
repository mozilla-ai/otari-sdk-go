# ResponsesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Background** | Pointer to **NullableBool** |  | [optional] 
**Conversation** | Pointer to [**NullableConversation**](Conversation.md) |  | [optional] 
**FrequencyPenalty** | Pointer to **NullableFloat32** |  | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**Input** | **interface{}** |  | 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxToolCalls** | Pointer to **NullableInt32** |  | [optional] 
**MaxToolIterations** | Pointer to **NullableInt32** |  | [optional] 
**McpServerIds** | Pointer to **[]string** |  | [optional] 
**McpServers** | Pointer to [**[]McpServerConfig**](McpServerConfig.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Model** | **string** |  | 
**ParallelToolCalls** | Pointer to **NullableBool** |  | [optional] 
**PresencePenalty** | Pointer to **NullableFloat32** |  | [optional] 
**PreviousResponseId** | Pointer to **NullableString** |  | [optional] 
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**PromptCacheRetention** | Pointer to **NullableString** |  | [optional] 
**Reasoning** | Pointer to **map[string]interface{}** |  | [optional] 
**ResponseFormat** | Pointer to **map[string]interface{}** |  | [optional] 
**SafetyIdentifier** | Pointer to **NullableString** |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 
**SessionLabel** | Pointer to **NullableString** | Optional caller-supplied label for cost attribution (per run, experiment, or conversation). In hybrid mode it is forwarded onto the platform usage report so spend can be sliced by session without standing up OpenTelemetry. Stripped before the request is forwarded upstream to the provider. Has no effect in standalone mode, where there is no platform to report it to. | [optional] 
**Store** | Pointer to **NullableBool** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**StreamOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**Text** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolChoice** | Pointer to [**NullableToolChoice1**](ToolChoice1.md) |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolsHeader** | Pointer to **NullableString** |  | [optional] 
**TopLogprobs** | Pointer to **NullableInt32** |  | [optional] 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**Truncation** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesRequest

`func NewResponsesRequest(input interface{}, model string, ) *ResponsesRequest`

NewResponsesRequest instantiates a new ResponsesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesRequestWithDefaults

`func NewResponsesRequestWithDefaults() *ResponsesRequest`

NewResponsesRequestWithDefaults instantiates a new ResponsesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackground

`func (o *ResponsesRequest) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *ResponsesRequest) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *ResponsesRequest) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *ResponsesRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *ResponsesRequest) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *ResponsesRequest) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetConversation

`func (o *ResponsesRequest) GetConversation() Conversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *ResponsesRequest) GetConversationOk() (*Conversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *ResponsesRequest) SetConversation(v Conversation)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *ResponsesRequest) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### SetConversationNil

`func (o *ResponsesRequest) SetConversationNil(b bool)`

 SetConversationNil sets the value for Conversation to be an explicit nil

### UnsetConversation
`func (o *ResponsesRequest) UnsetConversation()`

UnsetConversation ensures that no value is present for Conversation, not even an explicit nil
### GetFrequencyPenalty

`func (o *ResponsesRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ResponsesRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ResponsesRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ResponsesRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *ResponsesRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *ResponsesRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetGuardrails

`func (o *ResponsesRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ResponsesRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ResponsesRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ResponsesRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *ResponsesRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *ResponsesRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetInclude

`func (o *ResponsesRequest) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ResponsesRequest) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ResponsesRequest) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *ResponsesRequest) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### SetIncludeNil

`func (o *ResponsesRequest) SetIncludeNil(b bool)`

 SetIncludeNil sets the value for Include to be an explicit nil

### UnsetInclude
`func (o *ResponsesRequest) UnsetInclude()`

UnsetInclude ensures that no value is present for Include, not even an explicit nil
### GetInput

`func (o *ResponsesRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ResponsesRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ResponsesRequest) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *ResponsesRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ResponsesRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetInstructions

`func (o *ResponsesRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ResponsesRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ResponsesRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ResponsesRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ResponsesRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ResponsesRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetMaxOutputTokens

`func (o *ResponsesRequest) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ResponsesRequest) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ResponsesRequest) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ResponsesRequest) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *ResponsesRequest) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *ResponsesRequest) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetMaxToolCalls

`func (o *ResponsesRequest) GetMaxToolCalls() int32`

GetMaxToolCalls returns the MaxToolCalls field if non-nil, zero value otherwise.

### GetMaxToolCallsOk

`func (o *ResponsesRequest) GetMaxToolCallsOk() (*int32, bool)`

GetMaxToolCallsOk returns a tuple with the MaxToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolCalls

`func (o *ResponsesRequest) SetMaxToolCalls(v int32)`

SetMaxToolCalls sets MaxToolCalls field to given value.

### HasMaxToolCalls

`func (o *ResponsesRequest) HasMaxToolCalls() bool`

HasMaxToolCalls returns a boolean if a field has been set.

### SetMaxToolCallsNil

`func (o *ResponsesRequest) SetMaxToolCallsNil(b bool)`

 SetMaxToolCallsNil sets the value for MaxToolCalls to be an explicit nil

### UnsetMaxToolCalls
`func (o *ResponsesRequest) UnsetMaxToolCalls()`

UnsetMaxToolCalls ensures that no value is present for MaxToolCalls, not even an explicit nil
### GetMaxToolIterations

`func (o *ResponsesRequest) GetMaxToolIterations() int32`

GetMaxToolIterations returns the MaxToolIterations field if non-nil, zero value otherwise.

### GetMaxToolIterationsOk

`func (o *ResponsesRequest) GetMaxToolIterationsOk() (*int32, bool)`

GetMaxToolIterationsOk returns a tuple with the MaxToolIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolIterations

`func (o *ResponsesRequest) SetMaxToolIterations(v int32)`

SetMaxToolIterations sets MaxToolIterations field to given value.

### HasMaxToolIterations

`func (o *ResponsesRequest) HasMaxToolIterations() bool`

HasMaxToolIterations returns a boolean if a field has been set.

### SetMaxToolIterationsNil

`func (o *ResponsesRequest) SetMaxToolIterationsNil(b bool)`

 SetMaxToolIterationsNil sets the value for MaxToolIterations to be an explicit nil

### UnsetMaxToolIterations
`func (o *ResponsesRequest) UnsetMaxToolIterations()`

UnsetMaxToolIterations ensures that no value is present for MaxToolIterations, not even an explicit nil
### GetMcpServerIds

`func (o *ResponsesRequest) GetMcpServerIds() []string`

GetMcpServerIds returns the McpServerIds field if non-nil, zero value otherwise.

### GetMcpServerIdsOk

`func (o *ResponsesRequest) GetMcpServerIdsOk() (*[]string, bool)`

GetMcpServerIdsOk returns a tuple with the McpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerIds

`func (o *ResponsesRequest) SetMcpServerIds(v []string)`

SetMcpServerIds sets McpServerIds field to given value.

### HasMcpServerIds

`func (o *ResponsesRequest) HasMcpServerIds() bool`

HasMcpServerIds returns a boolean if a field has been set.

### SetMcpServerIdsNil

`func (o *ResponsesRequest) SetMcpServerIdsNil(b bool)`

 SetMcpServerIdsNil sets the value for McpServerIds to be an explicit nil

### UnsetMcpServerIds
`func (o *ResponsesRequest) UnsetMcpServerIds()`

UnsetMcpServerIds ensures that no value is present for McpServerIds, not even an explicit nil
### GetMcpServers

`func (o *ResponsesRequest) GetMcpServers() []McpServerConfig`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *ResponsesRequest) GetMcpServersOk() (*[]McpServerConfig, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *ResponsesRequest) SetMcpServers(v []McpServerConfig)`

SetMcpServers sets McpServers field to given value.

### HasMcpServers

`func (o *ResponsesRequest) HasMcpServers() bool`

HasMcpServers returns a boolean if a field has been set.

### SetMcpServersNil

`func (o *ResponsesRequest) SetMcpServersNil(b bool)`

 SetMcpServersNil sets the value for McpServers to be an explicit nil

### UnsetMcpServers
`func (o *ResponsesRequest) UnsetMcpServers()`

UnsetMcpServers ensures that no value is present for McpServers, not even an explicit nil
### GetMetadata

`func (o *ResponsesRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponsesRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponsesRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponsesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ResponsesRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResponsesRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *ResponsesRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponsesRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponsesRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetParallelToolCalls

`func (o *ResponsesRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *ResponsesRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *ResponsesRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *ResponsesRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### SetParallelToolCallsNil

`func (o *ResponsesRequest) SetParallelToolCallsNil(b bool)`

 SetParallelToolCallsNil sets the value for ParallelToolCalls to be an explicit nil

### UnsetParallelToolCalls
`func (o *ResponsesRequest) UnsetParallelToolCalls()`

UnsetParallelToolCalls ensures that no value is present for ParallelToolCalls, not even an explicit nil
### GetPresencePenalty

`func (o *ResponsesRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ResponsesRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ResponsesRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ResponsesRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *ResponsesRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *ResponsesRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetPreviousResponseId

`func (o *ResponsesRequest) GetPreviousResponseId() string`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *ResponsesRequest) GetPreviousResponseIdOk() (*string, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *ResponsesRequest) SetPreviousResponseId(v string)`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *ResponsesRequest) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *ResponsesRequest) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *ResponsesRequest) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetPromptCacheKey

`func (o *ResponsesRequest) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *ResponsesRequest) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *ResponsesRequest) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *ResponsesRequest) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *ResponsesRequest) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *ResponsesRequest) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetPromptCacheRetention

`func (o *ResponsesRequest) GetPromptCacheRetention() string`

GetPromptCacheRetention returns the PromptCacheRetention field if non-nil, zero value otherwise.

### GetPromptCacheRetentionOk

`func (o *ResponsesRequest) GetPromptCacheRetentionOk() (*string, bool)`

GetPromptCacheRetentionOk returns a tuple with the PromptCacheRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheRetention

`func (o *ResponsesRequest) SetPromptCacheRetention(v string)`

SetPromptCacheRetention sets PromptCacheRetention field to given value.

### HasPromptCacheRetention

`func (o *ResponsesRequest) HasPromptCacheRetention() bool`

HasPromptCacheRetention returns a boolean if a field has been set.

### SetPromptCacheRetentionNil

`func (o *ResponsesRequest) SetPromptCacheRetentionNil(b bool)`

 SetPromptCacheRetentionNil sets the value for PromptCacheRetention to be an explicit nil

### UnsetPromptCacheRetention
`func (o *ResponsesRequest) UnsetPromptCacheRetention()`

UnsetPromptCacheRetention ensures that no value is present for PromptCacheRetention, not even an explicit nil
### GetReasoning

`func (o *ResponsesRequest) GetReasoning() map[string]interface{}`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ResponsesRequest) GetReasoningOk() (*map[string]interface{}, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ResponsesRequest) SetReasoning(v map[string]interface{})`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ResponsesRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *ResponsesRequest) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *ResponsesRequest) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetResponseFormat

`func (o *ResponsesRequest) GetResponseFormat() map[string]interface{}`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ResponsesRequest) GetResponseFormatOk() (*map[string]interface{}, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ResponsesRequest) SetResponseFormat(v map[string]interface{})`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ResponsesRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### SetResponseFormatNil

`func (o *ResponsesRequest) SetResponseFormatNil(b bool)`

 SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil

### UnsetResponseFormat
`func (o *ResponsesRequest) UnsetResponseFormat()`

UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
### GetSafetyIdentifier

`func (o *ResponsesRequest) GetSafetyIdentifier() string`

GetSafetyIdentifier returns the SafetyIdentifier field if non-nil, zero value otherwise.

### GetSafetyIdentifierOk

`func (o *ResponsesRequest) GetSafetyIdentifierOk() (*string, bool)`

GetSafetyIdentifierOk returns a tuple with the SafetyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyIdentifier

`func (o *ResponsesRequest) SetSafetyIdentifier(v string)`

SetSafetyIdentifier sets SafetyIdentifier field to given value.

### HasSafetyIdentifier

`func (o *ResponsesRequest) HasSafetyIdentifier() bool`

HasSafetyIdentifier returns a boolean if a field has been set.

### SetSafetyIdentifierNil

`func (o *ResponsesRequest) SetSafetyIdentifierNil(b bool)`

 SetSafetyIdentifierNil sets the value for SafetyIdentifier to be an explicit nil

### UnsetSafetyIdentifier
`func (o *ResponsesRequest) UnsetSafetyIdentifier()`

UnsetSafetyIdentifier ensures that no value is present for SafetyIdentifier, not even an explicit nil
### GetServiceTier

`func (o *ResponsesRequest) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ResponsesRequest) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ResponsesRequest) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ResponsesRequest) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *ResponsesRequest) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *ResponsesRequest) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetSessionLabel

`func (o *ResponsesRequest) GetSessionLabel() string`

GetSessionLabel returns the SessionLabel field if non-nil, zero value otherwise.

### GetSessionLabelOk

`func (o *ResponsesRequest) GetSessionLabelOk() (*string, bool)`

GetSessionLabelOk returns a tuple with the SessionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLabel

`func (o *ResponsesRequest) SetSessionLabel(v string)`

SetSessionLabel sets SessionLabel field to given value.

### HasSessionLabel

`func (o *ResponsesRequest) HasSessionLabel() bool`

HasSessionLabel returns a boolean if a field has been set.

### SetSessionLabelNil

`func (o *ResponsesRequest) SetSessionLabelNil(b bool)`

 SetSessionLabelNil sets the value for SessionLabel to be an explicit nil

### UnsetSessionLabel
`func (o *ResponsesRequest) UnsetSessionLabel()`

UnsetSessionLabel ensures that no value is present for SessionLabel, not even an explicit nil
### GetStore

`func (o *ResponsesRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ResponsesRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ResponsesRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ResponsesRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ResponsesRequest) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ResponsesRequest) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStream

`func (o *ResponsesRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ResponsesRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ResponsesRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ResponsesRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamOptions

`func (o *ResponsesRequest) GetStreamOptions() map[string]interface{}`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *ResponsesRequest) GetStreamOptionsOk() (*map[string]interface{}, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *ResponsesRequest) SetStreamOptions(v map[string]interface{})`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *ResponsesRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.

### SetStreamOptionsNil

`func (o *ResponsesRequest) SetStreamOptionsNil(b bool)`

 SetStreamOptionsNil sets the value for StreamOptions to be an explicit nil

### UnsetStreamOptions
`func (o *ResponsesRequest) UnsetStreamOptions()`

UnsetStreamOptions ensures that no value is present for StreamOptions, not even an explicit nil
### GetTemperature

`func (o *ResponsesRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ResponsesRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ResponsesRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ResponsesRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *ResponsesRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *ResponsesRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetText

`func (o *ResponsesRequest) GetText() map[string]interface{}`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResponsesRequest) GetTextOk() (*map[string]interface{}, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResponsesRequest) SetText(v map[string]interface{})`

SetText sets Text field to given value.

### HasText

`func (o *ResponsesRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ResponsesRequest) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ResponsesRequest) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetToolChoice

`func (o *ResponsesRequest) GetToolChoice() ToolChoice1`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ResponsesRequest) GetToolChoiceOk() (*ToolChoice1, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ResponsesRequest) SetToolChoice(v ToolChoice1)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ResponsesRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *ResponsesRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *ResponsesRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *ResponsesRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ResponsesRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ResponsesRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *ResponsesRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *ResponsesRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *ResponsesRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolsHeader

`func (o *ResponsesRequest) GetToolsHeader() string`

GetToolsHeader returns the ToolsHeader field if non-nil, zero value otherwise.

### GetToolsHeaderOk

`func (o *ResponsesRequest) GetToolsHeaderOk() (*string, bool)`

GetToolsHeaderOk returns a tuple with the ToolsHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsHeader

`func (o *ResponsesRequest) SetToolsHeader(v string)`

SetToolsHeader sets ToolsHeader field to given value.

### HasToolsHeader

`func (o *ResponsesRequest) HasToolsHeader() bool`

HasToolsHeader returns a boolean if a field has been set.

### SetToolsHeaderNil

`func (o *ResponsesRequest) SetToolsHeaderNil(b bool)`

 SetToolsHeaderNil sets the value for ToolsHeader to be an explicit nil

### UnsetToolsHeader
`func (o *ResponsesRequest) UnsetToolsHeader()`

UnsetToolsHeader ensures that no value is present for ToolsHeader, not even an explicit nil
### GetTopLogprobs

`func (o *ResponsesRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ResponsesRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ResponsesRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ResponsesRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ResponsesRequest) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ResponsesRequest) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
### GetTopP

`func (o *ResponsesRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ResponsesRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ResponsesRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ResponsesRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *ResponsesRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *ResponsesRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetTruncation

`func (o *ResponsesRequest) GetTruncation() string`

GetTruncation returns the Truncation field if non-nil, zero value otherwise.

### GetTruncationOk

`func (o *ResponsesRequest) GetTruncationOk() (*string, bool)`

GetTruncationOk returns a tuple with the Truncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncation

`func (o *ResponsesRequest) SetTruncation(v string)`

SetTruncation sets Truncation field to given value.

### HasTruncation

`func (o *ResponsesRequest) HasTruncation() bool`

HasTruncation returns a boolean if a field has been set.

### SetTruncationNil

`func (o *ResponsesRequest) SetTruncationNil(b bool)`

 SetTruncationNil sets the value for Truncation to be an explicit nil

### UnsetTruncation
`func (o *ResponsesRequest) UnsetTruncation()`

UnsetTruncation ensures that no value is present for Truncation, not even an explicit nil
### GetUser

`func (o *ResponsesRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ResponsesRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ResponsesRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ResponsesRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ResponsesRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ResponsesRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


