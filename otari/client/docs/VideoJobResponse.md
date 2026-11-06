# VideoJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | Pointer to **NullableInt32** |  | [optional] 
**Error** | Pointer to [**NullableVideoError**](VideoError.md) |  | [optional] 
**Id** | **string** |  | 
**Model** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "video"]
**Progress** | Pointer to **NullableInt32** |  | [optional] 
**Provider** | **string** |  | 
**Seconds** | Pointer to **NullableInt32** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewVideoJobResponse

`func NewVideoJobResponse(id string, model string, provider string, status string, ) *VideoJobResponse`

NewVideoJobResponse instantiates a new VideoJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoJobResponseWithDefaults

`func NewVideoJobResponseWithDefaults() *VideoJobResponse`

NewVideoJobResponseWithDefaults instantiates a new VideoJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *VideoJobResponse) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *VideoJobResponse) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *VideoJobResponse) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *VideoJobResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *VideoJobResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *VideoJobResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *VideoJobResponse) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoJobResponse) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoJobResponse) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoJobResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *VideoJobResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *VideoJobResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetError

`func (o *VideoJobResponse) GetError() VideoError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VideoJobResponse) GetErrorOk() (*VideoError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VideoJobResponse) SetError(v VideoError)`

SetError sets Error field to given value.

### HasError

`func (o *VideoJobResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *VideoJobResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *VideoJobResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *VideoJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoJobResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *VideoJobResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VideoJobResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VideoJobResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *VideoJobResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VideoJobResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VideoJobResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *VideoJobResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProgress

`func (o *VideoJobResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VideoJobResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VideoJobResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VideoJobResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *VideoJobResponse) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *VideoJobResponse) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetProvider

`func (o *VideoJobResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VideoJobResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VideoJobResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSeconds

`func (o *VideoJobResponse) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *VideoJobResponse) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *VideoJobResponse) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *VideoJobResponse) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### SetSecondsNil

`func (o *VideoJobResponse) SetSecondsNil(b bool)`

 SetSecondsNil sets the value for Seconds to be an explicit nil

### UnsetSeconds
`func (o *VideoJobResponse) UnsetSeconds()`

UnsetSeconds ensures that no value is present for Seconds, not even an explicit nil
### GetSize

`func (o *VideoJobResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoJobResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoJobResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoJobResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VideoJobResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VideoJobResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *VideoJobResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoJobResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoJobResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


