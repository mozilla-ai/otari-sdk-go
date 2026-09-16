# InviteOrganizationMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "member"]
**WorkspaceAssignments** | Pointer to [**[]WorkspaceAssignmentRequest**](WorkspaceAssignmentRequest.md) |  | [optional] 

## Methods

### NewInviteOrganizationMemberRequest

`func NewInviteOrganizationMemberRequest(email string, ) *InviteOrganizationMemberRequest`

NewInviteOrganizationMemberRequest instantiates a new InviteOrganizationMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteOrganizationMemberRequestWithDefaults

`func NewInviteOrganizationMemberRequestWithDefaults() *InviteOrganizationMemberRequest`

NewInviteOrganizationMemberRequestWithDefaults instantiates a new InviteOrganizationMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteOrganizationMemberRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteOrganizationMemberRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteOrganizationMemberRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *InviteOrganizationMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteOrganizationMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteOrganizationMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *InviteOrganizationMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetWorkspaceAssignments

`func (o *InviteOrganizationMemberRequest) GetWorkspaceAssignments() []WorkspaceAssignmentRequest`

GetWorkspaceAssignments returns the WorkspaceAssignments field if non-nil, zero value otherwise.

### GetWorkspaceAssignmentsOk

`func (o *InviteOrganizationMemberRequest) GetWorkspaceAssignmentsOk() (*[]WorkspaceAssignmentRequest, bool)`

GetWorkspaceAssignmentsOk returns a tuple with the WorkspaceAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAssignments

`func (o *InviteOrganizationMemberRequest) SetWorkspaceAssignments(v []WorkspaceAssignmentRequest)`

SetWorkspaceAssignments sets WorkspaceAssignments field to given value.

### HasWorkspaceAssignments

`func (o *InviteOrganizationMemberRequest) HasWorkspaceAssignments() bool`

HasWorkspaceAssignments returns a boolean if a field has been set.

### SetWorkspaceAssignmentsNil

`func (o *InviteOrganizationMemberRequest) SetWorkspaceAssignmentsNil(b bool)`

 SetWorkspaceAssignmentsNil sets the value for WorkspaceAssignments to be an explicit nil

### UnsetWorkspaceAssignments
`func (o *InviteOrganizationMemberRequest) UnsetWorkspaceAssignments()`

UnsetWorkspaceAssignments ensures that no value is present for WorkspaceAssignments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


