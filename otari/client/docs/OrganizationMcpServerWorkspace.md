# OrganizationMcpServerWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Narrowed** | **bool** | The workspace&#39;s admins changed something of their own: turned the server off, or set its default access or whether it is attached | 
**OrganizationTools** | **[]string** | The organization&#39;s limit for the workspace; null sets none, [] turns the server off there | 
**Reached** | **bool** | The workspace&#39;s requests reach the server: neither the organization nor the workspace turned it off there, and the server itself is on | 
**WorkspaceEnabled** | **bool** | The workspace&#39;s own switch, which its admins set | 
**WorkspaceId** | **string** |  | 

## Methods

### NewOrganizationMcpServerWorkspace

`func NewOrganizationMcpServerWorkspace(name string, narrowed bool, organizationTools []string, reached bool, workspaceEnabled bool, workspaceId string, ) *OrganizationMcpServerWorkspace`

NewOrganizationMcpServerWorkspace instantiates a new OrganizationMcpServerWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMcpServerWorkspaceWithDefaults

`func NewOrganizationMcpServerWorkspaceWithDefaults() *OrganizationMcpServerWorkspace`

NewOrganizationMcpServerWorkspaceWithDefaults instantiates a new OrganizationMcpServerWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationMcpServerWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationMcpServerWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationMcpServerWorkspace) SetName(v string)`

SetName sets Name field to given value.


### GetNarrowed

`func (o *OrganizationMcpServerWorkspace) GetNarrowed() bool`

GetNarrowed returns the Narrowed field if non-nil, zero value otherwise.

### GetNarrowedOk

`func (o *OrganizationMcpServerWorkspace) GetNarrowedOk() (*bool, bool)`

GetNarrowedOk returns a tuple with the Narrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrowed

`func (o *OrganizationMcpServerWorkspace) SetNarrowed(v bool)`

SetNarrowed sets Narrowed field to given value.


### GetOrganizationTools

`func (o *OrganizationMcpServerWorkspace) GetOrganizationTools() []string`

GetOrganizationTools returns the OrganizationTools field if non-nil, zero value otherwise.

### GetOrganizationToolsOk

`func (o *OrganizationMcpServerWorkspace) GetOrganizationToolsOk() (*[]string, bool)`

GetOrganizationToolsOk returns a tuple with the OrganizationTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationTools

`func (o *OrganizationMcpServerWorkspace) SetOrganizationTools(v []string)`

SetOrganizationTools sets OrganizationTools field to given value.


### SetOrganizationToolsNil

`func (o *OrganizationMcpServerWorkspace) SetOrganizationToolsNil(b bool)`

 SetOrganizationToolsNil sets the value for OrganizationTools to be an explicit nil

### UnsetOrganizationTools
`func (o *OrganizationMcpServerWorkspace) UnsetOrganizationTools()`

UnsetOrganizationTools ensures that no value is present for OrganizationTools, not even an explicit nil
### GetReached

`func (o *OrganizationMcpServerWorkspace) GetReached() bool`

GetReached returns the Reached field if non-nil, zero value otherwise.

### GetReachedOk

`func (o *OrganizationMcpServerWorkspace) GetReachedOk() (*bool, bool)`

GetReachedOk returns a tuple with the Reached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReached

`func (o *OrganizationMcpServerWorkspace) SetReached(v bool)`

SetReached sets Reached field to given value.


### GetWorkspaceEnabled

`func (o *OrganizationMcpServerWorkspace) GetWorkspaceEnabled() bool`

GetWorkspaceEnabled returns the WorkspaceEnabled field if non-nil, zero value otherwise.

### GetWorkspaceEnabledOk

`func (o *OrganizationMcpServerWorkspace) GetWorkspaceEnabledOk() (*bool, bool)`

GetWorkspaceEnabledOk returns a tuple with the WorkspaceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceEnabled

`func (o *OrganizationMcpServerWorkspace) SetWorkspaceEnabled(v bool)`

SetWorkspaceEnabled sets WorkspaceEnabled field to given value.


### GetWorkspaceId

`func (o *OrganizationMcpServerWorkspace) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OrganizationMcpServerWorkspace) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OrganizationMcpServerWorkspace) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


