# AgentTelemetryGroupRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOther** | Pointer to **bool** |  | [optional] [default to false]
**Key** | **NullableString** |  | 
**Rows** | **int32** |  | 

## Methods

### NewAgentTelemetryGroupRow

`func NewAgentTelemetryGroupRow(key NullableString, rows int32, ) *AgentTelemetryGroupRow`

NewAgentTelemetryGroupRow instantiates a new AgentTelemetryGroupRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryGroupRowWithDefaults

`func NewAgentTelemetryGroupRowWithDefaults() *AgentTelemetryGroupRow`

NewAgentTelemetryGroupRowWithDefaults instantiates a new AgentTelemetryGroupRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOther

`func (o *AgentTelemetryGroupRow) GetIsOther() bool`

GetIsOther returns the IsOther field if non-nil, zero value otherwise.

### GetIsOtherOk

`func (o *AgentTelemetryGroupRow) GetIsOtherOk() (*bool, bool)`

GetIsOtherOk returns a tuple with the IsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOther

`func (o *AgentTelemetryGroupRow) SetIsOther(v bool)`

SetIsOther sets IsOther field to given value.

### HasIsOther

`func (o *AgentTelemetryGroupRow) HasIsOther() bool`

HasIsOther returns a boolean if a field has been set.

### GetKey

`func (o *AgentTelemetryGroupRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentTelemetryGroupRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentTelemetryGroupRow) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *AgentTelemetryGroupRow) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AgentTelemetryGroupRow) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetRows

`func (o *AgentTelemetryGroupRow) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *AgentTelemetryGroupRow) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *AgentTelemetryGroupRow) SetRows(v int32)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


