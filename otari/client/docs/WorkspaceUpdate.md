# WorkspaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RequestMcpServersAllowed** | Pointer to **NullableBool** | Let this workspace&#39;s keys name MCP servers by URL in a request (mcp_servers), beside the stored servers whose access lists the workspace controls. | [optional] 
**TraceContentEnabled** | Pointer to **NullableBool** | Keep what this workspace&#39;s agent sessions were about: the prompts, commands and file paths its agents report. Only where the deployment keeps content at all. | [optional] 
**TraceContentToolOutput** | Pointer to **NullableBool** | Also keep what each tool returned, where this workspace keeps content and the deployment keeps tool output. Turning it off keeps none from the next request on, whatever the deployment allows. | [optional] 

## Methods

### NewWorkspaceUpdate

`func NewWorkspaceUpdate() *WorkspaceUpdate`

NewWorkspaceUpdate instantiates a new WorkspaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceUpdateWithDefaults

`func NewWorkspaceUpdateWithDefaults() *WorkspaceUpdate`

NewWorkspaceUpdateWithDefaults instantiates a new WorkspaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkspaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspaceUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *WorkspaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkspaceUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkspaceUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRequestMcpServersAllowed

`func (o *WorkspaceUpdate) GetRequestMcpServersAllowed() bool`

GetRequestMcpServersAllowed returns the RequestMcpServersAllowed field if non-nil, zero value otherwise.

### GetRequestMcpServersAllowedOk

`func (o *WorkspaceUpdate) GetRequestMcpServersAllowedOk() (*bool, bool)`

GetRequestMcpServersAllowedOk returns a tuple with the RequestMcpServersAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMcpServersAllowed

`func (o *WorkspaceUpdate) SetRequestMcpServersAllowed(v bool)`

SetRequestMcpServersAllowed sets RequestMcpServersAllowed field to given value.

### HasRequestMcpServersAllowed

`func (o *WorkspaceUpdate) HasRequestMcpServersAllowed() bool`

HasRequestMcpServersAllowed returns a boolean if a field has been set.

### SetRequestMcpServersAllowedNil

`func (o *WorkspaceUpdate) SetRequestMcpServersAllowedNil(b bool)`

 SetRequestMcpServersAllowedNil sets the value for RequestMcpServersAllowed to be an explicit nil

### UnsetRequestMcpServersAllowed
`func (o *WorkspaceUpdate) UnsetRequestMcpServersAllowed()`

UnsetRequestMcpServersAllowed ensures that no value is present for RequestMcpServersAllowed, not even an explicit nil
### GetTraceContentEnabled

`func (o *WorkspaceUpdate) GetTraceContentEnabled() bool`

GetTraceContentEnabled returns the TraceContentEnabled field if non-nil, zero value otherwise.

### GetTraceContentEnabledOk

`func (o *WorkspaceUpdate) GetTraceContentEnabledOk() (*bool, bool)`

GetTraceContentEnabledOk returns a tuple with the TraceContentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceContentEnabled

`func (o *WorkspaceUpdate) SetTraceContentEnabled(v bool)`

SetTraceContentEnabled sets TraceContentEnabled field to given value.

### HasTraceContentEnabled

`func (o *WorkspaceUpdate) HasTraceContentEnabled() bool`

HasTraceContentEnabled returns a boolean if a field has been set.

### SetTraceContentEnabledNil

`func (o *WorkspaceUpdate) SetTraceContentEnabledNil(b bool)`

 SetTraceContentEnabledNil sets the value for TraceContentEnabled to be an explicit nil

### UnsetTraceContentEnabled
`func (o *WorkspaceUpdate) UnsetTraceContentEnabled()`

UnsetTraceContentEnabled ensures that no value is present for TraceContentEnabled, not even an explicit nil
### GetTraceContentToolOutput

`func (o *WorkspaceUpdate) GetTraceContentToolOutput() bool`

GetTraceContentToolOutput returns the TraceContentToolOutput field if non-nil, zero value otherwise.

### GetTraceContentToolOutputOk

`func (o *WorkspaceUpdate) GetTraceContentToolOutputOk() (*bool, bool)`

GetTraceContentToolOutputOk returns a tuple with the TraceContentToolOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceContentToolOutput

`func (o *WorkspaceUpdate) SetTraceContentToolOutput(v bool)`

SetTraceContentToolOutput sets TraceContentToolOutput field to given value.

### HasTraceContentToolOutput

`func (o *WorkspaceUpdate) HasTraceContentToolOutput() bool`

HasTraceContentToolOutput returns a boolean if a field has been set.

### SetTraceContentToolOutputNil

`func (o *WorkspaceUpdate) SetTraceContentToolOutputNil(b bool)`

 SetTraceContentToolOutputNil sets the value for TraceContentToolOutput to be an explicit nil

### UnsetTraceContentToolOutput
`func (o *WorkspaceUpdate) UnsetTraceContentToolOutput()`

UnsetTraceContentToolOutput ensures that no value is present for TraceContentToolOutput, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


