# ToolSettingField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** |  | 
**Service** | **string** |  | 
**Type** | **string** |  | 
**Value** | [**NullableValue1**](Value1.md) |  | 

## Methods

### NewToolSettingField

`func NewToolSettingField(key string, service string, type_ string, value NullableValue1, ) *ToolSettingField`

NewToolSettingField instantiates a new ToolSettingField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSettingFieldWithDefaults

`func NewToolSettingFieldWithDefaults() *ToolSettingField`

NewToolSettingFieldWithDefaults instantiates a new ToolSettingField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ToolSettingField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolSettingField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolSettingField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolSettingField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolSettingField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolSettingField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *ToolSettingField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ToolSettingField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ToolSettingField) SetKey(v string)`

SetKey sets Key field to given value.


### GetService

`func (o *ToolSettingField) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ToolSettingField) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ToolSettingField) SetService(v string)`

SetService sets Service field to given value.


### GetType

`func (o *ToolSettingField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolSettingField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolSettingField) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ToolSettingField) GetValue() Value1`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ToolSettingField) GetValueOk() (*Value1, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ToolSettingField) SetValue(v Value1)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ToolSettingField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ToolSettingField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


