# ProviderInfoSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**ProviderCapabilitiesSchema**](ProviderCapabilitiesSchema.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DocUrl** | Pointer to **NullableString** |  | [optional] 
**EnvKey** | Pointer to **NullableString** | Env var the credential is read from. | [optional] 
**Instance** | **string** | Configured provider key (may differ from the type). | 
**Name** | **string** | Human-friendly provider name. | 
**PricingUrls** | Pointer to **[]string** |  | [optional] 
**ProviderType** | **string** | Underlying any-llm provider type. | 

## Methods

### NewProviderInfoSchema

`func NewProviderInfoSchema(capabilities ProviderCapabilitiesSchema, instance string, name string, providerType string, ) *ProviderInfoSchema`

NewProviderInfoSchema instantiates a new ProviderInfoSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInfoSchemaWithDefaults

`func NewProviderInfoSchemaWithDefaults() *ProviderInfoSchema`

NewProviderInfoSchemaWithDefaults instantiates a new ProviderInfoSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ProviderInfoSchema) GetCapabilities() ProviderCapabilitiesSchema`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ProviderInfoSchema) GetCapabilitiesOk() (*ProviderCapabilitiesSchema, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ProviderInfoSchema) SetCapabilities(v ProviderCapabilitiesSchema)`

SetCapabilities sets Capabilities field to given value.


### GetDescription

`func (o *ProviderInfoSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderInfoSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderInfoSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderInfoSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProviderInfoSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProviderInfoSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocUrl

`func (o *ProviderInfoSchema) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *ProviderInfoSchema) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *ProviderInfoSchema) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *ProviderInfoSchema) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.

### SetDocUrlNil

`func (o *ProviderInfoSchema) SetDocUrlNil(b bool)`

 SetDocUrlNil sets the value for DocUrl to be an explicit nil

### UnsetDocUrl
`func (o *ProviderInfoSchema) UnsetDocUrl()`

UnsetDocUrl ensures that no value is present for DocUrl, not even an explicit nil
### GetEnvKey

`func (o *ProviderInfoSchema) GetEnvKey() string`

GetEnvKey returns the EnvKey field if non-nil, zero value otherwise.

### GetEnvKeyOk

`func (o *ProviderInfoSchema) GetEnvKeyOk() (*string, bool)`

GetEnvKeyOk returns a tuple with the EnvKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvKey

`func (o *ProviderInfoSchema) SetEnvKey(v string)`

SetEnvKey sets EnvKey field to given value.

### HasEnvKey

`func (o *ProviderInfoSchema) HasEnvKey() bool`

HasEnvKey returns a boolean if a field has been set.

### SetEnvKeyNil

`func (o *ProviderInfoSchema) SetEnvKeyNil(b bool)`

 SetEnvKeyNil sets the value for EnvKey to be an explicit nil

### UnsetEnvKey
`func (o *ProviderInfoSchema) UnsetEnvKey()`

UnsetEnvKey ensures that no value is present for EnvKey, not even an explicit nil
### GetInstance

`func (o *ProviderInfoSchema) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProviderInfoSchema) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProviderInfoSchema) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetName

`func (o *ProviderInfoSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderInfoSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderInfoSchema) SetName(v string)`

SetName sets Name field to given value.


### GetPricingUrls

`func (o *ProviderInfoSchema) GetPricingUrls() []string`

GetPricingUrls returns the PricingUrls field if non-nil, zero value otherwise.

### GetPricingUrlsOk

`func (o *ProviderInfoSchema) GetPricingUrlsOk() (*[]string, bool)`

GetPricingUrlsOk returns a tuple with the PricingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUrls

`func (o *ProviderInfoSchema) SetPricingUrls(v []string)`

SetPricingUrls sets PricingUrls field to given value.

### HasPricingUrls

`func (o *ProviderInfoSchema) HasPricingUrls() bool`

HasPricingUrls returns a boolean if a field has been set.

### GetProviderType

`func (o *ProviderInfoSchema) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ProviderInfoSchema) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ProviderInfoSchema) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


