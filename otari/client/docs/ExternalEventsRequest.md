# ExternalEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]ExternalUsageEvent**](ExternalUsageEvent.md) |  | 
**Source** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExternalEventsRequest

`func NewExternalEventsRequest(events []ExternalUsageEvent, source string, ) *ExternalEventsRequest`

NewExternalEventsRequest instantiates a new ExternalEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEventsRequestWithDefaults

`func NewExternalEventsRequestWithDefaults() *ExternalEventsRequest`

NewExternalEventsRequestWithDefaults instantiates a new ExternalEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ExternalEventsRequest) GetEvents() []ExternalUsageEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ExternalEventsRequest) GetEventsOk() (*[]ExternalUsageEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ExternalEventsRequest) SetEvents(v []ExternalUsageEvent)`

SetEvents sets Events field to given value.


### GetSource

`func (o *ExternalEventsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExternalEventsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExternalEventsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetUserId

`func (o *ExternalEventsRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEventsRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEventsRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExternalEventsRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExternalEventsRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExternalEventsRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


