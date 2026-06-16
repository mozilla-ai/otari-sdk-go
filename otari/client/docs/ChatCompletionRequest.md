# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrequencyPenalty** | Pointer to **NullableFloat32** |  | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**LogitBias** | Pointer to **map[string]float32** |  | [optional] 
**Logprobs** | Pointer to **NullableBool** |  | [optional] 
**MaxCompletionTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxToolIterations** | Pointer to **NullableInt32** |  | [optional] 
**McpServerIds** | Pointer to **[]string** |  | [optional] 
**McpServers** | Pointer to [**[]McpServerConfig**](McpServerConfig.md) |  | [optional] 
**Messages** | [**[]ChatMessageInput**](ChatMessageInput.md) |  | 
**Model** | **string** |  | 
**N** | Pointer to **NullableInt32** |  | [optional] 
**ParallelToolCalls** | Pointer to **NullableBool** |  | [optional] 
**PresencePenalty** | Pointer to **NullableFloat32** |  | [optional] 
**ReasoningEffort** | Pointer to **NullableString** |  | [optional] 
**ResponseFormat** | Pointer to **map[string]interface{}** |  | [optional] 
**Seed** | Pointer to **NullableInt32** |  | [optional] 
**Stop** | Pointer to [**NullableStop**](Stop.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**StreamOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**ToolChoice** | Pointer to [**NullableToolChoice**](ToolChoice.md) |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolsHeader** | Pointer to **NullableString** | Optional override for the lead-in that the gateway prepends before the per-tool hint block in the system message. Useful for expressing global tool-selection policy (e.g. &#39;prefer MCP tools over code_execution&#39;). Falls back to OTARI_TOOLS_HEADER env, then to the built-in default. | [optional] 
**TopLogprobs** | Pointer to **NullableInt32** |  | [optional] 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(messages []ChatMessageInput, model string, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequencyPenalty

`func (o *ChatCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ChatCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ChatCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *ChatCompletionRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *ChatCompletionRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetGuardrails

`func (o *ChatCompletionRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ChatCompletionRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ChatCompletionRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ChatCompletionRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *ChatCompletionRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *ChatCompletionRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetLogitBias

`func (o *ChatCompletionRequest) GetLogitBias() map[string]float32`

GetLogitBias returns the LogitBias field if non-nil, zero value otherwise.

### GetLogitBiasOk

`func (o *ChatCompletionRequest) GetLogitBiasOk() (*map[string]float32, bool)`

GetLogitBiasOk returns a tuple with the LogitBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogitBias

`func (o *ChatCompletionRequest) SetLogitBias(v map[string]float32)`

SetLogitBias sets LogitBias field to given value.

### HasLogitBias

`func (o *ChatCompletionRequest) HasLogitBias() bool`

HasLogitBias returns a boolean if a field has been set.

### SetLogitBiasNil

`func (o *ChatCompletionRequest) SetLogitBiasNil(b bool)`

 SetLogitBiasNil sets the value for LogitBias to be an explicit nil

### UnsetLogitBias
`func (o *ChatCompletionRequest) UnsetLogitBias()`

UnsetLogitBias ensures that no value is present for LogitBias, not even an explicit nil
### GetLogprobs

`func (o *ChatCompletionRequest) GetLogprobs() bool`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ChatCompletionRequest) GetLogprobsOk() (*bool, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ChatCompletionRequest) SetLogprobs(v bool)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ChatCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *ChatCompletionRequest) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *ChatCompletionRequest) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetMaxCompletionTokens

`func (o *ChatCompletionRequest) GetMaxCompletionTokens() int32`

GetMaxCompletionTokens returns the MaxCompletionTokens field if non-nil, zero value otherwise.

### GetMaxCompletionTokensOk

`func (o *ChatCompletionRequest) GetMaxCompletionTokensOk() (*int32, bool)`

GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompletionTokens

`func (o *ChatCompletionRequest) SetMaxCompletionTokens(v int32)`

SetMaxCompletionTokens sets MaxCompletionTokens field to given value.

### HasMaxCompletionTokens

`func (o *ChatCompletionRequest) HasMaxCompletionTokens() bool`

HasMaxCompletionTokens returns a boolean if a field has been set.

### SetMaxCompletionTokensNil

`func (o *ChatCompletionRequest) SetMaxCompletionTokensNil(b bool)`

 SetMaxCompletionTokensNil sets the value for MaxCompletionTokens to be an explicit nil

### UnsetMaxCompletionTokens
`func (o *ChatCompletionRequest) UnsetMaxCompletionTokens()`

UnsetMaxCompletionTokens ensures that no value is present for MaxCompletionTokens, not even an explicit nil
### GetMaxTokens

`func (o *ChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *ChatCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *ChatCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetMaxToolIterations

`func (o *ChatCompletionRequest) GetMaxToolIterations() int32`

GetMaxToolIterations returns the MaxToolIterations field if non-nil, zero value otherwise.

### GetMaxToolIterationsOk

`func (o *ChatCompletionRequest) GetMaxToolIterationsOk() (*int32, bool)`

GetMaxToolIterationsOk returns a tuple with the MaxToolIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolIterations

`func (o *ChatCompletionRequest) SetMaxToolIterations(v int32)`

SetMaxToolIterations sets MaxToolIterations field to given value.

### HasMaxToolIterations

`func (o *ChatCompletionRequest) HasMaxToolIterations() bool`

HasMaxToolIterations returns a boolean if a field has been set.

### SetMaxToolIterationsNil

`func (o *ChatCompletionRequest) SetMaxToolIterationsNil(b bool)`

 SetMaxToolIterationsNil sets the value for MaxToolIterations to be an explicit nil

### UnsetMaxToolIterations
`func (o *ChatCompletionRequest) UnsetMaxToolIterations()`

UnsetMaxToolIterations ensures that no value is present for MaxToolIterations, not even an explicit nil
### GetMcpServerIds

`func (o *ChatCompletionRequest) GetMcpServerIds() []string`

GetMcpServerIds returns the McpServerIds field if non-nil, zero value otherwise.

### GetMcpServerIdsOk

`func (o *ChatCompletionRequest) GetMcpServerIdsOk() (*[]string, bool)`

GetMcpServerIdsOk returns a tuple with the McpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerIds

`func (o *ChatCompletionRequest) SetMcpServerIds(v []string)`

SetMcpServerIds sets McpServerIds field to given value.

### HasMcpServerIds

`func (o *ChatCompletionRequest) HasMcpServerIds() bool`

HasMcpServerIds returns a boolean if a field has been set.

### SetMcpServerIdsNil

`func (o *ChatCompletionRequest) SetMcpServerIdsNil(b bool)`

 SetMcpServerIdsNil sets the value for McpServerIds to be an explicit nil

### UnsetMcpServerIds
`func (o *ChatCompletionRequest) UnsetMcpServerIds()`

UnsetMcpServerIds ensures that no value is present for McpServerIds, not even an explicit nil
### GetMcpServers

`func (o *ChatCompletionRequest) GetMcpServers() []McpServerConfig`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *ChatCompletionRequest) GetMcpServersOk() (*[]McpServerConfig, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *ChatCompletionRequest) SetMcpServers(v []McpServerConfig)`

SetMcpServers sets McpServers field to given value.

### HasMcpServers

`func (o *ChatCompletionRequest) HasMcpServers() bool`

HasMcpServers returns a boolean if a field has been set.

### SetMcpServersNil

`func (o *ChatCompletionRequest) SetMcpServersNil(b bool)`

 SetMcpServersNil sets the value for McpServers to be an explicit nil

### UnsetMcpServers
`func (o *ChatCompletionRequest) UnsetMcpServers()`

UnsetMcpServers ensures that no value is present for McpServers, not even an explicit nil
### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []ChatMessageInput`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]ChatMessageInput, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []ChatMessageInput)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *ChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetN

`func (o *ChatCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ChatCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ChatCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *ChatCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *ChatCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetParallelToolCalls

`func (o *ChatCompletionRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *ChatCompletionRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *ChatCompletionRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *ChatCompletionRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### SetParallelToolCallsNil

`func (o *ChatCompletionRequest) SetParallelToolCallsNil(b bool)`

 SetParallelToolCallsNil sets the value for ParallelToolCalls to be an explicit nil

### UnsetParallelToolCalls
`func (o *ChatCompletionRequest) UnsetParallelToolCalls()`

UnsetParallelToolCalls ensures that no value is present for ParallelToolCalls, not even an explicit nil
### GetPresencePenalty

`func (o *ChatCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ChatCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ChatCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *ChatCompletionRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *ChatCompletionRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetReasoningEffort

`func (o *ChatCompletionRequest) GetReasoningEffort() string`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *ChatCompletionRequest) GetReasoningEffortOk() (*string, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *ChatCompletionRequest) SetReasoningEffort(v string)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *ChatCompletionRequest) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.

### SetReasoningEffortNil

`func (o *ChatCompletionRequest) SetReasoningEffortNil(b bool)`

 SetReasoningEffortNil sets the value for ReasoningEffort to be an explicit nil

### UnsetReasoningEffort
`func (o *ChatCompletionRequest) UnsetReasoningEffort()`

UnsetReasoningEffort ensures that no value is present for ReasoningEffort, not even an explicit nil
### GetResponseFormat

`func (o *ChatCompletionRequest) GetResponseFormat() map[string]interface{}`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatCompletionRequest) GetResponseFormatOk() (*map[string]interface{}, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatCompletionRequest) SetResponseFormat(v map[string]interface{})`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### SetResponseFormatNil

`func (o *ChatCompletionRequest) SetResponseFormatNil(b bool)`

 SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil

### UnsetResponseFormat
`func (o *ChatCompletionRequest) UnsetResponseFormat()`

UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
### GetSeed

`func (o *ChatCompletionRequest) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *ChatCompletionRequest) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *ChatCompletionRequest) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *ChatCompletionRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### SetSeedNil

`func (o *ChatCompletionRequest) SetSeedNil(b bool)`

 SetSeedNil sets the value for Seed to be an explicit nil

### UnsetSeed
`func (o *ChatCompletionRequest) UnsetSeed()`

UnsetSeed ensures that no value is present for Seed, not even an explicit nil
### GetStop

`func (o *ChatCompletionRequest) GetStop() Stop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ChatCompletionRequest) GetStopOk() (*Stop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ChatCompletionRequest) SetStop(v Stop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ChatCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *ChatCompletionRequest) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *ChatCompletionRequest) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamOptions

`func (o *ChatCompletionRequest) GetStreamOptions() map[string]interface{}`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *ChatCompletionRequest) GetStreamOptionsOk() (*map[string]interface{}, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *ChatCompletionRequest) SetStreamOptions(v map[string]interface{})`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *ChatCompletionRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.

### SetStreamOptionsNil

`func (o *ChatCompletionRequest) SetStreamOptionsNil(b bool)`

 SetStreamOptionsNil sets the value for StreamOptions to be an explicit nil

### UnsetStreamOptions
`func (o *ChatCompletionRequest) UnsetStreamOptions()`

UnsetStreamOptions ensures that no value is present for StreamOptions, not even an explicit nil
### GetTemperature

`func (o *ChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *ChatCompletionRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *ChatCompletionRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetToolChoice

`func (o *ChatCompletionRequest) GetToolChoice() ToolChoice`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ChatCompletionRequest) GetToolChoiceOk() (*ToolChoice, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ChatCompletionRequest) SetToolChoice(v ToolChoice)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *ChatCompletionRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *ChatCompletionRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *ChatCompletionRequest) GetTools() []*map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatCompletionRequest) GetToolsOk() (*[]*map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatCompletionRequest) SetTools(v []*map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *ChatCompletionRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *ChatCompletionRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolsHeader

`func (o *ChatCompletionRequest) GetToolsHeader() string`

GetToolsHeader returns the ToolsHeader field if non-nil, zero value otherwise.

### GetToolsHeaderOk

`func (o *ChatCompletionRequest) GetToolsHeaderOk() (*string, bool)`

GetToolsHeaderOk returns a tuple with the ToolsHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsHeader

`func (o *ChatCompletionRequest) SetToolsHeader(v string)`

SetToolsHeader sets ToolsHeader field to given value.

### HasToolsHeader

`func (o *ChatCompletionRequest) HasToolsHeader() bool`

HasToolsHeader returns a boolean if a field has been set.

### SetToolsHeaderNil

`func (o *ChatCompletionRequest) SetToolsHeaderNil(b bool)`

 SetToolsHeaderNil sets the value for ToolsHeader to be an explicit nil

### UnsetToolsHeader
`func (o *ChatCompletionRequest) UnsetToolsHeader()`

UnsetToolsHeader ensures that no value is present for ToolsHeader, not even an explicit nil
### GetTopLogprobs

`func (o *ChatCompletionRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ChatCompletionRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ChatCompletionRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ChatCompletionRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ChatCompletionRequest) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ChatCompletionRequest) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
### GetTopP

`func (o *ChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *ChatCompletionRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *ChatCompletionRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetUser

`func (o *ChatCompletionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatCompletionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatCompletionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatCompletionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ChatCompletionRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ChatCompletionRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


