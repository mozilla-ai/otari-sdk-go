# UsageToolRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | **int32** |  | 
**Cost** | **float32** |  | 
**Errors** | **int32** |  | 
**Requests** | **int32** |  | 
**Tool** | **string** |  | 

## Methods

### NewUsageToolRow

`func NewUsageToolRow(calls int32, cost float32, errors int32, requests int32, tool string, ) *UsageToolRow`

NewUsageToolRow instantiates a new UsageToolRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageToolRowWithDefaults

`func NewUsageToolRowWithDefaults() *UsageToolRow`

NewUsageToolRowWithDefaults instantiates a new UsageToolRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *UsageToolRow) GetCalls() int32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *UsageToolRow) GetCallsOk() (*int32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *UsageToolRow) SetCalls(v int32)`

SetCalls sets Calls field to given value.


### GetCost

`func (o *UsageToolRow) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageToolRow) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageToolRow) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetErrors

`func (o *UsageToolRow) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UsageToolRow) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UsageToolRow) SetErrors(v int32)`

SetErrors sets Errors field to given value.


### GetRequests

`func (o *UsageToolRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageToolRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageToolRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetTool

`func (o *UsageToolRow) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *UsageToolRow) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *UsageToolRow) SetTool(v string)`

SetTool sets Tool field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


