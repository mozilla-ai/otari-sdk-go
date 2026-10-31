# HealthCheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRun** | [**NullableHealthCheckRunPublic**](HealthCheckRunPublic.md) |  | 
**Providers** | [**[]ProviderHealthSummary**](ProviderHealthSummary.md) |  | 
**Results** | [**[]HealthCheckResultPublic**](HealthCheckResultPublic.md) |  | 
**Running** | [**NullableHealthCheckRunPublic**](HealthCheckRunPublic.md) |  | 
**Settings** | [**HealthCheckSettingsPublic**](HealthCheckSettingsPublic.md) |  | 

## Methods

### NewHealthCheckStatus

`func NewHealthCheckStatus(lastRun NullableHealthCheckRunPublic, providers []ProviderHealthSummary, results []HealthCheckResultPublic, running NullableHealthCheckRunPublic, settings HealthCheckSettingsPublic, ) *HealthCheckStatus`

NewHealthCheckStatus instantiates a new HealthCheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckStatusWithDefaults

`func NewHealthCheckStatusWithDefaults() *HealthCheckStatus`

NewHealthCheckStatusWithDefaults instantiates a new HealthCheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastRun

`func (o *HealthCheckStatus) GetLastRun() HealthCheckRunPublic`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *HealthCheckStatus) GetLastRunOk() (*HealthCheckRunPublic, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *HealthCheckStatus) SetLastRun(v HealthCheckRunPublic)`

SetLastRun sets LastRun field to given value.


### SetLastRunNil

`func (o *HealthCheckStatus) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *HealthCheckStatus) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetProviders

`func (o *HealthCheckStatus) GetProviders() []ProviderHealthSummary`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *HealthCheckStatus) GetProvidersOk() (*[]ProviderHealthSummary, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *HealthCheckStatus) SetProviders(v []ProviderHealthSummary)`

SetProviders sets Providers field to given value.


### GetResults

`func (o *HealthCheckStatus) GetResults() []HealthCheckResultPublic`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HealthCheckStatus) GetResultsOk() (*[]HealthCheckResultPublic, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HealthCheckStatus) SetResults(v []HealthCheckResultPublic)`

SetResults sets Results field to given value.


### GetRunning

`func (o *HealthCheckStatus) GetRunning() HealthCheckRunPublic`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *HealthCheckStatus) GetRunningOk() (*HealthCheckRunPublic, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *HealthCheckStatus) SetRunning(v HealthCheckRunPublic)`

SetRunning sets Running field to given value.


### SetRunningNil

`func (o *HealthCheckStatus) SetRunningNil(b bool)`

 SetRunningNil sets the value for Running to be an explicit nil

### UnsetRunning
`func (o *HealthCheckStatus) UnsetRunning()`

UnsetRunning ensures that no value is present for Running, not even an explicit nil
### GetSettings

`func (o *HealthCheckStatus) GetSettings() HealthCheckSettingsPublic`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HealthCheckStatus) GetSettingsOk() (*HealthCheckSettingsPublic, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HealthCheckStatus) SetSettings(v HealthCheckSettingsPublic)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


