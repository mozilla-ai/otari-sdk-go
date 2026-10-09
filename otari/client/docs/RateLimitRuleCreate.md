# RateLimitRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseSec** | Pointer to **float32** | How long a max_concurrent slot is held at most. A slot is given back when its response ends; this bounds what a process that dies mid-request keeps. | [optional] [default to 900.0]
**MaxConcurrent** | Pointer to **NullableInt32** | Requests in flight at once. | [optional] 
**Models** | Pointer to **[]string** | The models a &#x60;per: model&#x60; rule limits, each as instance:model (the provider instance the model is called through, then the model), each counted on its own. Required there and refused on any other rule. | [optional] 
**Name** | **string** | Names the rule in a 429&#39;s detail and in the counter&#39;s key. Unique across rate_limits. | 
**Per** | **string** | What one count is shared by. | 
**Rpm** | Pointer to **NullableInt32** | Requests per minute. | [optional] 
**Tpm** | Pointer to **NullableInt32** | Tokens per minute, counted on what each request used; see tpm_admission for how a request is admitted. | [optional] 
**TpmAdmission** | Pointer to **string** | How a tpm limit admits a request. &#39;used&#39; (the default) admits a request while the minute&#39;s tokens are under the limit and counts what it used once it completes, as LiteLLM does. &#39;estimate&#39; holds the request&#39;s estimate (prompt plus max output, or budget_estimate_default_output_tokens) while it runs and refuses it when that does not fit, which is how providers count their own quotas: use it for a limit meant to stay under one. | [optional] [default to "used"]

## Methods

### NewRateLimitRuleCreate

`func NewRateLimitRuleCreate(name string, per string, ) *RateLimitRuleCreate`

NewRateLimitRuleCreate instantiates a new RateLimitRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRuleCreateWithDefaults

`func NewRateLimitRuleCreateWithDefaults() *RateLimitRuleCreate`

NewRateLimitRuleCreateWithDefaults instantiates a new RateLimitRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseSec

`func (o *RateLimitRuleCreate) GetLeaseSec() float32`

GetLeaseSec returns the LeaseSec field if non-nil, zero value otherwise.

### GetLeaseSecOk

`func (o *RateLimitRuleCreate) GetLeaseSecOk() (*float32, bool)`

GetLeaseSecOk returns a tuple with the LeaseSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseSec

`func (o *RateLimitRuleCreate) SetLeaseSec(v float32)`

SetLeaseSec sets LeaseSec field to given value.

### HasLeaseSec

`func (o *RateLimitRuleCreate) HasLeaseSec() bool`

HasLeaseSec returns a boolean if a field has been set.

### GetMaxConcurrent

`func (o *RateLimitRuleCreate) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *RateLimitRuleCreate) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *RateLimitRuleCreate) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.

### HasMaxConcurrent

`func (o *RateLimitRuleCreate) HasMaxConcurrent() bool`

HasMaxConcurrent returns a boolean if a field has been set.

### SetMaxConcurrentNil

`func (o *RateLimitRuleCreate) SetMaxConcurrentNil(b bool)`

 SetMaxConcurrentNil sets the value for MaxConcurrent to be an explicit nil

### UnsetMaxConcurrent
`func (o *RateLimitRuleCreate) UnsetMaxConcurrent()`

UnsetMaxConcurrent ensures that no value is present for MaxConcurrent, not even an explicit nil
### GetModels

`func (o *RateLimitRuleCreate) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *RateLimitRuleCreate) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *RateLimitRuleCreate) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *RateLimitRuleCreate) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *RateLimitRuleCreate) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *RateLimitRuleCreate) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetName

`func (o *RateLimitRuleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimitRuleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimitRuleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPer

`func (o *RateLimitRuleCreate) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimitRuleCreate) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimitRuleCreate) SetPer(v string)`

SetPer sets Per field to given value.


### GetRpm

`func (o *RateLimitRuleCreate) GetRpm() int32`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *RateLimitRuleCreate) GetRpmOk() (*int32, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *RateLimitRuleCreate) SetRpm(v int32)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *RateLimitRuleCreate) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### SetRpmNil

`func (o *RateLimitRuleCreate) SetRpmNil(b bool)`

 SetRpmNil sets the value for Rpm to be an explicit nil

### UnsetRpm
`func (o *RateLimitRuleCreate) UnsetRpm()`

UnsetRpm ensures that no value is present for Rpm, not even an explicit nil
### GetTpm

`func (o *RateLimitRuleCreate) GetTpm() int32`

GetTpm returns the Tpm field if non-nil, zero value otherwise.

### GetTpmOk

`func (o *RateLimitRuleCreate) GetTpmOk() (*int32, bool)`

GetTpmOk returns a tuple with the Tpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpm

`func (o *RateLimitRuleCreate) SetTpm(v int32)`

SetTpm sets Tpm field to given value.

### HasTpm

`func (o *RateLimitRuleCreate) HasTpm() bool`

HasTpm returns a boolean if a field has been set.

### SetTpmNil

`func (o *RateLimitRuleCreate) SetTpmNil(b bool)`

 SetTpmNil sets the value for Tpm to be an explicit nil

### UnsetTpm
`func (o *RateLimitRuleCreate) UnsetTpm()`

UnsetTpm ensures that no value is present for Tpm, not even an explicit nil
### GetTpmAdmission

`func (o *RateLimitRuleCreate) GetTpmAdmission() string`

GetTpmAdmission returns the TpmAdmission field if non-nil, zero value otherwise.

### GetTpmAdmissionOk

`func (o *RateLimitRuleCreate) GetTpmAdmissionOk() (*string, bool)`

GetTpmAdmissionOk returns a tuple with the TpmAdmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmAdmission

`func (o *RateLimitRuleCreate) SetTpmAdmission(v string)`

SetTpmAdmission sets TpmAdmission field to given value.

### HasTpmAdmission

`func (o *RateLimitRuleCreate) HasTpmAdmission() bool`

HasTpmAdmission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


