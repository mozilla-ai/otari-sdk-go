# BatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CompletionWindow** | **string** |  | 
**CreatedAt** | **int32** |  | 
**Endpoint** | **string** |  | 
**InputFileId** | **string** |  | 
**Object** | **string** |  | 
**Status** | **string** |  | 
**CancelledAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**CancellingAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**CompletedAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**ErrorFileId** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Errors** | Pointer to [**NullableBATCHErrors**](BATCHErrors.md) |  | [optional] 
**ExpiredAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**ExpiresAt** | Pointer to **NullableInt32** |  | [optional] 
**FailedAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**FinalizingAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**InProgressAt** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Model** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**OutputFileId** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**RequestCounts** | Pointer to [**NullableBATCHBatchRequestCounts**](BATCHBatchRequestCounts.md) |  | [optional] 
**Usage** | Pointer to [**NullableBATCHBatchUsage**](BATCHBatchUsage.md) |  | [optional] 
**Provider** | **string** | Provider instance that owns the batch. | 

## Methods

### NewBatchResponse

`func NewBatchResponse(id string, completionWindow string, createdAt int32, endpoint string, inputFileId string, object string, status string, provider string, ) *BatchResponse`

NewBatchResponse instantiates a new BatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseWithDefaults

`func NewBatchResponseWithDefaults() *BatchResponse`

NewBatchResponseWithDefaults instantiates a new BatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCompletionWindow

`func (o *BatchResponse) GetCompletionWindow() string`

GetCompletionWindow returns the CompletionWindow field if non-nil, zero value otherwise.

### GetCompletionWindowOk

`func (o *BatchResponse) GetCompletionWindowOk() (*string, bool)`

GetCompletionWindowOk returns a tuple with the CompletionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionWindow

`func (o *BatchResponse) SetCompletionWindow(v string)`

SetCompletionWindow sets CompletionWindow field to given value.


### GetCreatedAt

`func (o *BatchResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndpoint

`func (o *BatchResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BatchResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BatchResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetInputFileId

`func (o *BatchResponse) GetInputFileId() string`

GetInputFileId returns the InputFileId field if non-nil, zero value otherwise.

### GetInputFileIdOk

`func (o *BatchResponse) GetInputFileIdOk() (*string, bool)`

GetInputFileIdOk returns a tuple with the InputFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFileId

`func (o *BatchResponse) SetInputFileId(v string)`

SetInputFileId sets InputFileId field to given value.


### GetObject

`func (o *BatchResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BatchResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BatchResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *BatchResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCancelledAt

`func (o *BatchResponse) GetCancelledAt() int32`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *BatchResponse) GetCancelledAtOk() (*int32, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *BatchResponse) SetCancelledAt(v int32)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *BatchResponse) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *BatchResponse) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *BatchResponse) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetCancellingAt

`func (o *BatchResponse) GetCancellingAt() int32`

GetCancellingAt returns the CancellingAt field if non-nil, zero value otherwise.

### GetCancellingAtOk

`func (o *BatchResponse) GetCancellingAtOk() (*int32, bool)`

GetCancellingAtOk returns a tuple with the CancellingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellingAt

`func (o *BatchResponse) SetCancellingAt(v int32)`

SetCancellingAt sets CancellingAt field to given value.

### HasCancellingAt

`func (o *BatchResponse) HasCancellingAt() bool`

HasCancellingAt returns a boolean if a field has been set.

### SetCancellingAtNil

`func (o *BatchResponse) SetCancellingAtNil(b bool)`

 SetCancellingAtNil sets the value for CancellingAt to be an explicit nil

### UnsetCancellingAt
`func (o *BatchResponse) UnsetCancellingAt()`

UnsetCancellingAt ensures that no value is present for CancellingAt, not even an explicit nil
### GetCompletedAt

`func (o *BatchResponse) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponse) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponse) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BatchResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BatchResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BatchResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetErrorFileId

`func (o *BatchResponse) GetErrorFileId() string`

GetErrorFileId returns the ErrorFileId field if non-nil, zero value otherwise.

### GetErrorFileIdOk

`func (o *BatchResponse) GetErrorFileIdOk() (*string, bool)`

GetErrorFileIdOk returns a tuple with the ErrorFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFileId

`func (o *BatchResponse) SetErrorFileId(v string)`

SetErrorFileId sets ErrorFileId field to given value.

### HasErrorFileId

`func (o *BatchResponse) HasErrorFileId() bool`

HasErrorFileId returns a boolean if a field has been set.

### SetErrorFileIdNil

`func (o *BatchResponse) SetErrorFileIdNil(b bool)`

 SetErrorFileIdNil sets the value for ErrorFileId to be an explicit nil

### UnsetErrorFileId
`func (o *BatchResponse) UnsetErrorFileId()`

UnsetErrorFileId ensures that no value is present for ErrorFileId, not even an explicit nil
### GetErrors

`func (o *BatchResponse) GetErrors() BATCHErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponse) GetErrorsOk() (*BATCHErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponse) SetErrors(v BATCHErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BatchResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BatchResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetExpiredAt

`func (o *BatchResponse) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *BatchResponse) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *BatchResponse) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *BatchResponse) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *BatchResponse) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *BatchResponse) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetExpiresAt

`func (o *BatchResponse) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BatchResponse) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BatchResponse) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BatchResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *BatchResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *BatchResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFailedAt

`func (o *BatchResponse) GetFailedAt() int32`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *BatchResponse) GetFailedAtOk() (*int32, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *BatchResponse) SetFailedAt(v int32)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *BatchResponse) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *BatchResponse) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *BatchResponse) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetFinalizingAt

`func (o *BatchResponse) GetFinalizingAt() int32`

GetFinalizingAt returns the FinalizingAt field if non-nil, zero value otherwise.

### GetFinalizingAtOk

`func (o *BatchResponse) GetFinalizingAtOk() (*int32, bool)`

GetFinalizingAtOk returns a tuple with the FinalizingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizingAt

`func (o *BatchResponse) SetFinalizingAt(v int32)`

SetFinalizingAt sets FinalizingAt field to given value.

### HasFinalizingAt

`func (o *BatchResponse) HasFinalizingAt() bool`

HasFinalizingAt returns a boolean if a field has been set.

### SetFinalizingAtNil

`func (o *BatchResponse) SetFinalizingAtNil(b bool)`

 SetFinalizingAtNil sets the value for FinalizingAt to be an explicit nil

### UnsetFinalizingAt
`func (o *BatchResponse) UnsetFinalizingAt()`

UnsetFinalizingAt ensures that no value is present for FinalizingAt, not even an explicit nil
### GetInProgressAt

`func (o *BatchResponse) GetInProgressAt() int32`

GetInProgressAt returns the InProgressAt field if non-nil, zero value otherwise.

### GetInProgressAtOk

`func (o *BatchResponse) GetInProgressAtOk() (*int32, bool)`

GetInProgressAtOk returns a tuple with the InProgressAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressAt

`func (o *BatchResponse) SetInProgressAt(v int32)`

SetInProgressAt sets InProgressAt field to given value.

### HasInProgressAt

`func (o *BatchResponse) HasInProgressAt() bool`

HasInProgressAt returns a boolean if a field has been set.

### SetInProgressAtNil

`func (o *BatchResponse) SetInProgressAtNil(b bool)`

 SetInProgressAtNil sets the value for InProgressAt to be an explicit nil

### UnsetInProgressAt
`func (o *BatchResponse) UnsetInProgressAt()`

UnsetInProgressAt ensures that no value is present for InProgressAt, not even an explicit nil
### GetMetadata

`func (o *BatchResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BatchResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BatchResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BatchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BatchResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BatchResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *BatchResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BatchResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BatchResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BatchResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *BatchResponse) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *BatchResponse) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetOutputFileId

`func (o *BatchResponse) GetOutputFileId() string`

GetOutputFileId returns the OutputFileId field if non-nil, zero value otherwise.

### GetOutputFileIdOk

`func (o *BatchResponse) GetOutputFileIdOk() (*string, bool)`

GetOutputFileIdOk returns a tuple with the OutputFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileId

`func (o *BatchResponse) SetOutputFileId(v string)`

SetOutputFileId sets OutputFileId field to given value.

### HasOutputFileId

`func (o *BatchResponse) HasOutputFileId() bool`

HasOutputFileId returns a boolean if a field has been set.

### SetOutputFileIdNil

`func (o *BatchResponse) SetOutputFileIdNil(b bool)`

 SetOutputFileIdNil sets the value for OutputFileId to be an explicit nil

### UnsetOutputFileId
`func (o *BatchResponse) UnsetOutputFileId()`

UnsetOutputFileId ensures that no value is present for OutputFileId, not even an explicit nil
### GetRequestCounts

`func (o *BatchResponse) GetRequestCounts() BATCHBatchRequestCounts`

GetRequestCounts returns the RequestCounts field if non-nil, zero value otherwise.

### GetRequestCountsOk

`func (o *BatchResponse) GetRequestCountsOk() (*BATCHBatchRequestCounts, bool)`

GetRequestCountsOk returns a tuple with the RequestCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCounts

`func (o *BatchResponse) SetRequestCounts(v BATCHBatchRequestCounts)`

SetRequestCounts sets RequestCounts field to given value.

### HasRequestCounts

`func (o *BatchResponse) HasRequestCounts() bool`

HasRequestCounts returns a boolean if a field has been set.

### SetRequestCountsNil

`func (o *BatchResponse) SetRequestCountsNil(b bool)`

 SetRequestCountsNil sets the value for RequestCounts to be an explicit nil

### UnsetRequestCounts
`func (o *BatchResponse) UnsetRequestCounts()`

UnsetRequestCounts ensures that no value is present for RequestCounts, not even an explicit nil
### GetUsage

`func (o *BatchResponse) GetUsage() BATCHBatchUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BatchResponse) GetUsageOk() (*BATCHBatchUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BatchResponse) SetUsage(v BATCHBatchUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BatchResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *BatchResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *BatchResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetProvider

`func (o *BatchResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BatchResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BatchResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


