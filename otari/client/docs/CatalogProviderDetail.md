# CatalogProviderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | models.dev&#39;s name for the implementation, with the instance name after it where the two differ; the instance name where models.dev does not know it. | 
**ModelCount** | **int32** | Catalog models the caller may see that this instance offers. | 
**Models** | [**[]CatalogProviderModel**](CatalogProviderModel.md) | The requested page; &#x60;model_count&#x60; is the total. | 
**Name** | **string** | The instance name a selector starts with, such as &#x60;fireworks&#x60; or &#x60;openai-eu&#x60;. | 
**PrivacyUrl** | Pointer to **NullableString** | The provider&#39;s privacy policy; null where none is set. | [optional] 
**Provider** | **string** | The any-llm implementation behind the instance. | 
**TermsUrl** | Pointer to **NullableString** | The provider&#39;s terms of service; null where none is set. | [optional] 

## Methods

### NewCatalogProviderDetail

`func NewCatalogProviderDetail(displayName string, modelCount int32, models []CatalogProviderModel, name string, provider string, ) *CatalogProviderDetail`

NewCatalogProviderDetail instantiates a new CatalogProviderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogProviderDetailWithDefaults

`func NewCatalogProviderDetailWithDefaults() *CatalogProviderDetail`

NewCatalogProviderDetailWithDefaults instantiates a new CatalogProviderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CatalogProviderDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogProviderDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogProviderDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetModelCount

`func (o *CatalogProviderDetail) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *CatalogProviderDetail) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *CatalogProviderDetail) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.


### GetModels

`func (o *CatalogProviderDetail) GetModels() []CatalogProviderModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CatalogProviderDetail) GetModelsOk() (*[]CatalogProviderModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CatalogProviderDetail) SetModels(v []CatalogProviderModel)`

SetModels sets Models field to given value.


### GetName

`func (o *CatalogProviderDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogProviderDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogProviderDetail) SetName(v string)`

SetName sets Name field to given value.


### GetPrivacyUrl

`func (o *CatalogProviderDetail) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *CatalogProviderDetail) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *CatalogProviderDetail) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *CatalogProviderDetail) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *CatalogProviderDetail) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *CatalogProviderDetail) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetProvider

`func (o *CatalogProviderDetail) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CatalogProviderDetail) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CatalogProviderDetail) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetTermsUrl

`func (o *CatalogProviderDetail) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *CatalogProviderDetail) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *CatalogProviderDetail) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *CatalogProviderDetail) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *CatalogProviderDetail) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *CatalogProviderDetail) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


