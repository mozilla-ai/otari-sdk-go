# WorkspaceProviderPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireZeroDataRetention** | **bool** | Refuse every provider not declared zero data retention: an organization key or a hosted provider nobody declared, and every provider configured in config.yml or endpoint the workspace owns. | 

## Methods

### NewWorkspaceProviderPolicyRequest

`func NewWorkspaceProviderPolicyRequest(requireZeroDataRetention bool, ) *WorkspaceProviderPolicyRequest`

NewWorkspaceProviderPolicyRequest instantiates a new WorkspaceProviderPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceProviderPolicyRequestWithDefaults

`func NewWorkspaceProviderPolicyRequestWithDefaults() *WorkspaceProviderPolicyRequest`

NewWorkspaceProviderPolicyRequestWithDefaults instantiates a new WorkspaceProviderPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireZeroDataRetention

`func (o *WorkspaceProviderPolicyRequest) GetRequireZeroDataRetention() bool`

GetRequireZeroDataRetention returns the RequireZeroDataRetention field if non-nil, zero value otherwise.

### GetRequireZeroDataRetentionOk

`func (o *WorkspaceProviderPolicyRequest) GetRequireZeroDataRetentionOk() (*bool, bool)`

GetRequireZeroDataRetentionOk returns a tuple with the RequireZeroDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireZeroDataRetention

`func (o *WorkspaceProviderPolicyRequest) SetRequireZeroDataRetention(v bool)`

SetRequireZeroDataRetention sets RequireZeroDataRetention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


