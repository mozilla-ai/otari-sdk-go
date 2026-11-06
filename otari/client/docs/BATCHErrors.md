# BATCHErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BATCHBatchError**](BATCHBatchError.md) |  | [optional] 
**Object** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 

## Methods

### NewBATCHErrors

`func NewBATCHErrors() *BATCHErrors`

NewBATCHErrors instantiates a new BATCHErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBATCHErrorsWithDefaults

`func NewBATCHErrorsWithDefaults() *BATCHErrors`

NewBATCHErrorsWithDefaults instantiates a new BATCHErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BATCHErrors) GetData() []BATCHBatchError`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BATCHErrors) GetDataOk() (*[]BATCHBatchError, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BATCHErrors) SetData(v []BATCHBatchError)`

SetData sets Data field to given value.

### HasData

`func (o *BATCHErrors) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *BATCHErrors) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *BATCHErrors) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetObject

`func (o *BATCHErrors) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BATCHErrors) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BATCHErrors) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *BATCHErrors) HasObject() bool`

HasObject returns a boolean if a field has been set.

### SetObjectNil

`func (o *BATCHErrors) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *BATCHErrors) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


