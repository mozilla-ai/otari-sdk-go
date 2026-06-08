# GuardrailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] [default to "monitor"]
**On** | Pointer to **[]string** |  | [optional] 
**Profile** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**ValidateKwargs** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGuardrailConfig

`func NewGuardrailConfig(profile string, ) *GuardrailConfig`

NewGuardrailConfig instantiates a new GuardrailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailConfigWithDefaults

`func NewGuardrailConfigWithDefaults() *GuardrailConfig`

NewGuardrailConfigWithDefaults instantiates a new GuardrailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GuardrailConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GuardrailConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GuardrailConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GuardrailConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOn

`func (o *GuardrailConfig) GetOn() []string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *GuardrailConfig) GetOnOk() (*[]string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *GuardrailConfig) SetOn(v []string)`

SetOn sets On field to given value.

### HasOn

`func (o *GuardrailConfig) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetProfile

`func (o *GuardrailConfig) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GuardrailConfig) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GuardrailConfig) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetUrl

`func (o *GuardrailConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GuardrailConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GuardrailConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GuardrailConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GuardrailConfig) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GuardrailConfig) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidateKwargs

`func (o *GuardrailConfig) GetValidateKwargs() map[string]interface{}`

GetValidateKwargs returns the ValidateKwargs field if non-nil, zero value otherwise.

### GetValidateKwargsOk

`func (o *GuardrailConfig) GetValidateKwargsOk() (*map[string]interface{}, bool)`

GetValidateKwargsOk returns a tuple with the ValidateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateKwargs

`func (o *GuardrailConfig) SetValidateKwargs(v map[string]interface{})`

SetValidateKwargs sets ValidateKwargs field to given value.

### HasValidateKwargs

`func (o *GuardrailConfig) HasValidateKwargs() bool`

HasValidateKwargs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


