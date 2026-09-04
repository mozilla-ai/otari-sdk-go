# MRBetaAdvisorResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StopReason** | Pointer to **NullableString** |  | [optional] 
**Text** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaAdvisorResultBlock

`func NewMRBetaAdvisorResultBlock(text string, type_ string, ) *MRBetaAdvisorResultBlock`

NewMRBetaAdvisorResultBlock instantiates a new MRBetaAdvisorResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaAdvisorResultBlockWithDefaults

`func NewMRBetaAdvisorResultBlockWithDefaults() *MRBetaAdvisorResultBlock`

NewMRBetaAdvisorResultBlockWithDefaults instantiates a new MRBetaAdvisorResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopReason

`func (o *MRBetaAdvisorResultBlock) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *MRBetaAdvisorResultBlock) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *MRBetaAdvisorResultBlock) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *MRBetaAdvisorResultBlock) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### SetStopReasonNil

`func (o *MRBetaAdvisorResultBlock) SetStopReasonNil(b bool)`

 SetStopReasonNil sets the value for StopReason to be an explicit nil

### UnsetStopReason
`func (o *MRBetaAdvisorResultBlock) UnsetStopReason()`

UnsetStopReason ensures that no value is present for StopReason, not even an explicit nil
### GetText

`func (o *MRBetaAdvisorResultBlock) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MRBetaAdvisorResultBlock) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MRBetaAdvisorResultBlock) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *MRBetaAdvisorResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaAdvisorResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaAdvisorResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


