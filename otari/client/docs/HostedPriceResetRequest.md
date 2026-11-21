# HostedPriceResetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** |  | [optional] [default to false]
**Models** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHostedPriceResetRequest

`func NewHostedPriceResetRequest() *HostedPriceResetRequest`

NewHostedPriceResetRequest instantiates a new HostedPriceResetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedPriceResetRequestWithDefaults

`func NewHostedPriceResetRequestWithDefaults() *HostedPriceResetRequest`

NewHostedPriceResetRequestWithDefaults instantiates a new HostedPriceResetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *HostedPriceResetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *HostedPriceResetRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *HostedPriceResetRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *HostedPriceResetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetModels

`func (o *HostedPriceResetRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *HostedPriceResetRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *HostedPriceResetRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *HostedPriceResetRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *HostedPriceResetRequest) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *HostedPriceResetRequest) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetVendor

`func (o *HostedPriceResetRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostedPriceResetRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostedPriceResetRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HostedPriceResetRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *HostedPriceResetRequest) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *HostedPriceResetRequest) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


