# OrganizationSsoPreflightPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedCount** | **int32** | How many active members hold no sign-in link to this connection. | 
**AffectedMembers** | [**[]OrganizationSsoPreflightMemberPublic**](OrganizationSsoPreflightMemberPublic.md) | A sample of the affected members, for recognizing who they are. | 
**AffectedMembersTruncated** | **bool** | Whether affected_members leaves out some of the members affected_count counts. | 
**BreakGlassCandidates** | [**[]OrganizationSsoPreflightMemberPublic**](OrganizationSsoPreflightMemberPublic.md) | The owners who may be named to keep a route in that does not depend on this connection. An owner this connection provisioned is not among them, because exempting an account that cannot sign in without the connection keeps no route open. | 
**TotalMembers** | **int32** | The organization&#39;s active membership count. | 

## Methods

### NewOrganizationSsoPreflightPublic

`func NewOrganizationSsoPreflightPublic(affectedCount int32, affectedMembers []OrganizationSsoPreflightMemberPublic, affectedMembersTruncated bool, breakGlassCandidates []OrganizationSsoPreflightMemberPublic, totalMembers int32, ) *OrganizationSsoPreflightPublic`

NewOrganizationSsoPreflightPublic instantiates a new OrganizationSsoPreflightPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSsoPreflightPublicWithDefaults

`func NewOrganizationSsoPreflightPublicWithDefaults() *OrganizationSsoPreflightPublic`

NewOrganizationSsoPreflightPublicWithDefaults instantiates a new OrganizationSsoPreflightPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedCount

`func (o *OrganizationSsoPreflightPublic) GetAffectedCount() int32`

GetAffectedCount returns the AffectedCount field if non-nil, zero value otherwise.

### GetAffectedCountOk

`func (o *OrganizationSsoPreflightPublic) GetAffectedCountOk() (*int32, bool)`

GetAffectedCountOk returns a tuple with the AffectedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedCount

`func (o *OrganizationSsoPreflightPublic) SetAffectedCount(v int32)`

SetAffectedCount sets AffectedCount field to given value.


### GetAffectedMembers

`func (o *OrganizationSsoPreflightPublic) GetAffectedMembers() []OrganizationSsoPreflightMemberPublic`

GetAffectedMembers returns the AffectedMembers field if non-nil, zero value otherwise.

### GetAffectedMembersOk

`func (o *OrganizationSsoPreflightPublic) GetAffectedMembersOk() (*[]OrganizationSsoPreflightMemberPublic, bool)`

GetAffectedMembersOk returns a tuple with the AffectedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMembers

`func (o *OrganizationSsoPreflightPublic) SetAffectedMembers(v []OrganizationSsoPreflightMemberPublic)`

SetAffectedMembers sets AffectedMembers field to given value.


### GetAffectedMembersTruncated

`func (o *OrganizationSsoPreflightPublic) GetAffectedMembersTruncated() bool`

GetAffectedMembersTruncated returns the AffectedMembersTruncated field if non-nil, zero value otherwise.

### GetAffectedMembersTruncatedOk

`func (o *OrganizationSsoPreflightPublic) GetAffectedMembersTruncatedOk() (*bool, bool)`

GetAffectedMembersTruncatedOk returns a tuple with the AffectedMembersTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMembersTruncated

`func (o *OrganizationSsoPreflightPublic) SetAffectedMembersTruncated(v bool)`

SetAffectedMembersTruncated sets AffectedMembersTruncated field to given value.


### GetBreakGlassCandidates

`func (o *OrganizationSsoPreflightPublic) GetBreakGlassCandidates() []OrganizationSsoPreflightMemberPublic`

GetBreakGlassCandidates returns the BreakGlassCandidates field if non-nil, zero value otherwise.

### GetBreakGlassCandidatesOk

`func (o *OrganizationSsoPreflightPublic) GetBreakGlassCandidatesOk() (*[]OrganizationSsoPreflightMemberPublic, bool)`

GetBreakGlassCandidatesOk returns a tuple with the BreakGlassCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassCandidates

`func (o *OrganizationSsoPreflightPublic) SetBreakGlassCandidates(v []OrganizationSsoPreflightMemberPublic)`

SetBreakGlassCandidates sets BreakGlassCandidates field to given value.


### GetTotalMembers

`func (o *OrganizationSsoPreflightPublic) GetTotalMembers() int32`

GetTotalMembers returns the TotalMembers field if non-nil, zero value otherwise.

### GetTotalMembersOk

`func (o *OrganizationSsoPreflightPublic) GetTotalMembersOk() (*int32, bool)`

GetTotalMembersOk returns a tuple with the TotalMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMembers

`func (o *OrganizationSsoPreflightPublic) SetTotalMembers(v int32)`

SetTotalMembers sets TotalMembers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


