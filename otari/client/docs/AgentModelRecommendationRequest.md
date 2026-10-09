# AgentModelRecommendationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | **string** | The subagent type: a built-in such as &#x60;Explore&#x60; or &#x60;Plan&#x60;, or a custom agent&#39;s name. | 
**Description** | Pointer to **string** | The caller&#39;s short description of the task. | [optional] [default to ""]
**Harness** | **string** | The agent harness asking, such as &#x60;claude-code&#x60;. | 
**ParentModel** | **string** | The model the parent conversation runs on, as the harness names it. | 
**Prompt** | **string** | The task the subagent is given. | 
**RequestedModel** | Pointer to **NullableString** | The model the caller asked for, if any. Sent as a fact; the recommendation is the gateway&#39;s. | [optional] 
**SessionId** | **string** | The harness&#39;s own id for the session the spawn happens in. | 
**ToolUseId** | **string** | The harness&#39;s own id for the tool call that spawns the subagent. | 
**User** | Pointer to **NullableString** | User ID the decision is billed to when asking with the master key; not sent upstream. | [optional] 

## Methods

### NewAgentModelRecommendationRequest

`func NewAgentModelRecommendationRequest(agentType string, harness string, parentModel string, prompt string, sessionId string, toolUseId string, ) *AgentModelRecommendationRequest`

NewAgentModelRecommendationRequest instantiates a new AgentModelRecommendationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentModelRecommendationRequestWithDefaults

`func NewAgentModelRecommendationRequestWithDefaults() *AgentModelRecommendationRequest`

NewAgentModelRecommendationRequestWithDefaults instantiates a new AgentModelRecommendationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *AgentModelRecommendationRequest) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentModelRecommendationRequest) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentModelRecommendationRequest) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetDescription

`func (o *AgentModelRecommendationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentModelRecommendationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentModelRecommendationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentModelRecommendationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHarness

`func (o *AgentModelRecommendationRequest) GetHarness() string`

GetHarness returns the Harness field if non-nil, zero value otherwise.

### GetHarnessOk

`func (o *AgentModelRecommendationRequest) GetHarnessOk() (*string, bool)`

GetHarnessOk returns a tuple with the Harness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarness

`func (o *AgentModelRecommendationRequest) SetHarness(v string)`

SetHarness sets Harness field to given value.


### GetParentModel

`func (o *AgentModelRecommendationRequest) GetParentModel() string`

GetParentModel returns the ParentModel field if non-nil, zero value otherwise.

### GetParentModelOk

`func (o *AgentModelRecommendationRequest) GetParentModelOk() (*string, bool)`

GetParentModelOk returns a tuple with the ParentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentModel

`func (o *AgentModelRecommendationRequest) SetParentModel(v string)`

SetParentModel sets ParentModel field to given value.


### GetPrompt

`func (o *AgentModelRecommendationRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgentModelRecommendationRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgentModelRecommendationRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequestedModel

`func (o *AgentModelRecommendationRequest) GetRequestedModel() string`

GetRequestedModel returns the RequestedModel field if non-nil, zero value otherwise.

### GetRequestedModelOk

`func (o *AgentModelRecommendationRequest) GetRequestedModelOk() (*string, bool)`

GetRequestedModelOk returns a tuple with the RequestedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedModel

`func (o *AgentModelRecommendationRequest) SetRequestedModel(v string)`

SetRequestedModel sets RequestedModel field to given value.

### HasRequestedModel

`func (o *AgentModelRecommendationRequest) HasRequestedModel() bool`

HasRequestedModel returns a boolean if a field has been set.

### SetRequestedModelNil

`func (o *AgentModelRecommendationRequest) SetRequestedModelNil(b bool)`

 SetRequestedModelNil sets the value for RequestedModel to be an explicit nil

### UnsetRequestedModel
`func (o *AgentModelRecommendationRequest) UnsetRequestedModel()`

UnsetRequestedModel ensures that no value is present for RequestedModel, not even an explicit nil
### GetSessionId

`func (o *AgentModelRecommendationRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AgentModelRecommendationRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AgentModelRecommendationRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetToolUseId

`func (o *AgentModelRecommendationRequest) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *AgentModelRecommendationRequest) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *AgentModelRecommendationRequest) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetUser

`func (o *AgentModelRecommendationRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AgentModelRecommendationRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AgentModelRecommendationRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AgentModelRecommendationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AgentModelRecommendationRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AgentModelRecommendationRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


