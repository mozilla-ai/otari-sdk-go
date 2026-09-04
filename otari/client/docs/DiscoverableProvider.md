# DiscoverableProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **NullableString** | When this instance was last dialed, ISO 8601. Null when it has not been checked yet, which is what the first read after a restart sees while the background refresh runs. | [optional] 
**DiscoveryUnsupported** | Pointer to **bool** | True when discovery failed only because this backend serves no model-listing endpoint. The provider may still handle requests for models declared in config. | [optional] [default to false]
**Error** | Pointer to **NullableString** | Why discovery failed. Null when &#x60;ok&#x60; is true. | [optional] 
**Models** | [**[]DiscoverableModel**](DiscoverableModel.md) |  | 
**Ok** | **bool** | False when this instance could not be queried. | 
**Provider** | **string** |  | 

## Methods

### NewDiscoverableProvider

`func NewDiscoverableProvider(models []DiscoverableModel, ok bool, provider string, ) *DiscoverableProvider`

NewDiscoverableProvider instantiates a new DiscoverableProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoverableProviderWithDefaults

`func NewDiscoverableProviderWithDefaults() *DiscoverableProvider`

NewDiscoverableProviderWithDefaults instantiates a new DiscoverableProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *DiscoverableProvider) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *DiscoverableProvider) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *DiscoverableProvider) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *DiscoverableProvider) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### SetCheckedAtNil

`func (o *DiscoverableProvider) SetCheckedAtNil(b bool)`

 SetCheckedAtNil sets the value for CheckedAt to be an explicit nil

### UnsetCheckedAt
`func (o *DiscoverableProvider) UnsetCheckedAt()`

UnsetCheckedAt ensures that no value is present for CheckedAt, not even an explicit nil
### GetDiscoveryUnsupported

`func (o *DiscoverableProvider) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *DiscoverableProvider) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *DiscoverableProvider) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *DiscoverableProvider) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *DiscoverableProvider) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DiscoverableProvider) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DiscoverableProvider) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DiscoverableProvider) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DiscoverableProvider) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DiscoverableProvider) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetModels

`func (o *DiscoverableProvider) GetModels() []DiscoverableModel`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *DiscoverableProvider) GetModelsOk() (*[]DiscoverableModel, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *DiscoverableProvider) SetModels(v []DiscoverableModel)`

SetModels sets Models field to given value.


### GetOk

`func (o *DiscoverableProvider) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DiscoverableProvider) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DiscoverableProvider) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetProvider

`func (o *DiscoverableProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DiscoverableProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DiscoverableProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


