# OrgProviderAvailableModelsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryUnsupported** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to **NullableString** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Provider** | **string** |  | 

## Methods

### NewOrgProviderAvailableModelsPublic

`func NewOrgProviderAvailableModelsPublic(provider string, ) *OrgProviderAvailableModelsPublic`

NewOrgProviderAvailableModelsPublic instantiates a new OrgProviderAvailableModelsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgProviderAvailableModelsPublicWithDefaults

`func NewOrgProviderAvailableModelsPublicWithDefaults() *OrgProviderAvailableModelsPublic`

NewOrgProviderAvailableModelsPublicWithDefaults instantiates a new OrgProviderAvailableModelsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryUnsupported

`func (o *OrgProviderAvailableModelsPublic) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *OrgProviderAvailableModelsPublic) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *OrgProviderAvailableModelsPublic) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *OrgProviderAvailableModelsPublic) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *OrgProviderAvailableModelsPublic) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OrgProviderAvailableModelsPublic) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OrgProviderAvailableModelsPublic) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OrgProviderAvailableModelsPublic) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OrgProviderAvailableModelsPublic) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OrgProviderAvailableModelsPublic) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetModels

`func (o *OrgProviderAvailableModelsPublic) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *OrgProviderAvailableModelsPublic) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *OrgProviderAvailableModelsPublic) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *OrgProviderAvailableModelsPublic) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProvider

`func (o *OrgProviderAvailableModelsPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrgProviderAvailableModelsPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrgProviderAvailableModelsPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


