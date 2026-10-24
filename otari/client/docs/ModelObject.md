# ModelObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to [**NullableModelArchitecture**](ModelArchitecture.md) |  | [optional] 
**ContextWindow** | Pointer to **NullableInt32** |  | [optional] 
**Created** | **int32** |  | 
**DeploymentManaged** | Pointer to **bool** |  | [optional] [default to false]
**Id** | **string** |  | 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**OwnedBy** | **string** |  | 
**Pricing** | Pointer to [**NullableModelPricingInfo**](ModelPricingInfo.md) |  | [optional] 
**PricingSource** | Pointer to **string** |  | [optional] [default to "none"]
**SupportedParameters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelObject

`func NewModelObject(created int32, id string, ownedBy string, ) *ModelObject`

NewModelObject instantiates a new ModelObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelObjectWithDefaults

`func NewModelObjectWithDefaults() *ModelObject`

NewModelObjectWithDefaults instantiates a new ModelObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ModelObject) GetArchitecture() ModelArchitecture`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ModelObject) GetArchitectureOk() (*ModelArchitecture, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ModelObject) SetArchitecture(v ModelArchitecture)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ModelObject) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *ModelObject) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *ModelObject) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetContextWindow

`func (o *ModelObject) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *ModelObject) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *ModelObject) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *ModelObject) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### SetContextWindowNil

`func (o *ModelObject) SetContextWindowNil(b bool)`

 SetContextWindowNil sets the value for ContextWindow to be an explicit nil

### UnsetContextWindow
`func (o *ModelObject) UnsetContextWindow()`

UnsetContextWindow ensures that no value is present for ContextWindow, not even an explicit nil
### GetCreated

`func (o *ModelObject) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelObject) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelObject) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetDeploymentManaged

`func (o *ModelObject) GetDeploymentManaged() bool`

GetDeploymentManaged returns the DeploymentManaged field if non-nil, zero value otherwise.

### GetDeploymentManagedOk

`func (o *ModelObject) GetDeploymentManagedOk() (*bool, bool)`

GetDeploymentManagedOk returns a tuple with the DeploymentManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentManaged

`func (o *ModelObject) SetDeploymentManaged(v bool)`

SetDeploymentManaged sets DeploymentManaged field to given value.

### HasDeploymentManaged

`func (o *ModelObject) HasDeploymentManaged() bool`

HasDeploymentManaged returns a boolean if a field has been set.

### GetId

`func (o *ModelObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelObject) SetId(v string)`

SetId sets Id field to given value.


### GetMaxOutputTokens

`func (o *ModelObject) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *ModelObject) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *ModelObject) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *ModelObject) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *ModelObject) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *ModelObject) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetName

`func (o *ModelObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObject

`func (o *ModelObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ModelObject) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelObject) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelObject) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetPricing

`func (o *ModelObject) GetPricing() ModelPricingInfo`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ModelObject) GetPricingOk() (*ModelPricingInfo, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ModelObject) SetPricing(v ModelPricingInfo)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *ModelObject) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *ModelObject) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *ModelObject) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetPricingSource

`func (o *ModelObject) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *ModelObject) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *ModelObject) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.

### HasPricingSource

`func (o *ModelObject) HasPricingSource() bool`

HasPricingSource returns a boolean if a field has been set.

### GetSupportedParameters

`func (o *ModelObject) GetSupportedParameters() []string`

GetSupportedParameters returns the SupportedParameters field if non-nil, zero value otherwise.

### GetSupportedParametersOk

`func (o *ModelObject) GetSupportedParametersOk() (*[]string, bool)`

GetSupportedParametersOk returns a tuple with the SupportedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedParameters

`func (o *ModelObject) SetSupportedParameters(v []string)`

SetSupportedParameters sets SupportedParameters field to given value.

### HasSupportedParameters

`func (o *ModelObject) HasSupportedParameters() bool`

HasSupportedParameters returns a boolean if a field has been set.

### SetSupportedParametersNil

`func (o *ModelObject) SetSupportedParametersNil(b bool)`

 SetSupportedParametersNil sets the value for SupportedParameters to be an explicit nil

### UnsetSupportedParameters
`func (o *ModelObject) UnsetSupportedParameters()`

UnsetSupportedParameters ensures that no value is present for SupportedParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


