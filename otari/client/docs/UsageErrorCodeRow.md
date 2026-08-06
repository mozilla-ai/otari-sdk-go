# UsageErrorCodeRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorClass** | **string** |  | 
**Requests** | **int32** |  | 
**StatusCode** | **NullableInt32** |  | 

## Methods

### NewUsageErrorCodeRow

`func NewUsageErrorCodeRow(errorClass string, requests int32, statusCode NullableInt32, ) *UsageErrorCodeRow`

NewUsageErrorCodeRow instantiates a new UsageErrorCodeRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageErrorCodeRowWithDefaults

`func NewUsageErrorCodeRowWithDefaults() *UsageErrorCodeRow`

NewUsageErrorCodeRowWithDefaults instantiates a new UsageErrorCodeRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorClass

`func (o *UsageErrorCodeRow) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *UsageErrorCodeRow) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *UsageErrorCodeRow) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.


### GetRequests

`func (o *UsageErrorCodeRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageErrorCodeRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageErrorCodeRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetStatusCode

`func (o *UsageErrorCodeRow) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UsageErrorCodeRow) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UsageErrorCodeRow) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### SetStatusCodeNil

`func (o *UsageErrorCodeRow) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *UsageErrorCodeRow) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


