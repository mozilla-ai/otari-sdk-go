# PlaygroundToolStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | **bool** | Whether this deployment has a backend for the tool at all. | 
**Enabled** | **bool** | Whether the caller&#39;s workspace may attach it. | 
**Reason** | Pointer to **NullableString** | Why it cannot be attached. Null when it can. | [optional] 

## Methods

### NewPlaygroundToolStatus

`func NewPlaygroundToolStatus(configured bool, enabled bool, ) *PlaygroundToolStatus`

NewPlaygroundToolStatus instantiates a new PlaygroundToolStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundToolStatusWithDefaults

`func NewPlaygroundToolStatusWithDefaults() *PlaygroundToolStatus`

NewPlaygroundToolStatusWithDefaults instantiates a new PlaygroundToolStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *PlaygroundToolStatus) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *PlaygroundToolStatus) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *PlaygroundToolStatus) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.


### GetEnabled

`func (o *PlaygroundToolStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlaygroundToolStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlaygroundToolStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetReason

`func (o *PlaygroundToolStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlaygroundToolStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlaygroundToolStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlaygroundToolStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *PlaygroundToolStatus) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PlaygroundToolStatus) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


