# CatalogProviderSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | models.dev&#39;s name for the implementation, with the instance name after it where the two differ; the instance name where models.dev does not know it. | 
**ModelCount** | **int32** | Catalog models the caller may see that this instance offers. | 
**Name** | **string** | The instance name a selector starts with, such as &#x60;fireworks&#x60; or &#x60;openai-eu&#x60;. | 
**PrivacyUrl** | Pointer to **NullableString** | The provider&#39;s privacy policy; null where none is set. | [optional] 
**Provider** | **string** | The any-llm implementation behind the instance. | 
**TermsUrl** | Pointer to **NullableString** | The provider&#39;s terms of service; null where none is set. | [optional] 

## Methods

### NewCatalogProviderSummary

`func NewCatalogProviderSummary(displayName string, modelCount int32, name string, provider string, ) *CatalogProviderSummary`

NewCatalogProviderSummary instantiates a new CatalogProviderSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogProviderSummaryWithDefaults

`func NewCatalogProviderSummaryWithDefaults() *CatalogProviderSummary`

NewCatalogProviderSummaryWithDefaults instantiates a new CatalogProviderSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CatalogProviderSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogProviderSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogProviderSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetModelCount

`func (o *CatalogProviderSummary) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *CatalogProviderSummary) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *CatalogProviderSummary) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.


### GetName

`func (o *CatalogProviderSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogProviderSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogProviderSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPrivacyUrl

`func (o *CatalogProviderSummary) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *CatalogProviderSummary) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *CatalogProviderSummary) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *CatalogProviderSummary) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *CatalogProviderSummary) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *CatalogProviderSummary) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetProvider

`func (o *CatalogProviderSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CatalogProviderSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CatalogProviderSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetTermsUrl

`func (o *CatalogProviderSummary) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *CatalogProviderSummary) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *CatalogProviderSummary) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *CatalogProviderSummary) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *CatalogProviderSummary) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *CatalogProviderSummary) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


