# CreateBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionWindow** | Pointer to **string** |  | [optional] [default to "24h"]
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Model** | **string** |  | 
**Requests** | [**[]BatchRequestItem**](BatchRequestItem.md) |  | 

## Methods

### NewCreateBatchRequest

`func NewCreateBatchRequest(model string, requests []BatchRequestItem, ) *CreateBatchRequest`

NewCreateBatchRequest instantiates a new CreateBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchRequestWithDefaults

`func NewCreateBatchRequestWithDefaults() *CreateBatchRequest`

NewCreateBatchRequestWithDefaults instantiates a new CreateBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionWindow

`func (o *CreateBatchRequest) GetCompletionWindow() string`

GetCompletionWindow returns the CompletionWindow field if non-nil, zero value otherwise.

### GetCompletionWindowOk

`func (o *CreateBatchRequest) GetCompletionWindowOk() (*string, bool)`

GetCompletionWindowOk returns a tuple with the CompletionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionWindow

`func (o *CreateBatchRequest) SetCompletionWindow(v string)`

SetCompletionWindow sets CompletionWindow field to given value.

### HasCompletionWindow

`func (o *CreateBatchRequest) HasCompletionWindow() bool`

HasCompletionWindow returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateBatchRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateBatchRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateBatchRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateBatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateBatchRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateBatchRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *CreateBatchRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateBatchRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateBatchRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetRequests

`func (o *CreateBatchRequest) GetRequests() []BatchRequestItem`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CreateBatchRequest) GetRequestsOk() (*[]BatchRequestItem, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CreateBatchRequest) SetRequests(v []BatchRequestItem)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


