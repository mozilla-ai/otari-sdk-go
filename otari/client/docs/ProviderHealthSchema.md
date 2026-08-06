# ProviderHealthSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **NullableString** | ISO 8601 wall-clock time the provider&#39;s reachability was last checked (null if never). | [optional] 
**DiscoveryUnsupported** | Pointer to **bool** | True when the check failed only because this backend serves no model-listing endpoint. The provider may still handle requests; only model discovery is unavailable. | [optional] [default to false]
**Error** | Pointer to **NullableString** | Sanitized provider error when unreachable. | [optional] 
**Instance** | **string** |  | 
**ModelCount** | **int32** | Number of models the last successful listing returned. | 
**Ok** | **bool** | True when the provider&#39;s credentials could list models. | 

## Methods

### NewProviderHealthSchema

`func NewProviderHealthSchema(instance string, modelCount int32, ok bool, ) *ProviderHealthSchema`

NewProviderHealthSchema instantiates a new ProviderHealthSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderHealthSchemaWithDefaults

`func NewProviderHealthSchemaWithDefaults() *ProviderHealthSchema`

NewProviderHealthSchemaWithDefaults instantiates a new ProviderHealthSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *ProviderHealthSchema) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *ProviderHealthSchema) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *ProviderHealthSchema) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *ProviderHealthSchema) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### SetCheckedAtNil

`func (o *ProviderHealthSchema) SetCheckedAtNil(b bool)`

 SetCheckedAtNil sets the value for CheckedAt to be an explicit nil

### UnsetCheckedAt
`func (o *ProviderHealthSchema) UnsetCheckedAt()`

UnsetCheckedAt ensures that no value is present for CheckedAt, not even an explicit nil
### GetDiscoveryUnsupported

`func (o *ProviderHealthSchema) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *ProviderHealthSchema) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *ProviderHealthSchema) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *ProviderHealthSchema) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *ProviderHealthSchema) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProviderHealthSchema) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProviderHealthSchema) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ProviderHealthSchema) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ProviderHealthSchema) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ProviderHealthSchema) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetInstance

`func (o *ProviderHealthSchema) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProviderHealthSchema) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProviderHealthSchema) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetModelCount

`func (o *ProviderHealthSchema) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *ProviderHealthSchema) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *ProviderHealthSchema) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.


### GetOk

`func (o *ProviderHealthSchema) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ProviderHealthSchema) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ProviderHealthSchema) SetOk(v bool)`

SetOk sets Ok field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


