# ProviderHistoryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answered** | **int32** |  | 
**StartedAt** | **time.Time** |  | 
**Total** | **int32** |  | 

## Methods

### NewProviderHistoryPoint

`func NewProviderHistoryPoint(answered int32, startedAt time.Time, total int32, ) *ProviderHistoryPoint`

NewProviderHistoryPoint instantiates a new ProviderHistoryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderHistoryPointWithDefaults

`func NewProviderHistoryPointWithDefaults() *ProviderHistoryPoint`

NewProviderHistoryPointWithDefaults instantiates a new ProviderHistoryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswered

`func (o *ProviderHistoryPoint) GetAnswered() int32`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *ProviderHistoryPoint) GetAnsweredOk() (*int32, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *ProviderHistoryPoint) SetAnswered(v int32)`

SetAnswered sets Answered field to given value.


### GetStartedAt

`func (o *ProviderHistoryPoint) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ProviderHistoryPoint) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ProviderHistoryPoint) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTotal

`func (o *ProviderHistoryPoint) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProviderHistoryPoint) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProviderHistoryPoint) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


