# CheckVerdictRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**GateId** | **string** |  | 
**Outcome** | **string** |  | 

## Methods

### NewCheckVerdictRequest

`func NewCheckVerdictRequest(gateId string, outcome string, ) *CheckVerdictRequest`

NewCheckVerdictRequest instantiates a new CheckVerdictRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckVerdictRequestWithDefaults

`func NewCheckVerdictRequestWithDefaults() *CheckVerdictRequest`

NewCheckVerdictRequestWithDefaults instantiates a new CheckVerdictRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *CheckVerdictRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CheckVerdictRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CheckVerdictRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CheckVerdictRequest) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetGateId

`func (o *CheckVerdictRequest) GetGateId() string`

GetGateId returns the GateId field if non-nil, zero value otherwise.

### GetGateIdOk

`func (o *CheckVerdictRequest) GetGateIdOk() (*string, bool)`

GetGateIdOk returns a tuple with the GateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateId

`func (o *CheckVerdictRequest) SetGateId(v string)`

SetGateId sets GateId field to given value.


### GetOutcome

`func (o *CheckVerdictRequest) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *CheckVerdictRequest) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *CheckVerdictRequest) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


