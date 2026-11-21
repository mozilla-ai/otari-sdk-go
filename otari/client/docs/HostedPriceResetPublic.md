# HostedPriceResetPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**Changes** | [**[]HostedPriceResetChangePublic**](HostedPriceResetChangePublic.md) |  | 
**SkippedNoDefault** | **[]string** |  | 

## Methods

### NewHostedPriceResetPublic

`func NewHostedPriceResetPublic(applied bool, changes []HostedPriceResetChangePublic, skippedNoDefault []string, ) *HostedPriceResetPublic`

NewHostedPriceResetPublic instantiates a new HostedPriceResetPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedPriceResetPublicWithDefaults

`func NewHostedPriceResetPublicWithDefaults() *HostedPriceResetPublic`

NewHostedPriceResetPublicWithDefaults instantiates a new HostedPriceResetPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *HostedPriceResetPublic) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *HostedPriceResetPublic) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *HostedPriceResetPublic) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetChanges

`func (o *HostedPriceResetPublic) GetChanges() []HostedPriceResetChangePublic`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HostedPriceResetPublic) GetChangesOk() (*[]HostedPriceResetChangePublic, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HostedPriceResetPublic) SetChanges(v []HostedPriceResetChangePublic)`

SetChanges sets Changes field to given value.


### GetSkippedNoDefault

`func (o *HostedPriceResetPublic) GetSkippedNoDefault() []string`

GetSkippedNoDefault returns the SkippedNoDefault field if non-nil, zero value otherwise.

### GetSkippedNoDefaultOk

`func (o *HostedPriceResetPublic) GetSkippedNoDefaultOk() (*[]string, bool)`

GetSkippedNoDefaultOk returns a tuple with the SkippedNoDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedNoDefault

`func (o *HostedPriceResetPublic) SetSkippedNoDefault(v []string)`

SetSkippedNoDefault sets SkippedNoDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


