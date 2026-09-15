# PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Model name callers send, e.g. &#39;fast&#39;. | 
**RenameFrom** | Pointer to **NullableString** | Current name of the policy to rename, in the same scope, meaning the same workspace and the same user. A rename never moves a policy between them. The stored row keeps its id and created_at and takes &#x60;name&#x60; and &#x60;spec&#x60;. Sending it asserts that policy exists, so a name with no stored row is a 404 rather than a create, even when it equals &#x60;name&#x60;. Omit to create or update the policy named &#x60;name&#x60;. Renaming changes what callers must send as &#x60;model&#x60;; usage already recorded keeps the old name. | [optional] 
**Spec** | **map[string]interface{}** | The policy body: select (with exactly one &#x60;default&#x60; entry, last), optional on_failure and guardrails. Same schema as a &#x60;routing.policies&#x60; entry in config.yml, and closed to unknown keys, so a typo is a 400 rather than a silently ignored setting. | 
**UserId** | Pointer to **NullableString** | User this policy belongs to. Omit for a policy every caller in the workspace sees. A user-scoped policy resolves only for that user and shadows the workspace-wide one of the same name. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | Workspace this policy belongs to. Omit for the deployment&#39;s default workspace. The policy resolves only for requests in that workspace, so two workspaces can each define their own &#39;fast&#39;. | [optional] 

## Methods

### NewPolicyRequest

`func NewPolicyRequest(name string, spec map[string]interface{}, ) *PolicyRequest`

NewPolicyRequest instantiates a new PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRequestWithDefaults

`func NewPolicyRequestWithDefaults() *PolicyRequest`

NewPolicyRequestWithDefaults instantiates a new PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRenameFrom

`func (o *PolicyRequest) GetRenameFrom() string`

GetRenameFrom returns the RenameFrom field if non-nil, zero value otherwise.

### GetRenameFromOk

`func (o *PolicyRequest) GetRenameFromOk() (*string, bool)`

GetRenameFromOk returns a tuple with the RenameFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameFrom

`func (o *PolicyRequest) SetRenameFrom(v string)`

SetRenameFrom sets RenameFrom field to given value.

### HasRenameFrom

`func (o *PolicyRequest) HasRenameFrom() bool`

HasRenameFrom returns a boolean if a field has been set.

### SetRenameFromNil

`func (o *PolicyRequest) SetRenameFromNil(b bool)`

 SetRenameFromNil sets the value for RenameFrom to be an explicit nil

### UnsetRenameFrom
`func (o *PolicyRequest) UnsetRenameFrom()`

UnsetRenameFrom ensures that no value is present for RenameFrom, not even an explicit nil
### GetSpec

`func (o *PolicyRequest) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PolicyRequest) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PolicyRequest) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetUserId

`func (o *PolicyRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PolicyRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PolicyRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PolicyRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *PolicyRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PolicyRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *PolicyRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PolicyRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PolicyRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *PolicyRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *PolicyRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *PolicyRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


