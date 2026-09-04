# WorkspaceWebSearchConfigPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | **[]string** |  | 
**BlockedDomains** | **[]string** |  | 
**Configured** | **bool** |  | 
**CreatedAt** | **NullableString** |  | 
**Enabled** | **bool** |  | 
**MaxResults** | **NullableInt32** |  | 
**ProviderOptions** | **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | 
**PurposeHint** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**WebSearchConfigured** | **bool** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceWebSearchConfigPublic

`func NewWorkspaceWebSearchConfigPublic(allowedDomains []string, blockedDomains []string, configured bool, createdAt NullableString, enabled bool, maxResults NullableInt32, providerOptions map[string]interface{}, purposeHint NullableString, updatedAt NullableString, webSearchConfigured bool, workspaceId string, ) *WorkspaceWebSearchConfigPublic`

NewWorkspaceWebSearchConfigPublic instantiates a new WorkspaceWebSearchConfigPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWebSearchConfigPublicWithDefaults

`func NewWorkspaceWebSearchConfigPublicWithDefaults() *WorkspaceWebSearchConfigPublic`

NewWorkspaceWebSearchConfigPublicWithDefaults instantiates a new WorkspaceWebSearchConfigPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *WorkspaceWebSearchConfigPublic) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *WorkspaceWebSearchConfigPublic) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *WorkspaceWebSearchConfigPublic) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.


### SetAllowedDomainsNil

`func (o *WorkspaceWebSearchConfigPublic) SetAllowedDomainsNil(b bool)`

 SetAllowedDomainsNil sets the value for AllowedDomains to be an explicit nil

### UnsetAllowedDomains
`func (o *WorkspaceWebSearchConfigPublic) UnsetAllowedDomains()`

UnsetAllowedDomains ensures that no value is present for AllowedDomains, not even an explicit nil
### GetBlockedDomains

`func (o *WorkspaceWebSearchConfigPublic) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *WorkspaceWebSearchConfigPublic) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *WorkspaceWebSearchConfigPublic) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.


### SetBlockedDomainsNil

`func (o *WorkspaceWebSearchConfigPublic) SetBlockedDomainsNil(b bool)`

 SetBlockedDomainsNil sets the value for BlockedDomains to be an explicit nil

### UnsetBlockedDomains
`func (o *WorkspaceWebSearchConfigPublic) UnsetBlockedDomains()`

UnsetBlockedDomains ensures that no value is present for BlockedDomains, not even an explicit nil
### GetConfigured

`func (o *WorkspaceWebSearchConfigPublic) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *WorkspaceWebSearchConfigPublic) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *WorkspaceWebSearchConfigPublic) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.


### GetCreatedAt

`func (o *WorkspaceWebSearchConfigPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceWebSearchConfigPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceWebSearchConfigPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *WorkspaceWebSearchConfigPublic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WorkspaceWebSearchConfigPublic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEnabled

`func (o *WorkspaceWebSearchConfigPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceWebSearchConfigPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceWebSearchConfigPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxResults

`func (o *WorkspaceWebSearchConfigPublic) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *WorkspaceWebSearchConfigPublic) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *WorkspaceWebSearchConfigPublic) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.


### SetMaxResultsNil

`func (o *WorkspaceWebSearchConfigPublic) SetMaxResultsNil(b bool)`

 SetMaxResultsNil sets the value for MaxResults to be an explicit nil

### UnsetMaxResults
`func (o *WorkspaceWebSearchConfigPublic) UnsetMaxResults()`

UnsetMaxResults ensures that no value is present for MaxResults, not even an explicit nil
### GetProviderOptions

`func (o *WorkspaceWebSearchConfigPublic) GetProviderOptions() map[string]interface{}`

GetProviderOptions returns the ProviderOptions field if non-nil, zero value otherwise.

### GetProviderOptionsOk

`func (o *WorkspaceWebSearchConfigPublic) GetProviderOptionsOk() (*map[string]interface{}, bool)`

GetProviderOptionsOk returns a tuple with the ProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOptions

`func (o *WorkspaceWebSearchConfigPublic) SetProviderOptions(v map[string]interface{})`

SetProviderOptions sets ProviderOptions field to given value.


### SetProviderOptionsNil

`func (o *WorkspaceWebSearchConfigPublic) SetProviderOptionsNil(b bool)`

 SetProviderOptionsNil sets the value for ProviderOptions to be an explicit nil

### UnsetProviderOptions
`func (o *WorkspaceWebSearchConfigPublic) UnsetProviderOptions()`

UnsetProviderOptions ensures that no value is present for ProviderOptions, not even an explicit nil
### GetPurposeHint

`func (o *WorkspaceWebSearchConfigPublic) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceWebSearchConfigPublic) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceWebSearchConfigPublic) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.


### SetPurposeHintNil

`func (o *WorkspaceWebSearchConfigPublic) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceWebSearchConfigPublic) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetUpdatedAt

`func (o *WorkspaceWebSearchConfigPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceWebSearchConfigPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceWebSearchConfigPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *WorkspaceWebSearchConfigPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *WorkspaceWebSearchConfigPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWebSearchConfigured

`func (o *WorkspaceWebSearchConfigPublic) GetWebSearchConfigured() bool`

GetWebSearchConfigured returns the WebSearchConfigured field if non-nil, zero value otherwise.

### GetWebSearchConfiguredOk

`func (o *WorkspaceWebSearchConfigPublic) GetWebSearchConfiguredOk() (*bool, bool)`

GetWebSearchConfiguredOk returns a tuple with the WebSearchConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchConfigured

`func (o *WorkspaceWebSearchConfigPublic) SetWebSearchConfigured(v bool)`

SetWebSearchConfigured sets WebSearchConfigured field to given value.


### GetWorkspaceId

`func (o *WorkspaceWebSearchConfigPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceWebSearchConfigPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceWebSearchConfigPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


