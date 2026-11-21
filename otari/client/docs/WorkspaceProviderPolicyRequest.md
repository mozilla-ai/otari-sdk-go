# WorkspaceProviderPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDataLocations** | Pointer to **[]string** | The data locations a hosted provider region must declare to serve the workspace: ISO 3166-1 alpha-2 country codes, EU (which also admits each member state) or GLOBAL. Null accepts any location, and omitting the field keeps the stored list. A region whose location nobody declared never matches a list. | [optional] 
**RequireZeroDataRetention** | Pointer to **NullableBool** | Refuse every provider not declared zero data retention: an organization key or a hosted provider nobody declared, and every provider configured in config.yml or endpoint the workspace owns. Omitted keeps the stored setting. | [optional] 

## Methods

### NewWorkspaceProviderPolicyRequest

`func NewWorkspaceProviderPolicyRequest() *WorkspaceProviderPolicyRequest`

NewWorkspaceProviderPolicyRequest instantiates a new WorkspaceProviderPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceProviderPolicyRequestWithDefaults

`func NewWorkspaceProviderPolicyRequestWithDefaults() *WorkspaceProviderPolicyRequest`

NewWorkspaceProviderPolicyRequestWithDefaults instantiates a new WorkspaceProviderPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDataLocations

`func (o *WorkspaceProviderPolicyRequest) GetAllowedDataLocations() []string`

GetAllowedDataLocations returns the AllowedDataLocations field if non-nil, zero value otherwise.

### GetAllowedDataLocationsOk

`func (o *WorkspaceProviderPolicyRequest) GetAllowedDataLocationsOk() (*[]string, bool)`

GetAllowedDataLocationsOk returns a tuple with the AllowedDataLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDataLocations

`func (o *WorkspaceProviderPolicyRequest) SetAllowedDataLocations(v []string)`

SetAllowedDataLocations sets AllowedDataLocations field to given value.

### HasAllowedDataLocations

`func (o *WorkspaceProviderPolicyRequest) HasAllowedDataLocations() bool`

HasAllowedDataLocations returns a boolean if a field has been set.

### SetAllowedDataLocationsNil

`func (o *WorkspaceProviderPolicyRequest) SetAllowedDataLocationsNil(b bool)`

 SetAllowedDataLocationsNil sets the value for AllowedDataLocations to be an explicit nil

### UnsetAllowedDataLocations
`func (o *WorkspaceProviderPolicyRequest) UnsetAllowedDataLocations()`

UnsetAllowedDataLocations ensures that no value is present for AllowedDataLocations, not even an explicit nil
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

### HasRequireZeroDataRetention

`func (o *WorkspaceProviderPolicyRequest) HasRequireZeroDataRetention() bool`

HasRequireZeroDataRetention returns a boolean if a field has been set.

### SetRequireZeroDataRetentionNil

`func (o *WorkspaceProviderPolicyRequest) SetRequireZeroDataRetentionNil(b bool)`

 SetRequireZeroDataRetentionNil sets the value for RequireZeroDataRetention to be an explicit nil

### UnsetRequireZeroDataRetention
`func (o *WorkspaceProviderPolicyRequest) UnsetRequireZeroDataRetention()`

UnsetRequireZeroDataRetention ensures that no value is present for RequireZeroDataRetention, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


