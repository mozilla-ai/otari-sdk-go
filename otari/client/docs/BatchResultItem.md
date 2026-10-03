# BatchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomId** | **string** | Identifier supplied for this request in the batch. | 
**Result** | [**NullableChatCompletion**](ChatCompletion.md) | Serialized provider result, or null when the request failed. | 
**Error** | [**NullableBatchResultItemError**](BatchResultItemError.md) |  | 

## Methods

### NewBatchResultItem

`func NewBatchResultItem(customId string, result NullableChatCompletion, error_ NullableBatchResultItemError, ) *BatchResultItem`

NewBatchResultItem instantiates a new BatchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResultItemWithDefaults

`func NewBatchResultItemWithDefaults() *BatchResultItem`

NewBatchResultItemWithDefaults instantiates a new BatchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomId

`func (o *BatchResultItem) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *BatchResultItem) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *BatchResultItem) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetResult

`func (o *BatchResultItem) GetResult() ChatCompletion`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BatchResultItem) GetResultOk() (*ChatCompletion, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BatchResultItem) SetResult(v ChatCompletion)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *BatchResultItem) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *BatchResultItem) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *BatchResultItem) GetError() BatchResultItemError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchResultItem) GetErrorOk() (*BatchResultItemError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchResultItem) SetError(v BatchResultItemError)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *BatchResultItem) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BatchResultItem) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


