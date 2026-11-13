# HostedCatalogProviderRefreshPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | **[]string** |  | 
**CredentialUnreadable** | Pointer to **bool** |  | [optional] [default to false]
**DiscoveryUnsupported** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** |  | 
**Repriced** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] [default to "database"]

## Methods

### NewHostedCatalogProviderRefreshPublic

`func NewHostedCatalogProviderRefreshPublic(added []string, provider string, ) *HostedCatalogProviderRefreshPublic`

NewHostedCatalogProviderRefreshPublic instantiates a new HostedCatalogProviderRefreshPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCatalogProviderRefreshPublicWithDefaults

`func NewHostedCatalogProviderRefreshPublicWithDefaults() *HostedCatalogProviderRefreshPublic`

NewHostedCatalogProviderRefreshPublicWithDefaults instantiates a new HostedCatalogProviderRefreshPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *HostedCatalogProviderRefreshPublic) GetAdded() []string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *HostedCatalogProviderRefreshPublic) GetAddedOk() (*[]string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *HostedCatalogProviderRefreshPublic) SetAdded(v []string)`

SetAdded sets Added field to given value.


### GetCredentialUnreadable

`func (o *HostedCatalogProviderRefreshPublic) GetCredentialUnreadable() bool`

GetCredentialUnreadable returns the CredentialUnreadable field if non-nil, zero value otherwise.

### GetCredentialUnreadableOk

`func (o *HostedCatalogProviderRefreshPublic) GetCredentialUnreadableOk() (*bool, bool)`

GetCredentialUnreadableOk returns a tuple with the CredentialUnreadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialUnreadable

`func (o *HostedCatalogProviderRefreshPublic) SetCredentialUnreadable(v bool)`

SetCredentialUnreadable sets CredentialUnreadable field to given value.

### HasCredentialUnreadable

`func (o *HostedCatalogProviderRefreshPublic) HasCredentialUnreadable() bool`

HasCredentialUnreadable returns a boolean if a field has been set.

### GetDiscoveryUnsupported

`func (o *HostedCatalogProviderRefreshPublic) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *HostedCatalogProviderRefreshPublic) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *HostedCatalogProviderRefreshPublic) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *HostedCatalogProviderRefreshPublic) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *HostedCatalogProviderRefreshPublic) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HostedCatalogProviderRefreshPublic) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HostedCatalogProviderRefreshPublic) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HostedCatalogProviderRefreshPublic) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HostedCatalogProviderRefreshPublic) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HostedCatalogProviderRefreshPublic) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProvider

`func (o *HostedCatalogProviderRefreshPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HostedCatalogProviderRefreshPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HostedCatalogProviderRefreshPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRepriced

`func (o *HostedCatalogProviderRefreshPublic) GetRepriced() []string`

GetRepriced returns the Repriced field if non-nil, zero value otherwise.

### GetRepricedOk

`func (o *HostedCatalogProviderRefreshPublic) GetRepricedOk() (*[]string, bool)`

GetRepricedOk returns a tuple with the Repriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepriced

`func (o *HostedCatalogProviderRefreshPublic) SetRepriced(v []string)`

SetRepriced sets Repriced field to given value.

### HasRepriced

`func (o *HostedCatalogProviderRefreshPublic) HasRepriced() bool`

HasRepriced returns a boolean if a field has been set.

### GetSource

`func (o *HostedCatalogProviderRefreshPublic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HostedCatalogProviderRefreshPublic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HostedCatalogProviderRefreshPublic) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HostedCatalogProviderRefreshPublic) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


