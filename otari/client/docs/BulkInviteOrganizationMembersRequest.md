# BulkInviteOrganizationMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** |  | 
**Role** | Pointer to **string** |  | [optional] [default to "member"]
**WorkspaceAssignments** | Pointer to [**[]WorkspaceAssignmentRequest**](WorkspaceAssignmentRequest.md) |  | [optional] 

## Methods

### NewBulkInviteOrganizationMembersRequest

`func NewBulkInviteOrganizationMembersRequest(emails []string, ) *BulkInviteOrganizationMembersRequest`

NewBulkInviteOrganizationMembersRequest instantiates a new BulkInviteOrganizationMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteOrganizationMembersRequestWithDefaults

`func NewBulkInviteOrganizationMembersRequestWithDefaults() *BulkInviteOrganizationMembersRequest`

NewBulkInviteOrganizationMembersRequestWithDefaults instantiates a new BulkInviteOrganizationMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *BulkInviteOrganizationMembersRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *BulkInviteOrganizationMembersRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *BulkInviteOrganizationMembersRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetRole

`func (o *BulkInviteOrganizationMembersRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkInviteOrganizationMembersRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkInviteOrganizationMembersRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkInviteOrganizationMembersRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetWorkspaceAssignments

`func (o *BulkInviteOrganizationMembersRequest) GetWorkspaceAssignments() []WorkspaceAssignmentRequest`

GetWorkspaceAssignments returns the WorkspaceAssignments field if non-nil, zero value otherwise.

### GetWorkspaceAssignmentsOk

`func (o *BulkInviteOrganizationMembersRequest) GetWorkspaceAssignmentsOk() (*[]WorkspaceAssignmentRequest, bool)`

GetWorkspaceAssignmentsOk returns a tuple with the WorkspaceAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAssignments

`func (o *BulkInviteOrganizationMembersRequest) SetWorkspaceAssignments(v []WorkspaceAssignmentRequest)`

SetWorkspaceAssignments sets WorkspaceAssignments field to given value.

### HasWorkspaceAssignments

`func (o *BulkInviteOrganizationMembersRequest) HasWorkspaceAssignments() bool`

HasWorkspaceAssignments returns a boolean if a field has been set.

### SetWorkspaceAssignmentsNil

`func (o *BulkInviteOrganizationMembersRequest) SetWorkspaceAssignmentsNil(b bool)`

 SetWorkspaceAssignmentsNil sets the value for WorkspaceAssignments to be an explicit nil

### UnsetWorkspaceAssignments
`func (o *BulkInviteOrganizationMembersRequest) UnsetWorkspaceAssignments()`

UnsetWorkspaceAssignments ensures that no value is present for WorkspaceAssignments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


