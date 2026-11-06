# HealthCheckResultPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | **time.Time** |  | 
**Error** | **NullableString** |  | 
**FailedRuns** | Pointer to **int32** | How many runs in a row it has failed. | [optional] [default to 0]
**FailingSince** | Pointer to **NullableTime** | When a failing model&#39;s unbroken run of failures began, as far back as results are kept (48 hours). Null when it answered. | [optional] 
**LatencyMs** | **int32** |  | 
**Provider** | **string** |  | 
**Selector** | **string** |  | 
**Status** | **string** |  | 
**StatusCode** | **NullableInt32** |  | 

## Methods

### NewHealthCheckResultPublic

`func NewHealthCheckResultPublic(checkedAt time.Time, error_ NullableString, latencyMs int32, provider string, selector string, status string, statusCode NullableInt32, ) *HealthCheckResultPublic`

NewHealthCheckResultPublic instantiates a new HealthCheckResultPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResultPublicWithDefaults

`func NewHealthCheckResultPublicWithDefaults() *HealthCheckResultPublic`

NewHealthCheckResultPublicWithDefaults instantiates a new HealthCheckResultPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *HealthCheckResultPublic) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *HealthCheckResultPublic) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *HealthCheckResultPublic) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.


### GetError

`func (o *HealthCheckResultPublic) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HealthCheckResultPublic) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HealthCheckResultPublic) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *HealthCheckResultPublic) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HealthCheckResultPublic) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFailedRuns

`func (o *HealthCheckResultPublic) GetFailedRuns() int32`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *HealthCheckResultPublic) GetFailedRunsOk() (*int32, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *HealthCheckResultPublic) SetFailedRuns(v int32)`

SetFailedRuns sets FailedRuns field to given value.

### HasFailedRuns

`func (o *HealthCheckResultPublic) HasFailedRuns() bool`

HasFailedRuns returns a boolean if a field has been set.

### GetFailingSince

`func (o *HealthCheckResultPublic) GetFailingSince() time.Time`

GetFailingSince returns the FailingSince field if non-nil, zero value otherwise.

### GetFailingSinceOk

`func (o *HealthCheckResultPublic) GetFailingSinceOk() (*time.Time, bool)`

GetFailingSinceOk returns a tuple with the FailingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingSince

`func (o *HealthCheckResultPublic) SetFailingSince(v time.Time)`

SetFailingSince sets FailingSince field to given value.

### HasFailingSince

`func (o *HealthCheckResultPublic) HasFailingSince() bool`

HasFailingSince returns a boolean if a field has been set.

### SetFailingSinceNil

`func (o *HealthCheckResultPublic) SetFailingSinceNil(b bool)`

 SetFailingSinceNil sets the value for FailingSince to be an explicit nil

### UnsetFailingSince
`func (o *HealthCheckResultPublic) UnsetFailingSince()`

UnsetFailingSince ensures that no value is present for FailingSince, not even an explicit nil
### GetLatencyMs

`func (o *HealthCheckResultPublic) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *HealthCheckResultPublic) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *HealthCheckResultPublic) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.


### GetProvider

`func (o *HealthCheckResultPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HealthCheckResultPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HealthCheckResultPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSelector

`func (o *HealthCheckResultPublic) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *HealthCheckResultPublic) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *HealthCheckResultPublic) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetStatus

`func (o *HealthCheckResultPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResultPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResultPublic) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusCode

`func (o *HealthCheckResultPublic) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HealthCheckResultPublic) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HealthCheckResultPublic) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### SetStatusCodeNil

`func (o *HealthCheckResultPublic) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *HealthCheckResultPublic) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


