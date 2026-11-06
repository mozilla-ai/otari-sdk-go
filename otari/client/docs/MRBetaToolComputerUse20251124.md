# MRBetaToolComputerUse20251124

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayHeightPx** | **int32** |  | 
**DisplayWidthPx** | **int32** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**DisplayNumber** | Pointer to **NullableInt32** |  | [optional] 
**EnableZoom** | Pointer to **NullableBool** | Only end users that are, or are not, blocked. | [optional] 
**InputExamples** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMRBetaToolComputerUse20251124

`func NewMRBetaToolComputerUse20251124(displayHeightPx int32, displayWidthPx int32, name string, type_ string, ) *MRBetaToolComputerUse20251124`

NewMRBetaToolComputerUse20251124 instantiates a new MRBetaToolComputerUse20251124 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaToolComputerUse20251124WithDefaults

`func NewMRBetaToolComputerUse20251124WithDefaults() *MRBetaToolComputerUse20251124`

NewMRBetaToolComputerUse20251124WithDefaults instantiates a new MRBetaToolComputerUse20251124 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayHeightPx

`func (o *MRBetaToolComputerUse20251124) GetDisplayHeightPx() int32`

GetDisplayHeightPx returns the DisplayHeightPx field if non-nil, zero value otherwise.

### GetDisplayHeightPxOk

`func (o *MRBetaToolComputerUse20251124) GetDisplayHeightPxOk() (*int32, bool)`

GetDisplayHeightPxOk returns a tuple with the DisplayHeightPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayHeightPx

`func (o *MRBetaToolComputerUse20251124) SetDisplayHeightPx(v int32)`

SetDisplayHeightPx sets DisplayHeightPx field to given value.


### GetDisplayWidthPx

`func (o *MRBetaToolComputerUse20251124) GetDisplayWidthPx() int32`

GetDisplayWidthPx returns the DisplayWidthPx field if non-nil, zero value otherwise.

### GetDisplayWidthPxOk

`func (o *MRBetaToolComputerUse20251124) GetDisplayWidthPxOk() (*int32, bool)`

GetDisplayWidthPxOk returns a tuple with the DisplayWidthPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayWidthPx

`func (o *MRBetaToolComputerUse20251124) SetDisplayWidthPx(v int32)`

SetDisplayWidthPx sets DisplayWidthPx field to given value.


### GetName

`func (o *MRBetaToolComputerUse20251124) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaToolComputerUse20251124) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaToolComputerUse20251124) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaToolComputerUse20251124) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaToolComputerUse20251124) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaToolComputerUse20251124) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaToolComputerUse20251124) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaToolComputerUse20251124) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaToolComputerUse20251124) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaToolComputerUse20251124) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaToolComputerUse20251124) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaToolComputerUse20251124) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetCacheControl

`func (o *MRBetaToolComputerUse20251124) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaToolComputerUse20251124) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaToolComputerUse20251124) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaToolComputerUse20251124) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaToolComputerUse20251124) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaToolComputerUse20251124) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaToolComputerUse20251124) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaToolComputerUse20251124) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaToolComputerUse20251124) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaToolComputerUse20251124) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaToolComputerUse20251124) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaToolComputerUse20251124) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetDisplayNumber

`func (o *MRBetaToolComputerUse20251124) GetDisplayNumber() int32`

GetDisplayNumber returns the DisplayNumber field if non-nil, zero value otherwise.

### GetDisplayNumberOk

`func (o *MRBetaToolComputerUse20251124) GetDisplayNumberOk() (*int32, bool)`

GetDisplayNumberOk returns a tuple with the DisplayNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNumber

`func (o *MRBetaToolComputerUse20251124) SetDisplayNumber(v int32)`

SetDisplayNumber sets DisplayNumber field to given value.

### HasDisplayNumber

`func (o *MRBetaToolComputerUse20251124) HasDisplayNumber() bool`

HasDisplayNumber returns a boolean if a field has been set.

### SetDisplayNumberNil

`func (o *MRBetaToolComputerUse20251124) SetDisplayNumberNil(b bool)`

 SetDisplayNumberNil sets the value for DisplayNumber to be an explicit nil

### UnsetDisplayNumber
`func (o *MRBetaToolComputerUse20251124) UnsetDisplayNumber()`

UnsetDisplayNumber ensures that no value is present for DisplayNumber, not even an explicit nil
### GetEnableZoom

`func (o *MRBetaToolComputerUse20251124) GetEnableZoom() bool`

GetEnableZoom returns the EnableZoom field if non-nil, zero value otherwise.

### GetEnableZoomOk

`func (o *MRBetaToolComputerUse20251124) GetEnableZoomOk() (*bool, bool)`

GetEnableZoomOk returns a tuple with the EnableZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoom

`func (o *MRBetaToolComputerUse20251124) SetEnableZoom(v bool)`

SetEnableZoom sets EnableZoom field to given value.

### HasEnableZoom

`func (o *MRBetaToolComputerUse20251124) HasEnableZoom() bool`

HasEnableZoom returns a boolean if a field has been set.

### SetEnableZoomNil

`func (o *MRBetaToolComputerUse20251124) SetEnableZoomNil(b bool)`

 SetEnableZoomNil sets the value for EnableZoom to be an explicit nil

### UnsetEnableZoom
`func (o *MRBetaToolComputerUse20251124) UnsetEnableZoom()`

UnsetEnableZoom ensures that no value is present for EnableZoom, not even an explicit nil
### GetInputExamples

`func (o *MRBetaToolComputerUse20251124) GetInputExamples() []map[string]interface{}`

GetInputExamples returns the InputExamples field if non-nil, zero value otherwise.

### GetInputExamplesOk

`func (o *MRBetaToolComputerUse20251124) GetInputExamplesOk() (*[]map[string]interface{}, bool)`

GetInputExamplesOk returns a tuple with the InputExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputExamples

`func (o *MRBetaToolComputerUse20251124) SetInputExamples(v []map[string]interface{})`

SetInputExamples sets InputExamples field to given value.

### HasInputExamples

`func (o *MRBetaToolComputerUse20251124) HasInputExamples() bool`

HasInputExamples returns a boolean if a field has been set.

### SetInputExamplesNil

`func (o *MRBetaToolComputerUse20251124) SetInputExamplesNil(b bool)`

 SetInputExamplesNil sets the value for InputExamples to be an explicit nil

### UnsetInputExamples
`func (o *MRBetaToolComputerUse20251124) UnsetInputExamples()`

UnsetInputExamples ensures that no value is present for InputExamples, not even an explicit nil
### GetStrict

`func (o *MRBetaToolComputerUse20251124) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaToolComputerUse20251124) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaToolComputerUse20251124) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaToolComputerUse20251124) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaToolComputerUse20251124) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaToolComputerUse20251124) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


