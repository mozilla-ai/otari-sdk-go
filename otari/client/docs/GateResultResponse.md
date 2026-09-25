# GateResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **NullableString** |  | [optional] 
**Enforcement** | **string** |  | 
**GateId** | **string** |  | 
**Message** | **string** |  | 
**Outcome** | **string** |  | 

## Methods

### NewGateResultResponse

`func NewGateResultResponse(enforcement string, gateId string, message string, outcome string, ) *GateResultResponse`

NewGateResultResponse instantiates a new GateResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGateResultResponseWithDefaults

`func NewGateResultResponseWithDefaults() *GateResultResponse`

NewGateResultResponseWithDefaults instantiates a new GateResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *GateResultResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GateResultResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GateResultResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GateResultResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *GateResultResponse) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *GateResultResponse) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetEnforcement

`func (o *GateResultResponse) GetEnforcement() string`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *GateResultResponse) GetEnforcementOk() (*string, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *GateResultResponse) SetEnforcement(v string)`

SetEnforcement sets Enforcement field to given value.


### GetGateId

`func (o *GateResultResponse) GetGateId() string`

GetGateId returns the GateId field if non-nil, zero value otherwise.

### GetGateIdOk

`func (o *GateResultResponse) GetGateIdOk() (*string, bool)`

GetGateIdOk returns a tuple with the GateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateId

`func (o *GateResultResponse) SetGateId(v string)`

SetGateId sets GateId field to given value.


### GetMessage

`func (o *GateResultResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GateResultResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GateResultResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOutcome

`func (o *GateResultResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *GateResultResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *GateResultResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


