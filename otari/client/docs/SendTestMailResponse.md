# SendTestMailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Transport** | **string** |  | 

## Methods

### NewSendTestMailResponse

`func NewSendTestMailResponse(ok bool, transport string, ) *SendTestMailResponse`

NewSendTestMailResponse instantiates a new SendTestMailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTestMailResponseWithDefaults

`func NewSendTestMailResponseWithDefaults() *SendTestMailResponse`

NewSendTestMailResponseWithDefaults instantiates a new SendTestMailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *SendTestMailResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *SendTestMailResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *SendTestMailResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetReason

`func (o *SendTestMailResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SendTestMailResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SendTestMailResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SendTestMailResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *SendTestMailResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *SendTestMailResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetTransport

`func (o *SendTestMailResponse) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *SendTestMailResponse) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *SendTestMailResponse) SetTransport(v string)`

SetTransport sets Transport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


