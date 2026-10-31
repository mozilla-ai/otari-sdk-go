# SandboxCredentialPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvVar** | **string** | The environment variable a deployment that stores no key is read from. | 
**Last4** | Pointer to **NullableString** | The key&#39;s last four characters, to tell keys apart. | [optional] 
**Source** | **NullableString** | &#39;stored&#39; for a key set here, which wins; &#39;environment&#39; for the provider&#39;s own environment variable; null when there is neither and code execution fails until one is set. | 

## Methods

### NewSandboxCredentialPublic

`func NewSandboxCredentialPublic(envVar string, source NullableString, ) *SandboxCredentialPublic`

NewSandboxCredentialPublic instantiates a new SandboxCredentialPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCredentialPublicWithDefaults

`func NewSandboxCredentialPublicWithDefaults() *SandboxCredentialPublic`

NewSandboxCredentialPublicWithDefaults instantiates a new SandboxCredentialPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvVar

`func (o *SandboxCredentialPublic) GetEnvVar() string`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *SandboxCredentialPublic) GetEnvVarOk() (*string, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *SandboxCredentialPublic) SetEnvVar(v string)`

SetEnvVar sets EnvVar field to given value.


### GetLast4

`func (o *SandboxCredentialPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *SandboxCredentialPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *SandboxCredentialPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *SandboxCredentialPublic) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *SandboxCredentialPublic) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *SandboxCredentialPublic) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetSource

`func (o *SandboxCredentialPublic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SandboxCredentialPublic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SandboxCredentialPublic) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *SandboxCredentialPublic) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SandboxCredentialPublic) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


