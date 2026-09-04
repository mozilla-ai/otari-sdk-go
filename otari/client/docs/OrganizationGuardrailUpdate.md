# OrganizationGuardrailUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesToAllWorkspaces** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**OnUnavailable** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ValidateKwargs** | Pointer to **map[string]interface{}** |  | [optional] 
**WorkspaceIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrganizationGuardrailUpdate

`func NewOrganizationGuardrailUpdate() *OrganizationGuardrailUpdate`

NewOrganizationGuardrailUpdate instantiates a new OrganizationGuardrailUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailUpdateWithDefaults

`func NewOrganizationGuardrailUpdateWithDefaults() *OrganizationGuardrailUpdate`

NewOrganizationGuardrailUpdateWithDefaults instantiates a new OrganizationGuardrailUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailUpdate) GetAppliesToAllWorkspaces() bool`

GetAppliesToAllWorkspaces returns the AppliesToAllWorkspaces field if non-nil, zero value otherwise.

### GetAppliesToAllWorkspacesOk

`func (o *OrganizationGuardrailUpdate) GetAppliesToAllWorkspacesOk() (*bool, bool)`

GetAppliesToAllWorkspacesOk returns a tuple with the AppliesToAllWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailUpdate) SetAppliesToAllWorkspaces(v bool)`

SetAppliesToAllWorkspaces sets AppliesToAllWorkspaces field to given value.

### HasAppliesToAllWorkspaces

`func (o *OrganizationGuardrailUpdate) HasAppliesToAllWorkspaces() bool`

HasAppliesToAllWorkspaces returns a boolean if a field has been set.

### GetCredential

`func (o *OrganizationGuardrailUpdate) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OrganizationGuardrailUpdate) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OrganizationGuardrailUpdate) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OrganizationGuardrailUpdate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### SetCredentialNil

`func (o *OrganizationGuardrailUpdate) SetCredentialNil(b bool)`

 SetCredentialNil sets the value for Credential to be an explicit nil

### UnsetCredential
`func (o *OrganizationGuardrailUpdate) UnsetCredential()`

UnsetCredential ensures that no value is present for Credential, not even an explicit nil
### GetEnabled

`func (o *OrganizationGuardrailUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationGuardrailUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationGuardrailUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationGuardrailUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *OrganizationGuardrailUpdate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrganizationGuardrailUpdate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrganizationGuardrailUpdate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OrganizationGuardrailUpdate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOnUnavailable

`func (o *OrganizationGuardrailUpdate) GetOnUnavailable() string`

GetOnUnavailable returns the OnUnavailable field if non-nil, zero value otherwise.

### GetOnUnavailableOk

`func (o *OrganizationGuardrailUpdate) GetOnUnavailableOk() (*string, bool)`

GetOnUnavailableOk returns a tuple with the OnUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUnavailable

`func (o *OrganizationGuardrailUpdate) SetOnUnavailable(v string)`

SetOnUnavailable sets OnUnavailable field to given value.

### HasOnUnavailable

`func (o *OrganizationGuardrailUpdate) HasOnUnavailable() bool`

HasOnUnavailable returns a boolean if a field has been set.

### GetProfile

`func (o *OrganizationGuardrailUpdate) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OrganizationGuardrailUpdate) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OrganizationGuardrailUpdate) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *OrganizationGuardrailUpdate) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetUrl

`func (o *OrganizationGuardrailUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationGuardrailUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationGuardrailUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationGuardrailUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *OrganizationGuardrailUpdate) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *OrganizationGuardrailUpdate) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidateKwargs

`func (o *OrganizationGuardrailUpdate) GetValidateKwargs() map[string]interface{}`

GetValidateKwargs returns the ValidateKwargs field if non-nil, zero value otherwise.

### GetValidateKwargsOk

`func (o *OrganizationGuardrailUpdate) GetValidateKwargsOk() (*map[string]interface{}, bool)`

GetValidateKwargsOk returns a tuple with the ValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateKwargs

`func (o *OrganizationGuardrailUpdate) SetValidateKwargs(v map[string]interface{})`

SetValidateKwargs sets ValidateKwargs field to given value.

### HasValidateKwargs

`func (o *OrganizationGuardrailUpdate) HasValidateKwargs() bool`

HasValidateKwargs returns a boolean if a field has been set.

### SetValidateKwargsNil

`func (o *OrganizationGuardrailUpdate) SetValidateKwargsNil(b bool)`

 SetValidateKwargsNil sets the value for ValidateKwargs to be an explicit nil

### UnsetValidateKwargs
`func (o *OrganizationGuardrailUpdate) UnsetValidateKwargs()`

UnsetValidateKwargs ensures that no value is present for ValidateKwargs, not even an explicit nil
### GetWorkspaceIds

`func (o *OrganizationGuardrailUpdate) GetWorkspaceIds() []string`

GetWorkspaceIds returns the WorkspaceIds field if non-nil, zero value otherwise.

### GetWorkspaceIdsOk

`func (o *OrganizationGuardrailUpdate) GetWorkspaceIdsOk() (*[]string, bool)`

GetWorkspaceIdsOk returns a tuple with the WorkspaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIds

`func (o *OrganizationGuardrailUpdate) SetWorkspaceIds(v []string)`

SetWorkspaceIds sets WorkspaceIds field to given value.

### HasWorkspaceIds

`func (o *OrganizationGuardrailUpdate) HasWorkspaceIds() bool`

HasWorkspaceIds returns a boolean if a field has been set.

### SetWorkspaceIdsNil

`func (o *OrganizationGuardrailUpdate) SetWorkspaceIdsNil(b bool)`

 SetWorkspaceIdsNil sets the value for WorkspaceIds to be an explicit nil

### UnsetWorkspaceIds
`func (o *OrganizationGuardrailUpdate) UnsetWorkspaceIds()`

UnsetWorkspaceIds ensures that no value is present for WorkspaceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


