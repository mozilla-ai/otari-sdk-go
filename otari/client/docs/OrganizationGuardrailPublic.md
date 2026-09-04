# OrganizationGuardrailPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesToAllWorkspaces** | **bool** |  | 
**CreatedAt** | **string** |  | 
**Enabled** | **bool** |  | 
**HasCredential** | **bool** |  | 
**Id** | **string** |  | 
**Mode** | **string** |  | 
**OnUnavailable** | **string** |  | 
**OrganizationId** | **string** |  | 
**Profile** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Url** | **NullableString** |  | 
**ValidateKwargs** | **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | 
**WorkspaceIds** | **[]string** |  | 

## Methods

### NewOrganizationGuardrailPublic

`func NewOrganizationGuardrailPublic(appliesToAllWorkspaces bool, createdAt string, enabled bool, hasCredential bool, id string, mode string, onUnavailable string, organizationId string, profile string, updatedAt string, url NullableString, validateKwargs map[string]interface{}, workspaceIds []string, ) *OrganizationGuardrailPublic`

NewOrganizationGuardrailPublic instantiates a new OrganizationGuardrailPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailPublicWithDefaults

`func NewOrganizationGuardrailPublicWithDefaults() *OrganizationGuardrailPublic`

NewOrganizationGuardrailPublicWithDefaults instantiates a new OrganizationGuardrailPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailPublic) GetAppliesToAllWorkspaces() bool`

GetAppliesToAllWorkspaces returns the AppliesToAllWorkspaces field if non-nil, zero value otherwise.

### GetAppliesToAllWorkspacesOk

`func (o *OrganizationGuardrailPublic) GetAppliesToAllWorkspacesOk() (*bool, bool)`

GetAppliesToAllWorkspacesOk returns a tuple with the AppliesToAllWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToAllWorkspaces

`func (o *OrganizationGuardrailPublic) SetAppliesToAllWorkspaces(v bool)`

SetAppliesToAllWorkspaces sets AppliesToAllWorkspaces field to given value.


### GetCreatedAt

`func (o *OrganizationGuardrailPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationGuardrailPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationGuardrailPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *OrganizationGuardrailPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationGuardrailPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationGuardrailPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasCredential

`func (o *OrganizationGuardrailPublic) GetHasCredential() bool`

GetHasCredential returns the HasCredential field if non-nil, zero value otherwise.

### GetHasCredentialOk

`func (o *OrganizationGuardrailPublic) GetHasCredentialOk() (*bool, bool)`

GetHasCredentialOk returns a tuple with the HasCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCredential

`func (o *OrganizationGuardrailPublic) SetHasCredential(v bool)`

SetHasCredential sets HasCredential field to given value.


### GetId

`func (o *OrganizationGuardrailPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationGuardrailPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationGuardrailPublic) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *OrganizationGuardrailPublic) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrganizationGuardrailPublic) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrganizationGuardrailPublic) SetMode(v string)`

SetMode sets Mode field to given value.


### GetOnUnavailable

`func (o *OrganizationGuardrailPublic) GetOnUnavailable() string`

GetOnUnavailable returns the OnUnavailable field if non-nil, zero value otherwise.

### GetOnUnavailableOk

`func (o *OrganizationGuardrailPublic) GetOnUnavailableOk() (*string, bool)`

GetOnUnavailableOk returns a tuple with the OnUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUnavailable

`func (o *OrganizationGuardrailPublic) SetOnUnavailable(v string)`

SetOnUnavailable sets OnUnavailable field to given value.


### GetOrganizationId

`func (o *OrganizationGuardrailPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationGuardrailPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationGuardrailPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetProfile

`func (o *OrganizationGuardrailPublic) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OrganizationGuardrailPublic) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OrganizationGuardrailPublic) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetUpdatedAt

`func (o *OrganizationGuardrailPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationGuardrailPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationGuardrailPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *OrganizationGuardrailPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationGuardrailPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationGuardrailPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *OrganizationGuardrailPublic) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *OrganizationGuardrailPublic) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidateKwargs

`func (o *OrganizationGuardrailPublic) GetValidateKwargs() map[string]interface{}`

GetValidateKwargs returns the ValidateKwargs field if non-nil, zero value otherwise.

### GetValidateKwargsOk

`func (o *OrganizationGuardrailPublic) GetValidateKwargsOk() (*map[string]interface{}, bool)`

GetValidateKwargsOk returns a tuple with the ValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateKwargs

`func (o *OrganizationGuardrailPublic) SetValidateKwargs(v map[string]interface{})`

SetValidateKwargs sets ValidateKwargs field to given value.


### SetValidateKwargsNil

`func (o *OrganizationGuardrailPublic) SetValidateKwargsNil(b bool)`

 SetValidateKwargsNil sets the value for ValidateKwargs to be an explicit nil

### UnsetValidateKwargs
`func (o *OrganizationGuardrailPublic) UnsetValidateKwargs()`

UnsetValidateKwargs ensures that no value is present for ValidateKwargs, not even an explicit nil
### GetWorkspaceIds

`func (o *OrganizationGuardrailPublic) GetWorkspaceIds() []string`

GetWorkspaceIds returns the WorkspaceIds field if non-nil, zero value otherwise.

### GetWorkspaceIdsOk

`func (o *OrganizationGuardrailPublic) GetWorkspaceIdsOk() (*[]string, bool)`

GetWorkspaceIdsOk returns a tuple with the WorkspaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIds

`func (o *OrganizationGuardrailPublic) SetWorkspaceIds(v []string)`

SetWorkspaceIds sets WorkspaceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


