# WorkspaceProviderPolicyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDataLocations** | Pointer to **[]string** |  | [optional] 
**RequireZeroDataRetention** | **bool** |  | 

## Methods

### NewWorkspaceProviderPolicyPublic

`func NewWorkspaceProviderPolicyPublic(requireZeroDataRetention bool, ) *WorkspaceProviderPolicyPublic`

NewWorkspaceProviderPolicyPublic instantiates a new WorkspaceProviderPolicyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceProviderPolicyPublicWithDefaults

`func NewWorkspaceProviderPolicyPublicWithDefaults() *WorkspaceProviderPolicyPublic`

NewWorkspaceProviderPolicyPublicWithDefaults instantiates a new WorkspaceProviderPolicyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDataLocations

`func (o *WorkspaceProviderPolicyPublic) GetAllowedDataLocations() []string`

GetAllowedDataLocations returns the AllowedDataLocations field if non-nil, zero value otherwise.

### GetAllowedDataLocationsOk

`func (o *WorkspaceProviderPolicyPublic) GetAllowedDataLocationsOk() (*[]string, bool)`

GetAllowedDataLocationsOk returns a tuple with the AllowedDataLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataLocations

`func (o *WorkspaceProviderPolicyPublic) SetAllowedDataLocations(v []string)`

SetAllowedDataLocations sets AllowedDataLocations field to given value.

### HasAllowedDataLocations

`func (o *WorkspaceProviderPolicyPublic) HasAllowedDataLocations() bool`

HasAllowedDataLocations returns a boolean if a field has been set.

### SetAllowedDataLocationsNil

`func (o *WorkspaceProviderPolicyPublic) SetAllowedDataLocationsNil(b bool)`

 SetAllowedDataLocationsNil sets the value for AllowedDataLocations to be an explicit nil

### UnsetAllowedDataLocations
`func (o *WorkspaceProviderPolicyPublic) UnsetAllowedDataLocations()`

UnsetAllowedDataLocations ensures that no value is present for AllowedDataLocations, not even an explicit nil
### GetRequireZeroDataRetention

`func (o *WorkspaceProviderPolicyPublic) GetRequireZeroDataRetention() bool`

GetRequireZeroDataRetention returns the RequireZeroDataRetention field if non-nil, zero value otherwise.

### GetRequireZeroDataRetentionOk

`func (o *WorkspaceProviderPolicyPublic) GetRequireZeroDataRetentionOk() (*bool, bool)`

GetRequireZeroDataRetentionOk returns a tuple with the RequireZeroDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireZeroDataRetention

`func (o *WorkspaceProviderPolicyPublic) SetRequireZeroDataRetention(v bool)`

SetRequireZeroDataRetention sets RequireZeroDataRetention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


