# UpdateCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | [**[]UpdateCheckAdvisory**](UpdateCheckAdvisory.md) | Security advisories that affect the version running now, as of the last check. | 
**AnsweredBy** | **NullableString** | Who decided: config.yml or the environment (the dashboard cannot change it), an operator in the dashboard, or nobody yet. | 
**Checkable** | **bool** | Whether this build carries a release version. A development build is never checked. | 
**CurrentVersion** | **string** |  | 
**Enabled** | **NullableBool** | Whether checks are sent. Null until an operator answers. | 
**Endpoint** | **string** | Where checks are sent. | 
**InstallId** | **string** | The random id sent with each check. | 
**LastCheckedAt** | **NullableTime** |  | 
**LatestVersion** | **NullableString** |  | 
**NextCheckAt** | **NullableTime** |  | 
**PresetBy** | **NullableString** | When the configuration decided, the setting that did: the OTARI_UPDATE_CHECK variable, update_check in config.yml, or DO_NOT_TRACK or CI answering no. Null otherwise. | 
**ReleaseUrl** | **NullableString** |  | 
**UpdateAvailable** | **bool** |  | 

## Methods

### NewUpdateCheckStatus

`func NewUpdateCheckStatus(advisories []UpdateCheckAdvisory, answeredBy NullableString, checkable bool, currentVersion string, enabled NullableBool, endpoint string, installId string, lastCheckedAt NullableTime, latestVersion NullableString, nextCheckAt NullableTime, presetBy NullableString, releaseUrl NullableString, updateAvailable bool, ) *UpdateCheckStatus`

NewUpdateCheckStatus instantiates a new UpdateCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCheckStatusWithDefaults

`func NewUpdateCheckStatusWithDefaults() *UpdateCheckStatus`

NewUpdateCheckStatusWithDefaults instantiates a new UpdateCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisories

`func (o *UpdateCheckStatus) GetAdvisories() []UpdateCheckAdvisory`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *UpdateCheckStatus) GetAdvisoriesOk() (*[]UpdateCheckAdvisory, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *UpdateCheckStatus) SetAdvisories(v []UpdateCheckAdvisory)`

SetAdvisories sets Advisories field to given value.


### GetAnsweredBy

`func (o *UpdateCheckStatus) GetAnsweredBy() string`

GetAnsweredBy returns the AnsweredBy field if non-nil, zero value otherwise.

### GetAnsweredByOk

`func (o *UpdateCheckStatus) GetAnsweredByOk() (*string, bool)`

GetAnsweredByOk returns a tuple with the AnsweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredBy

`func (o *UpdateCheckStatus) SetAnsweredBy(v string)`

SetAnsweredBy sets AnsweredBy field to given value.


### SetAnsweredByNil

`func (o *UpdateCheckStatus) SetAnsweredByNil(b bool)`

 SetAnsweredByNil sets the value for AnsweredBy to be an explicit nil

### UnsetAnsweredBy
`func (o *UpdateCheckStatus) UnsetAnsweredBy()`

UnsetAnsweredBy ensures that no value is present for AnsweredBy, not even an explicit nil
### GetCheckable

`func (o *UpdateCheckStatus) GetCheckable() bool`

GetCheckable returns the Checkable field if non-nil, zero value otherwise.

### GetCheckableOk

`func (o *UpdateCheckStatus) GetCheckableOk() (*bool, bool)`

GetCheckableOk returns a tuple with the Checkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckable

`func (o *UpdateCheckStatus) SetCheckable(v bool)`

SetCheckable sets Checkable field to given value.


### GetCurrentVersion

`func (o *UpdateCheckStatus) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *UpdateCheckStatus) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *UpdateCheckStatus) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetEnabled

`func (o *UpdateCheckStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCheckStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCheckStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *UpdateCheckStatus) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateCheckStatus) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetEndpoint

`func (o *UpdateCheckStatus) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateCheckStatus) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateCheckStatus) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetInstallId

`func (o *UpdateCheckStatus) GetInstallId() string`

GetInstallId returns the InstallId field if non-nil, zero value otherwise.

### GetInstallIdOk

`func (o *UpdateCheckStatus) GetInstallIdOk() (*string, bool)`

GetInstallIdOk returns a tuple with the InstallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallId

`func (o *UpdateCheckStatus) SetInstallId(v string)`

SetInstallId sets InstallId field to given value.


### GetLastCheckedAt

`func (o *UpdateCheckStatus) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *UpdateCheckStatus) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *UpdateCheckStatus) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.


### SetLastCheckedAtNil

`func (o *UpdateCheckStatus) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *UpdateCheckStatus) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetLatestVersion

`func (o *UpdateCheckStatus) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *UpdateCheckStatus) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *UpdateCheckStatus) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### SetLatestVersionNil

`func (o *UpdateCheckStatus) SetLatestVersionNil(b bool)`

 SetLatestVersionNil sets the value for LatestVersion to be an explicit nil

### UnsetLatestVersion
`func (o *UpdateCheckStatus) UnsetLatestVersion()`

UnsetLatestVersion ensures that no value is present for LatestVersion, not even an explicit nil
### GetNextCheckAt

`func (o *UpdateCheckStatus) GetNextCheckAt() time.Time`

GetNextCheckAt returns the NextCheckAt field if non-nil, zero value otherwise.

### GetNextCheckAtOk

`func (o *UpdateCheckStatus) GetNextCheckAtOk() (*time.Time, bool)`

GetNextCheckAtOk returns a tuple with the NextCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCheckAt

`func (o *UpdateCheckStatus) SetNextCheckAt(v time.Time)`

SetNextCheckAt sets NextCheckAt field to given value.


### SetNextCheckAtNil

`func (o *UpdateCheckStatus) SetNextCheckAtNil(b bool)`

 SetNextCheckAtNil sets the value for NextCheckAt to be an explicit nil

### UnsetNextCheckAt
`func (o *UpdateCheckStatus) UnsetNextCheckAt()`

UnsetNextCheckAt ensures that no value is present for NextCheckAt, not even an explicit nil
### GetPresetBy

`func (o *UpdateCheckStatus) GetPresetBy() string`

GetPresetBy returns the PresetBy field if non-nil, zero value otherwise.

### GetPresetByOk

`func (o *UpdateCheckStatus) GetPresetByOk() (*string, bool)`

GetPresetByOk returns a tuple with the PresetBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetBy

`func (o *UpdateCheckStatus) SetPresetBy(v string)`

SetPresetBy sets PresetBy field to given value.


### SetPresetByNil

`func (o *UpdateCheckStatus) SetPresetByNil(b bool)`

 SetPresetByNil sets the value for PresetBy to be an explicit nil

### UnsetPresetBy
`func (o *UpdateCheckStatus) UnsetPresetBy()`

UnsetPresetBy ensures that no value is present for PresetBy, not even an explicit nil
### GetReleaseUrl

`func (o *UpdateCheckStatus) GetReleaseUrl() string`

GetReleaseUrl returns the ReleaseUrl field if non-nil, zero value otherwise.

### GetReleaseUrlOk

`func (o *UpdateCheckStatus) GetReleaseUrlOk() (*string, bool)`

GetReleaseUrlOk returns a tuple with the ReleaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseUrl

`func (o *UpdateCheckStatus) SetReleaseUrl(v string)`

SetReleaseUrl sets ReleaseUrl field to given value.


### SetReleaseUrlNil

`func (o *UpdateCheckStatus) SetReleaseUrlNil(b bool)`

 SetReleaseUrlNil sets the value for ReleaseUrl to be an explicit nil

### UnsetReleaseUrl
`func (o *UpdateCheckStatus) UnsetReleaseUrl()`

UnsetReleaseUrl ensures that no value is present for ReleaseUrl, not even an explicit nil
### GetUpdateAvailable

`func (o *UpdateCheckStatus) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *UpdateCheckStatus) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *UpdateCheckStatus) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


