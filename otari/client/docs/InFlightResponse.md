# InFlightResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]InFlightEntry**](InFlightEntry.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewInFlightResponse

`func NewInFlightResponse(requests []InFlightEntry, total int32, ) *InFlightResponse`

NewInFlightResponse instantiates a new InFlightResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInFlightResponseWithDefaults

`func NewInFlightResponseWithDefaults() *InFlightResponse`

NewInFlightResponseWithDefaults instantiates a new InFlightResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *InFlightResponse) GetRequests() []InFlightEntry`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *InFlightResponse) GetRequestsOk() (*[]InFlightEntry, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *InFlightResponse) SetRequests(v []InFlightEntry)`

SetRequests sets Requests field to given value.


### GetTotal

`func (o *InFlightResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InFlightResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InFlightResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


