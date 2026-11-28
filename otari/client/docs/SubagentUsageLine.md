# SubagentUsageLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CacheWriteTokens** | Pointer to **int32** |  | [optional] [default to 0]
**InputTokens** | **int32** | Input tokens that were neither read from nor written to a cache. | 
**Model** | **string** | The model, as the harness names it. | 
**OutputTokens** | **int32** |  | 

## Methods

### NewSubagentUsageLine

`func NewSubagentUsageLine(inputTokens int32, model string, outputTokens int32, ) *SubagentUsageLine`

NewSubagentUsageLine instantiates a new SubagentUsageLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubagentUsageLineWithDefaults

`func NewSubagentUsageLineWithDefaults() *SubagentUsageLine`

NewSubagentUsageLineWithDefaults instantiates a new SubagentUsageLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadTokens

`func (o *SubagentUsageLine) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *SubagentUsageLine) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *SubagentUsageLine) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.

### HasCacheReadTokens

`func (o *SubagentUsageLine) HasCacheReadTokens() bool`

HasCacheReadTokens returns a boolean if a field has been set.

### GetCacheWriteTokens

`func (o *SubagentUsageLine) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *SubagentUsageLine) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *SubagentUsageLine) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *SubagentUsageLine) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### GetInputTokens

`func (o *SubagentUsageLine) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *SubagentUsageLine) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *SubagentUsageLine) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetModel

`func (o *SubagentUsageLine) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SubagentUsageLine) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SubagentUsageLine) SetModel(v string)`

SetModel sets Model field to given value.


### GetOutputTokens

`func (o *SubagentUsageLine) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *SubagentUsageLine) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *SubagentUsageLine) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


