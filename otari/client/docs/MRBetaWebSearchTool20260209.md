# MRBetaWebSearchTool20260209

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**AllowedCallers** | Pointer to **[]string** |  | [optional] 
**AllowedDomains** | Pointer to **[]string** |  | [optional] 
**BlockedDomains** | Pointer to **[]string** |  | [optional] 
**CacheControl** | Pointer to [**NullableMRBetaCacheControlEphemeral**](MRBetaCacheControlEphemeral.md) |  | [optional] 
**DeferLoading** | Pointer to **NullableBool** |  | [optional] 
**MaxUses** | Pointer to **NullableInt32** |  | [optional] 
**Strict** | Pointer to **NullableBool** |  | [optional] 
**UserLocation** | Pointer to [**NullableMRBetaUserLocation**](MRBetaUserLocation.md) |  | [optional] 

## Methods

### NewMRBetaWebSearchTool20260209

`func NewMRBetaWebSearchTool20260209(name string, type_ string, ) *MRBetaWebSearchTool20260209`

NewMRBetaWebSearchTool20260209 instantiates a new MRBetaWebSearchTool20260209 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaWebSearchTool20260209WithDefaults

`func NewMRBetaWebSearchTool20260209WithDefaults() *MRBetaWebSearchTool20260209`

NewMRBetaWebSearchTool20260209WithDefaults instantiates a new MRBetaWebSearchTool20260209 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MRBetaWebSearchTool20260209) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MRBetaWebSearchTool20260209) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MRBetaWebSearchTool20260209) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MRBetaWebSearchTool20260209) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaWebSearchTool20260209) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaWebSearchTool20260209) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedCallers

`func (o *MRBetaWebSearchTool20260209) GetAllowedCallers() []string`

GetAllowedCallers returns the AllowedCallers field if non-nil, zero value otherwise.

### GetAllowedCallersOk

`func (o *MRBetaWebSearchTool20260209) GetAllowedCallersOk() (*[]string, bool)`

GetAllowedCallersOk returns a tuple with the AllowedCallers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCallers

`func (o *MRBetaWebSearchTool20260209) SetAllowedCallers(v []string)`

SetAllowedCallers sets AllowedCallers field to given value.

### HasAllowedCallers

`func (o *MRBetaWebSearchTool20260209) HasAllowedCallers() bool`

HasAllowedCallers returns a boolean if a field has been set.

### SetAllowedCallersNil

`func (o *MRBetaWebSearchTool20260209) SetAllowedCallersNil(b bool)`

 SetAllowedCallersNil sets the value for AllowedCallers to be an explicit nil

### UnsetAllowedCallers
`func (o *MRBetaWebSearchTool20260209) UnsetAllowedCallers()`

UnsetAllowedCallers ensures that no value is present for AllowedCallers, not even an explicit nil
### GetAllowedDomains

`func (o *MRBetaWebSearchTool20260209) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *MRBetaWebSearchTool20260209) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *MRBetaWebSearchTool20260209) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *MRBetaWebSearchTool20260209) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### SetAllowedDomainsNil

`func (o *MRBetaWebSearchTool20260209) SetAllowedDomainsNil(b bool)`

 SetAllowedDomainsNil sets the value for AllowedDomains to be an explicit nil

### UnsetAllowedDomains
`func (o *MRBetaWebSearchTool20260209) UnsetAllowedDomains()`

UnsetAllowedDomains ensures that no value is present for AllowedDomains, not even an explicit nil
### GetBlockedDomains

`func (o *MRBetaWebSearchTool20260209) GetBlockedDomains() []string`

GetBlockedDomains returns the BlockedDomains field if non-nil, zero value otherwise.

### GetBlockedDomainsOk

`func (o *MRBetaWebSearchTool20260209) GetBlockedDomainsOk() (*[]string, bool)`

GetBlockedDomainsOk returns a tuple with the BlockedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedDomains

`func (o *MRBetaWebSearchTool20260209) SetBlockedDomains(v []string)`

SetBlockedDomains sets BlockedDomains field to given value.

### HasBlockedDomains

`func (o *MRBetaWebSearchTool20260209) HasBlockedDomains() bool`

HasBlockedDomains returns a boolean if a field has been set.

### SetBlockedDomainsNil

`func (o *MRBetaWebSearchTool20260209) SetBlockedDomainsNil(b bool)`

 SetBlockedDomainsNil sets the value for BlockedDomains to be an explicit nil

### UnsetBlockedDomains
`func (o *MRBetaWebSearchTool20260209) UnsetBlockedDomains()`

UnsetBlockedDomains ensures that no value is present for BlockedDomains, not even an explicit nil
### GetCacheControl

`func (o *MRBetaWebSearchTool20260209) GetCacheControl() MRBetaCacheControlEphemeral`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *MRBetaWebSearchTool20260209) GetCacheControlOk() (*MRBetaCacheControlEphemeral, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *MRBetaWebSearchTool20260209) SetCacheControl(v MRBetaCacheControlEphemeral)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *MRBetaWebSearchTool20260209) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### SetCacheControlNil

`func (o *MRBetaWebSearchTool20260209) SetCacheControlNil(b bool)`

 SetCacheControlNil sets the value for CacheControl to be an explicit nil

### UnsetCacheControl
`func (o *MRBetaWebSearchTool20260209) UnsetCacheControl()`

UnsetCacheControl ensures that no value is present for CacheControl, not even an explicit nil
### GetDeferLoading

`func (o *MRBetaWebSearchTool20260209) GetDeferLoading() bool`

GetDeferLoading returns the DeferLoading field if non-nil, zero value otherwise.

### GetDeferLoadingOk

`func (o *MRBetaWebSearchTool20260209) GetDeferLoadingOk() (*bool, bool)`

GetDeferLoadingOk returns a tuple with the DeferLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferLoading

`func (o *MRBetaWebSearchTool20260209) SetDeferLoading(v bool)`

SetDeferLoading sets DeferLoading field to given value.

### HasDeferLoading

`func (o *MRBetaWebSearchTool20260209) HasDeferLoading() bool`

HasDeferLoading returns a boolean if a field has been set.

### SetDeferLoadingNil

`func (o *MRBetaWebSearchTool20260209) SetDeferLoadingNil(b bool)`

 SetDeferLoadingNil sets the value for DeferLoading to be an explicit nil

### UnsetDeferLoading
`func (o *MRBetaWebSearchTool20260209) UnsetDeferLoading()`

UnsetDeferLoading ensures that no value is present for DeferLoading, not even an explicit nil
### GetMaxUses

`func (o *MRBetaWebSearchTool20260209) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *MRBetaWebSearchTool20260209) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *MRBetaWebSearchTool20260209) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *MRBetaWebSearchTool20260209) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### SetMaxUsesNil

`func (o *MRBetaWebSearchTool20260209) SetMaxUsesNil(b bool)`

 SetMaxUsesNil sets the value for MaxUses to be an explicit nil

### UnsetMaxUses
`func (o *MRBetaWebSearchTool20260209) UnsetMaxUses()`

UnsetMaxUses ensures that no value is present for MaxUses, not even an explicit nil
### GetStrict

`func (o *MRBetaWebSearchTool20260209) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *MRBetaWebSearchTool20260209) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *MRBetaWebSearchTool20260209) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *MRBetaWebSearchTool20260209) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### SetStrictNil

`func (o *MRBetaWebSearchTool20260209) SetStrictNil(b bool)`

 SetStrictNil sets the value for Strict to be an explicit nil

### UnsetStrict
`func (o *MRBetaWebSearchTool20260209) UnsetStrict()`

UnsetStrict ensures that no value is present for Strict, not even an explicit nil
### GetUserLocation

`func (o *MRBetaWebSearchTool20260209) GetUserLocation() MRBetaUserLocation`

GetUserLocation returns the UserLocation field if non-nil, zero value otherwise.

### GetUserLocationOk

`func (o *MRBetaWebSearchTool20260209) GetUserLocationOk() (*MRBetaUserLocation, bool)`

GetUserLocationOk returns a tuple with the UserLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocation

`func (o *MRBetaWebSearchTool20260209) SetUserLocation(v MRBetaUserLocation)`

SetUserLocation sets UserLocation field to given value.

### HasUserLocation

`func (o *MRBetaWebSearchTool20260209) HasUserLocation() bool`

HasUserLocation returns a boolean if a field has been set.

### SetUserLocationNil

`func (o *MRBetaWebSearchTool20260209) SetUserLocationNil(b bool)`

 SetUserLocationNil sets the value for UserLocation to be an explicit nil

### UnsetUserLocation
`func (o *MRBetaWebSearchTool20260209) UnsetUserLocation()`

UnsetUserLocation ensures that no value is present for UserLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


