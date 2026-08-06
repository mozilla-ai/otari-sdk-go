# UsageGroupRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** |  | 
**IsOther** | Pointer to **bool** |  | [optional] [default to false]
**Key** | **NullableString** |  | 
**Requests** | **int32** |  | 
**Tokens** | **int32** |  | 

## Methods

### NewUsageGroupRow

`func NewUsageGroupRow(cost float32, key NullableString, requests int32, tokens int32, ) *UsageGroupRow`

NewUsageGroupRow instantiates a new UsageGroupRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageGroupRowWithDefaults

`func NewUsageGroupRowWithDefaults() *UsageGroupRow`

NewUsageGroupRowWithDefaults instantiates a new UsageGroupRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *UsageGroupRow) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageGroupRow) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageGroupRow) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetIsOther

`func (o *UsageGroupRow) GetIsOther() bool`

GetIsOther returns the IsOther field if non-nil, zero value otherwise.

### GetIsOtherOk

`func (o *UsageGroupRow) GetIsOtherOk() (*bool, bool)`

GetIsOtherOk returns a tuple with the IsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOther

`func (o *UsageGroupRow) SetIsOther(v bool)`

SetIsOther sets IsOther field to given value.

### HasIsOther

`func (o *UsageGroupRow) HasIsOther() bool`

HasIsOther returns a boolean if a field has been set.

### GetKey

`func (o *UsageGroupRow) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UsageGroupRow) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UsageGroupRow) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *UsageGroupRow) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UsageGroupRow) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetRequests

`func (o *UsageGroupRow) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *UsageGroupRow) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *UsageGroupRow) SetRequests(v int32)`

SetRequests sets Requests field to given value.


### GetTokens

`func (o *UsageGroupRow) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UsageGroupRow) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UsageGroupRow) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


