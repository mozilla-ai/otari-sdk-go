# KnownProviderSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultApiBase** | Pointer to **NullableString** | Built-in endpoint; blank means the SDK&#39;s default. | [optional] 
**EnvKey** | Pointer to **NullableString** | Env var the SDK reads this provider&#39;s key from. | [optional] 
**EnvKeyPresent** | Pointer to **bool** | True when env_key is already set on the server, so a pasted key is optional (env fallback). | [optional] [default to false]
**Id** | **string** | any-llm provider id, used as the default instance name. | 
**Name** | **string** | Human-friendly display name. | 
**RequiresApiKey** | **bool** | False for keyless local backends (Ollama, llama.cpp). | 

## Methods

### NewKnownProviderSchema

`func NewKnownProviderSchema(id string, name string, requiresApiKey bool, ) *KnownProviderSchema`

NewKnownProviderSchema instantiates a new KnownProviderSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownProviderSchemaWithDefaults

`func NewKnownProviderSchemaWithDefaults() *KnownProviderSchema`

NewKnownProviderSchemaWithDefaults instantiates a new KnownProviderSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultApiBase

`func (o *KnownProviderSchema) GetDefaultApiBase() string`

GetDefaultApiBase returns the DefaultApiBase field if non-nil, zero value otherwise.

### GetDefaultApiBaseOk

`func (o *KnownProviderSchema) GetDefaultApiBaseOk() (*string, bool)`

GetDefaultApiBaseOk returns a tuple with the DefaultApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiBase

`func (o *KnownProviderSchema) SetDefaultApiBase(v string)`

SetDefaultApiBase sets DefaultApiBase field to given value.

### HasDefaultApiBase

`func (o *KnownProviderSchema) HasDefaultApiBase() bool`

HasDefaultApiBase returns a boolean if a field has been set.

### SetDefaultApiBaseNil

`func (o *KnownProviderSchema) SetDefaultApiBaseNil(b bool)`

 SetDefaultApiBaseNil sets the value for DefaultApiBase to be an explicit nil

### UnsetDefaultApiBase
`func (o *KnownProviderSchema) UnsetDefaultApiBase()`

UnsetDefaultApiBase ensures that no value is present for DefaultApiBase, not even an explicit nil
### GetEnvKey

`func (o *KnownProviderSchema) GetEnvKey() string`

GetEnvKey returns the EnvKey field if non-nil, zero value otherwise.

### GetEnvKeyOk

`func (o *KnownProviderSchema) GetEnvKeyOk() (*string, bool)`

GetEnvKeyOk returns a tuple with the EnvKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvKey

`func (o *KnownProviderSchema) SetEnvKey(v string)`

SetEnvKey sets EnvKey field to given value.

### HasEnvKey

`func (o *KnownProviderSchema) HasEnvKey() bool`

HasEnvKey returns a boolean if a field has been set.

### SetEnvKeyNil

`func (o *KnownProviderSchema) SetEnvKeyNil(b bool)`

 SetEnvKeyNil sets the value for EnvKey to be an explicit nil

### UnsetEnvKey
`func (o *KnownProviderSchema) UnsetEnvKey()`

UnsetEnvKey ensures that no value is present for EnvKey, not even an explicit nil
### GetEnvKeyPresent

`func (o *KnownProviderSchema) GetEnvKeyPresent() bool`

GetEnvKeyPresent returns the EnvKeyPresent field if non-nil, zero value otherwise.

### GetEnvKeyPresentOk

`func (o *KnownProviderSchema) GetEnvKeyPresentOk() (*bool, bool)`

GetEnvKeyPresentOk returns a tuple with the EnvKeyPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvKeyPresent

`func (o *KnownProviderSchema) SetEnvKeyPresent(v bool)`

SetEnvKeyPresent sets EnvKeyPresent field to given value.

### HasEnvKeyPresent

`func (o *KnownProviderSchema) HasEnvKeyPresent() bool`

HasEnvKeyPresent returns a boolean if a field has been set.

### GetId

`func (o *KnownProviderSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnownProviderSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnownProviderSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KnownProviderSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnownProviderSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnownProviderSchema) SetName(v string)`

SetName sets Name field to given value.


### GetRequiresApiKey

`func (o *KnownProviderSchema) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *KnownProviderSchema) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *KnownProviderSchema) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


