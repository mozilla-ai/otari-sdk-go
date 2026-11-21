# HostedProviderRegionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ClientArgs** | Pointer to **map[string]interface{}** |  | [optional] 
**DataLocation** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Label** | **string** |  | 
**PrimaryLabel** | Pointer to **NullableString** |  | [optional] 
**PrimarySlug** | Pointer to **NullableString** |  | [optional] 
**Slug** | **string** |  | 
**SupportedCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHostedProviderRegionCreateRequest

`func NewHostedProviderRegionCreateRequest(label string, slug string, ) *HostedProviderRegionCreateRequest`

NewHostedProviderRegionCreateRequest instantiates a new HostedProviderRegionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedProviderRegionCreateRequestWithDefaults

`func NewHostedProviderRegionCreateRequestWithDefaults() *HostedProviderRegionCreateRequest`

NewHostedProviderRegionCreateRequestWithDefaults instantiates a new HostedProviderRegionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *HostedProviderRegionCreateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *HostedProviderRegionCreateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *HostedProviderRegionCreateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *HostedProviderRegionCreateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *HostedProviderRegionCreateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *HostedProviderRegionCreateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *HostedProviderRegionCreateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HostedProviderRegionCreateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HostedProviderRegionCreateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HostedProviderRegionCreateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *HostedProviderRegionCreateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HostedProviderRegionCreateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetClientArgs

`func (o *HostedProviderRegionCreateRequest) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *HostedProviderRegionCreateRequest) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *HostedProviderRegionCreateRequest) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *HostedProviderRegionCreateRequest) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *HostedProviderRegionCreateRequest) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *HostedProviderRegionCreateRequest) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetDataLocation

`func (o *HostedProviderRegionCreateRequest) GetDataLocation() string`

GetDataLocation returns the DataLocation field if non-nil, zero value otherwise.

### GetDataLocationOk

`func (o *HostedProviderRegionCreateRequest) GetDataLocationOk() (*string, bool)`

GetDataLocationOk returns a tuple with the DataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLocation

`func (o *HostedProviderRegionCreateRequest) SetDataLocation(v string)`

SetDataLocation sets DataLocation field to given value.

### HasDataLocation

`func (o *HostedProviderRegionCreateRequest) HasDataLocation() bool`

HasDataLocation returns a boolean if a field has been set.

### SetDataLocationNil

`func (o *HostedProviderRegionCreateRequest) SetDataLocationNil(b bool)`

 SetDataLocationNil sets the value for DataLocation to be an explicit nil

### UnsetDataLocation
`func (o *HostedProviderRegionCreateRequest) UnsetDataLocation()`

UnsetDataLocation ensures that no value is present for DataLocation, not even an explicit nil
### GetEnabled

`func (o *HostedProviderRegionCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedProviderRegionCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedProviderRegionCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HostedProviderRegionCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLabel

`func (o *HostedProviderRegionCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HostedProviderRegionCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HostedProviderRegionCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPrimaryLabel

`func (o *HostedProviderRegionCreateRequest) GetPrimaryLabel() string`

GetPrimaryLabel returns the PrimaryLabel field if non-nil, zero value otherwise.

### GetPrimaryLabelOk

`func (o *HostedProviderRegionCreateRequest) GetPrimaryLabelOk() (*string, bool)`

GetPrimaryLabelOk returns a tuple with the PrimaryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLabel

`func (o *HostedProviderRegionCreateRequest) SetPrimaryLabel(v string)`

SetPrimaryLabel sets PrimaryLabel field to given value.

### HasPrimaryLabel

`func (o *HostedProviderRegionCreateRequest) HasPrimaryLabel() bool`

HasPrimaryLabel returns a boolean if a field has been set.

### SetPrimaryLabelNil

`func (o *HostedProviderRegionCreateRequest) SetPrimaryLabelNil(b bool)`

 SetPrimaryLabelNil sets the value for PrimaryLabel to be an explicit nil

### UnsetPrimaryLabel
`func (o *HostedProviderRegionCreateRequest) UnsetPrimaryLabel()`

UnsetPrimaryLabel ensures that no value is present for PrimaryLabel, not even an explicit nil
### GetPrimarySlug

`func (o *HostedProviderRegionCreateRequest) GetPrimarySlug() string`

GetPrimarySlug returns the PrimarySlug field if non-nil, zero value otherwise.

### GetPrimarySlugOk

`func (o *HostedProviderRegionCreateRequest) GetPrimarySlugOk() (*string, bool)`

GetPrimarySlugOk returns a tuple with the PrimarySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySlug

`func (o *HostedProviderRegionCreateRequest) SetPrimarySlug(v string)`

SetPrimarySlug sets PrimarySlug field to given value.

### HasPrimarySlug

`func (o *HostedProviderRegionCreateRequest) HasPrimarySlug() bool`

HasPrimarySlug returns a boolean if a field has been set.

### SetPrimarySlugNil

`func (o *HostedProviderRegionCreateRequest) SetPrimarySlugNil(b bool)`

 SetPrimarySlugNil sets the value for PrimarySlug to be an explicit nil

### UnsetPrimarySlug
`func (o *HostedProviderRegionCreateRequest) UnsetPrimarySlug()`

UnsetPrimarySlug ensures that no value is present for PrimarySlug, not even an explicit nil
### GetSlug

`func (o *HostedProviderRegionCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *HostedProviderRegionCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *HostedProviderRegionCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSupportedCountries

`func (o *HostedProviderRegionCreateRequest) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *HostedProviderRegionCreateRequest) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *HostedProviderRegionCreateRequest) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *HostedProviderRegionCreateRequest) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.

### SetSupportedCountriesNil

`func (o *HostedProviderRegionCreateRequest) SetSupportedCountriesNil(b bool)`

 SetSupportedCountriesNil sets the value for SupportedCountries to be an explicit nil

### UnsetSupportedCountries
`func (o *HostedProviderRegionCreateRequest) UnsetSupportedCountries()`

UnsetSupportedCountries ensures that no value is present for SupportedCountries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


