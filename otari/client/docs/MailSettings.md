# MailSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether a transport is configured at all. | 
**FromEmail** | **NullableString** | The &#39;From&#39; address on outgoing mail, if one is configured. | 
**FromName** | **string** | The &#39;From&#39; display name on outgoing mail. | 
**Missing** | **[]string** | Settings that must be set before mail works, in config order. Empty exactly when &#39;ready&#39; is true, so the dashboard can name what to configure rather than only reporting that mail is off. | 
**PublicBaseUrl** | **NullableString** | This deployment&#39;s own externally-reachable URL, used to build links in outgoing mail. | 
**Ready** | **bool** | Whether a message carrying a link back to this deployment can be sent, which is what every message the control plane sends needs. Matches &#39;mail_ready&#39; on /v1/bootstrap. | 
**Transport** | **string** | The transport a send would use: &#39;smtp&#39;, &#39;console&#39; (logged, not delivered), or &#39;none&#39;. | 

## Methods

### NewMailSettings

`func NewMailSettings(enabled bool, fromEmail NullableString, fromName string, missing []string, publicBaseUrl NullableString, ready bool, transport string, ) *MailSettings`

NewMailSettings instantiates a new MailSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailSettingsWithDefaults

`func NewMailSettingsWithDefaults() *MailSettings`

NewMailSettingsWithDefaults instantiates a new MailSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MailSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MailSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MailSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFromEmail

`func (o *MailSettings) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *MailSettings) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *MailSettings) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### SetFromEmailNil

`func (o *MailSettings) SetFromEmailNil(b bool)`

 SetFromEmailNil sets the value for FromEmail to be an explicit nil

### UnsetFromEmail
`func (o *MailSettings) UnsetFromEmail()`

UnsetFromEmail ensures that no value is present for FromEmail, not even an explicit nil
### GetFromName

`func (o *MailSettings) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *MailSettings) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *MailSettings) SetFromName(v string)`

SetFromName sets FromName field to given value.


### GetMissing

`func (o *MailSettings) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *MailSettings) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *MailSettings) SetMissing(v []string)`

SetMissing sets Missing field to given value.


### GetPublicBaseUrl

`func (o *MailSettings) GetPublicBaseUrl() string`

GetPublicBaseUrl returns the PublicBaseUrl field if non-nil, zero value otherwise.

### GetPublicBaseUrlOk

`func (o *MailSettings) GetPublicBaseUrlOk() (*string, bool)`

GetPublicBaseUrlOk returns a tuple with the PublicBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicBaseUrl

`func (o *MailSettings) SetPublicBaseUrl(v string)`

SetPublicBaseUrl sets PublicBaseUrl field to given value.


### SetPublicBaseUrlNil

`func (o *MailSettings) SetPublicBaseUrlNil(b bool)`

 SetPublicBaseUrlNil sets the value for PublicBaseUrl to be an explicit nil

### UnsetPublicBaseUrl
`func (o *MailSettings) UnsetPublicBaseUrl()`

UnsetPublicBaseUrl ensures that no value is present for PublicBaseUrl, not even an explicit nil
### GetReady

`func (o *MailSettings) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MailSettings) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MailSettings) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetTransport

`func (o *MailSettings) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *MailSettings) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *MailSettings) SetTransport(v string)`

SetTransport sets Transport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


