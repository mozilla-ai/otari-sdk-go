# WorkspaceAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] [default to "member"]
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceAssignmentRequest

`func NewWorkspaceAssignmentRequest(workspaceId string, ) *WorkspaceAssignmentRequest`

NewWorkspaceAssignmentRequest instantiates a new WorkspaceAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceAssignmentRequestWithDefaults

`func NewWorkspaceAssignmentRequestWithDefaults() *WorkspaceAssignmentRequest`

NewWorkspaceAssignmentRequestWithDefaults instantiates a new WorkspaceAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *WorkspaceAssignmentRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WorkspaceAssignmentRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WorkspaceAssignmentRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *WorkspaceAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WorkspaceAssignmentRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceAssignmentRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceAssignmentRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


