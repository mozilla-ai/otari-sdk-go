# MRBetaToolTextEditor20250429

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**InputExamples** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMRBetaToolTextEditor20250429

`func NewMRBetaToolTextEditor20250429(name string, type_ string, ) *MRBetaToolTextEditor20250429`

NewMRBetaToolTextEditor20250429 instantiates a new MRBetaToolTextEditor20250429 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaToolTextEditor20250429WithDefaults

`func NewMRBetaToolTextEditor20250429WithDefaults() *MRBetaToolTextEditor20250429`

NewMRBetaToolTextEditor20250429WithDefaults instantiates a new MRBetaToolTextEditor20250429 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MRBetaToolTextEditor20250429) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaToolTextEditor20250429) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaToolTextEditor20250429) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaToolTextEditor20250429) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaToolTextEditor20250429) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaToolTextEditor20250429) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaToolTextEditor20250429) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaToolTextEditor20250429) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaToolTextEditor20250429) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaToolTextEditor20250429) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaToolTextEditor20250429) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaToolTextEditor20250429) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetCacheControl

`func (o *MRBetaToolTextEditor20250429) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaToolTextEditor20250429) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaToolTextEditor20250429) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaToolTextEditor20250429) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaToolTextEditor20250429) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaToolTextEditor20250429) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaToolTextEditor20250429) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaToolTextEditor20250429) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaToolTextEditor20250429) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaToolTextEditor20250429) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaToolTextEditor20250429) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaToolTextEditor20250429) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetInputExamples

`func (o *MRBetaToolTextEditor20250429) GetInputExamples() []map[string]interface{}`

GetInputExamples returns the InputExamples field if non-nil, zero value otherwise.

### GetInputExamplesOk

`func (o *MRBetaToolTextEditor20250429) GetInputExamplesOk() (*[]map[string]interface{}, bool)`

GetInputExamplesOk returns a tuple with the InputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputExamples

`func (o *MRBetaToolTextEditor20250429) SetInputExamples(v []map[string]interface{})`

SetInputExamples sets InputExamples field to given value.

### HasInputExamples

`func (o *MRBetaToolTextEditor20250429) HasInputExamples() bool`

HasInputExamples returns a boolean if a field has been set.

### SetInputExamplesNil

`func (o *MRBetaToolTextEditor20250429) SetInputExamplesNil(b bool)`

 SetInputExamplesNil sets the value for InputExamples to be an explicit nil

### UnsetInputExamples
`func (o *MRBetaToolTextEditor20250429) UnsetInputExamples()`

UnsetInputExamples ensures that no value is present for InputExamples, not even an explicit nil
### GetStrict

`func (o *MRBetaToolTextEditor20250429) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaToolTextEditor20250429) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaToolTextEditor20250429) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaToolTextEditor20250429) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaToolTextEditor20250429) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaToolTextEditor20250429) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


