# TestProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryUnsupported** | Pointer to **bool** | True when the test failed only because this backend serves no model-listing endpoint, so the credentials could not be verified this way but may still work for requests. | [optional] [default to false]
**Error** | Pointer to **NullableString** |  | [optional] 
**ModelCount** | **int32** |  | 
**Ok** | **bool** |  | 

## Methods

### NewTestProviderResponse

`func NewTestProviderResponse(modelCount int32, ok bool, ) *TestProviderResponse`

NewTestProviderResponse instantiates a new TestProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestProviderResponseWithDefaults

`func NewTestProviderResponseWithDefaults() *TestProviderResponse`

NewTestProviderResponseWithDefaults instantiates a new TestProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryUnsupported

`func (o *TestProviderResponse) GetDiscoveryUnsupported() bool`

GetDiscoveryUnsupported returns the DiscoveryUnsupported field if non-nil, zero value otherwise.

### GetDiscoveryUnsupportedOk

`func (o *TestProviderResponse) GetDiscoveryUnsupportedOk() (*bool, bool)`

GetDiscoveryUnsupportedOk returns a tuple with the DiscoveryUnsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryUnsupported

`func (o *TestProviderResponse) SetDiscoveryUnsupported(v bool)`

SetDiscoveryUnsupported sets DiscoveryUnsupported field to given value.

### HasDiscoveryUnsupported

`func (o *TestProviderResponse) HasDiscoveryUnsupported() bool`

HasDiscoveryUnsupported returns a boolean if a field has been set.

### GetError

`func (o *TestProviderResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TestProviderResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TestProviderResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TestProviderResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TestProviderResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TestProviderResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetModelCount

`func (o *TestProviderResponse) GetModelCount() int32`

GetModelCount returns the ModelCount field if non-nil, zero value otherwise.

### GetModelCountOk

`func (o *TestProviderResponse) GetModelCountOk() (*int32, bool)`

GetModelCountOk returns a tuple with the ModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelCount

`func (o *TestProviderResponse) SetModelCount(v int32)`

SetModelCount sets ModelCount field to given value.


### GetOk

`func (o *TestProviderResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *TestProviderResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *TestProviderResponse) SetOk(v bool)`

SetOk sets Ok field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


