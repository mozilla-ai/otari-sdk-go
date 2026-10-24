# ProviderEndpointCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | **string** | Base URL of the endpoint. Refused when it resolves to a private, loopback, link-local or reserved address. | 
**ApiKey** | Pointer to **NullableString** | Sent to the endpoint. Never returned, only its last four. | [optional] 
**DefaultParams** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Name** | **string** | What callers put before the colon to reach this endpoint, as &#39;&lt;name&gt;:&lt;model&gt;&#39;. Letters, digits, &#39;.&#39;, &#39;_&#39; and &#39;-&#39;, starting with a letter or digit. It may not be a provider&#39;s name or a configured instance&#39;s. | 
**Provider** | **string** | The implementation that speaks to the endpoint: &#39;openai&#39; for an OpenAI-compatible server, &#39;anthropic&#39; for an Anthropic-compatible one. Whichever it is, callers may use any of the chat, responses and messages routes. | 
**UserId** | Pointer to **NullableString** | User that owns the endpoint. Omit for one every caller in the workspace reaches. A user&#39;s endpoint shadows a workspace-wide one of the same name for that user alone. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Workspace that owns the endpoint. Omit for the deployment&#39;s default workspace. | [optional] 

## Methods

### NewProviderEndpointCreateRequest

`func NewProviderEndpointCreateRequest(apiBase string, name string, provider string, ) *ProviderEndpointCreateRequest`

NewProviderEndpointCreateRequest instantiates a new ProviderEndpointCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderEndpointCreateRequestWithDefaults

`func NewProviderEndpointCreateRequestWithDefaults() *ProviderEndpointCreateRequest`

NewProviderEndpointCreateRequestWithDefaults instantiates a new ProviderEndpointCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ProviderEndpointCreateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ProviderEndpointCreateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ProviderEndpointCreateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.


### GetApiKey

`func (o *ProviderEndpointCreateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ProviderEndpointCreateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ProviderEndpointCreateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ProviderEndpointCreateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *ProviderEndpointCreateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *ProviderEndpointCreateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetDefaultParams

`func (o *ProviderEndpointCreateRequest) GetDefaultParams() map[string]interface{}`

GetDefaultParams returns the DefaultParams field if non-nil, zero value otherwise.

### GetDefaultParamsOk

`func (o *ProviderEndpointCreateRequest) GetDefaultParamsOk() (*map[string]interface{}, bool)`

GetDefaultParamsOk returns a tuple with the DefaultParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParams

`func (o *ProviderEndpointCreateRequest) SetDefaultParams(v map[string]interface{})`

SetDefaultParams sets DefaultParams field to given value.

### HasDefaultParams

`func (o *ProviderEndpointCreateRequest) HasDefaultParams() bool`

HasDefaultParams returns a boolean if a field has been set.

### SetDefaultParamsNil

`func (o *ProviderEndpointCreateRequest) SetDefaultParamsNil(b bool)`

 SetDefaultParamsNil sets the value for DefaultParams to be an explicit nil

### UnsetDefaultParams
`func (o *ProviderEndpointCreateRequest) UnsetDefaultParams()`

UnsetDefaultParams ensures that no value is present for DefaultParams, not even an explicit nil
### GetName

`func (o *ProviderEndpointCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderEndpointCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderEndpointCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *ProviderEndpointCreateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderEndpointCreateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderEndpointCreateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUserId

`func (o *ProviderEndpointCreateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProviderEndpointCreateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProviderEndpointCreateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ProviderEndpointCreateRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ProviderEndpointCreateRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ProviderEndpointCreateRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *ProviderEndpointCreateRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ProviderEndpointCreateRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ProviderEndpointCreateRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ProviderEndpointCreateRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *ProviderEndpointCreateRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *ProviderEndpointCreateRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


