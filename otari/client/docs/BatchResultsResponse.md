# BatchResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]BatchResultItem**](BatchResultItem.md) |  | 

## Methods

### NewBatchResultsResponse

`func NewBatchResultsResponse(results []BatchResultItem, ) *BatchResultsResponse`

NewBatchResultsResponse instantiates a new BatchResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResultsResponseWithDefaults

`func NewBatchResultsResponseWithDefaults() *BatchResultsResponse`

NewBatchResultsResponseWithDefaults instantiates a new BatchResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchResultsResponse) GetResults() []BatchResultItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResultsResponse) GetResultsOk() (*[]BatchResultItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResultsResponse) SetResults(v []BatchResultItem)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


