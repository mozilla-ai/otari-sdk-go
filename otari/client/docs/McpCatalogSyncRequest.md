# McpCatalogSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | Pointer to **NullableString** | A workspace the caller may add MCP servers to. Required unless the caller operates the deployment. | [optional] 

## Methods

### NewMcpCatalogSyncRequest

`func NewMcpCatalogSyncRequest() *McpCatalogSyncRequest`

NewMcpCatalogSyncRequest instantiates a new McpCatalogSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCatalogSyncRequestWithDefaults

`func NewMcpCatalogSyncRequestWithDefaults() *McpCatalogSyncRequest`

NewMcpCatalogSyncRequestWithDefaults instantiates a new McpCatalogSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *McpCatalogSyncRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *McpCatalogSyncRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *McpCatalogSyncRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *McpCatalogSyncRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *McpCatalogSyncRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *McpCatalogSyncRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


