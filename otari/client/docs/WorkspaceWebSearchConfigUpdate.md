# WorkspaceWebSearchConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | Pointer to **[]string** | Results are kept only from these domains; intersected with any list the request sends | [optional] 
**BlockedDomains** | Pointer to **[]string** | Results from these domains are dropped; added to any list the request sends | [optional] 
**Enabled** | **bool** | False refuses web search for this workspace, both the otari_web_search tool and POST /v1/search. The fields below narrow the tool only. | 
**MaxResults** | Pointer to **NullableInt32** | Ceiling on results one search returns; only ever lowers the effective limit, so at most 20 | [optional] 
**ProviderOptions** | Pointer to **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | [optional] 
**PurposeHint** | Pointer to **NullableString** | Hint used when a request declares otari_web_search without one of its own | [optional] 

## Methods

### NewWorkspaceWebSearchConfigUpdate

`func NewWorkspaceWebSearchConfigUpdate(enabled bool, ) *WorkspaceWebSearchConfigUpdate`

NewWorkspaceWebSearchConfigUpdate instantiates a new WorkspaceWebSearchConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWebSearchConfigUpdateWithDefaults

`func NewWorkspaceWebSearchConfigUpdateWithDefaults() *WorkspaceWebSearchConfigUpdate`

NewWorkspaceWebSearchConfigUpdateWithDefaults instantiates a new WorkspaceWebSearchConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *WorkspaceWebSearchConfigUpdate) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *WorkspaceWebSearchConfigUpdate) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *WorkspaceWebSearchConfigUpdate) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *WorkspaceWebSearchConfigUpdate) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### SetAllowedDomainsNil

`func (o *WorkspaceWebSearchConfigUpdate) SetAllowedDomainsNil(b bool)`

 SetAllowedDomainsNil sets the value for AllowedDomains to be an explicit nil

### UnsetAllowedDomains
`func (o *WorkspaceWebSearchConfigUpdate) UnsetAllowedDomains()`

UnsetAllowedDomains ensures that no value is present for AllowedDomains, not even an explicit nil
### GetBlockedDomains

`func (o *WorkspaceWebSearchConfigUpdate) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *WorkspaceWebSearchConfigUpdate) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *WorkspaceWebSearchConfigUpdate) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.

### HasBlockedDomains

`func (o *WorkspaceWebSearchConfigUpdate) HasBlockedDomains() bool`

HasBlockedDomains returns a boolean if a field has been set.

### SetBlockedDomainsNil

`func (o *WorkspaceWebSearchConfigUpdate) SetBlockedDomainsNil(b bool)`

 SetBlockedDomainsNil sets the value for BlockedDomains to be an explicit nil

### UnsetBlockedDomains
`func (o *WorkspaceWebSearchConfigUpdate) UnsetBlockedDomains()`

UnsetBlockedDomains ensures that no value is present for BlockedDomains, not even an explicit nil
### GetEnabled

`func (o *WorkspaceWebSearchConfigUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkspaceWebSearchConfigUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkspaceWebSearchConfigUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxResults

`func (o *WorkspaceWebSearchConfigUpdate) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *WorkspaceWebSearchConfigUpdate) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *WorkspaceWebSearchConfigUpdate) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *WorkspaceWebSearchConfigUpdate) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### SetMaxResultsNil

`func (o *WorkspaceWebSearchConfigUpdate) SetMaxResultsNil(b bool)`

 SetMaxResultsNil sets the value for MaxResults to be an explicit nil

### UnsetMaxResults
`func (o *WorkspaceWebSearchConfigUpdate) UnsetMaxResults()`

UnsetMaxResults ensures that no value is present for MaxResults, not even an explicit nil
### GetProviderOptions

`func (o *WorkspaceWebSearchConfigUpdate) GetProviderOptions() map[string]interface{}`

GetProviderOptions returns the ProviderOptions field if non-nil, zero value otherwise.

### GetProviderOptionsOk

`func (o *WorkspaceWebSearchConfigUpdate) GetProviderOptionsOk() (*map[string]interface{}, bool)`

GetProviderOptionsOk returns a tuple with the ProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOptions

`func (o *WorkspaceWebSearchConfigUpdate) SetProviderOptions(v map[string]interface{})`

SetProviderOptions sets ProviderOptions field to given value.

### HasProviderOptions

`func (o *WorkspaceWebSearchConfigUpdate) HasProviderOptions() bool`

HasProviderOptions returns a boolean if a field has been set.

### SetProviderOptionsNil

`func (o *WorkspaceWebSearchConfigUpdate) SetProviderOptionsNil(b bool)`

 SetProviderOptionsNil sets the value for ProviderOptions to be an explicit nil

### UnsetProviderOptions
`func (o *WorkspaceWebSearchConfigUpdate) UnsetProviderOptions()`

UnsetProviderOptions ensures that no value is present for ProviderOptions, not even an explicit nil
### GetPurposeHint

`func (o *WorkspaceWebSearchConfigUpdate) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *WorkspaceWebSearchConfigUpdate) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *WorkspaceWebSearchConfigUpdate) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.

### HasPurposeHint

`func (o *WorkspaceWebSearchConfigUpdate) HasPurposeHint() bool`

HasPurposeHint returns a boolean if a field has been set.

### SetPurposeHintNil

`func (o *WorkspaceWebSearchConfigUpdate) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *WorkspaceWebSearchConfigUpdate) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


