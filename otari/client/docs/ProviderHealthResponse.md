# ProviderHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **NullableString** | ISO 8601 time of the most recent per-provider check (null if none yet). | [optional] 
**Degraded** | Pointer to **int32** | How many providers are not counted as reachable only because model discovery is unavailable for them. These may still serve requests. | [optional] [default to 0]
**Healthy** | **int32** | How many providers are currently reachable. | 
**Providers** | [**[]ProviderHealthSchema**](ProviderHealthSchema.md) |  | 
**Total** | **int32** | How many providers are configured. | 

## Methods

### NewProviderHealthResponse

`func NewProviderHealthResponse(healthy int32, providers []ProviderHealthSchema, total int32, ) *ProviderHealthResponse`

NewProviderHealthResponse instantiates a new ProviderHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderHealthResponseWithDefaults

`func NewProviderHealthResponseWithDefaults() *ProviderHealthResponse`

NewProviderHealthResponseWithDefaults instantiates a new ProviderHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *ProviderHealthResponse) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *ProviderHealthResponse) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *ProviderHealthResponse) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *ProviderHealthResponse) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### SetCheckedAtNil

`func (o *ProviderHealthResponse) SetCheckedAtNil(b bool)`

 SetCheckedAtNil sets the value for CheckedAt to be an explicit nil

### UnsetCheckedAt
`func (o *ProviderHealthResponse) UnsetCheckedAt()`

UnsetCheckedAt ensures that no value is present for CheckedAt, not even an explicit nil
### GetDegraded

`func (o *ProviderHealthResponse) GetDegraded() int32`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *ProviderHealthResponse) GetDegradedOk() (*int32, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *ProviderHealthResponse) SetDegraded(v int32)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *ProviderHealthResponse) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetHealthy

`func (o *ProviderHealthResponse) GetHealthy() int32`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *ProviderHealthResponse) GetHealthyOk() (*int32, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *ProviderHealthResponse) SetHealthy(v int32)`

SetHealthy sets Healthy field to given value.


### GetProviders

`func (o *ProviderHealthResponse) GetProviders() []ProviderHealthSchema`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ProviderHealthResponse) GetProvidersOk() (*[]ProviderHealthSchema, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ProviderHealthResponse) SetProviders(v []ProviderHealthSchema)`

SetProviders sets Providers field to given value.


### GetTotal

`func (o *ProviderHealthResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderHealthResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderHealthResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


