# HostedVendorModelsSwitchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**OnlyUnpriced** | Pointer to **bool** |  | [optional] [default to false]
**Vendor** | **string** |  | 

## Methods

### NewHostedVendorModelsSwitchRequest

`func NewHostedVendorModelsSwitchRequest(enabled bool, vendor string, ) *HostedVendorModelsSwitchRequest`

NewHostedVendorModelsSwitchRequest instantiates a new HostedVendorModelsSwitchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedVendorModelsSwitchRequestWithDefaults

`func NewHostedVendorModelsSwitchRequestWithDefaults() *HostedVendorModelsSwitchRequest`

NewHostedVendorModelsSwitchRequestWithDefaults instantiates a new HostedVendorModelsSwitchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HostedVendorModelsSwitchRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostedVendorModelsSwitchRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostedVendorModelsSwitchRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOnlyUnpriced

`func (o *HostedVendorModelsSwitchRequest) GetOnlyUnpriced() bool`

GetOnlyUnpriced returns the OnlyUnpriced field if non-nil, zero value otherwise.

### GetOnlyUnpricedOk

`func (o *HostedVendorModelsSwitchRequest) GetOnlyUnpricedOk() (*bool, bool)`

GetOnlyUnpricedOk returns a tuple with the OnlyUnpriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyUnpriced

`func (o *HostedVendorModelsSwitchRequest) SetOnlyUnpriced(v bool)`

SetOnlyUnpriced sets OnlyUnpriced field to given value.

### HasOnlyUnpriced

`func (o *HostedVendorModelsSwitchRequest) HasOnlyUnpriced() bool`

HasOnlyUnpriced returns a boolean if a field has been set.

### GetVendor

`func (o *HostedVendorModelsSwitchRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostedVendorModelsSwitchRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostedVendorModelsSwitchRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


