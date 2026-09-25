# OrgProviderModelsRefreshPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | **[]string** |  | 
**Count** | **int32** |  | 
**DiscoveryUnsupported** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to **NullableString** |  | [optional] 
**Repriced** | **[]string** |  | 

## Methods

### NewOrgProviderModelsRefreshPublic

`func NewOrgProviderModelsRefreshPublic(added []string, count int32, repriced []string, ) *OrgProviderModelsRefreshPublic`

NewOrgProviderModelsRefreshPublic instantiates a new OrgProviderModelsRefreshPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgProviderModelsRefreshPublicWithDefaults

`func NewOrgProviderModelsRefreshPublicWithDefaults() *OrgProviderModelsRefreshPublic`

NewOrgProviderModelsRefreshPublicWithDefaults instantiates a new OrgProviderModelsRefreshPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *OrgProviderModelsRefreshPublic) GetAdded() []string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *OrgProviderModelsRefreshPublic) GetAddedOk() (*[]string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *OrgProviderModelsRefreshPublic) SetAdded(v []string)`

SetAdded sets Added field to given value.


### GetCount

`func (o *OrgProviderModelsRefreshPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrgProviderModelsRefreshPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrgProviderModelsRefreshPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDiscoveryUnsupported

`func (o *OrgProviderModelsRefreshPublic) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *OrgProviderModelsRefreshPublic) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *OrgProviderModelsRefreshPublic) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *OrgProviderModelsRefreshPublic) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *OrgProviderModelsRefreshPublic) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OrgProviderModelsRefreshPublic) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OrgProviderModelsRefreshPublic) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OrgProviderModelsRefreshPublic) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OrgProviderModelsRefreshPublic) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OrgProviderModelsRefreshPublic) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRepriced

`func (o *OrgProviderModelsRefreshPublic) GetRepriced() []string`

GetRepriced returns the Repriced field if non-nil, zero value otherwise.

### GetRepricedOk

`func (o *OrgProviderModelsRefreshPublic) GetRepricedOk() (*[]string, bool)`

GetRepricedOk returns a tuple with the Repriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepriced

`func (o *OrgProviderModelsRefreshPublic) SetRepriced(v []string)`

SetRepriced sets Repriced field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


