# BulkInviteOrganizationMembersResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | [**[]BulkInvitationFailurePublic**](BulkInvitationFailurePublic.md) |  | 
**Invited** | [**[]InviteOrganizationMemberResultPublic**](InviteOrganizationMemberResultPublic.md) |  | 

## Methods

### NewBulkInviteOrganizationMembersResultPublic

`func NewBulkInviteOrganizationMembersResultPublic(failed []BulkInvitationFailurePublic, invited []InviteOrganizationMemberResultPublic, ) *BulkInviteOrganizationMembersResultPublic`

NewBulkInviteOrganizationMembersResultPublic instantiates a new BulkInviteOrganizationMembersResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkInviteOrganizationMembersResultPublicWithDefaults

`func NewBulkInviteOrganizationMembersResultPublicWithDefaults() *BulkInviteOrganizationMembersResultPublic`

NewBulkInviteOrganizationMembersResultPublicWithDefaults instantiates a new BulkInviteOrganizationMembersResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *BulkInviteOrganizationMembersResultPublic) GetFailed() []BulkInvitationFailurePublic`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkInviteOrganizationMembersResultPublic) GetFailedOk() (*[]BulkInvitationFailurePublic, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkInviteOrganizationMembersResultPublic) SetFailed(v []BulkInvitationFailurePublic)`

SetFailed sets Failed field to given value.


### GetInvited

`func (o *BulkInviteOrganizationMembersResultPublic) GetInvited() []InviteOrganizationMemberResultPublic`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *BulkInviteOrganizationMembersResultPublic) GetInvitedOk() (*[]InviteOrganizationMemberResultPublic, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *BulkInviteOrganizationMembersResultPublic) SetInvited(v []InviteOrganizationMemberResultPublic)`

SetInvited sets Invited field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


