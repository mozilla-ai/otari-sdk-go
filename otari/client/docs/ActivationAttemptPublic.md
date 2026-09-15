# ActivationAttemptPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostUsd** | Pointer to **NullableFloat32** |  | [optional] 
**ErrorCategory** | Pointer to **NullableString** | What kind of failure this was, for a failed attempt only. The dashboard renders its own sentence per category; the provider&#39;s own error text is never returned here. | [optional] 
**LatencyMs** | Pointer to **NullableInt32** |  | [optional] 
**Model** | Pointer to **NullableString** | Model the request named. | [optional] 
**OccurredAt** | **string** | When the request was recorded, UTC ISO-8601. | 
**Provider** | Pointer to **NullableString** | Provider instance that served it, when one did. | [optional] 
**RequestId** | **string** | The usage row&#39;s id, which the Activity page can be filtered by. | 
**Status** | **string** |  | 

## Methods

### NewActivationAttemptPublic

`func NewActivationAttemptPublic(occurredAt string, requestId string, status string, ) *ActivationAttemptPublic`

NewActivationAttemptPublic instantiates a new ActivationAttemptPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationAttemptPublicWithDefaults

`func NewActivationAttemptPublicWithDefaults() *ActivationAttemptPublic`

NewActivationAttemptPublicWithDefaults instantiates a new ActivationAttemptPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostUsd

`func (o *ActivationAttemptPublic) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ActivationAttemptPublic) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ActivationAttemptPublic) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ActivationAttemptPublic) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### SetCostUsdNil

`func (o *ActivationAttemptPublic) SetCostUsdNil(b bool)`

 SetCostUsdNil sets the value for CostUsd to be an explicit nil

### UnsetCostUsd
`func (o *ActivationAttemptPublic) UnsetCostUsd()`

UnsetCostUsd ensures that no value is present for CostUsd, not even an explicit nil
### GetErrorCategory

`func (o *ActivationAttemptPublic) GetErrorCategory() string`

GetErrorCategory returns the ErrorCategory field if non-nil, zero value otherwise.

### GetErrorCategoryOk

`func (o *ActivationAttemptPublic) GetErrorCategoryOk() (*string, bool)`

GetErrorCategoryOk returns a tuple with the ErrorCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCategory

`func (o *ActivationAttemptPublic) SetErrorCategory(v string)`

SetErrorCategory sets ErrorCategory field to given value.

### HasErrorCategory

`func (o *ActivationAttemptPublic) HasErrorCategory() bool`

HasErrorCategory returns a boolean if a field has been set.

### SetErrorCategoryNil

`func (o *ActivationAttemptPublic) SetErrorCategoryNil(b bool)`

 SetErrorCategoryNil sets the value for ErrorCategory to be an explicit nil

### UnsetErrorCategory
`func (o *ActivationAttemptPublic) UnsetErrorCategory()`

UnsetErrorCategory ensures that no value is present for ErrorCategory, not even an explicit nil
### GetLatencyMs

`func (o *ActivationAttemptPublic) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *ActivationAttemptPublic) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *ActivationAttemptPublic) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *ActivationAttemptPublic) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### SetLatencyMsNil

`func (o *ActivationAttemptPublic) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *ActivationAttemptPublic) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
### GetModel

`func (o *ActivationAttemptPublic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ActivationAttemptPublic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ActivationAttemptPublic) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ActivationAttemptPublic) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ActivationAttemptPublic) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ActivationAttemptPublic) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetOccurredAt

`func (o *ActivationAttemptPublic) GetOccurredAt() string`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ActivationAttemptPublic) GetOccurredAtOk() (*string, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ActivationAttemptPublic) SetOccurredAt(v string)`

SetOccurredAt sets OccurredAt field to given value.


### GetProvider

`func (o *ActivationAttemptPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ActivationAttemptPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ActivationAttemptPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ActivationAttemptPublic) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ActivationAttemptPublic) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ActivationAttemptPublic) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetRequestId

`func (o *ActivationAttemptPublic) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ActivationAttemptPublic) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ActivationAttemptPublic) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *ActivationAttemptPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivationAttemptPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivationAttemptPublic) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


