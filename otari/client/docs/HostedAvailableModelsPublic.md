# HostedAvailableModelsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryUnsupported** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to **NullableString** |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Provider** | **string** |  | 

## Methods

### NewHostedAvailableModelsPublic

`func NewHostedAvailableModelsPublic(provider string, ) *HostedAvailableModelsPublic`

NewHostedAvailableModelsPublic instantiates a new HostedAvailableModelsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedAvailableModelsPublicWithDefaults

`func NewHostedAvailableModelsPublicWithDefaults() *HostedAvailableModelsPublic`

NewHostedAvailableModelsPublicWithDefaults instantiates a new HostedAvailableModelsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryUnsupported

`func (o *HostedAvailableModelsPublic) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *HostedAvailableModelsPublic) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *HostedAvailableModelsPublic) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *HostedAvailableModelsPublic) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *HostedAvailableModelsPublic) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HostedAvailableModelsPublic) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HostedAvailableModelsPublic) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HostedAvailableModelsPublic) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HostedAvailableModelsPublic) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HostedAvailableModelsPublic) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetModels

`func (o *HostedAvailableModelsPublic) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *HostedAvailableModelsPublic) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *HostedAvailableModelsPublic) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *HostedAvailableModelsPublic) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProvider

`func (o *HostedAvailableModelsPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HostedAvailableModelsPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HostedAvailableModelsPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


