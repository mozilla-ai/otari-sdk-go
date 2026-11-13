# GuardrailDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Gates** | [**[]GateDefinitionRequest**](GateDefinitionRequest.md) |  | 

## Methods

### NewGuardrailDefinitionRequest

`func NewGuardrailDefinitionRequest(digest string, gates []GateDefinitionRequest, ) *GuardrailDefinitionRequest`

NewGuardrailDefinitionRequest instantiates a new GuardrailDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailDefinitionRequestWithDefaults

`func NewGuardrailDefinitionRequestWithDefaults() *GuardrailDefinitionRequest`

NewGuardrailDefinitionRequestWithDefaults instantiates a new GuardrailDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *GuardrailDefinitionRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GuardrailDefinitionRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GuardrailDefinitionRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetGates

`func (o *GuardrailDefinitionRequest) GetGates() []GateDefinitionRequest`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *GuardrailDefinitionRequest) GetGatesOk() (*[]GateDefinitionRequest, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *GuardrailDefinitionRequest) SetGates(v []GateDefinitionRequest)`

SetGates sets Gates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


