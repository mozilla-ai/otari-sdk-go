# VideoJobList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VideoJobSummary**](VideoJobSummary.md) |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]

## Methods

### NewVideoJobList

`func NewVideoJobList(data []VideoJobSummary, ) *VideoJobList`

NewVideoJobList instantiates a new VideoJobList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoJobListWithDefaults

`func NewVideoJobListWithDefaults() *VideoJobList`

NewVideoJobListWithDefaults instantiates a new VideoJobList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VideoJobList) GetData() []VideoJobSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideoJobList) GetDataOk() (*[]VideoJobSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideoJobList) SetData(v []VideoJobSummary)`

SetData sets Data field to given value.


### GetObject

`func (o *VideoJobList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VideoJobList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VideoJobList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *VideoJobList) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


