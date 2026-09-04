# MRBetaFallbackBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**MRBetaFallbackInfo**](MRBetaFallbackInfo.md) |  | 
**To** | [**MRBetaFallbackInfo**](MRBetaFallbackInfo.md) |  | 
**Trigger** | [**MRBetaFallbackRefusalTrigger**](MRBetaFallbackRefusalTrigger.md) |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaFallbackBlock

`func NewMRBetaFallbackBlock(from MRBetaFallbackInfo, to MRBetaFallbackInfo, trigger MRBetaFallbackRefusalTrigger, type_ string, ) *MRBetaFallbackBlock`

NewMRBetaFallbackBlock instantiates a new MRBetaFallbackBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaFallbackBlockWithDefaults

`func NewMRBetaFallbackBlockWithDefaults() *MRBetaFallbackBlock`

NewMRBetaFallbackBlockWithDefaults instantiates a new MRBetaFallbackBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MRBetaFallbackBlock) GetFrom() MRBetaFallbackInfo`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MRBetaFallbackBlock) GetFromOk() (*MRBetaFallbackInfo, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MRBetaFallbackBlock) SetFrom(v MRBetaFallbackInfo)`

SetFrom sets From field to given value.


### GetTo

`func (o *MRBetaFallbackBlock) GetTo() MRBetaFallbackInfo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MRBetaFallbackBlock) GetToOk() (*MRBetaFallbackInfo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MRBetaFallbackBlock) SetTo(v MRBetaFallbackInfo)`

SetTo sets To field to given value.


### GetTrigger

`func (o *MRBetaFallbackBlock) GetTrigger() MRBetaFallbackRefusalTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *MRBetaFallbackBlock) GetTriggerOk() (*MRBetaFallbackRefusalTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *MRBetaFallbackBlock) SetTrigger(v MRBetaFallbackRefusalTrigger)`

SetTrigger sets Trigger field to given value.


### GetType

`func (o *MRBetaFallbackBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaFallbackBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaFallbackBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


