# BuiltInGuardrailSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateBackends** | Pointer to **[]string** |  | [optional] 
**Backend** | [**BackendType**](BackendType.md) |  | 
**Categories** | **[]string** |  | 
**CreateParameters** | Pointer to [**[]GuardrailParameterSpec**](GuardrailParameterSpec.md) | Constructor arguments, which is where a vendor API key and an endpoint live | [optional] 
**DefaultLicense** | **string** |  | 
**Description** | **string** |  | 
**DisplayName** | **string** |  | 
**GuardrailName** | **string** | The any-guardrail class, and the name a stored guardrail selects | 
**Multilingual** | Pointer to **bool** |  | [optional] [default to false]
**Multimodal** | Pointer to **bool** |  | [optional] [default to false]
**OptionalValidateKwargs** | Pointer to **[]string** |  | [optional] 
**OutputShapes** | **[]string** |  | 
**PrimaryCategory** | [**GuardrailCategory**](GuardrailCategory.md) |  | 
**RequiredValidateKwargs** | Pointer to **[]string** |  | [optional] 
**RequirementGroups** | Pointer to [**[]RequirementGroup**](RequirementGroup.md) | One-of constraints that no single parameter&#39;s required flag can express. At least one member of each group must be supplied, or one of the environment variables that satisfies it | [optional] 
**RequiresApiKey** | Pointer to **bool** |  | [optional] [default to false]
**Stages** | **[]string** |  | 
**SupportsBatch** | Pointer to **bool** | Whether several inputs run as one real batched call, not a per-item loop | [optional] [default to false]
**ValidateParameters** | Pointer to [**[]GuardrailParameterSpec**](GuardrailParameterSpec.md) | Per-call arguments, sent with the text on every check | [optional] 
**VariantLicenses** | Pointer to **[]map[string]string** |  | [optional] 
**Vendor** | **string** |  | 

## Methods

### NewBuiltInGuardrailSpec

`func NewBuiltInGuardrailSpec(backend BackendType, categories []string, defaultLicense string, description string, displayName string, guardrailName string, outputShapes []string, primaryCategory GuardrailCategory, stages []string, vendor string, ) *BuiltInGuardrailSpec`

NewBuiltInGuardrailSpec instantiates a new BuiltInGuardrailSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltInGuardrailSpecWithDefaults

`func NewBuiltInGuardrailSpecWithDefaults() *BuiltInGuardrailSpec`

NewBuiltInGuardrailSpecWithDefaults instantiates a new BuiltInGuardrailSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateBackends

`func (o *BuiltInGuardrailSpec) GetAlternateBackends() []string`

GetAlternateBackends returns the AlternateBackends field if non-nil, zero value otherwise.

### GetAlternateBackendsOk

`func (o *BuiltInGuardrailSpec) GetAlternateBackendsOk() (*[]string, bool)`

GetAlternateBackendsOk returns a tuple with the AlternateBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateBackends

`func (o *BuiltInGuardrailSpec) SetAlternateBackends(v []string)`

SetAlternateBackends sets AlternateBackends field to given value.

### HasAlternateBackends

`func (o *BuiltInGuardrailSpec) HasAlternateBackends() bool`

HasAlternateBackends returns a boolean if a field has been set.

### GetBackend

`func (o *BuiltInGuardrailSpec) GetBackend() BackendType`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *BuiltInGuardrailSpec) GetBackendOk() (*BackendType, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *BuiltInGuardrailSpec) SetBackend(v BackendType)`

SetBackend sets Backend field to given value.


### GetCategories

`func (o *BuiltInGuardrailSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BuiltInGuardrailSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BuiltInGuardrailSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetCreateParameters

`func (o *BuiltInGuardrailSpec) GetCreateParameters() []GuardrailParameterSpec`

GetCreateParameters returns the CreateParameters field if non-nil, zero value otherwise.

### GetCreateParametersOk

`func (o *BuiltInGuardrailSpec) GetCreateParametersOk() (*[]GuardrailParameterSpec, bool)`

GetCreateParametersOk returns a tuple with the CreateParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateParameters

`func (o *BuiltInGuardrailSpec) SetCreateParameters(v []GuardrailParameterSpec)`

SetCreateParameters sets CreateParameters field to given value.

### HasCreateParameters

`func (o *BuiltInGuardrailSpec) HasCreateParameters() bool`

HasCreateParameters returns a boolean if a field has been set.

### GetDefaultLicense

`func (o *BuiltInGuardrailSpec) GetDefaultLicense() string`

GetDefaultLicense returns the DefaultLicense field if non-nil, zero value otherwise.

### GetDefaultLicenseOk

`func (o *BuiltInGuardrailSpec) GetDefaultLicenseOk() (*string, bool)`

GetDefaultLicenseOk returns a tuple with the DefaultLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicense

`func (o *BuiltInGuardrailSpec) SetDefaultLicense(v string)`

SetDefaultLicense sets DefaultLicense field to given value.


### GetDescription

`func (o *BuiltInGuardrailSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuiltInGuardrailSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuiltInGuardrailSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *BuiltInGuardrailSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuiltInGuardrailSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuiltInGuardrailSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetGuardrailName

`func (o *BuiltInGuardrailSpec) GetGuardrailName() string`

GetGuardrailName returns the GuardrailName field if non-nil, zero value otherwise.

### GetGuardrailNameOk

`func (o *BuiltInGuardrailSpec) GetGuardrailNameOk() (*string, bool)`

GetGuardrailNameOk returns a tuple with the GuardrailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailName

`func (o *BuiltInGuardrailSpec) SetGuardrailName(v string)`

SetGuardrailName sets GuardrailName field to given value.


### GetMultilingual

`func (o *BuiltInGuardrailSpec) GetMultilingual() bool`

GetMultilingual returns the Multilingual field if non-nil, zero value otherwise.

### GetMultilingualOk

`func (o *BuiltInGuardrailSpec) GetMultilingualOk() (*bool, bool)`

GetMultilingualOk returns a tuple with the Multilingual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultilingual

`func (o *BuiltInGuardrailSpec) SetMultilingual(v bool)`

SetMultilingual sets Multilingual field to given value.

### HasMultilingual

`func (o *BuiltInGuardrailSpec) HasMultilingual() bool`

HasMultilingual returns a boolean if a field has been set.

### GetMultimodal

`func (o *BuiltInGuardrailSpec) GetMultimodal() bool`

GetMultimodal returns the Multimodal field if non-nil, zero value otherwise.

### GetMultimodalOk

`func (o *BuiltInGuardrailSpec) GetMultimodalOk() (*bool, bool)`

GetMultimodalOk returns a tuple with the Multimodal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimodal

`func (o *BuiltInGuardrailSpec) SetMultimodal(v bool)`

SetMultimodal sets Multimodal field to given value.

### HasMultimodal

`func (o *BuiltInGuardrailSpec) HasMultimodal() bool`

HasMultimodal returns a boolean if a field has been set.

### GetOptionalValidateKwargs

`func (o *BuiltInGuardrailSpec) GetOptionalValidateKwargs() []string`

GetOptionalValidateKwargs returns the OptionalValidateKwargs field if non-nil, zero value otherwise.

### GetOptionalValidateKwargsOk

`func (o *BuiltInGuardrailSpec) GetOptionalValidateKwargsOk() (*[]string, bool)`

GetOptionalValidateKwargsOk returns a tuple with the OptionalValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalValidateKwargs

`func (o *BuiltInGuardrailSpec) SetOptionalValidateKwargs(v []string)`

SetOptionalValidateKwargs sets OptionalValidateKwargs field to given value.

### HasOptionalValidateKwargs

`func (o *BuiltInGuardrailSpec) HasOptionalValidateKwargs() bool`

HasOptionalValidateKwargs returns a boolean if a field has been set.

### GetOutputShapes

`func (o *BuiltInGuardrailSpec) GetOutputShapes() []string`

GetOutputShapes returns the OutputShapes field if non-nil, zero value otherwise.

### GetOutputShapesOk

`func (o *BuiltInGuardrailSpec) GetOutputShapesOk() (*[]string, bool)`

GetOutputShapesOk returns a tuple with the OutputShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputShapes

`func (o *BuiltInGuardrailSpec) SetOutputShapes(v []string)`

SetOutputShapes sets OutputShapes field to given value.


### GetPrimaryCategory

`func (o *BuiltInGuardrailSpec) GetPrimaryCategory() GuardrailCategory`

GetPrimaryCategory returns the PrimaryCategory field if non-nil, zero value otherwise.

### GetPrimaryCategoryOk

`func (o *BuiltInGuardrailSpec) GetPrimaryCategoryOk() (*GuardrailCategory, bool)`

GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCategory

`func (o *BuiltInGuardrailSpec) SetPrimaryCategory(v GuardrailCategory)`

SetPrimaryCategory sets PrimaryCategory field to given value.


### GetRequiredValidateKwargs

`func (o *BuiltInGuardrailSpec) GetRequiredValidateKwargs() []string`

GetRequiredValidateKwargs returns the RequiredValidateKwargs field if non-nil, zero value otherwise.

### GetRequiredValidateKwargsOk

`func (o *BuiltInGuardrailSpec) GetRequiredValidateKwargsOk() (*[]string, bool)`

GetRequiredValidateKwargsOk returns a tuple with the RequiredValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredValidateKwargs

`func (o *BuiltInGuardrailSpec) SetRequiredValidateKwargs(v []string)`

SetRequiredValidateKwargs sets RequiredValidateKwargs field to given value.

### HasRequiredValidateKwargs

`func (o *BuiltInGuardrailSpec) HasRequiredValidateKwargs() bool`

HasRequiredValidateKwargs returns a boolean if a field has been set.

### GetRequirementGroups

`func (o *BuiltInGuardrailSpec) GetRequirementGroups() []RequirementGroup`

GetRequirementGroups returns the RequirementGroups field if non-nil, zero value otherwise.

### GetRequirementGroupsOk

`func (o *BuiltInGuardrailSpec) GetRequirementGroupsOk() (*[]RequirementGroup, bool)`

GetRequirementGroupsOk returns a tuple with the RequirementGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementGroups

`func (o *BuiltInGuardrailSpec) SetRequirementGroups(v []RequirementGroup)`

SetRequirementGroups sets RequirementGroups field to given value.

### HasRequirementGroups

`func (o *BuiltInGuardrailSpec) HasRequirementGroups() bool`

HasRequirementGroups returns a boolean if a field has been set.

### GetRequiresApiKey

`func (o *BuiltInGuardrailSpec) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *BuiltInGuardrailSpec) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *BuiltInGuardrailSpec) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.

### HasRequiresApiKey

`func (o *BuiltInGuardrailSpec) HasRequiresApiKey() bool`

HasRequiresApiKey returns a boolean if a field has been set.

### GetStages

`func (o *BuiltInGuardrailSpec) GetStages() []string`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *BuiltInGuardrailSpec) GetStagesOk() (*[]string, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *BuiltInGuardrailSpec) SetStages(v []string)`

SetStages sets Stages field to given value.


### GetSupportsBatch

`func (o *BuiltInGuardrailSpec) GetSupportsBatch() bool`

GetSupportsBatch returns the SupportsBatch field if non-nil, zero value otherwise.

### GetSupportsBatchOk

`func (o *BuiltInGuardrailSpec) GetSupportsBatchOk() (*bool, bool)`

GetSupportsBatchOk returns a tuple with the SupportsBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsBatch

`func (o *BuiltInGuardrailSpec) SetSupportsBatch(v bool)`

SetSupportsBatch sets SupportsBatch field to given value.

### HasSupportsBatch

`func (o *BuiltInGuardrailSpec) HasSupportsBatch() bool`

HasSupportsBatch returns a boolean if a field has been set.

### GetValidateParameters

`func (o *BuiltInGuardrailSpec) GetValidateParameters() []GuardrailParameterSpec`

GetValidateParameters returns the ValidateParameters field if non-nil, zero value otherwise.

### GetValidateParametersOk

`func (o *BuiltInGuardrailSpec) GetValidateParametersOk() (*[]GuardrailParameterSpec, bool)`

GetValidateParametersOk returns a tuple with the ValidateParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateParameters

`func (o *BuiltInGuardrailSpec) SetValidateParameters(v []GuardrailParameterSpec)`

SetValidateParameters sets ValidateParameters field to given value.

### HasValidateParameters

`func (o *BuiltInGuardrailSpec) HasValidateParameters() bool`

HasValidateParameters returns a boolean if a field has been set.

### GetVariantLicenses

`func (o *BuiltInGuardrailSpec) GetVariantLicenses() []map[string]string`

GetVariantLicenses returns the VariantLicenses field if non-nil, zero value otherwise.

### GetVariantLicensesOk

`func (o *BuiltInGuardrailSpec) GetVariantLicensesOk() (*[]map[string]string, bool)`

GetVariantLicensesOk returns a tuple with the VariantLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLicenses

`func (o *BuiltInGuardrailSpec) SetVariantLicenses(v []map[string]string)`

SetVariantLicenses sets VariantLicenses field to given value.

### HasVariantLicenses

`func (o *BuiltInGuardrailSpec) HasVariantLicenses() bool`

HasVariantLicenses returns a boolean if a field has been set.

### GetVendor

`func (o *BuiltInGuardrailSpec) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BuiltInGuardrailSpec) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BuiltInGuardrailSpec) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


