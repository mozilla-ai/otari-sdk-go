# ConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**ExclusiveMinimum** | Pointer to **NullableFloat32** |  | [optional] 
**Group** | **string** |  | 
**Key** | **string** |  | 
**Minimum** | Pointer to **NullableFloat32** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**Settable** | **bool** |  | 
**Type** | **string** |  | 
**Value** | [**NullableValue**](Value.md) |  | 

## Methods

### NewConfigField

`func NewConfigField(group string, key string, settable bool, type_ string, value NullableValue, ) *ConfigField`

NewConfigField instantiates a new ConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFieldWithDefaults

`func NewConfigFieldWithDefaults() *ConfigField`

NewConfigFieldWithDefaults instantiates a new ConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConfigField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExclusiveMinimum

`func (o *ConfigField) GetExclusiveMinimum() float32`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *ConfigField) GetExclusiveMinimumOk() (*float32, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *ConfigField) SetExclusiveMinimum(v float32)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.

### HasExclusiveMinimum

`func (o *ConfigField) HasExclusiveMinimum() bool`

HasExclusiveMinimum returns a boolean if a field has been set.

### SetExclusiveMinimumNil

`func (o *ConfigField) SetExclusiveMinimumNil(b bool)`

 SetExclusiveMinimumNil sets the value for ExclusiveMinimum to be an explicit nil

### UnsetExclusiveMinimum
`func (o *ConfigField) UnsetExclusiveMinimum()`

UnsetExclusiveMinimum ensures that no value is present for ExclusiveMinimum, not even an explicit nil
### GetGroup

`func (o *ConfigField) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ConfigField) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ConfigField) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetKey

`func (o *ConfigField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigField) SetKey(v string)`

SetKey sets Key field to given value.


### GetMinimum

`func (o *ConfigField) GetMinimum() float32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *ConfigField) GetMinimumOk() (*float32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *ConfigField) SetMinimum(v float32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *ConfigField) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *ConfigField) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *ConfigField) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetOptions

`func (o *ConfigField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigField) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ConfigField) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ConfigField) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetSettable

`func (o *ConfigField) GetSettable() bool`

GetSettable returns the Settable field if non-nil, zero value otherwise.

### GetSettableOk

`func (o *ConfigField) GetSettableOk() (*bool, bool)`

GetSettableOk returns a tuple with the Settable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettable

`func (o *ConfigField) SetSettable(v bool)`

SetSettable sets Settable field to given value.


### GetType

`func (o *ConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigField) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ConfigField) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigField) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigField) SetValue(v Value)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ConfigField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConfigField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


