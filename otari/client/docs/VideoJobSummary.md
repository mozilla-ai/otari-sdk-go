# VideoJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billed** | **bool** | Whether a retrieval has found the job completed and billed it. | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Model** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "video"]
**Provider** | **string** |  | 
**Seconds** | **NullableInt32** | The length the request asked for, in seconds, when it named one. | 

## Methods

### NewVideoJobSummary

`func NewVideoJobSummary(billed bool, createdAt time.Time, id string, model string, provider string, seconds NullableInt32, ) *VideoJobSummary`

NewVideoJobSummary instantiates a new VideoJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoJobSummaryWithDefaults

`func NewVideoJobSummaryWithDefaults() *VideoJobSummary`

NewVideoJobSummaryWithDefaults instantiates a new VideoJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilled

`func (o *VideoJobSummary) GetBilled() bool`

GetBilled returns the Billed field if non-nil, zero value otherwise.

### GetBilledOk

`func (o *VideoJobSummary) GetBilledOk() (*bool, bool)`

GetBilledOk returns a tuple with the Billed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilled

`func (o *VideoJobSummary) SetBilled(v bool)`

SetBilled sets Billed field to given value.


### GetCreatedAt

`func (o *VideoJobSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoJobSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoJobSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *VideoJobSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoJobSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoJobSummary) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *VideoJobSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VideoJobSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VideoJobSummary) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *VideoJobSummary) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VideoJobSummary) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VideoJobSummary) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *VideoJobSummary) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProvider

`func (o *VideoJobSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VideoJobSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VideoJobSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSeconds

`func (o *VideoJobSummary) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *VideoJobSummary) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *VideoJobSummary) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.


### SetSecondsNil

`func (o *VideoJobSummary) SetSecondsNil(b bool)`

 SetSecondsNil sets the value for Seconds to be an explicit nil

### UnsetSeconds
`func (o *VideoJobSummary) UnsetSeconds()`

UnsetSeconds ensures that no value is present for Seconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


