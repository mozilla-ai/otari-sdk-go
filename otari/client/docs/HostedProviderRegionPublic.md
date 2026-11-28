# HostedProviderRegionPublic

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
**HasApiKey** | Pointer to **bool** |  | [optional] [default to false]
**Label** | **string** |  | 
**Position** | **int32** |  | 
**Selector** | **string** |  | 
**Slug** | **string** |  | 
**SupportedCountries** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewHostedProviderRegionPublic

`func NewHostedProviderRegionPublic(enabled bool, label string, position int32, selector string, slug string, ) *HostedProviderRegionPublic`

NewHostedProviderRegionPublic instantiates a new HostedProviderRegionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedProviderRegionPublicWithDefaults

`func NewHostedProviderRegionPublicWithDefaults() *HostedProviderRegionPublic`

NewHostedProviderRegionPublicWithDefaults instantiates a new HostedProviderRegionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *HostedProviderRegionPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *HostedProviderRegionPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *HostedProviderRegionPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *HostedProviderRegionPublic) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *HostedProviderRegionPublic) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *HostedProviderRegionPublic) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKeyLast4

`func (o *HostedProviderRegionPublic) GetApiKeyLast4() string`

GetApiKeyLast4 returns the ApiKeyLast4 field if non-nil, zero value otherwise.

### GetApiKeyLast4Ok

`func (o *HostedProviderRegionPublic) GetApiKeyLast4Ok() (*string, bool)`

GetApiKeyLast4Ok returns a tuple with the ApiKeyLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyLast4

`func (o *HostedProviderRegionPublic) SetApiKeyLast4(v string)`

SetApiKeyLast4 sets ApiKeyLast4 field to given value.

### HasApiKeyLast4

`func (o *HostedProviderRegionPublic) HasApiKeyLast4() bool`

HasApiKeyLast4 returns a boolean if a field has been set.

### SetApiKeyLast4Nil

`func (o *HostedProviderRegionPublic) SetApiKeyLast4Nil(b bool)`

 SetApiKeyLast4Nil sets the value for ApiKeyLast4 to be an explicit nil

### UnsetApiKeyLast4
`func (o *HostedProviderRegionPublic) UnsetApiKeyLast4()`

UnsetApiKeyLast4 ensures that no value is present for ApiKeyLast4, not even an explicit nil
### GetClientArgs

`func (o *HostedProviderRegionPublic) GetClientArgs() map[string]interface{}`

GetClientArgs returns the ClientArgs field if non-nil, zero value otherwise.

### GetClientArgsOk

`func (o *HostedProviderRegionPublic) GetClientArgsOk() (*map[string]interface{}, bool)`

GetClientArgsOk returns a tuple with the ClientArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientArgs

`func (o *HostedProviderRegionPublic) SetClientArgs(v map[string]interface{})`

SetClientArgs sets ClientArgs field to given value.

### HasClientArgs

`func (o *HostedProviderRegionPublic) HasClientArgs() bool`

HasClientArgs returns a boolean if a field has been set.

### SetClientArgsNil

`func (o *HostedProviderRegionPublic) SetClientArgsNil(b bool)`

 SetClientArgsNil sets the value for ClientArgs to be an explicit nil

### UnsetClientArgs
`func (o *HostedProviderRegionPublic) UnsetClientArgs()`

UnsetClientArgs ensures that no value is present for ClientArgs, not even an explicit nil
### GetCreatedAt

`func (o *HostedProviderRegionPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostedProviderRegionPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostedProviderRegionPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HostedProviderRegionPublic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HostedProviderRegionPublic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HostedProviderRegionPublic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDataLocation

`func (o *HostedProviderRegionPublic) GetDataLocation() string`

GetDataLocation returns the DataLocation field if non-nil, zero value otherwise.

### GetDataLocationOk

`func (o *HostedProviderRegionPublic) GetDataLocationOk() (*string, bool)`

GetDataLocationOk returns a tuple with the DataLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLocation

`func (o *HostedProviderRegionPublic) SetDataLocation(v string)`

SetDataLocation sets DataLocation field to given value.

### HasDataLocation

`func (o *HostedProviderRegionPublic) HasDataLocation() bool`

HasDataLocation returns a boolean if a field has been set.

### SetDataLocationNil

`func (o *HostedProviderRegionPublic) SetDataLocationNil(b bool)`

 SetDataLocationNil sets the value for DataLocation to be an explicit nil

### UnsetDataLocation
`func (o *HostedProviderRegionPublic) UnsetDataLocation()`

UnsetDataLocation ensures that no value is present for DataLocation, not even an explicit nil
### GetDecryptable

`func (o *HostedProviderRegionPublic) GetDecryptable() bool`

GetDecryptable returns the Decryptable field if non-nil, zero value otherwise.

### GetDecryptableOk

`func (o *HostedProviderRegionPublic) GetDecryptableOk() (*bool, bool)`

GetDecryptableOk returns a tuple with the Decryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptable

`func (o *HostedProviderRegionPublic) SetDecryptable(v bool)`

SetDecryptable sets Decryptable field to given value.

### HasDecryptable

`func (o *HostedProviderRegionPublic) HasDecryptable() bool`

HasDecryptable returns a boolean if a field has been set.

### GetEnabled

`func (o *HostedProviderRegionPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedProviderRegionPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedProviderRegionPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasApiKey

`func (o *HostedProviderRegionPublic) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *HostedProviderRegionPublic) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *HostedProviderRegionPublic) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.

### HasHasApiKey

`func (o *HostedProviderRegionPublic) HasHasApiKey() bool`

HasHasApiKey returns a boolean if a field has been set.

### GetLabel

`func (o *HostedProviderRegionPublic) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HostedProviderRegionPublic) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HostedProviderRegionPublic) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPosition

`func (o *HostedProviderRegionPublic) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *HostedProviderRegionPublic) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *HostedProviderRegionPublic) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetSelector

`func (o *HostedProviderRegionPublic) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *HostedProviderRegionPublic) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *HostedProviderRegionPublic) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetSlug

`func (o *HostedProviderRegionPublic) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *HostedProviderRegionPublic) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *HostedProviderRegionPublic) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSupportedCountries

`func (o *HostedProviderRegionPublic) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *HostedProviderRegionPublic) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *HostedProviderRegionPublic) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *HostedProviderRegionPublic) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.

### SetSupportedCountriesNil

`func (o *HostedProviderRegionPublic) SetSupportedCountriesNil(b bool)`

 SetSupportedCountriesNil sets the value for SupportedCountries to be an explicit nil

### UnsetSupportedCountries
`func (o *HostedProviderRegionPublic) UnsetSupportedCountries()`

UnsetSupportedCountries ensures that no value is present for SupportedCountries, not even an explicit nil
### GetUpdatedAt

`func (o *HostedProviderRegionPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostedProviderRegionPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostedProviderRegionPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HostedProviderRegionPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HostedProviderRegionPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HostedProviderRegionPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


