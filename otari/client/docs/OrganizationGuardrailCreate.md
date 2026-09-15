# OrganizationGuardrailCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesToAllWorkspaces** | Pointer to **bool** | True runs this in every workspace of the organization, including one created later; false runs it only in the workspaces named by workspace_ids | [optional] [default to false]
**Credential** | Pointer to **NullableString** | Bearer credential for this entry&#39;s endpoint. Requires url to be set, and https: an entry with no endpoint of its own falls back to the deployment&#39;s guardrails_url, which is commonly a same-host http sidecar. Encrypted at rest, never returned | [optional] 
**Enabled** | Pointer to **bool** | False stops the guardrail everywhere without discarding it | [optional] [default to true]
**Mode** | Pointer to **string** | block rejects a flagged request with 403; monitor annotates the response and forwards it | [optional] [default to "monitor"]
**OnUnavailable** | Pointer to **string** | What a block-mode entry does when the guardrails service cannot be reached at all | [optional] [default to "block"]
**Profile** | **string** | Profile name configured on the guardrails service, unique within the organization | 
**Url** | Pointer to **NullableString** | Guardrails endpoint for this entry; null uses the deployment&#39;s guardrails_url | [optional] 
**ValidateKwargs** | Pointer to **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | [optional] 
**WorkspaceIds** | Pointer to **[]string** | Workspaces this guardrail runs in. Must be empty when applies_to_all_workspaces is true | [optional] 

## Methods

### NewOrganizationGuardrailCreate

`func NewOrganizationGuardrailCreate(profile string, ) *OrganizationGuardrailCreate`

NewOrganizationGuardrailCreate instantiates a new OrganizationGuardrailCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailCreateWithDefaults

`func NewOrganizationGuardrailCreateWithDefaults() *OrganizationGuardrailCreate`

NewOrganizationGuardrailCreateWithDefaults instantiates a new OrganizationGuardrailCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailCreate) GetAppliesToAllWorkspaces() bool`

GetAppliesToAllWorkspaces returns the AppliesToAllWorkspaces field if non-nil, zero value otherwise.

### GetAppliesToAllWorkspacesOk

`func (o *OrganizationGuardrailCreate) GetAppliesToAllWorkspacesOk() (*bool, bool)`

GetAppliesToAllWorkspacesOk returns a tuple with the AppliesToAllWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailCreate) SetAppliesToAllWorkspaces(v bool)`

SetAppliesToAllWorkspaces sets AppliesToAllWorkspaces field to given value.

### HasAppliesToAllWorkspaces

`func (o *OrganizationGuardrailCreate) HasAppliesToAllWorkspaces() bool`

HasAppliesToAllWorkspaces returns a boolean if a field has been set.

### GetCredential

`func (o *OrganizationGuardrailCreate) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OrganizationGuardrailCreate) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OrganizationGuardrailCreate) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OrganizationGuardrailCreate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### SetCredentialNil

`func (o *OrganizationGuardrailCreate) SetCredentialNil(b bool)`

 SetCredentialNil sets the value for Credential to be an explicit nil

### UnsetCredential
`func (o *OrganizationGuardrailCreate) UnsetCredential()`

UnsetCredential ensures that no value is present for Credential, not even an explicit nil
### GetEnabled

`func (o *OrganizationGuardrailCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationGuardrailCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationGuardrailCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationGuardrailCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *OrganizationGuardrailCreate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrganizationGuardrailCreate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrganizationGuardrailCreate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OrganizationGuardrailCreate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOnUnavailable

`func (o *OrganizationGuardrailCreate) GetOnUnavailable() string`

GetOnUnavailable returns the OnUnavailable field if non-nil, zero value otherwise.

### GetOnUnavailableOk

`func (o *OrganizationGuardrailCreate) GetOnUnavailableOk() (*string, bool)`

GetOnUnavailableOk returns a tuple with the OnUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUnavailable

`func (o *OrganizationGuardrailCreate) SetOnUnavailable(v string)`

SetOnUnavailable sets OnUnavailable field to given value.

### HasOnUnavailable

`func (o *OrganizationGuardrailCreate) HasOnUnavailable() bool`

HasOnUnavailable returns a boolean if a field has been set.

### GetProfile

`func (o *OrganizationGuardrailCreate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OrganizationGuardrailCreate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OrganizationGuardrailCreate) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetUrl

`func (o *OrganizationGuardrailCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationGuardrailCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationGuardrailCreate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationGuardrailCreate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *OrganizationGuardrailCreate) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *OrganizationGuardrailCreate) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidateKwargs

`func (o *OrganizationGuardrailCreate) GetValidateKwargs() map[string]interface{}`

GetValidateKwargs returns the ValidateKwargs field if non-nil, zero value otherwise.

### GetValidateKwargsOk

`func (o *OrganizationGuardrailCreate) GetValidateKwargsOk() (*map[string]interface{}, bool)`

GetValidateKwargsOk returns a tuple with the ValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateKwargs

`func (o *OrganizationGuardrailCreate) SetValidateKwargs(v map[string]interface{})`

SetValidateKwargs sets ValidateKwargs field to given value.

### HasValidateKwargs

`func (o *OrganizationGuardrailCreate) HasValidateKwargs() bool`

HasValidateKwargs returns a boolean if a field has been set.

### SetValidateKwargsNil

`func (o *OrganizationGuardrailCreate) SetValidateKwargsNil(b bool)`

 SetValidateKwargsNil sets the value for ValidateKwargs to be an explicit nil

### UnsetValidateKwargs
`func (o *OrganizationGuardrailCreate) UnsetValidateKwargs()`

UnsetValidateKwargs ensures that no value is present for ValidateKwargs, not even an explicit nil
### GetWorkspaceIds

`func (o *OrganizationGuardrailCreate) GetWorkspaceIds() []string`

GetWorkspaceIds returns the WorkspaceIds field if non-nil, zero value otherwise.

### GetWorkspaceIdsOk

`func (o *OrganizationGuardrailCreate) GetWorkspaceIdsOk() (*[]string, bool)`

GetWorkspaceIdsOk returns a tuple with the WorkspaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIds

`func (o *OrganizationGuardrailCreate) SetWorkspaceIds(v []string)`

SetWorkspaceIds sets WorkspaceIds field to given value.

### HasWorkspaceIds

`func (o *OrganizationGuardrailCreate) HasWorkspaceIds() bool`

HasWorkspaceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


