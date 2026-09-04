# GatewaySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**[]ConfigField**](ConfigField.md) |  | 
**DefaultPricing** | **bool** |  | 
**MasterKeySource** | **string** | Whether the dashboard master key is configured at startup or generated and stored by Otari. | 
**Mode** | **string** |  | 
**ModelDiscovery** | **bool** |  | 
**RequirePricing** | **bool** |  | 
**SecretKeyConfigured** | **bool** | Whether OTARI_SECRET_KEY is set on the server. Provider credentials are encrypted at rest with it, so a deployment without it can store none. The dashboard reads the same fact from the membership context&#39;s provider_key_encryption_available, because this endpoint is operator-only and the provider-key pages are read by tenants. | 
**Version** | **string** |  | 

## Methods

### NewGatewaySettings

`func NewGatewaySettings(config []ConfigField, defaultPricing bool, masterKeySource string, mode string, modelDiscovery bool, requirePricing bool, secretKeyConfigured bool, version string, ) *GatewaySettings`

NewGatewaySettings instantiates a new GatewaySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySettingsWithDefaults

`func NewGatewaySettingsWithDefaults() *GatewaySettings`

NewGatewaySettingsWithDefaults instantiates a new GatewaySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GatewaySettings) GetConfig() []ConfigField`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GatewaySettings) GetConfigOk() (*[]ConfigField, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GatewaySettings) SetConfig(v []ConfigField)`

SetConfig sets Config field to given value.


### GetDefaultPricing

`func (o *GatewaySettings) GetDefaultPricing() bool`

GetDefaultPricing returns the DefaultPricing field if non-nil, zero value otherwise.

### GetDefaultPricingOk

`func (o *GatewaySettings) GetDefaultPricingOk() (*bool, bool)`

GetDefaultPricingOk returns a tuple with the DefaultPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPricing

`func (o *GatewaySettings) SetDefaultPricing(v bool)`

SetDefaultPricing sets DefaultPricing field to given value.


### GetMasterKeySource

`func (o *GatewaySettings) GetMasterKeySource() string`

GetMasterKeySource returns the MasterKeySource field if non-nil, zero value otherwise.

### GetMasterKeySourceOk

`func (o *GatewaySettings) GetMasterKeySourceOk() (*string, bool)`

GetMasterKeySourceOk returns a tuple with the MasterKeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeySource

`func (o *GatewaySettings) SetMasterKeySource(v string)`

SetMasterKeySource sets MasterKeySource field to given value.


### GetMode

`func (o *GatewaySettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewaySettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewaySettings) SetMode(v string)`

SetMode sets Mode field to given value.


### GetModelDiscovery

`func (o *GatewaySettings) GetModelDiscovery() bool`

GetModelDiscovery returns the ModelDiscovery field if non-nil, zero value otherwise.

### GetModelDiscoveryOk

`func (o *GatewaySettings) GetModelDiscoveryOk() (*bool, bool)`

GetModelDiscoveryOk returns a tuple with the ModelDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDiscovery

`func (o *GatewaySettings) SetModelDiscovery(v bool)`

SetModelDiscovery sets ModelDiscovery field to given value.


### GetRequirePricing

`func (o *GatewaySettings) GetRequirePricing() bool`

GetRequirePricing returns the RequirePricing field if non-nil, zero value otherwise.

### GetRequirePricingOk

`func (o *GatewaySettings) GetRequirePricingOk() (*bool, bool)`

GetRequirePricingOk returns a tuple with the RequirePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePricing

`func (o *GatewaySettings) SetRequirePricing(v bool)`

SetRequirePricing sets RequirePricing field to given value.


### GetSecretKeyConfigured

`func (o *GatewaySettings) GetSecretKeyConfigured() bool`

GetSecretKeyConfigured returns the SecretKeyConfigured field if non-nil, zero value otherwise.

### GetSecretKeyConfiguredOk

`func (o *GatewaySettings) GetSecretKeyConfiguredOk() (*bool, bool)`

GetSecretKeyConfiguredOk returns a tuple with the SecretKeyConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyConfigured

`func (o *GatewaySettings) SetSecretKeyConfigured(v bool)`

SetSecretKeyConfigured sets SecretKeyConfigured field to given value.


### GetVersion

`func (o *GatewaySettings) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewaySettings) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewaySettings) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


