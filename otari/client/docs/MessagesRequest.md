# MessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheControl** | Pointer to **map[string]interface{}** |  | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**MaxTokens** | **int32** |  | 
**MaxToolIterations** | Pointer to **NullableInt32** |  | [optional] 
**McpServerIds** | Pointer to **[]string** |  | [optional] 
**McpServers** | Pointer to [**[]McpServerConfig**](McpServerConfig.md) |  | [optional] 
**Messages** | **[]map[string]interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Model** | **string** |  | 
**SessionLabel** | Pointer to **NullableString** | Optional caller-supplied label for cost attribution (per run, experiment, or conversation). In hybrid mode it is forwarded onto the platform usage report so spend can be sliced by session without standing up OpenTelemetry. Stripped before the request is forwarded upstream to the provider. Has no effect in standalone mode, where there is no platform to report it to. | [optional] 
**StopSequences** | Pointer to **[]string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**System** | Pointer to [**NullableSystem**](System.md) |  | [optional] 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**Thinking** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolChoice** | Pointer to **map[string]interface{}** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ToolsHeader** | Pointer to **NullableString** |  | [optional] 
**TopK** | Pointer to **NullableInt32** |  | [optional] 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMessagesRequest

`func NewMessagesRequest(maxTokens int32, messages []map[string]interface{}, model string, ) *MessagesRequest`

NewMessagesRequest instantiates a new MessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesRequestWithDefaults

`func NewMessagesRequestWithDefaults() *MessagesRequest`

NewMessagesRequestWithDefaults instantiates a new MessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheControl

`func (o *MessagesRequest) GetCacheControl() map[string]interface{}`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MessagesRequest) GetCacheControlOk() (*map[string]interface{}, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MessagesRequest) SetCacheControl(v map[string]interface{})`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MessagesRequest) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MessagesRequest) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MessagesRequest) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetGuardrails

`func (o *MessagesRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *MessagesRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *MessagesRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *MessagesRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *MessagesRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *MessagesRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetMaxTokens

`func (o *MessagesRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *MessagesRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *MessagesRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.


### GetMaxToolIterations

`func (o *MessagesRequest) GetMaxToolIterations() int32`

GetMaxToolIterations returns the MaxToolIterations field if non-nil, zero value otherwise.

### GetMaxToolIterationsOk

`func (o *MessagesRequest) GetMaxToolIterationsOk() (*int32, bool)`

GetMaxToolIterationsOk returns a tuple with the MaxToolIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolIterations

`func (o *MessagesRequest) SetMaxToolIterations(v int32)`

SetMaxToolIterations sets MaxToolIterations field to given value.

### HasMaxToolIterations

`func (o *MessagesRequest) HasMaxToolIterations() bool`

HasMaxToolIterations returns a boolean if a field has been set.

### SetMaxToolIterationsNil

`func (o *MessagesRequest) SetMaxToolIterationsNil(b bool)`

 SetMaxToolIterationsNil sets the value for MaxToolIterations to be an explicit nil

### UnsetMaxToolIterations
`func (o *MessagesRequest) UnsetMaxToolIterations()`

UnsetMaxToolIterations ensures that no value is present for MaxToolIterations, not even an explicit nil
### GetMcpServerIds

`func (o *MessagesRequest) GetMcpServerIds() []string`

GetMcpServerIds returns the McpServerIds field if non-nil, zero value otherwise.

### GetMcpServerIdsOk

`func (o *MessagesRequest) GetMcpServerIdsOk() (*[]string, bool)`

GetMcpServerIdsOk returns a tuple with the McpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerIds

`func (o *MessagesRequest) SetMcpServerIds(v []string)`

SetMcpServerIds sets McpServerIds field to given value.

### HasMcpServerIds

`func (o *MessagesRequest) HasMcpServerIds() bool`

HasMcpServerIds returns a boolean if a field has been set.

### SetMcpServerIdsNil

`func (o *MessagesRequest) SetMcpServerIdsNil(b bool)`

 SetMcpServerIdsNil sets the value for McpServerIds to be an explicit nil

### UnsetMcpServerIds
`func (o *MessagesRequest) UnsetMcpServerIds()`

UnsetMcpServerIds ensures that no value is present for McpServerIds, not even an explicit nil
### GetMcpServers

`func (o *MessagesRequest) GetMcpServers() []McpServerConfig`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *MessagesRequest) GetMcpServersOk() (*[]McpServerConfig, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *MessagesRequest) SetMcpServers(v []McpServerConfig)`

SetMcpServers sets McpServers field to given value.

### HasMcpServers

`func (o *MessagesRequest) HasMcpServers() bool`

HasMcpServers returns a boolean if a field has been set.

### SetMcpServersNil

`func (o *MessagesRequest) SetMcpServersNil(b bool)`

 SetMcpServersNil sets the value for McpServers to be an explicit nil

### UnsetMcpServers
`func (o *MessagesRequest) UnsetMcpServers()`

UnsetMcpServers ensures that no value is present for McpServers, not even an explicit nil
### GetMessages

`func (o *MessagesRequest) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MessagesRequest) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MessagesRequest) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.


### GetMetadata

`func (o *MessagesRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MessagesRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MessagesRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MessagesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *MessagesRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *MessagesRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *MessagesRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MessagesRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MessagesRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSessionLabel

`func (o *MessagesRequest) GetSessionLabel() string`

GetSessionLabel returns the SessionLabel field if non-nil, zero value otherwise.

### GetSessionLabelOk

`func (o *MessagesRequest) GetSessionLabelOk() (*string, bool)`

GetSessionLabelOk returns a tuple with the SessionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLabel

`func (o *MessagesRequest) SetSessionLabel(v string)`

SetSessionLabel sets SessionLabel field to given value.

### HasSessionLabel

`func (o *MessagesRequest) HasSessionLabel() bool`

HasSessionLabel returns a boolean if a field has been set.

### SetSessionLabelNil

`func (o *MessagesRequest) SetSessionLabelNil(b bool)`

 SetSessionLabelNil sets the value for SessionLabel to be an explicit nil

### UnsetSessionLabel
`func (o *MessagesRequest) UnsetSessionLabel()`

UnsetSessionLabel ensures that no value is present for SessionLabel, not even an explicit nil
### GetStopSequences

`func (o *MessagesRequest) GetStopSequences() []string`

GetStopSequences returns the StopSequences field if non-nil, zero value otherwise.

### GetStopSequencesOk

`func (o *MessagesRequest) GetStopSequencesOk() (*[]string, bool)`

GetStopSequencesOk returns a tuple with the StopSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequences

`func (o *MessagesRequest) SetStopSequences(v []string)`

SetStopSequences sets StopSequences field to given value.

### HasStopSequences

`func (o *MessagesRequest) HasStopSequences() bool`

HasStopSequences returns a boolean if a field has been set.

### SetStopSequencesNil

`func (o *MessagesRequest) SetStopSequencesNil(b bool)`

 SetStopSequencesNil sets the value for StopSequences to be an explicit nil

### UnsetStopSequences
`func (o *MessagesRequest) UnsetStopSequences()`

UnsetStopSequences ensures that no value is present for StopSequences, not even an explicit nil
### GetStream

`func (o *MessagesRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *MessagesRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *MessagesRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *MessagesRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSystem

`func (o *MessagesRequest) GetSystem() System`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *MessagesRequest) GetSystemOk() (*System, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *MessagesRequest) SetSystem(v System)`

SetSystem sets System field to given value.

### HasSystem

`func (o *MessagesRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *MessagesRequest) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *MessagesRequest) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetTemperature

`func (o *MessagesRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *MessagesRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *MessagesRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *MessagesRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *MessagesRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *MessagesRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetThinking

`func (o *MessagesRequest) GetThinking() map[string]interface{}`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *MessagesRequest) GetThinkingOk() (*map[string]interface{}, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *MessagesRequest) SetThinking(v map[string]interface{})`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *MessagesRequest) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### SetThinkingNil

`func (o *MessagesRequest) SetThinkingNil(b bool)`

 SetThinkingNil sets the value for Thinking to be an explicit nil

### UnsetThinking
`func (o *MessagesRequest) UnsetThinking()`

UnsetThinking ensures that no value is present for Thinking, not even an explicit nil
### GetToolChoice

`func (o *MessagesRequest) GetToolChoice() map[string]interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *MessagesRequest) GetToolChoiceOk() (*map[string]interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *MessagesRequest) SetToolChoice(v map[string]interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *MessagesRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *MessagesRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *MessagesRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *MessagesRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MessagesRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MessagesRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *MessagesRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *MessagesRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *MessagesRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolsHeader

`func (o *MessagesRequest) GetToolsHeader() string`

GetToolsHeader returns the ToolsHeader field if non-nil, zero value otherwise.

### GetToolsHeaderOk

`func (o *MessagesRequest) GetToolsHeaderOk() (*string, bool)`

GetToolsHeaderOk returns a tuple with the ToolsHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsHeader

`func (o *MessagesRequest) SetToolsHeader(v string)`

SetToolsHeader sets ToolsHeader field to given value.

### HasToolsHeader

`func (o *MessagesRequest) HasToolsHeader() bool`

HasToolsHeader returns a boolean if a field has been set.

### SetToolsHeaderNil

`func (o *MessagesRequest) SetToolsHeaderNil(b bool)`

 SetToolsHeaderNil sets the value for ToolsHeader to be an explicit nil

### UnsetToolsHeader
`func (o *MessagesRequest) UnsetToolsHeader()`

UnsetToolsHeader ensures that no value is present for ToolsHeader, not even an explicit nil
### GetTopK

`func (o *MessagesRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *MessagesRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *MessagesRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *MessagesRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### SetTopKNil

`func (o *MessagesRequest) SetTopKNil(b bool)`

 SetTopKNil sets the value for TopK to be an explicit nil

### UnsetTopK
`func (o *MessagesRequest) UnsetTopK()`

UnsetTopK ensures that no value is present for TopK, not even an explicit nil
### GetTopP

`func (o *MessagesRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *MessagesRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *MessagesRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *MessagesRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *MessagesRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *MessagesRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


