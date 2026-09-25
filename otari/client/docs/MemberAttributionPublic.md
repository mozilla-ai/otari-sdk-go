# MemberAttributionPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** |  | [optional] 
**Blocked** | **bool** |  | 
**Reserved** | **float32** |  | 
**Spend** | **float32** |  | 

## Methods

### NewMemberAttributionPublic

`func NewMemberAttributionPublic(blocked bool, reserved float32, spend float32, ) *MemberAttributionPublic`

NewMemberAttributionPublic instantiates a new MemberAttributionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAttributionPublicWithDefaults

`func NewMemberAttributionPublicWithDefaults() *MemberAttributionPublic`

NewMemberAttributionPublicWithDefaults instantiates a new MemberAttributionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *MemberAttributionPublic) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *MemberAttributionPublic) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *MemberAttributionPublic) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *MemberAttributionPublic) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### SetAllowedModelsNil

`func (o *MemberAttributionPublic) SetAllowedModelsNil(b bool)`

 SetAllowedModelsNil sets the value for AllowedModels to be an explicit nil

### UnsetAllowedModels
`func (o *MemberAttributionPublic) UnsetAllowedModels()`

UnsetAllowedModels ensures that no value is present for AllowedModels, not even an explicit nil
### GetBlocked

`func (o *MemberAttributionPublic) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *MemberAttributionPublic) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *MemberAttributionPublic) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetReserved

`func (o *MemberAttributionPublic) GetReserved() float32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *MemberAttributionPublic) GetReservedOk() (*float32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *MemberAttributionPublic) SetReserved(v float32)`

SetReserved sets Reserved field to given value.


### GetSpend

`func (o *MemberAttributionPublic) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *MemberAttributionPublic) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *MemberAttributionPublic) SetSpend(v float32)`

SetSpend sets Spend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


