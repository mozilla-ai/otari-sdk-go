# RequestSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | **int32** |  | 
**CostUsd** | **NullableString** |  | 
**PromptTokens** | **int32** |  | 
**RequestId** | **string** |  | 
**RowCount** | **int32** |  | 
**Status** | **string** |  | 
**TotalTokens** | **int32** |  | 

## Methods

### NewRequestSettlement

`func NewRequestSettlement(completionTokens int32, costUsd NullableString, promptTokens int32, requestId string, rowCount int32, status string, totalTokens int32, ) *RequestSettlement`

NewRequestSettlement instantiates a new RequestSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSettlementWithDefaults

`func NewRequestSettlementWithDefaults() *RequestSettlement`

NewRequestSettlementWithDefaults instantiates a new RequestSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *RequestSettlement) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *RequestSettlement) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *RequestSettlement) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetCostUsd

`func (o *RequestSettlement) GetCostUsd() string`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *RequestSettlement) GetCostUsdOk() (*string, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *RequestSettlement) SetCostUsd(v string)`

SetCostUsd sets CostUsd field to given value.


### SetCostUsdNil

`func (o *RequestSettlement) SetCostUsdNil(b bool)`

 SetCostUsdNil sets the value for CostUsd to be an explicit nil

### UnsetCostUsd
`func (o *RequestSettlement) UnsetCostUsd()`

UnsetCostUsd ensures that no value is present for CostUsd, not even an explicit nil
### GetPromptTokens

`func (o *RequestSettlement) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *RequestSettlement) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *RequestSettlement) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetRequestId

`func (o *RequestSettlement) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequestSettlement) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequestSettlement) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRowCount

`func (o *RequestSettlement) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *RequestSettlement) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *RequestSettlement) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.


### GetStatus

`func (o *RequestSettlement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestSettlement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestSettlement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalTokens

`func (o *RequestSettlement) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *RequestSettlement) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *RequestSettlement) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


