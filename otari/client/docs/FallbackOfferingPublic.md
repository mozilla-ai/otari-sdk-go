# FallbackOfferingPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtariTools** | **[]string** | Built-in tools Otari runs in this provider&#39;s place, for that caller. | 
**OwnTools** | **[]string** | Built-in tools this provider runs itself, for a caller that declares them in the provider&#39;s own words and leaves the choice of executor to the deployment. | 
**Provider** | **string** | The provider instance the selector names. | 
**ProviderType** | **string** | The any-llm implementation behind the instance. | 
**Selector** | **string** | The offering, in &#x60;instance:model&#x60; form. | 

## Methods

### NewFallbackOfferingPublic

`func NewFallbackOfferingPublic(otariTools []string, ownTools []string, provider string, providerType string, selector string, ) *FallbackOfferingPublic`

NewFallbackOfferingPublic instantiates a new FallbackOfferingPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackOfferingPublicWithDefaults

`func NewFallbackOfferingPublicWithDefaults() *FallbackOfferingPublic`

NewFallbackOfferingPublicWithDefaults instantiates a new FallbackOfferingPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtariTools

`func (o *FallbackOfferingPublic) GetOtariTools() []string`

GetOtariTools returns the OtariTools field if non-nil, zero value otherwise.

### GetOtariToolsOk

`func (o *FallbackOfferingPublic) GetOtariToolsOk() (*[]string, bool)`

GetOtariToolsOk returns a tuple with the OtariTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtariTools

`func (o *FallbackOfferingPublic) SetOtariTools(v []string)`

SetOtariTools sets OtariTools field to given value.


### GetOwnTools

`func (o *FallbackOfferingPublic) GetOwnTools() []string`

GetOwnTools returns the OwnTools field if non-nil, zero value otherwise.

### GetOwnToolsOk

`func (o *FallbackOfferingPublic) GetOwnToolsOk() (*[]string, bool)`

GetOwnToolsOk returns a tuple with the OwnTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnTools

`func (o *FallbackOfferingPublic) SetOwnTools(v []string)`

SetOwnTools sets OwnTools field to given value.


### GetProvider

`func (o *FallbackOfferingPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FallbackOfferingPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FallbackOfferingPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderType

`func (o *FallbackOfferingPublic) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *FallbackOfferingPublic) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *FallbackOfferingPublic) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetSelector

`func (o *FallbackOfferingPublic) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *FallbackOfferingPublic) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *FallbackOfferingPublic) SetSelector(v string)`

SetSelector sets Selector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


