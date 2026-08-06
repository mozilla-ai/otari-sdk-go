# ExternalIngestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**Index** | **int32** |  | 
**SourceEventId** | **NullableString** |  | 

## Methods

### NewExternalIngestError

`func NewExternalIngestError(detail string, index int32, sourceEventId NullableString, ) *ExternalIngestError`

NewExternalIngestError instantiates a new ExternalIngestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIngestErrorWithDefaults

`func NewExternalIngestErrorWithDefaults() *ExternalIngestError`

NewExternalIngestErrorWithDefaults instantiates a new ExternalIngestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ExternalIngestError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ExternalIngestError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ExternalIngestError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetIndex

`func (o *ExternalIngestError) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ExternalIngestError) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ExternalIngestError) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetSourceEventId

`func (o *ExternalIngestError) GetSourceEventId() string`

GetSourceEventId returns the SourceEventId field if non-nil, zero value otherwise.

### GetSourceEventIdOk

`func (o *ExternalIngestError) GetSourceEventIdOk() (*string, bool)`

GetSourceEventIdOk returns a tuple with the SourceEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEventId

`func (o *ExternalIngestError) SetSourceEventId(v string)`

SetSourceEventId sets SourceEventId field to given value.


### SetSourceEventIdNil

`func (o *ExternalIngestError) SetSourceEventIdNil(b bool)`

 SetSourceEventIdNil sets the value for SourceEventId to be an explicit nil

### UnsetSourceEventId
`func (o *ExternalIngestError) UnsetSourceEventId()`

UnsetSourceEventId ensures that no value is present for SourceEventId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


