# MRBetaCodeExecutionTool20260521

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMRBetaCodeExecutionTool20260521

`func NewMRBetaCodeExecutionTool20260521(name string, type_ string, ) *MRBetaCodeExecutionTool20260521`

NewMRBetaCodeExecutionTool20260521 instantiates a new MRBetaCodeExecutionTool20260521 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCodeExecutionTool20260521WithDefaults

`func NewMRBetaCodeExecutionTool20260521WithDefaults() *MRBetaCodeExecutionTool20260521`

NewMRBetaCodeExecutionTool20260521WithDefaults instantiates a new MRBetaCodeExecutionTool20260521 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MRBetaCodeExecutionTool20260521) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaCodeExecutionTool20260521) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaCodeExecutionTool20260521) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaCodeExecutionTool20260521) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCodeExecutionTool20260521) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCodeExecutionTool20260521) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaCodeExecutionTool20260521) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaCodeExecutionTool20260521) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaCodeExecutionTool20260521) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaCodeExecutionTool20260521) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaCodeExecutionTool20260521) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaCodeExecutionTool20260521) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetCacheControl

`func (o *MRBetaCodeExecutionTool20260521) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaCodeExecutionTool20260521) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaCodeExecutionTool20260521) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaCodeExecutionTool20260521) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaCodeExecutionTool20260521) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaCodeExecutionTool20260521) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaCodeExecutionTool20260521) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaCodeExecutionTool20260521) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaCodeExecutionTool20260521) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaCodeExecutionTool20260521) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaCodeExecutionTool20260521) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaCodeExecutionTool20260521) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetStrict

`func (o *MRBetaCodeExecutionTool20260521) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaCodeExecutionTool20260521) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaCodeExecutionTool20260521) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaCodeExecutionTool20260521) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaCodeExecutionTool20260521) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaCodeExecutionTool20260521) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


