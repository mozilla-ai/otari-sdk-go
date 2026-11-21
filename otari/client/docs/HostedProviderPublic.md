# HostedProviderPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKeyLast4** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**DataLocation** | Pointer to **NullableString** |  | [optional] 
**Decryptable** | Pointer to **bool** |  | [optional] [default to true]
**Enabled** | **bool** |  | 
**HasApiKey** | Pointer to **bool** |  | [optional] [default to true]
**Id** | Pointer to **NullableString** |  | [optional] 
**Models** | Pointer to [**NullableHostedProviderModelSummary**](HostedProviderModelSummary.md) |  | [optional] 
**Name** | **string** |  | 
**PrivacyUrl** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** |  | 
**Regions** | Pointer to [**[]HostedProviderRegionPublic**](HostedProviderRegionPublic.md) |  | [optional] 
**RestrictedToOrganizationId** | Pointer to **NullableString** |  | [optional] 
**ServesUnlistedModels** | Pointer to **bool** |  | [optional] [default to false]
**ShadowsConfig** | Pointer to **bool** |  | [optional] [default to false]
**Source** | Pointer to **string** |  | [optional] [default to "database"]
**SupportedCountries** | Pointer to **[]string** |  | [optional] 
**TermsUrl** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ZeroDataRetention** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewHostedProviderPublic

`func NewHostedProviderPublic(enabled bool, name string, provider string, ) *HostedProviderPublic`

NewHostedProviderPublic instantiates a new HostedProviderPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedProviderPublicWithDefaults

`func NewHostedProviderPublicWithDefaults() *HostedProviderPublic`

NewHostedProviderPublicWithDefaults instantiates a new HostedProviderPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *HostedProviderPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *HostedProviderPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *HostedProviderPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *HostedProviderPublic) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *HostedProviderPublic) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *HostedProviderPublic) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKeyLast4

`func (o *HostedProviderPublic) GetApiKeyLast4() string`

GetApiKeyLast4 returns the ApiKeyLast4 field if non-nil, zero value otherwise.

### GetApiKeyLast4Ok

`func (o *HostedProviderPublic) GetApiKeyLast4Ok() (*string, bool)`

GetApiKeyLast4Ok returns a tuple with the ApiKeyLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyLast4

`func (o *HostedProviderPublic) SetApiKeyLast4(v string)`

SetApiKeyLast4 sets ApiKeyLast4 field to given value.

### HasApiKeyLast4

`func (o *HostedProviderPublic) HasApiKeyLast4() bool`

HasApiKeyLast4 returns a boolean if a field has been set.

### SetApiKeyLast4Nil

`func (o *HostedProviderPublic) SetApiKeyLast4Nil(b bool)`

 SetApiKeyLast4Nil sets the value for ApiKeyLast4 to be an explicit nil

### UnsetApiKeyLast4
`func (o *HostedProviderPublic) UnsetApiKeyLast4()`

UnsetApiKeyLast4 ensures that no value is present for ApiKeyLast4, not even an explicit nil
### GetClientArgs

`func (o *HostedProviderPublic) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *HostedProviderPublic) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *HostedProviderPublic) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *HostedProviderPublic) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *HostedProviderPublic) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *HostedProviderPublic) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetCreatedAt

`func (o *HostedProviderPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostedProviderPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostedProviderPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HostedProviderPublic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HostedProviderPublic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HostedProviderPublic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDataLocation

`func (o *HostedProviderPublic) GetDataLocation() string`

GetDataLocation returns the DataLocation field if non-nil, zero value otherwise.

### GetDataLocationOk

`func (o *HostedProviderPublic) GetDataLocationOk() (*string, bool)`

GetDataLocationOk returns a tuple with the DataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLocation

`func (o *HostedProviderPublic) SetDataLocation(v string)`

SetDataLocation sets DataLocation field to given value.

### HasDataLocation

`func (o *HostedProviderPublic) HasDataLocation() bool`

HasDataLocation returns a boolean if a field has been set.

### SetDataLocationNil

`func (o *HostedProviderPublic) SetDataLocationNil(b bool)`

 SetDataLocationNil sets the value for DataLocation to be an explicit nil

### UnsetDataLocation
`func (o *HostedProviderPublic) UnsetDataLocation()`

UnsetDataLocation ensures that no value is present for DataLocation, not even an explicit nil
### GetDecryptable

`func (o *HostedProviderPublic) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *HostedProviderPublic) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *HostedProviderPublic) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *HostedProviderPublic) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetEnabled

`func (o *HostedProviderPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedProviderPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedProviderPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasApiKey

`func (o *HostedProviderPublic) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *HostedProviderPublic) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *HostedProviderPublic) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.

### HasHasApiKey

`func (o *HostedProviderPublic) HasHasApiKey() bool`

HasHasApiKey returns a boolean if a field has been set.

### GetId

`func (o *HostedProviderPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostedProviderPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostedProviderPublic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostedProviderPublic) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HostedProviderPublic) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HostedProviderPublic) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModels

`func (o *HostedProviderPublic) GetModels() HostedProviderModelSummary`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *HostedProviderPublic) GetModelsOk() (*HostedProviderModelSummary, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *HostedProviderPublic) SetModels(v HostedProviderModelSummary)`

SetModels sets Models field to given value.

### HasModels

`func (o *HostedProviderPublic) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *HostedProviderPublic) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *HostedProviderPublic) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetName

`func (o *HostedProviderPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostedProviderPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostedProviderPublic) SetName(v string)`

SetName sets Name field to given value.


### GetPrivacyUrl

`func (o *HostedProviderPublic) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *HostedProviderPublic) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *HostedProviderPublic) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *HostedProviderPublic) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *HostedProviderPublic) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *HostedProviderPublic) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetProvider

`func (o *HostedProviderPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HostedProviderPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HostedProviderPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegions

`func (o *HostedProviderPublic) GetRegions() []HostedProviderRegionPublic`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *HostedProviderPublic) GetRegionsOk() (*[]HostedProviderRegionPublic, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *HostedProviderPublic) SetRegions(v []HostedProviderRegionPublic)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *HostedProviderPublic) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetRestrictedToOrganizationId

`func (o *HostedProviderPublic) GetRestrictedToOrganizationId() string`

GetRestrictedToOrganizationId returns the RestrictedToOrganizationId field if non-nil, zero value otherwise.

### GetRestrictedToOrganizationIdOk

`func (o *HostedProviderPublic) GetRestrictedToOrganizationIdOk() (*string, bool)`

GetRestrictedToOrganizationIdOk returns a tuple with the RestrictedToOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToOrganizationId

`func (o *HostedProviderPublic) SetRestrictedToOrganizationId(v string)`

SetRestrictedToOrganizationId sets RestrictedToOrganizationId field to given value.

### HasRestrictedToOrganizationId

`func (o *HostedProviderPublic) HasRestrictedToOrganizationId() bool`

HasRestrictedToOrganizationId returns a boolean if a field has been set.

### SetRestrictedToOrganizationIdNil

`func (o *HostedProviderPublic) SetRestrictedToOrganizationIdNil(b bool)`

 SetRestrictedToOrganizationIdNil sets the value for RestrictedToOrganizationId to be an explicit nil

### UnsetRestrictedToOrganizationId
`func (o *HostedProviderPublic) UnsetRestrictedToOrganizationId()`

UnsetRestrictedToOrganizationId ensures that no value is present for RestrictedToOrganizationId, not even an explicit nil
### GetServesUnlistedModels

`func (o *HostedProviderPublic) GetServesUnlistedModels() bool`

GetServesUnlistedModels returns the ServesUnlistedModels field if non-nil, zero value otherwise.

### GetServesUnlistedModelsOk

`func (o *HostedProviderPublic) GetServesUnlistedModelsOk() (*bool, bool)`

GetServesUnlistedModelsOk returns a tuple with the ServesUnlistedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServesUnlistedModels

`func (o *HostedProviderPublic) SetServesUnlistedModels(v bool)`

SetServesUnlistedModels sets ServesUnlistedModels field to given value.

### HasServesUnlistedModels

`func (o *HostedProviderPublic) HasServesUnlistedModels() bool`

HasServesUnlistedModels returns a boolean if a field has been set.

### GetShadowsConfig

`func (o *HostedProviderPublic) GetShadowsConfig() bool`

GetShadowsConfig returns the ShadowsConfig field if non-nil, zero value otherwise.

### GetShadowsConfigOk

`func (o *HostedProviderPublic) GetShadowsConfigOk() (*bool, bool)`

GetShadowsConfigOk returns a tuple with the ShadowsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowsConfig

`func (o *HostedProviderPublic) SetShadowsConfig(v bool)`

SetShadowsConfig sets ShadowsConfig field to given value.

### HasShadowsConfig

`func (o *HostedProviderPublic) HasShadowsConfig() bool`

HasShadowsConfig returns a boolean if a field has been set.

### GetSource

`func (o *HostedProviderPublic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HostedProviderPublic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HostedProviderPublic) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HostedProviderPublic) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSupportedCountries

`func (o *HostedProviderPublic) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *HostedProviderPublic) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *HostedProviderPublic) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *HostedProviderPublic) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.

### SetSupportedCountriesNil

`func (o *HostedProviderPublic) SetSupportedCountriesNil(b bool)`

 SetSupportedCountriesNil sets the value for SupportedCountries to be an explicit nil

### UnsetSupportedCountries
`func (o *HostedProviderPublic) UnsetSupportedCountries()`

UnsetSupportedCountries ensures that no value is present for SupportedCountries, not even an explicit nil
### GetTermsUrl

`func (o *HostedProviderPublic) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *HostedProviderPublic) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *HostedProviderPublic) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *HostedProviderPublic) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *HostedProviderPublic) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *HostedProviderPublic) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil
### GetUpdatedAt

`func (o *HostedProviderPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostedProviderPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostedProviderPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HostedProviderPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HostedProviderPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HostedProviderPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetZeroDataRetention

`func (o *HostedProviderPublic) GetZeroDataRetention() bool`

GetZeroDataRetention returns the ZeroDataRetention field if non-nil, zero value otherwise.

### GetZeroDataRetentionOk

`func (o *HostedProviderPublic) GetZeroDataRetentionOk() (*bool, bool)`

GetZeroDataRetentionOk returns a tuple with the ZeroDataRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroDataRetention

`func (o *HostedProviderPublic) SetZeroDataRetention(v bool)`

SetZeroDataRetention sets ZeroDataRetention field to given value.

### HasZeroDataRetention

`func (o *HostedProviderPublic) HasZeroDataRetention() bool`

HasZeroDataRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


