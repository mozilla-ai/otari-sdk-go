# ModelListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ModelObject**](ModelObject.md) |  | 
**Object** | Pointer to **string** |  | [optional] [default to "list"]

## Methods

### NewModelListResponse

`func NewModelListResponse(data []ModelObject, ) *ModelListResponse`

NewModelListResponse instantiates a new ModelListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelListResponseWithDefaults

`func NewModelListResponseWithDefaults() *ModelListResponse`

NewModelListResponseWithDefaults instantiates a new ModelListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModelListResponse) GetData() []ModelObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelListResponse) GetDataOk() (*[]ModelObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelListResponse) SetData(v []ModelObject)`

SetData sets Data field to given value.


### GetObject

`func (o *ModelListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelListResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelListResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


