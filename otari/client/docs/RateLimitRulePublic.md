# RateLimitRulePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseSec** | Pointer to **float32** | How long a max_concurrent slot is held at most. A slot is given back when its response ends; this bounds what a process that dies mid-request keeps. | [optional] [default to 900.0]
**MaxConcurrent** | Pointer to **NullableInt32** | Requests in flight at once. | [optional] 
**Models** | Pointer to **[]string** | The models a &#x60;per: model&#x60; rule limits, each as instance:model (the provider instance the model is called through, then the model), each counted on its own. Required there and refused on any other rule. | [optional] 
**Name** | **string** | Names the rule in a 429&#39;s detail and in the counter&#39;s key. Unique across rate_limits. | 
**Per** | **string** | What one count is shared by. | 
**Rpm** | Pointer to **NullableInt32** | Requests per minute. | [optional] 
**Source** | **string** | &#39;config&#39; for a rule from config.yml, which is read-only here; &#39;dashboard&#39; for a stored one. | 
**Tpm** | Pointer to **NullableInt32** | Tokens per minute, counted on what each request used; see tpm_admission for how a request is admitted. | [optional] 
**TpmAdmission** | Pointer to **string** | How a tpm limit admits a request. &#39;used&#39; (the default) admits a request while the minute&#39;s tokens are under the limit and counts what it used once it completes, as LiteLLM does. &#39;estimate&#39; holds the request&#39;s estimate (prompt plus max output, or budget_estimate_default_output_tokens) while it runs and refuses it when that does not fit, which is how providers count their own quotas: use it for a limit meant to stay under one. | [optional] [default to "used"]
**UpdatedAt** | Pointer to **NullableTime** | When a stored rule last changed. | [optional] 

## Methods

### NewRateLimitRulePublic

`func NewRateLimitRulePublic(name string, per string, source string, ) *RateLimitRulePublic`

NewRateLimitRulePublic instantiates a new RateLimitRulePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRulePublicWithDefaults

`func NewRateLimitRulePublicWithDefaults() *RateLimitRulePublic`

NewRateLimitRulePublicWithDefaults instantiates a new RateLimitRulePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseSec

`func (o *RateLimitRulePublic) GetLeaseSec() float32`

GetLeaseSec returns the LeaseSec field if non-nil, zero value otherwise.

### GetLeaseSecOk

`func (o *RateLimitRulePublic) GetLeaseSecOk() (*float32, bool)`

GetLeaseSecOk returns a tuple with the LeaseSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseSec

`func (o *RateLimitRulePublic) SetLeaseSec(v float32)`

SetLeaseSec sets LeaseSec field to given value.

### HasLeaseSec

`func (o *RateLimitRulePublic) HasLeaseSec() bool`

HasLeaseSec returns a boolean if a field has been set.

### GetMaxConcurrent

`func (o *RateLimitRulePublic) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *RateLimitRulePublic) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *RateLimitRulePublic) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.

### HasMaxConcurrent

`func (o *RateLimitRulePublic) HasMaxConcurrent() bool`

HasMaxConcurrent returns a boolean if a field has been set.

### SetMaxConcurrentNil

`func (o *RateLimitRulePublic) SetMaxConcurrentNil(b bool)`

 SetMaxConcurrentNil sets the value for MaxConcurrent to be an explicit nil

### UnsetMaxConcurrent
`func (o *RateLimitRulePublic) UnsetMaxConcurrent()`

UnsetMaxConcurrent ensures that no value is present for MaxConcurrent, not even an explicit nil
### GetModels

`func (o *RateLimitRulePublic) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *RateLimitRulePublic) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *RateLimitRulePublic) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *RateLimitRulePublic) HasModels() bool`

HasModels returns a boolean if a field has been set.

### SetModelsNil

`func (o *RateLimitRulePublic) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *RateLimitRulePublic) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil
### GetName

`func (o *RateLimitRulePublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimitRulePublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimitRulePublic) SetName(v string)`

SetName sets Name field to given value.


### GetPer

`func (o *RateLimitRulePublic) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *RateLimitRulePublic) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *RateLimitRulePublic) SetPer(v string)`

SetPer sets Per field to given value.


### GetRpm

`func (o *RateLimitRulePublic) GetRpm() int32`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *RateLimitRulePublic) GetRpmOk() (*int32, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *RateLimitRulePublic) SetRpm(v int32)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *RateLimitRulePublic) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### SetRpmNil

`func (o *RateLimitRulePublic) SetRpmNil(b bool)`

 SetRpmNil sets the value for Rpm to be an explicit nil

### UnsetRpm
`func (o *RateLimitRulePublic) UnsetRpm()`

UnsetRpm ensures that no value is present for Rpm, not even an explicit nil
### GetSource

`func (o *RateLimitRulePublic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RateLimitRulePublic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RateLimitRulePublic) SetSource(v string)`

SetSource sets Source field to given value.


### GetTpm

`func (o *RateLimitRulePublic) GetTpm() int32`

GetTpm returns the Tpm field if non-nil, zero value otherwise.

### GetTpmOk

`func (o *RateLimitRulePublic) GetTpmOk() (*int32, bool)`

GetTpmOk returns a tuple with the Tpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpm

`func (o *RateLimitRulePublic) SetTpm(v int32)`

SetTpm sets Tpm field to given value.

### HasTpm

`func (o *RateLimitRulePublic) HasTpm() bool`

HasTpm returns a boolean if a field has been set.

### SetTpmNil

`func (o *RateLimitRulePublic) SetTpmNil(b bool)`

 SetTpmNil sets the value for Tpm to be an explicit nil

### UnsetTpm
`func (o *RateLimitRulePublic) UnsetTpm()`

UnsetTpm ensures that no value is present for Tpm, not even an explicit nil
### GetTpmAdmission

`func (o *RateLimitRulePublic) GetTpmAdmission() string`

GetTpmAdmission returns the TpmAdmission field if non-nil, zero value otherwise.

### GetTpmAdmissionOk

`func (o *RateLimitRulePublic) GetTpmAdmissionOk() (*string, bool)`

GetTpmAdmissionOk returns a tuple with the TpmAdmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmAdmission

`func (o *RateLimitRulePublic) SetTpmAdmission(v string)`

SetTpmAdmission sets TpmAdmission field to given value.

### HasTpmAdmission

`func (o *RateLimitRulePublic) HasTpmAdmission() bool`

HasTpmAdmission returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RateLimitRulePublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RateLimitRulePublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RateLimitRulePublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RateLimitRulePublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *RateLimitRulePublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *RateLimitRulePublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


