# ToolSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]ToolSettingField**](ToolSettingField.md) |  | 
**SandboxProvider** | Pointer to [**NullableSandboxProvider**](SandboxProvider.md) | What runs generated code: &#39;protocol&#39; (a sandbox at sandbox_url) or &#39;e2b&#39;. Set at startup, not editable here. Null for a reader who does not operate the deployment. | [optional] 

## Methods

### NewToolSettingsResponse

`func NewToolSettingsResponse(fields []ToolSettingField, ) *ToolSettingsResponse`

NewToolSettingsResponse instantiates a new ToolSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSettingsResponseWithDefaults

`func NewToolSettingsResponseWithDefaults() *ToolSettingsResponse`

NewToolSettingsResponseWithDefaults instantiates a new ToolSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ToolSettingsResponse) GetFields() []ToolSettingField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ToolSettingsResponse) GetFieldsOk() (*[]ToolSettingField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ToolSettingsResponse) SetFields(v []ToolSettingField)`

SetFields sets Fields field to given value.


### GetSandboxProvider

`func (o *ToolSettingsResponse) GetSandboxProvider() SandboxProvider`

GetSandboxProvider returns the SandboxProvider field if non-nil, zero value otherwise.

### GetSandboxProviderOk

`func (o *ToolSettingsResponse) GetSandboxProviderOk() (*SandboxProvider, bool)`

GetSandboxProviderOk returns a tuple with the SandboxProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxProvider

`func (o *ToolSettingsResponse) SetSandboxProvider(v SandboxProvider)`

SetSandboxProvider sets SandboxProvider field to given value.

### HasSandboxProvider

`func (o *ToolSettingsResponse) HasSandboxProvider() bool`

HasSandboxProvider returns a boolean if a field has been set.

### SetSandboxProviderNil

`func (o *ToolSettingsResponse) SetSandboxProviderNil(b bool)`

 SetSandboxProviderNil sets the value for SandboxProvider to be an explicit nil

### UnsetSandboxProvider
`func (o *ToolSettingsResponse) UnsetSandboxProvider()`

UnsetSandboxProvider ensures that no value is present for SandboxProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


