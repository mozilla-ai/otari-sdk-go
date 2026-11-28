# MRBetaResponseToolInputSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMRBetaResponseToolInputSchema

`func NewMRBetaResponseToolInputSchema(type_ string, ) *MRBetaResponseToolInputSchema`

NewMRBetaResponseToolInputSchema instantiates a new MRBetaResponseToolInputSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaResponseToolInputSchemaWithDefaults

`func NewMRBetaResponseToolInputSchemaWithDefaults() *MRBetaResponseToolInputSchema`

NewMRBetaResponseToolInputSchemaWithDefaults instantiates a new MRBetaResponseToolInputSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MRBetaResponseToolInputSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaResponseToolInputSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaResponseToolInputSchema) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *MRBetaResponseToolInputSchema) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MRBetaResponseToolInputSchema) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MRBetaResponseToolInputSchema) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MRBetaResponseToolInputSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *MRBetaResponseToolInputSchema) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *MRBetaResponseToolInputSchema) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRequired

`func (o *MRBetaResponseToolInputSchema) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MRBetaResponseToolInputSchema) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MRBetaResponseToolInputSchema) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *MRBetaResponseToolInputSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *MRBetaResponseToolInputSchema) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *MRBetaResponseToolInputSchema) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


