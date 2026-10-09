# RateLimitRuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseSec** | Pointer to **NullableFloat32** | How long a max_concurrent slot is held at most. | [optional] 
**MaxConcurrent** | Pointer to **NullableInt32** | Requests in flight at once. | [optional] 
**Models** | Pointer to **[]string** | The instance:model names a per: model rule limits; null for any other rule. | [optional] 
**Per** | Pointer to **NullableString** | What one count is shared by. | [optional] 
**Rpm** | Pointer to **NullableInt32** | Requests per minute. | [optional] 
**Tpm** | Pointer to **NullableInt32** | Tokens per minute. | [optional] 
**TpmAdmission** | Pointer to **NullableString** | &#39;used&#39; counts only what a request used; &#39;estimate&#39; holds its estimate. | [optional] 

## Methods

### NewRateLimitRuleUpdate

`func NewRateLimitRuleUpdate() *RateLimitRuleUpdate`

NewRateLimitRuleUpdate instantiates a new RateLimitRuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRuleUpdateWithDefaults

`func NewRateLimitRuleUpdateWithDefaults() *RateLimitRuleUpdate`

NewRateLimitRuleUpdateWithDefaults instantiates a new RateLimitRuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseSec

`func (o *RateLimitRuleUpdate) GetLeaseSec() float32`

GetLeaseSec returns the LeaseSec field if non-nil, zero value otherwise.

### GetLeaseSecOk

`func (o *RateLimitRuleUpdate) GetLeaseSecOk() (*float32, bool)`

GetLeaseSecOk returns a tuple with the LeaseSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseSec

`func (o *RateLimitRuleUpdate) SetLeaseSec(v float32)`

SetLeaseSec sets LeaseSec field to given value.

### HasLeaseSec

`func (o *RateLimitRuleUpdate) HasLeaseSec() bool`

HasLeaseSec returns a boolean if a field has been set.

### SetLeaseSecNil

`func (o *RateLimitRuleUpdate) SetLeaseSecNil(b bool)`

 SetLeaseSecNil sets the value for LeaseSec to be an explicit nil

### UnsetLeaseSec
`func (o *RateLimitRuleUpdate) UnsetLeaseSec()`

UnsetLeaseSec ensures that no value is present for LeaseSec, not even an explicit nil
### GetMaxConcurrent

`func (o *RateLimitRuleUpdate) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *RateLimitRuleUpdate) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *RateLimitRuleUpdate) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.

### HasMaxConcurrent

`func (o *RateLimitRuleUpdate) HasMaxConcurrent() bool`

HasMaxConcurrent returns a boolean if a field has been set.

### SetMaxConcurrentNil

`func (o *RateLimitRuleUpdate) SetMaxConcurrentNil(b bool)`

 SetMaxConcurrentNil sets the value for MaxConcurrent to be an explicit nil

### UnsetMaxConcurrent
`func (o *RateLimitRuleUpdate) UnsetMaxConcurrent()`

UnsetMaxConcurrent ensures that no value is present for MaxConcurrent, not even an explicit nil
### GetModels

`func (o *RateLimitRuleUpdate) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *RateLimitRuleUpdate) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *RateLimitRuleUpdate) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *RateLimitRuleUpdate) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *RateLimitRuleUpdate) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *RateLimitRuleUpdate) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetPer

`func (o *RateLimitRuleUpdate) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimitRuleUpdate) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimitRuleUpdate) SetPer(v string)`

SetPer sets Per field to given value.

### HasPer

`func (o *RateLimitRuleUpdate) HasPer() bool`

HasPer returns a boolean if a field has been set.

### SetPerNil

`func (o *RateLimitRuleUpdate) SetPerNil(b bool)`

 SetPerNil sets the value for Per to be an explicit nil

### UnsetPer
`func (o *RateLimitRuleUpdate) UnsetPer()`

UnsetPer ensures that no value is present for Per, not even an explicit nil
### GetRpm

`func (o *RateLimitRuleUpdate) GetRpm() int32`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *RateLimitRuleUpdate) GetRpmOk() (*int32, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *RateLimitRuleUpdate) SetRpm(v int32)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *RateLimitRuleUpdate) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### SetRpmNil

`func (o *RateLimitRuleUpdate) SetRpmNil(b bool)`

 SetRpmNil sets the value for Rpm to be an explicit nil

### UnsetRpm
`func (o *RateLimitRuleUpdate) UnsetRpm()`

UnsetRpm ensures that no value is present for Rpm, not even an explicit nil
### GetTpm

`func (o *RateLimitRuleUpdate) GetTpm() int32`

GetTpm returns the Tpm field if non-nil, zero value otherwise.

### GetTpmOk

`func (o *RateLimitRuleUpdate) GetTpmOk() (*int32, bool)`

GetTpmOk returns a tuple with the Tpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpm

`func (o *RateLimitRuleUpdate) SetTpm(v int32)`

SetTpm sets Tpm field to given value.

### HasTpm

`func (o *RateLimitRuleUpdate) HasTpm() bool`

HasTpm returns a boolean if a field has been set.

### SetTpmNil

`func (o *RateLimitRuleUpdate) SetTpmNil(b bool)`

 SetTpmNil sets the value for Tpm to be an explicit nil

### UnsetTpm
`func (o *RateLimitRuleUpdate) UnsetTpm()`

UnsetTpm ensures that no value is present for Tpm, not even an explicit nil
### GetTpmAdmission

`func (o *RateLimitRuleUpdate) GetTpmAdmission() string`

GetTpmAdmission returns the TpmAdmission field if non-nil, zero value otherwise.

### GetTpmAdmissionOk

`func (o *RateLimitRuleUpdate) GetTpmAdmissionOk() (*string, bool)`

GetTpmAdmissionOk returns a tuple with the TpmAdmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmAdmission

`func (o *RateLimitRuleUpdate) SetTpmAdmission(v string)`

SetTpmAdmission sets TpmAdmission field to given value.

### HasTpmAdmission

`func (o *RateLimitRuleUpdate) HasTpmAdmission() bool`

HasTpmAdmission returns a boolean if a field has been set.

### SetTpmAdmissionNil

`func (o *RateLimitRuleUpdate) SetTpmAdmissionNil(b bool)`

 SetTpmAdmissionNil sets the value for TpmAdmission to be an explicit nil

### UnsetTpmAdmission
`func (o *RateLimitRuleUpdate) UnsetTpmAdmission()`

UnsetTpmAdmission ensures that no value is present for TpmAdmission, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


