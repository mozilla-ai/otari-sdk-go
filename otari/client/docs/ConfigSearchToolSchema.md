# ConfigSearchToolSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**HasApiKey** | **bool** | Whether the config entry carries an API key. The key itself is not shown. | 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**Shadowed** | Pointer to **bool** | True when a stored search tool of the same name overrides this entry. | [optional] [default to false]

## Methods

### NewConfigSearchToolSchema

`func NewConfigSearchToolSchema(hasApiKey bool, name string, provider string, ) *ConfigSearchToolSchema`

NewConfigSearchToolSchema instantiates a new ConfigSearchToolSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSearchToolSchemaWithDefaults

`func NewConfigSearchToolSchemaWithDefaults() *ConfigSearchToolSchema`

NewConfigSearchToolSchemaWithDefaults instantiates a new ConfigSearchToolSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ConfigSearchToolSchema) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ConfigSearchToolSchema) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ConfigSearchToolSchema) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *ConfigSearchToolSchema) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *ConfigSearchToolSchema) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *ConfigSearchToolSchema) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetHasApiKey

`func (o *ConfigSearchToolSchema) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *ConfigSearchToolSchema) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *ConfigSearchToolSchema) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.


### GetName

`func (o *ConfigSearchToolSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigSearchToolSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigSearchToolSchema) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *ConfigSearchToolSchema) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConfigSearchToolSchema) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConfigSearchToolSchema) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetShadowed

`func (o *ConfigSearchToolSchema) GetShadowed() bool`

GetShadowed returns the Shadowed field if non-nil, zero value otherwise.

### GetShadowedOk

`func (o *ConfigSearchToolSchema) GetShadowedOk() (*bool, bool)`

GetShadowedOk returns a tuple with the Shadowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowed

`func (o *ConfigSearchToolSchema) SetShadowed(v bool)`

SetShadowed sets Shadowed field to given value.

### HasShadowed

`func (o *ConfigSearchToolSchema) HasShadowed() bool`

HasShadowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


