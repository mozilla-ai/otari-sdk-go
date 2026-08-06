# ExternalIngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **int32** |  | [optional] [default to 0]
**Duplicate** | Pointer to **int32** |  | [optional] [default to 0]
**Errors** | Pointer to [**[]ExternalIngestError**](ExternalIngestError.md) |  | [optional] 
**Rejected** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewExternalIngestResult

`func NewExternalIngestResult() *ExternalIngestResult`

NewExternalIngestResult instantiates a new ExternalIngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIngestResultWithDefaults

`func NewExternalIngestResultWithDefaults() *ExternalIngestResult`

NewExternalIngestResultWithDefaults instantiates a new ExternalIngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *ExternalIngestResult) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *ExternalIngestResult) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *ExternalIngestResult) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *ExternalIngestResult) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetDuplicate

`func (o *ExternalIngestResult) GetDuplicate() int32`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *ExternalIngestResult) GetDuplicateOk() (*int32, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *ExternalIngestResult) SetDuplicate(v int32)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *ExternalIngestResult) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetErrors

`func (o *ExternalIngestResult) GetErrors() []ExternalIngestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExternalIngestResult) GetErrorsOk() (*[]ExternalIngestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExternalIngestResult) SetErrors(v []ExternalIngestError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExternalIngestResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetRejected

`func (o *ExternalIngestResult) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *ExternalIngestResult) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *ExternalIngestResult) SetRejected(v int32)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *ExternalIngestResult) HasRejected() bool`

HasRejected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


