# SearchProviderOptionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Enum** | Pointer to **[]string** | The only values it takes, when it is limited to a list. | [optional] 
**Name** | **string** |  | 
**OperatorOnly** | Pointer to **bool** | True when only an instance or a credential may set it, never a workspace or a request. | [optional] [default to false]
**Type** | **string** |  | 

## Methods

### NewSearchProviderOptionSchema

`func NewSearchProviderOptionSchema(name string, type_ string, ) *SearchProviderOptionSchema`

NewSearchProviderOptionSchema instantiates a new SearchProviderOptionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProviderOptionSchemaWithDefaults

`func NewSearchProviderOptionSchemaWithDefaults() *SearchProviderOptionSchema`

NewSearchProviderOptionSchemaWithDefaults instantiates a new SearchProviderOptionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *SearchProviderOptionSchema) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SearchProviderOptionSchema) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SearchProviderOptionSchema) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SearchProviderOptionSchema) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *SearchProviderOptionSchema) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *SearchProviderOptionSchema) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDescription

`func (o *SearchProviderOptionSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchProviderOptionSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchProviderOptionSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchProviderOptionSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnum

`func (o *SearchProviderOptionSchema) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *SearchProviderOptionSchema) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *SearchProviderOptionSchema) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *SearchProviderOptionSchema) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *SearchProviderOptionSchema) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *SearchProviderOptionSchema) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
### GetName

`func (o *SearchProviderOptionSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchProviderOptionSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchProviderOptionSchema) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorOnly

`func (o *SearchProviderOptionSchema) GetOperatorOnly() bool`

GetOperatorOnly returns the OperatorOnly field if non-nil, zero value otherwise.

### GetOperatorOnlyOk

`func (o *SearchProviderOptionSchema) GetOperatorOnlyOk() (*bool, bool)`

GetOperatorOnlyOk returns a tuple with the OperatorOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorOnly

`func (o *SearchProviderOptionSchema) SetOperatorOnly(v bool)`

SetOperatorOnly sets OperatorOnly field to given value.

### HasOperatorOnly

`func (o *SearchProviderOptionSchema) HasOperatorOnly() bool`

HasOperatorOnly returns a boolean if a field has been set.

### GetType

`func (o *SearchProviderOptionSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchProviderOptionSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchProviderOptionSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


