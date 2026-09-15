# SearchProviderSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultApiBase** | Pointer to **NullableString** | The endpoint a tool on this provider uses when it declares no api_base. Null means nothing supplies one, so an api_base is required. | [optional] 
**Id** | **string** | Value to send as &#39;provider&#39;. | 
**RequiresApiBase** | **bool** | True when this provider has no endpoint of its own, so the tool must say where the backend is. | 
**RequiresApiKey** | **bool** | True when a tool on this provider must carry an API key. | 

## Methods

### NewSearchProviderSchema

`func NewSearchProviderSchema(id string, requiresApiBase bool, requiresApiKey bool, ) *SearchProviderSchema`

NewSearchProviderSchema instantiates a new SearchProviderSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProviderSchemaWithDefaults

`func NewSearchProviderSchemaWithDefaults() *SearchProviderSchema`

NewSearchProviderSchemaWithDefaults instantiates a new SearchProviderSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultApiBase

`func (o *SearchProviderSchema) GetDefaultApiBase() string`

GetDefaultApiBase returns the DefaultApiBase field if non-nil, zero value otherwise.

### GetDefaultApiBaseOk

`func (o *SearchProviderSchema) GetDefaultApiBaseOk() (*string, bool)`

GetDefaultApiBaseOk returns a tuple with the DefaultApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApiBase

`func (o *SearchProviderSchema) SetDefaultApiBase(v string)`

SetDefaultApiBase sets DefaultApiBase field to given value.

### HasDefaultApiBase

`func (o *SearchProviderSchema) HasDefaultApiBase() bool`

HasDefaultApiBase returns a boolean if a field has been set.

### SetDefaultApiBaseNil

`func (o *SearchProviderSchema) SetDefaultApiBaseNil(b bool)`

 SetDefaultApiBaseNil sets the value for DefaultApiBase to be an explicit nil

### UnsetDefaultApiBase
`func (o *SearchProviderSchema) UnsetDefaultApiBase()`

UnsetDefaultApiBase ensures that no value is present for DefaultApiBase, not even an explicit nil
### GetId

`func (o *SearchProviderSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchProviderSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchProviderSchema) SetId(v string)`

SetId sets Id field to given value.


### GetRequiresApiBase

`func (o *SearchProviderSchema) GetRequiresApiBase() bool`

GetRequiresApiBase returns the RequiresApiBase field if non-nil, zero value otherwise.

### GetRequiresApiBaseOk

`func (o *SearchProviderSchema) GetRequiresApiBaseOk() (*bool, bool)`

GetRequiresApiBaseOk returns a tuple with the RequiresApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiBase

`func (o *SearchProviderSchema) SetRequiresApiBase(v bool)`

SetRequiresApiBase sets RequiresApiBase field to given value.


### GetRequiresApiKey

`func (o *SearchProviderSchema) GetRequiresApiKey() bool`

GetRequiresApiKey returns the RequiresApiKey field if non-nil, zero value otherwise.

### GetRequiresApiKeyOk

`func (o *SearchProviderSchema) GetRequiresApiKeyOk() (*bool, bool)`

GetRequiresApiKeyOk returns a tuple with the RequiresApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresApiKey

`func (o *SearchProviderSchema) SetRequiresApiKey(v bool)`

SetRequiresApiKey sets RequiresApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


