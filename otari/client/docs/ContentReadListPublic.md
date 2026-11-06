# ContentReadListPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | More reads follow this page. | 
**Items** | [**[]ContentReadPublic**](ContentReadPublic.md) |  | 

## Methods

### NewContentReadListPublic

`func NewContentReadListPublic(hasMore bool, items []ContentReadPublic, ) *ContentReadListPublic`

NewContentReadListPublic instantiates a new ContentReadListPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentReadListPublicWithDefaults

`func NewContentReadListPublicWithDefaults() *ContentReadListPublic`

NewContentReadListPublicWithDefaults instantiates a new ContentReadListPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ContentReadListPublic) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ContentReadListPublic) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ContentReadListPublic) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetItems

`func (o *ContentReadListPublic) GetItems() []ContentReadPublic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContentReadListPublic) GetItemsOk() (*[]ContentReadPublic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContentReadListPublic) SetItems(v []ContentReadPublic)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


